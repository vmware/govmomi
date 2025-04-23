#!/usr/bin/env bats

load test_helper

@test "vcsim rbvmomi" {
  if ! ruby -e "require 'rbvmomi'" ; then
    skip "requires rbvmomi"
  fi

  vcsim_env

  ruby ./vcsim_test.rb "$(govc env -x GOVC_URL_PORT)"
}

@test "vcsim powercli" {
  require_docker

  vcsim_env -l 0.0.0.0:0

  server=$(govc env -x GOVC_URL_HOST)
  port=$(govc env -x GOVC_URL_PORT)

  # docker run --rm projects.registry.vmware.com/pez/powerclicore@sha256:09b29f69c0653f871f6d569f7c4c03c952909f68a27e9792ef2f7c8653235668 /usr/bin/pwsh -f - <<EOF
  docker run --rm ghcr.io/embano1/powerclicore@sha256:09b29f69c0653f871f6d569f7c4c03c952909f68a27e9792ef2f7c8653235668 /usr/bin/pwsh -f - <<EOF
Set-PowerCLIConfiguration -InvalidCertificateAction Ignore -confirm:\$false | Out-Null
Connect-VIServer -Server $server -Port $port -User user -Password pass

Get-VM
Get-VIEvent
Get-VirtualNetwork
EOF
}

