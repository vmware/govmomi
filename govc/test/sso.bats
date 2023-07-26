#!/usr/bin/env bats

load test_helper

@test "sso.service.ls" {
  vcsim_env

  sts=$(govc option.ls config.vpxd.sso.sts.uri | awk '{print $2}')

  # Remove credentials from URL, lookup service allows anonymous access
  GOVC_URL="$(govc env GOVC_URL)"

  run govc sso.service.ls
  assert_success

  run govc sso.service.ls -l
  assert_success

  run govc sso.service.ls -json
  assert_success

  run govc sso.service.ls -dump
  assert_success

  [ -z "$(govc sso.service.ls -t enoent)" ]

  run govc sso.service.ls -t cs.identity -P wsTrust -U
  assert_success "$sts"

  run govc sso.service.ls -t sso:sts -U
  assert_success "$sts"

  cert=$(govc about.cert -show | grep -v CERTIFICATE | tr -d '\n')
  trust=$(govc sso.service.ls -json -t sso:sts | jq -r .[].ServiceEndpoints[].SslTrust[0])
  assert_equal "$cert" "$trust"

  govc sso.service.ls -t cs.identity | grep com.vmware.cis | grep -v https:
  govc sso.service.ls -t cs.identity -l | grep https:
  govc sso.service.ls -p com.vmware.cis -t cs.identity -P wsTrust -T com.vmware.cis.cs.identity.sso -l | grep wsTrust
  govc sso.service.ls -P vmomi | grep vcenterserver | grep -v https:
  govc sso.service.ls -P vmomi -l | grep https:
}

@test "sso.idp.ls" {
  vcsim_env

  run govc sso.idp.ls -json
  assert_success

  run govc sso.idp.ls
  assert_success
  [ ${#lines[@]} -eq 4 ]
  assert_matches "System Domain"
  assert_matches "Local OS"
  assert_matches "ActiveDirectory"
}

@test "sso.user" {
  vcsim_env

  run govc sso.user.ls
  assert_success

  run govc sso.user.create -p password govc
  assert_success

  run govc sso.user.ls
  assert_success
  assert_matches govc

  run govc sso.user.ls -s
  assert_success ""

  run govc sso.user.create -p password govc
  assert_failure # duplicate name

  run govc sso.user.update -p newpassword govc
  assert_success

  run govc sso.user.rm govc
  assert_success

  run govc sso.user.rm govc
  assert_failure # does not exist

  run govc sso.user.create -C dummy-cert govc
  assert_success

  run govc sso.user.update -C new-cert govc
  assert_success

  run govc sso.user.ls -s
  assert_success
  assert_matches govc
}

@test "sso.group" {
  vcsim_env

  run govc sso.group.ls
  assert_success

  run govc sso.group.create bats
  assert_success

  run govc sso.group.create -d "govc CLI" govc
  assert_success

  run govc sso.group.ls
  assert_success
  assert_matches "govc CLI"

  run govc sso.group.update -d "govmomi/govc CLI" govc
  assert_success
  run govc sso.group.ls
  assert_success
  assert_matches "govmomi/govc CLI"

  run govc sso.group.update -a user govc
  assert_success

  govc sso.user.id | grep "groups=govc"

  run govc sso.group.update -r user govc
  assert_success
  govc sso.user.id | grep -v "groups=govc"

  run govc sso.group.update -g -a govc bats
  assert_success

  run govc sso.group.ls govc
  assert_success
  assert_matches bats

  run govc sso.group.rm govc
  assert_success

  run govc sso.group.rm govc
  assert_failure # does not exist
}
