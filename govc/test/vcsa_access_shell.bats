#!/usr/bin/env bats

load test_helper

@test "vcsa.access.shell.get" {
  vcsim_env
  local output

  run govc vcsa.access.shell.get
  assert_success

  run govc vcsa.access.shell.get -json=True
  assert_success

  run govc vcsa.access.shell.get -xml=True
  assert_success
}

@test "vcsa.access.shell.set" {
  vcsim_env
  local output

  run govc vcsa.access.shell.set
  status=$(govc vcsa.access.shell.get| grep -ow "false"| awk '{print $1}')
  assert_equal $status 'false'

  run govc vcsa.access.shell.set -enabled=true -timeout=240
  status=$(govc vcsa.access.shell.get | grep -ow "true"| awk '{print $1}')
  assert_equal $status 'true'

  run govc vcsa.access.shell.set -enabled=false
  status=$(govc vcsa.access.shell.get | grep -ow "false"| awk '{print $1}')
  assert_equal $status 'false'
}
