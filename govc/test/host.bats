#!/usr/bin/env bats
load test_helper

@test "host info esx" {
  vcsim_env

  run govc host.info
  assert_success
  grep -q Manufacturer: <<<"$output"

  run govc host.info -host enoent
  assert_failure "govc: host 'enoent' not found"

  for opt in dns ip ipath uuid
  do
    run govc host.info "-host.$opt" enoent
    assert_failure "govc: no such host"
  done

  # avoid hardcoding the esxbox hostname
  local name=$(govc ls -t HostSystem '/*/host/*' | head -1)

  run govc host.info -host "$name"
  assert_success
  grep -q Manufacturer: <<<"$output"

  run govc host.info -host ${name##*/}
  assert_success
  grep -q Manufacturer: <<<"$output"

  run govc host.info -host.ipath "$name"
  assert_success

  run govc host.info -host.dns localhost
  assert_success

  uuid=$(govc host.info -json | jq -r .hostSystems[].hardware.systemInfo.uuid)

  run govc host.info -host.uuid "$uuid"
  assert_success

  run govc host.info "*"
  assert_success
}

@test "host info vc" {
  vcsim_env

  run govc host.info
  assert_success
  grep -q Manufacturer: <<<"$output"

  run govc host.info -host enoent
  assert_failure "govc: host 'enoent' not found"

  for opt in ipath uuid # dns ip # TODO: SearchIndex:SearchIndex does not implement: FindByDnsName
  do
    run govc host.info "-host.$opt" enoent
    assert_failure "govc: no such host"
  done

  local name=$GOVC_HOST

  unset GOVC_HOST
  run govc host.info
  assert_failure "govc: default host resolves to multiple instances, please specify"

  run govc host.info -host "$name"
  assert_success
  grep -q Manufacturer: <<<"$output"

  run govc host.info -host.ipath "$name"
  assert_success

  run govc host.info -host.dns $(basename "$name")
  assert_failure # TODO: SearchIndex:SearchIndex does not implement: FindByDnsName

  uuid=$(govc host.info -host "$name" -json | jq -r .hostSystems[].summary.hardware.uuid)
  run govc host.info -host.uuid "$uuid"
  assert_success

  # summary.host should have a reference to the generated moid, not the template esx.HostSystem.Self (ha-host)
  govc object.collect -s -type h / summary.host | grep -v ha-host
}

@test "host maintenance vc" {
  vcsim_env

  run govc host.info
  assert_success
  grep -q -v Maintenance <<<"$output"

  run govc host.maintenance.enter "$GOVC_HOST"
  assert_success

  run govc host.info
  assert_success
  grep -q Maintenance <<<"$output"

  run govc host.maintenance.exit "$GOVC_HOST"
  assert_success

  run govc host.info
  assert_success
  grep -q -v Maintenance <<<"$output"
}

@test "host.vnic.info" {
  vcsim_env

  run govc host.vnic.info
  assert_success

  govc host.vnic.info -json | jq .
}

@test "host.vnic.hint" {
  vcsim_env

  run govc host.vnic.hint -xml -host DC0_C0_H0
  assert_success

  run govc host.disconnect DC0_C0_H0
  assert_success

  run govc host.vnic.hint -xml -host DC0_C0_H0
  assert_failure
  assert_matches HostNotConnected "$output"
}

@test "host.vswitch.info" {
  vcsim_env -esx

  run govc host.vswitch.info
  assert_success

  run govc host.vswitch.info -json
  assert_success
}

@test "host.portgroup.info" {
  vcsim_env -esx

  run govc host.portgroup.info
  assert_success

  run govc host.portgroup.info -json
  assert_success
}

@test "host.storage.info" {
    vcsim_env

    run govc host.storage.info
    assert_success

    run govc host.storage.info -rescan -refresh -rescan-vmfs
    assert_success

    run govc host.storage.info -t hba
    assert_success

    names=$(govc host.storage.info -json | jq -r .storageDeviceInfo.scsiLun[].alternateName[].data)
    # given data is hex encoded []byte and:
    #   [0] == encoding
    #   [1] == type
    #   [2] == ?
    #   [3] == length
    # validate name is at least 2 char x 4
    for name in $names; do
      [ "${#name}" -ge 8 ]
    done
}

@test "host.options" {
  vcsim_env -esx
  govc host.option.ls
  run govc host.option.ls Config.HostAgent.log.level
  assert_success

  run govc host.option.ls Config.HostAgent.log.
  assert_success

  run govc host.option.ls -json Config.HostAgent.log.
  assert_success

  run govc host.option.ls Config.HostAgent.plugins.solo.ENOENT
  assert_failure
}

@test "host.service" {
  esx_env

  run govc host.service.ls
  assert_success

  run govc host.service.ls -json
  assert_success

  run govc host.service status TSM-SSH
  assert_success
}

@test "host.cert.info" {
  vcsim_env -esx

  run govc host.cert.info
  assert_success

  run govc host.cert.info -json
  assert_success

  expires=$(govc host.cert.info -json | jq -r .notAfter)
  about_expires=$(govc about.cert -json | jq -r .notAfter)
  assert_equal "$expires" "$about_expires"

  run govc host.cert.info -show
  assert_success

  run openssl x509 -text <<<"$output"
  assert_success
}

@test "host.cert.csr" {
  vcsim_env -esx

  #   Requested Extensions:
  #       X509v3 Subject Alternative Name:
  #       IP Address:...
  result=$(govc host.cert.csr -ip | openssl req -text -noout)
  assert_matches "IP Address:" "$result"
  ! assert_matches "DNS:" "$result"

  #   Requested Extensions:
  #       X509v3 Subject Alternative Name:
  #       DNS:...
  result=$(govc host.cert.csr | openssl req -text -noout)
  ! assert_matches "IP Address:" "$result"
  assert_matches "DNS:" "$result"
}

@test "host.cert.import" {
  vcsim_env -esx

  expires=$(govc host.cert.info -json | jq -r .notAfter)

  govc host.cert.csr -ip | ./host_cert_sign.sh | govc host.cert.import
  expires2=$(govc host.cert.info -json | jq -r .notAfter)

  # cert expiration should have changed
  [ "$expires" != "$expires2" ]

  # verify hostd is using the new cert too
  expires=$(govc host.cert.info -json | jq -r .notAfter)
  assert_equal "$expires" "$expires2"
}

@test "host.date.info" {
  esx_env

  run govc host.date.info
  assert_success

  run govc host.date.info -json
  assert_success
}

@test "host.disconnect and host.reconnect" {
  vcsim_env

  run govc host.info
  assert_success
  status=$(govc host.info| grep -i "State"| awk '{print $2}')
  assert_equal 'connected' $status

  run govc host.disconnect "$GOVC_HOST"
  assert_success

  run govc host.info
  assert_success
  status=$(govc host.info| grep -i "State"| awk '{print $2}')
  assert_equal 'disconnected' $status

  run govc host.reconnect "$GOVC_HOST"
  assert_success

  run govc host.info
  assert_success
  status=$(govc host.info| grep -i "State"| awk '{print $2}')
  assert_equal 'connected' $status
}

@test "host.tpm" {
  vcsim_env

  run govc host.tpm.info
  assert_success

  run govc host.tpm.report
  assert_success
}
