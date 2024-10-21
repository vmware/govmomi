#!/usr/bin/env bats

load test_helper

@test "alarms" {
  vcsim_env

  run govc alarms
  assert_success

  run govc alarms -d
  assert_success

  vm=/DC0/vm/DC0_H0_VM0

  run govc alarms $vm
  assert_success

  run govc collect -s $vm triggeredAlarmState
  assert_success "" # empty

  run env GOVC_SHOW_UNRELEASED=true govc event.post -s info -i vcsim.vm.success $vm
  assert_success

  run govc alarms $vm
  assert_success
  [ ${#lines[@]} -eq 1 ] # header only, no alarm triggered

  run env GOVC_SHOW_UNRELEASED=true govc event.post -s warning -i vcsim.vm.failure $vm
  assert_success

  run govc alarms $vm
  assert_success
  assert_matches "Warning"

  run govc alarms -l $vm
  assert_success

  run govc alarms -json $vm
  assert_success
  alarms="$output"
  run jq -r .[].event.eventTypeId <<<"$alarms"
  assert_success "vcsim.vm.failure"

  run govc collect -json -s $vm triggeredAlarmState
  assert_success
  state="$output"
  run jq -r .[].overallStatus <<<"$state"
  assert_success "yellow"
  run jq -r .[].acknowledged <<<"$state"
  assert_success "false"

  run govc alarms -ack
  assert_success

  run govc collect -json -s $vm triggeredAlarmState
  assert_success
  state="$output"
  run jq -r .[].overallStatus <<<"$state"
  assert_success "yellow"
  run jq -r .[].acknowledged <<<"$state"
  assert_success "true"

  run env GOVC_SHOW_UNRELEASED=true govc event.post -s info -i vcsim.vm.success $vm
  assert_success

  run govc collect -s $vm triggeredAlarmState
  assert_success "" # empty

  run govc collect -s / triggeredAlarmState
  assert_success "" # empty
}

@test "alarm.info" {
  vcsim_env

  run govc alarm.info
  assert_success

  run govc alarm.info -n alarm.VmErrorAlarm
  assert_success
  assert_matches "Virtual machine error"

  run govc alarm.info -n invalid
  assert_success ""
}

@test "alarm -esx" {
  vcsim_env -esx

  run govc collect -s - content.alarmManager
  assert_success "" # empty, no AlarmManager

  run govc alarm.info
  assert_failure # not supported
}

@test "alarm.create" {
  vcsim_env

  export GOVC_SHOW_UNRELEASED=true

  run govc alarm.create -n "My Alarm" -green my.alarm.success -yellow my.alarm.failure
  assert_success
  id="$output"

  self=$(govc alarm.info -json -n "My Alarm" | jq -r .[].self.value)
  assert_equal "$id" "$self"

  run govc alarm.info -n "My Alarm"
  assert_success

  run govc alarm.create -n "My Alarm"
  assert_failure # DuplicateName

  run govc alarm.create -n "My Alarm" -r -d "This is my alarm description"
  assert_success

  run govc alarm.rm "$id"
  assert_success

  run govc alarm.rm "$id"
  assert_failure

  run govc alarm.info -n "My Alarm"
  assert_success ""
}
