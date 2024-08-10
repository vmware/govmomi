#!/usr/bin/env bats

load test_helper

@test "session.ls" {
  vcsim_env

  run govc session.ls
  assert_success

  run govc session.ls -json
  assert_success

  # Test User-Agent
  govc session.ls | grep "$(govc version | tr ' ' /)"

  run govc session.ls -S
  assert_success

  run govc session.ls -u "$(govc env GOVC_USERNAME)@$(govc env GOVC_URL)" -S
  assert_failure # no password
}

@test "session.rm" {
  vcsim_env

  dir=$($mktemp --tmpdir -d govc-test-XXXXX 2>/dev/null || $mktemp -d -t govc-test-XXXXX)
  export GOVMOMI_HOME="$dir"
  export GOVC_PERSIST_SESSION=true

  run govc session.rm enoent
  assert_failure # NotFound

  # Can't remove the current session
  id=$(govc session.ls -json | jq -r .currentSession.key)
  run govc session.rm "$id"
  assert_failure

  thumbprint=$(govc about.cert -thumbprint)
  id=$(govc session.ls -json -k=false -tls-known-hosts <(echo "$thumbprint") | jq -r .currentSession.key)

  rm -rf "$dir"

  run govc session.rm "$id"
  assert_success
}

@test "session.persist" {
  vcsim_env

  dir=$($mktemp --tmpdir -d govc-test-XXXXX 2>/dev/null || $mktemp -d -t govc-test-XXXXX)
  export GOVMOMI_HOME="$dir"
  export GOVC_PERSIST_SESSION=true

  run govc role.ls
  assert_success

  run govc session.ls -r
  assert_success
  grep -v REST <<<"$output" # should not have a cached REST session

  run govc tags.ls
  assert_success # created a REST session

  run govc session.ls -r
  assert_success
  grep REST <<<"$output" # now we should have a cached REST session

  host=$(govc env GOVC_URL)
  user=$(govc env GOVC_USERNAME)
  run govc role.ls -u "$host" # url w/o user:pass
  assert_failure # NotAuthenticated

  run govc role.ls -u "$user@$host" # url w/o pass
  assert_success # authenticated via persisted session

  rm -rf "$dir"
}

@test "session.login" {
    vcsim_env

    # Remove username/password
    host=$(govc env GOVC_URL)

    # Validate auth is not required for service content
    run govc about -u "$host"
    assert_success

    # Auth is required here
    run govc ls -u "$host"
    assert_failure

    cookie=$(govc session.login -l)
    ticket=$(govc session.login -cookie "$cookie" -clone)

    run govc session.login -u "$host" -ticket "$ticket"
    assert_success

    cookie=$(govc session.login -r -l)
    run govc session.login -r -u "$host" -cookie "$cookie"
    assert_success

    user=$(govc env GOVC_USERNAME)
    dir=$($mktemp --tmpdir -d govc-test-XXXXX 2>/dev/null || $mktemp -d -t govc-test-XXXXX)
    export GOVMOMI_HOME="$dir"
    export GOVC_PERSIST_SESSION=true

    run govc session.login
    assert_success

    run govc role.ls -u "$user@$host" # url w/o pass
    assert_success # authenticated via persisted SOAP session

    run govc tags.ls -u "$user@$host" # url w/o pass
    assert_failure # no persisted REST session yet

    run govc session.login -r
    assert_success

    run govc tags.ls -u "$user@$host" # url w/o pass
    assert_success # authenticated via persisted REST session

    run govc session.logout -r
    assert_success

    run govc role.ls -u "$user@$host"
    assert_failure # logged out of persisted session

    run govc tags.ls -u "$user@$host"
    assert_failure # logged out of persisted session

    # ImpersonateUser

    run govc session.login -as vcsim
    assert_failure # user must have an existing session

    run govc session.login -u vcsim:pass@"$host"
    assert_success

    run govc session.login -as vcsim
    assert_success

    rm -rf "$dir"
}

@test "session.loginbytoken" {
  vcsim_env

  user=$(govc env GOVC_USERNAME)
  dir=$($mktemp --tmpdir -d govc-test-XXXXX 2>/dev/null || $mktemp -d -t govc-test-XXXXX)
  export GOVMOMI_HOME="$dir"
  export GOVC_PERSIST_SESSION=true

  # Remove username/password
  host=$(govc env GOVC_URL)
  # Token template, vcsim just checks Assertion.Subject.NameID
  token="<Assertion><Subject><NameID>%s</NameID></Subject></Assertion>"

  # shellcheck disable=2059
  run govc session.login -l -token "$(printf $token "")"
  assert_failure # empty NameID is a InvalidLogin fault

  # shellcheck disable=2059
  run govc session.login -l -token "$(printf $token root@localos)"
  assert_success # non-empty NameID is enough to login

  run govc role.ls -u "$user@$host" # url w/o pass
  assert_success # authenticated via persisted SOAP session

  run govc tags.ls -u "$user@$host" # url w/o pass
  assert_failure # no persisted REST session yet

  run govc session.login -r -token "$(printf $token root@localos)"
  assert_success

  run govc tags.ls -u "$user@$host" # url w/o pass
  assert_success # authenticated via persisted REST session

  id=$(new_id)
  run govc extension.setcert -cert-pem ++ "$id" # generate a cert for testing
  assert_success

  # Test with STS simulator issued token
  token="$(govc session.login -issue)"
  run govc session.login -cert "$id.crt" -key "$id.key" -l -token "$token"
  assert_success

  run govc session.login -cert "$id.crt" -key "$id.key" -l -renew
  assert_failure # missing -token

  run govc session.login -cert "$id.crt" -key "$id.key" -l -renew -lifetime 24h -token "$token"
  assert_success

  # remove generated cert and key
  rm "$id".{crt,key}
  rm -rf "$dir"
}

@test "session.loginextension" {
  vcsim_env -tunnel 0

  run govc session.login -extension com.vmware.vsan.health
  assert_failure # no certificate

  id=$(new_id)
  run govc extension.setcert -cert-pem ++ "$id" # generate a cert for testing
  assert_success

  # vcsim will login if any certificate is provided
  run govc session.login -extension com.vmware.vsan.health -cert "$id.crt" -key "$id.key"
  assert_success

  # remove generated cert and key
  rm "$id".{crt,key}
}

@test "session.curl" {
  vcsim_env

  run govc session.login /sdk/vimServiceVersions.xml
  assert_success

  run govc session.login /enoent
  assert_failure

  run govc session.login -r /rest/com/vmware/cis/session
  assert_success

  run govc session.login -r /enoent
  assert_failure

  cluster=$(govc find -i / -type c -name DC0_C0 | cut -d: -f2)

  run govc session.login -r -X POST "/rest/vcenter/cluster/modules" <<EOF
  {"spec": {"cluster": "$cluster"}}
EOF
  assert_success
  module=$(jq -r .value <<<"$output")

  members="/rest/vcenter/cluster/modules/vm/$module/members"
  vms=$(govc find -i /DC0/host/DC0_C0 -type m | cut -d: -f2 | jq --raw-input --slurp 'split("\n") | map(select(. != ""))')

  run govc session.login -r -X POST "$members?action=invalid" <<EOF
  {"vms": $vms}
EOF
  assert_failure # action=invalid

  run govc session.login -r -X POST "$members?action=add" <<EOF
  {"vms": $vms}
EOF
  assert_success
}
