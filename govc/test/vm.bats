#!/usr/bin/env bats

load test_helper

@test "vm.ip" {
  id=$(new_ttylinux_vm)

  run govc vm.power -on $id
  assert_success

  run govc vm.ip $id
  assert_success
}

@test "vm.ip -esxcli" {
  id=$(new_ttylinux_vm)

  run govc vm.power -on $id
  assert_success

  run govc vm.ip -esxcli $id
  assert_success

  ip_esxcli=$output

  run govc vm.ip $id
  assert_success
  ip_tools=$output

  assert_equal $ip_esxcli $ip_tools
}

@test "vm.create in cluster" {
  vcsim_env

  # using GOVC_HOST and its resource pool
  run govc vm.create -on=false $(new_id)
  assert_success

  # using no -host and the default resource pool for DC0
  unset GOVC_HOST
  run govc vm.create -on=false $(new_id)
  assert_success
}
