#!/usr/bin/env bats

load test_helper

@test "device.ls" {
  vcsim_env

  vm=$(new_empty_vm)

  result=$(govc device.ls -vm $vm ethernet-* | wc -l)
  [ $result -eq 1 ]
}

@test "device.info" {
  vcsim_env -esx

  vm=$(new_empty_vm)

  run govc device.info -vm $vm ide-200
  assert_success

  run govc device.info -vm $vm ide-20000
  assert_failure

  run govc device.info -vm $vm -net enoent
  assert_failure

  run govc device.info -vm $vm -net "VM Network" ide-200
  assert_failure

  result=$(govc device.info -vm $vm -net "VM Network" | grep "MAC Address" | wc -l)
  [ $result -eq 1 ]

  run govc device.info -vm $vm -json
  assert_matches ethernet-
  assert_matches '"name":' # injected field
  assert_matches '"type":' # injected field
}

@test "device.boot" {
  vcsim_env

  vm="DC0_H0_VM0"

  run govc device.remove -vm $vm cdrom-*
  assert_success

  result=$(govc device.ls -vm $vm -boot | wc -l)
  [ $result -eq 0 ]

  run govc device.boot -vm $vm -order floppy,cdrom,ethernet,disk
  assert_success

  result=$(govc device.ls -vm $vm -boot | wc -l)
  [ $result -eq 2 ]

  run govc device.cdrom.add -vm $vm
  assert_success

  run govc device.floppy.add -vm $vm
  assert_success

  run govc device.boot -vm $vm -order floppy,cdrom,ethernet,disk
  assert_success

  result=$(govc device.ls -vm $vm -boot | wc -l)
  [ $result -eq 4 ]

  run govc device.boot -vm $vm -order -
  assert_success

  result=$(govc device.ls -vm $vm -boot | wc -l)
  [ $result -eq 0 ]

  run govc device.boot -vm $vm -secure
  assert_failure

  run govc device.boot -vm $vm -secure -firmware efi
  assert_success

  run govc device.boot -vm $vm -order -
  assert_success

  firmware=$(govc object.collect -s vm/$vm config.firmware)
  assert_equal efi "$firmware"
}

@test "device.cdrom" {
  vcsim_env

  vm=$(new_empty_vm)

  result=$(govc device.ls -vm $vm | grep cdrom- | wc -l)
  [ $result -eq 0 ]

  run govc device.cdrom.add -vm $vm
  assert_success
  id=$output

  result=$(govc device.ls -vm $vm | grep $id | wc -l)
  [ $result -eq 1 ]

  run govc device.info -vm $vm $id
  assert_success

  run govc device.cdrom.insert -vm $vm -device $id x.iso
  assert_success

  run govc device.info -vm $vm $id
  assert_line "Summary: ISO [${GOVC_DATASTORE}] x.iso"

  run govc device.disconnect -vm $vm $id
  assert_success

  run govc device.connect -vm $vm $id
  assert_success

  run govc device.remove -vm $vm $id
  assert_success

  run govc device.disconnect -vm $vm $id
  assert_failure "govc: device '$id' not found"

  run govc device.cdrom.insert -vm $vm -device $id x.iso
  assert_failure "govc: device '$id' not found"

  run govc device.remove -vm $vm $id
  assert_failure "govc: device '$id' not found"
}

@test "device.floppy" {
  vcsim_env

  vm=$(new_empty_vm)

  result=$(govc device.ls -vm $vm | grep floppy- | wc -l)
  [ $result -eq 0 ]

  run govc device.floppy.add -vm $vm
  assert_success
  id=$output

  result=$(govc device.ls -vm $vm | grep $id | wc -l)
  [ $result -eq 1 ]

  run govc device.info -vm $vm $id
  assert_success

  run govc device.floppy.insert -vm $vm -device $id x.img
  assert_success

  run govc device.info -vm $vm $id
  assert_line "Summary: Image [${GOVC_DATASTORE}] x.img"

  run govc device.disconnect -vm $vm $id
  assert_success

  run govc device.connect -vm $vm $id
  assert_success

  run govc device.remove -vm $vm $id
  assert_success

  run govc device.disconnect -vm $vm $id
  assert_failure "govc: device '$id' not found"

  run govc device.floppy.insert -vm $vm -device $id x.img
  assert_failure "govc: device '$id' not found"

  run govc device.remove -vm $vm $id
  assert_failure "govc: device '$id' not found"
}

