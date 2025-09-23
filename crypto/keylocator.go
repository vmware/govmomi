// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package crypto

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"strconv"
	"strings"
)

// KeyLocatorType identifies different types of key locators
type KeyLocatorType int

const (
	KeyLocatorTypeInvalid KeyLocatorType = iota

	// Atomic types
	KeyLocatorTypeNull // the null key locator
	KeyLocatorTypeKey  // encodes a key directly

	// Indirect types
	KeyLocatorTypePassphrase // generates a key from a passphrase
	KeyLocatorTypeLDAP       // data in an LDAP server
	KeyLocatorTypeScript     // get key from external script
	KeyLocatorTypeRole       // data at a well known location
	KeyLocatorTypeFQID       // fully-qualified key ID

	// Compound types
	KeyLocatorTypeList // list of KLs (KeySafe possibly)
	KeyLocatorTypePair // A KL and associated encrypted data
)

// KeyLocatorClass represents the class of a key locator
type KeyLocatorClass int

const (
	KeyLocatorClassAtomic KeyLocatorClass = iota
	KeyLocatorClassIndirect
	KeyLocatorClassCompound
)

// KeyLocatorRole identifies well-known keys managed by this module
type KeyLocatorRole int

const (
	KeyLocatorRoleObfuscation KeyLocatorRole = iota
	KeyLocatorRoleAdminIdent
	KeyLocatorRoleAdminRecovery
	KeyLocatorRoleServer
)

// URL constants
const (
	VMwareURLPrefix             = "vmware:"
	VMwareURLDelim              = '/'
	VMwareURLCategoryKL         = "key"
	VMwareURLCompoundLeftDelim  = '('
	VMwareURLCompoundRightDelim = ')'
	VMwareURLCompoundDelim      = ','
	VMwareURLNullElem           = "<VMWARE-NULL>"
	VMwareURLEmptyStringElem    = "<VMWARE-EMPTYSTRING>"
)

// Type name mappings
var klTypeNames = map[KeyLocatorType]string{
	KeyLocatorTypeNull:       "null",
	KeyLocatorTypeKey:        "rawkey",
	KeyLocatorTypePassphrase: "phrase",
	KeyLocatorTypeLDAP:       "ldap",
	KeyLocatorTypeScript:     "script",
	KeyLocatorTypeRole:       "role",
	KeyLocatorTypeFQID:       "fqid",
	KeyLocatorTypeList:       "list",
	KeyLocatorTypePair:       "pair",
}

var klTypeNameReverse = make(map[string]KeyLocatorType)

// Role name mappings
var klRoleNames = map[KeyLocatorRole]string{
	KeyLocatorRoleObfuscation:   "obfuscation",
	KeyLocatorRoleAdminIdent:    "adminIdent",
	KeyLocatorRoleAdminRecovery: "adminRecovery",
	KeyLocatorRoleServer:        "server",
}

var klRoleNameReverse = make(map[string]KeyLocatorRole)

func init() {
	// Initialize reverse mappings
	for k, v := range klTypeNames {
		klTypeNameReverse[v] = k
	}
	for k, v := range klRoleNames {
		klRoleNameReverse[v] = k
	}
}

// KeyLocatorPassphraseParams holds parameters for passphrase-based key locators
type KeyLocatorPassphraseParams struct {
	KeyGenData     []byte
	KeyGenDataSize int
}

// KeyLocatorLDAPParams holds parameters for LDAP-based key locators
type KeyLocatorLDAPParams struct {
	Server string
	Domain string
	Port   uint32
	UseSSL bool
	Path   string
}

// KeyLocatorScriptParams holds parameters for script-based key locators
type KeyLocatorScriptParams struct {
	RelPath       string
	Signature     []byte
	SignatureSize int
}

// KeyLocatorFQIDParams holds parameters for FQID-based key locators
type KeyLocatorFQIDParams struct {
	KeyServerID string
	KeyID       string
}

// KeyLocatorIndirect represents the contents of an indirect key locator
type KeyLocatorIndirect struct {
	Type     KeyLocatorType
	UniqueID string

	// Union of type-specific parameters
	Phrase KeyLocatorPassphraseParams
	LDAP   KeyLocatorLDAPParams
	Script KeyLocatorScriptParams
	Role   KeyLocatorRole
	FQID   KeyLocatorFQIDParams
}

// KeyLocatorPair represents the contents of a pair key locator
type KeyLocatorPair struct {
	Locker         *KeyLocator
	CryptoMAC      string // Name of the MAC algorithm
	LockedData     []byte
	LockedDataSize int
}

