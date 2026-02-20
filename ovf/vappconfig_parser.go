// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package ovf

import (
	"fmt"
	"net"
	"regexp"
	"strconv"
	"strings"
)

const maxVAppPropStringLen = 65535

var (
	strWithMinLenRx      = regexp.MustCompile(`^(?:string|password)\(\.\.(\d+)\)$`)
	strWithMaxLenRx      = regexp.MustCompile(`^(?:string|password)\((\d+)\.\.\)$`)
	strWithMinMaxLenRx   = regexp.MustCompile(`^(?:string|password)\((\d+)\.\.(\d+)\)$`)
	intWithMinMaxSizeRx  = regexp.MustCompile(`^int\(([+-]?\d+)\.\.([+-]?\d+)\)$`)
	realWithMinMaxSizeRx = regexp.MustCompile(`^real\(([+-]?(?:\d+(?:\.\d*)?|\.\d+))\.\.([+-]?(?:\d+(?:\.\d*)?|\.\d+))\)$`)
)

func parseVAppConfigValue(p Property, val string) (string, error) {
	val = strings.TrimSpace(val)
	parsedValue := ""

	switch p.Type {
	case "string", "password":
		//
		// A generic string. Max length 65535 (64k).
		//
		if l := len(val); l > maxVAppPropStringLen {
			return "", newParseLenErr(p, l, 0, maxVAppPropStringLen)
		}
		parsedValue = val

	case "boolean":
		//
		// A boolean. The value can be "True" or "False".
		//
		if ok, _ := strconv.ParseBool(val); ok {
			parsedValue = "True"
		} else {
			parsedValue = "False"
		}

	case "int":
		//
		// An integer value. Is semantically equivalent to
		// int(-2147483648..2147483647) e.g. signed int32.
		//
		if _, err := strconv.ParseInt(val, 10, 32); err != nil {
			return "", newParseErr(p, val)
		}
		parsedValue = val

	case "real":
		//
		// An IEEE 8-byte floating-point value, i.e. a float64.
		//
		if _, err := strconv.ParseFloat(val, 64); err != nil {
			return "", newParseErr(p, val)
		}
		parsedValue = val

	case "ip":
		//
		// An IPv4 address in dot-decimal notation or an IPv6 address in
		// colon-hexadecimal notation.
		//
		if v, _, err := parseIP(val); v == nil || err != nil {
			return "", newParseErr(p, val)
		}
		parsedValue = val

	case "ip:network":
		//
		// An IP address in dot-notation (IPv4) and colon-hexadecimal (IPv6)
		// on a particular network. The behavior of this type depends on the
		// ipAllocationPolicy.
		//

		// TODO(akutz) Figure out the correct parsing strategy.
		parsedValue = val

	case "expression":
		//
		// The default value specifies an expression that is calculated
		// by the system.
		//

		// TODO(akutz) Figure out the correct parsing strategy.
		parsedValue = val

	default:
		if m := strWithMinLenRx.FindStringSubmatch(p.Type); len(m) > 0 {
			//
			// A string with minimum character length x.
			//
			minLen, _ := strconv.Atoi(m[1])
			if l := len(val); l < minLen {
				return "", newParseLenErr(p, l, minLen, maxVAppPropStringLen)
			}
			parsedValue = val

		} else if m := strWithMaxLenRx.FindStringSubmatch(p.Type); len(m) > 0 {
			//
			// A string with maximum character length x.
			//
			maxLen, _ := strconv.Atoi(m[1])
			if l := len(val); l > maxLen {
				return "", newParseLenErr(p, l, 0, maxLen)
			}
			parsedValue = val

		} else if m := strWithMinMaxLenRx.FindStringSubmatch(p.Type); len(m) > 0 {
			//
			// A string with minimum character length x and maximum
			// character length y.
			//
			minLen, _ := strconv.Atoi(m[1])
			maxLen, _ := strconv.Atoi(m[2])

			if minLen > maxLen {
				return "", newParseMinMaxErr(p, int64(minLen), int64(maxLen))
			}

			if l := len(val); l < minLen || l > maxLen {
				return "", newParseLenErr(p, l, minLen, maxLen)
			}
			parsedValue = val

		} else if m := intWithMinMaxSizeRx.FindStringSubmatch(p.Type); len(m) > 0 {
			//
			// An integer value with a minimum size x and a maximum size y.
			// For example int(0..255) is a number between 0 and 255 both
			// included. This is also a way to specify that the number must
			// be a uint8. There is always a lower and lower bound. Max
			// number of digits is 100 including any sign. If exported to
			// OVF the value will be truncated to max of uint64 or int64.
			//
			minSize, _ := strconv.ParseInt(m[1], 10, 64)
			maxSize, _ := strconv.ParseInt(m[2], 10, 64)

			if minSize > maxSize {
				return "", newParseMinMaxErr(p, minSize, maxSize)
			}

			if minSize >= 0 {
				v, err := strconv.ParseUint(val, 10, 64)
				if err != nil {
					return "", newParseErr(p, val)
				}
				umin, umax := uint64(minSize), uint64(maxSize) //nolint:gosec
				if v < umin || v > umax {
					return "", newParseUintSizeErr(p, v, umin, umax)
				}
			} else {
				v, err := strconv.ParseInt(val, 10, 64)
				if err != nil {
					return "", newParseErr(p, val)
				}
				if v < minSize || v > maxSize {
					return "", newParseIntSizeErr(p, v, minSize, maxSize)
				}
			}

			parsedValue = val

		} else if m := realWithMinMaxSizeRx.FindStringSubmatch(p.Type); len(m) > 0 {
			//
			// An IEEE 8-byte floating-point value with a minimum size x and
			// a maximum size y. For example real(-1.5..1.5) must be a
			// number between -1.5 and 1.5. Because of the nature of float
			// some conversions can truncate the value. Real must be encoded
			// according to CIM.
			//
			minSize, _ := strconv.ParseFloat(m[1], 64)
			maxSize, _ := strconv.ParseFloat(m[2], 64)

			if minSize > maxSize {
				return "", newParseRealMinMaxErr(p, minSize, maxSize)
			}

			v, err := strconv.ParseFloat(val, 64)
			if err != nil {
				return "", newParseErr(p, val)
			}
			if v < minSize || v > maxSize {
				return "", newParseRealSizeErr(p, v, minSize, maxSize)
			}

			parsedValue = val

		} else {
			parsedValue = val
		}
	}

	return parsedValue, nil
}

