#!/usr/bin/env bats

load test_helper

@test "vcsa.health.mem.get" {
  vcsim_env
  local output

  run govc vcsa.health.mem.get
  assert_success

  run govc vcsa.health.mem.get -json=True
  assert_success

  run govc vcsa.health.mem.get -xml=True
  assert_success

}
