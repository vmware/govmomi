#!/usr/bin/env bats

load test_helper

@test "vsan.change" {
  vcsim_env -cluster 2

  run govc vsan.info -json DC0_C0
  assert_success
  config=$(jq .Clusters[].Info.UnmapConfig <<<"$output")
  assert_equal null "$config"

  run govc vsan.change DC0_C0
  assert_failure # no flags specified

  run govc vsan.change -unmap-enabled DC0_C0
  assert_success

  run govc vsan.info -json DC0_C0
  assert_success

  config=$(jq .Clusters[].Info.UnmapConfig.Enable <<<"$output")
  assert_equal true "$config"
}
