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
	"bufio"
	"context"
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"strings"
	"sync"
	"text/tabwriter"
	"time"

	"github.com/kr/pretty"

	"github.com/vmware/govmomi/vim25/debug"
	"github.com/vmware/govmomi/vim25/soap"
	"github.com/vmware/govmomi/vim25/types"
)

type cmdFormat struct {
	path string
	err  error
	args []string
}

func (c *cmdFormat) lookPath(file string, args ...string) {
	c.args = args
	c.path, c.err = exec.LookPath(file)
}

func (c *cmdFormat) cmd() (*exec.Cmd, error) {
	if c.err != nil {
		return nil, c.err
	}
	return exec.Command(c.path, c.args...), nil
}

type DebugFlag struct {
	common

	enable  bool
	trace   bool
	verbose bool
	dump    bool
	xml     cmdFormat
	json    cmdFormat
}

var debugFlagKey = flagKey("debug")

func NewDebugFlag(ctx context.Context) (*DebugFlag, context.Context) {
	if v := ctx.Value(debugFlagKey); v != nil {
		return v.(*DebugFlag), ctx
	}

	v := &DebugFlag{}
	ctx = context.WithValue(ctx, debugFlagKey, v)
	return v, ctx
}

func (flag *DebugFlag) Verbose() bool {
	return flag.verbose
}

func (flag *DebugFlag) Register(ctx context.Context, f *flag.FlagSet) {
	flag.RegisterOnce(func() {
		env := "GOVC_DEBUG"
		enable := false
		switch env := strings.ToLower(os.Getenv(env)); env {
		case "1", "true":
			enable = true
		}

		usage := fmt.Sprintf("Store debug logs [%s]", env)
		f.BoolVar(&flag.enable, "debug", enable, usage)
		f.BoolVar(&flag.trace, "trace", false, "Write SOAP/REST traffic to stderr")
		f.BoolVar(&flag.verbose, "verbose", false, "Write request/response data to stderr")
	})
}

type cmdFormatCloser struct {
	rc  io.Closer
	in  io.Closer
	cmd *exec.Cmd
	wg  *sync.WaitGroup
}

func (c *cmdFormatCloser) Close() error {
	_ = c.rc.Close()
	_ = c.in.Close()
	c.wg.Wait()
	return c.cmd.Wait()
}

func (flag *DebugFlag) newFormatReader(rc io.ReadCloser, w io.Writer, ext string) (io.ReadCloser, error) {
	var err error
	var cmd *exec.Cmd

	switch ext {
	case "json":
		cmd, err = flag.json.cmd()
		if err != nil {
			return nil, err
		}
	case "xml":
		cmd, err = flag.xml.cmd()
		if err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("unsupported type %s", ext)
	}

	cmd.Stderr = os.Stderr

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return nil, err
	}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}

	err = cmd.Start()
	if err != nil {
		return nil, err
	}

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		_, _ = io.Copy(w, stdout)
		wg.Done()
	}()

	return debug.ReadCloser{
		Reader: io.TeeReader(rc, stdin),
		Closer: &cmdFormatCloser{rc, stdin, cmd, &wg},
	}, nil
}

func (flag *DebugFlag) debugTrace(rc io.ReadCloser, w io.Writer, ext string) io.ReadCloser {
	fr, err := flag.newFormatReader(rc, w, ext)
	if err != nil {
		return debug.NewTeeReader(rc, w)
	}
	return fr
}

