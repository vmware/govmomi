// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package crypto

import (
	"fmt"
	"io"
	"strconv"
	"strings"
)

// URLEscape escapes all non-alphanumeric characters in the string for URL
// inclusion.
func URLEscape(input []byte) string {
	if len(input) == 0 {
		return ""
	}

	var result strings.Builder
	result.Grow(len(input) * 3) // Worst case: every byte needs escaping

	for _, b := range input {
		if isAlnum(b) {
			result.WriteByte(b)
		} else {
			result.WriteString(fmt.Sprintf("%%%02x", b))
		}
	}

	return result.String()
}

// URLUnescape unescapes a URL-encoded string.
func URLUnescape(input string) ([]byte, error) {
	if len(input) == 0 {
		return nil, nil
	}

	result := make([]byte, 0, len(input))
	i := 0

	for i < len(input) {
		if input[i] == '%' {
			if i+2 >= len(input) {
				return nil, fmt.Errorf("invalid URL escape sequence")
			}

			hex := input[i+1 : i+3]
			if !isHexDigit(input[i+1]) || !isHexDigit(input[i+2]) {
				return nil, fmt.Errorf("invalid hex digits in escape sequence: %s", hex)
			}

			val, err := strconv.ParseUint(hex, 16, 8)
			if err != nil {
				return nil, fmt.Errorf("failed to parse hex value %s: %v", hex, err)
			}

			result = append(result, byte(val))
			i += 3
		} else {
			result = append(result, input[i])
			i++
		}
	}

	return result, nil
}

// isAlnum checks if a byte is alphanumeric
func isAlnum(b byte) bool {
	return (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z') || (b >= '0' && b <= '9')
}

// isHexDigit checks if a byte is a hex digit
func isHexDigit(b byte) bool {
	return (b >= '0' && b <= '9') || (b >= 'a' && b <= 'f') || (b >= 'A' && b <= 'F')
}

// escapeAndAdd escapes a string and adds it to the buffer with optional delimiter
func escapeAndAdd(s []byte, delim byte, buf io.Writer) error {
	var escaped string

	if s == nil {
		// Add NULL identifier
		escaped = VMwareURLNullElem
	} else if len(s) == 0 {
		// Add empty string identifier
		escaped = VMwareURLEmptyStringElem
	} else {
		// Escape the string
		escaped = URLEscape(s)
	}

	buf.Write([]byte(escaped))

	// Add delimiter if specified
	if delim != 0 {
		buf.Write([]byte{delim})
	}

	return nil
}

// escapeAndAddString is a convenience wrapper for string inputs
func escapeAndAddString(s string, delim byte, buf io.Writer) error {
	if s == "" {
		return escapeAndAdd(nil, delim, buf)
	}
	return escapeAndAdd([]byte(s), delim, buf)
}

// consumeToDelim advances through the string until the specified delimiter
func consumeToDelim(allowNullElems bool, delim byte, s *string) (string, error) {
	if len(*s) == 0 {
		return "", fmt.Errorf("empty string")
	}

	// Find delimiter or end of string
	pos := 0
	for pos < len(*s) && (*s)[pos] != delim {
		pos++
	}

	if pos == 0 {
		return "", fmt.Errorf("no characters consumed")
	}

	// Extract the consumed portion
	consumed := (*s)[:pos]

	// Advance the string pointer
	*s = (*s)[pos:]
	if len(*s) > 0 && (*s)[0] == delim {
		*s = (*s)[1:] // Skip delimiter
	}

	// Unescape the consumed string
	unescaped, err := URLUnescape(consumed)
	if err != nil {
		return "", err
	}

	consumedStr := string(unescaped)

	// Handle special cases
	if consumedStr == VMwareURLNullElem {
		if !allowNullElems {
			return "", fmt.Errorf("null elements not allowed")
		}
		return "", nil // Return empty string for null
	} else if consumedStr == VMwareURLEmptyStringElem {
		return "", nil
	}

	return consumedStr, nil
}

// consumeInNested consumes a nested delimited string (e.g., parentheses)
func consumeInNested(allowNullElems bool, left, right, delim byte, s *string) (string, error) {
	if len(*s) == 0 || (*s)[0] != left {
		return "", fmt.Errorf("string does not start with left delimiter")
	}

	level := 1
	pos := 1

	// Find matching right delimiter
	for pos < len(*s) && level > 0 {
		if (*s)[pos] == left {
			level++
		} else if (*s)[pos] == right {
			level--
		}
		if level > 0 {
			pos++
		}
	}

	if level > 0 {
		return "", fmt.Errorf("unmatched left delimiter")
	}

	// pos now points to the matching right delimiter
	if pos >= len(*s) {
		return "", fmt.Errorf("unexpected end of string")
	}

	// Extract content between delimiters (excluding the delimiters themselves)
	content := (*s)[1:pos]

	// Advance past the right delimiter
	pos++
	*s = (*s)[pos:]

	// Skip trailing delimiter if present
	if delim != 0 && len(*s) > 0 && (*s)[0] == delim {
		*s = (*s)[1:]
	}

	// Handle special cases
	if content == VMwareURLNullElem {
		if !allowNullElems {
			return "", fmt.Errorf("null elements not allowed")
		}
		return "", nil
	} else if content == VMwareURLEmptyStringElem {
		return "", nil
	}

	return content, nil
}

// parseBool parses a boolean string (case-insensitive)
func parseBool(s string) (bool, error) {
	switch strings.ToUpper(s) {
	case "TRUE":
		return true, nil
	case "FALSE":
		return false, nil
	default:
		return false, fmt.Errorf("invalid boolean value: %s", s)
	}
}

// formatBool formats a boolean as a string
func formatBool(b bool) string {
	if b {
		return "TRUE"
	}
	return "FALSE"
}
