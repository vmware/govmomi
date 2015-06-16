# changelog

### (unreleased)

...

* Add retry functionality to vim25 package

* Change finder functions to no longer take varargs

    The `find` package had functions to return a list of objects, given a
    variable number of patterns. This makes it impossible to distinguish which
    patterns produced results and which ones didn't.

    In particular for govc, where multiple arguments can be passed from the
    command line, it is useful to let the user know which ones produce results
    and which ones don't.

    To evaluate multiple patterns, the user should call the find functions
    multiple times (either serially or in parallel).

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
