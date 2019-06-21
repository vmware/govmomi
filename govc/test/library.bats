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

  run govc session.ls
  assert_success
  govc session.ls -json | jq .

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

  run govc library.export "/my-content/$TTYLINUX_NAME/*.ovf" -
  assert_success "$(cat "$GOVC_IMAGES/$TTYLINUX_NAME.ovf")"

  name="$BATS_TMPDIR/govc-$id-export"
  run govc library.export "/my-content/$TTYLINUX_NAME/*.ovf" "$name"
  assert_equal "$(cat "$GOVC_IMAGES/$TTYLINUX_NAME.ovf")" "$(cat "$name")"
  rm "$name"

  mkdir "$name"
  run govc library.export "/my-content/$TTYLINUX_NAME" "$name"
  assert_success
  assert_equal "$(cat "$GOVC_IMAGES/$TTYLINUX_NAME.ovf")" "$(cat "$name/$TTYLINUX_NAME.ovf")"
  rm -r "$name"

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

  run govc session.ls
  assert_success
  govc session.ls -json | jq .
}

@test "library.deploy" {
  vcsim_env

  run govc library.create my-content
  assert_success
  library_id="$output"

  # link ovf/ova to datastore so we can test library.import with an http source
  dir=$(govc datastore.info -json | jq -r .Datastores[].Info.Url)
  ln -s "$GOVC_IMAGES/$TTYLINUX_NAME."* "$dir"

  run govc library.import -pull my-content "https://$(govc env GOVC_URL)/folder/$TTYLINUX_NAME.ovf"
  assert_success

  run govc library.deploy "my-content/$TTYLINUX_NAME" ttylinux
  assert_success

  run govc vm.info ttylinux
  assert_success

  run govc library.import -pull -n ttylinux-unpacked my-content "https://$(govc env GOVC_URL)/folder/$TTYLINUX_NAME.ova"
  assert_success

  item_id=$(govc library.info -json /my-content/ttylinux-unpacked | jq -r .[].id)
  assert_equal "$(cat "$GOVC_IMAGES/$TTYLINUX_NAME.ovf")" "$(cat "$dir/contentlib-$library_id/$item_id/$TTYLINUX_NAME.ovf")"

  cat > "$BATS_TMPDIR/ttylinux.json" <<EOF
{
  "DiskProvisioning": "flat",
  "IPAllocationPolicy": "dhcpPolicy",
  "IPProtocol": "IPv4",
  "NetworkMapping": [
    {
      "Name": "nat",
      "Network": "DC0_DVPG0"
    }
  ],
  "MarkAsTemplate": false,
  "PowerOn": false,
  "InjectOvfEnv": false,
  "WaitForIP": false,
  "Name": "ttylinux2"
}
EOF

  run govc library.deploy -options "$BATS_TMPDIR/ttylinux.json" "my-content/$TTYLINUX_NAME"
  assert_success
  rm "$BATS_TMPDIR/ttylinux.json"

  run govc vm.info -r ttylinux2
  assert_success
  assert_matches DC0_DVPG0

  run govc vm.destroy ttylinux ttylinux2
  assert_success

  item_id=$(govc library.info -json "/my-content/$TTYLINUX_NAME" | jq -r .[].id)

  run govc datastore.rm "contentlib-$library_id/$item_id" # remove library files out-of-band, forcing a deploy error below
  assert_success

  run govc library.deploy "my-content/$TTYLINUX_NAME" ttylinux2
  assert_failure
}
