// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package crypto_test

import (
	"fmt"
	"log"

	"github.com/vmware/govmomi/crypto"
)

// ExampleImportKeyLocator demonstrates importing and parsing a KeyLocator URL.
func ExampleImportKeyLocator() {
	testURL := `vmware:key/list/(pair/(fqid/<VMWARE-NULL>/local/ASEAAgEAIexBYS7MTFOGo6XyS0PZfQEIAAwAEAAgAAQAQUVTLTI1NgCmT8mZIAuAGqLaFVU3pBlaT7%2fzDJm3%2fy%2f05n9y9%2bxv1aVSfBY9e6rBrhvKIkB2G%2fsvF7L%2bLGpoojr136%2bghgEA,HMAC%2dSHA%2d256,kIc7z%2fJowrpzVUAand6fC4ixT5BY6KwNTbAPFQErmxFRkmhJNOTp1VyQFnkn5kLgvKpt7KJKlm%2fvLkO6YxkVe61EMdtdsR2nL9DWMsDWov9syEh%2ftVED%2fzCct1fFpUaqSa29J%2fFk9%2bD22HiA0%2flumBPwt9M5aW0HB9T9lEMxNEVpSOBPmOW63DzLzAq1EC7%2fIWuCimTL%2b15%2be4uwDvxEYI5RDofZ2fm9oyM9MLHDTYPo%2fsFo8GU1LK%2frLsQcj20XijOe%2bfLnDlbJcH1nCmyoO8tweHwDs%2fmwhbpQudvXbGVM3jboiXoPj9rki%2boGeE8clTcBUyRxHE6n56MuZ6HmH1GHt9tBLyAHvk4oj2wNGGc%3d))`

	// Import the KeyLocator
	kl, err := crypto.ImportKeyLocator(testURL)
	if err != nil {
		log.Fatalf("Failed to import KeyLocator: %v", err)
	}

	// Analyze the structure
	fmt.Printf("Root type: %s\n", crypto.KeyLocatorTypeToName(kl.Type))
	fmt.Printf("List contains %d element(s)\n", len(kl.List))

	// Get the first (and only) element - it's a pair
	pair := kl.List[0]
	fmt.Printf("First element type: %s\n", crypto.KeyLocatorTypeToName(pair.Type))

	// Examine the locker
	locker := pair.Pair.Locker
	fmt.Printf("Locker type: %s\n", crypto.KeyLocatorTypeToName(locker.Type))
	fmt.Printf("Key server: %s\n", locker.Indirect.FQID.KeyServerID)
	fmt.Printf("Key ID length: %d characters\n", len(locker.Indirect.FQID.KeyID))

	// Examine the pair details
	fmt.Printf("Crypto MAC: %s\n", pair.Pair.CryptoMAC)
	fmt.Printf("Locked data size: %d bytes\n", len(pair.Pair.LockedData))

	// Output:
	// Root type: list
	// List contains 1 element(s)
	// First element type: pair
	// Locker type: fqid
	// Key server: local
	// Key ID length: 140 characters
	// Crypto MAC: HMAC-SHA-256
	// Locked data size: 272 bytes
}

// ExampleKeyLocator_String demonstrates exporting a KeyLocator back to URL
// format.
func ExampleKeyLocator_String() {
	// First import a KeyLocator
	testURL := `vmware:key/fqid/unique-123/server1/key-abc`

	kl, err := crypto.ImportKeyLocator(testURL)
	if err != nil {
		log.Fatalf("Failed to import: %v", err)
	}

	// Now export it back
	exported := kl.String()

	fmt.Printf("Original:  %s\n", testURL)
	fmt.Printf("Exported:  %s\n", exported)

	// Output:
	// Original:  vmware:key/fqid/unique-123/server1/key-abc
	// Exported:  vmware:key/fqid/unique%2d123/server1/key%2dabc
}

// Example demonstrates creating a simple KeyLocator programmatically
func ExampleKeyLocator_manual() {
	// Create an FQID KeyLocator manually
	kl := &crypto.KeyLocator{
		Type: crypto.KeyLocatorTypeFQID,
		Indirect: &crypto.KeyLocatorIndirect{
			Type:     crypto.KeyLocatorTypeFQID,
			UniqueID: "my-unique-id",
			FQID: crypto.KeyLocatorFQIDParams{
				KeyServerID: "production-server",
				KeyID:       "encryption-key-001",
			},
		},
	}

	// Export it to URL format
	url := kl.String()
	fmt.Printf("Generated URL: %s\n", url)

	// Import it back to verify
	imported, err := crypto.ImportKeyLocator(string(url))
	if err != nil {
		log.Fatalf("Failed to import back: %v", err)
	}

	fmt.Printf("Round-trip successful: %t\n",
		imported.Indirect.FQID.KeyServerID == kl.Indirect.FQID.KeyServerID &&
			imported.Indirect.FQID.KeyID == kl.Indirect.FQID.KeyID)

	// Output:
	// Generated URL: vmware:key/fqid/my%2dunique%2did/production%2dserver/encryption%2dkey%2d001
	// Round-trip successful: true
}
