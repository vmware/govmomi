#!/usr/bin/env bats

load test_helper

@test "vcsa.health.storage.get" {
  vcsim_env
  local output

  run govc vcsa.health.storage.get
  assert_success

  run govc vcsa.health.storage.get -json=True
  assert_success

  run govc vcsa.health.storage.get -xml=True
  assert_success

}
