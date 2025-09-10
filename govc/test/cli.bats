#!/usr/bin/env bats

load test_helper

@test "about" {
  vcsim_env -api-version uE53DA

  run govc about
  assert_success
  assert_line "Vendor: VMware, Inc."

  run env GOVC_TLS_HANDSHAKE_TIMEOUT=10s govc about
  assert_success
  assert_line "Vendor: VMware, Inc."

  run env GOVC_TLS_HANDSHAKE_TIMEOUT=NOT_A_DURATION govc about
  assert_failure

  run govc about -json
  assert_success

  run govc about -json -l
  assert_success

  run govc about -dump
  assert_success

  run govc about -dump -l
  assert_success

  version=$(govc about -json -c | jq -r .client.version)
  assert_equal 9.1.0.0 "$version" # govc's default version

  version=$(govc about -json -c -vim-version "" | jq -r .client.version)
  assert_equal uE53DA "$version" # vcsim's service version

  version=$(env GOVC_VIM_VERSION=- govc about -json -c | jq -r .client.version)
  assert_equal uE53DA "$version" # vcsim's service version

  version=$(govc about -json -c -vim-version 6.8.2 | jq -r .client.version)
  assert_equal 6.8.2 "$version" # client specified version

  run govc about -trace
  assert_success

  run env GOVC_DEBUG_FORMAT=false govc about -trace
  assert_success

  run env GOVC_DEBUG_XML=enoent GOVC_DEBUG_JSON=enoent govc library.ls -trace
  assert_success
}

@test "about.cert" {
  vcsim_env -esx

  run govc about.cert
  assert_success

  run govc about.cert -json
  assert_success

  run govc about.cert -show
  assert_success

  # with -k=true we get sha256 thumbprint output and exit 0
  thumbprint=$(govc about.cert -k=true -thumbprint)

  # with -k=true we get thumbprint output and exit 60
  run govc about.cert -k=false -thumbprint
  if [ "$status" -ne 60 ]; then
    flunk $(printf "expected failed exit status=60, got status=%d" $status)
  fi
  assert_output "$thumbprint"

  run govc about -k=false
  assert_failure

  run govc about -k=false -tls-known-hosts <(echo "$thumbprint")
  assert_success

  run govc about -k=false -tls-known-hosts <(echo "nope nope")
  assert_failure

  # sha1 backwards compatibility
  host=$(awk '{print $1}'<<<"$thumbprint")
  sha1=$(govc about.cert -k=true -json | jq -r .thumbprintSHA1)
  run govc about -k=false -tls-known-hosts <(echo "$host $sha1")
  assert_success
}

@test "version" {
  vcsim_env -esx

  run govc version
  assert_success

  v=$(govc version | awk '{print $NF}')
  run govc version -require "$v"
  assert_success

  run govc version -require "not-a-version-string"
  assert_failure

  run govc version -require 100.0.0
  assert_failure
}

@test "login attempt without credentials" {
  vcsim_env -esx

  host=$(govc env -x GOVC_URL_HOST)
  port=$(govc env -x GOVC_URL_PORT)
  run govc about -u "enoent@$host:$port"
  assert_failure "govc: ServerFaultCode: Login failure"
}

@test "login attempt with GOVC_URL, GOVC_USERNAME, and GOVC_PASSWORD" {
  vcsim_env -esx

  govc_url_to_vars
  run govc about
  assert_success
}

@test "govc env" {
  output="$(govc env -x -u 'user:pass@enoent:99999?key=val#anchor')"
  assert grep -q GOVC_URL=enoent:99999 <<<${output}
  assert grep -q GOVC_USERNAME=user <<<${output}
  assert grep -q GOVC_PASSWORD=pass <<<${output}

  assert grep -q GOVC_URL_SCHEME=https <<<${output}
  assert grep -q GOVC_URL_HOST=enoent <<<${output}
  assert grep -q GOVC_URL_PORT=99999 <<<${output}
  assert grep -q GOVC_URL_PATH=/sdk <<<${output}
  assert grep -q GOVC_URL_QUERY=key=val <<<${output}
  assert grep -q GOVC_URL_FRAGMENT=anchor <<<${output}

  password="pa\$sword!ok"
  run govc env -u "user:${password}@enoent:99999" GOVC_PASSWORD
  assert_output "$password"
}

