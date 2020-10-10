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

  # test tags attach,ls,detach for libraries
  run govc tags.category.create -m "$(new_id)"
  assert_success
  category="$output"

  tag_name=$(new_id)
  run govc tags.create -c "$category" "$tag_name"
  assert_success
  tag=$output

  run govc tags.attach "$tag" /my-content
  assert_success

  run govc tags.attached.ls "$tag_name"
  assert_success "com.vmware.content.Library:$id"

  run govc tags.attached.ls -r /my-content
  assert_success "$tag_name"

  run govc tags.attached.ls -r "com.vmware.content.Library:$id"
  assert_success "$tag_name"

  run govc tags.detach "$tag" /my-content
  assert_success

  run govc tags.attached.ls "$tag_name"
  assert_success ""

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
  assert_success
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

  run govc library.info -l "/my-content/ttylinux-live/$TTYLINUX_NAME.iso"
  assert_success

  run govc library.info -L /my-content/ttylinux-live/
  assert_success
  assert_matches contentlib-
  file="$output"
  run govc datastore.ls "$file"
  assert_matches "$TTYLINUX_NAME.iso"

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

  run env GOVC_DATASTORE="" govc library.deploy "my-content/$TTYLINUX_NAME" ttylinux3 # datastore is not required
  assert_success

  run govc vm.destroy ttylinux ttylinux2 ttylinux3
  assert_success

  item_id=$(govc library.info -json "/my-content/$TTYLINUX_NAME" | jq -r .[].id)

  run govc datastore.rm "contentlib-$library_id/$item_id" # remove library files out-of-band, forcing a deploy error below
  assert_success

  run govc library.deploy "my-content/$TTYLINUX_NAME" ttylinux2
  assert_failure
}

@test "library.clone ovf" {
  vcsim_env

  vm=DC0_H0_VM0
  item="${vm}_item"

  run govc library.create my-content
  assert_success

  run govc library.clone -vm $vm -ovf my-content $item
  assert_success

  run govc vm.destroy $vm
  assert_success

  run govc library.ls my-content/
  assert_success /my-content/$item

  run govc library.create mirror
  assert_success

  run govc library.cp /my-content/$item /mirror
  assert_success

  run govc library.ls mirror/
  assert_success /mirror/$item
}

@test "library.deploy vmtx" {
  vcsim_env

  vm=DC0_H0_VM0
  item="${vm}_item"

  run govc vm.clone -library enoent -vm $vm $item
  assert_failure # library does not exist

  run govc library.create my-content
  assert_success

  run govc library.deploy my-content/$item my-vm
  assert_failure # vmtx item does not exist

  run govc library.clone -vm $vm my-content $item
  assert_success

  run govc library.deploy my-content/$item my-vm
  assert_success

  run govc library.checkout my-content/enoent my-vm-checkout
  assert_failure # vmtx item does not exist

  run govc library.checkout my-content/$item my-vm-checkout
  assert_success

  run govc library.checkin -vm my-vm-checkout my-content/enoent
  assert_failure # vmtx item does not exist

  run govc library.checkin -vm my-vm-checkout my-content/$item
  assert_success

  run govc object.collect -s vm/$item config.template
  assert_success "true"

  run govc object.collect -s vm/$item summary.config.template
  assert_success "true"

  run govc vm.destroy $item
  assert_success # expected to delete the CL item too

  run govc library.deploy my-content/$item my-vm2
  assert_failure # $item no longer exists
}

@test "library.pubsub" {
  vcsim_env

  run govc library.create -pub published-content
  assert_success
  id="$output"

  url="https://$(govc env GOVC_URL)/cls/vcsp/lib/$id"

  run govc library.info published-content
  assert_success
  assert_matches "Publication:"
  assert_matches "$url"

  run govc library.import published-content "$GOVC_IMAGES/ttylinux-latest.ova"
  assert_success

  run govc library.create -sub "$url" my-content
  assert_success

  run govc library.info my-content
  assert_success
  assert_matches "Subscription:"
  assert_matches "$url"

  run govc library.import my-content "$GOVC_IMAGES/$TTYLINUX_NAME.iso"
  assert_failure # cannot add items to subscribed libraries

  run govc library.ls my-content/ttylinux-latest/
  assert_success
  assert_matches "/my-content/ttylinux-latest/ttylinux-pc_i486-16.1.ovf"

  run govc library.sync my-content
  assert_success
}

