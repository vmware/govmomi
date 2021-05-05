# Functional Tests for govc

The govc tests use [bats](https://github.com/sstephenson/bats/), a framework that provides a simple way to verify *NIX programs behave as expected.

## Enabling ESX Tests

Some of the govc tests require an ESX instance. These tests may be enabled with any ESX host with the following environment variable:

```
GOVC_TEST_URL=user:pass@<ESX_HOST>
```

## Running Tests

There are two ways to run the tests locally:

1. [_Via GitHub Actions in Docker_](#run-tests-via-github-actions-in-docker) - mimics how GitHub actions will execute the tests
2. [_Natively on Linux and macOS_](#run-tests-natively-on-linux-and-macos) - the fastest way to run the tests

### Run Tests via GitHub Actions in Docker

This method of running the govc functional tests mimics how the project's associated GitHub actions will execute the tests when new pull requests are opened:

1. Install [`act`](https://github.com/nektos/act):

    * Linux (curl bash)

        ```shell
        curl https://raw.githubusercontent.com/nektos/act/master/install.sh | sudo bash
        ```

    * macOS

        ```shell
        brew install act
        ```

    * [_Documentation for other installation methods_](https://github.com/nektos/act#installation)

2. Run the `govc-tests` action:

    ```shell
    act --env USER=user -j govc-tests
    ```

    ---

    **Note**: To run the ESX tests with `act`, execute the following command:

    ```shell
    act --env USER=user --env GOVC_TEST_URL="user:pass@<ESX_HOST>" -j govc-tests
    ```

    ---

### Run Tests Natively on Linux and macOS

The fastest way to run the tests is to do so natively on Linux or macOS:

1. Install [bats](https://github.com/sstephenson/bats/):

    * Debian / Ubuntu Linux

        ```shell
        apt-get install bats
        ```

    * macOS

        ```shell
        brew install bats
        ```

    * [_Installing bats from source_](https://github.com/sstephenson/bats#installing-bats-from-source)


2. Some tests depend on the `ttylinux` images, and they can be downloaded with the following command:

    ```shell
    ./govc/test/images/update.sh
    ```

    The images are uploaded to the `$GOVC_TEST_URL` as needed by tests and can be removed with the following command:

    ```shell
    ./govc/test/images/clean.sh
    ```

---

**Note**: Users of macOS will want to install `gxargs`, `greadlink`, and `gmktemp` with the following command:

```shell
brew install coreutils findutils
```

---

3. The easiest way to run the tests is with the top-level Makefile:

    ```shell
    make govc-test
    ```

    The tests can also be run directly with the `bats` command:

    ```shell
    bats -t ./govc/test/
    ```

    Finally, it's also possible to run the tests individually with:

    ```shell
    ./govc/test/cli.bats
    ```
