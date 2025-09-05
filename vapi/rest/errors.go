// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package rest

import "encoding/json"

type Type string

const (
	ErrError Type = "ERROR"

	ErrorAlreadyExists Type = "ALREADY_EXISTS"

	ErrAlreadyInDesiredState Type = "ALREADY_IN_DESIRED_STATE"

	ErrCanceled Type = "CANCELED"

	ErrConcurrentChange Type = "CONCURRENT_CHANGE"

	ErrFeatureInUse Type = "FEATURE_IN_USE"

	ErrInternalServer Type = "INTERNAL_SERVER_ERROR"

	ErrInvalidArgument Type = "INVALID_ARGUMENT"

	ErrInvalidElementConfiguration Type = "INVALID_ELEMENT_CONFIGURATION"

	ErrInvalidElementType Type = "INVALID_ELEMENT_TYPE"

	ErrInvalidRequest Type = "INVALID_REQUEST"

	ErrNotAllowedInCurrentState Type = "NOT_ALLOWED_IN_CURRENT_STATE"

	ErrNotFound Type = "NOT_FOUND"

	ErrOperationNotFound Type = "OPERATION_NOT_FOUND"

	ErrResourceBusy Type = "RESOURCE_BUSY"

	ErrResourceInUse Type = "RESOURCE_IN_USE"

	ErrResourceInaccessible Type = "RESOURCE_INACCESSIBLE"

	ErrServiceUnavailable Type = "SERVICE_UNAVAILABLE"

	ErrTimedOut Type = "TIMED_OUT"

	ErrUnableToAllocateResource Type = "UNABLE_TO_ALLOCATE_RESOURCE"

	ErrUnauthenticated Type = "UNAUTHENTICATED"

	ErrUnauthorized Type = "UNAUTHORIZED"

	ErrUnexpectedInput Type = "UNEXPECTED_INPUT"

	ErrUnsupported Type = "UNSUPPORTED"

	ErrUnverifiedPeer Type = "UNVERIFIED_PEER"
)

type Error struct {
	Messages []LocalizableMessage `json:"messages"`

	Data json.RawMessage `json:"data"`

	ErrorType Type `json:"error_type"`
}
