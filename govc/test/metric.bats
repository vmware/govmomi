#!/usr/bin/env bats

load test_helper

@test "metric.ls" {
  vcsim_env -esx

  run govc metric.ls
  assert_failure

  run govc metric.ls enoent
  assert_failure

  host=$(govc ls -t HostSystem ./... | head -n 1)
  pool=$(govc ls -t ResourcePool ./... | head -n 1)
  vm=$(govc ls -t VirtualMachine ./... | head -n 1)

  run govc metric.ls "$host"
  assert_success

  run govc metric.ls -json "$host"
  assert_success

  run govc metric.ls "$pool"
  assert_success

  run govc metric.ls "$vm"
  assert_success
}

@test "metric.sample" {
  vcsim_env

  host=$(govc ls -t HostSystem ./... | head -n 1)
  metrics=($(govc metric.ls "$host"))

  run govc metric.sample "$host" enoent
  assert_failure

  run govc metric.sample "$host" "${metrics[@]}"
  assert_success

  run govc metric.sample -instance - "$host" "${metrics[@]}"
  assert_success

  run govc metric.sample -json "$host" "${metrics[@]}"
  assert_success

  vm=vm/DC0_H0_VM0

  metrics=($(govc metric.ls "$vm"))

  run govc metric.sample -i day "$vm" "${metrics[@]}"
  assert_success

  run govc metric.sample -i 300 -json "$vm" "${metrics[@]}"
  assert_success

  run govc metric.sample $vm "${metrics[@]}"
  assert_success
}

@test "metric.info" {
  vcsim_env

  host=$(govc ls -t HostSystem ./... | head -n 1)
  metrics=($(govc metric.ls "$host"))

  run govc metric.info "$host" enoent
  assert_failure

  run govc metric.info "$host"
  assert_success

  run govc metric.info -json "$host"
  assert_success

  run govc metric.info -dump "$host"
  assert_success

  run govc metric.sample "$host" "${metrics[@]}"
  assert_success

  run govc metric.info "$host" "${metrics[@]}"
  assert_success

  run govc metric.info - "${metrics[@]}"
  assert_success
}

@test "metric manager" {
  vcsim_env

  moid=$(govc object.collect -s - content.perfManager)

  govc object.collect -json "$moid" | jq .
}
