// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package ovf

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Tests for vApp property value parsing and validation. Covers all supported OVF property types
// and their constraints.

func TestParseVAppConfigValue(t *testing.T) {
	t.Run("string type", func(t *testing.T) {
		t.Run("valid value", func(t *testing.T) {
			got, err := parseVAppConfigValue(Property{Key: "string-prop", Type: "string"}, "test-value")
			require.NoError(t, err)
			assert.Equal(t, "test-value", got)
		})
		t.Run("value too long", func(t *testing.T) {
			_, err := parseVAppConfigValue(Property{Key: "string-prop", Type: "string"}, string(make([]byte, 65536)))
			assert.EqualError(t, err, fmt.Sprintf("failed to parse prop=%q, type=%s due to length: len=%d, min=%d, max=%d", "string-prop", "string", 65536, 0, 65535))
		})
		t.Run("empty value", func(t *testing.T) {
			got, err := parseVAppConfigValue(Property{Key: "string-prop", Type: "string"}, "")
			require.NoError(t, err)
			assert.Equal(t, "", got)
		})
	})

	t.Run("password type", func(t *testing.T) {
		t.Run("valid value", func(t *testing.T) {
			got, err := parseVAppConfigValue(Property{Key: "password-prop", Type: "password"}, "secret-password")
			require.NoError(t, err)
			assert.Equal(t, "secret-password", got)
		})
		t.Run("value too long does not leak password in error", func(t *testing.T) {
			longVal := string(make([]byte, 65536))
			_, err := parseVAppConfigValue(Property{Key: "password-prop", Type: "password"}, longVal)
			require.Error(t, err)
			assert.NotContains(t, err.Error(), longVal, "error must not contain the password value")
		})
	})

	t.Run("boolean type", func(t *testing.T) {
		t.Run("true value", func(t *testing.T) {
			got, err := parseVAppConfigValue(Property{Key: "bool-prop", Type: "boolean"}, "true")
			require.NoError(t, err)
			assert.Equal(t, "True", got)
		})
		t.Run("false value", func(t *testing.T) {
			got, err := parseVAppConfigValue(Property{Key: "bool-prop", Type: "boolean"}, "false")
			require.NoError(t, err)
			assert.Equal(t, "False", got)
		})
		t.Run("invalid value treated as false", func(t *testing.T) {
			got, err := parseVAppConfigValue(Property{Key: "bool-prop", Type: "boolean"}, "not-a-bool")
			require.NoError(t, err)
			assert.Equal(t, "False", got)
		})
	})

	t.Run("int type", func(t *testing.T) {
		t.Run("valid positive value", func(t *testing.T) {
			got, err := parseVAppConfigValue(Property{Key: "int-prop", Type: "int"}, "42")
			require.NoError(t, err)
			assert.Equal(t, "42", got)
		})
		t.Run("valid negative value", func(t *testing.T) {
			got, err := parseVAppConfigValue(Property{Key: "int-prop", Type: "int"}, "-42")
			require.NoError(t, err)
			assert.Equal(t, "-42", got)
		})
		t.Run("invalid value", func(t *testing.T) {
			_, err := parseVAppConfigValue(Property{Key: "int-prop", Type: "int"}, "not-a-number")
			assert.EqualError(t, err, fmt.Sprintf("failed to parse prop=%q, type=%s, value=%v", "int-prop", "int", "not-a-number"))
		})
		t.Run("value too large for int32", func(t *testing.T) {
			_, err := parseVAppConfigValue(Property{Key: "int-prop", Type: "int"}, "2147483648")
			assert.EqualError(t, err, fmt.Sprintf("failed to parse prop=%q, type=%s, value=%v", "int-prop", "int", "2147483648"))
		})
		t.Run("value too small for int32", func(t *testing.T) {
			_, err := parseVAppConfigValue(Property{Key: "int-prop", Type: "int"}, "-2147483649")
			assert.EqualError(t, err, fmt.Sprintf("failed to parse prop=%q, type=%s, value=%v", "int-prop", "int", "-2147483649"))
		})
	})

	t.Run("real type", func(t *testing.T) {
		t.Run("valid value", func(t *testing.T) {
			got, err := parseVAppConfigValue(Property{Key: "real-prop", Type: "real"}, "3.14159")
			require.NoError(t, err)
			assert.Equal(t, "3.14159", got)
		})
		t.Run("invalid value", func(t *testing.T) {
			_, err := parseVAppConfigValue(Property{Key: "real-prop", Type: "real"}, "not-a-float")
			assert.EqualError(t, err, fmt.Sprintf("failed to parse prop=%q, type=%s, value=%v", "real-prop", "real", "not-a-float"))
		})
	})

	t.Run("ip type", func(t *testing.T) {
		t.Run("valid IPv4", func(t *testing.T) {
			got, err := parseVAppConfigValue(Property{Key: "ip-prop", Type: "ip"}, "192.168.1.1")
			require.NoError(t, err)
			assert.Equal(t, "192.168.1.1", got)
		})
		t.Run("invalid IPv4", func(t *testing.T) {
			_, err := parseVAppConfigValue(Property{Key: "ip-prop", Type: "ip"}, "192.168.1.1.1")
			assert.EqualError(t, err, fmt.Sprintf("failed to parse prop=%q, type=%s, value=%v", "ip-prop", "ip", "192.168.1.1.1"))
		})
		t.Run("valid IPv6", func(t *testing.T) {
			got, err := parseVAppConfigValue(Property{Key: "ip-prop", Type: "ip"}, "2001:db8::1")
			require.NoError(t, err)
			assert.Equal(t, "2001:db8::1", got)
		})
		t.Run("invalid IPv6", func(t *testing.T) {
			_, err := parseVAppConfigValue(Property{Key: "ip-prop", Type: "ip"}, "2001:db8::1:::::::::")
			assert.EqualError(t, err, fmt.Sprintf("failed to parse prop=%q, type=%s, value=%v", "ip-prop", "ip", "2001:db8::1:::::::::"))
		})
		t.Run("invalid CIDR", func(t *testing.T) {
			_, err := parseVAppConfigValue(Property{Key: "ip-prop", Type: "ip"}, "192.168.1.1/33")
			assert.EqualError(t, err, fmt.Sprintf("failed to parse prop=%q, type=%s, value=%v", "ip-prop", "ip", "192.168.1.1/33"))
		})
	})

	t.Run("ip:network type", func(t *testing.T) {
		t.Run("valid CIDR passes through", func(t *testing.T) {
			got, err := parseVAppConfigValue(Property{Key: "ip-network-prop", Type: "ip:network"}, "192.168.1.0/24")
			require.NoError(t, err)
			assert.Equal(t, "192.168.1.0/24", got)
		})
	})

	t.Run("expression type", func(t *testing.T) {
		t.Run("valid expression passes through", func(t *testing.T) {
			got, err := parseVAppConfigValue(Property{Key: "expr-prop", Type: "expression"}, "some-expression")
			require.NoError(t, err)
			assert.Equal(t, "some-expression", got)
		})
	})

	t.Run("unknown type passes through", func(t *testing.T) {
		got, err := parseVAppConfigValue(Property{Key: "unknown-prop", Type: "unknown-type"}, "some-value")
		require.NoError(t, err)
		assert.Equal(t, "some-value", got)
	})

	t.Run("string with minimum length constraint", func(t *testing.T) {
		t.Run("valid - meets minimum", func(t *testing.T) {
			got, err := parseVAppConfigValue(Property{Key: "str-min-prop", Type: "string(..3)"}, "abc")
			require.NoError(t, err)
			assert.Equal(t, "abc", got)
		})
		t.Run("invalid - below minimum", func(t *testing.T) {
			_, err := parseVAppConfigValue(Property{Key: "str-min-prop", Type: "string(..3)"}, "ab")
			assert.EqualError(t, err, fmt.Sprintf("failed to parse prop=%q, type=%s due to length: len=%d, min=%d, max=%d", "str-min-prop", "string(..3)", 2, 3, 65535))
		})
		t.Run("valid - exact minimum", func(t *testing.T) {
			got, err := parseVAppConfigValue(Property{Key: "str-min-prop", Type: "string(..3)"}, "abc")
			require.NoError(t, err)
			assert.Equal(t, "abc", got)
		})
	})

	t.Run("string with maximum length constraint", func(t *testing.T) {
		t.Run("valid - within maximum", func(t *testing.T) {
			got, err := parseVAppConfigValue(Property{Key: "str-max-prop", Type: "string(3..)"}, "abc")
			require.NoError(t, err)
			assert.Equal(t, "abc", got)
		})
		t.Run("invalid - exceeds maximum", func(t *testing.T) {
			_, err := parseVAppConfigValue(Property{Key: "str-max-prop", Type: "string(3..)"}, "abcd")
			assert.EqualError(t, err, fmt.Sprintf("failed to parse prop=%q, type=%s due to length: len=%d, min=%d, max=%d", "str-max-prop", "string(3..)", 4, 0, 3))
		})
	})

	t.Run("string with min/max length constraint", func(t *testing.T) {
		t.Run("valid - within range", func(t *testing.T) {
			got, err := parseVAppConfigValue(Property{Key: "str-range-prop", Type: "string(2..4)"}, "abc")
			require.NoError(t, err)
			assert.Equal(t, "abc", got)
		})
		t.Run("invalid - too short", func(t *testing.T) {
			_, err := parseVAppConfigValue(Property{Key: "str-range-prop", Type: "string(2..4)"}, "a")
			assert.EqualError(t, err, fmt.Sprintf("failed to parse prop=%q, type=%s due to length: len=%d, min=%d, max=%d", "str-range-prop", "string(2..4)", 1, 2, 4))
		})
		t.Run("invalid - too long", func(t *testing.T) {
			_, err := parseVAppConfigValue(Property{Key: "str-range-prop", Type: "string(2..4)"}, "abcde")
			assert.EqualError(t, err, fmt.Sprintf("failed to parse prop=%q, type=%s due to length: len=%d, min=%d, max=%d", "str-range-prop", "string(2..4)", 5, 2, 4))
		})
		t.Run("invalid - min greater than max", func(t *testing.T) {
			_, err := parseVAppConfigValue(Property{Key: "str-range-prop", Type: "string(4..2)"}, "abc")
			assert.EqualError(t, err, fmt.Sprintf("failed to parse prop=%q, type=%s due to min=%d > max=%d", "str-range-prop", "string(4..2)", 4, 2))
		})
		t.Run("valid - exact length", func(t *testing.T) {
			got, err := parseVAppConfigValue(Property{Key: "str-exact-prop", Type: "string(3..3)"}, "abc")
			require.NoError(t, err)
			assert.Equal(t, "abc", got)
		})
	})

	t.Run("password with length constraints", func(t *testing.T) {
		t.Run("valid - minimum length", func(t *testing.T) {
			got, err := parseVAppConfigValue(Property{Key: "pass-min-prop", Type: "password(..3)"}, "abc")
			require.NoError(t, err)
			assert.Equal(t, "abc", got)
		})
		t.Run("invalid - below minimum does not leak password in error", func(t *testing.T) {
			secret := "ab"
			_, err := parseVAppConfigValue(Property{Key: "pass-min-prop", Type: "password(..3)"}, secret)
			require.Error(t, err)
			assert.NotContains(t, err.Error(), secret, "error must not contain the password value")
		})
		t.Run("valid - maximum length", func(t *testing.T) {
			got, err := parseVAppConfigValue(Property{Key: "pass-max-prop", Type: "password(3..)"}, "abc")
			require.NoError(t, err)
			assert.Equal(t, "abc", got)
		})
		t.Run("valid - min/max range", func(t *testing.T) {
			got, err := parseVAppConfigValue(Property{Key: "pass-range-prop", Type: "password(2..4)"}, "abc")
			require.NoError(t, err)
			assert.Equal(t, "abc", got)
		})
	})

	t.Run("int with range constraint", func(t *testing.T) {
		t.Run("valid positive", func(t *testing.T) {
			got, err := parseVAppConfigValue(Property{Key: "int-range-prop", Type: "int(0..255)"}, "100")
			require.NoError(t, err)
			assert.Equal(t, "100", got)
		})
		t.Run("valid - zero boundary", func(t *testing.T) {
			got, err := parseVAppConfigValue(Property{Key: "int-range-prop", Type: "int(0..10)"}, "0")
			require.NoError(t, err)
			assert.Equal(t, "0", got)
		})
		t.Run("valid - max boundary", func(t *testing.T) {
			got, err := parseVAppConfigValue(Property{Key: "int-range-prop", Type: "int(0..10)"}, "10")
			require.NoError(t, err)
			assert.Equal(t, "10", got)
		})
		t.Run("valid negative range", func(t *testing.T) {
			got, err := parseVAppConfigValue(Property{Key: "int-range-prop", Type: "int(-100..100)"}, "-50")
			require.NoError(t, err)
			assert.Equal(t, "-50", got)
		})
		t.Run("valid - negative min boundary", func(t *testing.T) {
			got, err := parseVAppConfigValue(Property{Key: "int-range-prop", Type: "int(-10..10)"}, "-10")
			require.NoError(t, err)
			assert.Equal(t, "-10", got)
		})
		t.Run("valid - negative max boundary", func(t *testing.T) {
			got, err := parseVAppConfigValue(Property{Key: "int-range-prop", Type: "int(-10..10)"}, "10")
			require.NoError(t, err)
			assert.Equal(t, "10", got)
		})
		t.Run("invalid - too high", func(t *testing.T) {
			_, err := parseVAppConfigValue(Property{Key: "int-range-prop", Type: "int(0..255)"}, "300")
			assert.EqualError(t, err, fmt.Sprintf("failed to parse prop=%q, type=%s due to size: val=%d, min=%d, max=%d", "int-range-prop", "int(0..255)", 300, 0, 255))
		})
		t.Run("invalid - negative value in unsigned range", func(t *testing.T) {
			_, err := parseVAppConfigValue(Property{Key: "int-range-prop", Type: "int(0..255)"}, "-10")
			assert.EqualError(t, err, fmt.Sprintf("failed to parse prop=%q, type=%s, value=%v", "int-range-prop", "int(0..255)", "-10"))
		})
		t.Run("invalid - min greater than max", func(t *testing.T) {
			_, err := parseVAppConfigValue(Property{Key: "int-range-prop", Type: "int(255..0)"}, "100")
			assert.EqualError(t, err, fmt.Sprintf("failed to parse prop=%q, type=%s due to min=%d > max=%d", "int-range-prop", "int(255..0)", 255, 0))
		})
		t.Run("invalid - non-numeric value", func(t *testing.T) {
			_, err := parseVAppConfigValue(Property{Key: "int-range-prop", Type: "int(0..255)"}, "not-a-number")
			assert.EqualError(t, err, fmt.Sprintf("failed to parse prop=%q, type=%s, value=%v", "int-range-prop", "int(0..255)", "not-a-number"))
		})
		t.Run("invalid - negative too low", func(t *testing.T) {
			_, err := parseVAppConfigValue(Property{Key: "int-range-prop", Type: "int(-10..10)"}, "-11")
			assert.EqualError(t, err, fmt.Sprintf("failed to parse prop=%q, type=%s due to size: val=%d, min=%d, max=%d", "int-range-prop", "int(-10..10)", -11, -10, 10))
		})
		t.Run("invalid - negative too high", func(t *testing.T) {
			_, err := parseVAppConfigValue(Property{Key: "int-range-prop", Type: "int(-10..10)"}, "11")
			assert.EqualError(t, err, fmt.Sprintf("failed to parse prop=%q, type=%s due to size: val=%d, min=%d, max=%d", "int-range-prop", "int(-10..10)", 11, -10, 10))
		})
		t.Run("invalid - int64 overflow", func(t *testing.T) {
			_, err := parseVAppConfigValue(Property{Key: "int-range-prop", Type: "int(-10..10)"}, "9223372036854775808")
			assert.EqualError(t, err, fmt.Sprintf("failed to parse prop=%q, type=%s, value=%v", "int-range-prop", "int(-10..10)", "9223372036854775808"))
		})
	})

	t.Run("real with range constraint", func(t *testing.T) {
		t.Run("valid - within range", func(t *testing.T) {
			got, err := parseVAppConfigValue(Property{Key: "real-range-prop", Type: "real(-2.0..2.0)"}, "1.5")
			require.NoError(t, err)
			assert.Equal(t, "1.5", got)
		})
		t.Run("valid - min boundary", func(t *testing.T) {
			got, err := parseVAppConfigValue(Property{Key: "real-range-prop", Type: "real(-2.0..2.0)"}, "-2.0")
			require.NoError(t, err)
			assert.Equal(t, "-2.0", got)
		})
		t.Run("valid - max boundary", func(t *testing.T) {
			got, err := parseVAppConfigValue(Property{Key: "real-range-prop", Type: "real(-2.0..2.0)"}, "2.0")
			require.NoError(t, err)
			assert.Equal(t, "2.0", got)
		})
		t.Run("invalid - too high", func(t *testing.T) {
			_, err := parseVAppConfigValue(Property{Key: "real-range-prop", Type: "real(-2.0..2.0)"}, "3.0")
			assert.EqualError(t, err, fmt.Sprintf("failed to parse prop=%q, type=%s due to size: val=%f, min=%f, max=%f", "real-range-prop", "real(-2.0..2.0)", 3.0, -2.0, 2.0))
		})
		t.Run("invalid - too low", func(t *testing.T) {
			_, err := parseVAppConfigValue(Property{Key: "real-range-prop", Type: "real(-2.0..2.0)"}, "-3.0")
			assert.EqualError(t, err, fmt.Sprintf("failed to parse prop=%q, type=%s due to size: val=%f, min=%f, max=%f", "real-range-prop", "real(-2.0..2.0)", -3.0, -2.0, 2.0))
		})
		t.Run("invalid - min greater than max", func(t *testing.T) {
			_, err := parseVAppConfigValue(Property{Key: "real-range-prop", Type: "real(2.0..-2.0)"}, "1.0")
			assert.EqualError(t, err, fmt.Sprintf("failed to parse prop=%q, type=%s due to min=%f > max=%f", "real-range-prop", "real(2.0..-2.0)", 2.0, -2.0))
		})
		t.Run("invalid - non-numeric value", func(t *testing.T) {
			_, err := parseVAppConfigValue(Property{Key: "real-range-prop", Type: "real(-2.0..2.0)"}, "not-a-float")
			assert.EqualError(t, err, fmt.Sprintf("failed to parse prop=%q, type=%s, value=%v", "real-range-prop", "real(-2.0..2.0)", "not-a-float"))
		})
	})
}

