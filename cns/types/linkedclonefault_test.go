// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"strings"
	"testing"

	"github.com/vmware/govmomi/vim25/types"
	"github.com/vmware/govmomi/vim25/xml"
)

// A CnsVolumeOperationResult fault as returned by CNS when a linked clone create-from-snapshot request is rejected
// by the FCD layer: the typed fault carries the original NotSupported fault as its cause.
const linkedCloneNotSupportedFaultXML = `
<fault xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
  <fault xsi:type="CnsLinkedCloneNotSupportedFault">
    <faultCause>
      <fault xsi:type="NotSupported">
        <faultMessage>
          <key>com.vmware.vim.vslm.vStorageObjectManager.EncryptionOpOnLCNotSupported</key>
          <message>Crypto ops are not supported when creating Linked Clone FCD.</message>
        </faultMessage>
      </fault>
      <localizedMessage>The operation is not supported on the object.</localizedMessage>
    </faultCause>
    <reason>Task failed with fault NotSupported</reason>
  </fault>
  <localizedMessage>Task failed with fault NotSupported</localizedMessage>
</fault>`

// TestCnsLinkedCloneNotSupportedFaultDecode verifies that the fault is registered, so that a client decodes it to its
// own type, with the NotSupported fault as its cause, rather than to a nil Fault.
func TestCnsLinkedCloneNotSupportedFaultDecode(t *testing.T) {
	dec := xml.NewDecoder(strings.NewReader(linkedCloneNotSupportedFaultXML))
	dec.TypeFunc = types.TypeFunc()

	var lmf types.LocalizedMethodFault
	if err := dec.Decode(&lmf); err != nil {
		t.Fatal(err)
	}

	fault, ok := lmf.Fault.(*CnsLinkedCloneNotSupportedFault)
	if !ok {
		t.Fatalf("fault is %T, want *CnsLinkedCloneNotSupportedFault", lmf.Fault)
	}
	if fault.Reason != "Task failed with fault NotSupported" {
		t.Errorf("reason = %q", fault.Reason)
	}
	// The typed fault is a CnsFault, so callers that only know CnsFault still see it as one.
	var _ BaseCnsFault = fault

	cause := fault.FaultCause
	if cause == nil {
		t.Fatal("fault cause is nil")
	}
	notSupported, ok := cause.Fault.(*types.NotSupported)
	if !ok {
		t.Fatalf("fault cause is %T, want *NotSupported", cause.Fault)
	}
	if len(notSupported.FaultMessage) != 1 ||
		notSupported.FaultMessage[0].Message != "Crypto ops are not supported when creating Linked Clone FCD." {
		t.Errorf("fault messages = %+v", notSupported.FaultMessage)
	}
}
