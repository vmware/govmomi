#!/usr/bin/env bats

load test_helper

@test "tags.category" {
  vcsim_env

  category_name=$(new_id)
  des="new_des"

  run govc tags.category.create -d ${des} -m ${category_name}
  assert_success

  category_id=${output}
  run govc tags.category.ls

  ls_result=$(govc tags.category.ls | grep ${category_name} | wc -l)
  [ ${ls_result} -eq 1 ]

  id_res=$(govc tags.category.ls -json | jq -r '.[].CategoryID')
  assert_matches ${id_res} ${category_id}

  run govc tags.category.info ${category_name}
  assert_success

  update_des="update_des"
  update_n="update_d_n"

  run govc tags.category.update -n ${update_n} -d ${update_des} ${category_id}
  assert_success

  info_des=$(govc tags.category.info ${update_n} | grep ${update_des} | wc -l)
  [ ${info_des} -eq 1 ]
}

@test "tags" {
  vcsim_env

  run govc tags.category.create -d "desc" -m "$(new_id)"
  assert_success

  category="$output"
  test_name="test_name"
  des_tag="update_des_tag"

  run govc tags.create -d "desc" ${test_name} ${category}
  assert_success

  tag_id=${output}

  ls_result=$(govc tags.ls | grep ${test_name} | wc -l)
  [ ${ls_result} -eq 1 ]

  id_res=$(govc tags.ls -json | jq -r '.[].TagID')
  assert_matches ${id_res} ${tag_id}

  update_tag_name="update_name"
  run govc tags.update -d ${des_tag} -n ${update_tag_name} ${tag_id}
  assert_success

  des_result=$(govc tags.info ${update_tag_name} ${category} | grep ${des_tag} | wc -l)
  [ ${des_result} -eq 1 ]

  des_result=$(govc tags.info -i ${tag_id} | grep ${des_tag} | wc -l)
  [ ${des_result} -eq 1 ]
}

@test "tags.association" {
  vcsim_env

  run govc tags.category.create -d "desc" -m "$(new_id)"
  assert_success
  category="$output"

  run govc tags.create -d "desc" "$(new_id)" "$category"
  assert_success
  tag=$output

  tag_name=$(govc tags.ls -json | jq -r ".[] | select(.TagID == \"$tag\") | .Name")
  run govc find . -type h
  object=${lines[0]}

  run govc tags.attach ${tag} ${object}
  assert_success

  result=$(govc tags.association.ls ${object})
  assert_matches ${result} ${tag_name}

  result=$(govc tags.association.ls -json ${object} | jq -r '.[].Name')
  assert_matches ${result} ${tag_name}

  run govc tags.rm ${tag}
  assert_failure

  run govc tags.detach ${tag} ${object}
  assert_success

  run govc tags.rm ${tag}
  assert_success

  run govc tags.category.rm ${category}
  assert_success
}
