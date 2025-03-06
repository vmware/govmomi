module github.com/vmware/govmomi/vcsim

go 1.23

replace github.com/vmware/govmomi => ../

require (
	github.com/google/uuid v1.6.0
	github.com/vmware/govmomi v0.0.0-00010101000000-000000000000
)

require github.com/Masterminds/semver/v3 v3.3.1 // indirect
