#!/usr/bin/env bats

load test_helper

@test "vcsa.health.software_packages.get" {
  vcsim_env
  local output

  run govc vcsa.health.software_packages.get
  assert_success

  run govc vcsa.health.software_packages.get -json=True
  assert_success

  run govc vcsa.health.software_packages.get -xml=True
  assert_success

}