// KeyLocator represents a key locator
type KeyLocator struct {
	Type KeyLocatorType

	// Union of type-specific data
	Key      []byte              // for atomic class (raw key data)
	Indirect *KeyLocatorIndirect // for indirect class
	Pair     *KeyLocatorPair     // for compound class, pair type
	List     []*KeyLocator       // for compound class, list type
}

// GetClass returns the class of the key locator type
func (kl *KeyLocator) GetClass() KeyLocatorClass {
	return KeyLocatorTypeToClass(kl.Type)
}

// KeyLocatorTypeToClass returns the class for a given type
func KeyLocatorTypeToClass(klType KeyLocatorType) KeyLocatorClass {
	switch klType {
	case KeyLocatorTypeNull, KeyLocatorTypeKey:
		return KeyLocatorClassAtomic
	case KeyLocatorTypePassphrase, KeyLocatorTypeLDAP, KeyLocatorTypeScript, KeyLocatorTypeRole, KeyLocatorTypeFQID:
		return KeyLocatorClassIndirect
	case KeyLocatorTypeList, KeyLocatorTypePair:
		return KeyLocatorClassCompound
	default:
		return KeyLocatorClassAtomic // default fallback
	}
}

// keyLocatorTypeToName returns the name for a key locator type.
func KeyLocatorTypeToName(klType KeyLocatorType) string {
	return klTypeNames[klType]
}

// KeyLocatorNameToType returns the type for a key locator name
func KeyLocatorNameToType(name string) (KeyLocatorType, bool) {
	klType, ok := klTypeNameReverse[strings.ToLower(name)]
	return klType, ok
}

// KeyLocatorRoleToName returns the name for a key locator role
func KeyLocatorRoleToName(role KeyLocatorRole) string {
	return klRoleNames[role]
}

// KeyLocatorNameToRole returns the role for a key locator name
func KeyLocatorNameToRole(name string) (KeyLocatorRole, bool) {
	role, ok := klRoleNameReverse[strings.ToLower(name)]
	return role, ok
}

// ImportKeyLocator constructs a KeyLocator from the specified opaque string,
// which was originally exported by KeyLocator.String.
func ImportKeyLocator(s string) (*KeyLocator, error) {
	var kl KeyLocator
	if err := kl.UnmarshalText([]byte(s)); err != nil {
		return nil, err
	}
	return &kl, nil
}

func (kl KeyLocator) String() string {
	data, err := kl.MarshalText()
	if err != nil {
		return err.Error()
	}
	return string(data)
}

func (kl KeyLocator) MarshalText() ([]byte, error) {
	var buf bytes.Buffer

	// Add URL prefix and category
	buf.WriteString(VMwareURLPrefix)
	buf.WriteString(VMwareURLCategoryKL)
	buf.WriteByte(VMwareURLDelim)

	// Export the key locator
	err := exportKeyLocator(&kl, &buf)
	if err != nil {
		return nil, fmt.Errorf("failed to export key locator: %v", err)
	}

	return buf.Bytes(), nil
}

func (kl *KeyLocator) UnmarshalText(text []byte) error {
	if len(text) == 0 {
		return fmt.Errorf("empty input string")
	}

	curS := string(text)

	// Consume prefix and category
	consumed, err := consumeToDelim(false, byte(VMwareURLDelim), &curS)
	if err != nil {
		return fmt.Errorf("failed to parse URL prefix: %v", err)
	}

	expectedPrefix := VMwareURLPrefix + VMwareURLCategoryKL
	if !strings.EqualFold(consumed, expectedPrefix) {
		return fmt.Errorf("invalid URL prefix, expected %s, got %s", expectedPrefix, consumed)
	}

	// Import the key locator.
	if err := importKeyLocator(false, &curS, kl); err != nil {
		return fmt.Errorf("failed to import key locator: %v", err)
	}

	return nil
}

// importKeyLocator imports a key locator from a URL string
func importKeyLocator(
	compound bool,
	s *string,
	kl *KeyLocator) error {

	// Consume type
	typeName, err := consumeToDelim(false, byte(VMwareURLDelim), s)
	if err != nil {
		return fmt.Errorf("failed to parse type: %v", err)
	}

	// Match type
	klType, ok := KeyLocatorNameToType(typeName)
	if !ok {
		return fmt.Errorf("unknown key locator type: %s", typeName)
	}

	// Create skeleton key locator
	kl.Type = klType

	// Continue importing based on class
	switch kl.GetClass() {
	case KeyLocatorClassAtomic:
		return importAtomicKL(compound, s, kl)
	case KeyLocatorClassIndirect:
		return importIndirectKL(compound, s, kl)
	case KeyLocatorClassCompound:
		return importCompoundKL(compound, s, kl)
	default:
		return fmt.Errorf("unknown key locator class for type %s", typeName)
	}
}

