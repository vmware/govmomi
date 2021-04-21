# How to create a new `govmomi` Release on Github

On every new tag matching `v*` pushed to the repository a Github Action Release
Workflow is executed. 

The Github Actions release [workflow](.github/workflows/govmomi-release.yaml)
uses [`goreleaser`](http://goreleaser.com/) ([configuration
file](.goreleaser.yml)) and automatically creates/pushes:

- Release artifacts for `govc` and `vcsim` to the
  [release](https://github.com/vmware/govmomi/releases) page, including
  `LICENSE.txt`, `README` and `CHANGELOG`
- Docker images for `vmware/govc` and `vmware/vcsim` to Docker Hub
- Source code

## Verify `master` branch is up to date with the remote

```console
$ git checkout master
$ git fetch -avp
$ git diff master origin/master

# if your local and remote branches diverge run
$ git pull origin/master
```

⚠️ **Note:** These steps assume `origin` to point to the remote
`https://github.com/vmware/govmomi`, respectively
`git@github.com:vmware/govmomi`.

## Create the new Version Tag


Create a new release tag adhering to the [semantic
versioning](https://semver.org/) scheme:

```console
# Example
TAG=v0.25.0
git tag -a ${TAG} -m "Release ${TAG}"
```

## Push the new Tag

```console
# Will trigger Github Actions Release Workflow
git push origin refs/tags/${TAG}
```

## Verify Github Action Release Workflow

After pushing a new release tag, the status of the
workflow can be inspected
[here](https://github.com/vmware/govmomi/actions/workflows/govmomi-release.yaml).

![Release](static/release-workflow.png "Successful Release Run")