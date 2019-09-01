/*
Copyright (c) 2019 VMware, Inc. All Rights Reserved.

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

package library

import "fmt"

// ErrMultiMatch is an error returned when a query returns more than one result.
type ErrMultiMatch struct {
	// Type is the type of object being queried.
	Type string

	// Key is the key used to perform the query.
	Key string

	// Val is the value used to perform the query.
	Val string

	// Count is the number of objects returned.
	Count int
}

// Error returns the error string.
func (e ErrMultiMatch) Error() string {
	return e.String()
}

// String returns the same value as Error().
func (e ErrMultiMatch) String() string {
	return fmt.Sprintf("%q=%q matches %d items, %q id must be specified", e.Key, e.Val, e.Count, e.Type)
}
