#!/usr/bin/env bats

load test_helper

#testing tags.category.* commands
@test "tags.category.*" {
    esx_env
    export GOVC_DATASTORE=vsanDatastore

    category_name=$(new_id)
    des="new_des"

    run govc tags.category.create -d ${des} -m ${category_name}
    assert_success

    category_id=${output}
    run govc tags.category.ls
    ls_result=$(govc tags.category.ls | grep ${category_id} | wc -l)
     [ ${ls_result} -eq 1 ]

    run govc tags.category.get ${category_name}
    assert_success

    update_des="update_des"
    update_n="update_d_n"

    run govc tags.category.update ${category_id}
    assert_success

    run govc tags.category.update -n ${update_n} -d ${update_des} ${category_id}
    assert_success

    get_result=$(govc tags.category.get ${update_n})
    assert_matches ${category_id} ${get_result}

}

#testing tags.* commands
@test "tags.*" {
    esx_env
    export GOVC_DATASTORE=vsanDatastore
    
    category=$(govc tags.category.ls | awk '{printf}')
    test_name="test_name"
    des_tag="update_des_tag"

    run govc tags.create ${test_name} ${category}
    assert_success

    tag_id=${output}

    ls_result=$(govc tags.ls | grep ${tag_id} | wc -l)
    [ ${ls_result} -eq 1 ]


    ls_result=$(govc tags.ls -i ${category} ${test_name} | grep ${tag_id} | wc -l)
    [ ${ls_result} -eq 1 ]

    update_tag_name="update_name"
    run govc tags.update -d ${des_tag} -n ${update_tag_name} ${tag_id}
    assert_success

    name_result=$(govc tags.info ${tag_id} | awk 'NR==1{printf $2}')
    assert_matches ${name_result} ${update_tag_name}

    des_result=$(govc tags.info -n ${update_tag_name} ${category} | awk 'NR==3{printf $2}')
    assert_matches ${des_result} ${des_tag}

}

#testing tags.association.* commands
@test "tags.association.*" {
    esx_env
    export GOVC_DATASTORE=vsanDatastore
    
    category=$(govc tags.category.ls | awk '{printf}')
    tag=$(govc tags.ls | awk '{printf}')
    object=$(govc find . -type h | awk 'NR==1{printf}')

    run govc tags.attach ${tag} ${object}
    assert_success

    run govc tags.association.ls ${object}
    result=${output}
    assert_matches ${result} ${tag}

    run govc tags.rm ${tag}
    assert_failure

    run govc tags.detach ${tag} ${object}
    assert_success

    run govc tags.rm ${tag}

    ls_result=$(govc tags.ls | grep ${tag} | wc -l)
    [ ${ls_result} -eq 0 ]

    run govc tags.category.rm ${category}
    ls_result=$(govc tags.category.ls | grep ${category_id} | wc -l)
    [ ${ls_result} -eq 0 ]

}