// importAtomicKL imports an atomic key locator
func importAtomicKL(compound bool, s *string, kl *KeyLocator) error {
	delim := byte(VMwareURLDelim)
	if compound {
		delim = byte(VMwareURLCompoundDelim)
	}

	if kl.Type == KeyLocatorTypeKey {
		// Consume exported key
		consumed, err := consumeToDelim(false, delim, s)
		if err != nil {
			return fmt.Errorf("failed to consume key data: %v", err)
		}

		// For now, store the consumed data as raw bytes
		// In a real implementation, this would be imported using the crypto library
		kl.Key = []byte(consumed)

	} else if kl.Type == KeyLocatorTypeNull {
		// Should consume empty string
		consumed, err := consumeToDelim(false, delim, s)
		if err != nil {
			return fmt.Errorf("failed to consume null element: %v", err)
		}

		if consumed != "" {
			return fmt.Errorf("null key locator should have empty content, got: %s", consumed)
		}
	} else {
		return fmt.Errorf("unsupported atomic key locator type: %d", kl.Type)
	}

	return nil
}

// importIndirectKL imports an indirect key locator
func importIndirectKL(compound bool, s *string, kl *KeyLocator) error {
	kl.Indirect = &KeyLocatorIndirect{Type: kl.Type}

	// Read unique ID (can be null for some types)
	// Allow null elements for cacheable types (all except passphrase)
	allowNull := kl.Type != KeyLocatorTypePassphrase
	uniqueID, err := consumeToDelim(allowNull, byte(VMwareURLDelim), s)
	if err != nil {
		return fmt.Errorf("failed to consume unique ID: %v", err)
	}
	kl.Indirect.UniqueID = uniqueID

	delim := byte(VMwareURLDelim)
	if compound {
		delim = byte(VMwareURLCompoundDelim)
	}

	// Read rest based on type
	switch kl.Type {
	case KeyLocatorTypePassphrase:
		keyGenData, err := consumeToDelim(false, delim, s)
		if err != nil {
			return fmt.Errorf("failed to consume key generation data: %v", err)
		}
		kl.Indirect.Phrase.KeyGenData = []byte(keyGenData)
		kl.Indirect.Phrase.KeyGenDataSize = len(keyGenData)

	case KeyLocatorTypeLDAP:
		// Read server
		server, err := consumeToDelim(true, byte(VMwareURLDelim), s)
		if err != nil {
			return fmt.Errorf("failed to consume LDAP server: %v", err)
		}
		kl.Indirect.LDAP.Server = server

		// Read domain
		domain, err := consumeToDelim(true, byte(VMwareURLDelim), s)
		if err != nil {
			return fmt.Errorf("failed to consume LDAP domain: %v", err)
		}
		kl.Indirect.LDAP.Domain = domain

		// Read port
		portStr, err := consumeToDelim(false, byte(VMwareURLDelim), s)
		if err != nil {
			return fmt.Errorf("failed to consume LDAP port: %v", err)
		}
		port, err := strconv.ParseUint(portStr, 10, 32)
		if err != nil {
			return fmt.Errorf("failed to parse LDAP port: %v", err)
		}
		kl.Indirect.LDAP.Port = uint32(port)

		// Read useSSL
		useSSLStr, err := consumeToDelim(false, byte(VMwareURLDelim), s)
		if err != nil {
			return fmt.Errorf("failed to consume LDAP useSSL: %v", err)
		}
		useSSL, err := parseBool(useSSLStr)
		if err != nil {
			return fmt.Errorf("failed to parse LDAP useSSL: %v", err)
		}
		kl.Indirect.LDAP.UseSSL = useSSL

		// Read path
		path, err := consumeToDelim(false, delim, s)
		if err != nil {
			return fmt.Errorf("failed to consume LDAP path: %v", err)
		}
		kl.Indirect.LDAP.Path = path

	case KeyLocatorTypeScript:
		// Read relative path
		relPath, err := consumeToDelim(false, byte(VMwareURLDelim), s)
		if err != nil {
			return fmt.Errorf("failed to consume script relative path: %v", err)
		}
		kl.Indirect.Script.RelPath = relPath

		// Read signature
		signature, err := consumeToDelim(false, delim, s)
		if err != nil {
			return fmt.Errorf("failed to consume script signature: %v", err)
		}
		kl.Indirect.Script.Signature = []byte(signature)
		kl.Indirect.Script.SignatureSize = len(signature)

	case KeyLocatorTypeRole:
		// Read role
		roleName, err := consumeToDelim(false, delim, s)
		if err != nil {
			return fmt.Errorf("failed to consume role name: %v", err)
		}
		role, ok := KeyLocatorNameToRole(roleName)
		if !ok {
			return fmt.Errorf("unknown role name: %s", roleName)
		}
		kl.Indirect.Role = role

	case KeyLocatorTypeFQID:
		// Read key server ID
		keyServerID, err := consumeToDelim(true, byte(VMwareURLDelim), s)
		if err != nil {
			return fmt.Errorf("failed to consume key server ID: %v", err)
		}
		kl.Indirect.FQID.KeyServerID = keyServerID

		// Read key ID
		keyID, err := consumeToDelim(false, delim, s)
		if err != nil {
			return fmt.Errorf("failed to consume key ID: %v", err)
		}
		kl.Indirect.FQID.KeyID = keyID

	default:
		return fmt.Errorf("unsupported indirect key locator type: %d", kl.Type)
	}

	return nil
}

