#!/usr/bin/env bats

load test_helper

@test "network dvs" {
  vcsim_env

  run govc dvs.create -discovery-protocol cdp -product-version 6.6.0 -mtu 1500 DVS1
  assert_success

  dvs=$(govc object.collect -o -json network/DVS1)

  assert_equal cdp "$(jq -r .config.linkDiscoveryProtocolConfig.protocol <<<"$dvs")"
  assert_equal 1500 "$(jq -r .config.maxMtu <<<"$dvs")"
  assert_equal 6.6.0 "$(jq -r .summary.productInfo.version <<<"$dvs")"

  run govc dvs.add -dvs DVS1 DC0_H0
  assert_success

  run govc events -type DvsHostJoinedEvent
  assert_success
  assert_matches "DC0_H0 joined the vSphere Distributed Switch DVS1"
}

@test "network dvs backing" {
  vcsim_env

  # DVS backed network by default (from vcsim_env)
  vm=$(new_empty_vm)

  eth0=$(govc device.ls -vm $vm | grep ethernet- | awk '{print $1}')
  run govc device.info -vm $vm $eth0
  assert_success

  summary=$(govc device.info -vm $vm $eth0 | grep Summary: | awk '{print $2}')
  assert_equal "DVSwitch:" $summary

  run govc device.remove -vm $vm $eth0
  assert_success

  eth0=$(govc device.ls -vm $vm | grep ethernet- | awk '{print $1}')
  [ -z "$eth0" ]

  # Standard network backing
  run govc vm.network.add -vm $vm -net "VM Network"
  assert_success

  eth0=$(govc device.ls -vm $vm | grep ethernet- | awk '{print $1}')

  run govc device.info -vm $vm $eth0
  assert_success

  summary=$(govc device.info -vm $vm $eth0 | grep Summary: | awk -F: '{print $2}')
  assert_equal "VM Network" "$(collapse_ws $summary)"

  run govc device.remove -vm $vm $eth0
  assert_success

  run govc device.remove -vm $vm $eth0
  assert_failure "govc: device '$eth0' not found"

  # Test PG's with the same name
  run govc dvs.create DVS1 # DVS0 already exists
  assert_success

  run govc dvs.portgroup.add -dvs DVS0 -type ephemeral NSX-dvpg
  assert_success

  uuid=$(govc object.collect -s network/NSX-dvpg config.logicalSwitchUuid)
  sid=$(govc object.collect -s network/NSX-dvpg config.segmentId)
  moid=$(govc ls -i network/NSX-dvpg)

  run govc dvs.portgroup.add -dvs DVS1 -type ephemeral NSX-dvpg
  assert_success

  run govc vm.network.add -vm $vm -net NSX-dvpg
  assert_failure # resolves to multiple networks

  run govc vm.network.add -vm $vm -net DVS0/NSX-dvpg
  assert_success # switch_name/portgroup_name is unique

  # Add a 2nd PG to the same switch, with the same name
  run govc dvs.portgroup.add -dvs DVS0 -type ephemeral NSX-dvpg
  assert_success

  run govc vm.network.add -vm $vm -net NSX-dvpg
  assert_failure # resolves to multiple networks

  run govc vm.network.add -vm $vm -net DVS0/NSX-dvpg
  assert_failure # switch_name/portgroup_name not is unique

  run govc vm.network.add -vm $vm -net "$uuid"
  assert_success # switch uuid is unique

  run govc vm.network.add -vm $vm -net "$sid"
  assert_success # segment id is unique

  run govc vm.network.add -vm $vm -net "$moid"
  assert_success # moid is unique
}

@test "network change backing" {
  vcsim_env

  vm=$(new_empty_vm)

  eth0=$(govc device.ls -vm $vm | grep ethernet- | awk '{print $1}')
  run govc vm.network.change -vm $vm $eth0 enoent
  assert_failure "govc: network 'enoent' not found"

  run govc vm.network.change -vm $vm enoent "VM Network"
  assert_failure "govc: device 'enoent' not found"

  run govc vm.network.change -vm $vm $eth0 "VM Network"
  assert_success

  run govc vm.network.change -vm $vm $eth0
  assert_success

  unset GOVC_NETWORK
  run govc vm.network.change -vm $vm $eth0
  assert_failure "govc: default network resolves to multiple instances, please specify"

  run govc vm.power -on $vm
  assert_success
  run govc vm.power -off $vm

  mac=$(vm_mac $vm)
  run govc vm.network.change -vm $vm -net "VM Network" $eth0
  assert_success

  # verify we didn't change the mac address
  run govc vm.power -on $vm
  assert_success
  assert_equal $mac $(vm_mac $vm)
}

@test "network standard backing" {
  vcsim_env

  vm=$(new_empty_vm)

  run govc device.info -vm $vm ethernet-*
  assert_success

  run govc device.remove -vm $vm ethernet-*
  assert_success

  run govc device.info -vm $vm ethernet-*
  assert_failure

  run govc vm.network.add -vm $vm enoent
  assert_failure "govc: network 'enoent' not found"

  run govc vm.network.add -vm $vm "VM Network"
  assert_success

  run govc device.info -vm $vm ethernet-*
  assert_success

  dups=$(govc vm.info -json '*' | jq -r '.virtualMachines[].config.hardware.device[].macAddress | select(. != null)' | uniq -d)
  if [ -n "$dups" ] ; then
    flunk "duplicate MACs: $dups"
  fi
}

