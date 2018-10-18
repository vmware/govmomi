#!/usr/bin/env bats

load test_helper

@test "disk.ls" {
  esx_env

  run govc disk.ls
  assert_success

  run govc disk.ls enoent
  assert_failure
}

@test "disk.create" {
  esx_env

  name=$(new_id)

  run govc disk.create -size 10M "$name"
  assert_success
  id="${lines[1]}"

  run govc disk.ls "$id"
  assert_success

  govc disk.ls -json "$id" | jq .

  run govc disk.rm "$id"
  assert_success

  run govc disk.rm "$id"
  assert_failure
}

@test "disk.register" {
  esx_env

  id=$(new_id)
  vmdk="$id/$id.vmdk"

  run govc datastore.mkdir "$id"
  assert_success

  # create with VirtualDiskManager
  run govc datastore.disk.create -size 10M "$vmdk"
  assert_success

  run govc disk.register "$vmdk" "$id"
  assert_success
  id="$output"

  run govc disk.register "$vmdk" "$id"
  assert_failure

  run govc disk.rm "$id"
  assert_success

  run govc disk.rm "$id"
  assert_failure
}

@test "disk.snapshot" {
  esx_env

  name=$(new_id)

  run govc disk.create -size 10M "$name"
  assert_success
  id="${lines[1]}"

  run govc disk.snapshot.ls "$id"
  assert_success

  run govc disk.snapshot.create "$id"
  assert_success
  sid="${lines[1]}"

  govc disk.snapshot.ls "$id" | grep "$sid"

  govc disk.snapshot.ls -json "$id" | jq .

  run govc disk.snapshot.rm "$id" "$sid"
  assert_success

  run govc disk.snapshot.rm "$id" "$sid"
  assert_failure

  run govc disk.rm "$id"
  assert_success
}