@test "device.serial" {
  vcsim_env

  vm=$(new_empty_vm)

  result=$(govc device.ls -vm $vm | grep serial- | wc -l)
  [ $result -eq 0 ]

  run govc device.serial.add -vm $vm
  assert_success
  id=$output

  result=$(govc device.ls -vm $vm | grep $id | wc -l)
  [ $result -eq 1 ]

  run govc device.info -vm $vm $id
  assert_success

  run govc device.serial.connect -vm $vm -
  assert_success

  run govc device.info -vm $vm $id
  assert_line "Summary: File [$GOVC_DATASTORE] $vm/${id}.log"

  uri=telnet://:33233
  run govc device.serial.connect -vm $vm -device $id $uri
  assert_success

  run govc device.info -vm $vm $id
  assert_line "Summary: Remote $uri"

  run govc device.serial.disconnect -vm $vm -device $id
  assert_success

  run govc device.info -vm $vm $id
  assert_line "Summary: Remote localhost:0"

  run govc device.disconnect -vm $vm $id
  assert_success

  run govc device.connect -vm $vm $id
  assert_success

  run govc device.remove -vm $vm $id
  assert_success

  run govc device.disconnect -vm $vm $id
  assert_failure "govc: device '$id' not found"

  run govc device.serial.connect -vm $vm -device $id $uri
  assert_failure "govc: device '$id' not found"

  run govc device.remove -vm $vm $id
  assert_failure "govc: device '$id' not found"
}

@test "device.scsi" {
  vcsim_env

  vm=$(new_empty_vm)

  result=$(govc device.ls -vm $vm | grep lsilogic- | wc -l)
  [ $result -eq 1 ]

  run govc device.scsi.add -vm $vm
  assert_success
  id=$output

  result=$(govc device.ls -vm $vm | grep $id | wc -l)
  [ $result -eq 1 ]

  result=$(govc device.ls -vm $vm | grep lsilogic- | wc -l)
  [ $result -eq 2 ]

  run govc device.scsi.add -vm $vm -type pvscsi
  assert_success
  id=$output

  result=$(govc device.ls -vm $vm | grep $id | wc -l)
  [ $result -eq 1 ]
}

@test "device.usb" {
  vcsim_env

  vm=$(new_empty_vm)

  result=$(govc device.ls -vm $vm | grep usb | wc -l)
  [ $result -eq 0 ]

  run govc device.usb.add -type enoent -vm $vm
  assert_failure

  run govc device.usb.add -vm $vm
  assert_success
  id=$output

  result=$(govc device.ls -vm $vm | grep $id | wc -l)
  [ $result -eq 1 ]

  run govc device.usb.add -type xhci -vm $vm
  assert_success
  id=$output

  result=$(govc device.ls -vm $vm | grep $id | wc -l)
  [ $result -eq 1 ]
}

@test "device.clock" {
  vcsim_env

  vm=$(new_empty_vm)

  result=$(govc device.ls -vm "$vm" | grep clock | wc -l)
  [ "$result" -eq 0 ]

  run govc device.clock.add -vm "$vm"
  assert_success
  id=$output

  result=$(govc device.ls -vm "$vm" | grep "$id" | wc -l)
  [ "$result" -eq 1 ]
}

@test "device.scsi slots" {
  vcsim_env

  vm=$(new_empty_vm)

  for i in $(seq 1 15) ; do
    name="disk-${i}"
    run govc vm.disk.create -vm "$vm" -name "$name" -size 1K
    assert_success
    result=$(govc device.ls -vm "$vm" | grep disk- | wc -l)
    [ "$result" -eq "$i" ]
  done
}

@test "device.match" {
  vcsim_env

  vm=DC0_H0_VM0

  run govc device.ls -vm $vm enoent-*
  assert_failure

  run govc device.info -vm $vm enoent-*
  assert_failure

  run govc device.info -vm $vm disk-*
  assert_success

  run govc vm.disk.create -vm $vm -name $vm/disk2.vmdk -size 10M
  assert_success

  run govc device.ls -vm $vm disk-*
  assert_success

  [ ${#lines[@]} -eq 2 ]

  run govc device.remove -vm $vm enoent-*
  assert_failure

  run govc device.remove -vm $vm disk-*
  assert_success

  run govc device.ls -vm $vm disk-*
  assert_failure
}