// importCompoundKL imports a compound key locator
func importCompoundKL(compound bool, s *string, kl *KeyLocator) error {
	delim := byte(VMwareURLDelim)
	if compound {
		delim = byte(VMwareURLCompoundDelim)
	}

	// Extract content from parentheses
	content, err := consumeInNested(false, byte(VMwareURLCompoundLeftDelim), byte(VMwareURLCompoundRightDelim), delim, s)
	if err != nil {
		return fmt.Errorf("failed to consume nested content: %v", err)
	}

	curS := content

	switch kl.Type {
	case KeyLocatorTypeList:
		kl.List = make([]*KeyLocator, 0)

		// Parse all locators in the list
		for len(curS) > 0 {
			var newKL KeyLocator
			if err := importKeyLocator(true, &curS, &newKL); err != nil {
				return fmt.Errorf("failed to import list element: %v", err)
			}
			kl.List = append(kl.List, &newKL)
		}

	case KeyLocatorTypePair:
		kl.Pair = &KeyLocatorPair{}

		// Import locker
		var locker KeyLocator
		if err := importKeyLocator(true, &curS, &locker); err != nil {
			return fmt.Errorf("failed to import pair locker: %v", err)
		}
		kl.Pair.Locker = &locker

		// Import crypto MAC name
		cryptoMAC, err := consumeToDelim(false, byte(VMwareURLCompoundDelim), &curS)
		if err != nil {
			return fmt.Errorf("failed to consume crypto MAC: %v", err)
		}
		kl.Pair.CryptoMAC = cryptoMAC

		// Import locked data (base64 encoded)
		base64LockedData, err := consumeToDelim(false, 0, &curS)
		if err != nil {
			return fmt.Errorf("failed to consume locked data: %v", err)
		}

		lockedData, err := base64.StdEncoding.DecodeString(base64LockedData)
		if err != nil {
			return fmt.Errorf("failed to decode base64 locked data: %v", err)
		}

		kl.Pair.LockedData = lockedData
		kl.Pair.LockedDataSize = len(lockedData)

	default:
		return fmt.Errorf("unsupported compound key locator type: %d", kl.Type)
	}

	return nil
}

// exportKeyLocator exports a key locator to a string buffer
func exportKeyLocator(kl *KeyLocator, buf io.Writer) error {
	// Add type name
	typeName := KeyLocatorTypeToName(kl.Type)
	if typeName == "" {
		return fmt.Errorf("unknown key locator type: %d", kl.Type)
	}

	err := escapeAndAddString(typeName, byte(VMwareURLDelim), buf)
	if err != nil {
		return fmt.Errorf("failed to add type name: %v", err)
	}

	// Add content based on class
	switch kl.GetClass() {
	case KeyLocatorClassAtomic:
		return exportAtomicKL(kl, buf)
	case KeyLocatorClassIndirect:
		return exportIndirectKL(kl, buf)
	case KeyLocatorClassCompound:
		return exportCompoundKL(kl, buf)
	default:
		return fmt.Errorf("unknown key locator class")
	}
}

// exportAtomicKL exports an atomic key locator
func exportAtomicKL(kl *KeyLocator, buf io.Writer) error {
	if kl.Type == KeyLocatorTypeKey {
		// Export key data
		// In a real implementation, this would use the crypto library to export the key
		return escapeAndAdd(kl.Key, 0, buf)
	} else if kl.Type == KeyLocatorTypeNull {
		// Export empty string element
		return escapeAndAddString(VMwareURLEmptyStringElem, 0, buf)
	}

	return fmt.Errorf("unsupported atomic key locator type: %d", kl.Type)
}

