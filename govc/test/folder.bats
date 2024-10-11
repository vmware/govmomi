#!/usr/bin/env bats

load test_helper

@test "folder.info" {
  vcsim_env -esx

  for name in / vm host network datastore ; do
    run govc folder.info $name
    assert_success

    govc folder.info -json $name
    assert_success
  done

  result=$(govc folder.info '*' | grep -c Name:)
  [ "$result" -eq 4 ]

  run govc info.info /enoent
  assert_failure
}

@test "folder.create" {
    vcsim_env

    name=$(new_id)

    # relative to $GOVC_DATACENTER
    run govc folder.create $name
    assert_failure

    run govc folder.create vm/$name
    assert_success

    run govc folder.info vm/$name
    assert_success

    run govc folder.info /$GOVC_DATACENTER/vm/$name
    assert_success

    run govc folder.create vm/$name
    assert_failure # duplicate name

    run govc object.destroy vm/$name
    assert_success

    unset GOVC_DATACENTER
    # relative to /

    run govc folder.create /$name
    assert_success

    run govc folder.info /$name
    assert_success

    child=$(new_id)
    run govc folder.create /$child
    assert_success

    run govc folder.info /$name/$child
    assert_failure

    run govc object.mv /$child /$name
    assert_success

    run govc folder.info /$name/$child
    assert_success

    new=$(new_id)
    run govc object.rename /$name $new
    assert_success
    name=$new

    run govc folder.info /$name
    assert_success

    run govc object.destroy /$name
    assert_success
}

@test "folder.place" {
    vcsim_env

    export GOVC_SHOW_UNRELEASED=true

    run govc folder.place
    assert_failure # -vm must be specified

    run govc folder.place -vm invalid
    assert_failure # vm must be valid

    vmid=$(new_id)
    run govc vm.create -cluster DC0_C0 "$vmid"
    assert_success

    run govc folder.place $name -vm $id
    assert_failure # -pool must be specified

    poolid=$(new_id)

    path="/DC0/host/DC0_C0/Resources/$poolid"
    run govc pool.create $path
    assert_success

    run govc folder.place $name -vm $vmid -pool $poolid
    assert_failure # -Type must be specified

    run govc folder.place $name -vm $vmid -pool $poolid -type relocate
    assert_success

    run govc folder.place -json $name -vm $vmid -pool $poolid -type relocate
    assert_success
}
