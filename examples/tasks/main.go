// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"context"
	"flag"
	"time"

	"github.com/dougm/pretty"

	"github.com/vmware/govmomi/task"

	"github.com/vmware/govmomi/examples"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/types"
)

func main() {
	// example use against vCenter with optional time filters:
	// go run main.go -url $GOVMOMI_URL -insecure $GOVMOMI_INSECURE -b 8h -f
	begin := flag.Duration("b", 10*time.Minute, "Begin time (filtered by started time)") // default BeginTime is 10min ago
	end := flag.Duration("e", 0, "End time (filtered by started time)")
	follow := flag.Bool("f", false, "Follow task stream")

	examples.Run(func(ctx context.Context, c *vim25.Client) error {
		m := task.NewManager(c)

		ref := c.ServiceContent.RootFolder

		now, err := methods.GetCurrentTime(ctx, c) // vCenter server time (UTC)
		if err != nil {
			return err
		}

		filter := types.TaskFilterSpec{
			Entity: &types.TaskFilterSpecByEntity{
				Entity:    ref,
				Recursion: types.TaskFilterSpecRecursionOptionAll,
			},
			Time: &types.TaskFilterSpecByTime{
				TimeType:  types.TaskFilterSpecTimeOptionStartedTime,
				BeginTime: types.NewTime(now.Add(*begin * -1)),
			},
		}

		if *end != 0 {
			filter.Time.EndTime = types.NewTime(now.Add(*end * -1))
		}

		collector, err := m.CreateCollectorForTasks(ctx, filter)
		if err != nil {
			return err
		}

		defer collector.Destroy(ctx)

		for {
			tasks, err := collector.ReadNextTasks(ctx, 10)
			if err != nil {
				return err
			}

			if len(tasks) == 0 {
				if *follow {
					time.Sleep(time.Second)
					continue
				}
				break
			}

			for i := range tasks {
				pretty.Println(tasks[i])
			}
		}

		return nil
	})
}
