/*
Copyright (c) 2021 VMware, Inc. All Rights Reserved.

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

// Package util implements some helper functions for marshaling/unmarshaling
// vim objects to/from XML.
package util

import (
	"bytes"
	"fmt"
	"io"
	"reflect"
	"strings"

	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/types"
	"github.com/vmware/govmomi/vim25/xml"
)

// VimObjectsFromBytes returns all vim objects decoded from the provided bytes.
//
// This is a helper function that wraps VimObjectsFromDecoder. Please see
// the documentation on VimObjectsFromDecoder for additional information.
func VimObjectsFromBytes(v []byte) (<-chan interface{}, <-chan error) {
	return VimObjectsFromReader(bytes.NewReader(v))
}

// VimObjectsFromString returns all vim objects decoded from the provided
// string.
//
// This is a helper function that wraps VimObjectsFromDecoder. Please see
// the documentation on VimObjectsFromDecoder for additional information.
func VimObjectsFromString(v string) (<-chan interface{}, <-chan error) {
	return VimObjectsFromReader(strings.NewReader(v))
}

// VimObjectsFromReader returns all vim objects decoded from the provided input
// stream.
//
// This is a helper function that wraps VimObjectsFromDecoder. Please see
// the documentation on VimObjectsFromDecoder for additional information.
func VimObjectsFromReader(r io.Reader) (<-chan interface{}, <-chan error) {
	return VimObjectsFromDecoder(xml.NewDecoder(r))
}

// VimObjectsFromDecoder returns all vim objects from the provided decoder.
//
// Any decode errors are sent on the error channel.
//
// Both channels are closed when io.EOF is reached.
//
// The decoder's TypeFunc is used to construct types from their XML data.
// If the decoder's TypeFunc is nil then vim25.TypeFunc is used.
func VimObjectsFromDecoder(d *xml.Decoder) (<-chan interface{}, <-chan error) {

	var (
		chanObj = make(chan interface{})
		chanErr = make(chan error)
	)

	// If the decoder did not provide a TypeFunc then use the one from the
	// vim25.types package.
	if d.TypeFunc == nil {
		d.TypeFunc = types.TypeFunc()
	}

	// Returns the value of the "xsi:type" attribute if it exists, otherwise
	// the value of the "type" attribute, otherwise an empty string.
	getVimTypeName := func(attrs []xml.Attr) string {
		var vimTypeName string
		for _, a := range attrs {
			if a.Name.Local == "type" {
				if a.Name.Space == xml.SchemaInstanceURI {
					vimTypeName = a.Value
				} else if vimTypeName == "" {
					vimTypeName = a.Value
				}
			}
		}
		return vimTypeName
	}

	// Decode the input stream in the background, returning tokens/errors on
	// the chanObj and chanErr channels.
	go func() {

		defer func() {
			// Once there are no more tokens and/or an error occurred, the loop
			// will exit. At this point we should notify readers by closing the
			// channels to indicate the operation has completed.
			close(chanObj)
			close(chanErr)
		}()

		// Keep iterating until there are no more tokens.
		for {

			// Get the next token.
			tok, err := d.Token()

			// If there are no more tokens then break out othe loop.
			if tok == nil && err == io.EOF {
				return
			}

			// If there was an error then send that over the channel and jump to
			// the next iteration of the loop.
			if err != nil {
				chanErr <- err
				continue
			}

			// If the token is not a start element then we can jump to the next
			// iteration of the loop.
			startElement, ok := tok.(xml.StartElement)
			if !ok {
				continue
			}

			// If there is no vim type name then jump to the next iteration of
			// the loop.
			vimTypeName := getVimTypeName(startElement.Attr)
			if vimTypeName == "" {
				continue
			}

			// At this point we have what looks like an XML element with a vim
			// type name, so let's try and get that from the registered lookup
			// function.
			//
			// If we cannot find a registered type with that name then we send
			// an error over the channel and jump to the next iteration of the
			// loop.
			vimType, ok := d.TypeFunc(vimTypeName)
			if !ok {
				chanErr <- fmt.Errorf("vim type not found: %s", vimTypeName)
				continue
			}

			// Create a new object from the vim type.
			vimObj := reflect.New(vimType).Interface()

			// Try to decode the element into the new vim object. If this
			// operation fails then the entire decode function must cease here
			// because the position of the underlying stream is not advanced
			// until the object can be decoded successfully.
			if err := d.DecodeElement(vimObj, &startElement); err != nil {
				chanErr <- err
				return
			}

			// Send the decoded object into the channel.
			chanObj <- vimObj
		}
	}()

	return chanObj, chanErr
}

// VimObjectsToBytes marshals the provided vim objects to XML and returns the
// result as a byte slice.
//
// This is a helper function that wraps VimObjectsToEncoder. Please see
// the documentation on VimObjectsToEncoder for additional information.
func VimObjectsToBytes(objs ...interface{}) ([]byte, error) {
	w := &bytes.Buffer{}
	if err := VimObjectsToWriter(w, objs...); err != nil {
		return nil, err
	}
	return w.Bytes(), nil
}

// VimObjectsToString marshals the provided vim objects to XML and returns the
// result as a string.
//
// This is a helper function that wraps VimObjectsToEncoder. Please see
// the documentation on VimObjectsToEncoder for additional information.
func VimObjectsToString(objs ...interface{}) (string, error) {
	w := &bytes.Buffer{}
	if err := VimObjectsToWriter(w, objs...); err != nil {
		return "", err
	}
	return w.String(), nil
}

// VimObjectsToWriter marshals the provided vim objects to XML using the
// provided io.Writer.
//
// This is a helper function that wraps VimObjectsToEncoder. Please see
// the documentation on VimObjectsToEncoder for additional information.
func VimObjectsToWriter(w io.Writer, objs ...interface{}) error {
	return VimObjectsToEncoder(xml.NewEncoder(w), objs...)
}

// VimObjectsToEncoder marshals vim objects using an xml.Encoder so the result
// can be unmarshaled by other vSphere SDKs, such as Java, Python, Ruby, etc.
func VimObjectsToEncoder(e *xml.Encoder, objs ...interface{}) error {

	for i := range objs {
		vimTypeName := fmt.Sprintf(
			"%s:%s",
			vim25.Namespace,
			reflect.TypeOf(objs[i]).Elem().Name())

		startElement := xml.StartElement{
			Name: xml.Name{
				Local: "obj",
			},
			Attr: []xml.Attr{
				{
					Name:  xml.Name{Local: "xmlns:" + vim25.Namespace},
					Value: "urn:" + vim25.Namespace,
				},
				{
					Name:  xml.Name{Local: "xmlns:xsi"},
					Value: xml.SchemaInstanceURI,
				},
				{
					Name:  xml.Name{Local: "xsi:type"},
					Value: vimTypeName,
				},
			},
		}
		if err := e.EncodeElement(objs[i], startElement); err != nil {
			return err
		}
	}

	return nil
}
