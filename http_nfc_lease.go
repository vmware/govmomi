/*
Copyright (c) 2014 VMware, Inc. All Rights Reserved.

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

package govmomi

import (
	"errors"
	"fmt"

	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"
)

func (o HttpNfcLease) Wait() (*types.HttpNfcLeaseInfo, error) {
	var lease mo.HttpNfcLease

	err := o.c.WaitForProperties(o.Reference(), []string{"state", "info", "error"}, func(pc []types.PropertyChange) bool {
		done := false

		for _, c := range pc {
			if c.Val == nil {
				continue
			}

			switch c.Name {
			case "error":
				val := c.Val.(types.LocalizedMethodFault)
				lease.Error = &val
				done = true
			case "info":
				val := c.Val.(types.HttpNfcLeaseInfo)
				lease.Info = &val
			case "state":
				lease.State = c.Val.(types.HttpNfcLeaseState)
				if lease.State != types.HttpNfcLeaseStateInitializing {
					done = true
				}
			}
		}

		return done
	})

	if err != nil {
		return nil, err
	}

	if lease.State == types.HttpNfcLeaseStateReady {
		return lease.Info, nil
	}

	if lease.Error != nil {
		return nil, errors.New(lease.Error.LocalizedMessage)
	}

	return nil, fmt.Errorf("unexpected nfc lease state: %s", lease.State)
}
