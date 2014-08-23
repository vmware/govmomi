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

type test2A struct {
	V int
}

func (t *test2A) Set() {
	t.V++
}

type test2B struct {
	t *testing.T

	F1 *test2A
	F2 *test2A
	F3 *test2A
	F4 *test2A
}

func (t *test2B) Set() {
	// Assert that this function gets called before the walker recurses.
	if t.F1 != nil || t.F3 != nil {
		t.t.Errorf("Expected user function to be called before recursing")
	}
}

type test2Set interface {
	Set()
}

func TestWalkDoesntOverwrite(t *testing.T) {
	var t2a test2A
	var t2b test2B
	var testSetType = reflect.TypeOf((*test2Set)(nil)).Elem()

	// Set elements 2 and 4
	t2b.t = t
	t2b.F2 = &t2a
	t2b.F4 = &t2a

	Walk(&t2b, testSetType, func(v interface{}) error {
		v.(test2Set).Set()
		return nil
	})

	// Set fields should remain untouched
	if t2b.F2 != &t2a {
		t.Errorf("Expected t2b.F2 to be left intact")
	}

	if t2b.F4 != &t2a {
		t.Errorf("Expected t2b.F4 to be left intact")
	}

	// Unset fields should be filled in with the same value
	if t2b.F1 != t2b.F3 {
		t.Errorf("Expected t2b.F1 to be equal to t2b.F3")
	}

	// All fields should be traversed.
	if t2b.F1.V != 2 || t2b.F2.V != 2 {
		t.Errorf("Expected all fields to be traversed")
	}
}
