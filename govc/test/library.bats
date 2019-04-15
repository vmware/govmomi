#!/usr/bin/env bats

load test_helper

@test "library" {
  vcsim_env

  run govc library.create my-content
  assert_success
  id="$output"

  jid=$(govc library.ls -json /my-content | jq -r .[].id)
  assert_equal "$id" "$jid"

  # run govc library.ls enoent
  # assert_failure # TODO: currently exit 0's

  run govc library.ls my-content
  assert_success "/my-content"

  run govc library.info /my-content
  assert_success
  assert_matches "$id"

  run govc library.rm /my-content
  assert_success
}

@test "library.import" {
  vcsim_env

  run govc library.create my-content
  assert_success
  library_id="$output"

  # run govc library.ls enoent
  # assert_failure # TODO: currently exit 0's

  # run govc library.info enoent
  # assert_failure # TODO: currently exit 0's

  run govc library.ls /my-content/*
  assert_success ""

  run govc library.import -n library.bats /my-content library.bats # any file will do
  assert_success

  run govc library.info /my-content/my-item
  assert_success
  assert_matches "$id"

  run govc library.info /my-content/*
  assert_success
  assert_matches "$id"

  id=$(govc library.ls -json "/my-content/library.bats" | jq -r .[].id)

  run govc library.info "/my-content/library.bats"
  assert_success
  assert_matches "$id"

  run govc library.rm /my-content/library.bats
  assert_success

  run govc library.import /my-content "$GOVC_IMAGES/$TTYLINUX_NAME.ova"
  assert_success

  run govc library.ls "/my-content/$TTYLINUX_NAME/*"
  assert_success
  assert_matches "$TTYLINUX_NAME.ovf"
  assert_matches "$TTYLINUX_NAME-disk1.vmdk"

  run govc library.import /my-content "$GOVC_IMAGES/$TTYLINUX_NAME.ovf"
  assert_failure # already_exists

  run govc library.rm "/my-content/$TTYLINUX_NAME"
  assert_success

  run govc library.import /my-content "$GOVC_IMAGES/$TTYLINUX_NAME.ovf"
  assert_success

  run govc library.ls "/my-content/$TTYLINUX_NAME/*"
  assert_success
  assert_matches "$TTYLINUX_NAME.ovf"
  assert_matches "$TTYLINUX_NAME-disk1.vmdk"

  run govc library.import /my-content "$GOVC_IMAGES/$TTYLINUX_NAME.iso"
  assert_failure # already_exists

  run govc library.import -n ttylinux-live /my-content "$GOVC_IMAGES/$TTYLINUX_NAME.iso"
  assert_success

  run govc library.ls "/my-content/ttylinux-live/*"
  assert_success
  assert_matches "$TTYLINUX_NAME.iso"

  if [ ! -e "$GOVC_IMAGES/ttylinux-latest.ova" ] ; then
    ln -s "$GOVC_IMAGES/$TTYLINUX_NAME.ova" "$GOVC_IMAGES/ttylinux-latest.ova"
  fi
  # test where $name.{ovf,mf} differs from ova name
  run govc library.import -m my-content "$GOVC_IMAGES/ttylinux-latest.ova"
  assert_success
  run govc library.ls "/my-content/ttylinux-latest/*"
  assert_success
  assert_matches "$TTYLINUX_NAME.ovf"
  assert_matches "$TTYLINUX_NAME-disk1.vmdk"
}

@test "library.deploy" {
  vcsim_env

  run govc library.create my-content
  assert_success
  library_id="$output"

  run govc library.import my-content "$GOVC_IMAGES/$TTYLINUX_NAME.ova"
  assert_success

  run govc vcenter.deploy my-content "$TTYLINUX_NAME" ttylinux
  assert_success

  run govc vm.info ttylinux
  assert_success

  run govc vm.destroy ttylinux
  assert_success

  item_id=$(govc library.info -json "/my-content/$TTYLINUX_NAME" | jq -r .[].id)

  run govc datastore.rm "contentlib-$library_id/$item_id" # remove library files out-of-band, forcing a deploy error below
  assert_success

  run govc vcenter.deploy my-content "$TTYLINUX_NAME.ova" ttylinux2
  assert_failure
}
