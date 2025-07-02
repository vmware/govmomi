#!/bin/bash

# © Broadcom. All Rights Reserved.
# The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
# SPDX-License-Identifier: Apache-2.0

set -e

# This script shows for every request in a debug trace how long it took
# and the name of the request body.

function body-name {
  (
    xmllint --shell $1 <<EOS
    setns soapenv=http://schemas.xmlsoap.org/soap/envelope/
    xpath name(//soapenv:Body/*)
EOS
  )  | head -1 | sed 's/.*Object is a string : \(.*\)$/\1/'
}

if [ -n "$1" ]; then
  cd $1
fi

for req in $(find . -name '*.req.xml'); do
  base=$(basename $req .req.xml)
  session=$(echo $base | awk -F'-' "{printf \"%d\", \$1}")
  number=$(echo $base | awk -F'-' "{printf \"%d\", \$2}")
  client_log=$(dirname $req)/${session}-client.log
  took=$(awk "/ ${number} took / { print \$4 }" ${client_log})

  printf "%s %8s: %s\n" ${base} ${took} $(body-name $req)
done
