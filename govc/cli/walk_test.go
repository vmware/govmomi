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
	"reflect"
	"testing"
)

type wC struct {
	C string
}

func (w *wC) Set() {
	w.C = "C"
}

type wB struct {
	C *wC

	B string
}

func (w *wB) Set() {
	w.B = "B"
}

type wA struct {
	B *wB
	C *wC

	A string
}

func (w *wA) Set() {
	w.A = "A"
}

type wSet interface {
	Set()
}

func TestWalk(t *testing.T) {
	var w wA
	var wSetType = reflect.TypeOf((*wSet)(nil)).Elem()

	Walk(&w, wSetType, func(v interface{}) error {
		v.(wSet).Set()
		return nil
	})

	if w.A != "A" {
		t.Errorf("x.A not set")
	}

	if w.B.B != "B" {
		t.Errorf("x.B not set")
	}

	if w.C.C != "C" {
		t.Errorf("x.C not set")
	}

	if w.C != w.B.C {
		t.Errorf("Expected Walk to remember value for type")
	}
}
