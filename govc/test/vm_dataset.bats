#!/usr/bin/env bats

load test_helper

# Data set support requires the virtual machine be at virtual hardware version vmx-20 or later
new_unsupported_hardware_vm() {
  id=$(new_id)
  govc vm.create -version=vmx-19 $id
  echo $id
}

new_supported_hardware_vm() {
  id=$(new_id)
  govc vm.create -version=vmx-20 $id
  echo $id
}

@test "vm.dataset.ls" {
  vcsim_env

  old_vm=$(new_unsupported_hardware_vm)
  vm=$(new_supported_hardware_vm)

  run govc vm.dataset.ls -vm $vm
  assert_success ""

  # non-existing VM
  run govc vm.dataset.ls -vm enoent
  assert_failure "govc: vm 'enoent' not found"

  # VM hardware too old
  run govc vm.dataset.ls -vm $old_vm
  assert_failure "govc: 400 Bad Request: {\"error_type\":\"UNSUPPORTED\", \"messages\":[]}"
}

@test "vm.dataset.create" {
  vcsim_env

  old_vm=$(new_unsupported_hardware_vm)
  vm=$(new_supported_hardware_vm)

  # create
  id=$(new_id)
  run govc vm.dataset.create -vm $vm $id
  assert_success "$id"

  # list the created data set
  run govc vm.dataset.ls -vm $vm
  assert_success "$id"

  # create second data set
  id2=$(new_id)
  run govc vm.dataset.create -vm $vm $id2
  assert_success "$id2"

  # list the two data sets
  run govc vm.dataset.ls -vm $vm
  assert_success
  assert_line "$id"
  assert_line "$id2"

  # create with duplicate name
  run govc vm.dataset.create -vm $vm $id
  assert_failure "govc: 400 Bad Request: {\"error_type\":\"ALREADY_EXISTS\", \"messages\":[]}"

  # create on non-existing VM
  run govc vm.dataset.create -vm enoent $(new_id)
  assert_failure "govc: vm 'enoent' not found"

  # create on unsupported VM
  run govc vm.dataset.create -vm $old_vm $(new_id)
  assert_failure "govc: 400 Bad Request: {\"error_type\":\"UNSUPPORTED\", \"messages\":[]}"

  # create on suspended VM
  govc vm.power -suspend $vm
  run govc vm.dataset.create -vm $vm $(new_id)
  assert_failure "govc: 400 Bad Request: {\"error_type\":\"NOT_ALLOWED_IN_CURRENT_STATE\", \"messages\":[]}"
}

@test "vm.dataset.info" {
  vcsim_env

  vm=$(new_supported_hardware_vm)

  # create
  id=$(new_id)
  run govc vm.dataset.create -vm $vm -d "Some description." -guest-access READ_ONLY -omit-from-snapshot=false $id
  assert_success "$id"

  # get
  run govc vm.dataset.info -vm $vm $id
  assert_success
  assert_line "Name: $id"
  assert_line "Description: Some description."
  assert_line "Host: READ_WRITE"
  assert_line "Guest: READ_ONLY"
  assert_line "Used: 0"
  assert_line "OmitFromSnapshotAndClone: false"

  # create an entry
  run govc vm.dataset.entry.set -vm $vm -dataset $id key1 val1
  assert_success

  # get to verify the 'Used' field includes the footprint of the entry
  run govc vm.dataset.info -vm $vm $id
  assert_success
  assert_line "Name: $id"
  assert_line "Description: Some description."
  assert_line "Host: READ_WRITE"
  assert_line "Guest: READ_ONLY"
  assert_line "Used: 8"
  assert_line "OmitFromSnapshotAndClone: false"

  # non-existing VM
  run govc vm.dataset.info -vm enoent $id
  assert_failure "govc: vm 'enoent' not found"

  # non-existing data set
  run govc vm.dataset.info -vm $vm enoent
  assert_failure
  assert_matches "404 Not Found"
}

