// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package fault_test

import (
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/vmware/govmomi/fault"
	"github.com/vmware/govmomi/task"
	"github.com/vmware/govmomi/vim25/soap"
	"github.com/vmware/govmomi/vim25/types"
)

func asPtr[T any](args ...T) *T {
	if len(args) == 0 {
		var t T
		return &t
	}
	return &args[0]
}

var (
	nilSystemErr       = asPtr[*types.SystemError]()
	nilInvalidVmConfig = asPtr[*types.InvalidVmConfig]()
)

type nilErrWrapper struct{}

func (e nilErrWrapper) Unwrap() error {
	return nil
}

type unwrappableErrSlice []error

func (e unwrappableErrSlice) Unwrap() []error {
	return e
}

type nilBaseMethodFault struct{}

func (f nilBaseMethodFault) GetMethodFault() *types.MethodFault {
	return nil
}

type asFaulter struct {
	val  any
	okay bool
	msg  string
}

func (f asFaulter) AsFault(target any) (string, bool) {
	if !f.okay {
		return "", false
	}
	targetVal := reflect.ValueOf(target)
	targetVal.Elem().Set(reflect.ValueOf(f.val))
	return f.msg, true
}

type isFaulter bool

func (f isFaulter) GetMethodFault() *types.MethodFault {
	return nil
}

func (f isFaulter) IsFault(target types.BaseMethodFault) bool {
	return bool(f)
}

