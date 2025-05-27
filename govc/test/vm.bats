#!/usr/bin/env bats

load test_helper

@test "vm.ip" {
  vcsim_env -autostart=false

  id=/DC0/vm/DC0_H0_VM0

  mac=00:50:56:83:3a:5d
  run govc vm.customize -vm $id -mac $mac -ip 10.0.0.1 -netmask 255.255.0.0 -type Linux
  assert_success

  run govc vm.power -on $id
  assert_success

  run govc vm.ip -wait 5s $id
  assert_success

  run govc vm.ip -wait 5s -a -v4 $id
  assert_success

  run govc vm.ip -wait 5s -n $mac $id
  assert_success

  run govc vm.ip -wait 5s -n ethernet-0 $id
  assert_success

  ip=$(govc vm.ip -wait 5s $id)

  # add a second nic
  run govc vm.network.add -vm $id "VM Network"
  assert_success

  res=$(govc vm.ip -wait 5s -n ethernet-0 $id)
  assert_equal $ip $res
}

@test "vm.ip capital MAC" {
  vcsim_env -autostart=false

  id=/DC0/vm/DC0_H0_VM0

  mac=00:50:56:83:3A:5D
  run govc vm.customize -vm $id -mac $mac -ip 10.0.0.1 -netmask 255.255.0.0 -type Linux
  assert_success

  run govc vm.power -on $id
  assert_success

  run govc vm.ip -wait 5s $id
  assert_success

  run govc vm.ip -wait 5s -a -v4 $id
  assert_success

  run govc vm.ip -wait 5s -n $mac $id
  assert_success

  run govc vm.ip -wait 5s -n ethernet-0 $id
  assert_success

  ip=$(govc vm.ip -wait 5s $id)

  # add a second nic
  run govc vm.network.add -vm $id "VM Network"
  assert_success

  res=$(govc vm.ip -wait 5s -n ethernet-0 $id)
  assert_equal $ip $res
}

