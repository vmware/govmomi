#!/usr/bin/env bats

load test_helper

@test "compute.policy" {
  vcsim_env

  run govc compute.policy.ls
  assert_success ""

  run govc compute.policy.ls -c
  assert_success

  run govc compute.policy.ls -c -json
  assert_success

  capability=$(jq -r .[0].capability <<<"$output")

  run govc compute.policy.create -n my-policy -d my-desc
  assert_failure

  run govc compute.policy.create -n my-policy -d my-desc enoent
  assert_failure

  run govc compute.policy.create -n my-policy -d my-desc "$capability"
  assert_failure

  run govc compute.policy.create -n my-policy -d my-desc -host my-tag "$capability"
  assert_success
  id="$output"

  run govc compute.policy.ls
  assert_success
  assert_matches "$id"

  run govc compute.policy.rm enoent
  assert_failure

  run govc compute.policy.rm "$id"
  assert_success
}
