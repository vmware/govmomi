#!/bin/sh

# © Broadcom. All Rights Reserved.
# The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
# SPDX-License-Identifier: Apache-2.0

set -e

DEFAULT_COMMAND="govc"
DEFAULT_ARGS="version"

if [ "$#" -eq 0 ]; then
    exec "$DEFAULT_COMMAND" $DEFAULT_ARGS
fi

if [ "$1" = "govc" ]; then
    exec "$@"
else
    exec /bin/sh -c "$*"
fi
