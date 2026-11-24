// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package soap

import (
	"crypto/x509"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/vmware/govmomi/vim25/types"
)

type soapFaultError struct {
	fault *Fault
}

func (s soapFaultError) Error() string {
	msg := s.fault.String

	if msg == "" {
		if s.fault.Detail.Fault == nil {
			msg = "unknown fault"
		} else {
			msg = reflect.TypeOf(s.fault.Detail.Fault).Name()
		}
	}

	return fmt.Sprintf("%s: %s", s.fault.Code, msg)
}

func (s soapFaultError) MarshalJSON() ([]byte, error) {
	out := struct {
		Fault *Fault
	}{
		Fault: s.fault,
	}
	return json.Marshal(out)
}

func (s soapFaultError) Fault() types.BaseMethodFault {
	if s.fault != nil {
		fault := s.fault.Detail.Fault
		if fault == nil {
			return nil
		}
		if f, ok := fault.(types.BaseMethodFault); ok {
			return f
		}
		if val := reflect.ValueOf(fault); val.Kind() != reflect.Pointer {
			ptrVal := reflect.New(val.Type())
			ptrVal.Elem().Set(val)
			if f, ok := ptrVal.Interface().(types.BaseMethodFault); ok {
				return f
			}
		}
	}
	return nil
}

type vimFaultError struct {
	fault types.BaseMethodFault
}

func (v vimFaultError) Error() string {
	typ := reflect.TypeOf(v.fault)
	for typ.Kind() == reflect.Pointer {
		typ = typ.Elem()
	}

	return typ.Name()
}

func (v vimFaultError) Fault() types.BaseMethodFault {
	return v.fault
}

func WrapSoapFault(f *Fault) error {
	return soapFaultError{f}
}

// IsSoapFault returns true if err is, or wraps, a SOAP fault.
func IsSoapFault(err error) bool {
	var s soapFaultError
	return errors.As(err, &s)
}

// ToSoapFault returns the SOAP fault in err's chain, or nil if there is none.
func ToSoapFault(err error) *Fault {
	var s soapFaultError
	if errors.As(err, &s) {
		return s.fault
	}
	return nil
}

func WrapVimFault(v types.BaseMethodFault) error {
	return vimFaultError{v}
}

// IsVimFault returns true if err is, or wraps, a vim fault.
func IsVimFault(err error) bool {
	var v vimFaultError
	return errors.As(err, &v)
}

// ToVimFault returns the vim fault in err's chain, or nil if there is none.
func ToVimFault(err error) types.BaseMethodFault {
	var v vimFaultError
	if errors.As(err, &v) {
		return v.fault
	}
	return nil
}

func IsCertificateUntrusted(err error) bool {
	// golang 1.20 introduce a new type to wrap 509 errors. So instead of
	// casting the type, now we check the error chain contains the
	// x509 error or not.
	if errors.As(err, &x509.UnknownAuthorityError{}) {
		return true
	}

	if errors.As(err, &x509.HostnameError{}) {
		return true
	}

	// The err variable may not be a special type of x509 or HTTP
	// error that can be validated by a type assertion. The err variable is
	// in fact be an *errors.errorString.

	msgs := []string{
		"certificate is not trusted",
		"certificate signed by unknown authority",
	}

	for _, msg := range msgs {
		if strings.HasSuffix(err.Error(), msg) {
			return true
		}
	}

	return false
}
