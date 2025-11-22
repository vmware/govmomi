module github.com/vmware/govmomi/vcsim

go 1.24.0

replace github.com/vmware/govmomi => ../

require (
	github.com/google/uuid v1.6.0
	github.com/vmware/govmomi v0.0.0-00010101000000-000000000000
)

require golang.org/x/text v0.31.0 // indirect
