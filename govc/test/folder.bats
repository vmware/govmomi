#!/usr/bin/env bats

load test_helper

@test "folder.info" {
    for name in / vm host network datastore ; do
        run govc folder.info $name
        assert_success

        govc folder.info -json $name
        assert_success
    done

    result=$(govc folder.info '*' | grep Name: | wc -l)
    [ $result -eq 4 ]

    run govc info.info /enoent
    assert_failure
}
