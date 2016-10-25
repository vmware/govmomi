#!/usr/bin/env bats

load test_helper

@test "object.destroy" {
    run govc object.destroy "/enoent"
    assert_failure

    run govc object.destroy
    assert_failure

    vm=$(new_id)
    run govc vm.create "$vm"
    assert_success

    # fails when powered on
    run govc object.destroy "vm/$vm"
    assert_failure

    run govc vm.power -off "$vm"
    assert_success

    run govc object.destroy "vm/$vm"
    assert_success
}

@test "object.rename" {
  run govc object.rename "/enoent" "nope"
  assert_failure

  vm=$(new_id)
  run govc vm.create -on=false "$vm"
  assert_success

  run govc object.rename "vm/$vm" "${vm}-renamed"
  assert_success

  run govc object.rename "vm/$vm" "${vm}-renamed"
  assert_failure

  run govc object.destroy "vm/${vm}-renamed"
  assert_success
}

@test "object.mv" {
  vcsim_env

  folder=$(new_id)

  run govc folder.create "vm/$folder"
  assert_success

  for _ in $(seq 1 3) ; do
    vm=$(new_id)
    run govc vm.create -folder "$folder" "$vm"
    assert_success
  done

  result=$(govc ls "vm/$folder" | wc -l)
  [ "$result" -eq "3" ]

  run govc folder.create "vm/${folder}-2"
  assert_success

  run govc object.mv "vm/$folder/*" "vm/${folder}-2"
  assert_success

  result=$(govc ls "vm/${folder}-2" | wc -l)
  [ "$result" -eq "3" ]

  result=$(govc ls "vm/$folder" | wc -l)
  [ "$result" -eq "0" ]
}
