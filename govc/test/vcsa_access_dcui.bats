#!/usr/bin/env bats

load test_helper

@test "vcsa.access.dcui.get" {
  vcsim_env
  local output

  run govc vcsa.access.dcui.get
  assert_success

  run govc vcsa.access.dcui.get -json=True
  assert_success

  run govc vcsa.access.dcui.get -xml=True
  assert_success
}

@test "vcsa.access.dcui.set" {
  vcsim_env
  local output

  run govc vcsa.access.dcui.set
  status=$(govc vcsa.access.dcui.get| grep -ow "false"| awk '{print $1}')
  assert_equal $status 'false'

  run govc vcsa.access.dcui.set -enabled=true
  status=$(govc vcsa.access.dcui.get | grep -ow "true"| awk '{print $1}')
  assert_equal $status 'true'

  run govc vcsa.access.dcui.set -enabled=false
  status=$(govc vcsa.access.dcui.get | grep -ow "false"| awk '{print $1}')
  assert_equal $status 'false'
  
}