// exportIndirectKL exports an indirect key locator
func exportIndirectKL(kl *KeyLocator, buf io.Writer) error {
	if kl.Indirect == nil {
		return fmt.Errorf("indirect key locator has nil indirect data")
	}

	// Add unique ID (can be null)
	err := escapeAndAddString(kl.Indirect.UniqueID, byte(VMwareURLDelim), buf)
	if err != nil {
		return fmt.Errorf("failed to add unique ID: %v", err)
	}

	// Add type-dependent contents
	switch kl.Type {
	case KeyLocatorTypePassphrase:
		return escapeAndAdd(kl.Indirect.Phrase.KeyGenData, 0, buf)

	case KeyLocatorTypeLDAP:
		// Add server
		err := escapeAndAddString(kl.Indirect.LDAP.Server, byte(VMwareURLDelim), buf)
		if err != nil {
			return err
		}

		// Add domain
		err = escapeAndAddString(kl.Indirect.LDAP.Domain, byte(VMwareURLDelim), buf)
		if err != nil {
			return err
		}

		// Add port
		portStr := strconv.FormatUint(uint64(kl.Indirect.LDAP.Port), 10)
		err = escapeAndAddString(portStr, byte(VMwareURLDelim), buf)
		if err != nil {
			return err
		}

		// Add useSSL
		useSSLStr := formatBool(kl.Indirect.LDAP.UseSSL)
		err = escapeAndAddString(useSSLStr, byte(VMwareURLDelim), buf)
		if err != nil {
			return err
		}

		// Add path
		return escapeAndAddString(kl.Indirect.LDAP.Path, 0, buf)

	case KeyLocatorTypeScript:
		// Add relative path
		err := escapeAndAddString(kl.Indirect.Script.RelPath, byte(VMwareURLDelim), buf)
		if err != nil {
			return err
		}

		// Add signature
		return escapeAndAdd(kl.Indirect.Script.Signature, 0, buf)

	case KeyLocatorTypeRole:
		// Add role name
		roleName := KeyLocatorRoleToName(kl.Indirect.Role)
		if roleName == "" {
			return fmt.Errorf("unknown role: %d", kl.Indirect.Role)
		}
		return escapeAndAddString(roleName, 0, buf)

	case KeyLocatorTypeFQID:
		// Add key server ID
		err := escapeAndAddString(kl.Indirect.FQID.KeyServerID, byte(VMwareURLDelim), buf)
		if err != nil {
			return err
		}

		// Add key ID
		return escapeAndAddString(kl.Indirect.FQID.KeyID, 0, buf)

	default:
		return fmt.Errorf("unsupported indirect key locator type: %d", kl.Type)
	}
}

// exportCompoundKL exports a compound key locator
func exportCompoundKL(kl *KeyLocator, buf io.Writer) error {
	// Add left delimiter
	buf.Write([]byte{VMwareURLCompoundLeftDelim})

	// Add compound contents
	switch kl.Type {
	case KeyLocatorTypeList:
		for i, elem := range kl.List {
			// Add separator between elements
			if i > 0 {
				buf.Write([]byte{VMwareURLCompoundDelim})
			}

			// Export this element
			err := exportKeyLocator(elem, buf)
			if err != nil {
				return fmt.Errorf("failed to export list element %d: %v", i, err)
			}
		}

	case KeyLocatorTypePair:
		if kl.Pair == nil {
			return fmt.Errorf("pair key locator has nil pair data")
		}

		// Export the locker
		err := exportKeyLocator(kl.Pair.Locker, buf)
		if err != nil {
			return fmt.Errorf("failed to export pair locker: %v", err)
		}

		// Add separator
		buf.Write([]byte{VMwareURLCompoundDelim})

		// Add crypto MAC name
		err = escapeAndAddString(kl.Pair.CryptoMAC, byte(VMwareURLCompoundDelim), buf)
		if err != nil {
			return fmt.Errorf("failed to add crypto MAC: %v", err)
		}

		// Add base64-encoded locked data
		base64Data := base64.StdEncoding.EncodeToString(kl.Pair.LockedData)
		err = escapeAndAddString(base64Data, 0, buf)
		if err != nil {
			return fmt.Errorf("failed to add locked data: %v", err)
		}

	default:
		return fmt.Errorf("unsupported compound key locator type: %d", kl.Type)
	}

	// Add right delimiter
	buf.Write([]byte{VMwareURLCompoundRightDelim})

	return nil
}
