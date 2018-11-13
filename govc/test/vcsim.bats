#!/usr/bin/env bats

load test_helper

@test "vcsim rbvmomi" {
  if ! ruby -e "require 'rbvmomi'" ; then
    skip "requires rbvmomi"
  fi

  vcsim_env

  ruby ./vcsim_test.rb "$(govc env -x GOVC_URL_PORT)"
}

@test "vcsim examples" {
  vcsim_env

  # compile + run examples against vcsim
  for main in ../../examples/*/main.go ; do
    run go run "$main" -insecure -url "$GOVC_URL"
    assert_success
  done
}

@test "vcsim about" {
  vcsim_env -dc 2 -cluster 3 -vm 0 -ds 0

  url="https://$(govc env GOVC_URL)"

  run curl -skf "$url/about"
  assert_matches "CurrentTime" # 1 param (without Context)
  assert_matches "TerminateSession" # 2 params (with Context)

  run curl -skf "$url/debug/vars"
  assert_success

  model=$(curl -sfk "$url/debug/vars" | jq .vcsim.Model)
  [ "$(jq .Datacenter <<<"$model")" == "2" ]
  [ "$(jq .Cluster <<<"$model")" == "6" ]
  [ "$(jq .Machine <<<"$model")" == "0" ]
  [ "$(jq .Datastore <<<"$model")" == "0" ]
}

@test "vcsim host placement" {
  vcsim_start -dc 0

  # https://github.com/vmware/govmomi/issues/1258
  id=$(new_id)
  govc datacenter.create DC0
  govc cluster.create comp
  govc cluster.add -cluster comp -hostname test.host.com -username user -password pass
  govc cluster.add -cluster comp -hostname test2.host.com -username user -password pass
  govc datastore.create -type local -name vol6 -path "$TMPDIR" test.host.com
  govc pool.create comp/Resources/testPool
  govc vm.create -c 1 -ds vol6 -g centos64Guest -pool testPool -m 4096 "$id"
  govc vm.destroy "$id"
}

@test "vcsim set vm properties" {
  vcsim_env

  vm=/DC0/vm/DC0_H0_VM0

  run govc object.collect $vm guest.ipAddress
  assert_success ""

  run govc vm.change -vm $vm -e SET.guest.ipAddress=127.0.0.1
  assert_success

  run govc object.collect -s $vm guest.ipAddress
  assert_success "127.0.0.1"

  run govc vm.info -vm.ip 127.0.0.1
  assert_success

  run govc object.collect -s $vm guest.hostName
  assert_success ""

  run govc vm.change -vm $vm -e SET.guest.hostName=localhost.localdomain
  assert_success

  run govc object.collect -s $vm guest.hostName
  assert_success "localhost.localdomain"

  run govc vm.info -vm.dns localhost.localdomain
  assert_success

  uuid=$(uuidgen)
  run govc vm.change -vm $vm -e SET.config.uuid="$uuid"
  assert_success

  run govc object.collect -s $vm config.uuid
  assert_success "$uuid"
}
