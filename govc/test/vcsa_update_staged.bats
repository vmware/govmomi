#!/usr/bin/env bats

load test_helper

@test "update.staged.get" {
  vcsim_env
  local output

  run govc vcsa.update.staged.get
  assert_success

  run govc vcsa.update.staged.get -json=True
  assert_success

  run govc vcsa.update.staged.get -xml=True
  assert_success
}

@test "update.staged.delete" {
  vcsim_env
  local output

  run govc vcsa.update.staged.delete
  assert_success
}
