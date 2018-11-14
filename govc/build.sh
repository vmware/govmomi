#!/bin/bash -e

git_version=$(git describe --dirty)
 if [[ $git_version == *-dirty ]] ; then
  echo 'Working tree is dirty.'
  echo 'NOTE: This script is meant for building govc releases via release.sh'
  echo 'To build govc from source see: https://github.com/vmware/govmomi/blob/master/govc/README.md#source'
  exit 1
fi

PROJECT_PKG="github.com/vmware/govmomi"
PROGRAM_PKG="${PROJECT_PKG}/$(basename "$(dirname "${0}")")"

CDIR=$(cd "$(dirname "${0}")" && pwd)
cd "$CDIR"
# Workaround when GOPATH is not defined
mkdir -p "gopath/src/$(dirname "${PROJECT_PKG}")"
if [ ! -s "gopath/src/${PROJECT_PKG}" ]; then
  ln -sf ../../../../../ "gopath/src/${PROJECT_PKG}"
fi

export GOPATH="${CDIR}/gopath"
export LDFLAGS="-w -X ${PROGRAM_PKG}/version.gitVersion=${git_version}"
export BUILD_OS="${BUILD_OS:-darwin linux windows freebsd}"
export BUILD_ARCH="${BUILD_ARCH:-386 amd64}"

set -x
make -C "${GOPATH}/src/${PROGRAM_PKG}" -j build-all
set +x
