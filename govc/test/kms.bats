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

@test "kms.key" {
  vcsim_env

  run govc kms.add -N nkp
  assert_success

  host=$(govc env -x GOVC_URL_HOST)

  run govc kms.add -n my-server -a "$host" skp
  assert_success

  export GOVC_SHOW_UNRELEASED=true

  run govc kms.key.create nkp
  assert_failure # Cannot generate keys with native key provider

  run govc kms.key.create skp
  assert_success
  skey="$output"

  run govc kms.key.info -p skp "$skey"
  assert_success

  run govc kms.key.info -json "$skey"
  assert_success

  run jq .status[].keyAvailable <<<"$output"
  assert_success "false" # provider not specified

  run govc kms.key.info -json -p skp "$skey"
  assert_success

  run jq .status[].keyAvailable <<<"$output"
  assert_success "true"

  run govc kms.key.info -p nkp "$skey"
  assert_success

  run govc kms.key.info -json -p nkp "$skey"
  assert_success

  run jq .status[].keyAvailable <<<"$output"
  assert_success "false" # wrong provider for key
}
