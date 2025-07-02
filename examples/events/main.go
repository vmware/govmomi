// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"context"
	"flag"
	"fmt"
	"reflect"
	"time"

	"github.com/vmware/govmomi/event"
	"github.com/vmware/govmomi/examples"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/types"
)

func main() {
	// example use against simulator: go run main.go -b 8h -f
	// example use against vCenter with optional event filters:
	// go run main.go -url $GOVMOMI_URL -insecure $GOVMOMI_INSECURE -b 8h -f VmEvent UserLoginSessionEvent
	begin := flag.Duration("b", 10*time.Minute, "Begin time") // default BeginTime is 10min ago
	end := flag.Duration("e", 0, "End time")
	follow := flag.Bool("f", false, "Follow event stream")

	examples.Run(func(ctx context.Context, c *vim25.Client) error {
		m := event.NewManager(c)

		ref := c.ServiceContent.RootFolder

		now, err := methods.GetCurrentTime(ctx, c) // vCenter server time (UTC)
		if err != nil {
			return err
		}

		filter := types.EventFilterSpec{
			EventTypeId: flag.Args(), // e.g. VmEvent
			Entity: &types.EventFilterSpecByEntity{
				Entity:    ref,
				Recursion: types.EventFilterSpecRecursionOptionAll,
			},
			Time: &types.EventFilterSpecByTime{
				BeginTime: types.NewTime(now.Add(*begin * -1)),
			},
		}
		if *end != 0 {
			filter.Time.EndTime = types.NewTime(now.Add(*end * -1))
		}

		collector, err := m.CreateCollectorForEvents(ctx, filter)
		if err != nil {
			return err
		}

		defer collector.Destroy(ctx)

		for {
			events, err := collector.ReadNextEvents(ctx, 100)
			if err != nil {
				return err
			}

			if len(events) == 0 {
				if *follow {
					time.Sleep(time.Second)
					continue
				}
				break
			}

			for i := range events {
				event := events[i].GetEvent()
				kind := reflect.TypeOf(events[i]).Elem().Name()
				fmt.Printf("%d [%s] [%s] %s\n", event.Key, event.CreatedTime.Format(time.ANSIC), kind, event.FullFormattedMessage)
			}
		}

		return nil
	})
}
