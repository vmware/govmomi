#!/usr/bin/env bats

load test_helper

# These tests should only run against a server running an evaluation license.
verify_evaluation() {
  if [ "$(govc license.ls -json | jq -r .[0].editionKey)" != "eval" ]; then
    skip "requires evaluation license"
  fi
}

get_key() {
  jq ".[] | select(.licenseKey == \"$1\")"
}

get_property() {
  jq -r ".properties[] | select(.key == \"$1\") | .value"
}

get_label() {
  govc license.ls -json | jq -r ".[] | select(.licenseKey == \"$1\") | .labels[] | select(.key == \"$2\") | .value"
}

get_nlabel() {
  govc license.ls -json | jq ".[] | select(.licenseKey == \"$1\") | .labels[].key" | wc -l
}

@test "license.add" {
  esx_env

  verify_evaluation

  run govc license.add -json 00000-00000-00000-00000-00001 00000-00000-00000-00000-00002
  assert_success

  # Expect to see an entry for both the first and the second key
  assert_equal "License is not valid for this product" "$(get_key 00000-00000-00000-00000-00001 <<<${output} | get_property diagnostic)"
  assert_equal "License is not valid for this product" "$(get_key 00000-00000-00000-00000-00002 <<<${output} | get_property diagnostic)"
}

@test "license.assign" {
  vcsim_env

  run govc license.assign -cluster DC0_C0 00000-00000-00000-00000-00000
  assert_success

  run govc license.assigned.ls
  assert_success
}

@test "license.remove" {
  vcsim_env

  verify_evaluation

  run govc license.remove -json 00000-00000-00000-00000-00001
  assert_success
}

@test "license.ls" {
  vcsim_env

  verify_evaluation

  run govc license.ls -json
  assert_success

  # Expect the test instance to run in evaluation mode
  mode="$(get_key 00000-00000-00000-00000-00000 <<<"$output" | jq -r ".name")"
  assert_equal "Evaluation Mode" "$mode"

  name=$(jq -r '.[].properties[] | select(.key == "ProductName") | .value' <<<"$output")
  assert_equal "$(govc about -json | jq -r .about.licenseProductName)" "$name"

  name=$(jq -r '.[].properties[] | select(.key == "ProductVersion") | .value' <<<"$output")
  assert_equal "$(govc about -json | jq -r .about.licenseProductVersion)" "$name"
}

@test "license.decode" {
  vcsim_env

  verify_evaluation

  key=00000-00000-00000-00000-00000
  assert_equal "eval" $(govc license.decode $key | grep $key | awk '{print $2}')
}

@test "license.label.set" {
  vcsim_env

  key=00000-00000-00000-00000-00000

  assert_equal 0 "$(get_nlabel $key)"
  assert_equal "" "$(get_label $key foo)"

  run govc license.label.set $key foo bar
  assert_success

  assert_equal 1 "$(get_nlabel $key)"
  assert_equal bar "$(get_label $key foo)"

  run govc license.label.set $key biz baz
  assert_success
  run govc license.label.set $key foo bar2
  assert_success

  assert_equal 2 "$(get_nlabel $key)"
  assert_equal bar2 "$(get_label $key foo)"

  run govc license.label.set $key foo ""
  assert_success

  assert_equal 1 "$(get_nlabel $key)"
  assert_equal "" "$(get_label $key foo)"
}
