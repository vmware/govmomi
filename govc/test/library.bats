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

  run govc library.update -n new-content my-content
  assert_success

  run govc library.info new-content
  assert_success
  assert_matches "$id"

  run govc library.rm /new-content
  assert_success
}

@test "library.import" {
  vcsim_env
  unset GOVC_HOST # else datastore.download tries via ESX

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

  run govc library.info -l -L "/my-content/$TTYLINUX_NAME/*.ovf"
  assert_success
  ovf="$output"

  # validate ovf.Envelope as json
  run govc datastore.download -json "$ovf" -
  assert_success

  run jq -r .virtualSystem.operatingSystemSection.osType <<<"$output"
  assert_success "otherLinuxGuest"

  run govc library.info -L "/my-content/$TTYLINUX_NAME/*.vmdk"
  assert_success
  vmdk="$output"

  # validate vmdk.Descriptor as json
  run govc datastore.download -json "$vmdk" -
  assert_success

  base=$(basename "$vmdk")
  run jq -r .extent[].info <<<"$output"
  assert_success "${base/.vmdk/-flat.vmdk}"

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
  run govc library.update -n ttylinux_latest my-content/ttylinux-latest
  assert_success
  run govc library.ls "/my-content/ttylinux_latest/*"
  assert_success
  assert_matches "$TTYLINUX_NAME.ovf"
  assert_matches "$TTYLINUX_NAME-disk1.vmdk"

  summary="ISO \[${GOVC_DATASTORE}\] contentlib-.*${TTYLINUX_NAME}.iso"

  run govc vm.create -on=false -iso "library:/my-content/ttylinux-live/$TTYLINUX_NAME.iso" library-iso-test
  assert_success

  run govc device.info -vm library-iso-test cdrom-*
  assert_success
  assert_matches "$summary"

  run govc device.cdrom.eject -vm library-iso-test
  assert_success

  run govc device.cdrom.insert -vm library-iso-test "library:/my-content/ttylinux-live/$TTYLINUX_NAME.iso"
  assert_success

  run govc device.info -vm library-iso-test cdrom-*
  assert_success
  assert_matches "$summary"
}

