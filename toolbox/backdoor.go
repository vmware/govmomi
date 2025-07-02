// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package toolbox

import (
	"errors"

	"github.com/vmware/vmw-guestinfo/message"
	"github.com/vmware/vmw-guestinfo/vmcheck"
)

const (
	rpciProtocol uint32 = 0x49435052
	tcloProtocol uint32 = 0x4f4c4354
)

var (
	ErrNotVirtualWorld = errors.New("not in a virtual world")
)

type backdoorChannel struct {
	protocol uint32

	*message.Channel
}

func (b *backdoorChannel) Start() error {
	if !vmcheck.IsVirtualCPU() {
		return ErrNotVirtualWorld
	}

	channel, err := message.NewChannel(b.protocol)
	if err != nil {
		return err
	}

	b.Channel = channel

	return nil
}

func (b *backdoorChannel) Stop() error {
	if b.Channel == nil {
		return nil
	}

	err := b.Channel.Close()

	b.Channel = nil

	return err
}

// NewBackdoorChannelOut creates a Channel for use with the RPCI protocol
func NewBackdoorChannelOut() Channel {
	return &backdoorChannel{
		protocol: rpciProtocol,
	}
}

// NewBackdoorChannelIn creates a Channel for use with the TCLO protocol
func NewBackdoorChannelIn() Channel {
	return &backdoorChannel{
		protocol: tcloProtocol,
	}
}