@test "govc help" {
  run govc
  assert_failure

  run govc -h
  assert_success
  assert_matches "Usage: govc <COMMAND> \[COMMON OPTIONS\] \[PATH\]..."

  run govc --help
  assert_success
  assert_matches "Usage: govc <COMMAND> \[COMMON OPTIONS\] \[PATH\]..."

  run govc -enoent
  assert_failure

  run govc vm.create
  assert_failure

  run govc vm.create -h
  assert_success

  run govc vm.create -enoent
  assert_failure

  run govc nope
  assert_failure
  assert_matches "Usage: govc <COMMAND> \[COMMON OPTIONS\] \[PATH\]..."

  run govc power
  assert_failure
  assert_matches "did you mean:"
}

@test "govc format error" {
  vcsim_env -host 1

  vm=DC0_H0_VM0

  run govc vm.power -json -on $vm
  assert_failure
  jq . <<<"$output"

  run govc vm.power -xml -on $vm
  assert_failure
  if type xmlstarlet ; then
    xmlstarlet fo <<<"$output"
  fi

  run govc vm.power -dump -on $vm
  assert_failure
  gofmt <<<"$output"

  run govc cluster.change -vsan-enabled DC0_C0
  assert_success

  run govc vm.create -ds vsanDatastore "$(new_id)"
  assert_failure
  assert_matches "requires 2 more usable fault domains"
}

@test "insecure cookies" {
  vcsim_start -tls=false

  run govc ls
  assert_success

  vcsim_stop

  VCSIM_SECURE_COOKIES=true vcsim_start -tls=false

  run govc ls
  assert_failure
  assert_matches NotAuthenticated # Go's cookiejar won't send Secure cookies if scheme != https

  vcsim_stop

  VCSIM_SECURE_COOKIES=true vcsim_start -tls=false

  run env GOVMOMI_INSECURE_COOKIES=true govc ls
  assert_success # soap.Client will set Cookie.Secure=false

  vcsim_stop
}

@test "govc verbose" {
  vcsim_env

  vm=DC0_H0_VM0

  run govc vm.power -verbose -off $vm
  assert_success

  run govc vm.power -verbose -off $vm
  assert_failure

  run govc vm.power -verbose -on $vm
  assert_success

  run govc vm.power -verbose -on $vm
  assert_failure

  run govc vm.info -verbose '*'
  assert_success

  run govc device.ls -vm DC0_H0_VM0 -verbose
  assert_success

  run govc device.info -vm DC0_H0_VM0 -verbose
  assert_success

  run govc vm.destroy -verbose $vm
  assert_success

  run govc host.info -verbose '*'
  assert_success

  run govc cluster.group.create -verbose -cluster DC0_C0 -name cgroup -vm DC0_C0_RP0_VM{0,1}
  assert_success

  run govc cluster.group.ls -cluster DC0_C0
  assert_success

  run govc cluster.create -verbose ClusterA
  assert_success

  run govc metric.ls -verbose /DC0/host/DC0_C0
  assert_success

  run govc metric.info -verbose /DC0/host/DC0_C0
  assert_success

  run govc metric.sample -verbose /DC0/host/DC0_C0 cpu.usage.average
  assert_success

  run govc session.login -verbose -issue # sts.Client
  assert_success

  run govc sso.service.ls -verbose # lookup.Client
  assert_success

  run govc storage.policy.ls -verbose # pbm.Client
  assert_success

  run govc volume.ls -verbose # cns.Client
  assert_success
}
