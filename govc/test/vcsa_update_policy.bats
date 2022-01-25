#!/usr/bin/env bats

load test_helper

@test "update.policy.get" {
  vcsim_env
  local output

  run govc vcsa.update.policy.get
  assert_success

  run govc vcsa.update.policy.get -json=True
  assert_success

  run govc vcsa.update.policy.get -xml=True
  assert_success

}

@test "update.policy.set" {
  vcsim_env
  local output

  run govc vcsa.update.policy.set MONDAY
  assert_success

  run govc vcsa.update.policy.set -auto_stage=false MONDAY
  assert_success

  run govc vcsa.update.policy.set -certificate_check=true MONDAY
  assert_success

  run govc vcsa.update.policy.set --custom_URL="https://test-url" MONDAY
  assert_success

  run govc vcsa.update.policy.set -username="vmware" -password="vmware" MONDAY
  assert_success

  run govc vcsa.update.policy.set  -minute=54 -hour=9 MONDAY
  assert_success

  run govc vcsa.update.policy.set
  assert_failure
}
