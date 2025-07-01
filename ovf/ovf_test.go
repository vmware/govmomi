// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package ovf

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"path"
	"strconv"
	"testing"
	"text/tabwriter"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"

	"github.com/vmware/govmomi/vim25/xml"
)

func testEnvelope(t *testing.T, fn string) *Envelope {
	f, err := os.Open(fn)
	if err != nil {
		t.Fatalf("error opening %s %s", fn, err)
	}
	defer f.Close()

	e, err := Unmarshal(f)
	if err != nil {
		t.Fatalf("error unmarshaling test file %s", err)
	}

	if e == nil {
		t.Fatal("empty envelope")
	}

	return e
}

func TestUnmarshal(t *testing.T) {
	e := testEnvelope(t, "fixtures/ttylinux.ovf")

	hw := e.VirtualSystem.VirtualHardware[0]
	if n := len(hw.Config); n != 3 {
		t.Errorf("Config=%d", n)
	}
	if n := len(hw.ExtraConfig); n != 2 {
		t.Errorf("ExtraConfig=%d", n)
	}
	for i, c := range append(hw.Config, hw.ExtraConfig...) {
		if *c.Required {
			t.Errorf("%d: Required=%t", i, *c.Required)
		}
		if c.Key == "" {
			t.Errorf("%d: key=''", i)
		}
		if c.Value == "" {
			t.Errorf("%d: value=''", i)
		}
	}
}

func TestDeploymentOptions(t *testing.T) {
	fn := os.Getenv("OVF_TEST_FILE")
	if fn == "" {
		t.Skip("OVF_TEST_FILE not specified")
	}
	e := testEnvelope(t, fn)

	if e.DeploymentOption == nil {
		t.Fatal("DeploymentOptionSection empty")
	}

	var b bytes.Buffer
	tw := tabwriter.NewWriter(&b, 2, 0, 2, ' ', 0)
	fmt.Fprintf(tw, "\n")
	for _, c := range e.DeploymentOption.Configuration {
		fmt.Fprintf(tw, "id=%s\t", c.ID)
		fmt.Fprintf(tw, "label=%s\t", c.Label)

		d := false
		if c.Default != nil {
			d = *c.Default
		}

		fmt.Fprintf(tw, "default=%t\t", d)
		fmt.Fprintf(tw, "\n")
	}
	tw.Flush()
	t.Log(b.String())
}

func TestVirtualSystemCollection(t *testing.T) {
	e := testEnvelope(t, "fixtures/virtualsystemcollection.ovf")

	assert.Nil(t, e.VirtualSystem)
	assert.NotNil(t, e.VirtualSystemCollection)
	assert.Len(t, e.VirtualSystemCollection.VirtualSystem, 2)
	assert.Equal(t, e.VirtualSystemCollection.VirtualSystem[0].ID, "storage server")
	assert.Equal(t, e.VirtualSystemCollection.VirtualSystem[1].ID, "web-server")
}

func TestMultipleDeploymentConfigs(t *testing.T) {
	e := testEnvelope(t, "fixtures/configspec.ovf")

	assert.NotNil(t, e.VirtualSystem)
	assert.Nil(t, e.VirtualSystemCollection)
	assert.NotNil(t, e.DeploymentOption)
	assert.Len(t, e.DeploymentOption.Configuration, 2)

	assert.NotNil(t, e.DeploymentOption.Configuration[0].Default)
	assert.True(t, *e.DeploymentOption.Configuration[0].Default)
	assert.Equal(t, "default", e.DeploymentOption.Configuration[0].ID)

	assert.Nil(t, e.DeploymentOption.Configuration[1].Default)
	assert.Equal(t, "frontend", e.DeploymentOption.Configuration[1].ID)

	assert.Len(t, e.VirtualSystem.VirtualHardware, 1)
	assert.Len(t, e.VirtualSystem.VirtualHardware[0].Item, 24)
	assert.Len(t, e.VirtualSystem.VirtualHardware[0].Item[2].Config, 1)
	assert.NotNil(t, e.VirtualSystem.VirtualHardware[0].Item[2].Config[0].Required)
	assert.False(t, *e.VirtualSystem.VirtualHardware[0].Item[2].Config[0].Required)
	assert.Equal(t, "slotInfo.pciSlotNumber", e.VirtualSystem.VirtualHardware[0].Item[2].Config[0].Key)
	assert.Equal(t, "128", e.VirtualSystem.VirtualHardware[0].Item[2].Config[0].Value)
}

func TestJSONEncoder(t *testing.T) {
	t.Parallel()

	testCases := []string{
		"fixtures/ttylinux.ovf",
		"fixtures/configspec.ovf",
		"fixtures/photon5.ovf",
		"fixtures/ubuntu24.10.ovf",
		"fixtures/virtualsystemcollection.ovf",
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(path.Base(tc), func(t *testing.T) {

			t.Parallel()

			// Unmarshal the OVF envelope from XML.
			decodedFromXML := testEnvelope(t, tc)

			// Marshal the OVF envelope to JSON.
			data, err := json.MarshalIndent(decodedFromXML, "", "  ")

			if assert.NoError(t, err) {

				if ok, _ := strconv.ParseBool(os.Getenv("DEBUG")); ok {
					t.Log(string(data))
				}

				// Unmarshal the OVF envelop from JSON.
				var decodedFromJSON Envelope
				if assert.NoError(t, json.Unmarshal(data, &decodedFromJSON)) {

					// Assert the OVF envelopes unmarshaled from XML and from
					// JSON are equal.
					if assert.True(
						t,
						cmp.Equal(*decodedFromXML, decodedFromJSON),
						cmp.Diff(*decodedFromXML, decodedFromJSON)) {

						// Take the OVF envelope that was unmarshaled from
						// JSON and marshal it *back* to XML.
						data, err := xml.Marshal(decodedFromJSON)
						if assert.NoError(t, err) {

							// Take the OVF envelope that was unmarshaled from
							// JSON, then back to XML, and unmarshal it from XML.
							var decodedFromXMLFromJSON Envelope
							if assert.NoError(
								t,
								xml.Unmarshal(data, &decodedFromXMLFromJSON)) {

								// Assert this envelope is equal to the
								// original.
								assert.True(
									t,
									cmp.Equal(*decodedFromXML, decodedFromXMLFromJSON),
									cmp.Diff(*decodedFromXML, decodedFromXMLFromJSON))
							}
						}
					}
				}
			}
		})
	}
}
