#!/usr/bin/env bats

load test_helper

@test "ls" {
  run govc ls
  assert_success
  # /dc/{vm,network,host,datastore}
  [ ${#lines[@]} -ge 4 ]

  run govc ls host
  assert_success
  [ ${#lines[@]} -ge 1 ]

  run govc ls enoent
  assert_success
  [ ${#lines[@]} -eq 0 ]
}

@test "ls vm" {
  vm=$(new_empty_vm)

  run govc ls vm
  assert_success
  [ ${#lines[@]} -ge 1 ]

  run govc ls vm/$vm
  assert_success
  [ ${#lines[@]} -eq 1 ]

  run govc ls /*/vm/$vm
  assert_success
  [ ${#lines[@]} -eq 1 ]
}

@test "ls network" {
  run govc ls network
  assert_success
  [ ${#lines[@]} -ge 1 ]

  local path=${lines[0]}
  run govc ls $path
  assert_success
  [ ${#lines[@]} -eq 1 ]

  run govc ls network/$(basename $path)
  assert_success
  [ ${#lines[@]} -eq 1 ]

  run govc ls /*/network/$(basename $path)
  assert_success
  [ ${#lines[@]} -eq 1 ]
}

@test "ls multi ds" {
  vcsim_env

  run govc ls /DC*
  assert_success
  # /DC[0,1]/{vm,network,host,datastore}
  [ ${#lines[@]} -eq 8 ]
}
