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
    assert_matches mock-service-2
    assert_matches DEACTIVATED
}

@test "namespace.service.info notfound" {
    vcsim_env

    run govc namespace.service.info non-existing-service
    assert_failure
    assert_matches "404 Not Found"
}

@test "namespace.service.info" {
    vcsim_env

    run govc namespace.service.info service1
    assert_success
    assert_matches mock-service-1
    assert_matches ACTIVATED

    run govc namespace.service.info -json service2
    assert_matches DEACTIVATED
}

@test "namespace.service.version.ls" {
    vcsim_env

    versions=$(govc namespace.service.version.ls -json service1)
    assert_equal "2" $(echo $versions | jq '. |  length')

    run govc namespace.service.version.ls service1
    assert_success
    assert_matches "1.0.0"
    assert_matches "2.0.0"
}

@test "namespace.service.version.info notfound" {
    vcsim_env

    run govc namespace.service.version.info service2 9.9.9
    assert_failure
    assert_matches "404 Not Found"
}

@test "namespace.service.version.info" {
    vcsim_env

    v=$(govc namespace.service.version.info -json service2 1.1.0 | jq)

    assert_equal "ACTIVATED" $(echo $v |  jq -r .state)
    assert_equal "This is service 2 version 1.1.0" "$(echo $v |  jq -r .description)"
    assert_equal "CARVEL_APPS_YAML" $(echo $v |  jq -r .content_type)
    assert_equal "abc" $(echo $v | jq -r .content)
}

@test "namespace.create" {
    vcsim_env

    run govc namespace.create -cluster DC0_C0 test-namespace-1
    assert_success

    ns=$(govc namespace.info -json test-namespace-1 | jq)
    id=$(govc find -i -maxdepth 0 /DC0/host/DC0_C0 | cut -d: -f2)
    assert_equal "$id" $(echo $ns | jq -r '."cluster"')
    assert_equal "0" $(echo $ns | jq -r '."vm_service_spec"."content_libraries"' | jq length)

    run govc namespace.create -cluster DC0_C0 -library=lib1 -library=lib2 test-namespace-2
    assert_success

    ns=$(govc namespace.info -json test-namespace-2 | jq)
    assert_equal "2" $(echo $ns | jq -r '."vm_service_spec"."content_libraries"' | jq length)

    run govc namespace.create -cluster DC0_C0 -storage MyStoragePolicy test-namespace-2
    assert_failure # storage policy does not exist

    run govc storage.policy.create -category my_cat -tag my_tag MyStoragePolicy
    assert_success

    run govc namespace.create -cluster DC0_C0 -storage MyStoragePolicy test-namespace-2
    assert_success
}

@test "namespace.update" {
    vcsim_env

    govc namespace.create -cluster DC0_C0 test-namespace-1

    run govc namespace.update -library=lib1 -library=lib2 -vmclass=class1 test-namespace-1
    assert_success

    ns=$(govc namespace.info -json test-namespace-1 | jq)
    assert_equal "2" $(echo $ns | jq -r '."vm_service_spec"."content_libraries"' | jq length)
    assert_matches "lib[0-9]+" $(echo $ns | jq -r '."vm_service_spec"."content_libraries"[0]')
    assert_matches "lib[0-9]+" $(echo $ns | jq -r '."vm_service_spec"."content_libraries"[1]')
    assert_equal "1" $(echo $ns | jq -r '."vm_service_spec"."vm_classes"' | jq length)
    assert_equal "class1" $(echo $ns | jq -r '."vm_service_spec"."vm_classes"[0]')

    run govc namespace.update -library=lib3 test-namespace-1
    assert_success

    ns=$(govc namespace.info -json test-namespace-1 | jq)
    assert_equal "1" $(echo $ns | jq -r '."vm_service_spec"."content_libraries"' | jq length)
    assert_equal "lib3" $(echo $ns | jq -r '."vm_service_spec"."content_libraries"[0]')
    assert_equal "0" $(echo $ns | jq -r '."vm_service_spec"."vm_classes"' | jq length)

    run govc namespace.update -vmclass=class3 test-namespace-1
    assert_success

    ns=$(govc namespace.info -json test-namespace-1 | jq)
    assert_equal "0" $(echo $ns | jq -r '."vm_service_spec"."content_libraries"' | jq length)
    assert_equal "1" $(echo $ns | jq -r '."vm_service_spec"."vm_classes"' | jq length)
    assert_equal "class3" $(echo $ns | jq -r '."vm_service_spec"."vm_classes"[0]')

    run govc namespace.update -library=lib4 -vmclass=class4 test-namespace-1
    assert_success

    ns=$(govc namespace.info -json test-namespace-1 | jq)
    assert_equal "1" $(echo $ns | jq -r '."vm_service_spec"."content_libraries"' | jq length)
    assert_equal "lib4" $(echo $ns | jq -r '."vm_service_spec"."content_libraries"[0]')
    assert_equal "1" $(echo $ns | jq -r '."vm_service_spec"."vm_classes"' | jq length)
    assert_equal "class4" $(echo $ns | jq -r '."vm_service_spec"."vm_classes"[0]')

    run govc namespace.update -library=lib1 -library=lib2 -vmclass=class1 non-existing-namespace
    assert_failure
    assert_matches "404 Not Found"
}

