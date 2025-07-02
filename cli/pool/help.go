// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package pool

var poolNameHelp = `
POOL may be an absolute or relative path to a resource pool or a (clustered)
compute host. If it resolves to a compute host, the associated root resource
pool is returned. If a relative path is specified, it is resolved with respect
to the current datacenter's "host" folder (i.e. /ha-datacenter/host).

Paths to nested resource pools must traverse through the root resource pool of
the selected compute host, i.e. "compute-host/Resources/nested-pool".

The same globbing rules that apply to the "ls" command apply here. For example,
POOL may be specified as "*/Resources/*" to expand to all resource pools that
are nested one level under the root resource pool, on all (clustered) compute
hosts in the current datacenter.`

var poolCreateHelp = `
POOL may be an absolute or relative path to a resource pool. The parent of the
specified POOL must be an existing resource pool. If a relative path is
specified, it is resolved with respect to the current datacenter's "host"
folder (i.e. /ha-datacenter/host). The basename of the specified POOL is used
as the name for the new resource pool.

The same globbing rules that apply to the "ls" command apply here. For example,
the path to the parent resource pool in POOL may be specified as "*/Resources"
to expand to the root resource pools on all (clustered) compute hosts in the
current datacenter.

For example:
  */Resources/test             Create resource pool "test" on all (clustered)
                               compute hosts in the current datacenter.
  somehost/Resources/*/nested  Create resource pool "nested" in every
                               resource pool that is a direct descendant of
                               the root resource pool on "somehost".`
