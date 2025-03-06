#!/usr/bin/env bats

load test_helper

@test "volume.ls" {
  vcsim_env

  run govc volume.ls
  assert_success ""

  for i in $(seq 1 6); do
    if [ $(( i % 2)) -eq 0 ] ; then
      seq=even
    else
      seq=odd
    fi

    run env GOVC_SHOW_UNRELEASED=true govc volume.create \
        -cluster-id super-$seq -label "pid=$$" -label "seq=$seq" my-volume-"$i"
    assert_success
  done

  run govc volume.ls
  assert_success
  assert_output_lines 6

  run govc volume.ls -L -n my-volume-2
  assert_success
  assert_output_lines 1

  run govc volume.ls $output
  assert_success
  assert_output_lines 1

  run govc volume.ls -label "pid=$$"
  assert_success
  assert_output_lines 6

  run govc volume.ls -L -label seq=even
  assert_success
  assert_output_lines 3

  run govc volume.ls $output
  assert_success
  assert_output_lines 3

  run govc volume.ls -H green
  assert_success
  assert_output_lines 6

  run govc volume.ls -H red
  assert_success ""

  run govc volume.ls -profile "vSAN Default Storage Policy"
  assert_success ""

  run env GOVC_SHOW_UNRELEASED=true govc volume.create \
      -cluster-id my-cluster -profile "vSAN Default Storage Policy" another-volume
  assert_success
  id=$output

  run govc volume.ls -profile "vSAN Default Storage Policy"
  assert_success
  assert_output_lines 1

  run govc volume.ls -c invalid-cluster
  assert_success ""

  run govc volume.ls -c super-odd
  assert_success
  assert_output_lines 3

  run govc volume.ls -c super-even -c my-cluster
  assert_success
  assert_output_lines 4

  run govc volume.ls -json
  assert_success

  run govc disk.metadata.ls -K cns.containerCluster.clusterId "$id"
  assert_success
  assert_matches "my-cluster"

  run govc volume.ls -b "$id" # -b invokes CnsQueryVolumeInfo
  assert_success

  run govc volume.ls -json -b -n another-volume
  assert_success

  run jq -r .info[].volumeInfo.vStorageObject.config.backing.filePath <<<"$output"
  assert_success
  path="$output"

  run govc datastore.disk.info "$path"
  assert_success
}

@test "volume.create -disk-id" {
  vcsim_env

  run env GOVC_SHOW_UNRELEASED=true govc volume.create \
      -cluster-id my-cluster -disk-id invalid my-volume
  assert_failure

  run govc disk.create -size 10M my-disk
  assert_success
  id="$output"

  run govc disk.ls "$id"
  assert_success

  run env GOVC_SHOW_UNRELEASED=true govc volume.create \
      -cluster-id my-cluster -disk-id "$id" my-volume
  assert_success
  vol="$output"

  run govc volume.ls "$vol"
  assert_success

  run govc volume.extend -size 20M "$id"
  assert_success

  run govc volume.rm "$vol"
  assert_success

  run govc disk.ls "$id"
  assert_failure

  run govc volume.extend -size 30M "$id"
  assert_failure
}

@test "volume.rm" {
  vcsim_env

  export GOVC_SHOW_UNRELEASED=true

  run govc volume.create -cluster-id my-cluster my-volume
  assert_success
  id="$output"

  run govc disk.ls "$id"
  assert_success

  run govc volume.rm invalid
  assert_failure

  run govc volume.rm "$id"
  assert_success

  run govc volume.rm "$id"
  assert_failure

  run govc disk.ls "$id"
  assert_failure

  run govc volume.create -cluster-id my-cluster my-volume
  assert_success
  id="$output"

  run govc volume.rm -keep "$id"
  assert_success

  run govc disk.ls "$id"
  assert_success
}

@test "volume.snapshot" {
  vcsim_env

  run govc volume.snapshot.ls
  assert_failure

  run govc volume.snapshot.rm
  assert_failure

  run govc volume.snapshot.create
  assert_failure
}
