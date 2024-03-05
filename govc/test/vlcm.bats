#!/usr/bin/env bats

load test_helper

location="http://127.0.0.1:9999/files/custom_driver.zip"

@test "vlcm.depot.offline.ls" {
  vcsim_env

  run govc vlcm.depot.offline.ls
  assert_success
}

@test "vlcm.depot.offline.create" {
  vcsim_env

  run govc vlcm.depot.offline.create -l=$location -source-type=PULL
  assert_success

  depot=$(govc vlcm.depot.offline.ls | jq '."depot-1"')
  l=$(echo $depot | jq -r '.location')
  assert_equal $location $l
  srcType=$(echo $depot | jq -r '."source_type"')
  assert_equal "PULL" $srcType
}

@test "vlcm.depot.offline.info" {
  vcsim_env

  run govc vlcm.depot.offline.create -l=$location -source-type=PULL
  assert_success ""

  depot=$(govc vlcm.depot.offline.info -depot-id=depot-1)
  comp=$(echo $depot | jq '."metadata_bundles"."dummy-content"[0]."independent_components"."dummy-component"')
  name=$(echo $comp | jq -r '."display_name"')
  assert_equal "DummyComponent" $name
  version=$(echo $comp | jq -r '.versions[0].version')
  displayVersion=$(echo $comp | jq -r '.versions[0]."display_version"')
  assert_equal "1.0.0" $version
  assert_equal "1.0.0" $displayVersion
}

@test "vlcm.depot.offline.rm" {
  vcsim_env

  run govc vlcm.depot.offline.create -l=$location -source-type=PULL
  assert_success

  run govc vlcm.depot.offline.rm -depot-id=depot-1
  assert_success

  run govc vlcm.depot.offline.info -depot-id=depot-1
  assert_failure
  assert_matches "404 Not Found"
}

@test "cluster.draft.ls" {
  vcsim_env

  run govc cluster.draft.ls -cluster-id=domain-c21
  assert_success
}

@test "cluster.draft.create" {
  vcsim_env

  run govc cluster.draft.create -cluster-id=domain-c21
  assert_success

  draft=$(govc cluster.draft.ls -cluster-id=domain-c21 | jq '."1"')
  assert_equal "" $(echo $draft | jq -r '.owner')
}

@test "cluster.draft.info" {
  vcsim_env

  run govc cluster.draft.create -cluster-id=domain-c21
  assert_success

  run govc cluster.draft.info -cluster-id=domain-c21 -draft-id=1
  assert_success
}

@test "cluster.draft.rm" {
  vcsim_env

  run govc cluster.draft.create -cluster-id=domain-c21
  assert_success

  run govc cluster.draft.rm -cluster-id=domain-c21 -draft-id=1

  run govc cluster.draft.info -cluster-id=domain-c21 -draft-id=1
  assert_failure
  assert_matches "404 Not Found"
}

@test "cluster.draft.commit" {
  vcsim_env

  run govc cluster.draft.create -cluster-id=domain-c21
  assert_success

  run govc cluster.draft.commit -cluster-id=domain-c21 -draft-id=1

  run govc cluster.draft.info -cluster-id=domain-c21 -draft-id=1
  assert_failure
  assert_matches "404 Not Found"
}

@test "cluster.draft.component.ls" {
  vcsim_env

  run govc cluster.draft.create -cluster-id=domain-c21
  assert_success

  run govc cluster.draft.component.ls -cluster-id=domain-c21 -draft-id=1
  assert_success
}

@test "cluster.draft.component.add" {
  vcsim_env

  run govc cluster.draft.create -cluster-id=domain-c21
  assert_success

  run govc cluster.draft.component.add -cluster-id=domain-c21 -draft-id=1 -component-id=comp-id -component-version=1.2.3.4
  assert_success

  comp=$(govc cluster.draft.component.ls -cluster-id=domain-c21 -draft-id=1 | jq '."comp-id"')
  assert_equal "1.2.3.4" $(echo $comp | jq -r '.version')
  assert_equal "DummyComponent" $(echo $comp | jq -r '.details.display_name')
}

@test "cluster.draft.component.info" {
  vcsim_env

  run govc cluster.draft.create -cluster-id=domain-c21
  assert_success

  run govc cluster.draft.component.add -cluster-id=domain-c21 -draft-id=1 -component-id=comp-id -component-version=1.2.3.4
  assert_success

  comp=$(govc cluster.draft.component.info -cluster-id=domain-c21 -draft-id=1 -component-id=comp-id)
  assert_equal "1.2.3.4" $(echo $comp | jq -r '.version')
  assert_equal "DummyComponent" $(echo $comp | jq -r '.details.display_name')
  run govc cluster.draft.component.info -cluster-id=domain-c21 -draft-id=1 -component-id=invalid-id
  assert_failure
  assert_matches "404 Not Found"
}

@test "cluster.draft.component.rm" {
  vcsim_env

  run govc cluster.draft.create -cluster-id=domain-c21
  assert_success

  run govc cluster.draft.component.add -cluster-id=domain-c21 -draft-id=1 -component-id=comp-id -component-version=1.2.3.4
  assert_success

  run govc cluster.draft.component.rm -cluster-id=domain-c21 -draft-id=1 -component-id=comp-id

  run govc cluster.draft.component.info -cluster-id=domain-c21 -draft-id=1 -component-id=comp-id
  assert_failure
  assert_matches "404 Not Found"
}

@test "cluster.draft.baseimage.set" {
  vcsim_env

  run govc cluster.draft.create -cluster-id=domain-c21
  assert_success

  run govc cluster.draft.baseimage.set -cluster-id=domain-c21 -draft-id=1 -version=1.0.0
  assert_success

  baseimg=$(govc cluster.draft.baseimage.info -cluster-id=domain-c21 -draft-id=1)

  assert_equal "1.0.0" $(echo $baseimg | jq -r '.version')
}

@test "cluster.vlcm.enable" {
  vcsim_env

  res=$(govc cluster.vlcm.info -cluster-id=domain-c21)
  assert_equal "false" $(echo $res | jq -r '.enabled')

  run govc cluster.vlcm.enable -cluster-id=domain-c21

  res=$(govc cluster.vlcm.info -cluster-id=domain-c21)
  assert_equal "true" $(echo $res | jq -r '.enabled')
}