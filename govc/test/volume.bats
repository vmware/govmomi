#!/usr/bin/env bats

load test_helper

@test "volume.ls" {
  vcsim_env

  run govc volume.ls
  assert_success ""
}
