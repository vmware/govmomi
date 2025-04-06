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

    run env GOVC_SHOW_UNRELEASED=true govc storage.policy.info -json -i "VM Encryption Policy"
    assert_success

    kind="$(jq -r .policies[].filterMap[].iofilters[].filterType <<<"$output")"
    assert_equal "ENCRYPTION" "$kind"
}

@test "storage.policy.create" {
  vcsim_env

  run govc storage.policy.create MyStoragePolicy
  assert_failure # at least one of -z or -tag required

  run govc storage.policy.create -category my_cat -tag my_tag MyStoragePolicy
  assert_success

  run govc storage.policy.info MyStoragePolicy
  assert_success

  govc storage.policy.create -z MyZonalPolicy
  assert_success

  run govc storage.policy.info MyZonalPolicy
  assert_success

  run govc storage.policy.create -e MyEncryptionPolicy
  assert_success

  run govc storage.policy.info MyEncryptionPolicy
  assert_success

  run govc storage.policy.create -category my_cat -tag my_tag -z -e MyCombinedPolicy
  assert_success

  run govc storage.policy.info MyCombinedPolicy
  assert_success
}

@test "vm.policy.ls" {
  vcsim_env

  run govc vm.policy.ls -vm DC0_H0_VM0
  assert_success

  run govc vm.policy.ls -vm DC0_H0_VM0 -json
  assert_success
}
