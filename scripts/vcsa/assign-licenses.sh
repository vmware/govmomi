#!/bin/bash -e

# © Broadcom. All Rights Reserved.
# The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
# SPDX-License-Identifier: Apache-2.0

export GOVC_INSECURE=true

govc license.add "$VCSA_LICENSE" "$ESX_LICENSE" >/dev/null

govc license.assign "$VCSA_LICENSE" >/dev/null

govc find / -type h | xargs -I% -n1 govc license.assign -host % "$ESX_LICENSE" >/dev/null

echo "Assigned licenses..."
govc license.assigned.ls

echo ""
echo "License usage..."
govc license.ls
