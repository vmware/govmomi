#!/usr/bin/env bats

load test_helper

@test "vcsa.access.ssh.get" {
  vcsim_env
  local output

  run govc vcsa.access.ssh.get
  assert_success

  run govc vcsa.access.ssh.get -json=True
  assert_success

  run govc vcsa.access.ssh.get -xml=True
  assert_success
}

@test "vcsa.access.ssh.set" {
  vcsim_env
  local output

  run govc vcsa.access.ssh.set
  status=$(govc vcsa.access.ssh.get| grep -ow "false"| awk '{print $1}')
  assert_equal $status 'false'

  run govc vcsa.access.ssh.set -enabled=true
  status=$(govc vcsa.access.ssh.get | grep -ow "true"| awk '{print $1}')
  assert_equal $status 'true'

  run govc vcsa.access.ssh.set -enabled=false
  status=$(govc vcsa.access.ssh.get | grep -ow "false"| awk '{print $1}')
  assert_equal $status 'false'
}
