#!/usr/bin/env bats

load test_helper

@test "vm.ip" {
  id=$(new_ttylinux_vm)

  run govc vm.power -on $id
  assert_success

  run govc vm.ip $id
  assert_success
}
