#!/usr/bin/env bats

load test_helper

vcsim_guest() {
  require_docker
  vcsim_env -autostart=false

  vm=DC0_H0_VM0

  name=$(docker_name $vm)

  if docker inspect "$name" ; then
    flunk "$vm container still exists"
  fi

  export GOVC_VM=$vm GOVC_GUEST_LOGIN=user:pass

  run govc object.collect -s vm/$vm guest.toolsStatus
  assert_success toolsNotInstalled

  run govc object.collect -s vm/$vm guest.toolsRunningStatus
  assert_success guestToolsNotRunning

  govc vm.change -vm $vm -e "RUN.container=--tmpfs /tmp nginx"

  govc vm.power -on $vm

  if ! docker inspect "$name" ; then
    flunk "$vm container does not exist"
  fi

  run govc object.collect -s vm/$vm guest.toolsStatus
  assert_success toolsOk

  run govc object.collect -s vm/$vm guest.toolsRunningStatus
  assert_success guestToolsRunning
}

@test "guest file manager" {
  vcsim_guest

  run govc guest.mkdir /tmp/foo
  assert_success

  run govc guest.mkdir /tmp/foo/bar/baz
  assert_failure

  run govc guest.mkdir -p /tmp/foo/bar/baz
  assert_success

  run govc guest.rmdir /tmp/foo
  assert_failure # not empty

  run govc guest.rmdir -r /tmp/foo
  assert_success

  run govc guest.mktemp -d
  assert_success
  tmp="$output"

  run govc guest.rmdir "$tmp"
  assert_success

  run govc guest.rmdir "$tmp"
  assert_failure # does not exist

  run govc guest.mktemp
  assert_success
  tmp="$output"

  run govc guest.ls
  assert_failure # InvalidArgument

  run govc guest.ls /enoent
  assert_failure # InvalidArgument

  run govc guest.ls "$tmp"
  assert_success
  assert_matches "$tmp"
  assert_matches "rw-------" # 0600


  run govc guest.chmod 0644 "$tmp"
  assert_success

  run govc guest.chown 0:0 "$tmp"
  assert_success

  run govc guest.ls "$tmp"
  assert_success
  assert_matches "$tmp"
  assert_matches "rw-r--r--" # 0644

  run govc guest.ls "$(dirname "$tmp")"
  assert_success
  assert_matches "$tmp"

  run govc guest.mv "$tmp" "$tmp-new"
  assert_success

  run govc guest.ls "$tmp"
  assert_failure

  run govc guest.ls "$tmp-new"
  assert_success

  run govc guest.upload -l '' README.md "$tmp"
  assert_failure # unauthenticated

  run govc guest.upload README.md "$tmp"
  assert_success

  run govc guest.download "$tmp" -
  assert_success "$(cat README.md)"

  run govc guest.rmdir "$tmp"
  assert_failure # not a directory

  run govc guest.rm "$tmp"
  assert_success

  run govc guest.rm "$tmp"
  assert_failure # does not exist
}

@test "guest process manager" {
  vcsim_guest

  run govc guest.kill -p 123456
  assert_failure # process does not exist

  run govc guest.start -l '' /bin/df
  assert_failure # unauthenticated

  run govc guest.start /bin/df -h
  assert_success
  pid="$output"

  run govc guest.ps -p "$pid"
  assert_success
  assert_matches /bin/df

  run govc guest.ps -x
  assert_success

  run govc guest.kill -p "$pid"
  assert_success

  run govc guest.run uname -a
  assert_success
  assert_matches Linux

  run govc guest.run -e FOO=bar -C /tmp env
  assert_success
  assert_matches FOO=bar
  assert_matches PWD=/tmp
}

@test "guest tools status" {
  vcsim_guest

  run govc vm.power -r $GOVC_VM
  assert_success

  run govc object.collect -s vm/$vm guest.toolsStatus
  assert_success toolsOk

  run govc object.collect -s vm/$vm guest.toolsRunningStatus
  assert_success guestToolsRunning

  run govc vm.power -off $GOVC_VM
  assert_success

  run govc object.collect -s vm/$vm guest.toolsStatus
  assert_success toolsNotRunning

  run govc object.collect -s vm/$vm guest.toolsRunningStatus
  assert_success guestToolsNotRunning

  run govc guest.run uname -a
  assert_failure # powered off
}
