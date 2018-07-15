#!/bin/bash -e

curl -fsSL https://download.docker.com/linux/ubuntu/gpg | apt-key add -

add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu \
   $(lsb_release -cs) \
   stable"

apt-get update
DEBIAN_FRONTEND=noninteractive apt-get -y -o Dpkg::Options::="--force-confdef" -o Dpkg::Options::="--force-confold" upgrade
apt-get -y install docker-ce

usermod -aG docker vagrant

if [ ! -d "/usr/local/go" ] ; then
  (cd /usr/local &&
     (curl --silent -L https://dl.google.com/go/go1.10.3.linux-amd64.tar.gz | tar -zxf -) &&
     ln -s /usr/local/go/bin/* /usr/local/bin/)
fi