func TestAs(t *testing.T) {

	testCases := []struct {
		name                     string
		err                      any
		target                   any
		expectedLocalizedMessage string
		expectedOkay             bool
		expectedPanic            any
		expectedTarget           any
	}{

		{
			name:           "err is nil",
			err:            nil,
			target:         asPtr[*types.SystemError](),
			expectedTarget: asPtr[*types.SystemError](),
		},
		{
			name:           "err is not supported",
			err:            struct{}{},
			target:         asPtr[*types.SystemError](),
			expectedTarget: asPtr[*types.SystemError](),
		},
		{
			name:           "target is nil",
			err:            errors.New("error"),
			target:         nil,
			expectedTarget: nil,
			expectedPanic:  "fault: target cannot be nil",
		},
		{
			name:           "target is not pointer",
			err:            errors.New("error"),
			target:         types.SystemError{},
			expectedTarget: types.SystemError{},
			expectedPanic:  "fault: target must be a non-nil pointer",
		},
		{
			name:           "target is not pointer to expected type",
			err:            errors.New("error"),
			target:         &types.SystemError{},
			expectedTarget: &types.SystemError{},
			expectedPanic:  "fault: *target must be interface or implement BaseMethodFault",
		},
		{
			name: "err is task.Error with fault",
			err: task.Error{
				LocalizedMethodFault: &types.LocalizedMethodFault{
					Fault: &types.SystemError{},
				},
			},
			target:                   asPtr[*types.SystemError](),
			expectedTarget:           asPtr(&types.SystemError{}),
			expectedOkay:             true,
			expectedLocalizedMessage: "",
		},
		{
			name:           "err unwraps to nil error",
			err:            nilErrWrapper{},
			target:         asPtr[*types.SystemError](),
			expectedTarget: asPtr[*types.SystemError](),
		},
		{
			name:           "err unwraps to nil error slice",
			err:            unwrappableErrSlice{},
			target:         asPtr[*types.SystemError](),
			expectedTarget: asPtr[*types.SystemError](),
		},
		{
			name: "err is wrapped task.Error with fault",
			err: fmt.Errorf("my error: %w", task.Error{
				LocalizedMethodFault: &types.LocalizedMethodFault{
					Fault: &types.SystemError{},
				},
			}),
			target:                   asPtr[*types.SystemError](),
			expectedTarget:           asPtr(&types.SystemError{}),
			expectedOkay:             true,
			expectedLocalizedMessage: "",
		},
		{
			name: "err is wrapped nil error and task.Error with fault",
			err: unwrappableErrSlice{
				nil,
				task.Error{
					LocalizedMethodFault: &types.LocalizedMethodFault{
						Fault: &types.SystemError{},
					},
				},
			},
			target:                   asPtr[*types.SystemError](),
			expectedTarget:           asPtr(&types.SystemError{}),
			expectedOkay:             true,
			expectedLocalizedMessage: "",
		},
		{
			name:                     "err is types.BaseMethodFault",
			err:                      &types.SystemError{},
			target:                   asPtr[*types.SystemError](),
			expectedTarget:           asPtr(&types.SystemError{}),
			expectedOkay:             true,
			expectedLocalizedMessage: "",
		},
		{
			name:                     "err is types.BaseMethodFault that returns nil",
			err:                      nilBaseMethodFault{},
			target:                   asPtr[*types.SystemError](),
			expectedTarget:           asPtr[*types.SystemError](),
			expectedLocalizedMessage: "",
		},
		{
			name:                     "err is asFaulter that returns true and fault",
			err:                      asFaulter{okay: true, msg: "Hello, world.", val: &types.SystemError{}},
			target:                   asPtr[*types.SystemError](),
			expectedTarget:           asPtr(&types.SystemError{}),
			expectedOkay:             true,
			expectedLocalizedMessage: "Hello, world.",
		},
		{
			name:                     "err is asFaulter that returns false",
			err:                      asFaulter{okay: false, msg: "Hello, world."},
			target:                   asPtr[*types.SystemError](),
			expectedTarget:           asPtr[*types.SystemError](),
			expectedLocalizedMessage: "",
		},
		{
			name: "err is *types.LocalizedMethodFault",
			err: &types.LocalizedMethodFault{
				LocalizedMessage: "fake",
				Fault:            &types.SystemError{},
			},
			target:                   asPtr[*types.SystemError](),
			expectedTarget:           asPtr(&types.SystemError{}),
			expectedOkay:             true,
			expectedLocalizedMessage: "fake",
		},
		{
			name: "err is *types.LocalizedMethodFault w nil fault",
			err: &types.LocalizedMethodFault{
				LocalizedMessage: "fake",
				Fault:            nil,
			},
			target:         asPtr[*types.SystemError](),
			expectedTarget: asPtr[*types.SystemError](),
			expectedOkay:   false,
		},
		{
			name: "err is task.Error with nested fault",
			err: &types.LocalizedMethodFault{
				LocalizedMessage: "fake1",
				Fault: &types.SystemError{
					RuntimeFault: types.RuntimeFault{
						MethodFault: types.MethodFault{
							FaultCause: &types.LocalizedMethodFault{
								LocalizedMessage: "fake2",
								Fault:            &types.InvalidVmConfig{},
							},
						},
					},
				},
			},
			target:                   asPtr[*types.InvalidVmConfig](),
			expectedTarget:           asPtr(&types.InvalidVmConfig{}),
			expectedOkay:             true,
			expectedLocalizedMessage: "fake2",
		},
		{
			name: "err is soap fault",
			err: soap.WrapSoapFault(&soap.Fault{
				Detail: struct {
					Fault types.AnyType "xml:\",any,typeattr\""
				}{
					Fault: &types.SystemError{
						RuntimeFault: types.RuntimeFault{
							MethodFault: types.MethodFault{
								FaultCause: &types.LocalizedMethodFault{
									Fault: &types.InvalidVmConfig{},
								},
							},
						},
					},
				},
			}),
			target:         asPtr[*types.InvalidVmConfig](),
			expectedTarget: asPtr(&types.InvalidVmConfig{}),
			expectedOkay:   true,
		},
		{
			name: "err is soap fault with nil vim fault",
			err: soap.WrapSoapFault(&soap.Fault{
				Detail: struct {
					Fault types.AnyType "xml:\",any,typeattr\""
				}{
					Fault: nil,
				},
			}),
			target:         asPtr[*types.InvalidVmConfig](),
			expectedTarget: asPtr[*types.InvalidVmConfig](),
			expectedOkay:   false,
		},
		{
			name: "err is vim fault",
			err: soap.WrapVimFault(&types.SystemError{
				RuntimeFault: types.RuntimeFault{
					MethodFault: types.MethodFault{
						FaultCause: &types.LocalizedMethodFault{
							Fault: &types.InvalidVmConfig{},
						},
					},
				},
			}),
			target:         asPtr[*types.InvalidVmConfig](),
			expectedTarget: asPtr(&types.InvalidVmConfig{}),
			expectedOkay:   true,
		},
		{
			name: "err is vim fault with nil vim fault",
			err: soap.WrapVimFault(&types.SystemError{
				RuntimeFault: types.RuntimeFault{
					MethodFault: types.MethodFault{
						FaultCause: &types.LocalizedMethodFault{
							Fault: nil,
						},
					},
				},
			}),
			target:         asPtr[*types.InvalidVmConfig](),
			expectedTarget: asPtr[*types.InvalidVmConfig](),
			expectedOkay:   false,
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			var (
				okay             bool
				localizedMessage string
			)

			if tc.expectedPanic != nil {
				assert.PanicsWithValue(
					t,
					tc.expectedPanic,
					func() {
						localizedMessage, okay = fault.As(tc.err, tc.target)
					})
			} else {
				localizedMessage, okay = fault.As(tc.err, tc.target)
			}

			assert.Equal(t, tc.expectedOkay, okay)
			assert.Equal(t, tc.expectedLocalizedMessage, localizedMessage)
			assert.Equal(t, tc.expectedTarget, tc.target)
		})
	}
}

