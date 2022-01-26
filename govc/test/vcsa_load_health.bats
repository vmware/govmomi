#!/usr/bin/env bats

load test_helper

@test "vcsa.health.load.get" {
  vcsim_env
  local output

  run govc vcsa.health.load.get
  assert_success

  run govc vcsa.health.load.get -json=True
  assert_success

  run govc vcsa.health.load.get -xml=True
  assert_success

}
