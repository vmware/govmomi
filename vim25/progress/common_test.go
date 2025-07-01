// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package progress

type dummySinker struct {
	ch chan Report
}

func (d dummySinker) Sink() chan<- Report {
	return d.ch
}

type dummyReport struct {
	p float32
	d string
	e error
}

func (p dummyReport) Percentage() float32 {
	return p.p
}

func (p dummyReport) Detail() string {
	return p.d
}

func (p dummyReport) Error() error {
	return p.e
}
