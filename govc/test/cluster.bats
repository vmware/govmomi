#!/usr/bin/env bats

load test_helper

@test "cluster.group" {
  vcsim_env -cluster 2 -host 4 -vm 8

  run govc cluster.group.ls -cluster DC0_C0
  assert_success "" # no groups

  run govc cluster.group.ls -cluster DC0_C0 -name my_vm_group
  assert_failure # group does not exist

  run govc cluster.group.create -cluster DC0_C0 -name my_vm_group -vm DC0_C0_H{0,1}
  assert_failure # -vm or -host required

  run govc cluster.group.create -cluster DC0_C0 -name my_vm_group -vm DC0_C0_H{0,1}
  assert_failure # -vm with HostSystem type args

  run govc cluster.group.create -cluster DC0_C0 -name my_vm_group -vm DC0_C0_RP0_VM{0,1}
  assert_success

  run govc cluster.group.ls -cluster DC0_C0
  assert_success "my_vm_group"

  run govc cluster.group.create -cluster DC0_C0 -name my_vm_group -vm DC0_C0_RP0_VM{0,1}
  assert_failure # group exists

  run govc cluster.group.ls -cluster DC0_C0 -name my_vm_group
  assert_success "$(printf "%s\n" DC0_C0_RP0_VM{0,1})"

  run govc cluster.group.change -cluster DC0_C0 -name my_vm_group DC0_C0_RP0_VM{0,1,2}
  assert_success

  run govc cluster.group.ls -cluster DC0_C0 -name my_vm_group
  assert_success "$(printf "%s\n" DC0_C0_RP0_VM{0,1,2})"

  run govc cluster.group.create -cluster DC0_C0 -name my_host_group -host DC0_C0_RP0_VM{0,1}
  assert_failure # -host with VirtualMachine type args

  run govc cluster.group.create -cluster DC0_C0 -name my_host_group -host DC0_C0_H{0,1}
  assert_success

  run govc cluster.group.ls -cluster DC0_C0 -name my_host_group
  assert_success

  run govc cluster.group.remove -cluster DC0_C0 -name my_vm_group
  assert_success

  run govc cluster.group.remove -cluster DC0_C0 -name my_vm_group
  assert_failure # group does not exist

  run govc cluster.group.ls -cluster DC0_C0
  assert_success "my_host_group"
}

@test "cluster.rule" {
  vcsim_env -cluster 2 -host 4 -vm 8

  run govc cluster.rule.ls -cluster DC0_C0
  assert_success "" # no rules

  run govc object.collect -json /DC0/host/DC0_C0 configurationEx.rule
  assert_success

  run govc cluster.rule.ls -cluster DC0_C0 -name pod1
  assert_failure # rule does not exist

  run govc cluster.rule.create -cluster DC0_C0 -name pod1 -affinity DC0_C0_RP0_VM0
  assert_failure # requires >= 2 VMs

  run govc cluster.rule.create -cluster DC0_C0 -name pod1 -affinity DC0_C0_RP0_VM{0,1,2,3}
  assert_success

  run govc cluster.rule.ls -cluster DC0_C0 -name pod1
  assert_success "$(printf "%s\n" DC0_C0_RP0_VM{0,1,2,3})"

  run govc cluster.rule.change -cluster DC0_C0 -name pod1 DC0_C0_RP0_VM{2,3,4}
  assert_success

  run govc cluster.rule.ls -cluster DC0_C0 -name pod1
  assert_success "$(printf "%s\n" DC0_C0_RP0_VM{2,3,4})"

  run govc object.collect -json /DC0/host/DC0_C0 configurationEx.rule
  assert_success

  run govc cluster.group.create -cluster DC0_C0 -name my_vms -vm DC0_C0_RP0_VM{0,1,2,3}
  assert_success

  run govc cluster.group.create -cluster DC0_C0 -name even_hosts -host DC0_C0_H{0,2}
  assert_success

  run govc cluster.group.create -cluster DC0_C0 -name odd_hosts -host DC0_C0_H{1,3}
  assert_success

  run govc cluster.rule.create -cluster DC0_C0 -name pod2 -enable -mandatory -vm-host -vm-group my_vms -host-affine-group even_hosts -host-anti-affine-group odd_hosts
  assert_success

  run govc cluster.rule.remove -cluster DC0_C0 -name pod1
  assert_success

  run govc cluster.rule.remove -cluster DC0_C0 -name pod1
  assert_failure # rule does not exist
}
