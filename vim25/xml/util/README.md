# XML Utilities

This package provides helper functions for marshaling/unmarshaling vim objects to and from XML across the various languages that have vSphere API clients, such as Golang and Python.

## Notes

* Does not need to know about the type of object being decoded ahead of time!
* Can decode _multiple_ VIM objects from one XML string and return them into a channel.
* Produces output compatible with the Python and Java SDKs for vSphere.

## Examples

* The directory [`./examples`](./examples) illustrates how to share VIM objects between Go and Python using XML!
* The file [`util_examples_test.go`](util_examples_test.go) includes examples for how to decode/encode objects from/to XML using these new helper functions.
* The file [`util_test.go`](util_test.go) includes even more examples of the new helper functions.
