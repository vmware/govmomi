#!/bin/bash -e

# Generate vSphere Cloud Provider config from govc env and run local-up-cluster.sh
# Assumes the create.sh NFS configuration has been applied.

GOVC_NETWORK=${GOVC_NETWORK:-"VM Network"}
GOVC_DATACENTER=${GOVC_DATACENTER:-"$(govc find / -type d)"}
GOVC_DATACENTER="$(basename "$GOVC_DATACENTER")"
GOVC_CLUSTER=${GOVC_CLUSTER:-"$(govc find / -type c -type r)"}

oneline() {
  awk '{printf "%s\\n", $0}' "$1" # make gcfg happy
}

username="$(govc env GOVC_USERNAME)"
password="$(govc env GOVC_PASSWORD)"
if [ -n "$GOVC_CERTIFICATE" ] ; then
  username="$(oneline "$GOVC_CERTIFICATE")"
  password="$(oneline "$GOVC_PRIVATE_KEY")"
fi

cat <<EOF | tee vcp.conf
[Global]
        insecure-flag = "$(govc env GOVC_INSECURE)"

[VirtualCenter "$(govc env -x GOVC_URL_HOST)"]
        user = "$username"
        password = "$password"
        port = "$(govc env -x GOVC_URL_PORT)"
        datacenters = "$(basename "$GOVC_DATACENTER")"

[Workspace]
        server = "$(govc env -x GOVC_URL_HOST)"
        datacenter = "$GOVC_DATACENTER"
        folder = "vm"
        default-datastore = "$GOVC_DATACENTER"
        resourcepool-path = "$GOVC_CLUSTER/Resources"
[Disk]
        scsicontrollertype = pvscsi

[Network]
        public-network = "$GOVC_NETWORK"
EOF

k8s="$GOPATH/src/k8s.io/kubernetes"

ip=$(govc vm.ip -a -v4 "$USER-ubuntu-16.04")

ssh-add ~/.vagrant.d/insecure_private_key
ssh "vagrant@$ip" mkdir -p "$k8s"
rsync -auvz "$k8s" "vagrant@$ip:$(dirname "$k8s")"
rsync -auvz "$PWD/vcp.conf" "vagrant@$ip:$k8s"
ssh "vagrant@$ip" "GOPATH=$GOPATH" "$k8s/hack/install-etcd.sh"

# shellcheck disable=2029
ssh -tt </dev/null -L 8080:127.0.0.1:8080 "vagrant@$ip" \
    CLOUD_PROVIDER=vsphere CLOUD_CONFIG="$k8s/vcp.conf" \
    PATH="$PATH:$k8s/third_party/etcd" "$k8s/hack/local-up-cluster.sh"
