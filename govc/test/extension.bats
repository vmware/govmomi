#!/usr/bin/env bats

load test_helper

@test "extension" {
  vcsim_env

  run govc extension.info enoent
  assert_failure

  run govc extension.info
  assert_success

  id=$(new_id)

  # register extension
  run govc extension.register $id <<EOS
  {
    "Description": {
      "Label": "govc",
      "Summary": "Go interface to vCenter"
    },
    "Key": "${id}",
    "Company": "VMware, Inc.",
    "Version": "0.2.0"
  }
EOS
  assert_success

  # check info output is legit
  run govc extension.info $id
  assert_line "Name: $id"

  json=$(govc extension.info -json $id)
  label=$(jq -r .extensions[].description.label <<<"$json")
  assert_equal "govc" "$label"

  # change label and update extension
  json=$(jq -r '.extensions[] | .description.label = "novc"' <<<"$json")
  run govc extension.register -update $id <<<"$json"
  assert_success

  # check label changed in info output
  json=$(govc extension.info -json $id)
  label=$(jq -r .extensions[].description.label <<<"$json")
  assert_equal "novc" "$label"

  # set extension certificate to generated certificate
  run govc extension.setcert -cert-pem '+' $id
  assert_success

  # client certificate authentication is tested in session.bats

  # remove generated cert and key
  rm ${id}.{crt,key}

  run govc extension.info $(govc extension.info -json | jq -r .extensions[].key)
  assert_success

  run govc extension.unregister $id
  assert_success

  run govc extension.info $id
  assert_failure
}
