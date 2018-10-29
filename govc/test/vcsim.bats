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
