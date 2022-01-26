#!/usr/bin/env bats

load test_helper

@test "vcsa.health.database.get" {
  vcsim_env
  local output

  run govc vcsa.health.database.get
  assert_success

  run govc vcsa.health.database.get -json=True
  assert_success

  run govc vcsa.health.database.get -xml=True
  assert_success

}
