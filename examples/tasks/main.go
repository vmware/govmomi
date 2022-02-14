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