@test "vcsim examples" {
  vcsim_env

  # compile + run examples against vcsim
  for main in ../../examples/*/main.go ; do
    # TODO: #2476
    if [[ $main =~ alarm ]]; then
      continue
    fi
    run go run "$main" -insecure -url "$GOVC_URL"
    assert_success
  done
}

@test "vcsim about" {
  vcsim_env -dc 2 -cluster 3 -vm 0 -ds 0

  url="https://$(govc env GOVC_URL)"

  run curl -skf "$url/about"
  assert_matches "CurrentTime" # 1 param (without Context)
  assert_matches "TerminateSession" # 2 params (with Context)
  assert_matches "CnsAttachVolume" # method from namespace vsan
  assert_matches "PbmCreate" # method from namespace pbm

  run curl -skf "$url/debug/vars"
  assert_success

  model=$(curl -sfk "$url/debug/vars" | jq .vcsim.model)
  [ "$(jq .datacenter <<<"$model")" == "2" ]
  [ "$(jq .cluster <<<"$model")" == "6" ]
  [ "$(jq .machine <<<"$model")" == "0" ]
  [ "$(jq .datastore <<<"$model")" == "0" ]

  run govc about
  assert_success
  assert_matches "govmomi simulator"
}

@test "truncated WaitForUpdatesEx" {
  total=142 # 100 is the default MaxObjectUpdates
  vcsim_env -standalone-host 0 -vm $total

  url="https://$(govc env GOVC_URL)"

  model=$(curl -sfk "$url/debug/vars" | jq .vcsim.model)
  vms=$(jq .machine <<<"$model")
  assert_equal $total "$vms" # sanity check

  run govc object.collect -type m / name
  assert_success
  [ ${#lines[@]} -eq "$total" ]
}

@test "vcsim host placement" {
  vcsim_start -dc 0

  # https://github.com/vmware/govmomi/issues/1258
  id=$(new_id)
  govc datacenter.create DC0
  govc cluster.create comp
  govc cluster.add -cluster comp -hostname test.host.com -username user -password pass
  govc cluster.add -cluster comp -hostname test2.host.com -username user -password pass
  govc datastore.create -type local -name vol6 -path "$BATS_TMPDIR" test.host.com
  govc pool.create comp/Resources/testPool
  govc vm.create -c 1 -ds vol6 -g centos64Guest -pool testPool -m 4096 "$id"
  govc vm.destroy "$id"
}

@test "vcsim host config.port" {
  vcsim_start -dc 0
  url=$(govc env GOVC_URL)
  port=$(govc env -x GOVC_URL_PORT)
  vcsim_stop

  vcsim_start -l "$url" # reuse free port selection from above

  run govc object.collect -s -type h host/DC0_H0 summary.config.port
  assert_success "$port"
  ports=$(govc object.collect -s -type h / summary.config.port | uniq -u | wc -l)
  assert_equal 0 "$ports" # all host ports should be the same value

  vcsim_stop

  VCSIM_HOST_PORT_UNIQUE=true vcsim_start -l "$url"

  hosts=$(curl -sk "https://$url/debug/vars" | jq .vcsim.model.host)
  ports=$(govc object.collect -s -type h / summary.config.port | uniq -u | wc -l)
  assert_equal "$ports" "$hosts" # all host ports should be unique
  grep -v "$port" <<<"$ports" # host ports should not include vcsim port
}

@test "vcsim set vm properties" {
  vcsim_env

  vm=/DC0/vm/DC0_H0_VM0

  run govc object.collect $vm guest.ipAddress
  assert_success ""

  run govc vm.change -vm $vm -e SET.guest.ipAddress=10.0.0.1
  assert_success

  run govc object.collect -s $vm guest.ipAddress
  assert_success "10.0.0.1"

  run govc vm.ip $vm
  assert_success "10.0.0.1"

  run govc object.collect -s $vm summary.guest.ipAddress
  assert_success "10.0.0.1"

  netip=$(govc object.collect -json -o $vm guest.net | jq -r .guest.net[].ipAddress[0])
  [ "$netip" = "10.0.0.1" ]

  run govc vm.info -vm.ip 10.0.0.1
  assert_success

  run govc object.collect -s $vm guest.hostName
  assert_success ""

  run govc vm.change -vm $vm -e SET.guest.hostName=localhost.localdomain
  assert_success

  run govc object.collect -s $vm guest.hostName
  assert_success "localhost.localdomain"

  run govc object.collect -s $vm summary.guest.hostName
  assert_success "localhost.localdomain"

  run govc vm.info -vm.dns localhost.localdomain
  assert_success

  uuid=$(vcsim uuidgen)
  run govc vm.change -vm $vm -e SET.config.uuid="$uuid"
  assert_success

  run govc object.collect -s $vm config.uuid
  assert_success "$uuid"

  govc import.ovf -options - "$GOVC_IMAGES/$TTYLINUX_NAME.ovf" <<EOF
{
  "PropertyMapping": [
    {
      "Key": "SET.guest.ipAddress",
      "Value": "10.0.0.42"
    }
  ],
  "PowerOn": true,
  "WaitForIP": true
}
EOF

  run govc vm.ip "$TTYLINUX_NAME"
  assert_success "10.0.0.42"

  run govc vm.destroy "$TTYLINUX_NAME"
  assert_success

  govc import.ovf -options - "$GOVC_IMAGES/$TTYLINUX_NAME.ovf" <<EOF
{
  "PropertyMapping": [
    {
      "Key": "ip0",
      "Value": "10.0.0.43"
    }
  ],
  "PowerOn": true,
  "WaitForIP": true
}
EOF

  run govc vm.ip "$TTYLINUX_NAME"
  assert_success "10.0.0.43"
}

@test "vcsim vm.create" {
  vcsim_env

  # VM uuids are stable, based on path to .vmx
  run govc object.collect -s vm/DC0_H0_VM0 config.uuid config.instanceUuid
  assert_success "$(printf "63d40cc5-9cba-5cc1-884a-f7f19070ecea\nb170c191-7587-5f8e-9a15-08c6a0a11ca5")"

  dups=$(govc object.collect -s -type m / config.uuid | sort | uniq -d | wc -l)
  assert_equal 0 "$dups"

  run govc object.collect -s host/DC0_H0/DC0_H0 summary.hardware.uuid
  assert_success efc5827c-ee19-5d35-84ef-a77d6ea6ee4c

  dups=$(govc object.collect -s -type m / summary.hardware.uuid | sort | uniq -d | wc -l)
  assert_equal 0 "$dups"

  run govc vm.create foo.yakity
  assert_success

  run govc vm.create bar.yakity
  assert_success
}

@test "vcsim issue #1251" {
  vcsim_env

  govc object.collect -type ComputeResource -n 1 / name &
  pid=$!
  # give some time for the above to park in WaitForUpdates,
  # otherwise the rename may happen first and collect will hang as it missed the update
  sleep .2

  run govc object.rename /DC0/host/DC0_C0 DC0_C0b
  assert_success

  wait $pid

  govc object.collect -type ClusterComputeResource -n 1 / name &
  pid=$!
  sleep .2

  run govc object.rename /DC0/host/DC0_C0b DC0_C0
  assert_success

  wait $pid
}

@test "vcsim issue #3396" {
  vcsim_env

  govc pool.create /DC0/host/DC0_C0/Resources/foo
  govc pool.create /DC0/host/DC0_C0/Resources/foo/bar
  govc find -l /DC0/host/DC0_C0/Resources
  govc pool.destroy /DC0/host/DC0_C0/Resources/foo
  # prior to the fix, "bar"'s Parent was still "foo", causing the following to hang
  govc find -l /DC0/host/DC0_C0/Resources
}

@test "vcsim run container" {
  require_docker

  vcsim_env -autostart=false

  vm=DC0_H0_VM0
  name=$(docker_name $vm)

  if docker inspect "$name" ; then
    flunk "$vm container still exists"
  fi

  run govc vm.change -vm $vm -e RUN.container=nginx
  assert_success

  run govc vm.power -on $vm
  assert_success

  if ! docker inspect "$name" ; then
    flunk "$vm container does not exist"
  fi

  ip=$(docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' "$name")
  run govc collect -s vm/$vm guest.ipAddress
  assert_success "$ip"

  run govc collect -s vm/$vm summary.guest.ipAddress
  assert_success "$ip"

  netip=$(govc collect -json -o vm/$vm guest.net | jq -r .guest.net[].ipAddress[0])
  [ "$netip" = "$ip" ]

  run docker inspect -f '{{.Config.Hostname}}' "$name"
  assert_success
  hostname="$output"

  run docker inspect -f '{{.NetworkSettings.Gateway}}' "$name"
  assert_success
  gateway="$output"

  run govc collect -s -json vm/$vm guest.ipStack
  assert_success
  ipStack="$output"

  run jq -r .[].dnsConfig.hostName <<<"$ipStack"
  assert_success "$hostname"

  run govc collect -s vm/$vm guest.hostName
  assert_success "$hostname"

  run jq -r .[].ipRouteConfig.ipRoute[].gateway.ipAddress <<<"$ipStack"
  assert_success "$gateway"

  run govc vm.ip $vm # covers VirtualMachine.WaitForIP
  assert_success "$ip"

  run govc vm.ip -a $vm # covers VirtualMachine.WaitForNetIP
  assert_success "$ip"

  run govc vm.ip -n ethernet-0 $vm # covers VirtualMachine.WaitForNetIP
  assert_success "$ip"

  run govc vm.power -s $vm
  assert_success

  # wait for power state == off after guest shutdown above
  run govc collect -s vm/$vm -runtime.powerState poweredOff
  assert_success

  run docker inspect -f '{{.State.Status}}' "$name"
  assert_success "exited"

  run govc vm.power -on $vm
  assert_success

  run docker inspect -f '{{.State.Status}}' "$name"
  assert_success "running"

  run govc vm.destroy $vm
  assert_success

  if docker inspect "$name" ; then
    flunk "$vm container still exists"
  fi

  vm=DC0_H0_VM1
  name=$(docker_name $vm)

  # test json encoded args
  run govc vm.change -vm $vm -e RUN.container="[\"-v\", \"$PWD:/usr/share/nginx/html:ro\", \"nginx\"]"
  assert_success

  # test bash -c args parsing
  run govc vm.change -vm $vm -e RUN.container="-v '$PWD:/usr/share/nginx/html:ro' nginx"
  assert_success

  run govc vm.power -on $vm
  assert_success

  run docker inspect "$name"
  assert_success

  ip=$(govc collect -s vm/$vm guest.ipAddress)
  run docker run --rm curlimages/curl curl -f "http://$ip/vcsim.bats"
  assert_success

  # test suspend/resume
  run docker inspect -f '{{.State.Status}}' "$name"
  assert_success "running"

  run govc vm.power -suspend $vm
  assert_success

  run docker inspect -f '{{.State.Status}}' "$name"
  assert_success "paused"

  run govc vm.power -on $vm
  assert_success

  run docker inspect -f '{{.State.Status}}' "$name"
  assert_success "running"

  run docker volume inspect "$name--dmi"
  assert_success

  run govc vm.destroy $vm
  assert_success

  run docker volume inspect "$name"
  assert_failure

  vm=DC0_C0_RP0_VM0
  name=$(docker_name $vm)

  run govc vm.change -vm $vm -e RUN.container="busybox grep VMware- /sys/class/dmi/id/product_serial"
  assert_success

  run govc vm.power -on $vm
  assert_success

  run docker inspect -f '{{.State.ExitCode}}' "$name"
  assert_success "0"

  run govc vm.destroy $vm
  assert_success

  vm=DC0_C0_RP0_VM1
  name=$(docker_name $vm)

  run govc vm.change -vm $vm -e RUN.container="busybox sh -c 'sleep \$VMX_GUESTINFO_SLEEP'" -e guestinfo.sleep=500
  assert_success

  run govc vm.power -on $vm
  assert_success

  run docker inspect -f '{{.State.Status}}' "$name"
  assert_success "running"

  # stopping vcsim should remove the containers and volumes
  vcsim_stop

  run docker inspect "$name"
  assert_failure

  run docker volume inspect "$name"
  assert_failure
}

@test "vcsim listen" {
  vcsim_start -dc 0
  url=$(govc option.ls vcsim.server.url)
  [[ "$url" == *"https://127.0.0.1:"* ]]
  vcsim_stop

  # Symantec WSS Agent may block this :shrug:
  if [ -z "$(pidof wssa-ui_netext)" ] ; then
    vcsim_start -dc 0 -l 0.0.0.0:0
    url=$(govc option.ls vcsim.server.url)
    [[ "$url" != *"https://127.0.0.1:"* ]]
    [[ "$url" != *"https://[::]:"* ]]
    vcsim_stop
  fi
}

@test "vcsim vapi auth" {
  vcsim_env

  url=$(govc env GOVC_URL)

  run curl -fsk "https://$url/rest/com/vmware/cis/tagging/tag"
  [ "$status" -ne 0 ] # not authenticated

  run curl -fsk -X POST "https://$url/rest/com/vmware/cis/session"
  [ "$status" -ne 0 ] # no basic auth header

  run curl -fsk -X POST --user user: "https://$url/rest/com/vmware/cis/session"
  [ "$status" -ne 0 ] # no password

  run curl -fsk -X POST --user "$USER:pass" "https://$url/rest/com/vmware/cis/session"
  assert_success # login with user:pass

  id=$(jq -r .value <<<"$output")

  run curl -fsk "https://$url/rest/com/vmware/cis/session"
  [ "$status" -ne 0 ] # no header or cookie

  run curl -fsk "https://$url/rest/com/vmware/cis/session" -H "vmware-api-session-id:$id"
  assert_success # valid session header

  user=$(jq -r .value.user <<<"$output")
  assert_equal "$USER" "$user"
}

@test "vcsim auth" {
  vcsim_start

  run env GOVC_USERNAME=anybody GOVC_PASSWORD=anything govc ls -u "$(govc env GOVC_URL)"
  assert_success

  vcsim_stop

  vcsim_start -username nobody -password nothing

  run govc ls
  assert_success

  run govc tags.ls
  assert_success

  run env GOVC_USERNAME=nobody GOVC_PASSWORD=nothing govc ls -u "$(govc env GOVC_URL)"
  assert_success

  run env GOVC_USERNAME=nobody GOVC_PASSWORD=nothing govc tags.ls -u "$(govc env GOVC_URL)"
  assert_success

  run govc ls -u "user:pass@$(govc env GOVC_URL)"
  assert_failure

  run govc tags.ls -u "user:pass@$(govc env GOVC_URL)"
  assert_failure

  run env GOVC_USERNAME=user GOVC_PASSWORD=pass govc ls -u "$(govc env GOVC_URL)"
  assert_failure

  run env GOVC_USERNAME=user GOVC_PASSWORD=pass govc tags.ls -u "$(govc env GOVC_URL)"
  assert_failure

  run govc sso.user.create -p pass user
  assert_success

  run govc ls -u "user:pass@$(govc env GOVC_URL)"
  assert_success

  run govc tags.ls -u "user:pass@$(govc env GOVC_URL)"
  assert_success

  vcsim_stop

  dir=$($mktemp --tmpdir -d govc-test-XXXXX 2>/dev/null || $mktemp -d -t govc-test-XXXXX)
  echo nobody > "$dir/username"
  echo nothing > "$dir/password"

  vcsim_start -username "$dir/username" -password "$dir/password"

  run govc ls
  assert_success

  run env GOVC_USERNAME="$dir/username" GOVC_PASSWORD="$dir/password" govc ls -u "$(govc env GOVC_URL)"
  assert_success

  run govc ls -u "user:pass@$(govc env GOVC_URL)"
  assert_failure

  vcsim_stop

  rm -rf "$dir"
}

@test "vcsim ovftool" {
  if ! ovftool -h >/dev/null ; then
    skip "requires ovftool"
  fi

  vcsim_env

  url=$(govc env GOVC_URL)

  run ovftool --noSSLVerify --acceptAllEulas -ds=LocalDS_0 --network=DC0_DVPG0 "$GOVC_IMAGES/$TTYLINUX_NAME.ova" "vi://user:pass@$url/DC0/host/DC0_C0/DC0_C0_H1"
  assert_success

  run govc vm.destroy "$TTYLINUX_NAME"
  assert_success
}

@test "vcsim model load" {
  vcsim_start
  dir="$BATS_TMPDIR/$(new_id)"
  govc object.save -v -d "$dir"
  vcsim_stop

  vcsim_env -load "$dir"
  rm -rf "$dir"

  govc object.collect -s -type h / configManager.networkSystem | xargs -n1 -I% govc object.collect -s % dnsConfig

  objs=$(govc find / | wc -l)
  assert_equal 23 "$objs"

  run govc cluster.add -cluster DC0_C0 -hostname DC0_C0_H0-clone -username DC0_C0_H0 -password pass -noverify
  assert_success
  objs=$(govc find / | wc -l)
  assert_equal 24 "$objs"

  run govc host.add -hostname DC0_H0-clone -username DC0_H0 -password pass -noverify
  assert_success
  objs=$(govc find / | wc -l)
  assert_equal 27 "$objs" # ComputeResource + ResourcePool + HostSystem

  run govc host.portgroup.add -host DC0_H0 -vswitch vSwitch0 bridge
  assert_success # issue #2016

  # make sure we can write the save/load-ed Datastore
  run govc import.ova "$GOVC_IMAGES/$TTYLINUX_NAME.ova"
  assert_success
}

@test "vcsim trace file" {
  file="$BATS_TMPDIR/$(new_id).trace"

  vcsim_start -trace-file "$file"

  run govc ls
  assert_success

  vcsim_stop

  run ls -l "$file"
  assert_success

  rm "$file"
}

@test "vcsim jsonrpc" {
  vcsim_start

  url=$(govc env -x GOVC_URL) # w/o username:password@

  run govc session.login -u "$url" -X POST /api <<EOF
{
  "jsonrpc": "2.0",
  "method": "invoke",
  "params": {
    "serviceId": "com.vmware.cis.session",
    "operationId": "create",
    "ctx": {
      "securityCtx": {
        "schemeId": "com.vmware.vapi.std.security.user_pass",
        "userName": "$(govc env GOVC_USERNAME)",
        "password": "$(govc env GOVC_PASSWORD)"
      }
    }
  },
  "id": "0"
}
EOF
  assert_success

  session=$(jq -r .result.output.SECRET <<<"$output")
  [ -n "$session" ]

  run govc session.login -X POST /api <<EOF
{
  "jsonrpc": "2.0",
  "method": "invoke",
  "params": {
    "serviceId": "com.vmware.cis.session",
    "operationId": "delete"
  },
  "id": "123"
}
EOF
  assert_success

  id=$(jq -r .id <<<"$output")
  assert_equal "$id" "123"

  run govc session.login -r -X GET /rest/vcenter/certificate-authority/get-root
  assert_success

  run jq -r .value <<<"$output"
  assert_success

  run openssl x509 -text <<<"$output"
  assert_success

  run govc host.cert.csr -host DC0_H0
  assert_success
  csr="${output//$'\n'/\\n}"

  run govc session.login -r -X POST /rest/vcenter/certificate-authority/sign-cert <<EOF
{ "csr": "$csr", "duration": "30"}
EOF
  assert_success

  run jq -r .value <<<"$output"
  assert_success

  cert="$output"
  run openssl x509 -text <<<"$cert"
  assert_success
  assert_matches DNS:DC0_H0
}
