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

package importx

import (
	"fmt"
	"path"
	"strings"
)

type importable string

func (i importable) Ext() string {
	return strings.ToLower(path.Ext(string(i)))
}

func (i importable) Base() string {
	return path.Base(string(i))
}

func (i importable) Dir() string {
	return path.Dir(string(i))
}

func (i importable) BaseClean() string {
	b := path.Base(string(i))
	e := path.Ext(string(i))
	return b[:len(b)-len(e)]
}

func (i importable) RemoteVMDK() string {
	bc := i.BaseClean()
	return fmt.Sprintf("%s-vmdk/%s.vmdk", bc, bc)
}

func (i importable) RemoteDst() string {
	bc := i.BaseClean()
	return fmt.Sprintf("%s/%s.vmdk", bc, bc)
}
