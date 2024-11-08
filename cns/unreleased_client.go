/*
Copyright (c) 2024-2024 VMware, Inc. All Rights Reserved.

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

package cns

import (
	"context"

	"github.com/vmware/govmomi/cns/methods"
	cnstypes "github.com/vmware/govmomi/cns/types"
	"github.com/vmware/govmomi/object"
)

// UpdateVolumeCrypto updates container volumes with given crypto specifications.
func (c *Client) UpdateVolumeCrypto(ctx context.Context, updateSpecList []cnstypes.CnsVolumeCryptoUpdateSpec) (*object.Task, error) {
	req := cnstypes.CnsUpdateVolumeCrypto{
		This:        CnsVolumeManagerInstance,
		UpdateSpecs: updateSpecList,
	}
	res, err := methods.CnsUpdateVolumeCrypto(ctx, c, &req)
	if err != nil {
		return nil, err
	}
	return object.NewTask(c.vim25Client, res.Returnval), nil
}