@test "namespace.info" {
    vcsim_env

    govc namespace.create -cluster DC0_C0 test-namespace-1

    ns=$(govc namespace.info -json test-namespace-1 | jq)
    id=$(govc find -i -maxdepth 0 /DC0/host/DC0_C0 | cut -d: -f2)
    assert_equal "$id" $(echo $ns | jq -r '."cluster"')

    run govc namespace.info test-namespace-1
    assert_success

    run govc namespace.info non-existing-namespace
    assert_failure
    assert_matches "404 Not Found"
}

@test "namespace.ls" {
    vcsim_env

    run govc namespace.ls
    assert_success ""

    ls=$(govc namespace.ls -json)
    assert_equal "0" $(echo $ls | jq length)

    govc namespace.create -cluster DC0_C0 test-namespace-1
    ls=$(govc namespace.ls -json)
    assert_equal "1" $(echo $ls | jq length)
    assert_equal "test-namespace-1" $(echo $ls | jq -r '.[0]."namespace"')

    run govc namespace.ls
    assert_success test-namespace-1

    govc namespace.create -cluster DC0_C0 test-namespace-2
    ls=$(govc namespace.ls -json)
    assert_equal "2" $(echo $ls | jq length)
    id=$(govc find -i -maxdepth 0 /DC0/host/DC0_C0 | cut -d: -f2)
    assert_equal "$id" $(echo $ls | jq -r '.[0]."cluster"')
    assert_matches "test-namespace-[0-9]+" $(echo $ls | jq -r '.[0]."namespace"')
    assert_equal "$id" $(echo $ls | jq -r '.[1]."cluster"')
    assert_matches "test-namespace-[0-9]+" $(echo $ls | jq -r '.[1]."namespace"')
}

@test "namespace.rm" {
    vcsim_env

    run govc namespace.rm non-existing-namespace
    assert_failure
    assert_matches "404 Not Found"

    run govc namespace.create -cluster DC0_C0 test-namespace-1
    assert_success

    run govc namespace.rm test-namespace-1
    assert_success
}

@test "namespace.vmclass.create" {
    vcsim_env

    run govc namespace.vmclass.create -cpus=16 -memory=16000 test-class-1
    assert_success

    run govc namespace.vmclass.info -json test-class-1
    assert_success
    class="$output"

    run jq .cpu_count <<<"$class"
    assert_success 16

    run jq .memory_mb <<<"$class"
    assert_success 16000

    run govc namespace.vmclass.create -spec '{"numCPUs":6,"memoryMB":1024}' test-class-2
    assert_success

    run govc namespace.vmclass.info -json test-class-2
    assert_success
    class="$output"

    run jq .cpu_count <<<"$class"
    assert_success 6

    run jq .memory_mb <<<"$class"
    assert_success 1024

    run govc namespace.vmclass.info test-class-2
    assert_success

    spec=$(grep "ConfigSpec:" <<<"$output" | awk '{print $2}')

    run jq .numCPUs <<<"$spec"
    assert_success 6

    run jq .memoryMB <<<"$spec"
    assert_success 1024

}

@test "namespace.vmclass.update" {
    vcsim_env

    govc namespace.vmclass.create -cpus=16 -memory=16000 test-class-1

    govc namespace.vmclass.update -cpus=24 -memory=24000 test-class-1
    c=$(govc namespace.vmclass.info -json test-class-1 | jq)
    assert_equal "24" $(echo $c | jq -r '."cpu_count"')
    assert_equal "24000" $(echo $c | jq -r '."memory_mb"')
}

@test "namespace.vmclass.info" {
    vcsim_env

    run govc namespace.vmclass.info
    assert_failure

    run govc namespace.vmclass.create -cpus=16 -memory=16000 test-class-1
    assert_success

    run govc namespace.vmclass.info test-class-1
    assert_success

    c=$(govc namespace.vmclass.info -json test-class-1 | jq)
    assert_equal "16" $(echo $c | jq -r '."cpu_count"')
    assert_equal "16000" $(echo $c | jq -r '."memory_mb"')

    run govc namespace.vmclass.info non-existing-class
    assert_failure
    assert_matches "404 Not Found"
}

@test "namespace.vmclass.ls" {
    vcsim_env

    run govc namespace.vmclass.ls
    assert_success ""

    ls=$(govc namespace.vmclass.ls -json)
    assert_equal "0" $(echo $ls | jq length)

    govc namespace.vmclass.create -cpus=16 -memory=16000 test-class-1
    ls=$(govc namespace.vmclass.ls -json)
    assert_equal "1" $(echo $ls | jq length)

    govc namespace.vmclass.create -cpus=16 -memory=16000 test-class-2
    ls=$(govc namespace.vmclass.ls -json)
    assert_equal "2" $(echo $ls | jq length)
}

@test "namespace.vmclass.rm" {
    vcsim_env

    run govc namespace.vmclass.rm non-existing-class
    assert_failure
    assert_matches "404 Not Found"

    govc namespace.vmclass.create -cpus=16 -memory=16000 test-class-1

    run govc namespace.vmclass.rm test-class-1
    assert_success
}

@test "namespace.registervm" {
  vcsim_env

  vm=DC0_C0_RP0_VM0

  run govc namespace.create -cluster DC0_C0 test-namespace-1
  assert_success

  run govc namespace.registervm -vm $vm test-namespace-1
  assert_failure # missing resource.yaml

  run govc vm.change -vm $vm -e vmservice.virtualmachine.resource.yaml=b64
  assert_success

  run govc namespace.registervm -vm $vm test-namespace-1
  assert_success
}
