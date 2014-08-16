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

package cli

import (
	"fmt"
	"path/filepath"
	"reflect"
)

var commands = map[string]Command{}

func name(c Command) string {
	t := reflect.TypeOf(c).Elem()
	base := filepath.Base(t.PkgPath())
	if base == t.Name() {
		return t.Name()
	}
	return fmt.Sprintf("%s.%s", base, t.Name())
}

func Register(c Command) {
	commands[name(c)] = c
}
