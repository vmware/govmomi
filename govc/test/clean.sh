#!/bin/bash

# cleanup any VMs or .{vmdk,iso,img} artifacts created by govc

. $(dirname $0)/test_helper.bash

teardown

datastore_rm() {
  name=$1
  govc datastore.rm $name 2> /dev/null
}

datastore_rm $GOVC_TEST_IMG
datastore_rm $GOVC_TEST_ISO
datastore_rm $GOVC_TEST_VMDK
datastore_rm $(echo $GOVC_TEST_VMDK | sed 's/.vmdk/-flat.vmdk/')

govc datastore.ls