@test "library.deploy" {
  vcsim_env

  run govc library.create my-content
  assert_success
  library_id="$output"

  # link ovf/ova to datastore so we can test library.import with an http source
  dir=$(govc datastore.info -json | jq -r .datastores[].info.url)
  ln -s "$GOVC_IMAGES/$TTYLINUX_NAME."* "$dir"

  run govc library.import -c fake -pull my-content -n invalid-sha1 "https://$(govc env GOVC_URL)/folder/$TTYLINUX_NAME.ovf"
  assert_failure # invalid checksum

  sum=$(sha256sum "$GOVC_IMAGES/$TTYLINUX_NAME.ovf" | awk '{print $1}')
  run govc library.import -c "$sum" -pull my-content "https://$(govc env GOVC_URL)/folder/$TTYLINUX_NAME.ovf"
  assert_success

  run govc library.info -s "/my-content/ttylinux-live/$TTYLINUX_NAME.ovf"
  assert_success

  run env GOVC_SHOW_UNRELEASED=true govc library.info -S "/my-content/ttylinux-live/$TTYLINUX_NAME.ovf"
  assert_success

  run govc library.info -l -s /my-content/$TTYLINUX_NAME/$TTYLINUX_NAME.ovf
  assert_success

  run govc library.deploy "my-content/$TTYLINUX_NAME" ttylinux
  assert_success

  run govc vm.info ttylinux
  assert_success

  run govc library.import -pull -c fake -a MD5 -n invalid-md5 my-content "https://$(govc env GOVC_URL)/folder/$TTYLINUX_NAME.ova"
  assert_failure # invalid checksum

  sum=$(md5sum "$GOVC_IMAGES/$TTYLINUX_NAME.ova" | awk '{print $1}')
  run govc library.import -pull -c "$sum" -a MD5 -n ttylinux-unpacked my-content "https://$(govc env GOVC_URL)/folder/$TTYLINUX_NAME.ova"
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

  run govc library.deploy "my-content/$TTYLINUX_NAME" -options "$BATS_TMPDIR/ttylinux.json"
  assert_failure # see issue #2599

  run govc library.deploy -options "$BATS_TMPDIR/ttylinux.json" "my-content/$TTYLINUX_NAME"
  assert_success
  rm "$BATS_TMPDIR/ttylinux.json"

  run govc vm.info -r ttylinux2
  assert_success
  assert_matches DC0_DVPG0
  assert_matches 32MB
  assert_matches "1 vCPU"

  run env GOVC_DATASTORE="" govc library.deploy "my-content/$TTYLINUX_NAME" ttylinux3 # datastore is not required
  assert_success

  run govc vm.destroy ttylinux ttylinux2 ttylinux3
  assert_success

  config=$(base64 <<<'
<obj xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:type="vim25:VirtualMachineConfigSpec" xmlns:vim25="urn:vim25">
 <numCPUs>4</numCPUs>
 <memoryMB>2048</memoryMB>
</obj>')

  run env GOVC_SHOW_UNRELEASED=true govc library.deploy -config "$config" "my-content/$TTYLINUX_NAME" ttylinux
  assert_success

  run govc vm.info -r ttylinux
  assert_success
  assert_matches 2048MB
  assert_matches "4 vCPU"

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

  run govc library.clone -vm $vm -ovf -e -m my-content $item
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

@test "library.vmtx.info" {
  vcsim_env

  vm=DC0_H0_VM0
  item="${vm}_item"

  run govc library.create my-content
  assert_success

  run govc library.clone -vm $vm my-content $item
  assert_success

  run govc library.vmtx.info my-content/$item
  assert_success
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

  run govc library.info -L published-content/ttylinux-latest/*.vmdk
  assert_success
  vmdk="$output"

  run govc datastore.ls "$vmdk"
  assert_success

  run govc datastore.ls "${vmdk/.vmdk/-flat.vmdk}"
  assert_success

  run govc library.create -sub "$url" my-content
  assert_success

  run govc library.info my-content
  assert_success
  assert_matches "Subscription:"
  assert_matches "$url"

  run govc library.info -L my-content/ttylinux-latest/*.vmdk
  assert_success
  vmdk="$output"

  run govc datastore.ls "$vmdk"
  assert_success

  run govc datastore.ls "${vmdk/.vmdk/-flat.vmdk}"
  assert_success

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
  export GOVC_NETWORK=/DC0/network/DC0_DVPG0

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

@test "library.create.withpolicy" {
  vcsim_env

  policy_id=$(govc library.policy.ls -json | jq '.[][0].policy' -r)
  echo "$policy_id"

  run govc library.create -policy=foo secure-content
  assert_failure

  run govc library.create -policy "$policy_id" secure-content
  assert_success

  library_secpol=$(govc library.info -json secure-content | jq '.[].security_policy_id' -r)
  assert_equal "$library_secpol" "$policy_id"

  run govc library.import secure-content "$GOVC_IMAGES/ttylinux-latest.ova"
  assert_success

  run govc library.info -json secure-content/ttylinux-latest
  assert_success

  assert_equal false "$(jq -r <<<"$output" .[].security_compliance)"

  assert_equal NOT_AVAILABLE "$(jq -r <<<"$output" .[].certificate_verification_info.status)"

  run govc library.rm secure-content
  assert_success
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

@test "library.trust" {
  vcsim_env

  run govc library.trust.ls
  assert_success

  run govc library.trust.info enoent
  assert_failure # id does not exist

  run govc library.trust.rm enoent
  assert_failure # id does not exist

  pem=$(new_id)
  run govc extension.setcert -cert-pem ++ -org govc-library-trust "$pem" # generate a cert for testing
  assert_success

  run govc library.trust.create "$pem.crt"
  assert_success

  id=$(govc library.trust.ls | grep O=govc-library-trust | awk '{print $1}')
  run govc library.trust.info "$id"
  assert_success

  run govc library.trust.rm "$id"
  assert_success

  run govc library.trust.info "$id"
  assert_failure # id does not exist

  date > "$pem.crt"
  run govc library.trust.create "$id.crt"
  assert_failure # invalid cert

  # remove generated cert and key
  rm "$pem".{crt,key}
}

@test "library.session" {
  vcsim_env

  run govc library.session.ls
  assert_success

  run govc library.create my-content
  assert_success

  run govc library.import /my-content "$GOVC_IMAGES/$TTYLINUX_NAME.ova"
  assert_success

  run govc library.session.ls
  assert_success
  assert_matches ttylinux

  run govc library.session.ls -json
  assert_success

  run govc library.session.ls -json -i
  assert_success

  n=$(govc library.session.ls -json -i | jq '.files[] | length')
  assert_equal 2 "$n" # .ovf + .vmdk

  id=$(govc library.session.ls -json | jq -r .sessions[].id)

  run govc library.session.rm -i "$id" ttylinux-pc_i486-16.1.ovf
  assert_failure # removeFile not allowed in state DONE
  assert_matches "500 Internal Server Error"

  run govc library.session.rm "$id"
  assert_success
}

@test "library.probe" {
  vcsim_env

  export GOVC_SHOW_UNRELEASED=true

  run govc library.probe
  assert_failure

  run govc library.probe https://www.vmware.com
  assert_success

  run govc library.probe -f ftp://www.vmware.com
  if [ "$status" -ne 22 ]; then
    flunk $(printf "expected failed exit status=22, got status=%d" $status)
  fi

  run govc library.probe -json ftp://www.vmware.com
  assert_success
  assert_matches INVALID_URL
}

@test "library.evict" {
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

  run govc library.create -sub "$url" -sub-ondemand=true subscribed-content
  assert_success

  run govc library.info subscribed-content
  assert_success
  assert_matches "Subscription:"
  assert_matches "$url"

  run govc library.ls subscribed-content/ttylinux-latest/
  assert_success
  assert_matches "/subscribed-content/ttylinux-latest/ttylinux-pc_i486-16.1.ovf"

  run govc library.sync subscribed-content/ttylinux-latest
  assert_success

  # assert cached is false after item sync for ondemand library sans -f=true (force)
  cached=$(govc library.info subscribed-content/ttylinux-latest | grep Cached: | awk '{print $2}')
  assert_equal "false" "$cached"

  run govc library.sync -f=true subscribed-content/ttylinux-latest
  assert_success

  # assert cached is true after item sync with -f=true (force)
  cached=$(govc library.info subscribed-content/ttylinux-latest | grep Cached: | awk '{print $2}')
  assert_equal "true" "$cached"

  run govc library.evict subscribed-content/ttylinux-latest
  assert_success

  # assert cached is false after library item evict
  cached=$(govc library.info subscribed-content/ttylinux-latest | grep Cached: | awk '{print $2}')
  assert_equal "false" "$cached"
}