@test "vm.dataset.update" {
  vcsim_env

  vm=$(new_supported_hardware_vm)

  # create
  id=$(new_id)
  run govc vm.dataset.create -vm $vm -d "Initial description." $id
  assert_success "$id"

  # update description
  run govc vm.dataset.update -vm $vm -d "Updated description." $id
  assert_success ""

  # get the updated
  run govc vm.dataset.info -vm $vm $id
  assert_success
  assert_line "Name: $id"
  assert_line "Description: Updated description."

  # update on non-existing VM
  run govc vm.dataset.update -vm enoent -d "Even newer description." $id
  assert_failure "govc: vm 'enoent' not found"

  # update non-existing data set
  run govc vm.dataset.update -vm $vm -d "Even newer description." enoent
  assert_failure
  assert_matches "404 Not Found"

  # update on suspended VM
  govc vm.power -suspend $vm
  run govc vm.dataset.update -vm $vm -d "Even newer description." $id
  assert_failure "govc: 400 Bad Request: {\"error_type\":\"NOT_ALLOWED_IN_CURRENT_STATE\", \"messages\":[]}"
}

@test "vm.dataset.rm" {
  vcsim_env

  vm=$(new_supported_hardware_vm)

  # create
  id=$(new_id)
  run govc vm.dataset.create -vm $vm -d "Initial description." $id
  assert_success "$id"

  # delete on non-existing VM
  run govc vm.dataset.rm -vm enoent $id
  assert_failure "govc: vm 'enoent' not found"

  # delete non-existing data set
  run govc vm.dataset.rm -vm $vm enoent
  assert_failure
  assert_matches "404 Not Found"

  # delete on suspended VM
  govc vm.power -suspend $vm
  run govc vm.dataset.rm -vm $vm $id
  assert_failure "govc: 400 Bad Request: {\"error_type\":\"NOT_ALLOWED_IN_CURRENT_STATE\", \"messages\":[]}"
  govc vm.power -on $vm

  # delete
  run govc vm.dataset.rm -vm $vm $id
  assert_success ""

  # list to verify
  run govc vm.dataset.ls -vm $vm
  assert_success ""

  # create a data set with an entry
  id=$(new_id)
  run govc vm.dataset.create -vm $vm -d "Initial description." $id
  assert_success "$id"
  run govc vm.dataset.entry.set -vm $vm -dataset $id key1 val1
  assert_success

  # try to delete the non-empty data set
  run govc vm.dataset.rm -vm $vm $id
  assert_failure "govc: 400 Bad Request: {\"error_type\":\"RESOURCE_IN_USE\", \"messages\":[]}"

  # delete the non-empty data set with force
  run govc vm.dataset.rm -vm $vm -force=true $id
  assert_success ""
}

@test "vm.dataset.vmclone" {
  vcsim_env

  vm=$(new_supported_hardware_vm)

  # create data set which is included in clones/snapshots
  id=$(new_id)
  run govc vm.dataset.create -vm $vm -d "First data set." -omit-from-snapshot=false $id
  assert_success "$id"

  # create data set which is excluded from clones/snapshots
  id2=$(new_id)
  run govc vm.dataset.create -vm $vm -d "Second data set." -omit-from-snapshot=true $id2
  assert_success "$id2"

  # clone the VM
  cloned_vm=$(new_id)
  run govc vm.clone -vm $vm $cloned_vm
  assert_success

  # update the description of the data set on the original VM
  run govc vm.dataset.update -vm $vm -d "Updated description." $id
  assert_success ""
  
  # create another data set which is included in clones/snapshots on the original VM
  id3=$(new_id)
  run govc vm.dataset.create -vm $vm -d "Third data set." -omit-from-snapshot=false $id3
  assert_success "$id3"

  # list the data sets on the cloned VM
  run govc vm.dataset.ls -vm $cloned_vm
  assert_success "$id"

  # verify the data set on the cloned VM was not affected by the update
  run govc vm.dataset.info -vm $cloned_vm $id
  assert_success
  assert_line "Name: $id"
  assert_line "Description: First data set."
}

