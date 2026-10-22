#!/usr/bin/env bats

load test_helper

@test "vm.vapp.info" {
  vcsim_env

  id=/DC0/vm/DC0_H0_VM0

  # a stock vcsim VM has no vApp properties until deployed/cloned from an OVF
  run govc vm.vapp.info -vm $id
  assert_success
  assert_equal "Key  Id  Label  Type  Value" "$output"

  run govc vm.vapp.info -vm $id -json
  assert_success
  assert_equal "null" "$output"
}

@test "vm.vapp.set unknown property" {
  vcsim_env

  id=/DC0/vm/DC0_H0_VM0

  run govc vm.vapp.set -vm $id -property guestinfo.hostname=my-vm
  assert_failure
  assert_matches "guestinfo.hostname"
  assert_matches "does not exist"
}

@test "vm.vapp.set requires -property" {
  vcsim_env

  id=/DC0/vm/DC0_H0_VM0

  run govc vm.vapp.set -vm $id
  assert_failure
}
