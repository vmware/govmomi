// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/dougm/pretty"

	"github.com/vmware/govmomi/examples"
	"github.com/vmware/govmomi/property"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"
)

func main() {
	var limit int
	flag.IntVar(&limit, "limit", 10, "maximum number of alarms to retrieve")

	examples.Run(func(ctx context.Context, c *vim25.Client) error {
		pc := property.DefaultCollector(c)
		am := c.ServiceContent.AlarmManager

		fmt.Println("retrieving all alarms")
		alarms, err := methods.GetAlarm(ctx, c, &types.GetAlarm{
			This:   *am,
			Entity: nil, // if not set, alarms are returned for all visible entities
		})
		if err != nil {
			return fmt.Errorf("could not get alarms: %w", err)
		}

		counter := 0
		for _, a := range alarms.Returnval {
			counter++
			fmt.Printf("retrieving details for alarm %q\n", a.String())

			var info mo.Alarm
			if err = pc.RetrieveOne(ctx, a, nil, &info); err != nil {
				return fmt.Errorf("retrieve alarm info: %w", err)
			}

			_, err = pretty.Println(info)
			if err != nil {
				return fmt.Errorf("print alarm: %w", err)
			}

			if counter == limit {
				fmt.Printf("reached maximum number of alarms to read (limit=%d)\n", limit)
				break
			}
		}

		return nil
	})
}
