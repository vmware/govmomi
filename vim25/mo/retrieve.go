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

package mo

import (
	"fmt"
	"reflect"
	"regexp"

	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/soap"
	"github.com/vmware/govmomi/vim25/types"
)

var arrayOfRegexp = regexp.MustCompile("ArrayOf(.*)$")

func anyTypeToValue(t interface{}) reflect.Value {
	rt := reflect.TypeOf(t)
	rv := reflect.ValueOf(t)

	// Dereference if ArrayOfXYZ type
	m := arrayOfRegexp.FindStringSubmatch(rt.Name())
	if len(m) > 0 {
		// ArrayOfXYZ type has single field named XYZ
		rv = rv.FieldByName(m[1])
		if !rv.IsValid() {
			panic(fmt.Sprintf("expected %s type to have field %s", m[0], m[1]))
		}
	}

	return rv
}

func buildValueMap(v reflect.Value, m map[string]reflect.Value) {
	t := v.Type().Elem()
	for i := 0; i < t.NumField(); i++ {
		sft := t.Field(i)
		if sft.Anonymous {
			// Recurse into embedded field
			buildValueMap(v.Elem().Field(i).Addr(), m)
			continue
		}

		tag := sft.Tag.Get("mo")
		if tag == "" {
			continue
		}

		m[tag] = v.Elem().Field(i)
	}
}

// Returns pointer to type t.
func objectContentToType(o types.ObjectContent, t reflect.Type) reflect.Value {
	v := reflect.New(t)

	// Build map of property names to assignable reflect.Value
	m := make(map[string]reflect.Value)
	buildValueMap(v, m)

	for _, p := range o.PropSet {
		rv, ok := m[p.Name]
		if ok {
			pv := anyTypeToValue(p.Val)

			// If type is a pointer, create new instance of type
			if rv.Kind() == reflect.Ptr {
				rv.Set(reflect.New(rv.Type().Elem()))
				rv = rv.Elem()
			}

			// If type is an interface, check if pv implements it
			if rv.Kind() == reflect.Interface {
				rt := rv.Type()
				pt := pv.Type()
				if !pt.Implements(rt) {
					// Check if pointer to pv implements it
					if reflect.PtrTo(pt).Implements(rt) {
						npv := reflect.New(pt)
						npv.Elem().Set(pv)
						pv = npv
					} else {
						panic(fmt.Sprintf("type %s doesn't implement %s", pt.Name(), rt.Name()))
					}
				}
			}

			rv.Set(pv)
		}
	}

	return v
}

func RetrieveProperties(r soap.RoundTripper, pc, obj types.ManagedObjectReference, dst interface{}) error {
	rt := reflect.TypeOf(dst)
	if rt.Kind() != reflect.Ptr {
		panic("need pointer")
	}

	rv := reflect.ValueOf(dst).Elem()
	if !rv.CanSet() {
		panic("cannot set dst")
	}

	var isSlice bool

	switch rt.Elem().Kind() {
	case reflect.Struct:
	case reflect.Slice:
		isSlice = true
	default:
		panic("unexpected type")
	}

	req := types.RetrieveProperties{
		This: pc,
		SpecSet: []types.PropertyFilterSpec{
			{
				ObjectSet: []types.ObjectSpec{
					{
						Obj:  obj,
						Skip: false,
					},
				},
				PropSet: []types.PropertySpec{
					{
						All:  true,
						Type: obj.Type,
					},
				},
			},
		},
	}

	res, err := methods.RetrieveProperties(r, &req)
	if err != nil {
		return err
	}

	if isSlice {
		panic("TODO(PN)")
	} else {
		switch len(res.Returnval) {
		case 0:
		case 1:
			v := objectContentToType(res.Returnval[0], rt.Elem())
			rv.Set(v.Elem())
		default:
			// If dst is not a slice, expect to receive 0 or 1 results
			panic("more than 1 result")
		}
	}

	return nil
}
