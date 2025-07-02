// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package toolbox

import (
	"errors"
	"testing"
)

func TestPowerCommandHandler(t *testing.T) {
	shutdown = "/bin/echo"

	in := new(mockChannelIn)
	out := new(mockChannelOut)

	service := NewService(in, out)
	power := service.Power

	// cover nil Handler and out.Receive paths
	_, _ = power.Halt.Dispatch(nil)

	out.reply = append(out.reply, rpciOK, rpciOK)

	power.Halt.Handler = Halt
	power.Reboot.Handler = Reboot
	power.Suspend.Handler = func() error {
		return errors.New("an error")
	}

	commands := []PowerCommand{
		power.Halt,
		power.Reboot,
		power.Suspend,
	}

	for _, cmd := range commands {
		_, _ = cmd.Dispatch(nil)
	}
}
