#!/bin/bash

pushd $(dirname $0)

# Sadly, the ttylinux project was abandoned in late 2015.
# But this release still serves us well.
base_url=https://github.com/dougm/packer-ttylinux/releases/download/16.1u2
ttylinux="ttylinux-pc_i486-16.1"
files="${ttylinux}.iso ${ttylinux}-live.ova ${ttylinux}.ova"

for name in $files ; do
  wget -qO $name $base_url/$name
done

wget -qN https://github.com/icebreaker/floppybird/raw/master/build/floppybird.img

# extract ova so we can also use the .vmdk and .ovf files directly
tar -xvf ${ttylinux}.ova

popd
