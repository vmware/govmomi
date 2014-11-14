#!/usr/bin/env bats

load test_helper

@test "about" {
  run govc about
  assert_success
  assert_line "Vendor: VMware, Inc."
}

@test "login attempt without credentials" {
  run govc about -u $(echo $GOVC_URL | awk -F@ '{print $2}')
  assert_failure "Error: ServerFaultCode: Cannot complete login due to an incorrect user name or password."
}

@test "connect to an endpoint with a non-supported API version" {
  run env GOVC_MIN_API_VERSION=24.4 govc about
  assert grep -q "^Error: Require API version 24.4," <<<${output}
}
