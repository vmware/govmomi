#!/usr/bin/env bats

load test_helper

@test "vcsa.health.applmgmt.get" {
  vcsim_env
  local output

  run govc vcsa.health.applmgmt.get
  assert_success

  run govc vcsa.health.applmgmt.get -json=True
  assert_success

  run govc vcsa.health.applmgmt.get -xml=True
  assert_success

}
