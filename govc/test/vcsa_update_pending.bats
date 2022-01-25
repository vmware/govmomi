#!/usr/bin/env bats

load test_helper

@test "update.pending.list" {
  vcsim_env
  local output

  run govc vcsa.update.pending.list LAST_CHECK
  assert_success

  run govc vcsa.update.pending.list LOCAL
  assert_success

  run govc vcsa.update.pending.list LOCAL_AND_ONLINE
  assert_success

  run govc vcsa.update.pending.list  -json=True LAST_CHECK
  assert_success

  run govc vcsa.update.pending.list -json=True LOCAL
  assert_success

  run govc vcsa.update.pending.list -json=True LOCAL_AND_ONLINE
  assert_success

  run govc vcsa.update.pending.list -xml=True LAST_CHECK
  assert_success

  run govc vcsa.update.pending.list -xml=True LOCAL
  assert_success

  run govc vcsa.update.pending.list -xml=True LOCAL_AND_ONLINE
  assert_success

  run govc vcsa.update.pending.list
  assert_failure

}

@test "update.pending.get" {
  vcsim_env
  local output

  run govc vcsa.update.pending.get 7.0.3.00000
  assert_success

  run govc vcsa.update.pending.get -json=True 7.0.3.00000
  assert_success

  run govc vcsa.update.pending.get -xml=True 7.0.3.00000
  assert_success

  run govc vcsa.update.pending.get
  assert_failure
}

@test "update.pending.precheck" {
  vcsim_env
  local output

  run govc vcsa.update.pending.precheck 7.0.3.00000
  assert_success

  run govc vcsa.update.pending.precheck -json=True 7.0.3.00000
  assert_success

  run govc vcsa.update.pending.precheck -xml=True 7.0.3.00000
  assert_success

  run govc vcsa.update.pending.precheck
  assert_failure
}

@test "update.pending.stage" {
  vcsim_env
  local output

  run govc vcsa.update.pending.stage 7.0.3.00000
  assert_success

  run govc vcsa.update.pending.stage
  assert_failure
}

@test "update.pending.validate" {
  vcsim_env
  local output

  run govc vcsa.update.pending.validate 7.0.3.00000
  assert_failure

  run govc vcsa.update.pending.validate "key1=val1,key2=val2"
  assert_failure

  run govc vcsa.update.pending.validate 7.0.3.00000 "key1=val1,key2=val2"
  assert_success
}

@test "update.pending.stage-and-install" {
  vcsim_env
  local output

  run govc vcsa.update.pending.stage-and-install 7.0.3.00000
  assert_failure

  run govc vcsa.update.pending.stage-and-install "key1=val1,key2=val2"
  assert_failure

  run govc vcsa.update.pending.stage-and-install 7.0.3.00000 "key1=val1,key2=val2"
  assert_success
}

@test "update.pending.install" {
  vcsim_env
  local output

  run govc vcsa.update.pending.install 7.0.3.00000
  assert_failure

  run govc vcsa.update.pending.install "key1=val1,key2=val2"
  assert_failure

  run govc vcsa.update.pending.install 7.0.3.00000 "key1=val1,key2=val2"
  assert_success
}
