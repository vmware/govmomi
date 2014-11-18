#!/usr/bin/env bats

load test_helper

get_key() {
  jq ".[] | select(.LicenseKey == \"$1\")"
}

get_property() {
  jq -r ".Properties[] | select(.Key == \"$1\") | .Value"
}

@test "license.add" {
  run govc license.add -json 00000-00000-00000-00000-00001 00000-00000-00000-00000-00002
  assert_success

  # Expect to see an entry for both the first and the second key
  assert_equal "License is not valid for this product" $(get_key 00000-00000-00000-00000-00001 <<<${output} | get_property diagnostic)
  assert_equal "License is not valid for this product" $(get_key 00000-00000-00000-00000-00002 <<<${output} | get_property diagnostic)
}

@test "license.remove" {
  run govc license.remove -json 00000-00000-00000-00000-00001
  assert_success
}

@test "license.list" {
  run govc license.list -json
  assert_success

  # Expect the test instance to run in evaluation mode
  assert_equal "Evaluation Mode" $(get_key 00000-00000-00000-00000-00000 <<<$output | jq -r ".Name")
}
