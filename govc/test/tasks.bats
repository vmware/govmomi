#!/usr/bin/env bats

load test_helper

@test "tasks" {
  vcsim_env

  run govc tasks
  assert_success
}

@test "tasks host" {
  vcsim_env

  run govc tasks 'host/*'
  assert_success
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
