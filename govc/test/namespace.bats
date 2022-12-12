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

@test "namespace.cluster.enable" {
    vcsim_env

    # need to set up some dependencies
    govc cluster.create WCP-Cluster
    assert_success

    govc dvs.create "DVPG-Management Network"
    assert_success

    govc namespace.cluster.enable \
      --service-cidr 10.96.0.0/23 \
      --pod-cidrs 10.244.0.0/20 \
      --cluster "WCP-Cluster" \
      --control-plane-dns 8.8.8.8 \
      --worker-dns 8.8.8.8 \
      --control-plane-dns-search-domains example.com \
      --control-plane-dns-names wcp.example.com \
      --control-plane-ntp-servers pool.ntp.org \
      --network-provider "NSXT_CONTAINER_PLUGIN" \
      --workload-network.egress-cidrs 10.0.0.128/26 \
      --workload-network.ingress-cidrs "10.0.0.64/26" \
      --workload-network.switch VDS \
      --workload-network.edge-cluster Edge-Cluster-1 \
      --size TINY   \
      --mgmt-network.mode STATICRANGE \
      --mgmt-network.network "DVPG-Management Network" \
      --mgmt-network.gateway 10.0.0.1 \
      --mgmt-network.starting-address 10.0.0.45 \
      --mgmt-network.subnet-mask 255.255.255.0 \
      --ephemeral-storage-policy "vSAN Default Storage Policy" \
      --control-plane-storage-policy "vSAN Default Storage Policy" \
      --image-storage-policy "vSAN Default Storage Policy"
    assert_success
}

@test "namespace.cluster.disable" {
    vcsim_env

    govc cluster.create WCP-Cluster
    assert_success

    govc namespace.cluster.disable --cluster WCP-Cluster
    assert_success
}

@test "namespace.logs" {
  vcsim_env

  id=$(govc find -i -maxdepth 0 host/DC0_C0 | awk -F: '{print $2}')

  run govc namespace.logs.download -cluster DC0_C0
  assert_success

  rm "wcp-support-bundle-$id-"*.tar

  govc namespace.logs.download -cluster DC0_C0 - | tar -xvOf-
}

@test "namespace.service.ls" {
    vcsim_env

    run govc namespace.service.ls
    assert_success
    assert_matches service1
    assert_matches service2

    run govc namespace.service.ls -l
    assert_success
    assert_matches ACTIVATED
    assert_matches mock-service-1
}

@test "namespace.service.info" {
    vcsim_env

    run govc namespace.service.info service1
    assert_success
    assert_matches mock-service-1
    assert_matches ACTIVATED
    assert_matches "Description of service1"

    run govc namespace.service.info -json service2
    assert_matches DE-ACTIVATED

}