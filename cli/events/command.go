// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package events

import (
	"context"
	"flag"
	"fmt"
	"io"
	"os"
	"reflect"
	"strings"
	"time"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/event"
	"github.com/vmware/govmomi/vim25/types"
)

type events struct {
	*flags.DatacenterFlag

	Max   int32
	Tail  bool
	Force bool
	Long  bool
	Kind  kinds
}

type kinds []string

func (e *kinds) String() string {
	return fmt.Sprint(*e)
}

func (e *kinds) Set(value string) error {
	*e = append(*e, value)
	return nil
}

func init() {
	// initialize with the maximum allowed objects set
	cli.Register("events", &events{})
}

func (cmd *events) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.DatacenterFlag, ctx = flags.NewDatacenterFlag(ctx)
	cmd.DatacenterFlag.Register(ctx, f)

	cmd.Max = 25 // default
	f.Var(flags.NewInt32(&cmd.Max), "n", "Output the last N events")
	f.BoolVar(&cmd.Tail, "f", false, "Follow event stream")
	f.BoolVar(&cmd.Force, "force", false, "Disable number objects to monitor limit")
	f.BoolVar(&cmd.Long, "l", false, "Long listing format")
	f.Var(&cmd.Kind, "type", "Include only the specified event types")
}

func (cmd *events) Description() string {
	return `Display events.

Examples:
  govc events vm/my-vm1 vm/my-vm2
  govc events /dc1/vm/* /dc2/vm/*
  govc events -type VmPoweredOffEvent -type VmPoweredOnEvent
  govc ls -t HostSystem host/* | xargs govc events | grep -i vsan`
}

func (cmd *events) Usage() string {
	return "[PATH]..."
}

func (cmd *events) printEvents(ctx context.Context, obj *types.ManagedObjectReference, page []types.BaseEvent, m *event.Manager) error {
	event.Sort(page)
	source := ""
	if obj != nil {
		source = obj.String()
		if !cmd.JSON {
			// print the object reference
			fmt.Fprintf(os.Stdout, "\n==> %s <==\n", source)
		}
	}
	for _, e := range page {
		cat, err := m.EventCategory(ctx, e)
		if err != nil {
			return err
		}

		event := e.GetEvent()
		r := &record{
			Object:      source,
			CreatedTime: event.CreatedTime,
			Category:    cat,
			Message:     strings.TrimSpace(event.FullFormattedMessage),
			event:       e,
		}

		if cmd.Long {
			r.Type = reflect.TypeOf(e).Elem().Name()
		}

		if cmd.Long {
			r.Key = event.Key
		}

		switch x := e.(type) {
		case *types.TaskEvent:
			// some tasks won't have this information, so just use the event message
			if x.Info.Entity != nil {
				r.Message = fmt.Sprintf("%s (target=%s %s)", r.Message, x.Info.Entity.Type, x.Info.EntityName)
			}
		case *types.EventEx:
			if r.Message == "" {
				r.Message = x.Message
			}
			if x.ObjectId != "" {
				r.Message = fmt.Sprintf("%s (%s)", r.Message, x.ObjectId)
			}
			if cmd.Long {
				r.Type = x.EventTypeId
			}
		}

		if err = cmd.WriteResult(r); err != nil {
			return err
		}
	}
	return nil
}

type record struct {
	Object      string    `json:"object,omitempty"`
	Type        string    `json:"type,omitempty"`
	CreatedTime time.Time `json:"createdTime"`
	Category    string    `json:"category"`
	Message     string    `json:"message"`
	Key         int32     `json:"key,omitempty"`

	event types.BaseEvent
}

// Dump the raw Event rather than the record struct.
func (r *record) Dump() any {
	return r.event
}

func (r *record) Write(w io.Writer) error {
	when := r.CreatedTime.Local().Format(time.ANSIC)
	var kind, key string
	if r.Type != "" {
		kind = fmt.Sprintf(" [%s]", r.Type)
	}
	if r.Key != 0 {
		key = fmt.Sprintf(" [%d]", r.Key)
	}
	_, err := fmt.Fprintf(w, "[%s] [%s]%s%s %s\n", when, r.Category, key, kind, r.Message)
	return err
}

func (cmd *events) Run(ctx context.Context, f *flag.FlagSet) error {
	c, err := cmd.Client()
	if err != nil {
		return err
	}

	objs, err := cmd.ManagedObjects(ctx, f.Args())
	if err != nil {
		return err
	}

	m := event.NewManager(c)

	return cmd.WithCancel(ctx, func(wctx context.Context) error {
		return m.Events(wctx, objs, cmd.Max, cmd.Tail, cmd.Force,
			func(obj types.ManagedObjectReference, ee []types.BaseEvent) error {
				var o *types.ManagedObjectReference
				if len(objs) > 1 {
					o = &obj
				}

				return cmd.printEvents(ctx, o, ee, m)
			}, cmd.Kind...)
	})
}