@test "library.subscriber example" {
  vcsim_start -ds 3

  ds0=LocalDS_0
  ds1=LocalDS_1
  ds2=LocalDS_2
  pool=DC0_C0/Resources

  # Create a published library with OVA items
  govc library.create -ds $ds0 -pub ttylinux-pub-ovf

  govc library.import ttylinux-pub-ovf "$GOVC_IMAGES/ttylinux-pc_i486-16.1.ova"

  url="$(govc library.info -U ttylinux-pub-ovf)"
  echo "$url"

  # Create a library subscribed to the publish-content library
  govc library.create -ds $ds0 -sub "$url" govc-sub-ovf

  # Create a library to contain VM Templates, and enabling publishing
  govc library.create -ds $ds0 -pub govc-pub-vmtx

  url="$(govc library.info -U govc-pub-vmtx)"
  echo "$url"

  # Create vm inventory folder to contain govc-pub-vmtx library templates
  govc folder.create vm/govc-pub-vmtx

  # Convert govc-sub-ovf's OVA items to VMTX items in the govc-pub-vmtx library
  govc library.sync -folder govc-pub-vmtx -pool $pool -vmtx govc-pub-vmtx govc-sub-ovf

  # No existing subscribers
  govc library.subscriber.ls govc-pub-vmtx

  for ds in $ds1 $ds2 ; do
    # Create a library subscribed to the govc-pub-vmtx library
    govc library.create -ds $ds -sub "$url" govc-sub-vmtx-$ds

    # Create vm inventory folder to contain sub-content library templates
    govc folder.create vm/govc-sub-vmtx-$ds

    # Create a subscriber to which the VM Templates can be published
    govc library.subscriber.create -folder govc-sub-vmtx-$ds -pool $pool govc-pub-vmtx govc-sub-vmtx-$ds
  done

  govc library.subscriber.ls govc-pub-vmtx | grep govc-sub-vmtx-$ds1
  govc library.subscriber.ls govc-pub-vmtx | grep govc-sub-vmtx-$ds2

  # Expect 1 VM: govc-pub-vmtx/ttylinux-pc_i486-16.1
  govc find vm -type f -name govc-* | $xargs -n1 -r govc find -type m

  # Publish entire library
  govc library.publish govc-pub-vmtx

  # Publish a specific item
  govc library.publish govc-pub-vmtx/ttylinux-pc_i486-16.1

  # Expect 2 more VMs: govc-sub-vmtx-{$ds1,$ds2}
  govc find vm -type f -name govc-* | $xargs -n1 govc find -type m

  for ds in $ds1 $ds2 ; do
    govc vm.clone -link -vm govc-sub-vmtx-$ds/ttylinux-pc_i486-16.1 -ds $ds -pool $pool -folder govc-sub-vmtx-$ds ttylinux
  done
}

@test "library.findbyid" {
  vcsim_env

  run govc library.create my-content
  assert_success
  id="$output"

  run govc library.create my-content
  assert_success

  run govc library.import my-content library.bats
  assert_failure # "my-content" matches 2 items

  run govc library.import "$id" library.bats
  assert_success # using id to find library

  n=$(govc library.info my-content | grep -c Name:)
  [ "$n" == 2 ]

  n=$(govc library.info "$id" | grep -c Name:)
  [ "$n" == 1 ]

  run govc library.rm my-content
  assert_failure # "my-content" matches 2 items

  run govc library.rm "$id"
  assert_success

  n=$(govc library.info my-content | grep -c Name:)
  [ "$n" == 1 ]
}
