#!/usr/bin/env bats

load test_helper

@test "datastore.ls" {
  file=$(mktemp --tmpdir govc-test-XXXXX)
  name=$(basename ${file})
  echo "Hello world!" > ${file}

  run govc datastore.upload "${file}" "${name}"
  assert_success

  rm -f "${file}"

  # Single argument
  run govc datastore.ls "${name}"
  assert_success
  [ ${#lines[@]} -eq 1 ]

  # Multiple arguments
  run govc datastore.ls "${name}" "${name}"
  assert_success
  [ ${#lines[@]} -eq 2 ]

  # Pattern argument
  run govc datastore.ls "./govc-test-*"
  assert_success
  [ ${#lines[@]} -ge 1 ]

  # Long listing
  run govc datastore.ls -l "./govc-test-*"
  assert_success
  assert_equal "13" $(awk '{ print $1 }' <<<${output})
}
