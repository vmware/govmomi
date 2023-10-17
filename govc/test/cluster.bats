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

  run govc cluster.group.ls -cluster DC0_C0 -l
  assert_success
  [ ${#lines[@]} -eq 2 ]

  run bash -c 'govc cluster.group.ls -cluster DC0_C0 -l | grep ClusterVmGroup'
  assert_success
  [ ${#lines[@]} -eq 1 ]

  run bash -c 'govc cluster.group.ls -cluster DC0_C0 -l | grep ClusterHostGroup'
  assert_success
  [ ${#lines[@]} -eq 1 ]

  run govc cluster.group.remove -cluster DC0_C0 -name my_vm_group
  assert_success

  run govc cluster.group.remove -cluster DC0_C0 -name my_vm_group
  assert_failure # group does not exist

  run govc cluster.group.ls -cluster DC0_C0
  assert_success "my_host_group"

  run govc cluster.group.ls -cluster DC0_C0 -l
  assert_success
  [ ${#lines[@]} -eq 1 ]
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

  run govc cluster.rule.ls -cluster DC0_C0
  assert_success "pod1"

  run govc cluster.rule.ls -cluster DC0_C0 -l=true
  assert_success "pod1 (ClusterAffinityRuleSpec)"

  run govc cluster.rule.ls -cluster DC0_C0 -name pod1
  assert_success "$(printf "%s\n" DC0_C0_RP0_VM{0,1,2,3})"

  run govc cluster.rule.ls -cluster DC0_C0 -name pod1 -l=true
  assert_success "$(printf "%s (VM)\n" DC0_C0_RP0_VM{0,1,2,3})"

  run govc cluster.rule.info -cluster DC0_C0
  assert_success "$(cat <<_EOF_
Rule: pod1
  Type: ClusterAffinityRuleSpec
  VM: DC0_C0_RP0_VM0
  VM: DC0_C0_RP0_VM1
  VM: DC0_C0_RP0_VM2
  VM: DC0_C0_RP0_VM3
_EOF_
)"

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

  run govc cluster.rule.ls -cluster DC0_C0 -l
  assert_success "pod2 (ClusterVmHostRuleInfo)"

  run govc cluster.rule.ls -cluster DC0_C0 -name pod2
  assert_success "$(printf "%s\n" {my_vms,even_hosts,odd_hosts})"

  run govc cluster.rule.ls -cluster DC0_C0 -name pod2 -l
  assert_success "$(printf "%s\n" {'my_vms (vmGroupName)','even_hosts (affineHostGroupName)','odd_hosts (antiAffineHostGroupName)'})"

  run govc cluster.rule.remove -cluster DC0_C0 -name pod1 -depends
  assert_failure # rule does not exist

  run govc cluster.rule.create -cluster DC0_C0 -name my_deps -depends
  assert_failure # requires 2 groups

  run govc cluster.group.create -cluster DC0_C0 -name my_app -vm DC0_C0_RP0_VM{4,5}
  assert_success

  run govc cluster.group.create -cluster DC0_C0 -name my_db -vm DC0_C0_RP0_VM{6,7}
  assert_success

  run govc cluster.rule.create -cluster DC0_C0 -name my_deps -depends my_app my_db
  assert_success

  run govc cluster.rule.ls -cluster DC0_C0 -l
  assert_success "$(printf "%s\n" {'pod2 (ClusterVmHostRuleInfo)','my_deps (ClusterDependencyRuleInfo)'})"

  run govc cluster.rule.ls -cluster DC0_C0 -name my_deps
  assert_success "$(printf "%s\n" {'my_app','my_db'})"

  run govc cluster.rule.ls -cluster DC0_C0 -name my_deps -l
  assert_success "$(printf "%s\n" {'my_app (VmGroup)','my_db (DependsOnVmGroup)'})"

  run govc cluster.rule.info -cluster DC0_C0
  assert_success "$(cat <<_EOF_
Rule: pod2
  Type: ClusterVmHostRuleInfo
  vmGroupName: my_vms
  affineHostGroupName even_hosts
  antiAffineHostGroupName odd_hosts
Rule: my_deps
  Type: ClusterDependencyRuleInfo
  VmGroup my_app
  DependsOnVmGroup my_db
_EOF_
)"

}

@test "cluster.vm" {
  vcsim_env -host 4 -vm 8

  run govc cluster.override.info
  assert_success "" # no overrides == empty output

  run govc cluster.override.change
  assert_failure # -vm required

  run govc cluster.override.change -vm DC0_C0_RP0_VM0
  assert_failure # no changes specified

  # DRS override
  query=".overrides[] | select(.name == \"DC0_C0_RP0_VM0\") | .drs.enabled"

  run govc cluster.override.change -vm DC0_C0_RP0_VM0 -drs-enabled=false
  assert_success
  [ "$(govc cluster.override.info -json | jq "$query")" == "false" ]

  run govc cluster.override.change -vm DC0_C0_RP0_VM0 -drs-enabled=true
  assert_success
  [ "$(govc cluster.override.info -json | jq "$query")" == "true" ]

  run govc cluster.override.change -vm DC0_C0_RP0_VM0 -drs-mode=manual
  assert_success

  # DAS override
  query=".overrides[] | select(.name == \"DC0_C0_RP0_VM0\") | .das.dasSettings.restartPriority"

  [ "$(govc cluster.override.info -json | jq -r "$query")" != "high" ]

  run govc cluster.override.change -vm DC0_C0_RP0_VM0 -ha-restart-priority high
  assert_success
  [ "$(govc cluster.override.info -json | jq -r "$query")" == "high" ]

  # Orchestration override
  query=".overrides[] | select(.name == \"DC0_C0_RP0_VM0\") | .orchestration.vmReadiness.postReadyDelay"

  run govc cluster.override.change -vm DC0_C0_RP0_VM0 -ha-additional-delay 60
  assert_success
  [ "$(govc cluster.override.info -json | jq -r "$query")" == "60" ]

  query=".overrides[] | select(.name == \"DC0_C0_RP0_VM0\") | .orchestration.vmReadiness.readyCondition"

  run govc cluster.override.change -vm DC0_C0_RP0_VM0 -ha-ready-condition poweredOn
  assert_success
  [ "$(govc cluster.override.info -json | jq -r "$query")" == "poweredOn" ]

  # remove overrides
  run govc cluster.override.remove -vm DC0_C0_RP0_VM0
  assert_success
  run govc cluster.override.info
  assert_success "" # no overrides == empty output
}

@test "cluster.add" {
    vcsim_env
    unset GOVC_HOST

    ip=$(govc object.collect -o -json host/DC0_C0/DC0_C0_H0 | jq -r .config.network.vnic[].spec.ip.ipAddress)
    assert_equal 127.0.0.1 "$ip"

    govc cluster.add -cluster DC0_C0 -hostname 10.0.0.42 -username user -password pass
    assert_success

    ip=$(govc object.collect -o -json host/DC0_C0/10.0.0.42 | jq -r .config.network.vnic[].spec.ip.ipAddress)

    assert_equal 10.0.0.42 "$ip"
    govc host.info -json '*' | jq -r .hostSystems[].config.network.vnic[].spec.ip
    name=$(govc host.info -json -host.ip 10.0.0.42 | jq -r .hostSystems[].name)
    assert_equal 10.0.0.42 "$name"
}

@test "cluster.usage" {
  vcsim_env -host 4

  run govc cluster.usage enoent
  assert_failure

  run govc cluster.usage DC0_C0
  assert_success

  memory=$(govc cluster.usage -json DC0_C0 | jq -r .memory.summary.usage)
  [ "$memory" = "34.3" ]
}

@test "cluster.stretch" {
  vcsim_env -host 4

  run govc cluster.stretch -witness DC0_H0 -first-fault-domain-hosts=DC0_C0_H1,DC0_C0_H2 -second-fault-domain-hosts DC0_C0_H2,DC0_C0_H3
  assert_failure # no cluster specified

  run govc cluster.stretch -witness DC0_H0 -first-fault-domain-hosts=DC0_C0_H1,DC0_C0_H2 DC0_C0
  assert_failure # no second-fault-domain-hosts specified

  run govc cluster.stretch -witness DC0_H0 -first-fault-domain-hosts=DC0_C0_H1,DC0_C0_H2 -second-fault-domain-hosts DC0_C0_H2,DC0_C0_H3 DC0_C0
  assert_success
}

@test "cluster.module" {
  vcsim_env
  local output

  run govc cluster.module.ls
  assert_success # no tags defined yet

  run govc cluster.module.ls -k=false
  assert_failure

  run govc cluster.module.ls -id enoent
  assert_failure # specific module does not exist

  run govc cluster.module.create -cluster DC0_C0
  assert_success

  id="$output"

  run govc cluster.module.ls -id $id
  assert_success

  vm="/DC0/vm/DC0_C0_RP0_VM0"

  run govc cluster.module.vm.add -id $id $vm
  assert_success

  count=$(govc cluster.module.ls -id $id | grep -c VirtualMachine)
  [ "$count" = "1" ]

  run govc cluster.module.vm.add -id $id $vm
  assert_failure # already a member

  run govc cluster.module.vm.rm -id $id $vm
  assert_success

  run govc cluster.module.ls -id $id
  [ -z "$output" ]

  run govc cluster.module.rm $id
  assert_success

  run govc cluster.module.ls -id $id
  assert_failure

  run govc cluster.module.rm "does_not_exist"
  assert_failure

  run govc cluster.module.rm --ignore-not-found "does_not_exist"
  assert_success

  run govc cluster.module.create -cluster DC0_C0
  assert_success

  id="$output"

  run govc cluster.module.create -cluster DC0_C0
  assert_success

  id2="$output"

  run govc cluster.module.rm --ignore-not-found "-" <<_EOF_
$id
does_not_exist
$id2
_EOF_
  assert_success

}

@test "cluster.mv" {
  vcsim_env -host 4 -cluster 2

  # start with 4 hosts in each cluster
  run govc find -type h /DC0/host/DC0_C0
  assert_success
  [ ${#lines[@]} -eq 4 ]

  run govc find -type h /DC0/host/DC0_C1
  assert_success
  [ ${#lines[@]} -eq 4 ]

  run govc cluster.mv -cluster DC0_C1 DC0_C0_H*
  assert_failure

  run govc host.maintenance.enter DC0_C0_H*
  assert_success

  # move 1 host from C0 to C1
  run govc cluster.mv -cluster DC0_C1 DC0_C0_H2
  assert_success

  run govc find -type h /DC0/host/DC0_C0
  assert_success
  [ ${#lines[@]} -eq 3 ]

  run govc find -type h /DC0/host/DC0_C1
  assert_success
  [ ${#lines[@]} -eq 5 ]

  # move remaining 3 hosts from C0 to C1
  run govc cluster.mv -cluster DC0_C1 DC0_C0_H{0,1,3}
  assert_success

  run govc find -type h /DC0/host/DC0_C0
  assert_success
  [ ${#lines[@]} -eq 0 ]

  run govc find -type h /DC0/host/DC0_C1
  assert_success
  [ ${#lines[@]} -eq 8 ]

  # move a standalone host into the cluster
  run govc cluster.mv -cluster DC0_C1 DC0_H0
  assert_success

  run govc find -type h /DC0/host/DC0_C1
  assert_success
  [ ${#lines[@]} -eq 9 ]

  run govc cluster.mv -cluster DC0_C1 DC0_C0_H*
  assert_failure # hosts are already in the cluster

  # TODO: vcsim's MoveIntoFolder_Task only supports moving from folders
  # # move a cluster host to a standalone host
  # run govc object.mv /DC0/host/DC0_C1/DC0_H0 /DC0/host
  # assert_success
}

@test "cluster.change" {
  vcsim_env

  run govc cluster.change -drs-enabled -ha-enabled -ha-admission-control-enabled=false /DC0/host/DC0_C0
  assert_success

  config=$(govc object.collect -o -json /DC0/host/DC0_C0 | jq .configurationEx)
  assert_equal true "$(jq -r .drsConfig.enabled <<<"$config")"
  assert_equal true "$(jq -r .dasConfig.enabled <<<"$config")"
  assert_equal false "$(jq -r .dasConfig.admissionControlEnabled <<<"$config")"
}
