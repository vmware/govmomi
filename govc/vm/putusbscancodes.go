/*
Copyright (c) 2014-2016 VMware, Inc. All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package vm

import (
	"context"
	"flag"
	"fmt"
	"strconv"

	"github.com/vmware/govmomi/govc/cli"
	"github.com/vmware/govmomi/govc/flags"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/vim25/types"
)

type putusbscancode struct {
	*flags.ClientFlag
	*flags.SearchFlag

	UsbHidCodeValue int32
	UsbHidCode      string
	UsbHideString   string
	LeftControl     bool
	LeftShift       bool
	LeftAlt         bool
	LeftGui         bool
	RightControl    bool
	RightShift      bool
	RightAlt        bool
	RightGui        bool
}

func init() {
	cli.Register("vm.putusbscancode", &putusbscancode{})
}

func (cmd *putusbscancode) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)

	cmd.SearchFlag, ctx = flags.NewSearchFlag(ctx, flags.SearchVirtualMachines)
	cmd.SearchFlag.Register(ctx, f)

	f.Var(flags.NewInt32(&cmd.UsbHidCodeValue), "r", "Raw USB HID Code Value (int32)")
	f.StringVar(&cmd.UsbHidCode, "c", "0x00", "USB HID Code (hex)")
	f.StringVar(&cmd.UsbHideString, "s", "", "Raw String to Send")
	f.BoolVar(&cmd.LeftControl, "lc", false, "Enable/Disable Left Control")
	f.BoolVar(&cmd.LeftShift, "ls", false, "Enable/Disable Left Shift")
	f.BoolVar(&cmd.LeftAlt, "la", false, "Enable/Disable Left Alt")
	f.BoolVar(&cmd.LeftGui, "lg", false, "Enable/Disable Left Gui")
	f.BoolVar(&cmd.RightControl, "rc", false, "Enable/Disable Right Control")
	f.BoolVar(&cmd.RightShift, "rs", false, "Enable/Disable Right Shift")
	f.BoolVar(&cmd.RightAlt, "ra", false, "Enable/Disable Right Alt")
	f.BoolVar(&cmd.RightGui, "rg", false, "Enable/Disable Right Gui")
}

func (cmd *putusbscancode) Usage() string {
	return ""
}

func (cmd *putusbscancode) Description() string {
	return `Send Keystroke to VM.

Examples:
 Default Scenario
  govc vm.putusbscancode $vm -r 1376263 (writes an 'r' to the console)
  govc vm.putusbscancode -c 0x15 (writes an 'r' to the console)
  govc vm.putusbscancode -s "root" (writes 'root' to the console)
  govc vm.putusbscancode -c 0x58 (presses ENTER on the console)
  govc vm.putusbscancode -c 0x4c -la true -lc true (sends CTRL+ALT+DEL to console)`
}

func (cmd *putusbscancode) Process(ctx context.Context) error {
	if err := cmd.ClientFlag.Process(ctx); err != nil {
		return err
	}
	if err := cmd.SearchFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *putusbscancode) Run(ctx context.Context, f *flag.FlagSet) error {
	vms, err := cmd.VirtualMachines(f.Args())
	if err != nil {
		return err
	}

	for _, vm := range vms {
		err := cmd.processUserInput(ctx, vm)
		if err != nil {
			return err
		}
	}
	return err
}

func (cmd *putusbscancode) processUserInput(ctx context.Context, vm *object.VirtualMachine) error {
	if err := cmd.checkValidInputs(); err != nil {
		return err
	}
	modifiers := types.UsbScanCodeSpecModifierType{
		LeftControl:  &cmd.LeftControl,
		LeftShift:    &cmd.LeftShift,
		LeftAlt:      &cmd.LeftAlt,
		LeftGui:      &cmd.LeftGui,
		RightControl: &cmd.RightControl,
		RightShift:   &cmd.RightShift,
		RightAlt:     &cmd.RightAlt,
		RightGui:     &cmd.RightGui,
	}
	codes, err := cmd.processUsbCode()
	if err != nil {
		return err
	}
	var keyEventArray []types.UsbScanCodeSpecKeyEvent
	for code := range codes {
		keyEvent := types.UsbScanCodeSpecKeyEvent{
			UsbHidCode: int32(code),
			Modifiers:  &modifiers,
		}

		keyEventArray = append(keyEventArray, keyEvent)
	}
	spec := types.UsbScanCodeSpec{
		KeyEvents: keyEventArray,
	}
	_, err = vm.PutUsbScanCodes(ctx, spec)
	return err
}

func (cmd *putusbscancode) processUsbCode() ([]int32, error) {
	// check to ensure only 1 input is specified
	if cmd.UsbHidCode != "" &&
		cmd.UsbHidCodeValue != 0 &&
		cmd.UsbHideString != "" {
		return nil, fmt.Errorf("Specify only 1 argument for HID code")
	}
	if cmd.rawCodeProvided() {
		return []int32{cmd.UsbHidCodeValue}, nil
	}
	if cmd.hexCodeProvided() {
		s, err := hexToHidCode(cmd.UsbHidCode)
		if err != nil {
			return nil, err
		}
		return []int32{int32(s)}, nil
	}
	if cmd.stringProvided() {
		return nil, fmt.Errorf("Not yet supported")
	}
	return nil, nil
}

func hexToHidCode(hex string) (int32, error) {
	s, err := strconv.ParseInt(hex, 0, 32)
	if err != nil {
		return 0, err
	}
	s = s << 16
	s = s | 7
	return int32(s), nil
}

func (cmd *putusbscancode) checkValidInputs() error {
	// poor man's boolean XOR -> A xor B xor C = A'BC' + AB'C' + A'B'C + ABC
	if (!cmd.rawCodeProvided() && cmd.hexCodeProvided() && !cmd.stringProvided()) || // A'BC'
		(cmd.rawCodeProvided() && !cmd.hexCodeProvided() && !cmd.stringProvided()) || // AB'C'
		(!cmd.rawCodeProvided() && !cmd.hexCodeProvided() && cmd.stringProvided()) || // A'B'C
		(cmd.rawCodeProvided() && cmd.hexCodeProvided() && cmd.stringProvided()) { // ABC
		return nil
	}
	return fmt.Errorf("Specify only 1 argument")
}

func (cmd putusbscancode) rawCodeProvided() bool {
	return cmd.UsbHidCodeValue != 0
}

func (cmd putusbscancode) hexCodeProvided() bool {
	return cmd.UsbHidCode != ""
}

func (cmd putusbscancode) stringProvided() bool {
	return cmd.UsbHideString != ""
}
