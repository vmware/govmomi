#!/usr/bin/env bats

load test_helper

@test "disk.ls" {
  vcsim_start -ds 2

  run govc disk.ls
  assert_success

  run govc disk.ls -ds LocalDS_0
  assert_success

  run govc disk.ls enoent
  assert_failure
}

@test "disk.create" {
  vcsim_env

  name=$(new_id)

  run govc disk.create -size 10M "$name"
  assert_success
  id="$output"

  run govc disk.ls "$id"
  assert_success

  govc disk.ls -json "$id" | jq .

  vm=DC0_H0_VM0

  run govc disk.attach -vm $vm "$id"
  assert_success

  run govc disk.detach -vm $vm "$id"
  assert_success

  run govc disk.rm "$id"
  assert_success

  run govc disk.rm "$id"
  assert_failure

  name=$(new_id)
  run govc disk.create -profile enoent "$name"
  assert_failure # profile does not exist

  run govc disk.create -profile "vSAN Default Storage Policy" "$name"
  assert_success
}

@test "disk.create -datastore-cluster" {
  vcsim_env -pod 1 -ds 3 -cluster 2

  pod=/DC0/datastore/DC0_POD0
  id=$(new_id)

  run govc disk.create -datastore-cluster $pod "$id"
  assert_failure

  run govc find $pod -type s
  assert_success
  [ ${#lines[@]} -eq 0 ]

  run govc object.mv /DC0/datastore/LocalDS_{1,2} $pod
  assert_success

  run govc find $pod -type s
  assert_success
  [ ${#lines[@]} -eq 2 ]

  run govc disk.create -datastore-cluster $pod -size 10M "$id"
  assert_success

  id=$(new_id)
  pool=$GOVC_RESOURCE_POOL
  unset GOVC_RESOURCE_POOL
  run govc disk.create -datastore-cluster $pod -size 10M "$id"
  assert_failure # -pool is required

  run govc disk.create -datastore-cluster $pod -size 10M -pool "$pool" "$id"
  assert_success
}

@test "disk.register" {
  vcsim_env

  id=$(new_id)
  vmdk="$id/$id.vmdk"

  run govc datastore.mkdir "$id"
  assert_success

  # create with VirtualDiskManager
  run govc datastore.disk.create -size 10M "$vmdk"
  assert_success

  run govc disk.register "$id" "$id"
  assert_failure # expect fail for directory

  run govc disk.register "" "$id"
  assert_failure # expect fail for empty path

  run govc disk.register "$vmdk" "$id"
  assert_success
  id="$output"

  run govc disk.ls "$id"
  assert_success

  run govc disk.register "$vmdk" "$id"
  assert_failure

  run govc disk.rm "$id"
  assert_success

  run govc disk.rm "$id"
  assert_failure
}

@test "disk.snapshot" {
  vcsim_env

  name=$(new_id)

  run govc disk.create -size 10M "$name"
  assert_success
  id="$output"

  run govc disk.snapshot.ls "$id"
  assert_success

  run govc disk.snapshot.create "$id"
  assert_success
  sid="$output"

  govc disk.snapshot.ls "$id" | grep "$sid"

  govc disk.snapshot.ls -json "$id" | jq .

  run govc disk.snapshot.rm "$id" "$sid"
  assert_success

  run govc disk.snapshot.rm "$id" "$sid"
  assert_failure

  run govc disk.rm "$id"
  assert_success
}

@test "disk.tags" {
  vcsim_env

  run govc tags.category.create region
  assert_success

  run govc tags.create -c region US-WEST
  assert_success

  name=$(new_id)

  run govc disk.create -size 10M "$name"
  assert_success
  id="$output"

  run govc disk.ls "$id"
  assert_success

  run govc disk.ls -c region -t US-WEST
  assert_success ""

  govc disk.ls -T | grep -v US-WEST

  run govc disk.tags.attach -c region US-WEST "$id"
  assert_success

  run govc disk.ls -c region -t US-WEST
  assert_success
  assert_matches "$id"

  run govc disk.ls -T
  assert_success
  assert_matches US-WEST

  run govc disk.tags.detach -c region enoent "$id"
  assert_failure

  run govc disk.tags.detach -c region US-WEST "$id"
  assert_success

  govc disk.ls -T | grep -v US-WEST
}

@test "disk.metadata" {
  vcsim_env

  run govc disk.create -size 10M my-disk
  assert_success
  id="$output"

  run govc disk.metadata.ls "$id"
  assert_success ""

  run govc disk.metadata.ls -json "$id"
  assert_success null

  run govc disk.metadata.update "$id" foo=bar biz=baz
  assert_success

  run govc disk.metadata.ls "$id"
  assert_success
  assert_output_lines 2

  run govc disk.metadata.ls -json "$id"
  assert_success

  run govc disk.metadata.ls -K foo "$id"
  assert_success
  assert_output_lines 1

  run govc disk.metadata.update "$id" foo2=bar2 biz2=baz2
  assert_success

  run govc disk.metadata.ls "$id"
  assert_success
  assert_output_lines 4

  run govc disk.metadata.ls -p foo "$id"
  assert_success
  assert_output_lines 2

  run govc disk.metadata.update -d foo2 -d biz2 "$id"
  assert_success

  run govc disk.metadata.ls "$id"
  assert_success
  assert_output_lines 2
}

@test "disk.reconcile" {
  vcsim_env

  name=$(new_id)

  run govc disk.create -size 10M "$name"
  assert_success
  id="$output"

  run govc disk.create -size 10M "$(new_id)"
  assert_success

  run govc disk.ls
  assert_success
  [ ${#lines[@]} -eq 2 ]

  path=$(govc disk.ls -json "$id" | jq -r .objects[].config.backing.filePath)
  run govc datastore.rm "$path"
  assert_success

  # file backing was removed without using disk.rm, results in NotFound fault
  run govc disk.ls "$id"
  assert_failure

  run govc disk.ls
  assert_success
  [ ${#lines[@]} -eq 1 ]

  run govc disk.ls -a
  assert_success
  [ ${#lines[@]} -eq 2 ]
  assert_matches "not found"

  run govc disk.ls -R
  assert_success
}

@test "disk global catalog" {
  vcsim_start -ds 3

  run govc disk.create -ds LocalDS_0 -size 10M disk-0
  assert_success
  id0="$output"

  run govc disk.create -ds LocalDS_1 -size 10M disk-1
  assert_success
  id1="$output"

  run govc disk.snapshot.create "$id0" snapshot-0
  assert_success
  sid0="$output"

  run govc disk.snapshot.ls "$id0"
  assert_success
  assert_matches "$sid0"
  assert_matches "snapshot-0"

  run govc disk.snapshot.rm "$id0" "$sid0"
  assert_success
  sid0="$output"

  run govc disk.ls
  assert_success
  [ ${#lines[@]} -eq 2 ]

  run govc disk.ls -ds LocalDS_0
  assert_success
  [ ${#lines[@]} -eq 1 ]

  run govc disk.ls -ds LocalDS_1
  assert_success
  [ ${#lines[@]} -eq 1 ]

  run govc disk.ls -ds LocalDS_2
  assert_success
  [ ${#lines[@]} -eq 0 ]

  vm=DC0_H0_VM0

  run govc disk.attach -vm $vm "$id0"
  assert_success

  run govc disk.detach -vm $vm "$id0"
  assert_success

  run govc disk.rm -ds LocalDS_1 "$id0"
  assert_failure

  run govc disk.rm -ds LocalDS_0 "$id0"
  assert_success

  run govc disk.rm "$id1"
  assert_success

  run govc disk.ls
  assert_success
  [ ${#lines[@]} -eq 0 ]
}

@test "disk query" {
  vcsim_start -ds 3

  name=0
  size=10

  for ds in $(govc find / -type s) ; do
    for prefix in alpha beta ; do
      run govc disk.create -ds "$ds" -size ${size}M $prefix-disk-$name
      assert_success
      id="$output"

      run govc disk.metadata.update "$id" \
          namespace=ns-$prefix \
          name=vol-$prefix-$name
      assert_success
    done

    name=$((name + 1))
    size=$((size + 10))
  done

  run govc disk.ls
  assert_success
  assert_output_lines 6

  run govc disk.ls -q capacity.eq=10
  assert_success
  assert_output_lines 2

  run govc disk.ls -q capacity.gt=10
  assert_success
  assert_output_lines 4

  run govc disk.ls -q capacity.ge=10 -q name.sw=alpha-
  assert_success
  assert_output_lines 3

  run govc disk.ls -q metadataKey.eq=namespace -q metadataValue.eq=ns-alpha
  assert_success
  assert_output_lines 3
}
