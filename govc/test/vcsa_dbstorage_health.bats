#!/usr/bin/env bats

load test_helper

@test "vcsa.health.database_storage.get" {
  vcsim_env
  local output

  run govc vcsa.health.database_storage.get
  assert_success

  run govc vcsa.health.database_storage.get -json=True
  assert_success

  run govc vcsa.health.database_storage.get -xml=True
  assert_success

}
