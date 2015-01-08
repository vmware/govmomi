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
  id=$(new_id)

  run govc pool.create $id
  assert_success

  path="*/Resources/$id"

  run govc vm.create -pool $path $id
  assert_success

  run govc pool.destroy $path
  assert_success
}
