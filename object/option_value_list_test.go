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

package object_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/vim25/types"
)

type nillableOptionValue struct{}

func (ov nillableOptionValue) GetOptionValue() *types.OptionValue {
	return nil
}

func TestOptionValueList(t *testing.T) {
	const (
		sza   = "a"
		szb   = "b"
		szc   = "c"
		szd   = "d"
		sze   = "e"
		sz1   = "1"
		sz2   = "2"
		sz3   = "3"
		sz4   = "4"
		sz5   = "5"
		i32_1 = int32(1)
		u64_2 = uint64(2)
		f32_3 = float32(3)
		f64_4 = float64(4)
		b_5   = byte(5) //nolint:revive,stylecheck
	)

	var (
		psz1   = &[]string{sz1}[0]
		pu64_2 = &[]uint64{u64_2}[0]
		pf32_3 = &[]float32{f32_3}[0]
		pb_5   = &[]byte{b_5}[0] //nolint:revive,stylecheck
	)

	t.Run("OptionValueListFromMap", func(t *testing.T) {

		t.Run("a nil map should return nil", func(t *testing.T) {
			assert.Nil(t, object.OptionValueListFromMap[any](nil))
		})

		t.Run("a map with string values should return OptionValues", func(t *testing.T) {
			assert.ElementsMatch(
				t,
				object.OptionValueListFromMap(map[string]string{
					szc: sz3,
					sza: sz1,
					szb: sz2,
				}),
				[]types.BaseOptionValue{
					&types.OptionValue{Key: szb, Value: sz2},
					&types.OptionValue{Key: sza, Value: sz1},
					&types.OptionValue{Key: szc, Value: sz3},
				},
			)
		})

		t.Run("a map with values of varying numeric types should return OptionValues", func(t *testing.T) {
			assert.ElementsMatch(
				t,
				object.OptionValueListFromMap(map[string]any{
					szc: f32_3,
					sza: i32_1,
					szb: u64_2,
				}),
				[]types.BaseOptionValue{
					&types.OptionValue{Key: szb, Value: u64_2},
					&types.OptionValue{Key: sza, Value: i32_1},
					&types.OptionValue{Key: szc, Value: f32_3},
				},
			)
		})

		t.Run("a map with pointer values should return OptionValues", func(t *testing.T) {
			assert.ElementsMatch(
				t,
				object.OptionValueListFromMap(map[string]any{
					szc: pf32_3,
					sza: psz1,
					szb: pu64_2,
				}),
				[]types.BaseOptionValue{
					&types.OptionValue{Key: szb, Value: pu64_2},
					&types.OptionValue{Key: sza, Value: psz1},
					&types.OptionValue{Key: szc, Value: pf32_3},
				},
			)
		})
	})

	t.Run("Get", func(t *testing.T) {
		testCases := []struct {
			name string
			left object.OptionValueList
			key  string
			out  any
			ok   bool
		}{
			{
				name: "a nil receiver should not panic and return nil, false",
				left: nil,
				key:  "",
				out:  nil,
				ok:   false,
			},
			{
				name: "a non-existent key should return nil, false",
				left: object.OptionValueList{},
				key:  "",
				out:  nil,
				ok:   false,
			},
			{
				name: "an existing key should return its value, true",
				left: object.OptionValueList{
					&types.OptionValue{Key: sza, Value: sz1},
				},
				key: sza,
				out: sz1,
				ok:  true,
			},
			{
				name: "an existing key should return its value, true when data includes a nillable types.BaseOptionValue",
				left: object.OptionValueList{
					nillableOptionValue{},
					&types.OptionValue{Key: sza, Value: sz1},
				},
				key: sza,
				out: sz1,
				ok:  true,
			},
		}

		for i := range testCases {
			tc := testCases[i]
			t.Run(tc.name, func(t *testing.T) {
				var (
					out any
					ok  bool
				)
				assert.NotPanics(t, func() { out, ok = tc.left.Get(tc.key) })
				assert.Equal(t, tc.out, out)
				assert.Equal(t, tc.ok, ok)
			})
		}
	})

	t.Run("GetString", func(t *testing.T) {
		testCases := []struct {
			name string
			left object.OptionValueList
			key  string
			out  string
			ok   bool
		}{
			{
				name: "a nil receiver should not panic and return \"\", false",
				left: nil,
				key:  "",
				out:  "",
				ok:   false,
			},
			{
				name: "a non-existent key should return \"\", false",
				left: object.OptionValueList{},
				key:  "",
				out:  "",
				ok:   false,
			},
			{
				name: "an existing key for a string value should return its string value, true",
				left: object.OptionValueList{
					&types.OptionValue{Key: sza, Value: sz1},
				},
				key: sza,
				out: sz1,
				ok:  true,
			},
			{
				name: "an existing key for a *string value that is not nil should return its string value, true",
				left: object.OptionValueList{
					&types.OptionValue{Key: sza, Value: psz1},
				},
				key: sza,
				out: sz1,
				ok:  true,
			},
			{
				name: "an existing key for a *string value that is nil should return \"\", true",
				left: object.OptionValueList{
					&types.OptionValue{Key: sza, Value: (*string)(nil)},
				},
				key: sza,
				out: "",
				ok:  true,
			},
			{
				name: "an existing key for an int32 value should return its string value, true",
				left: object.OptionValueList{
					&types.OptionValue{Key: sza, Value: i32_1},
				},
				key: sza,
				out: sz1,
				ok:  true,
			},
			{
				name: "an existing key for a *uint64 value that is not nil should return its string value, true",
				left: object.OptionValueList{
					&types.OptionValue{Key: sza, Value: pu64_2},
				},
				key: sza,
				out: sz2,
				ok:  true,
			},
			{
				name: "an existing key for a *uint64 value that is nil should return \"\", true",
				left: object.OptionValueList{
					&types.OptionValue{Key: sza, Value: (*uint64)(nil)},
				},
				key: sza,
				out: "",
				ok:  true,
			},
			{
				name: "an existing key for a string value should return its string value, true when data includes a nillable types.BaseOptionValue",
				left: object.OptionValueList{
					nillableOptionValue{},
					&types.OptionValue{Key: sza, Value: sz1},
				},
				key: sza,
				out: sz1,
				ok:  true,
			},
		}

		for i := range testCases {
			tc := testCases[i]
			t.Run(tc.name, func(t *testing.T) {
				var (
					out string
					ok  bool
				)
				assert.NotPanics(t, func() { out, ok = tc.left.GetString(tc.key) })
				assert.Equal(t, tc.out, out)
				assert.Equal(t, tc.ok, ok)
			})
		}
	})

	t.Run("Map", func(t *testing.T) {
		testCases := []struct {
			name string
			left object.OptionValueList
			out  map[string]any
		}{
			{
				name: "a nil receiver should not panic and return nil",
				left: nil,
				out:  nil,
			},
			{
				name: "data with homogeneous values should return a map",
				left: object.OptionValueList{
					&types.OptionValue{Key: sza, Value: sz1},
				},
				out: map[string]any{
					sza: sz1,
				},
			},
			{
				name: "data with heterogeneous values should return a map",
				left: object.OptionValueList{
					&types.OptionValue{Key: sza, Value: sz1},
					&types.OptionValue{Key: szb, Value: u64_2},
					&types.OptionValue{Key: szc, Value: pf32_3},
				},
				out: map[string]any{
					sza: sz1,
					szb: u64_2,
					szc: pf32_3,
				},
			},
			{
				name: "data with just a nillable types.BaseOptionValue should return nil",
				left: object.OptionValueList{
					nillableOptionValue{},
				},
				out: nil,
			},
		}

		for i := range testCases {
			tc := testCases[i]
			t.Run(tc.name, func(t *testing.T) {
				var out map[string]any
				assert.NotPanics(t, func() { out = tc.left.Map() })
				assert.Equal(t, tc.out, out)
			})
		}
	})

	t.Run("StringMap", func(t *testing.T) {
		testCases := []struct {
			name string
			left object.OptionValueList
			out  map[string]string
		}{
			{
				name: "a nil receiver should not panic and return nil",
				left: nil,
				out:  nil,
			},
			{
				name: "data with homogeneous values should return a map",
				left: object.OptionValueList{
					&types.OptionValue{Key: sza, Value: sz1},
				},
				out: map[string]string{
					sza: sz1,
				},
			},
			{
				name: "data with heterogeneous values should return a map",
				left: object.OptionValueList{
					&types.OptionValue{Key: sza, Value: sz1},
					&types.OptionValue{Key: szb, Value: u64_2},
					&types.OptionValue{Key: szc, Value: pf32_3},
				},
				out: map[string]string{
					sza: sz1,
					szb: sz2,
					szc: sz3,
				},
			},
			{
				name: "data with just a nillable types.BaseOptionValue should return nil",
				left: object.OptionValueList{
					nillableOptionValue{},
				},
				out: nil,
			},
		}

		for i := range testCases {
			tc := testCases[i]
			t.Run(tc.name, func(t *testing.T) {
				var out map[string]string
				assert.NotPanics(t, func() { out = tc.left.StringMap() })
				assert.Equal(t, tc.out, out)
			})
		}
	})

	t.Run("Additions", func(t *testing.T) {
		testCases := []struct {
			name  string
			left  object.OptionValueList
			right object.OptionValueList
			out   object.OptionValueList
		}{
			{
				name:  "a nil receiver and nil input should not panic and return nil",
				left:  nil,
				right: nil,
				out:   nil,
			},
			{
				name: "a nil receiver and non-nil input should not panic and return the diff",
				left: nil,
				right: object.OptionValueList{
					&types.OptionValue{Key: szb, Value: ""},
					&types.OptionValue{Key: szd, Value: f64_4},
					&types.OptionValue{Key: sze, Value: pb_5},
				},
				out: object.OptionValueList{
					&types.OptionValue{Key: szb, Value: ""},
					&types.OptionValue{Key: szd, Value: f64_4},
					&types.OptionValue{Key: sze, Value: pb_5},
				},
			},
			{
				name: "a non-nil receiver and nil input should return nil",
				left: object.OptionValueList{
					&types.OptionValue{Key: sza, Value: sz1},
					&types.OptionValue{Key: szb, Value: sz2},
					&types.OptionValue{Key: szc, Value: sz3},
				},
				right: nil,
				out:   nil,
			},
			{
				name: "a non-nil receiver and non-nil input should return the diff",
				left: object.OptionValueList{
					&types.OptionValue{Key: sza, Value: sz1},
					&types.OptionValue{Key: szb, Value: sz2},
					&types.OptionValue{Key: szc, Value: sz3},
				},
				right: object.OptionValueList{
					&types.OptionValue{Key: szb, Value: ""},
					&types.OptionValue{Key: szd, Value: f64_4},
					&types.OptionValue{Key: sze, Value: pb_5},
				},
				out: object.OptionValueList{
					&types.OptionValue{Key: szd, Value: f64_4},
					&types.OptionValue{Key: sze, Value: pb_5},
				},
			},
		}

		for i := range testCases {
			tc := testCases[i]
			t.Run(tc.name, func(t *testing.T) {
				var out object.OptionValueList
				assert.NotPanics(t, func() { out = tc.left.Additions(tc.right...) })
				assert.Equal(t, tc.out, out)
			})
		}
	})

	t.Run("Diff", func(t *testing.T) {
		testCases := []struct {
			name  string
			left  object.OptionValueList
			right object.OptionValueList
			out   object.OptionValueList
		}{
			{
				name:  "a nil receiver and nil input should not panic and return nil",
				left:  nil,
				right: nil,
				out:   nil,
			},
			{
				name: "a nil receiver and non-nil input should not panic and return the diff",
				left: nil,
				right: object.OptionValueList{
					&types.OptionValue{Key: szb, Value: ""},
					&types.OptionValue{Key: szd, Value: f64_4},
					&types.OptionValue{Key: sze, Value: pb_5},
				},
				out: object.OptionValueList{
					&types.OptionValue{Key: szb, Value: ""},
					&types.OptionValue{Key: szd, Value: f64_4},
					&types.OptionValue{Key: sze, Value: pb_5},
				},
			},
			{
				name: "a non-nil receiver and nil input should return nil",
				left: object.OptionValueList{
					&types.OptionValue{Key: sza, Value: sz1},
					&types.OptionValue{Key: szb, Value: sz2},
					&types.OptionValue{Key: szc, Value: sz3},
				},
				right: nil,
				out:   nil,
			},
			{
				name: "a non-nil receiver and non-nil input should return the diff",
				left: object.OptionValueList{
					&types.OptionValue{Key: sza, Value: sz1},
					&types.OptionValue{Key: szb, Value: sz2},
					&types.OptionValue{Key: szc, Value: sz3},
				},
				right: object.OptionValueList{
					&types.OptionValue{Key: szb, Value: ""},
					&types.OptionValue{Key: szd, Value: f64_4},
					&types.OptionValue{Key: sze, Value: pb_5},
				},
				out: object.OptionValueList{
					&types.OptionValue{Key: szb, Value: ""},
					&types.OptionValue{Key: szd, Value: f64_4},
					&types.OptionValue{Key: sze, Value: pb_5},
				},
			},
		}

		for i := range testCases {
			tc := testCases[i]
			t.Run(tc.name, func(t *testing.T) {
				var out object.OptionValueList
				assert.NotPanics(t, func() { out = tc.left.Diff(tc.right...) })
				assert.Equal(t, tc.out, out)
			})
		}
	})

	t.Run("Join", func(t *testing.T) {
		testCases := []struct {
			name  string
			left  object.OptionValueList
			right object.OptionValueList
			out   object.OptionValueList
		}{
			{
				name:  "a nil receiver and nil input should not panic and return nil",
				left:  nil,
				right: nil,
				out:   nil,
			},
			{
				name: "a nil receiver and non-nil input should not panic and return the joined data",
				left: nil,
				right: object.OptionValueList{
					&types.OptionValue{Key: szb, Value: ""},
					&types.OptionValue{Key: szd, Value: f64_4},
					&types.OptionValue{Key: sze, Value: pb_5},
				},
				out: object.OptionValueList{
					&types.OptionValue{Key: szb, Value: ""},
					&types.OptionValue{Key: szd, Value: f64_4},
					&types.OptionValue{Key: sze, Value: pb_5},
				},
			},
			{
				name: "a non-nil receiver and nil input should return the joined data",
				left: object.OptionValueList{
					&types.OptionValue{Key: sza, Value: sz1},
					&types.OptionValue{Key: szb, Value: sz2},
					&types.OptionValue{Key: szc, Value: sz3},
				},
				right: nil,
				out: object.OptionValueList{
					&types.OptionValue{Key: sza, Value: sz1},
					&types.OptionValue{Key: szb, Value: sz2},
					&types.OptionValue{Key: szc, Value: sz3},
				},
			},
			{
				name: "a non-nil receiver and non-nil input should return the joined data",
				left: object.OptionValueList{
					&types.OptionValue{Key: sza, Value: sz1},
					&types.OptionValue{Key: szb, Value: sz2},
					&types.OptionValue{Key: szc, Value: sz3},
				},
				right: object.OptionValueList{
					&types.OptionValue{Key: szb, Value: ""},
					&types.OptionValue{Key: szd, Value: f64_4},
					&types.OptionValue{Key: sze, Value: pb_5},
				},
				out: object.OptionValueList{
					&types.OptionValue{Key: sza, Value: sz1},
					&types.OptionValue{Key: szb, Value: sz2},
					&types.OptionValue{Key: szc, Value: sz3},
					&types.OptionValue{Key: szd, Value: f64_4},
					&types.OptionValue{Key: sze, Value: pb_5},
				},
			},
			{
				name: "a non-nil receiver and non-nil input, flipping left and right, should return the joined data",
				left: object.OptionValueList{
					&types.OptionValue{Key: szb, Value: ""},
					&types.OptionValue{Key: szd, Value: f64_4},
					&types.OptionValue{Key: sze, Value: pb_5},
				},
				right: object.OptionValueList{
					&types.OptionValue{Key: sza, Value: sz1},
					&types.OptionValue{Key: szb, Value: sz2},
					&types.OptionValue{Key: szc, Value: sz3},
				},
				out: object.OptionValueList{
					&types.OptionValue{Key: szb, Value: ""},
					&types.OptionValue{Key: szd, Value: f64_4},
					&types.OptionValue{Key: sze, Value: pb_5},
					&types.OptionValue{Key: sza, Value: sz1},
					&types.OptionValue{Key: szc, Value: sz3},
				},
			},
		}

		for i := range testCases {
			tc := testCases[i]
			t.Run(tc.name, func(t *testing.T) {
				var out object.OptionValueList
				assert.NotPanics(t, func() { out = tc.left.Join(tc.right...) })
				assert.Equal(t, tc.out, out)
			})
		}
	})
}
