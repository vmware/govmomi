#!/bin/bash

# © Broadcom. All Rights Reserved.
# The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
# SPDX-License-Identifier: Apache-2.0

set -e

ensure_rb_vmodl() {
  mkdir -p ./rbvmomi
  [ -f ./rbvmomi/vmodl.db ] || \
    curl -sSLo \
      ./rbvmomi/vmodl.db \
      https://github.com/vmware-archive/rbvmomi/raw/master/vmodl.db
}

generate() {
  dst="$1"
  wsdl="$2"
  modl="$3"

  pkgs=(types methods)
  if [ -n "$modl" ] ; then
    pkgs+=(mo)
  fi

  for p in "${pkgs[@]}"
  do
    mkdir -p "$dst/$p"
  done

  echo "generating $dst/..."

  bundle exec ruby gen_from_wsdl.rb "$dst" "$wsdl"
  if [ -n "$modl" ] ; then
    bundle exec ruby gen_from_vmodl.rb "$dst" "$wsdl" "$modl"
  fi

  for p in "${pkgs[@]}"
  do
    pushd "$dst/$p" >/dev/null
    goimports -w ./*.go
    go install
    popd >/dev/null
  done
}

update_vim_version() {
  sed -i'.bak' -e 's~^[[:blank:]]\{1,\}Version[[:blank:]]\{1,\}=.\{1,\}$~Version = '"\"${1}\""'~g' ../vim25/client.go
  rm -f ../vim25/client.go.bak
  go fmt ../vim25/client.go

  sed -i'.bak' -e 's/assert_equal .* ".*govc'"'"'s default version/assert_equal '"${1}"' "$version" # govc'"'"'s default version/' ../govc/test/cli.bats
  rm -f ../govc/test/cli.bats.bak
}

#
# Make sure the vmodl.db file exists.
#
ensure_rb_vmodl

#
# The VIM API version used by the vim25 client.
#
VIM_VERSION="${VIM_VERSION:-9.1.0.0}"

#
# Update the vim25 client's VIM version.
#
update_vim_version "${VIM_VERSION}"


#
# All types derive from vSphere 9.1.0.0, main..
#
export COPYRIGHT_DATE_RANGE="2014-2025"

#
# FORCE_BASE_INTERFACE_FOR_TYPES defines the types that we want to
# generate base interfaces for. The type names should be comma separated -
# e.g. "TypeA,TypeB,TypeC".
#
export FORCE_BASE_INTERFACE_FOR_TYPES="AgencyConfigInfo"

# ./sdk/ contains the contents of wsdl.zip from main.
generate "../vim25" "vim" "./rbvmomi/vmodl.db" # from github.com/vmware/rbvmomi@v3.0.0
generate "../pbm" "pbm"
generate "../vslm" "vslm"
generate "../sms" "sms"

# ./sdk/ contains the files eam-messagetypes.xsd and eam-types.xsd from
# eam-wsdl.zip, from main.
#
# Please note the EAM files are also available at the following, public URL --
# http://bit.ly/eam-sdk, therefore the WSDL resource for EAM are in fact
# public. A specific build was obtained in order to match the same build as
# used for the file from above, wsdl.zip.
COPYRIGHT_DATE_RANGE="2021-2025" generate "../eam" "eam"

# originally generated, then manually pruned as there are several vim25 types
# that are duplicated.
#
# generate "../lookup" "lookup" # lookup.wsdl from build 4571810
# originally generated, then manually pruned.
# generate "../ssoadmin" "ssoadmin" # ssoadmin.wsdl from PSC
