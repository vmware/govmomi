# Functional tests for govc

## Bats

Install [Bats](https://github.com/sstephenson/bats/)

## Download test images

Some tests depend on [ttylinux](http://ttylinux.net) images, these can be downloaded by running:

```
./images/update.sh
```

## GOVC_TEST_URL

The govc tests need an ESX instance to run against.  The default
`GOVC_TEST_URL` is that of the vagrant box in this directory:

```
vagrant up
```

Any other ESX box can be used by exporting the following variable:

```
export GOVC_TEST_URL=user:pass@hostname
```

## Running tests

The *govc* binary should be in your `PATH`; the test helper also prepends ../govc to `PATH`.

The tests can be run from any directory, as *govc* is found related to
`PATH` and *images* are found relative to `$BATS_TEST_DIRNAME`.
