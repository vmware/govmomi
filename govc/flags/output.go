/*
Copyright (c) 2014-2016 VMware, Inc. All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package flags

import (
	"context"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"reflect"
	"sync"
	"time"

	"github.com/kr/pretty"
	"github.com/vmware/govmomi/task"
	"github.com/vmware/govmomi/vim25/progress"
	"github.com/vmware/govmomi/vim25/soap"
	"github.com/vmware/govmomi/vim25/xml"
)

type OutputWriter interface {
	Write(io.Writer) error
}

type OutputFlag struct {
	common

	JSON bool
	XML  bool
	TTY  bool
	Dump bool
	Out  io.Writer

	formatError  bool
	formatIndent bool
}

var outputFlagKey = flagKey("output")

func NewOutputFlag(ctx context.Context) (*OutputFlag, context.Context) {
	if v := ctx.Value(outputFlagKey); v != nil {
		return v.(*OutputFlag), ctx
	}

	v := &OutputFlag{Out: os.Stdout}
	ctx = context.WithValue(ctx, outputFlagKey, v)
	return v, ctx
}

func (flag *OutputFlag) Register(ctx context.Context, f *flag.FlagSet) {
	flag.RegisterOnce(func() {
		f.BoolVar(&flag.JSON, "json", false, "Enable JSON output")
		f.BoolVar(&flag.XML, "xml", false, "Enable XML output")
		f.BoolVar(&flag.Dump, "dump", false, "Enable Go output")
		// Avoid adding more flags for now..
		flag.formatIndent = os.Getenv("GOVC_INDENT") != "false"      // Default to indented output
		flag.formatError = os.Getenv("GOVC_FORMAT_ERROR") != "false" // Default to formatted errors
	})
}

func (flag *OutputFlag) Process(ctx context.Context) error {
	return flag.ProcessOnce(func() error {
		if !flag.All() {
			// Assume we have a tty if not outputting JSON
			flag.TTY = true
		}

		return nil
	})
}

// Log outputs the specified string, prefixed with the current time.
// A newline is not automatically added. If the specified string
// starts with a '\r', the current line is cleared first.
func (flag *OutputFlag) Log(s string) (int, error) {
	if len(s) > 0 && s[0] == '\r' {
		flag.Write([]byte{'\r', 033, '[', 'K'})
		s = s[1:]
	}

	return flag.WriteString(time.Now().Format("[02-01-06 15:04:05] ") + s)
}

func (flag *OutputFlag) Write(b []byte) (int, error) {
	if !flag.TTY {
		return 0, nil
	}

	n, err := os.Stdout.Write(b)
	os.Stdout.Sync()
	return n, err
}

func (flag *OutputFlag) WriteString(s string) (int, error) {
	return flag.Write([]byte(s))
}

func (flag *OutputFlag) All() bool {
	return flag.JSON || flag.XML || flag.Dump
}

func dumpValue(val interface{}) interface{} {
	type dumper interface {
		Dump() interface{}
	}

	if d, ok := val.(dumper); ok {
		return d.Dump()
	}

	rval := reflect.ValueOf(val)
	if rval.Type().Kind() != reflect.Ptr {
		return val
	}

	rval = rval.Elem()
	if rval.Type().Kind() == reflect.Struct {
		f := rval.Field(0)
		if f.Type().Kind() == reflect.Slice {
			// common case for the various 'type infoResult'
			if f.Len() == 1 {
				return f.Index(0).Interface()
			}
			return f.Interface()
		}

		if rval.NumField() == 1 && rval.Type().Field(0).Anonymous {
			// common case where govc type wraps govmomi type to implement OutputWriter
			return f.Interface()
		}
	}

	return val
}

func (flag *OutputFlag) WriteResult(result OutputWriter) error {
	var err error

	switch {
	case flag.Dump:
		format := "%#v\n"
		if flag.formatIndent {
			format = "%# v\n"
		}
		_, err = pretty.Fprintf(flag.Out, format, dumpValue(result))
	case flag.JSON:
		e := json.NewEncoder(flag.Out)
		if flag.formatIndent {
			e.SetIndent("", "  ")
		}
		err = e.Encode(result)
	case flag.XML:
		e := xml.NewEncoder(flag.Out)
		if flag.formatIndent {
			e.Indent("", "  ")
		}
		err = e.Encode(dumpValue(result))
		if err == nil {
			fmt.Fprintln(flag.Out)
		}
	default:
		err = result.Write(flag.Out)
	}

	return err
}

func (flag *OutputFlag) WriteError(err error) bool {
	if flag.formatError {
		flag.Out = os.Stderr
		return flag.WriteResult(&errorOutput{err}) == nil
	}
	return false
}

type errorOutput struct {
	error
}

func (e errorOutput) Write(w io.Writer) error {
	_, ferr := fmt.Fprintf(w, "%s: %s\n", os.Args[0], e.error)
	return ferr
}

func (e errorOutput) Dump() interface{} {
	if f, ok := e.error.(task.Error); ok {
		return f.LocalizedMethodFault
	}
	if soap.IsSoapFault(e.error) {
		return soap.ToSoapFault(e.error)
	}
	if soap.IsVimFault(e.error) {
		return soap.ToVimFault(e.error)
	}
	return e
}

func (e errorOutput) canEncode() bool {
	switch e.error.(type) {
	case task.Error:
		return true
	}
	return soap.IsSoapFault(e.error) || soap.IsVimFault(e.error)
}

// cannotEncode causes cli.Run to output err.Error() as it would without an error format specified
var cannotEncode = errors.New("cannot encode error")

func (e errorOutput) MarshalJSON() ([]byte, error) {
	_, ok := e.error.(json.Marshaler)
	if ok || e.canEncode() {
		return json.Marshal(e.error)
	}
	return nil, cannotEncode
}

func (e errorOutput) MarshalXML(encoder *xml.Encoder, start xml.StartElement) error {
	_, ok := e.error.(xml.Marshaler)
	if ok || e.canEncode() {
		return encoder.Encode(e.error)
	}
	return cannotEncode
}

type progressLogger struct {
	flag   *OutputFlag
	prefix string

	wg sync.WaitGroup

	sink chan chan progress.Report
	done chan struct{}
}

func newProgressLogger(flag *OutputFlag, prefix string) *progressLogger {
	p := &progressLogger{
		flag:   flag,
		prefix: prefix,

		sink: make(chan chan progress.Report),
		done: make(chan struct{}),
	}

	p.wg.Add(1)

	go p.loopA()

	return p
}

// loopA runs before Sink() has been called.
func (p *progressLogger) loopA() {
	var err error

	defer p.wg.Done()

	tick := time.NewTicker(100 * time.Millisecond)
	defer tick.Stop()

	called := false

	for stop := false; !stop; {
		select {
		case ch := <-p.sink:
			err = p.loopB(tick, ch)
			stop = true
			called = true
		case <-p.done:
			stop = true
		case <-tick.C:
			line := fmt.Sprintf("\r%s", p.prefix)
			p.flag.Log(line)
		}
	}

	if err != nil && err != io.EOF {
		p.flag.Log(fmt.Sprintf("\r%sError: %s\n", p.prefix, err))
	} else if called {
		p.flag.Log(fmt.Sprintf("\r%sOK\n", p.prefix))
	}
}

// loopA runs after Sink() has been called.
func (p *progressLogger) loopB(tick *time.Ticker, ch <-chan progress.Report) error {
	var r progress.Report
	var ok bool
	var err error

	for ok = true; ok; {
		select {
		case r, ok = <-ch:
			if !ok {
				break
			}
			err = r.Error()
		case <-tick.C:
			line := fmt.Sprintf("\r%s", p.prefix)
			if r != nil {
				line += fmt.Sprintf("(%.0f%%", r.Percentage())
				detail := r.Detail()
				if detail != "" {
					line += fmt.Sprintf(", %s", detail)
				}
				line += ")"
			}
			p.flag.Log(line)
		}
	}

	return err
}

func (p *progressLogger) Sink() chan<- progress.Report {
	ch := make(chan progress.Report)
	p.sink <- ch
	return ch
}

func (p *progressLogger) Wait() {
	close(p.done)
	p.wg.Wait()
}

func (flag *OutputFlag) ProgressLogger(prefix string) *progressLogger {
	return newProgressLogger(flag, prefix)
}
