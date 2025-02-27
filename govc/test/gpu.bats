#!/usr/bin/env bats

load test_helper

@test "gpu.vm" {
  vcsim_env

  run govc gpu.vm.info
  assert_failure

  run govc gpu.vm.info '*'
  assert_failure

  run govc gpu.vm.info -vm DC0_H0_VM0
  assert_success ""
}

@test "gpu.host" {
  vcsim_start

  run govc gpu.host.info
  assert_failure

  run govc gpu.host.info -host DC0_C0_H0
  assert_success ""

  run govc gpu.host.profile.ls -host DC0_C0_H0
  assert_failure
}
