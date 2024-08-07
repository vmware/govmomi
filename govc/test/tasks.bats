#!/usr/bin/env bats

load test_helper

@test "tasks" {
  vcsim_env

  run govc tasks
  assert_success

  run govc tasks vm/DC0_H0_VM0
  assert_success
  assert_matches PowerOn

  run govc tasks vm/DC0_H0_VM0 vm/DC0_H0_VM1
  assert_failure # > 1 arg

  run govc tasks 'host/*'
  assert_failure # matches 2 objects

  run govc tasks -b 1h
  assert_success
  [ ${#lines[@]} -gt 10 ]

  run govc tasks -r /DC0/vm
  assert_success
  assert_matches CreateVm
  assert_matches PowerOn
}

@test "tasks esx" {
  vcsim_env -esx

  run govc tasks
  assert_success

  run govc tasks -b 1h
  assert_failure # TaskHistoryCollector not supported on ESX

  run govc tasks -r
  assert_failure # TaskHistoryCollector not supported on ESX
}

@test "task.create" {
  vcsim_env

  export GOVC_SHOW_UNRELEASED=true

  run govc task.create enoent
  assert_failure

  id=$(govc extension.info -json | jq -r '.extensions[].taskList | select(. != null) | .[].taskID' | head -1)
  assert_equal com.vmware.govmomi.simulator.test "$id"

  run govc task.create "$id"
  assert_success
  task="$output"

  run govc task.cancel "$task"
  assert_success
  cancelled=$(govc tasks -json | jq ".tasks[] | select(.key == \"$task\") | .cancelled")
  assert_equal "true" "$cancelled"

  run govc task.create "$id"
  assert_success
  task="$output"

  run govc task.set -s running "$task"
  assert_success

  govc tasks | grep "$id"

  run govc task.set -d "$id.init" -m "task init" "$task"
  assert_success

  run govc task.set -p 42 "$task"
  assert_success

  run govc task.set -s success "$task"
  assert_success

  run govc task.set -s running "$task"
  assert_failure
}
