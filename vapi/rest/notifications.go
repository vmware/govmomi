// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package rest

import "time"

// Notification contains fields to describe any info/warning/error messages
// that Tasks can raise.
type Notification struct {
	// Id is the notification id.
	Id string `json:"id"`

	// Time the notification was raised/found.
	Time time.Time `json:"time"`

	Message LocalizableMessage `json:"message"`

	Resolution LocalizableMessage `json:"resolution"`
}

// Notifications contains info/warning/error messages that can be reported.
type Notifications struct {
	// Info notification messages reported.
	Info []Notification `json:"info"`

	// Warning notification messages reported.
	Warnings []Notification `json:"warnings"`

	// Errors notification messages reported.
	Errors []Notification `json:"errors"`
}
