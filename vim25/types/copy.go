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

package types

import "bytes"

// Copy creates a copy of src into dst by encoding src using the vim
// JSON encoder to encode src into JSON, then decoding the JSON into
// dst.
func Copy[T AnyType](src T) (T, error) {
	var (
		w   bytes.Buffer
		dst T
	)
	e := NewJSONEncoder(&w)
	if err := e.Encode(src); err != nil {
		return dst, err
	}
	d := NewJSONDecoder(&w)
	if err := d.Decode(&dst); err != nil {
		return dst, err
	}
	return dst, nil
}

// MustCopy panics if Copy returns an error.
func MustCopy[T AnyType](src T) T {
	dst, err := Copy(src)
	if err != nil {
		panic(err)
	}
	return dst
}
