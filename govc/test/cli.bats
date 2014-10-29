#!/usr/bin/env bats

load test_helper

@test "about" {
  run govc about
  assert_success
  assert_line "Vendor:       VMware, Inc."
}

@test "login attempt without credentials" {
  run govc about -u $(echo $GOVC_URL | awk -F@ '{print $2}')
  assert_failure "Error: ServerFaultCode: Cannot complete login due to an incorrect user name or password."
}
