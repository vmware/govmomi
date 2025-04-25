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
  dir=$(govc datastore.info -json | jq -r .datastores[].info.url)
  ln -s "$GOVC_IMAGES/$TTYLINUX_NAME"* "$dir"

  run govc import.spec "https://$(govc env GOVC_URL)/folder/$TTYLINUX_NAME.ovf"
  assert_success

  run govc import.spec "https://$(govc env GOVC_URL)/folder/$TTYLINUX_NAME.ova"
  assert_success

  proto=$(jq -r .IPProtocol <<<"$output")
  assert_equal IPv4 "$proto"

  run govc import.ovf "https://$(govc env GOVC_URL)/folder/$TTYLINUX_NAME.ovf"
  assert_success

  run govc vm.destroy "$TTYLINUX_NAME"
  assert_success

  run govc import.ova -verbose "https://$(govc env GOVC_URL)/folder/$TTYLINUX_NAME.ova"
  assert_success

  run govc device.ls -vm "$TTYLINUX_NAME"
  assert_success
  assert_matches "disk-"

  run govc vm.destroy "$TTYLINUX_NAME"
  assert_success

  run govc import.ova -verbose -pool /DC0/host/DC0_C0/Resources/DC0_C0_APP0 "$GOVC_IMAGES/$TTYLINUX_NAME.ova"
  assert_success
}

@test "import.ova with checksum validation" {
  vcsim_env -app 1

  run govc import.ova -name=bad-checksum-vm -m "$GOVC_IMAGES/$TTYLINUX_NAME-bad-checksum.ova"
  assert_failure # vmdk checksum mismatch

  run govc import.ova -name=good-checksum-vm -m "$GOVC_IMAGES/$TTYLINUX_NAME.ova"
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
  dir=$($mktemp --tmpdir -d govc-test-XXXXX 2>/dev/null || $mktemp -d -t govc-test-XXXXX)
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
  file=$($mktemp --tmpdir govc-test-XXXXX 2>/dev/null || $mktemp -t govc-test-XXXXX)
  echo "{ \"Name\": \"${name}\"}" > ${file}

  run govc import.ovf -options="${file}" $GOVC_IMAGES/${TTYLINUX_NAME}.ovf
  assert_success

  run govc vm.destroy "${name}"
  assert_success

  rm -f ${file}
}

@test "import.ovf with import.spec result" {
  vcsim_env

  file=$($mktemp --tmpdir govc-test-XXXXX 2>/dev/null || $mktemp -t govc-test-XXXXX)
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

@test "import.vmdk -i" {
  # ClientFlag.Process errors otherwise, but not using client with -i flag
  export GOVC_URL="unused"

  run govc import.vmdk -i "$GOVC_IMAGES/${TTYLINUX_NAME}.ovf"
  assert_failure

  run govc import.vmdk -i "$GOVC_TEST_VMDK_SRC"
  assert_success

  run govc import.vmdk -json -i "$GOVC_TEST_VMDK_SRC"
  assert_success

  run jq . <<<"$output"
  assert_success
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
  grep -v "invalid NetworkMapping.Name" <<<"$output"

  options=$(jq ".NetworkMapping[].Network = \"/DC0/host/DC0_C0/DC0_DVPG0\"" <<<"$spec")

  run govc import.ovf -name ttylinux3 -options - "$ovf" <<<"$options"
  assert_success # cluster path is unique

  switch=$(govc find -i network -name DVS0)
  id=$(govc find -i network -config.distributedVirtualSwitch "$switch" -name NSX-dvpg)
  options=$(jq ".NetworkMapping[].Network = \"$id\"" <<<"$spec")
  options=$(jq ".NetworkMapping[].Name = \"enoent\"" <<<"$options")

  run govc import.ovf -options - "$ovf" <<<"$options"
  assert_success # using raw MO id
  grep "invalid NetworkMapping.Name" <<<"$output"

  run govc import.ovf -name netflag -net enoent "$ovf"
  assert_failure

  run govc import.ovf -name netflag -net "VM Network" "$ovf"
  assert_success
}

@test "import invalid disk provisioning" {
  vcsim_env

  ovf="$GOVC_IMAGES/$TTYLINUX_NAME.ovf"

  spec=$(govc import.spec "$GOVC_IMAGES/$TTYLINUX_NAME.ovf")

  options=$(jq ".DiskProvisioning = \"enoent\"" <<<"$spec")

  run govc import.ovf -options - "$ovf" <<<"$options"
  assert_failure "govc: Disk provisioning type not supported: enoent"

  options=$(jq ".DiskProvisioning = \"monolithicSparse\"" <<<"$spec")

  run govc import.ovf -options - "$ovf" <<<"$options"
  assert_failure
  assert_matches DeviceUnsupportedForVmPlatform
}

@test "import properties" {
  vcsim_env

  ovf=../../ovf/fixtures/properties.ovf

  govc import.spec $ovf | grep -q -v nfs_mount

  options=$(govc import.spec -hidden $ovf)

  grep -q -v vm.name <<<"$options"
  grep -q nfs_mount <<<"$options"

  run govc import.ovf -name "$(new_id)" -options - "$ovf" <<<"$options"
  assert_success # hidden options but no value changes

  run govc import.ovf -options - "$ovf" <<<"${options//transfer/other}"
  assert_failure # userConfigurable=false

  id=$(new_id)
  run govc import.ovf -name "$id" -hidden -options - "$ovf" <<<"${options//transfer/other}"
  assert_success # userConfigurable=true

  config=$(govc object.collect -o -json "vm/$id" | jq .config.vAppConfig)
  name=$(jq -r .product[0].name <<<"$config")
  version=$(jq -r .product[0].version <<<"$config")
  assert_equal ttylinux "$name"
  assert_equal 16.1 "$version"
}

@test "import datastore cluster" {
  vcsim_env -pod 1 -ds 3

  unset GOVC_DATASTORE

  pod=DC0_POD0

  run govc import.ova -ds $pod "$GOVC_IMAGES/$TTYLINUX_NAME.ova"
  assert_failure # datastore cluster "DC0_POD0" has no datastores

  run govc object.mv /DC0/datastore/LocalDS_* $pod
  assert_success

  run govc import.ova -ds $pod "$GOVC_IMAGES/$TTYLINUX_NAME.ova"
  assert_success

  run govc import.ova -ds $pod -name test-url -lease -json "$GOVC_IMAGES/$TTYLINUX_NAME.ova"
  assert_success

  run jq -r .info.deviceUrl[].url <<<"$output"
  assert_success

  # Device URL.Host should be the same as vcsim's https://127.0.0.1:$port/nfc/.../disk-0.vmdk
  run govc env -x -u "$output" GOVC_URL_HOST
  assert_success "$(govc env -x GOVC_URL_HOST)"
}

@test "import esx" {
  vcsim_env -esx

  run govc import.ova "$GOVC_IMAGES/$TTYLINUX_NAME.ova"
  assert_success

  run govc import.ova -name test-url -lease -json "$GOVC_IMAGES/$TTYLINUX_NAME.ova"
  assert_success

  run jq -r .info.deviceUrl[].url <<<"$output"
  assert_success

  # Device URL.Host should be '*' https://*:$port/nfc/.../disk-0.vmdk
  run govc env -x -u "$output" GOVC_URL_HOST
  assert_success "*"
}
