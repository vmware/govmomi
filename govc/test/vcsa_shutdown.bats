#!/usr/bin/env bats

load test_helper

@test "vcsa.shutdown.poweroff" {
  vcsim_env
  local output

  run govc vcsa.shutdown.poweroff -delay 50 "powering off for maintenance"
  assert_success

  status=$(govc vcsa.shutdown.get| grep -ow "poweroff"| awk '{print $1}')
  assert_equal $status 'poweroff'
   
  run govc vcsa.shutdown.poweroff -delay 50 "powering off for maintenance"
  assert_success

  status=$(govc vcsa.shutdown.get| grep -ow "poweroff"| awk '{print $1}')
  assert_equal $status 'poweroff'

  run govc vcsa.shutdown.poweroff
  assert_failure

}

@test "vcsa.shutdown.reboot" {
  vcsim_env
  local output

  run govc vcsa.shutdown.reboot -delay 50 "rebooting for maintenance"
  assert_success

  status=$(govc vcsa.shutdown.get| grep -ow "reboot"| awk '{print $1}')
  assert_equal $status 'reboot'

  run govc vcsa.shutdown.reboot -delay 50 "rebooting for maintenance"
  assert_success

  status=$(govc vcsa.shutdown.get| grep -ow "reboot"| awk '{print $1}')
  assert_equal $status 'reboot'

  run govc vcsa.shutdown.reboot
  assert_failure

}


@test "vcsa.shutdown.get" {
  vcsim_env
  local output

  run govc vcsa.shutdown.get
  assert_success

  run govc vcsa.shutdown.get -json
  assert_success

  run govc vcsa.shutdown.status -yaml
  assert_failure
}


@test "vcsa.shutdown.cancel" {
  vcsim_env
  local output

  run govc vcsa.shutdown.cancel
  assert_success

  run govc vcsa.shutdown.reboot -delay 50 "rebooting for maintenance"
  assert_success
  run govc vcsa.shutdown.cancel
  status=$(govc vcsa.shutdown.get| grep -i "Action"| awk '{print $2}')
  assert_equal $status ''
  status=$(govc vcsa.shutdown.get| grep -i "ShutDownTime"| awk '{print $2}')
  assert_equal $status ''
  status=$(govc vcsa.shutdown.get| grep -i "Reason"| awk '{print $2}')
  

  run govc vcsa.shutdown.cancel -reason "cancelling shutdown"
  assert_failure
}
