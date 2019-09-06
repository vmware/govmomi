#!/usr/bin/env bats

load test_helper

@test "datacenter.info" {
  vcsim_env -esx

  dc=$(govc ls -t Datacenter / | head -n1)
  run govc datacenter.info "$dc"
  assert_success

  run govc datacenter.info -json "$dc"
  assert_success

  run govc datacenter.info /enoent
  assert_failure
}

@test "datacenter.info with folders" {
  vcsim_start -cluster 3 -folder 1

  info=$(govc datacenter.info DC0)

  hosts=$(govc find -type h | wc -l)
  clusters=$(govc find -type c | wc -l)
  vms=$(govc find -type m | wc -l)
  datastores=$(govc find -type s | wc -l)


  assert_equal "$hosts" "$(grep Hosts: <<<"$info" | awk '{print $2}')"
  assert_equal "$clusters" "$(grep Clusters: <<<"$info" | awk '{print $2}')"
  assert_equal "$vms" "$(grep "Virtual Machines": <<<"$info" | awk '{print $3}')"
  assert_equal "$datastores" "$(grep Datastores: <<<"$info" | awk '{print $2}')"
}

@test "datacenter.create" {
  vcsim_env
  unset GOVC_DATACENTER

  # name not specified
  run govc datacenter.create
  assert_failure

  dcs=($(new_id) $(new_id))
  run govc datacenter.create "${dcs[@]}"
  assert_success

  for dc in ${dcs[*]}; do
    run govc ls "/$dc"
    assert_success
    # /<datacenter>/{vm,network,host,datastore}
    [ ${#lines[@]} -eq 4 ]

    run govc datacenter.info "/$dc"
    assert_success
  done

  run govc object.destroy "/$dc"
  assert_success
}

@test "datacenter commands fail against ESX" {
  vcsim_env -esx

  run govc datacenter.create something
  assert_failure

  run govc object.destroy /ha-datacenter
  assert_failure
}
