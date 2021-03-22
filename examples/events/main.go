/*
Copyright (c) 2020 VMware, Inc. All Rights Reserved.

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
