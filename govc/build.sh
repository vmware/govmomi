#!/bin/bash -e

if ! which gox > /dev/null; then
  echo "gox is not installed..."
  exit 1
fi

git_version=$(git describe --tags)
if git_status=$(git status --porcelain 2>/dev/null) && [ -z "${git_status}" ]; then
  git_version="${git_version}-dirty"
fi

ldflags="-X github.com/vmware/govmomi/govc/version.gitVersion ${git_version}"
os="darwin linux windows freebsd"
arch="386 amd64"

gox \
  -parallel=1 \
  -ldflags="${ldflags}" \
  -os="${os}" \
  -arch="${arch}" \
  github.com/vmware/govmomi/govc