@test "vm.dataset.vmsnapshot" {
  vcsim_env

  vm=$(new_supported_hardware_vm)

  # create data set which is included in clones/snapshots
  id=$(new_id)
  run govc vm.dataset.create -vm $vm -d "First data set." -omit-from-snapshot=false $id
  assert_success "$id"

  # create data set which is excluded from clones/snapshots
  id2=$(new_id)
  run govc vm.dataset.create -vm $vm -d "Second data set." -omit-from-snapshot=true $id2
  assert_success "$id2"

  # create a snapshot
  snapshot=$(new_id)
  run govc snapshot.create -vm $vm $snapshot
  assert_success

  # update the description of the data set on the original VM
  run govc vm.dataset.update -vm $vm -d "Updated description." $id
  assert_success ""

  # create another data set which is included in clones/snapshots
  id3=$(new_id)
  run govc vm.dataset.create -vm $vm -d "Third data set." -omit-from-snapshot=false $id3
  assert_success "$id3"

  # revert the VM to the snapshot
  run govc snapshot.revert -vm $vm $snapshot
  assert_success

  # list the data sets on the VM
  run govc vm.dataset.ls -vm $vm
  assert_success "$id"

  # verify the data set does not contain the update
  run govc vm.dataset.info -vm $vm $id
  assert_success
  assert_line "Name: $id"
  assert_line "Description: First data set."  
}

@test "vm.dataset.entry.ls" {
  vcsim_env

  # setup VM and data set
  vm=$(new_supported_hardware_vm)
  ds=$(new_id)
  run govc vm.dataset.create -vm $vm $ds
  assert_success "$ds"

  run govc vm.dataset.entry.ls -vm $vm -dataset $ds
  assert_success ""

  # non-existing VM
  run govc vm.dataset.entry.ls -vm enoent -dataset $ds
  assert_failure "govc: vm 'enoent' not found"

  # non-existing data set
  run govc vm.dataset.entry.ls -vm $vm -dataset enoent
  assert_failure
  assert_matches "404 Not Found"

  # host does not have read access to the data set
  run govc vm.dataset.update -vm $vm -host-access NONE $ds
  assert_success ""
  run govc vm.dataset.entry.ls -vm $vm -dataset $ds
  assert_failure "govc: 400 Bad Request: {\"error_type\":\"UNAUTHORIZED\", \"messages\":[]}"
}

@test "vm.dataset.entry.get" {
  vcsim_env

  # setup VM and data set
  vm=$(new_supported_hardware_vm)
  ds=$(new_id)
  run govc vm.dataset.create -vm $vm $ds
  assert_success "$ds"

  # non-existing VM
  run govc vm.dataset.entry.get -vm enoent -dataset $ds somekey
  assert_failure "govc: vm 'enoent' not found"

  # non-existing data set
  run govc vm.dataset.entry.get -vm $vm -dataset enoent somekey
  assert_failure
  assert_matches "404 Not Found"

  # non-existing entry
  run govc vm.dataset.entry.get -vm $vm -dataset $ds enoent
  assert_failure
  assert_matches "404 Not Found"
}

@test "vm.dataset.entry.set" {
  vcsim_env

  # setup VM and data set
  vm=$(new_supported_hardware_vm)
  ds=$(new_id)
  run govc vm.dataset.create -vm $vm $ds
  assert_success "$ds"

  # create new entry
  run govc vm.dataset.entry.set -vm $vm -dataset $ds key1 val1
  assert_success

  # list the created entry
  run govc vm.dataset.entry.ls -vm $vm -dataset $ds
  assert_success "key1"

  # get the value of the created entry
  run govc vm.dataset.entry.get -vm $vm -dataset $ds key1
  assert_success "val1"

  # update the value of the entry
  run govc vm.dataset.entry.set -vm $vm -dataset $ds key1 val1b
  assert_success

  # get the updated value
  run govc vm.dataset.entry.get -vm $vm -dataset $ds key1
  assert_success "val1b"

  # non-existing VM
  run govc vm.dataset.entry.set -vm enoent -dataset $ds key2 val2
  assert_failure "govc: vm 'enoent' not found"

  # non-existing data set
  run govc vm.dataset.entry.set -vm $vm -dataset enoent key2 val2
  assert_failure
  assert_matches "404 Not Found"

  # suspended VM
  govc vm.power -suspend $vm
  run govc vm.dataset.entry.set -vm $vm -dataset $ds key2 val2
  assert_failure "govc: 400 Bad Request: {\"error_type\":\"NOT_ALLOWED_IN_CURRENT_STATE\", \"messages\":[]}"
  govc vm.power -on $vm
}

