#!/usr/bin/env bats

load test_helper

@test "vcsa.health.swap.get" {
  vcsim_env
  local output

  run govc vcsa.health.swap.get
  assert_success

  run govc vcsa.health.swap.get -json=True
  assert_success

  run govc vcsa.health.swap.get -xml=True
  assert_success

}