// parseIP returns the parsed IP address and optional network. Please note, this
// function supports parsing IP addresses with or without the network length.
func parseIP(s string) (net.IP, *net.IPNet, error) {
	if strings.Contains(s, "/") {
		return net.ParseCIDR(s)
	}
	ip := net.ParseIP(s)
	return ip, nil, nil
}

func newParseErr(p Property, val string) error {
	if strings.HasPrefix(p.Type, "password") {
		return fmt.Errorf(
			"failed to parse prop=%q, type=%s",
			p.Key,
			p.Type)
	}
	return fmt.Errorf(
		"failed to parse prop=%q, type=%s, value=%v",
		p.Key,
		p.Type,
		val)
}

func newParseLenErr(p Property, actLen, minLen, maxLen int) error {
	return fmt.Errorf(
		"failed to parse prop=%q, type=%s due to length: "+
			"len=%d, min=%d, max=%d",
		p.Key,
		p.Type,
		actLen,
		minLen,
		maxLen)
}

func newParseIntSizeErr(p Property, val, minSize, maxSize int64) error {
	return fmt.Errorf(
		"failed to parse prop=%q, type=%s due to size: "+
			"val=%d, min=%d, max=%d",
		p.Key,
		p.Type,
		val,
		minSize,
		maxSize)
}

func newParseUintSizeErr(p Property, val, minSize, maxSize uint64) error {
	return fmt.Errorf(
		"failed to parse prop=%q, type=%s due to size: "+
			"val=%d, min=%d, max=%d",
		p.Key,
		p.Type,
		val,
		minSize,
		maxSize)
}

func newParseRealSizeErr(p Property, val, minSize, maxSize float64) error {
	return fmt.Errorf(
		"failed to parse prop=%q, type=%s due to size: "+
			"val=%f, min=%f, max=%f",
		p.Key,
		p.Type,
		val,
		minSize,
		maxSize)
}

func newParseMinMaxErr(p Property, minSize, maxSize int64) error {
	return fmt.Errorf(
		"failed to parse prop=%q, type=%s due to min=%d > max=%d",
		p.Key,
		p.Type,
		minSize,
		maxSize)
}

func newParseRealMinMaxErr(p Property, minSize, maxSize float64) error {
	return fmt.Errorf(
		"failed to parse prop=%q, type=%s due to min=%f > max=%f",
		p.Key,
		p.Type,
		minSize,
		maxSize)
}
