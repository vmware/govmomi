#!/usr/bin/env bats

load test_helper

@test "device.ls" {
  vm=$(new_empty_vm)

  result=$(govc device.ls -vm $vm | grep ethernet-0 | wc -l)
  [ $result -eq 1 ]
}

@test "device.info" {
  vm=$(new_empty_vm)

  run govc device.info -vm $vm ide-200
  assert_success

  run govc device.info -vm $vm ide-20000
  assert_failure
}

@test "device.cdrom" {
  vm=$(new_empty_vm)

  run govc device.cdrom.add -vm $vm
  assert_success
  id=$output

  run govc device.info -vm $vm $id
  assert_success

  run govc device.cdrom.insert -vm $vm -device $id x.iso
  assert_success

  run govc device.disconnect -vm $vm $id
  assert_success

  run govc device.connect -vm $vm $id
  assert_success

  run govc device.remove -vm $vm $id
  assert_success

  run govc device.disconnect -vm $vm $id
  assert_failure "Error: device '$id' not found"

  run govc device.cdrom.insert -vm $vm -device $id x.iso
  assert_failure "Error: device '$id' not found"

  run govc device.remove -vm $vm $id
  assert_failure "Error: device '$id' not found"
}
