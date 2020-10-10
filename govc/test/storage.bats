#!/usr/bin/env bats

load test_helper

@test "storage.policy.info" {
    vcsim_env

    run govc storage.policy.ls
    assert_success

    run govc storage.policy.ls enoent
    assert_failure

    run govc storage.policy.ls "vSAN Default Storage Policy"
    assert_success

    run govc storage.policy.ls -i "vSAN Default Storage Policy"
    assert_success

    run govc storage.policy.info
    assert_success

    run govc storage.policy.info enoent
    assert_failure

    run govc storage.policy.info "vSAN Default Storage Policy"
    assert_success
}
