#!/usr/bin/env bats

load test_helper

@test "pool.create" {
  id=$(new_id)

  run govc pool.create -cpu.shares low -mem.reservation 500 $id
  assert_success

  path="*/Resources/$id"
  run govc pool.info $path
  assert_success

  assert_line "Name: $id"
  assert_line "CPU Shares: low"
  assert_line "Mem Reservation: 500MB (expandable=true)"

  run govc pool.destroy $path
  assert_success
}

@test "pool.change" {
  id=$(new_id)

  run govc pool.create $id
  assert_success
  path="*/Resources/$id"

  run govc pool.change -pool $path -mem.shares high
  assert_success
  run govc pool.info $path
  assert_success
  assert_line "Mem Shares: high"
  assert_line "CPU Shares: normal"

  nid=$(new_id)
  run govc pool.change -pool $path -name $nid
  assert_success
  path="*/Resources/$nid"

  run govc pool.info $path
  assert_success
  assert_line "Name: $nid"

  run govc pool.destroy $path
  assert_success
}

@test "pool.destroy" {
  id=$(new_id)

  # should not be any existing test pools
  result=$(govc ls "host/*/Resources/govc-test-*" | wc -l)
  [ $result -eq 0 ]

  # parent pool
  run govc pool.create $id
  assert_success

  path="*/Resources/$id"

  result=$(govc ls "host/$path/*" | wc -l)
  [ $result -eq 0 ]

  # child pools
  run govc pool.create -pool $path $(new_id)
  assert_success

  run govc pool.create -pool $path $(new_id)
  assert_success

  # 2 child pools
  result=$(govc ls "host/$path/*" | wc -l)
  [ $result -eq 2 ]

  # 1 parent pool
  result=$(govc ls "host/*/Resources/govc-test-*" | wc -l)
  [ $result -eq 1 ]

  run govc pool.destroy -r $path
  assert_success

  # if we didn't -r, the child pools would end up here
  result=$(govc ls "host/*/Resources/govc-test-*" | wc -l)
  [ $result -eq 0 ]
}

@test "vm.create -pool" {
  # test with full inventory path to pools
  parent_path=$(govc ls 'host/*/Resources')
  parent_name=$(basename $parent_path)
  [ "$parent_name" = "Resources" ]

  child_name=$(new_id)
  child_path="$parent_path/$child_name"

  grand_child_name=$(new_id)
  grand_child_path="$child_path/$grand_child_name"

  run govc pool.create -pool $parent_path $child_name
  assert_success

  run govc pool.create -pool $child_path $grand_child_name
  assert_success

  for path in $parent_path $child_path $grand_child_path
  do
    run govc vm.create -on=false -pool $path $(new_id)
    assert_success
  done

  # test with glob inventory path to pools
  parent_path="*/$parent_name"
  child_path="$parent_path/$child_name"
  grand_child_path="$child_path/$grand_child_name"

  for path in $grand_child_path $child_path
  do
    run govc pool.destroy $path
    assert_success
  done
}

@test "vm.create -pool host" {
  id=$(new_id)

  path=$(govc ls host)

  run govc vm.create -on=false -pool enoent $id
  assert_failure "Error: resource pool 'enoent' not found"

  run govc vm.create -on=false -pool $path $id
  assert_success
}

@test "vm.create -pool cluster" {
  vcsim_env

  id=$(new_id)

  path=$(dirname $GOVC_HOST)

  unset GOVC_HOST
  unset GOVC_RESOURCE_POOL

  run govc vm.create -on=false -pool enoent $id
  assert_failure "Error: resource pool 'enoent' not found"

  run govc vm.create -on=false -pool $path $id
  assert_success
}
