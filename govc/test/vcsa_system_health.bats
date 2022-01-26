#!/usr/bin/env bats

load test_helper

@test "vcsa.health.system.get" {
  vcsim_env
  local output

  run govc vcsa.health.system.get
  assert_success

  run govc vcsa.health.system.get -json=True
  assert_success

  run govc vcsa.health.system.get -xml=True
  assert_success

}

@test "vcsa.health.system.last_check" {
  vcsim_env
  local output

  run govc vcsa.health.system.last_check
  assert_success

  run govc vcsa.health.system.last_check -json=True
  assert_success

  run govc vcsa.health.system.last_check -xml=True
  assert_success

}
