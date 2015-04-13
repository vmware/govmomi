# changelog

### (unreleased)

...

* Make optional boolean fields pointers (`vim25/types`).

    False is the zero value of a boolean field, which means they are not serialized
    if the field is marked "omitempty". If the field is a pointer instead, the zero
    value will be the nil pointer, and both true and false values are serialized.

### 0.1.0 (2015-03-17)

Prior to this version the API of this library was in flux.

Notable changes w.r.t. the state of this library before March 2015 are:

* All functions that may execute a request take a `context.Context` parameter.
* The `vim25` package contains a minimal client implementation.
* The property collector and its convenience functions live in the `property` package.