func (flag *DebugFlag) Process(ctx context.Context) error {
	// Base path for storing debug logs.
	r := os.Getenv("GOVC_DEBUG_PATH")

	if flag.trace {
		if flag.verbose {
			flag.dump = true // output req/res as Go code
			return nil
		}
		r = "-"
		flag.enable = true
		if os.Getenv("GOVC_DEBUG_FORMAT") != "false" {
			debugXML := os.Getenv("GOVC_DEBUG_XML")
			if debugXML == "" {
				debugXML = "xmlstarlet"
			}
			flag.xml.lookPath(debugXML, "fo")

			debugJSON := os.Getenv("GOVC_DEBUG_JSON")
			if debugJSON == "" {
				debugJSON = "jq"
			}
			flag.json.lookPath(debugJSON, ".")

			soap.Trace = flag.debugTrace
		}
	}

	if !flag.enable {
		return nil
	}

	return flag.ProcessOnce(func() error {
		switch r {
		case "-":
			debug.SetProvider(&debug.LogProvider{})
			return nil
		case "":
			r = home
		}
		r = filepath.Join(r, "debug")

		// Path for this particular run.
		run := os.Getenv("GOVC_DEBUG_PATH_RUN")
		if run == "" {
			now := time.Now().Format("2006-01-02T15-04-05.999999999")
			r = filepath.Join(r, now)
		} else {
			// reuse the same path
			r = filepath.Join(r, run)
			_ = os.RemoveAll(r)
		}

		err := os.MkdirAll(r, 0700)
		if err != nil {
			return err
		}

		p := debug.FileProvider{
			Path: r,
		}

		debug.SetProvider(&p)
		return nil
	})
}

type dump struct {
	roundTripper soap.RoundTripper
}

func (d *dump) RoundTrip(ctx context.Context, req, res soap.HasFault) error {
	vreq := reflect.ValueOf(req).Elem().FieldByName("Req").Elem()

	pretty.Fprintf(os.Stderr, "%# v\n", vreq.Interface())

	err := d.roundTripper.RoundTrip(ctx, req, res)
	if err != nil {
		if fault := res.Fault(); fault != nil {
			pretty.Fprintf(os.Stderr, "%# v\n", fault)
		}
		return err
	}

	vres := reflect.ValueOf(res).Elem().FieldByName("Res").Elem()
	pretty.Fprintf(os.Stderr, "%# v\n", vres.Interface())

	return nil
}

type verbose struct {
	roundTripper soap.RoundTripper
}

func (*verbose) mor(ref types.ManagedObjectReference) string {
	if strings.HasPrefix(ref.Value, "session") {
		ref.Value = "session[...]"
		return ref.String()
	}
	return ref.Value
}

func (*verbose) str(val reflect.Value) string {
	if !val.IsValid() {
		return ""
	}

	switch val.Kind() {
	case reflect.Ptr, reflect.Interface:
		if val.IsNil() {
			return "nil"
		}
	}

	p := ""

	switch pval := val.Interface().(type) {
	case fmt.Stringer:
		p = pval.String()
	case string:
		if len(pval) > 45 {
			pval = pval[:42] + "..."
		}
		p = fmt.Sprintf("%s", pval)
	case []string:
		p = fmt.Sprintf("%v", pval)
	case []types.ManagedObjectReference:
		refs := make([]string, len(pval))
		for i := range pval {
			refs[i] = pval[i].Value
		}
		p = fmt.Sprintf("%v", refs)
	default:
		return ""
	}

	return p
}

func (v *verbose) value(val types.AnyType) string {
	rval := reflect.ValueOf(val)
	if rval.Kind() == reflect.Ptr && !rval.IsNil() {
		rval = rval.Elem()
	}
	if rval.Kind() == reflect.Struct {
		if strings.HasPrefix(rval.Type().Name(), "ArrayOf") {
			rval = rval.Field(0)
			val = rval.Interface()
		}
	}
	s := v.str(rval)
	if s != "" {
		return s
	}
	return v.prettyPrint(val)
}

func (v *verbose) propertyValue(obj types.ManagedObjectReference, name string, pval types.AnyType) string {
	val := v.value(pval)
	if obj.Type != "Task" && !strings.HasPrefix(obj.Value, "session") {
		if len(val) > 512 {
			val = fmt.Sprintf("`govc object.collect -dump %s %s`", obj, name)
		}
	}
	return fmt.Sprintf("%s\t%s:\t%s", v.mor(obj), name, val)
}

