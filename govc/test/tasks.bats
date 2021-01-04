#!/usr/bin/env bats

load test_helper

@test "tasks" {
  vcsim_env

  run govc tasks
  assert_success
}

@test "tasks host" {
  vcsim_env

  run govc tasks 'host/*'
  assert_success
}
