#!/bin/bash -e

# This test is not run via bats.
# A VNC session will be opened to observe the VM boot order:
# 1) from floppy  (followed by: eject floppy, reboot)
# 2) from cdrom   (followed by: eject cdrom, reboot)
# 3) from network (will timeout)
# 4) from disk

. $(dirname $0)/test_helper.bash

upload_img
upload_iso

id=$(new_ttylinux_vm)

function cleanup() {
  govc vm.destroy $id
  quit_vnc $vnc
}

trap cleanup EXIT

govc device.cdrom.add -vm $id > /dev/null
govc device.cdrom.insert -vm $id $GOVC_TEST_ISO

govc device.floppy.add -vm $id > /dev/null
govc device.floppy.insert -vm $id $GOVC_TEST_IMG

govc device.boot -vm $id -delay 2000 -order floppy,cdrom,ethernet,disk

vnc=$(govc vm.vnc -vm $id -port 21122 -password govmomi | awk '{print $5}')

echo "booting from floppy..."
govc vm.power -on $id

open_vnc $vnc

sleep 10

govc vm.power -off $id

govc device.floppy.eject -vm $id

# this is ttylinux-live, notice the 'boot:' prompt vs 'login:' prompt when booted from disk
echo "booting from cdrom..."
govc vm.power -on $id

sleep 10

govc vm.power -off $id

govc device.cdrom.eject -vm $id

echo "booting from network, will timeout then boot from disk..."
govc vm.power -on $id

ip=$(govc vm.ip $id)

echo "VM booted from disk (ip=$ip)"

sleep 10