func (v *verbose) missingSet(o types.ManagedObjectReference, m []types.MissingProperty) []string {
	var s []string
	for _, p := range m {
		s = append(s, fmt.Sprintf("%s\t%s:\t%s", v.mor(o), p.Path, v.prettyPrint(p.Fault.Fault)))
	}
	return s
}

func (v *verbose) updateSet(u *types.UpdateSet) []string {
	var s []string
	for _, f := range u.FilterSet {
		for _, o := range f.ObjectSet {
			for _, c := range o.ChangeSet {
				s = append(s, v.propertyValue(o.Obj, c.Name, c.Val))
			}
			s = append(s, v.missingSet(o.Obj, o.MissingSet)...)
		}
	}
	return s
}

func (v *verbose) objectContent(content []types.ObjectContent) []string {
	var s []string
	for _, o := range content {
		for _, p := range o.PropSet {
			s = append(s, v.propertyValue(o.Obj, p.Name, p.Val))
		}
		s = append(s, v.missingSet(o.Obj, o.MissingSet)...)
	}
	return s
}

func (v *verbose) prettyPrint(val interface{}) string {
	p := pretty.Sprintf("%# v\n", val)
	var res []string
	scanner := bufio.NewScanner(strings.NewReader(p))
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "nil,") || strings.Contains(line, "(nil),") {
			continue // nil pointer field
		}
		if strings.Contains(line, `"",`) {
			continue // empty string field
		}
		if strings.Contains(line, `{},`) {
			continue // empty embedded struct
		}
		if strings.Contains(line, "[context]") {
			continue // noisy base64 encoded backtrace
		}
		res = append(res, line)
	}
	return strings.Join(res, "\n")
}

func (v *verbose) table(vals []string) {
	tw := tabwriter.NewWriter(os.Stderr, 2, 0, 1, ' ', 0)
	for _, val := range vals {
		fmt.Fprintf(tw, "...%s\n", val)
	}
	tw.Flush()
}

func (v *verbose) RoundTrip(ctx context.Context, req, res soap.HasFault) error {
	vreq := reflect.ValueOf(req).Elem().FieldByName("Req").Elem()
	param := []string{""}
	switch f := vreq.Field(0).Interface().(type) {
	case types.ManagedObjectReference:
		param[0] = v.mor(f)
	default:
		param[0] = fmt.Sprintf("%v", f)
	}

	for i := 1; i < vreq.NumField(); i++ {
		val := vreq.Field(i)

		if val.Kind() == reflect.Interface {
			val = val.Elem()
		}

		p := v.str(val)
		if p == "" {
			switch val.Kind() {
			case reflect.Ptr, reflect.Slice, reflect.Struct:
				p = val.Type().String()
			default:
				p = fmt.Sprintf("%v", val.Interface())
			}
		}

		param = append(param, p)
	}

	fmt.Fprintf(os.Stderr, "%s(%s)...\n", vreq.Type().Name(), strings.Join(param, ", "))

	err := v.roundTripper.RoundTrip(ctx, req, res)
	if err != nil {
		if fault := res.Fault(); fault != nil {
			fmt.Fprintln(os.Stderr, v.prettyPrint(fault))
		} else {
			fmt.Fprintf(os.Stderr, "...%s\n", err)
		}
		return err
	}

	vres := reflect.ValueOf(res).Elem().FieldByName("Res").Elem()
	ret := vres.FieldByName("Returnval")
	var s interface{} = "void"

	if ret.IsValid() {
		switch x := ret.Interface().(type) {
		case types.ManagedObjectReference:
			s = v.mor(x)
		case *types.UpdateSet:
			s = v.updateSet(x)
		case []types.ObjectContent:
			s = v.objectContent(x)
		case fmt.Stringer:
			s = x.String()
		default:
			s = v.value(x)
		}
	}
	if vals, ok := s.([]string); ok {
		v.table(vals)
	} else {
		fmt.Fprintf(os.Stderr, "...%s\n", s)
	}
	fmt.Fprintln(os.Stderr)

	return err
}
