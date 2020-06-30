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