func TestParseVAppConfigValueWhitespaceTrimming(t *testing.T) {
	// For every type, leading/trailing whitespace must be stripped before
	// parsing. The returned value must also be the trimmed form.
	cases := []struct {
		name     string
		propType string
		input    string
		want     string
	}{
		// Basic types
		{"string", "string", "  hello  ", "hello"},
		{"password", "password", "  secret  ", "secret"},
		{"boolean true", "boolean", "  true  ", "True"},
		{"boolean false", "boolean", "  false  ", "False"},
		{"int", "int", "  42  ", "42"},
		{"real", "real", "  3.14  ", "3.14"},
		{"ip IPv4", "ip", "  192.168.1.1  ", "192.168.1.1"},
		{"ip IPv6", "ip", "  2001:db8::1  ", "2001:db8::1"},
		{"ip:network", "ip:network", "  192.168.1.0/24  ", "192.168.1.0/24"},
		{"expression", "expression", "  expr  ", "expr"},
		{"unknown type", "unknown-type", "  val  ", "val"},
		// Constrained string types
		{"string min length", "string(..5)", "  hello  ", "hello"},
		{"string max length", "string(3..)", "  abc  ", "abc"},
		{"string min/max length", "string(2..6)", "  abc  ", "abc"},
		// Constrained password types
		{"password min length", "password(..5)", "  hello  ", "hello"},
		{"password max length", "password(3..)", "  abc  ", "abc"},
		{"password min/max length", "password(2..6)", "  abc  ", "abc"},
		// Constrained int and real types
		{"int range", "int(0..255)", "  42  ", "42"},
		{"int negative range", "int(-100..100)", "  -50  ", "-50"},
		{"real range", "real(-2.0..2.0)", "  1.5  ", "1.5"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := parseVAppConfigValue(Property{Key: "prop", Type: tc.propType}, tc.input)
			require.NoError(t, err)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestParseVAppConfigValuePasswordNoLeak(t *testing.T) {
	// Verify that no error variant leaks password values.
	secret := "my-secret-password"
	types := []string{"password", "password(..3)", "password(20..)", "password(2..4)"}
	for _, typ := range types {
		t.Run(typ, func(t *testing.T) {
			_, err := parseVAppConfigValue(Property{Key: "p", Type: typ}, secret)
			if err != nil {
				assert.False(t, strings.Contains(err.Error(), secret),
					"error message must not contain the password value")
			}
		})
	}
}
