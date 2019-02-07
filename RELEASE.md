# govmomi release process

Releasing a new version of govmomi and its binaries (currently `govc` and `vcsim`) requires three steps:

1. Bump the `Version` version var in [`govc/flags/version.go`](https://github.com/vmware/govmomi/blob/master/govc/flags/version.go).
2. Run `make prepare-release` locally, commit, and merge to `master`.
3. Tag a version with `git tag -m <version> <version>` and push the tag back to origin with `git push origin "refs/tags/<version>`

Once the tag is pushed to origin, Travis CI will use `goreleaser` to generate the artifacts and publish them.

Currently we are publishing:

Github releases:
- Available on [this Github page](https://github.com/vmware/govmomi/releases)

Docker images:

- `docker pull vmware/govc`
- `docker pull vmware/vcsim`

Homebrew formulas:

- `brew install govmomi/tap/govc`
- `brew install govmomi/tap/vcsim`