#!/usr/bin/env bats

load test_helper

@test "esxcli network vm list" {
  vcsim_env -esx

  # make sure there's at least 1 VM so we get a table header to count against
  vm=$(new_empty_vm)
  govc vm.power -on $vm

  nlines=$(govc host.esxcli network vm list | wc -l)

  vm=$(new_empty_vm)
  govc vm.power -on $vm

  xlines=$(govc host.esxcli network vm list | wc -l)

  # test that we see a new row
  [ $(($nlines + 1)) -eq $xlines ]

  run govc host.esxcli network vm list enoent
  assert_failure

  run govc host.esxcli network vm list
  assert_success
  assert_matches "VM Network"
  assert_matches VM0
  assert_matches VM1
}

@test "esxcli network ip connection list" {
  vcsim_env

  run govc host.esxcli -- network ip connection list -t tcp
  assert_success

  # test that we get the expected number of table columns
  nf=$(echo "${lines[1]}" | awk '{print NF}')
  [ $nf -eq 9 ]

  run govc host.esxcli -- network ip connection list -t enoent
  assert_failure
}

@test "esxcli system settings advanced list" {
  vcsim_env

  run govc host.esxcli -- system settings advanced list -o /Net/GuestIPHack
  assert_success
  assert_line "Path: /Net/GuestIPHack"

  run govc host.esxcli -- system settings advanced list -o /Net/ENOENT
  assert_failure
}

@test "esxcli software vib" {
  vcsim_env

  run govc host.esxcli software enoent
  assert_failure

  run govc host.esxcli software vib enoent
  assert_failure

  run govc host.esxcli software vib get -n esx-ui
  assert_success
}

@test "esxcli native types" {
  vcsim_env

  run govc host.esxcli hardware clock get # xsd:string
  assert_success
  assert_matches Z

  run govc host.esxcli iscsi software get # xsd:boolean
  assert_success false

  run govc host.esxcli system stats uptime get # xsd:long
  assert_success
}

@test "esxcli vm process" {
  vcsim_env -esx

  run govc host.esxcli vm process list
  assert_success
  assert_matches VM0
  assert_matches VM1
}

@test "esxcli firewall" {
  vcsim_env -esx

  run govc host.esxcli network firewall get
  assert_success
  assert_matches "DROP"
}

@test "esxcli model" {
  vcsim_env -esx

  export GOVC_SHOW_UNRELEASED=true

  run govc host.esxcli.model
  assert_success

  run govc host.esxcli.model network.vm network.ip.connection
  assert_success

  run govc host.esxcli.model -dump
  assert_success

  run govc host.esxcli.model -c
  assert_success

  run govc host.esxcli.model -i
  assert_success
}