type faultResult struct {
	fault               types.BaseMethodFault
	localizedMessage    string
	localizableMessages []types.LocalizableMessage
}

type testHarness struct {
	isNil                bool
	numCalls             int
	numCallsToReturnTrue int
	results              []faultResult
}

func (h *testHarness) Callback(
	fault types.BaseMethodFault,
	localizedMessage string,
	localizableMessages []types.LocalizableMessage) bool {

	h.numCalls++

	rvFault := reflect.ValueOf(fault)
	if rvFault.Kind() == reflect.Pointer {
		rvFault = rvFault.Elem()
	}
	emptyFault := reflect.New(rvFault.Type())

	h.results = append(h.results, faultResult{
		fault:               emptyFault.Interface().(types.BaseMethodFault),
		localizedMessage:    localizedMessage,
		localizableMessages: localizableMessages,
	})
	return h.numCallsToReturnTrue == h.numCalls
}

func TestIn(t *testing.T) {
	const unsupported = "fault: err must implement error, types.BaseMethodFault, or types.HasLocalizedMethodFault"

	testCases := []struct {
		name             string
		err              any
		callback         *testHarness
		expectedNumCalls int
		expectedResults  []faultResult
		expectedPanic    any
	}{
		{
			name:             "err is nil",
			err:              nil,
			callback:         &testHarness{},
			expectedNumCalls: 0,
			expectedPanic:    unsupported,
		},
		{
			name:             "callback is nil",
			err:              errors.New("error"),
			callback:         &testHarness{isNil: true},
			expectedNumCalls: 0,
			expectedPanic:    "fault: onFaultFn must not be nil",
		},
		{
			name:             "err is unsupported",
			err:              struct{}{},
			callback:         &testHarness{},
			expectedNumCalls: 0,
			expectedPanic:    unsupported,
		},
		{
			name:             "err is unsupported but still an err",
			err:              errors.New("error"),
			callback:         &testHarness{},
			expectedNumCalls: 0,
		},
		{
			name:             "error is task.Error",
			err:              task.Error{},
			callback:         &testHarness{},
			expectedNumCalls: 0,
		},
		{
			name: "error is task.Error with localized message and fault",
			err: task.Error{
				LocalizedMethodFault: &types.LocalizedMethodFault{
					LocalizedMessage: "fake",
					Fault:            &types.SystemError{},
				},
			},
			callback:         &testHarness{},
			expectedNumCalls: 1,
			expectedResults: []faultResult{
				{
					localizedMessage: "fake",
					fault:            &types.SystemError{},
				},
			},
		},
		{
			name:             "error is types.SystemError",
			err:              &types.SystemError{},
			callback:         &testHarness{},
			expectedNumCalls: 1,
			expectedResults: []faultResult{
				{
					fault: &types.SystemError{},
				},
			},
		},
		{
			name: "error is task.Error with a localized message but no fault",
			err: task.Error{
				LocalizedMethodFault: &types.LocalizedMethodFault{
					LocalizedMessage: "fake",
				},
			},
			callback:         &testHarness{},
			expectedNumCalls: 0,
		},
		{
			name: "error is multiple, wrapped errors",
			err: unwrappableErrSlice{
				nil,
				task.Error{
					LocalizedMethodFault: &types.LocalizedMethodFault{
						LocalizedMessage: "fake",
						Fault:            &types.SystemError{},
					},
				},
			},
			callback:         &testHarness{},
			expectedNumCalls: 1,
			expectedResults: []faultResult{
				{
					localizedMessage: "fake",
					fault:            &types.SystemError{},
				},
			},
		},
		{
			name: "error has nested task.Error with localized message and fault",
			err: fmt.Errorf(
				"error 1: %w",
				fmt.Errorf(
					"error 2: %w",
					task.Error{
						LocalizedMethodFault: &types.LocalizedMethodFault{
							LocalizedMessage: "fake",
							Fault:            &types.SystemError{},
						},
					},
				),
			),
			callback:         &testHarness{},
			expectedNumCalls: 1,
			expectedResults: []faultResult{
				{
					localizedMessage: "fake",
					fault:            &types.SystemError{},
				},
			},
		},
		{
			name: "error is task.Error with localized message and fault with localizable messages",
			err: task.Error{
				LocalizedMethodFault: &types.LocalizedMethodFault{
					LocalizedMessage: "fake",
					Fault: &types.SystemError{
						RuntimeFault: types.RuntimeFault{
							MethodFault: types.MethodFault{
								FaultMessage: []types.LocalizableMessage{
									{
										Key:     "fake.id",
										Message: "fake",
									},
								},
							},
						},
					},
				},
			},
			callback:         &testHarness{},
			expectedNumCalls: 1,
			expectedResults: []faultResult{
				{
					fault:            &types.SystemError{},
					localizedMessage: "fake",
					localizableMessages: []types.LocalizableMessage{
						{
							Key:     "fake.id",
							Message: "fake",
						},
					},
				},
			},
		},
		{
			name: "error is task.Error with nested faults",
			err: task.Error{
				LocalizedMethodFault: &types.LocalizedMethodFault{
					LocalizedMessage: "fake1",
					Fault: &types.SystemError{
						RuntimeFault: types.RuntimeFault{
							MethodFault: types.MethodFault{
								FaultMessage: []types.LocalizableMessage{
									{
										Key:     "fake1.id",
										Message: "fake1",
									},
								},
								FaultCause: &types.LocalizedMethodFault{
									LocalizedMessage: "fake2",
									Fault: &types.NotSupported{
										RuntimeFault: types.RuntimeFault{
											MethodFault: types.MethodFault{
												FaultMessage: []types.LocalizableMessage{
													{
														Key:     "fake2.id",
														Message: "fake2",
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			callback:         &testHarness{},
			expectedNumCalls: 2,
			expectedResults: []faultResult{
				{
					fault:            &types.SystemError{},
					localizedMessage: "fake1",
					localizableMessages: []types.LocalizableMessage{
						{
							Key:     "fake1.id",
							Message: "fake1",
						},
					},
				},
				{
					fault:            &types.NotSupported{},
					localizedMessage: "fake2",
					localizableMessages: []types.LocalizableMessage{
						{
							Key:     "fake2.id",
							Message: "fake2",
						},
					},
				},
			},
		},
		{
			name: "error is task.Error with nested faults but returns after single fault",
			err: task.Error{
				LocalizedMethodFault: &types.LocalizedMethodFault{
					LocalizedMessage: "fake1",
					Fault: &types.SystemError{
						RuntimeFault: types.RuntimeFault{
							MethodFault: types.MethodFault{
								FaultMessage: []types.LocalizableMessage{
									{
										Key:     "fake1.id",
										Message: "fake1",
									},
								},
								FaultCause: &types.LocalizedMethodFault{
									LocalizedMessage: "fake2",
									Fault: &types.NotSupported{
										RuntimeFault: types.RuntimeFault{
											MethodFault: types.MethodFault{
												FaultMessage: []types.LocalizableMessage{
													{
														Key:     "fake2.id",
														Message: "fake2",
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			callback:         &testHarness{numCallsToReturnTrue: 1},
			expectedNumCalls: 1,
			expectedResults: []faultResult{
				{
					fault:            &types.SystemError{},
					localizedMessage: "fake1",
					localizableMessages: []types.LocalizableMessage{
						{
							Key:     "fake1.id",
							Message: "fake1",
						},
					},
				},
			},
		},
		{
			name: "err is soap fault",
			err: soap.WrapSoapFault(&soap.Fault{
				Detail: struct {
					Fault types.AnyType "xml:\",any,typeattr\""
				}{
					Fault: &types.SystemError{},
				},
			}),
			callback:         &testHarness{numCallsToReturnTrue: 1},
			expectedNumCalls: 1,
			expectedResults: []faultResult{
				{
					fault: &types.SystemError{},
				},
			},
		},
		{
			name:             "err is vim fault",
			err:              soap.WrapVimFault(&types.SystemError{}),
			callback:         &testHarness{numCallsToReturnTrue: 1},
			expectedNumCalls: 1,
			expectedResults: []faultResult{
				{
					fault: &types.SystemError{},
				},
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			var callback fault.OnFaultFn

			if !tc.callback.isNil {
				callback = tc.callback.Callback
			}

			if tc.expectedPanic != nil {
				assert.PanicsWithValue(
					t,
					tc.expectedPanic,
					func() {
						fault.In(tc.err, callback)
					})
			} else {
				fault.In(tc.err, callback)
			}

			assert.Equal(t, tc.expectedNumCalls, tc.callback.numCalls)
			assert.Equal(t, tc.expectedResults, tc.callback.results)
		})
	}
}

func TestIs(t *testing.T) {

	testCases := []struct {
		name          string
		err           any
		target        types.BaseMethodFault
		expectedOkay  bool
		expectedPanic any
	}{
		{
			name:         "err is nil",
			err:          nil,
			target:       &types.SystemError{},
			expectedOkay: false,
		},
		{
			name:         "target is nil",
			err:          &types.SystemError{},
			target:       nil,
			expectedOkay: false,
		},
		{
			name:         "target and error are nil",
			err:          nil,
			target:       nil,
			expectedOkay: true,
		},
		{
			name:         "err is not supported",
			err:          struct{}{},
			expectedOkay: false,
		},
		{
			name:         "err and target are same value type",
			err:          nilBaseMethodFault{},
			target:       nilBaseMethodFault{},
			expectedOkay: true,
		},
		{
			name:         "target implements IsFault",
			err:          isFaulter(true),
			target:       &types.SystemError{},
			expectedOkay: true,
		},
		{
			name:         "err is *LocalizedMethodFault with nil fault",
			err:          &types.LocalizedMethodFault{},
			target:       &types.SystemError{},
			expectedOkay: false,
		},
		{
			name: "err is *LocalizedMethodFault with SystemError fault",
			err: &types.LocalizedMethodFault{
				Fault: &types.SystemError{},
			},
			target:       &types.SystemError{},
			expectedOkay: true,
		},
		{
			name: "err is *LocalizedMethodFault with InvalidVmConfig fault",
			err: &types.LocalizedMethodFault{
				Fault: &types.SystemError{},
			},
			target:       &types.InvalidVmConfig{},
			expectedOkay: false,
		},
		{
			name: "err is task.Error with nested InvalidVmConfig fault",
			err: task.Error{
				LocalizedMethodFault: &types.LocalizedMethodFault{
					Fault: &types.SystemError{
						RuntimeFault: types.RuntimeFault{
							MethodFault: types.MethodFault{
								FaultCause: &types.LocalizedMethodFault{
									Fault: &types.InvalidVmConfig{},
								},
							},
						},
					},
				},
			},
			target:       &types.InvalidVmConfig{},
			expectedOkay: true,
		},
		{
			name: "err is *LocalizedMethodFault with nested InvalidVmConfig fault",
			err: &types.LocalizedMethodFault{
				Fault: &types.SystemError{
					RuntimeFault: types.RuntimeFault{
						MethodFault: types.MethodFault{
							FaultCause: &types.LocalizedMethodFault{
								Fault: &types.InvalidVmConfig{},
							},
						},
					},
				},
			},
			target:       &types.InvalidVmConfig{},
			expectedOkay: true,
		},
		{
			name: "err is *LocalizedMethodFault with nested nil fault",
			err: &types.LocalizedMethodFault{
				Fault: &types.SystemError{
					RuntimeFault: types.RuntimeFault{
						MethodFault: types.MethodFault{
							FaultCause: &types.LocalizedMethodFault{
								Fault: nilBaseMethodFault{},
							},
						},
					},
				},
			},
			target:       &types.InvalidVmConfig{},
			expectedOkay: false,
		},
		{
			name: "err is wrapped task.Error with nested InvalidVmConfig fault",
			err: fmt.Errorf("my error: %w", task.Error{
				LocalizedMethodFault: &types.LocalizedMethodFault{
					Fault: &types.SystemError{
						RuntimeFault: types.RuntimeFault{
							MethodFault: types.MethodFault{
								FaultCause: &types.LocalizedMethodFault{
									Fault: &types.InvalidVmConfig{},
								},
							},
						},
					},
				},
			}),
			target:       &types.InvalidVmConfig{},
			expectedOkay: true,
		},
		{
			name:         "err is wrapped nil error",
			err:          nilErrWrapper{},
			target:       &types.InvalidVmConfig{},
			expectedOkay: false,
		},
		{
			name: "err is wrapped error slice with expected value",
			err: unwrappableErrSlice{
				nil,
				task.Error{
					LocalizedMethodFault: &types.LocalizedMethodFault{
						Fault: &types.SystemError{
							RuntimeFault: types.RuntimeFault{
								MethodFault: types.MethodFault{
									FaultCause: &types.LocalizedMethodFault{
										Fault: &types.InvalidVmConfig{},
									},
								},
							},
						},
					},
				},
			},
			target:       &types.InvalidVmConfig{},
			expectedOkay: true,
		},
		{
			name: "err is wrapped error slice sans expected value",
			err: unwrappableErrSlice{
				nil,
				task.Error{
					LocalizedMethodFault: &types.LocalizedMethodFault{
						Fault: &types.SystemError{
							RuntimeFault: types.RuntimeFault{
								MethodFault: types.MethodFault{
									FaultCause: &types.LocalizedMethodFault{
										Fault: &types.SystemError{},
									},
								},
							},
						},
					},
				},
			},
			target:       &types.InvalidVmConfig{},
			expectedOkay: false,
		},
		{
			name: "err is soap fault",
			err: soap.WrapSoapFault(&soap.Fault{
				Detail: struct {
					Fault types.AnyType "xml:\",any,typeattr\""
				}{
					Fault: &types.SystemError{
						RuntimeFault: types.RuntimeFault{
							MethodFault: types.MethodFault{
								FaultCause: &types.LocalizedMethodFault{
									Fault: &types.InvalidVmConfig{},
								},
							},
						},
					},
				},
			}),
			target:       &types.InvalidVmConfig{},
			expectedOkay: true,
		},
		{
			name: "err is soap fault with nil vim fault",
			err: soap.WrapSoapFault(&soap.Fault{
				Detail: struct {
					Fault types.AnyType "xml:\",any,typeattr\""
				}{
					Fault: nil,
				},
			}),
			target:       &types.InvalidVmConfig{},
			expectedOkay: false,
		},
		{
			name: "err is vim fault",
			err: soap.WrapVimFault(&types.SystemError{
				RuntimeFault: types.RuntimeFault{
					MethodFault: types.MethodFault{
						FaultCause: &types.LocalizedMethodFault{
							Fault: &types.InvalidVmConfig{},
						},
					},
				},
			}),
			target:       &types.InvalidVmConfig{},
			expectedOkay: true,
		},
		{
			name: "err is vim fault with nil vim fault",
			err: soap.WrapVimFault(&types.SystemError{
				RuntimeFault: types.RuntimeFault{
					MethodFault: types.MethodFault{
						FaultCause: &types.LocalizedMethodFault{
							Fault: nil,
						},
					},
				},
			}),
			target:       &types.InvalidVmConfig{},
			expectedOkay: false,
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			var okay bool

			if tc.expectedPanic != nil {
				assert.PanicsWithValue(
					t,
					tc.expectedPanic,
					func() {
						okay = fault.Is(tc.err, tc.target)
					})
			} else {
				okay = fault.Is(tc.err, tc.target)
			}

			assert.Equal(t, tc.expectedOkay, okay)
		})
	}
}
