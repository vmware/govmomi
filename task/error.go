// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package task

import (
	"strings"

	"github.com/vmware/govmomi/vim25/types"
)

type Error struct {
	*types.LocalizedMethodFault
	Description *types.LocalizableMessage
}

// Error returns the task's localized fault message, followed by the details
// carried in the fault itself, see FaultString.
func (e Error) Error() string {
	return FaultString(e.LocalizedMethodFault)
}

func (e Error) Fault() types.BaseMethodFault {
	return e.LocalizedMethodFault.Fault
}

// FaultString builds a message from f's localized message and the details
// carried by its fault: every FaultMessage entry and, for NoCompatibleHost,
// the nested per-host errors (recursively). Details are appended after the
// localized message as "msg: detail1; detail2". Empty messages, "[context]"
// entries and duplicates of the localized message are omitted.
func FaultString(f *types.LocalizedMethodFault) string {
	if f == nil {
		return ""
	}

	var details []string
	collectFaultDetails(f, &details, 0)

	msg := f.LocalizedMessage
	var extra []string
	seen := map[string]bool{msg: true}
	for _, d := range details {
		if !seen[d] {
			seen[d] = true
			extra = append(extra, d)
		}
	}

	if len(extra) == 0 {
		return msg
	}
	if msg == "" {
		return strings.Join(extra, "; ")
	}
	// Trim trailing colons and whitespace from localized message to avoid ": :" artifacts
	trimmed := strings.TrimRight(msg, ": \n\r\t")
	return trimmed + ": " + strings.Join(extra, "; ")
}

// maxFaultDepth guards against pathological or cyclic nesting.
const maxFaultDepth = 16

func collectFaultDetails(f *types.LocalizedMethodFault, out *[]string, depth int) {
	if f == nil || f.Fault == nil || depth > maxFaultDepth {
		return
	}

	for _, m := range f.Fault.GetMethodFault().FaultMessage {
		if m.Message != "" && !strings.HasPrefix(m.Message, "[context]") {
			*out = append(*out, m.Message)
		}
	}

	if n, ok := f.Fault.(types.BaseNoCompatibleHost); ok {
		for i := range n.GetNoCompatibleHost().Error {
			nested := &n.GetNoCompatibleHost().Error[i]
			if nested.LocalizedMessage != "" {
				*out = append(*out, nested.LocalizedMessage)
			}
			collectFaultDetails(nested, out, depth+1)
		}
	}
}
