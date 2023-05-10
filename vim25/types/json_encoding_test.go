/*
Copyright (c) 2023-2023 VMware, Inc. All Rights Reserved.

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

package types

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOptionValueSerialization(t *testing.T) {
	options := []struct {
		name    string
		wire    string
		binding OptionValue
	}{
		{
			name: "boolean",
			wire: `{"_typeName": "OptionValue","key": "option1",
				"value": {"_typeName": "boolean","_value": true}
			}`,
			binding: OptionValue{Key: "option1", Value: true},
		},
		{
			name: "byte",
			wire: `{"_typeName": "OptionValue","key": "option1",
				"value": {"_typeName": "byte","_value": 16}
			}`,
			binding: OptionValue{Key: "option1", Value: uint8(16)},
		},
		{
			name: "short",
			wire: `{"_typeName": "OptionValue","key": "option1",
				"value": {"_typeName": "short","_value": 300}
			}`,
			binding: OptionValue{Key: "option1", Value: int16(300)},
		},
		{
			name: "int",
			wire: `{"_typeName": "OptionValue","key": "option1",
				"value": {"_typeName": "int","_value": 300}}`,
			binding: OptionValue{Key: "option1", Value: int32(300)},
		},
		{
			name: "long",
			wire: `{"_typeName": "OptionValue","key": "option1",
				"value": {"_typeName": "long","_value": 300}}`,
			binding: OptionValue{Key: "option1", Value: int64(300)},
		},
		{
			name: "float",
			wire: `{"_typeName": "OptionValue","key": "option1",
				"value": {"_typeName": "float","_value": 30.5}}`,
			binding: OptionValue{Key: "option1", Value: float32(30.5)},
		},
		{
			name: "double",
			wire: `{"_typeName": "OptionValue","key": "option1",
				"value": {"_typeName": "double","_value": 12.2}}`,
			binding: OptionValue{Key: "option1", Value: float64(12.2)},
		},
		{
			name: "string",
			wire: `{"_typeName": "OptionValue","key": "option1",
				"value": {"_typeName": "string","_value": "test"}}`,
			binding: OptionValue{Key: "option1", Value: "test"},
		},
	}

	for _, opt := range options {
		t.Run("Serialize "+opt.name, func(t *testing.T) {
			r := strings.NewReader(opt.wire)
			dec := NewJSONDecoder(r)
			v := OptionValue{}
			e := dec.Decode(&v)
			if e != nil {
				assert.Fail(t, "Cannot read json", "json %v err %v", opt.wire, e)
				return
			}
			assert.Equal(t, opt.binding, v)
		})

		t.Run("De-serialize "+opt.name, func(t *testing.T) {
			var w bytes.Buffer
			enc := NewJSONEncoder(&w)
			enc.Encode(opt.binding)
			s := w.String()
			assert.JSONEq(t, opt.wire, s)
		})
	}
}