@test "vm.ip -esxcli" {
  esx_env

  ok=$(govc host.esxcli system settings advanced list -o /Net/GuestIPHack | grep ^IntValue: | awk '{print $2}')
  if [ "$ok" != "1" ] ; then
    skip "/Net/GuestIPHack=0"
  fi
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

@test "vm.create" {
  vcsim_start
  export GOVC_NETWORK=/DC0/network/DC0_DVPG0

  run govc cluster.create empty-cluster
  assert_success

  id=$(new_id)
  run govc vm.create -on=false "$id"
  assert_failure # -pool must be specified

  run govc vm.create -pool DC0_C0/Resources "$id"
  assert_success

  id=$(new_id)
  run govc vm.create -cluster enoent "$id"
  assert_failure # cluster does not exist

  run govc vm.create -cluster empty-cluster "$id"
  assert_failure # cluster has no hosts

  run govc vm.create -cluster DC0_C0 "$id"
  assert_success

  run govc vm.create -force -cluster DC0_C0 "$id"
  assert_success # create vm with the same name

  run govc vm.create -cluster DC0_C0 "my:vm"
  assert_success # vm has special characters (moref)

  run govc object.collect -s vm/my:vm name
  assert_success my:vm
}

@test "vm.change" {
  vcsim_env

  id=DC0_H0_VM0

  run govc vm.change -g ubuntu64Guest -m 1024 -c 2 -vm $id
  assert_success

  run govc vm.info $id
  assert_success
  assert_matches "buntu"
  assert_line "Memory: 1024MB"
  assert_line "CPU: 2 vCPU(s)"

  # test extraConfig
  run govc vm.change -e "guestinfo.a=1" -e "guestinfo.b=2" -vm $id
  assert_success

  run govc vm.info -e $id
  assert_success
  assert_line "guestinfo.a: 1"
  assert_line "guestinfo.b: 2"

  # test extraConfigFile
  run govc vm.change -f "guestinfo.c=this_is_not_an_existing.file" -vm $id
  assert_failure

  echo -n "3" > "$BATS_TMPDIR/extraConfigFile.conf"
  run govc vm.change -f "guestinfo.d=$BATS_TMPDIR/extraConfigFile.conf" -vm $id
  assert_success

  run govc vm.info -e $id
  assert_success
  assert_line "guestinfo.d: 3"

  run govc vm.change -sync-time-with-host=false -vm $id
  assert_success

  run govc vm.info -t $id
  assert_success
  assert_line "SyncTimeWithHost: false"

  run govc vm.change -sync-time-with-host=true -vm $id
  assert_success

  run govc vm.info -t $id
  assert_success
  assert_line "SyncTimeWithHost: true"

  run govc object.collect -s "vm/$id" config.memoryAllocation.reservation
  assert_success 0

  govc vm.change -vm "$id" -mem.reservation 1024

  run govc object.collect -s "vm/$id" config.memoryAllocation.reservation
  assert_success 1024

  run govc vm.change -annotation $$ -vm "$id"
  assert_success

  run govc object.collect -s "vm/$id" config.annotation
  assert_success $$

  uuid=$(vcsim uuidgen)
  run govc vm.change -vm $id -uuid "$uuid"
  assert_success
  run govc object.collect -s "vm/$id" config.uuid
  assert_success "$uuid"

  run govc vm.change -vm $id -managed-by com.vmware.govmomi.simulator
  assert_success

  run govc vm.info -json $id
  assert_success

  run govc object.collect -s "vm/$id" config.managedBy.extensionKey
  assert_success com.vmware.govmomi.simulator

  run govc vm.change -vm $id -managed-by -
  assert_success

  run govc object.collect -s "vm/$id" config.managedBy
  assert_success ""

  run govc vm.change -vm $id -migrate-encryption required -ft-encryption-mode ftEncryptionRequired
  assert_success

  run govc collect -s "vm/$id" config.migrateEncryption
  assert_success "required"

  run govc collect -s "vm/$id" config.FtEncryptionMode
  assert_success "ftEncryptionRequired"

  nid=$(new_id)
  run govc vm.change -name $nid -vm $id
  assert_success

  run govc vm.info $id
  [ ${#lines[@]} -eq 0 ]

  run govc vm.info $nid
  [ ${#lines[@]} -gt 0 ]
}

@test "vm.change vcsim" {
  vcsim_env

  run govc vm.change -vm DC0_H0_VM0 -latency fail
  assert_failure

  run govc object.collect -s vm/DC0_H0_VM0 config.latencySensitivity.level
  assert_success normal

  run govc vm.change -vm DC0_H0_VM0 -latency high
  assert_success

  run govc object.collect -s vm/DC0_H0_VM0 config.latencySensitivity.level
  assert_success high
}

@test "vm.power" {
  vcsim_env -autostart=false

  vm=DC0_H0_VM0

  run vm_power_state $vm
  assert_success "poweredOff"

  run govc vm.power $vm
  assert_failure

  run govc vm.power -on -off $vm
  assert_failure

  # off -> on
  run govc vm.power -on $vm
  assert_success
  run vm_power_state $vm
  assert_success "poweredOn"
  run govc vm.power -on $vm
  assert_failure # already powered on

  # on -> shutdown
  run govc vm.power -s $vm
  assert_success
  run vm_power_state $vm
  assert_success "poweredOff"
  run govc vm.power -off $vm
  assert_failure # already powered off
  run govc vm.power -on $vm
  assert_success

  # on -> suspended
  run govc vm.power -suspend $vm
  assert_success
  run vm_power_state $vm
  assert_success "suspended"
  run govc vm.power -suspend $vm
  assert_failure # already suspended

  # suspended -> on
  run govc vm.power -on $vm
  assert_success
  run vm_power_state $vm
  assert_success "poweredOn"

  # on -> standby
  run govc vm.power -standby $vm
  assert_success
  run vm_power_state $vm
  assert_success "suspended"
  run govc vm.power -standby $vm
  assert_failure # already suspended
}

@test "vm.power -on -M" {
  for esx in true false ; do
    vcsim_env -esx=$esx -autostart=false

    vms=($(govc find / -type m | sort))

    # All VMs are off with -autostart=false
    off=($(govc find / -type m -runtime.powerState poweredOff | sort))
    assert_equal "${vms[*]}" "${off[*]}"

    # Power on 1 VM to test that -M is idempotent
    run govc vm.power -on "${vms[0]}"
    assert_success

    run govc vm.power -on -M "${vms[@]}"
    assert_success

    # All VMs should be powered on now
    on=($(govc find / -type m -runtime.powerState poweredOn | sort))
    assert_equal "${vms[*]}" "${on[*]}"

    vcsim_stop
  done
}

@test "vm.power -force" {
  vcsim_env

  vm=$(new_id)
  govc vm.create $vm

  run govc vm.power -r $vm
  assert_failure

  run govc vm.power -r -force $vm
  assert_success

  run govc vm.power -s $vm
  assert_success

  run govc vm.power -off $vm
  assert_failure

  run govc vm.power -off -force $vm
  assert_success

  run govc vm.destroy $vm
  assert_success

  run govc vm.power -off $vm
  assert_failure

  run govc vm.power -off -force $vm
  assert_failure
}

@test "vm.destroy" {
  vcsim_env

  vm=$(new_id)
  govc vm.create $vm

  # destroy powers off vm before destruction
  run govc vm.destroy $vm
  assert_success

  run govc vm.destroy '*'
  assert_success

  run govc find / -type m
  assert_success "" # expect all VMs are gone
}

@test "vm.create pvscsi" {
  vcsim_env

  vm=$(new_id)
  govc vm.create -on=false -disk.controller pvscsi $vm

  result=$(govc device.ls -vm $vm | grep pvscsi- | wc -l)
  [ $result -eq 1 ]

  result=$(govc device.ls -vm $vm | grep lsilogic- | wc -l)
  [ $result -eq 0 ]

  vm=$(new_id)
  govc vm.create -on=false -disk.controller pvscsi -disk=1GB $vm
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

@test "vm.create -datastore-cluster" {
  vcsim_env -pod 1 -ds 3

  pod=/DC0/datastore/DC0_POD0
  id=$(new_id)

  run govc vm.create -disk 10M -datastore-cluster $pod "$id"
  assert_failure

  run govc object.mv /DC0/datastore/LocalDS_{1,2} $pod
  assert_success

  run govc vm.create -disk 10M -datastore-cluster $pod "$id"
  assert_success
}

@test "vm.info" {
  vcsim_env -esx

  local num=3

  local prefix=$(new_id)

  for x in $(seq $num)
  do
    local id="${prefix}-${x}"

    # If VM is not found: No output, exit code==0
    run govc vm.info $id
    assert_success
    [ ${#lines[@]} -eq 0 ]

    # If VM is not found (using -json flag): Valid json output, exit code==0
    run env GOVC_INDENT=false govc vm.info -json $id
    assert_success
    assert_line "{\"virtualMachines\":null}"

    run govc vm.info -dump $id
    assert_success

    run govc vm.create -on=false $id
    assert_success

    local info=$(govc vm.info -r $id)
    local found=$(grep Name: <<<"$info" | wc -l)
    [ "$found" -eq 1 ]

    # test that mo names are printed
    found=$(grep Host: <<<"$info" | awk '{print $2}')
    [ -n "$found" ]
    found=$(grep Storage: <<<"$info" | awk '{print $2}')
    [ -n "$found" ]
    found=$(grep Network: <<<"$info" | awk '{print $2}')
    [ -n "$found" ]
  done

  # test find slice
  local slice=$(govc vm.info ${prefix}-*)
  local found=$(grep Name: <<<"$slice" | wc -l)
  [ "$found" -eq $num ]

  # test -r
  found=$(grep Storage: <<<"$slice" | wc -l)
  [ "$found" -eq 0 ]
  found=$(grep Network: <<<"$slice" | wc -l)
  [ "$found" -eq 0 ]
  slice=$(govc vm.info -r ${prefix}-*)
  found=$(grep Storage: <<<"$slice" | wc -l)
  [ "$found" -eq $num ]
  found=$(grep Network: <<<"$slice" | wc -l)
  [ "$found" -eq $num ]

  # test extraConfig
  run govc vm.change -e "guestinfo.a=2" -vm $id
  assert_success
  run govc vm.info -e $id
  assert_success
  assert_line "guestinfo.a: 2"
  run govc vm.change -e "guestinfo.a=" -vm $id
  assert_success
  refute_line "guestinfo.a: 2"

  # test extraConfigFile
  run govc vm.change -f "guestinfo.b=this_is_not_an_existing.file" -vm $id
  assert_failure
  echo -n "3" > "$BATS_TMPDIR/extraConfigFile.conf"
  run govc vm.change -f "guestinfo.b=$BATS_TMPDIR/extraConfigFile.conf" -vm $id
  assert_success
  run govc vm.info -e $id
  assert_success
  assert_line "guestinfo.b: 3"
  run govc vm.change -f "guestinfo.b=" -vm $id
  assert_success
  refute_line "guestinfo.b: 3"

  # test optional bool Config
  run govc vm.change -nested-hv-enabled=true -vm "$id"
  assert_success

  hv=$(govc vm.info -json "$id" | jq '.[][0].config.nestedHVEnabled')
  assert_equal "$hv" "true"
}

@test "vm.info multi dc" {
  vcsim_start -dc 2

  run govc vm.info /DC1/vm/DC1_H0_VM1
  assert_success

  run govc vm.info DC1_H0_VM1
  assert_failure

  run govc vm.info -vm.ipath /DC1/vm/DC1_H0_VM1
  assert_success
  uuid=$(grep "UUID:" <<<"$output" | awk '{print $2}')

  run govc vm.info -vm.ipath DC1_H0_VM1
  assert_failure

  run govc vm.info -vm.uuid enoent
  assert_failure

  run govc vm.info -vm.uuid "$uuid"
  assert_failure

  run govc vm.info -dc DC1 -vm.uuid "$uuid"
  assert_success
}

@test "vm.create linked ide disk" {
  esx_env

  import_ttylinux_vmdk

  vm=$(new_id)

  run govc vm.create -disk $GOVC_TEST_VMDK -disk.controller ide -on=false $vm
  assert_success

  run govc device.info -vm $vm disk-200-0
  assert_success
  assert_line "Controller: ide-200"
}

@test "vm.create linked scsi disk" {
  esx_env

  import_ttylinux_vmdk

  vm=$(new_id)

  run govc vm.create -disk enoent -on=false $vm
  assert_failure "govc: cannot stat '[${GOVC_DATASTORE##*/}] enoent': No such file"

  run govc vm.create -disk $GOVC_TEST_VMDK -on=false $vm
  assert_success

  run govc device.info -vm $vm disk-1000-0
  assert_success
  assert_line "Controller: lsilogic-1000"
  assert_line "Parent: [${GOVC_DATASTORE##*/}] $GOVC_TEST_VMDK"
  assert_line "File: [${GOVC_DATASTORE##*/}] $vm/${vm}.vmdk"
}

@test "vm.create scsi disk" {
  esx_env

  import_ttylinux_vmdk

  vm=$(new_id)

  run govc vm.create -disk enoent -on=false $vm
  assert_failure "govc: cannot stat '[${GOVC_DATASTORE##*/}] enoent': No such file"

  run govc vm.create -disk $GOVC_TEST_VMDK -on=false -link=false $vm
  assert_success

  run govc device.info -vm $vm disk-1000-0
  assert_success
  assert_line "Controller: lsilogic-1000"
  refute_line "Parent: [${GOVC_DATASTORE##*/}] $GOVC_TEST_VMDK"
  assert_line "File: [${GOVC_DATASTORE##*/}] $GOVC_TEST_VMDK"
}

@test "vm.create scsi disk with datastore argument" {
  esx_env

  import_ttylinux_vmdk

  vm=$(new_id)

  run govc vm.create -disk="${GOVC_TEST_VMDK}" -disk-datastore="${GOVC_DATASTORE}" -on=false -link=false $vm
  assert_success

  run govc device.info -vm $vm disk-1000-0
  assert_success
  assert_line "File: [${GOVC_DATASTORE##*/}] $GOVC_TEST_VMDK"
}

@test "vm.create iso" {
  vcsim_env -esx

  upload_iso

  vm=$(new_id)

  run govc vm.create -iso enoent -on=false $vm
  assert_failure "govc: cannot stat '[${GOVC_DATASTORE##*/}] enoent': No such file"

  run govc vm.create -iso $GOVC_TEST_ISO -on=false $vm
  assert_success

  run govc device.info -vm $vm cdrom-*
  assert_success
  assert_line "Type: VirtualCdrom"
  assert_line "Summary: ISO [${GOVC_DATASTORE##*/}] $GOVC_TEST_ISO"
}

@test "vm.create iso with datastore argument" {
  vcsim_env

  upload_iso

  vm=$(new_id)

  run govc vm.create -iso="${GOVC_TEST_ISO}" -iso-datastore="${GOVC_DATASTORE}" -on=false $vm
  assert_success

  run govc device.info -vm $vm cdrom-*
  assert_success
  assert_line "Summary: ISO [${GOVC_DATASTORE##*/}] $GOVC_TEST_ISO"
}

@test "vm.disk.create empty vm" {
  vcsim_env

  vm=$(new_empty_vm)

  local name=$(new_id)

  run govc vm.disk.create -vm "$vm" -name "$name" -size 1G
  assert_success
  result=$(govc device.ls -vm "$vm" | grep -c disk-)
  [ "$result" -eq 1 ]
  govc device.info -json -vm "$vm" disk-* | jq .devices[].backing.sharing | grep -v sharingMultiWriter

  name=$(new_id)

  run govc vm.disk.create -vm "$vm" -name "$vm/$name" -size 2G
  assert_success

  result=$(govc device.ls -vm "$vm" | grep -c disk-)
  [ "$result" -eq 2 ]
}

@test "vm.disk.share" {
  esx_env

  vm=$(new_empty_vm)

  run govc vm.disk.create -vm "$vm" -name "$vm/shared.vmdk" -size 1G -eager -thick -sharing sharingMultiWriter
  assert_success
  govc device.info -json -vm "$vm" disk-* | jq .devices[].backing.sharing | grep sharingMultiWriter

  run govc vm.power -on "$vm"
  assert_success

  vm2=$(new_empty_vm)

  run govc vm.disk.attach -vm "$vm2" -link=false -disk "$vm/shared.vmdk"
  assert_success

  run govc vm.power -on "$vm2"
  assert_failure # requires sharingMultiWriter

  run govc device.remove -vm "$vm2" -keep disk-1000-0
  assert_success

  run govc vm.disk.attach -vm "$vm2" -link=false -sharing sharingMultiWriter -disk "$vm/shared.vmdk"
  assert_success

  run govc vm.power -on "$vm2"
  assert_success

  run govc vm.power -off "$vm"
  assert_success

  run govc vm.disk.change -vm "$vm" -disk.filePath "[$GOVC_DATASTORE] $vm/shared.vmdk" -sharing sharingNone
  assert_success

  ! govc device.info -json -vm "$vm" disk-* | jq .devices[].backing.sharing | grep sharingMultiWriter
}

@test "vm.disk.create" {
  vcsim_env

  vm=$(new_id)

  govc vm.create -on=false "$vm"
  assert_success

  name=$(new_id)

  run govc vm.disk.create -vm "$vm" -name "$vm/$name" -size 1M
  assert_success
  disk=$(govc device.ls -vm "$vm" disk-* | awk '{print $1}')
  result=$(grep -c disk- <<<"$disk")
  [ "$result" -eq 1 ]

  run govc vm.disk.change -vm "$vm" -disk.name "$disk" -size 2M
  assert_success

  run govc vm.disk.change -vm "$vm" -disk.name "$disk" -size 1M
  assert_failure # cannot shrink disk

  name=$(new_id)
  run govc vm.disk.create -vm "$vm" -name "$vm/$name" -profile enoent
  assert_failure # profile does not exist

  run govc vm.disk.create -vm "$vm" -name "$vm/$name" -profile "vSAN Default Storage Policy"
  assert_success
}

@test "vm.disk.attach" {
  esx_env

  import_ttylinux_vmdk

  vm=$(new_id)

  govc vm.create -disk $GOVC_TEST_VMDK -on=false $vm
  result=$(govc device.ls -vm $vm | grep disk- | wc -l)
  [ $result -eq 1 ]

  id=$(new_id)
  run govc import.vmdk $GOVC_TEST_VMDK_SRC $id
  assert_success

  run govc vm.disk.attach -vm $vm -link=false -disk enoent.vmdk
  assert_failure "govc: File [${GOVC_DATASTORE##*/}] enoent.vmdk was not found"

  run govc vm.disk.attach -vm $vm -disk enoent.vmdk
  assert_failure "govc: Invalid configuration for device '0'."

  run govc vm.disk.attach -vm $vm -disk $id/$(basename $GOVC_TEST_VMDK) -controller lsilogic-1000
  assert_success
  result=$(govc device.ls -vm $vm | grep disk- | wc -l)
  [ $result -eq 2 ]
}

@test "vm.disk.promote" {
  vcsim_env

  export GOVC_VM=DC0_H0_VM0

  run govc vm.disk.promote
  assert_failure

  run govc vm.disk.promote invalid-disk-name
  assert_failure

  run govc device.info disk-*-0
  assert_success
  grep -v "Parent:" <<<"$output" # No parent disk

  run govc vm.disk.promote disk-*-0
  assert_success

  run govc disk.create -size 10M my-disk
  assert_success
  id="$output"

  run govc disk.ls -json "$id"
  assert_success

  run jq -r .objects[].config.backing.filePath <<<"$output"
  assert_success
  path="$output"

  run govc vm.disk.attach -link -disk "$path"
  assert_success

  run govc device.info disk-*-1
  assert_success
  assert_line "Parent: $path" # Has parent disk

  run govc vm.disk.promote disk-*-1
  assert_success

  run govc device.info disk-*-1
  assert_success
  grep -v "Parent:" <<<"$output" # No more parent disk
}

@test "vm.create new disk with datastore argument" {
  vcsim_env

  vm=$(new_id)

  run govc vm.create -disk="1GiB" -ds="${GOVC_DATASTORE}" -on=false -link=false $vm
  assert_success

  run govc device.info -vm $vm disk-*
  assert_success
  assert_line "File: [${GOVC_DATASTORE##*/}] ${vm}/${vm}.vmdk"
}

@test "vm.create new disk with datastore cluster argument" {
  vcsim_env -pod 1 -ds 3

  vm=$(new_id)

  run govc object.mv /DC0/datastore/LocalDS_{1,2} /DC0/datastore/DC0_POD0
  assert_success

  run govc vm.create -disk="1GiB" -datastore-cluster=/DC0/datastore/DC0_POD0 -on=false -link=false "$vm"
  assert_success

  run govc device.info -vm $vm disk-*
  assert_success
}

@test "vm.register" {
  vcsim_env

  run govc vm.unregister enoent
  assert_failure

  vm=$(new_empty_vm)

  run govc vm.unregister "$vm"
  assert_success

  run govc vm.register "$vm/${vm}.vmx"
  assert_success
}

@test "vm.register vcsim" {
  vcsim_env -autostart=false

  host=$GOVC_HOST
  pool=$GOVC_RESOURCE_POOL

  unset GOVC_HOST GOVC_RESOURCE_POOL

  vm=DC0_H0_VM0

  run govc vm.unregister $vm
  assert_success

  run govc vm.register "$vm/${vm}.vmx"
  assert_failure # -pool is required

  run govc vm.register -pool "$pool" "$vm/${vm}.vmx"
  assert_success

  run govc vm.unregister $vm
  assert_success

  run govc vm.register -template -pool "$pool" "$vm/${vm}.vmx"
  assert_failure # -pool is not allowed w/ template

  run govc vm.register -template -host "$host" "$vm/${vm}.vmx"
  assert_success
}

@test "vm.clone" {
  vcsim_env

  vm="DC0_H0_VM0"
  clone=$(new_id)

  run govc vm.clone -vm "$vm" -host.ipath /DC0/host/DC0_C0/DC0_C0_H0 -annotation $$ "$clone"
  assert_success

  backing=$(govc device.info -json -vm "$clone" disk-* | jq .devices[].backing)
  assert_equal false "$(jq .eagerlyScrub <<<"$backing")"
  assert_equal true "$(jq .thinProvisioned <<<"$backing")"

  run govc object.collect -s "/$GOVC_DATACENTER/vm/$clone" config.annotation
  assert_success $$

  clone=$(new_id)
  run govc vm.clone -vm "$vm" -snapshot X "$clone"
  assert_failure

  run govc snapshot.create -vm "$vm" X
  assert_success

  run govc vm.clone -vm "$vm" -snapshot X "$clone"
  assert_success

  clone=$(new_id)
  run govc vm.clone -cluster enoent -vm "$vm" "$clone"
  assert_failure

  run govc vm.clone -cluster DC0_C0 -vm "$vm" "$clone"
  assert_success

  run govc vm.clone -cluster DC0_C0 -vm "$vm" "$clone"
  assert_failure # already exists

  run govc datastore.cp "$clone"/"$clone".vmx "$clone"/"$clone".vmx.copy
  run govc vm.destroy "$clone"
  run govc datastore.mv "$clone"/"$clone".vmx.copy "$clone"/"$clone".vmx # leave vmx file
  run govc vm.clone -force -vm "$vm" "$clone"
  assert_success # clone vm with the same name vmx file

  vm=$(new_empty_vm)
  run govc vm.disk.create -vm "$vm" -thick -eager -size 10M -name "$vm/data.vmdk"
  assert_success

  backing=$(govc device.info -json -vm "$vm" disk-* | jq .devices[].backing)
  assert_equal true "$(jq .eagerlyScrub <<<"$backing")"
  assert_equal false "$(jq .thinProvisioned <<<"$backing")"

  clone=$(new_id)
  run govc vm.clone -vm "$vm" "$clone"
  assert_success

  backing=$(govc device.info -json -vm "$clone" disk-* | jq .devices[].backing)
  assert_equal true "$(jq .eagerlyScrub <<<"$backing")"
  assert_equal false "$(jq .thinProvisioned <<<"$backing")"

  # test that each vm has a unique vmdk path
  for item in fileName uuid;  do
    items=$(govc object.collect -json -type m / config.hardware.device | \
              jq ".changeSet[].val._value[].backing.$item | select(. != null)")

    nitems=$(wc -l <<<"$items")
    uitems=$(sort -u <<<"$items" | wc -l)
    assert_equal "$nitems" "$uitems"
  done
}

@test "vm.clone change resources" {
  vcsim_env

  vm=$(new_empty_vm)
  clone=$(new_id)

  run govc vm.info -r "$vm"
  assert_success
  assert_line "Network: $(basename "$GOVC_NETWORK")" # DVPG0

  run govc vm.clone -m 1024 -c 2 -net "VM Network" -vm "$vm" "$clone"
  assert_success

  run govc vm.info -r "$clone"
  assert_success
  assert_line "Memory: 1024MB"
  assert_line "CPU: 2 vCPU(s)"
  assert_line "Network: VM Network"

  # Remove all NICs from source vm
  run govc device.remove -vm "$vm" "$(govc device.ls -vm "$vm" | grep ethernet- | awk '{print $1}')"
  assert_success

  clone=$(new_id)

  mac=00:00:0f:a7:a0:f1
  run govc vm.clone -net "VM Network" -net.address $mac -vm "$vm" "$clone"
  assert_success

  run govc vm.info -r "$clone"
  assert_success
  assert_line "Network: VM Network"

  run govc device.info -vm "$clone"
  assert_success
  assert_line "MAC Address: $mac"
}

@test "vm.clone usage" {
  # validate we require -vm flag
  run govc vm.clone enoent
  assert_failure
}

@test "vm.migrate" {
  vcsim_env -cluster 2

  host0=/DC0/host/DC0_C0/DC0_C0_H0
  host1=/DC0/host/DC0_C0/DC0_C0_H1
  moid0=$(govc find -maxdepth 0 -i $host0)
  moid1=$(govc find -maxdepth 0 -i $host1)

  vm=$(new_id)
  run govc vm.create -on=false -host $host0 "$vm"
  assert_success

  # assert VM is on H0
  run govc object.collect "vm/$vm" -runtime.host "$moid0"
  assert_success

  # WaitForUpdates until the VM runtime.host changes to H1
  govc object.collect "vm/$vm" -runtime.host "$moid1" &
  pid=$!

  # migrate from H0 to H1
  run govc vm.migrate -host $host1 "$vm"
  assert_success

  run govc events -type VmMigratedEvent "vm/$vm"
  assert_success
  assert_matches "Migration of virtual machine"

  wait $pid

  # (re-)assert VM is now on H1
  run govc object.collect "vm/$vm" -runtime.host "$moid1"
  assert_success

  # migrate from C0 to C1
  run govc vm.migrate -pool DC0_C1/Resources "$vm"
  assert_success

  run govc folder.create vm/new-folder
  assert_success

  run govc object.collect -s "vm/$vm" parent
  assert_success

  uuid=$(govc object.collect -s "vm/$vm" config.uuid)
  run govc vm.migrate -folder vm/new-folder -vm.uuid "$uuid"
  assert_success

  run govc object.collect -s "vm/new-folder/$vm" parent
  assert_success

  run govc vm.info -r "$vm"
  assert_matches "Network: +DC0_DVPG0"

  run govc vm.migrate -host "$host0" -net "VM Network" "$vm"
  assert_success

  run govc vm.info -r "$vm"
  assert_matches "Network: +VM Network"
}

@test "object name with slash" {
  vcsim_env

  vm=DC0_H0_VM0

  name="$vm/with-slash"

  # rename VM to include a '/'
  run govc vm.change -vm "$vm" -name "$name"
  assert_success

  path=$(govc ls "vm/$name")

  run govc vm.info "$name"
  assert_success
  assert_line "Name: $name"
  assert_line "Path: $path"

  run govc vm.info "$path"
  assert_success
  assert_line "Name: $name"
  assert_line "Path: $path"

  run govc find vm -name "$name"
  assert_success "vm/$name"

  # create a portgroup where name includes a '/'
  net=$(new_id)/with-slash

  run govc host.portgroup.add -vswitch vSwitch0 "$net"
  assert_success

  run govc vm.network.change -vm "$name" -net "$net" ethernet-0
  assert_success

  # change VM eth0 to use network that includes a '/' in the name
  run govc device.info -vm "$name" ethernet-0
  assert_success
  assert_line "Summary: $net"

  run govc host.portgroup.remove "$net"
  assert_success
}

@test "vm.console" {
  esx_env

  vm=$(new_empty_vm)

  run govc vm.console "$vm"
  assert_success

  run govc vm.console -wss "$vm"
  assert_failure

  run govc vm.power -on "$vm"
  assert_success

  run govc vm.console "$vm"
  assert_success

  run govc vm.console -wss "$vm"
  assert_success

  run govc vm.console -capture - "$vm"
  assert_success
}

@test "vm.upgrade" {
  vcsim_env

  vm=$(new_id)

  run govc vm.create -on=false -version 0.5 "$vm"
  assert_failure

  run govc vm.create -on=false -version 5.5 "$vm"
  assert_success

  run govc object.collect -s "vm/$vm" config.version
  assert_success "vmx-10"

  run govc vm.upgrade -vm "$vm"
  assert_success

  version=$(govc object.collect -s "vm/$vm" config.version)
  [[ "$version" > "vmx-10" ]]

  run govc vm.upgrade -vm "$vm"
  assert_failure

  run govc vm.create -on=false -version vmx-11 "$(new_id)"
  assert_success
}

@test "vm.markastemplate" {
  vcsim_env

  id=$(new_id)

  run govc vm.create -on=true "$id"
  assert_success

  run govc vm.change -vm "$id" -e testing=123
  assert_success

  run govc vm.markastemplate "$id"
  assert_failure

  run govc vm.power -off "$id"
  assert_success

  run govc vm.markasvm "$id"
  assert_failure # already a vm

  run govc vm.markastemplate "$id"
  assert_success

  run govc vm.markastemplate "$id"
  assert_failure # already a template

  run govc vm.change -vm "$id" -e testing=456
  assert_failure # template reconfigure only allows name and annotation change

  run govc vm.change -vm "$id" -annotation testing123
  assert_success

  run govc vm.power -on "$id"
  assert_failure

  run govc vm.clone -vm "$id" -on=false new-vm
  assert_success

  run govc vm.markasvm "$id"
  assert_success
}

@test "vm.option.info" {
  vcsim_env

  run govc vm.option.info -host "$GOVC_HOST"
  assert_success

  run govc vm.option.info -cluster "$(dirname "$GOVC_HOST")"
  assert_success

  run govc vm.option.info -vm DC0_H0_VM0
  assert_success

  family=$(govc vm.option.info -json ubuntu64Guest | jq -r .guestOSDescriptor[].family)
  assert_equal linuxGuest "$family"

  family=$(govc vm.option.info -json windows8_64Guest | jq -r .guestOSDescriptor[].family)
  assert_equal windowsGuest "$family"

  run govc vm.option.info enoent
  assert_success  # returns the entire simulator.GuestID list
  [ ${#lines[@]} -ge 100 ]
}

@test "vm.target.info" {
  vcsim_env

  run govc vm.target.info -host "$GOVC_HOST"
  assert_success

  run govc vm.target.info -cluster "$(dirname "$GOVC_HOST")"
  assert_success

  run govc vm.target.info -vm DC0_H0_VM0
  assert_success

  run govc vm.target.info -json
  assert_success
}

@test "vm.customize" {
  vcsim_env

  run govc vm.customize -vm DC0_H0_VM0 -ip 10.0.0.42 -netmask 255.255.0.0 vcsim-linux-static
  assert_failure # power must be off

  run govc vm.power -off DC0_H0_VM0
  assert_success

  run govc vm.customize -vm DC0_H0_VM0 -ip 10.0.0.42 -netmask 255.255.0.0 vcsim-linux-static
  assert_success

  run govc vm.customize -vm DC0_H0_VM0 -ip 10.0.0.42 -netmask 255.255.0.0 vcsim-linux-static
  assert_failure # pending customization

  run govc vm.power -on DC0_H0_VM0
  assert_success

  run govc object.collect -s vm/DC0_H0_VM0 guest.ipAddress
  assert_success 10.0.0.42

  run govc object.collect -s vm/DC0_H0_VM0 guest.hostName
  assert_success vcsim-1

  run govc vm.power -off DC0_H0_VM0
  assert_success

  run govc vm.customize -vm DC0_H0_VM0 -ip 10.0.0.43 -netmask 255.255.0.0 -domain HOME -tz D -name windoze vcsim-windows-static
  assert_success

  run govc vm.power -on DC0_H0_VM0
  assert_success

  run govc object.collect -s vm/DC0_H0_VM0 guest.ipAddress
  assert_success 10.0.0.43

  run govc object.collect -s vm/DC0_H0_VM0 summary.guest.ipAddress
  assert_success 10.0.0.43

  run govc object.collect -s vm/DC0_H0_VM0 summary.guest.hostName
  assert_success windoze

  run govc vm.info DC0_H0_VM0
  assert_success
  assert_matches 10.0.0.43

  run govc object.collect -s vm/DC0_H0_VM0 guest.hostName
  assert_success windoze

  run govc vm.power -off DC0_H0_VM0
  assert_success

  run govc vm.customize -vm DC0_H0_VM0 -ip 10.0.0.44 -netmask 255.255.0.0 -type Windows
  assert_success

  run govc vm.power -on DC0_H0_VM0
  assert_success

  run govc object.collect -s vm/DC0_H0_VM0 guest.ipAddress
  assert_success 10.0.0.44

  run govc vm.power -off DC0_H0_VM0
  assert_success

  run govc vm.customize -vm DC0_H0_VM0 -type Linux
  assert_failure # no -ip specified

  run govc vm.customize -vm DC0_H0_VM0 -ip 10.0.0.45 -netmask 255.255.0.0 -type Linux -dns-server 1.1.1.1 -dns-suffix example.com
  assert_success

  run govc vm.power -on DC0_H0_VM0
  assert_success

  run govc object.collect -s vm/DC0_H0_VM0 guest.ipAddress
  assert_success 10.0.0.45

  host=$(govc ls -L "$(govc object.collect -s vm/DC0_H0_VM0 runtime.host)")
  run govc host.maintenance.enter "$host"
  assert_success

  run govc vm.power -off DC0_H0_VM0
  assert_success

  run govc vm.power -on DC0_H0_VM0
  assert_failure # InvalidState

  run govc vm.customize -vm DC0_H0_VM0 -ip 10.0.0.45 -netmask 255.255.0.0 -type Linux
  assert_failure # InvalidState

  run govc host.maintenance.exit "$host"
  assert_success

  run govc vm.customize -vm DC0_H0_VM0 -ip 10.0.0.45 -netmask 255.255.0.0 -type Linux
  assert_success
}

@test "vm.customize failures" {
  vcsim_env -autostart=false

  vm=DC0_H0_VM0

  run govc vm.network.add -vm $vm -net "VM Network"
  assert_success

  run govc vm.customize -vm $vm -ip 10.0.0.42 -netmask 255.255.0.0 -ip 10.0.0.43 -netmask 255.255.0.0
  assert_success

  run govc collect -s vm/$vm config.tools.pendingCustomization
  assert_success
  [ "$output" != "" ]

  run govc device.remove -vm $vm ethernet-0
  assert_success

  run govc vm.power -on $vm
  assert_success

  run govc collect -s vm/$vm config.tools.pendingCustomization
  assert_success
  [ "$output" == "" ]

  run govc events -type CustomizationNetworkSetupFailed vm/$vm
  assert_success
  assert_output_lines 1
}

@test "vm.check.config" {
  vcsim_env -cluster 2

  export GOVC_SHOW_UNRELEASED=true

  vm=DC0_C0_RP0_VM0
  run govc vm.create -spec -pool DC0_C0/Resources $vm
  assert_success
  spec="$output"

  run govc vm.check.config -vm $vm <<<"$spec"
  assert_success

  # we can use any host in the VM's cluster
  for host in $(govc find /DC0/host/DC0_C0 -type f) ; do
    run govc vm.check.config -host "$host" <<<"$spec"
    assert_success
  done

  run govc vm.check.config -host DC0_C1_H0 <<<"$spec"
  assert_failure # pool and host do not belong to the same compute resource

  run govc vm.check.config -pool DC0_C1/Resources <<<"$spec"
  assert_failure # pool and host do not belong to the same compute resource

  # spec.memoryMB
  max_mem=$(govc object.collect -s DC0_C1_H0 capability.maxSupportedVmMemory)
  run govc vm.create -spec -m "$((max_mem+100))" $vm
  assert_success
  spec="$output"

  run govc vm.check.config -vm $vm -json <<<"$spec"
  assert_success
  assert_matches "outside the range"

  # spec.numCPUs
  max_cpu=$(govc object.collect -s "$host" summary.hardware.numCpuCores)
  run govc vm.create -spec -c "$((max_cpu+100))" $vm
  assert_success
  spec="$output"

  run govc vm.check.config -vm $vm -json <<<"$spec"
  assert_success
  assert_matches "vm requires 100 CPUs"

  # spec.guestId
  run govc vm.create -spec -g ttylinux -pool DC0_C0/Resources $vm
  assert_success
  spec="$output"

  run govc vm.check.config -vm $vm -json <<<"$spec"
  assert_success
  assert_matches unsupported
}

@test "vm.check.compat" {
  vcsim_env -cluster 2

  export GOVC_SHOW_UNRELEASED=true

  vm=DC0_C0_RP0_VM0

  run govc vm.check.compat -vm $vm -pool DC0_C0/Resources
  assert_success

  run govc vm.check.compat -vm $vm -pool DC0_C1/Resources
  assert_failure # InvalidArgument spec.pool
}

@test "vm.check.relocate" {
  vcsim_env -cluster 2

  export GOVC_SHOW_UNRELEASED=true

  vm=DC0_C0_RP0_VM0
  run govc vm.migrate -spec -pool DC0_C0/Resources $vm
  assert_success
  spec="$output"

  run govc vm.check.relocate -vm $vm <<<"$spec"
  assert_success
}
