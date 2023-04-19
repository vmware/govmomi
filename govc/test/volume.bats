#!/usr/bin/env bats

load test_helper

@test "volume.ls" {
  vcsim_env

  run govc volume.ls
  assert_success ""
}

@test "volume.snapshot" {
  vcsim_env

  run govc volume.snapshot.ls
  assert_failure

  run govc volume.snapshot.rm
  assert_failure

  run govc volume.snapshot.create
  assert_failure
}
