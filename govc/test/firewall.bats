#!/usr/bin/env bats

load test_helper

@test "firewall.ruleset.find" {
  # Assumes sshServer ruleset is enabled
  run govc firewall.ruleset.find -direction inbound -port 22
  assert_success

  # Assumes sshClient ruleset is disabled
  run govc firewall.ruleset.find -direction outbound -port 22
  assert_failure

  run govc firewall.ruleset.find -direction outbound -port 22 -enabled=false
  assert_success

  # find disabled should include sshClient ruleset in output
  result=$(govc firewall.ruleset.find -direction outbound -port 22 -enabled=false | grep sshClient | wc -l)
  [ $result -eq 1 ]
}
