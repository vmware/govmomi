#!/usr/bin/env bats

load test_helper

@test "namespace.cluster.ls" {
    vcsim_env

    run govc namespace.cluster.ls
    assert_success ""

    run govc namespace.cluster.ls -json
    assert_success "[]"

    run govc cluster.create WCP-cluster
    assert_success

    run govc namespace.cluster.ls
    assert_success /DC0/host/WCP-cluster

    run govc namespace.cluster.ls -l
    assert_success
    assert_matches RUNNING
    assert_matches READY

    id=$(govc namespace.cluster.ls -json | jq -r .[].cluster)

    run govc object.collect -s "ClusterComputeResource:$id" name
    assert_success WCP-cluster
}

@test "namespace.logs" {
  vcsim_env

  id=$(govc find -i -maxdepth 0 host/DC0_C0 | awk -F: '{print $2}')

  run govc namespace.logs.download -cluster DC0_C0
  assert_success

  rm "wcp-support-bundle-$id-"*.tar

  govc namespace.logs.download -cluster DC0_C0 - | tar -xvOf-
}
