package convert

import (
	"fmt"

	mo "github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"

	vimv1 "github.com/vmware/govmomi/crd/pkg/vim/api/v1alpha1"
)

func ConvertVIMToCRD(dst any, src any) error {
	switch tDst := dst.(type) {

	case *vimv1.VirtualMachine:
		var tSrc mo.VirtualMachine

		switch tempTSrc := src.(type) {
		case *mo.VirtualMachine:
			tSrc = *tempTSrc
		case mo.VirtualMachine:
			tSrc = tempTSrc
		default:
			return fmt.Errorf(
				"src is not a %[1]s or *%[1]s",
				"mo.VirtualMachine")
		}

		if tSrc.Config != nil {
			tDst.Spec.Config = &vimv1.VirtualMachineConfigInfoSpec{}
			if err := ConvertVIMToCRD(
				tDst.Spec.Config,
				tSrc.Config); err != nil {

				return fmt.Errorf("failed to convert moVM.Config: %w", err)
			}
		}

	case *vimv1.VirtualMachineConfigInfoSpec:
		var tSrc types.VirtualMachineConfigInfo

		switch tempTSrc := src.(type) {
		case *types.VirtualMachineConfigInfo:
			tSrc = *tempTSrc
		case types.VirtualMachineConfigInfo:
			tSrc = tempTSrc
		default:
			return fmt.Errorf(
				"src is not a %[1]s or *%[1]s",
				"types.VirtualMachineConfigInfo")
		}

		tDst.AlternateGuestName = tSrc.AlternateGuestName
		tDst.Annotation = tSrc.Annotation

	}
	return nil
}
