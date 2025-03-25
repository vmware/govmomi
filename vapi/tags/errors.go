// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package tags

import (
	"fmt"
)

const (
	errFormat = "[error: %d type: %s reason: %s]"
	separator = "," // concat multiple error strings
)

// BatchError is an error returned for a single item which failed in a batch
// operation
type BatchError struct {
	Type    string `json:"id"`
	Message string `json:"default_message"`
}

// BatchErrors contains all errors which occurred in a batch operation
type BatchErrors []BatchError

func (b BatchErrors) Error() string {
	if len(b) == 0 {
		return ""
	}

	var errString string
	for i := range b {
		errType := b[i].Type
		reason := b[i].Message
		errString += fmt.Sprintf(errFormat, i, errType, reason)

		// no separator after last item
		if i+1 < len(b) {
			errString += separator
		}
	}
	return errString
}