@test "vm.dataset.entry.rm" {
  vcsim_env

  # setup VM, data set and entry
  vm=$(new_supported_hardware_vm)
  ds=$(new_id)
  run govc vm.dataset.create -vm $vm $ds
  assert_success "$ds"
  run govc vm.dataset.entry.set -vm $vm -dataset $ds key1 val1
  assert_success

  # delete on non-existing VM
  run govc vm.dataset.entry.rm -vm enoent -dataset $ds key1
  assert_failure "govc: vm 'enoent' not found"

  # delete on non-existing data set
  run govc vm.dataset.entry.rm -vm $vm -dataset enoent key1
  assert_failure
  assert_matches "404 Not Found"

  # delete on suspended VM
  govc vm.power -suspend $vm
  run govc vm.dataset.entry.rm -vm $vm -dataset $ds key1
  assert_failure "govc: 400 Bad Request: {\"error_type\":\"NOT_ALLOWED_IN_CURRENT_STATE\", \"messages\":[]}"
  govc vm.power -on $vm

  # delete
  run govc vm.dataset.entry.rm -vm $vm -dataset $ds key1
  assert_success

  # list to verify
  run govc vm.dataset.entry.ls -vm $vm -dataset $ds
  assert_success ""
}

@test "vm.dataset.entry.access" {
  vcsim_env

  # setup VM, data set and entry
  vm=$(new_supported_hardware_vm)
  ds=$(new_id)
  run govc vm.dataset.create -vm $vm $ds
  assert_success "$ds"
  run govc vm.dataset.entry.set -vm $vm -dataset $ds key1 val1
  assert_success

  # change the host access to NONE (govc calls the VC API, so the guest access does not matter here)
  # the default access mode is READ_WRITE and has already been covered by previous tests
  run govc vm.dataset.update -vm $vm -host-access NONE $ds
  assert_success

  # list
  run govc vm.dataset.entry.ls -vm $vm -dataset $ds
  assert_failure "govc: 400 Bad Request: {\"error_type\":\"UNAUTHORIZED\", \"messages\":[]}"

  # get
  run govc vm.dataset.entry.get -vm $vm -dataset $ds key1
  assert_failure "govc: 400 Bad Request: {\"error_type\":\"UNAUTHORIZED\", \"messages\":[]}"

  # set
  run govc vm.dataset.entry.set -vm $vm -dataset $ds key1 val1b
  assert_failure "govc: 400 Bad Request: {\"error_type\":\"UNAUTHORIZED\", \"messages\":[]}"

  # delete
  run govc vm.dataset.entry.rm -vm $vm -dataset $ds key1
  assert_failure "govc: 400 Bad Request: {\"error_type\":\"UNAUTHORIZED\", \"messages\":[]}"

  # change the host access to READ_ONLY
  run govc vm.dataset.update -vm $vm -host-access READ_ONLY $ds
  assert_success

  # list
  run govc vm.dataset.entry.ls -vm $vm -dataset $ds
  assert_success "key1"

  # get
  run govc vm.dataset.entry.get -vm $vm -dataset $ds key1
  assert_success "val1"

  # set
  run govc vm.dataset.entry.set -vm $vm -dataset $ds key1 val1b
  assert_failure "govc: 400 Bad Request: {\"error_type\":\"UNAUTHORIZED\", \"messages\":[]}"

  # delete
  run govc vm.dataset.entry.rm -vm $vm -dataset $ds key1
  assert_failure "govc: 400 Bad Request: {\"error_type\":\"UNAUTHORIZED\", \"messages\":[]}"
}
