#!/usr/bin/env bats

load test_helper

@test "kms standard" {
  vcsim_env

  run govc kms.ls
  assert_success

  run govc kms.ls -json
  assert_success

  run govc kms.ls enoent
  assert_failure

  host=$(govc env -x GOVC_URL_HOST)

  run govc kms.add vcsim-kp
  assert_failure # InvalidProperty server.info.name

  run govc kms.add -n my-server -a "$host" vcsim-kp
  assert_success

  run govc kms.add -n my-server vcsim-kp
  assert_failure # already registered

  run govc kms.ls
  assert_success
  assert_matches vcsim-kp

  run govc kms.ls vcsim-kp
  assert_success

  run govc kms.default
  assert_failure

  run govc kms.default vcsim-kp
  assert_success

  run govc kms.default -
  assert_success

  run govc kms.rm -s my-server vcsim-kp
  assert_success

  run govc kms.add -n my-server -a "$host" vcsim-kp
  assert_success

  run govc kms.export vcsim-kp
  assert_failure # export is only supported for native
  assert_matches "400 Bad Request"

  run govc session.login -r -X DELETE "/api/vcenter/crypto-manager/kms/providers/vcsim-kp"
  assert_failure # vapi can only delete native providers
  assert_matches "400 Bad Request"

  run govc kms.rm vcsim-kp
  assert_success

  run govc kms.rm vcsim-kp
  assert_failure # does not exist
}

@test "kms native" {
  vcsim_env

  run govc kms.add -N nkp
  assert_success

  run govc kms.ls nkp
  assert_success

  run govc kms.export -f /dev/null nkp
  assert_success

  run govc kms.default nkp
  assert_success

  run govc kms.rm nkp
  assert_success
}
