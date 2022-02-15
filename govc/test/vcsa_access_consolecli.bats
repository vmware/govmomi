#!/usr/bin/env bats

load test_helper

@test "vcsa.access.consolecli.get" {
  vcsim_env
  local output

  run govc vcsa.access.consolecli.get
  assert_success

  run govc vcsa.access.consolecli.get -json=True
  assert_success

  run govc vcsa.access.consolecli.get -xml=True
  assert_success
}

@test "vcsa.access.consolecli.set" {
  vcsim_env
  local output

  run govc vcsa.access.consolecli.set
  status=$(govc vcsa.access.consolecli.get| grep "false"| awk '{print $1}')
  assert_equal $status 'false'

  run govc vcsa.access.consolecli.set -enabled=true
  status=$(govc vcsa.access.consolecli.get | grep "true"| awk '{print $0}')
  assert_equal $status 'true'

  run govc vcsa.access.consolecli.set -enabled=false
  status=$(govc vcsa.access.consolecli.get | grep "false"| awk '{print $1}')
  assert_equal $status 'false'

}