@test "network adapter" {
  vcsim_env

  vm=$(new_id)
  run govc vm.create -on=false -net.adapter=enoent $vm
  assert_failure "govc: unknown ethernet card type 'enoent'"

  vm=$(new_id)
  run govc vm.create -on=false -net.adapter=vmxnet3 $vm
  assert_success

  eth0=$(govc device.ls -vm $vm | grep ethernet- | awk '{print $1}')
  type=$(govc device.info -vm $vm $eth0 | grep Type: | awk -F: '{print $2}')
  assert_equal "VirtualVmxnet3" $(collapse_ws $type)

  run govc vm.network.add -vm $vm -net.adapter e1000e "VM Network"
  assert_success

  eth1=$(govc device.ls -vm $vm | grep ethernet- | grep -v $eth0 | awk '{print $1}')
  type=$(govc device.info -vm $vm $eth1 | grep Type: | awk -F: '{print $2}')
  assert_equal "VirtualE1000e" $(collapse_ws $type)

  # validate each NIC has a unique MAC
  macs=$(govc device.info -vm "$vm" -json ethernet-* | jq -r .devices[].macAddress | uniq | wc -l)
  assert_equal 2 "$macs"

  # validate -net.protocol. VM Network not compatible with vmxnet3vrdma, so create on dvgp under existing DVS0
  run govc dvs.portgroup.add -dvs DVS0 -type ephemeral NSX-dvpg
  assert_success

  # add a valid vmxnet3vrdma adapter with valid protocal
  run govc vm.network.add -vm $vm -net.adapter vmxnet3vrdma -net "DVS0/NSX-dvpg" -net.protocol=rocev2
  assert_success

  # add a valid vmxnet3vrdma adapter with valid protocal
  run govc vm.network.add -vm $vm -net.adapter vmxnet3vrdma -net "DVS0/NSX-dvpg" -net.protocol=rocev1
  assert_success

  # invalid value for -net.protocol
  run govc vm.network.add -vm $vm -net.adapter vmxnet3vrdma -net "DVS0/NSX-dvpg" -net.protocol=what
  assert_failure "govc: invalid device protocol 'what'"

  # invalid combination for -net.adapter and -net.protocol
  run govc vm.network.add -vm $vm -net.adapter e1000e -net "DVS0/NSX-dvpg" -net.protocol=rocev2
  assert_failure "govc: device protocol is only supported for vmxnet3vrdma at the moment"
}

@test "network flag required" {
  vcsim_env

  # -net flag is required when there are multiple networks
  unset GOVC_NETWORK
  run govc vm.create -on=false $(new_id)
  assert_failure "govc: default network resolves to multiple instances, please specify"
}

@test "network change hardware address" {
  vcsim_env -esx

  mac="00:00:0f$(dd bs=1 count=3 if=/dev/random 2>/dev/null | hexdump -v -e '/1 ":%02x"')"
  vm=$(new_id)
  run govc vm.create -on=false $vm
  assert_success

  run govc vm.network.change -vm $vm -net.address $mac ethernet-0
  assert_success

  run govc vm.power -on $vm
  assert_success

  assert_equal $mac $(vm_mac $vm)
}

@test "dvs.portgroup" {
  vcsim_env
  id=$(new_id)

  run govc dvs.create "$id"
  assert_success

  run govc events -type DvsCreatedEvent
  assert_success
  assert_matches "vSphere Distributed Switch $id was created"

  run govc events -type DVPortgroupCreatedEvent
  assert_success
  assert_matches "was added to switch"

  local host=$GOVC_HOST

  run govc dvs.add -dvs "$id" "$host"
  assert_success

  run govc dvs.portgroup.add -dvs "$id" -type earlyBinding -nports 16 "${id}-ExternalNetwork"
  assert_success

  nports=$(govc object.collect -s "network/${id}-ExternalNetwork" portKeys | awk -F, '{print NF}')
  [ "$nports" = "16" ]

  run govc dvs.portgroup.add -dvs "$id" -type ephemeral -vlan 3122 "${id}-InternalNetwork"
  assert_success

  info=$(govc dvs.portgroup.info "$id" | grep VlanId: | uniq | grep 3122)
  [ -n "$info" ]

  run govc dvs.portgroup.change -vlan 3123 "${id}-InternalNetwork"
  assert_success

  info=$(govc dvs.portgroup.info "$id" | grep VlanId: | uniq | grep 3123)
  [ -n "$info" ]

  info=$(govc dvs.portgroup.info -json "$id" | jq  '.port[].config.setting.vlan | select(.vlanId == 3123)')
  [ -n "$info" ]

  info=$(govc dvs.portgroup.info -json "$id" | jq  '.port[].config.setting.vlan | select(.vlanId == 7777)')
  [ -z "$info" ]

  run govc object.destroy "network/${id}-ExternalNetwork" "network/${id}-InternalNetwork" "network/${id}"
  assert_success

  run govc events -type DvsDestroyedEvent
  assert_success
  assert_matches "vSphere Distributed Switch $id in DC0 was deleted"

  run govc events -type DVPortgroupDestroyedEvent
  assert_success
  [ ${#lines[@]} -eq 2 ]
}
