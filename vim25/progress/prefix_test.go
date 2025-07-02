// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package progress

import "testing"

func TestPrefix(t *testing.T) {
	var r Report

	ch := make(chan Report, 1)
	s := Prefix(dummySinker{ch}, "prefix").Sink()

	// No detail
	s <- dummyReport{d: ""}
	r = <-ch
	if r.Detail() != "prefix" {
		t.Errorf("Expected detail to be prefixed")
	}

	// With detail
	s <- dummyReport{d: "something"}
	r = <-ch
	if r.Detail() != "prefix: something" {
		t.Errorf("Expected detail to be prefixed")
	}
}
