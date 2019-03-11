#!/usr/bin/env bats

load test_helper

@test "library" {
  vcsim_env

  run govc library.create my-content
  assert_success
  id="$output"

  jid=$(govc library.ls -json /my-content | jq -r .[].id)
  assert_equal "$id" "$jid"

  # run govc library.ls enoent
  # assert_failure # TODO: currently exit 0's

  run govc library.ls my-content
  assert_success "/my-content"

  run govc library.info /my-content
  assert_success
  assert_matches "$id"

  run govc library.rm "$id"
  assert_success
}

@test "library.item" {
  vcsim_env

  # run govc library.ls enoent
  # assert_failure # TODO: currently exit 0's

  # run govc library.info enoent
  # assert_failure # TODO: currently exit 0's

  run govc library.item.create enoent my-item
  assert_failure

  run govc library.create my-content
  assert_success

  run govc library.ls /my-content/*
  assert_success ""

  run govc library.item.create my-content my-item
  assert_success
  id="$output"

  run govc library.ls /my-content/my-item
  assert_success /my-content/my-item

  run govc library.info /my-content/my-item
  assert_success
  assert_matches "$id"

  run govc library.info /my-content/*
  assert_success
  assert_matches "$id"

  run govc library.item.upload my-content library.bats # any file will do
  assert_success

  id=$(govc library.ls -json "/my-content/library.bats" | jq -r .[].id)

  run govc library.info "/my-content/library.bats"
  assert_success
  assert_matches "$id"

  run govc library.item.rm my-content library.bats
  assert_success
}

@test "library.deploy" {
  vcsim_env

  run govc library.create my-content
  assert_success

  run govc library.ova my-content "$GOVC_IMAGES/${TTYLINUX_NAME}.ova"
  assert_success
}
