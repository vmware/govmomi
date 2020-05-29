#!/usr/bin/env bats

load test_helper

@test "import.ova" {
  vcsim_env -app 1

  run govc import.ova "$GOVC_IMAGES/$TTYLINUX_NAME.ova"
  assert_success

  run govc device.ls -vm "$TTYLINUX_NAME"
  assert_success
  assert_matches "disk-"

  run govc vm.destroy "$TTYLINUX_NAME"
  assert_success

  # link ovf/ova to datastore so we can test with an http source
  dir=$(govc datastore.info -json | jq -r .Datastores[].Info.Url)
  ln -s "$GOVC_IMAGES/$TTYLINUX_NAME."* "$dir"

  run govc import.spec "https://$(govc env GOVC_URL)/folder/$TTYLINUX_NAME.ovf"
  assert_success

  proto=$(jq -r .IPProtocol <<<"$output")
  assert_equal IPv4 "$proto"

  run govc import.ova "https://$(govc env GOVC_URL)/folder/$TTYLINUX_NAME.ova"
  assert_success

  run govc device.ls -vm "$TTYLINUX_NAME"
  assert_success
  assert_matches "disk-"

  run govc vm.destroy "$TTYLINUX_NAME"
  assert_success

  run govc import.ova -pool /DC0/host/DC0_C0/Resources/DC0_C0_APP0 "$GOVC_IMAGES/$TTYLINUX_NAME.ova"
  assert_success
}

@test "import.ova with iso" {
  vcsim_env

  run govc import.ova "$GOVC_IMAGES/${TTYLINUX_NAME}-live.ova"
  assert_success

  run govc device.ls -vm "${TTYLINUX_NAME}-live"
  assert_success
  assert_matches "disk-"
  assert_matches "cdrom-.* ${TTYLINUX_NAME}-live/_deviceImage0.iso"

  run govc vm.destroy "${TTYLINUX_NAME}-live"
  assert_success
}

@test "import.ovf" {
  vcsim_env

  run govc import.ovf $GOVC_IMAGES/${TTYLINUX_NAME}.ovf
  assert_success

  run govc vm.destroy ${TTYLINUX_NAME}
  assert_success

  # test w/ relative dir
  pushd $BATS_TEST_DIRNAME >/dev/null
  run govc import.ovf ./images/${TTYLINUX_NAME}.ovf
  assert_success
  popd >/dev/null

  run govc vm.destroy ${TTYLINUX_NAME}
  assert_success

  # ensure vcsim doesn't panic without capacityAllocationUnits
  dir=$($mktemp --tmpdir -d govc-test-XXXXX)
  sed -e s/capacityAllocationUnits/invalid/ "$GOVC_IMAGES/$TTYLINUX_NAME.ovf" > "$dir/$TTYLINUX_NAME.ovf"
  touch "$dir/$TTYLINUX_NAME-disk1.vmdk" # .vmdk contents don't matter to vcsim
  run govc import.ovf "$dir/$TTYLINUX_NAME.ovf"
  assert_success
  rm -rf "$dir"
}

@test "import.ovf -host.ipath" {
  vcsim_env

  run govc import.ovf -host.ipath="$(govc find / -type h | head -1)" "$GOVC_IMAGES/${TTYLINUX_NAME}.ovf"
  assert_success

  run govc vm.destroy "$TTYLINUX_NAME"
  assert_success
}

@test "import.ovf with name in options" {
  vcsim_env

  name=$(new_id)
  file=$($mktemp --tmpdir govc-test-XXXXX)
  echo "{ \"Name\": \"${name}\"}" > ${file}

  run govc import.ovf -options="${file}" $GOVC_IMAGES/${TTYLINUX_NAME}.ovf
  assert_success

  run govc vm.destroy "${name}"
  assert_success

  rm -f ${file}
}

@test "import.ovf with import.spec result" {
  vcsim_env

  file=$($mktemp --tmpdir govc-test-XXXXX)
  name=$(new_id)

  govc import.spec $GOVC_IMAGES/${TTYLINUX_NAME}.ovf > ${file}

  run govc import.ovf -name="${name}" -options="${file}" $GOVC_IMAGES/${TTYLINUX_NAME}.ovf
  assert_success

  run govc vm.destroy "${name}"
  assert_success
}

@test "import.ovf with name as argument" {
  vcsim_env

  name=$(new_id)

  run govc import.ova -name="${name}" $GOVC_IMAGES/${TTYLINUX_NAME}.ova
  assert_success

  run govc vm.destroy "${name}"
  assert_success
}

@test "import.vmdk" {
  esx_env

  name=$(new_id)

  run govc import.vmdk "$GOVC_TEST_VMDK_SRC" "$name"
  assert_success

  run govc import.vmdk "$GOVC_TEST_VMDK_SRC" "$name"
  assert_failure # exists

  run govc import.vmdk -force "$GOVC_TEST_VMDK_SRC" "$name"
  assert_success # exists, but -force was used
}

@test "import duplicate dvpg names" {
  vcsim_env

  run govc dvs.create DVS1 # DVS0 already exists
  assert_success

  run govc dvs.portgroup.add -dvs DVS0 -type ephemeral NSX-dvpg
  assert_success

  run govc dvs.portgroup.add -dvs DVS1 -type ephemeral NSX-dvpg
  assert_success

  ovf="$GOVC_IMAGES/$TTYLINUX_NAME.ovf"

  spec=$(govc import.spec "$GOVC_IMAGES/$TTYLINUX_NAME.ovf")

  run govc import.ovf -name ttylinux -options - "$ovf" <<<"$spec"
  assert_success # no network specified

  options=$(jq ".NetworkMapping[].Network = \"enoent\"" <<<"$spec")

  run govc import.ovf -options - "$ovf" <<<"$options"
  assert_failure # network not found

  options=$(jq ".NetworkMapping[].Network = \"NSX-dvpg\"" <<<"$spec")

  run govc import.ovf -options - "$ovf" <<<"$options"
  assert_failure # 2 networks have the same name

  options=$(jq ".NetworkMapping[].Network = \"DVS0/NSX-dvpg\"" <<<"$spec")

  run govc import.ovf -name ttylinux2 -options - "$ovf" <<<"$options"
  assert_success # switch_name/portgroup_name is unique

  switch=$(govc find -i network -name DVS0)
  id=$(govc find -i network -config.distributedVirtualSwitch "$switch" -name NSX-dvpg)
  options=$(jq ".NetworkMapping[].Network = \"$id\"" <<<"$spec")

  run govc import.ovf -options - "$ovf" <<<"$options"
  assert_success # using raw MO id
}
