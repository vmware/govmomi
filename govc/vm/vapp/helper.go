package vapp

import (
	"context"

	"github.com/vmware/govmomi/govc/flags"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"
)

func retrieveVAppConfig(ctx context.Context, vm *object.VirtualMachine) (*types.VmConfigInfo, error) {
	var props mo.VirtualMachine
	if err := vm.Properties(ctx, vm.Reference(), []string{"config.vAppConfig"}, &props); err != nil {
		return nil, err
	}
	var info *types.VmConfigInfo
	if props.Config != nil && props.Config.VAppConfig != nil {
		info = props.Config.VAppConfig.GetVmConfigInfo()
	}
	return info, nil
}

func hasVAppConfig(ctx context.Context, vm *object.VirtualMachine) (bool, error) {
	config, err := retrieveVAppConfig(ctx, vm)
	switch {
	case err != nil:
		return false, err
	case config == nil:
		return false, nil
	}
	return true, nil
}

func waitLog(ctx context.Context, cmd *flags.OutputFlag, task *object.Task, msg string) error {
	logger := cmd.ProgressLogger(msg)
	defer logger.Wait()

	_, err := task.WaitForResult(ctx, logger)
	return err
}
