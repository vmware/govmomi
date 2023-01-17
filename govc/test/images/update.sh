#!/bin/bash

pushd $(dirname $0)

# Sadly, the ttylinux project was abandoned in late 2015.
# But this release still serves us well.
base_url=https://github.com/dougm/packer-ttylinux/releases/download/16.1u2
ttylinux="ttylinux-pc_i486-16.1"
files="${ttylinux}.iso ${ttylinux}-live.ova ${ttylinux}.ova"
tarbin="tar"
if [ "$(uname)" = "Darwin" ]; then
  # macOS BSD tar creates additional entries that govc tar archive reader
  # chokes on (instead of ignoring). Use GNU tar instead.
  tarbin="gtar"
fi

for name in $files ; do
  wget -qO $name $base_url/$name
done

wget -qN https://github.com/icebreaker/floppybird/raw/master/build/floppybird.img

# extract ova so we can also use the .vmdk and .ovf files directly
$tarbin -xvf ${ttylinux}.ova

# create an ova with "bad" checksum in manifest by
# modifying/replacing .mf in copy of ${ttylinux}.ova
mkdir -p "$(pwd)/tmp"
$tarbin xv -C "$(pwd)/tmp" -f ${ttylinux}.ova
pushd "$(pwd)/tmp"
sed 's/=.*$/= 5e82716003a1bff5b1d27bbd7d1e83addc881503/g' < ${ttylinux}.mf > ${ttylinux}-bad-checksum.mf
mv ${ttylinux}-bad-checksum.mf ${ttylinux}.mf
popd
$tarbin cv -C "$(pwd)/tmp" -f ${ttylinux}-bad-checksum.ova .
rm -fr "$(pwd)/tmp"

popd
