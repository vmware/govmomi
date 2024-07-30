/*
Copyright (c) 2024-2024 VMware, Inc. All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package ovf

import (
	"math"
	"regexp"
	"strconv"
	"strings"
)

// These are used to validate the overall structure of the string being parsed and to differentiate tokens as we are
// processing them
var (
	blankRegexp         = regexp.MustCompile(`[[:blank:]]`)
	validIntegerRegexp  = regexp.MustCompile(`^([1-9]\d*)$`)
	validExponentRegexp = regexp.MustCompile(`^([1-9]\d*\^[1-9]\d*)$`)
	validByteUnitRegexp = regexp.MustCompile(`((^|kilo|kibi|mega|mebi|giga|gibi)byte(s?)$)`)
	validCapacityRegexp = regexp.MustCompile(`^[[:blank:]]*((([1-9]\d*\^[1-9]\d*)|([1-9]\d*))($|[[:blank:]]*\*[[:blank:]]*))*(([a-zA-Z]*(b|B)(y|Y)(t|T)(e|E)(s|S)?)($|([[:blank:]]*\*[[:blank:]]*(([1-9]\d*\^[1-9]\d*)|([1-9]\d*)))*))?$`)
)

// We only handle kilo, kibi, mega, mebi, giga, gibi prefixes due to size constraints of int64/uint64, but more
// importantly because prefixes larger than giga & gibi don't make sense for our use-case
var prefixMultipliers = map[string]int64{
	"byte":     1,                      // byte
	"kilobyte": 1 * 1000,               // byte * 1000
	"kibibyte": 1 * 1024,               // byte * 1024
	"megabyte": 1 * 1000 * 1000,        // byte * 1000 * 1000 = kilobyte * 1000
	"mebibyte": 1 * 1024 * 1024,        // byte * 1024 * 1024 = kibibyte * 1024
	"gigabyte": 1 * 1000 * 1000 * 1000, // byte * 1000 * 1000 * 1000 = kilobyte * 1000 * 1000 = megabyte * 1000
	"gibibyte": 1 * 1024 * 1024 * 1024, // byte * 1024 * 1024 * 1024 = kibibyte * 1024 * 1024 = mebibyte * 1024
}

// ParseCapacityAllocationUnits validates the string s is a valid programmatic unit with respect to the base unit 'byte'
// and parses the string to return the number of bytes s represents
func ParseCapacityAllocationUnits(s string) int64 {
	// Any strings which don't match against the regular expression are deemed invalid and zero is returned as the result
	if !validCapacityString(s) {
		return 0
	}
	var capacityBytes int64 = 1
	// Remove any whitespace in s and lowercase any alphabetic characters. Removal of whitespace is done after
	// validating against the regular expression because whitespace is valid for the most part, but is not valid
	// for exponential terms, e.g 2 ^ 10
	s = strings.ToLower(blankRegexp.ReplaceAllString(s, ""))
	// Split s on multiplication operator (*) so that we can just calculate integer multipliers. Each token will
	// then be either an integer, an exponential term to be converted to an integer, or a unit term to be converted
	// to an integer
	tokens := strings.Split(s, "*")

	// Loop through all tokens and convert any to integers if necessary and use to compute a running product
	for _, token := range tokens {
		switch {
		// "" should be treated identically to "byte". capacityBytes is already set to 1 so there is nothing to do
		case len(token) == 0:
			continue
		case validByteUnitString(token):
			capacityBytes = capacityBytes * prefixMultipliers[strings.TrimSuffix(token, "s")]
		case validExponentString(token):
			p := strings.Split(token, "^")
			b, _ := strconv.ParseInt(p[0], 10, 64)
			e, _ := strconv.ParseInt(p[1], 10, 64)
			capacityBytes = capacityBytes * int64(math.Pow(float64(b), float64(e)))
		case validIntegerString(token):
			n, _ := strconv.ParseInt(token, 10, 64)
			capacityBytes = capacityBytes * n
		default:
			// This should be unreachable. validCapacityString should have filtered out anything that cannot be
			// matched by the non-default cases
			capacityBytes = 0
		}
	}
	return capacityBytes
}

// validIntegerString matches the string s against the regular expression `^([1-9]\d*)$`; i.e. s should be of the form:
// any non-zero digit ([1-9]), followed by zero or more digits (\d*)
func validIntegerString(s string) bool {
	return validIntegerRegexp.MatchString(s)
}

// validExponentString matches the string s against the regular expression `^([1-9]\d*\^[1-9]\d*)$`; i.e. s should be of
// the form: any non-zero digit ([1-9]), followed by a caret (^) followed by any non-zero digit ([1-9]), followed by zero
// or more digits (\d*)
func validExponentString(s string) bool {
	return validExponentRegexp.MatchString(s)
}

// validByteUnitString matches the string s against a regular expression which only allows a unit of byte
// (optionally plural) with a valid decimal or binary prefix. See prefixMultipliers
func validByteUnitString(s string) bool {
	return validByteUnitRegexp.MatchString(s)
}

// validCapacityString matches the string s against the regular expression validCapacityRegexp and verifies that s is a
// valid programmatic unit with respect to the base unit 'byte'.
//
// Per the OVF schema defined in DSP8023: "If not specified default value is bytes. Value shall match a recognized value
// for the UNITS qualifier in DSP0004"
//
// DSP004 defines a programmatic unit as:
//
// programmatic-unit = [ sign ] *S unit-element *( *S unit-operator *S unit-element )
// sign = HYPHEN
// unit-element = number / [ prefix ] base-unit [ CARET exponent ]
// unit-operator = "*" / "/"
// number = floatingpoint-number / exponent-number
//
// ; An exponent shall be interpreted as a floating point number
// ; with the specified decimal base and exponent and a mantissa of 1
// exponent-number = base CARET exponent
// base = integer-number
// exponent = [ sign ] integer-number
//
// ; An integer shall be interpreted as a decimal integer number
// integer-number = NON-ZERO-DIGIT *( DIGIT )
//
// ; A float shall be interpreted as a decimal floating point number
// floatingpoint-number = 1*( DIGIT ) [ "." ] *( DIGIT )
//
// ; A prefix for a base unit (e.g. "kilo"). The numeric equivalents of
// ; these prefixes shall be interpreted as multiplication factors for the
// ; directly succeeding base unit. In other words, if a prefixed base
// ; unit is in the denominator of the overall programmatic unit, the
// ; numeric equivalent of that prefix is also in the denominator
// prefix = decimal-prefix / binary-prefix
//
// ; SI decimal prefixes as defined in ISO 1000
// decimal-prefix =
//
//	  "deca" ; 10^1
//	/ "hecto" ; 10^2
//	/ "kilo" ; 10^3
//	/ "mega" ; 10^6
//	/ "giga" ; 10^9
//	/ "tera" ; 10^12
//	/ "peta" ; 10^15
//	/ "exa" ; 10^18
//	/ "zetta" ; 10^21
//	/ "yotta" ; 10^24
//	/ "deci" ; 10^-1
//	/ "centi" ; 10^-2
//	/ "milli" ; 10^-3
//	/ "micro" ; 10^-6
//	/ "nano" ; 10^-9
//	/ "pico" ; 10^-12
//	/ "femto" ; 10^-15
//	/ "atto" ; 10^-18
//	/ "zepto" ; 10^-21
//	/ "yocto" ; 10^-24
//
// ; IEC binary prefixes as defined in IEC 80000-13
// binary-prefix =
//
//	  "kibi" ; 2^10
//	/ "mebi" ´ ; 2^20
//	/ "gibi" ; 2^30
//	/ "tebi" ; 2^40
//	/ "pebi" ; 2^50
//	/ "exbi" ; 2^60
//	/ "zebi" ; 2^70
//	/ "yobi" ; 2^80
//
// ; The name of a base unit
// base-unit = standard-unit / extension-unit
//
// ; The name of a standard base unit
// standard-unit = UNIT-IDENTIFIER
//
// ; The name of an extension base unit. If UNIT-IDENTIFIER begins with a
// ; prefix (see prefix ABNF rule), the meaning of that prefix shall not be
// ; changed by the extension base unit (examples of this for standard base
// ; units are "decibel" or "kilogram")
// ; extension-unit = org-id COLON UNIT-IDENTIFIER
//
// ; org-id shall include a copyrighted, trademarked, or otherwise unique
// ; name that is owned by the business entity that is defining the
// ; extension unit, or that is a registered ID assigned to the business
// ; entity by a recognized global authority. org-id shall not begin with
// ; a prefix (see prefix ABNF rule)
// org-id = UNIT-IDENTIFIER
// UNIT-IDENTIFIER = FIRST-UNIT-CHAR [ *( MID-UNIT-CHAR )
// LAST-UNIT-CHAR ]
// FIRST-UNIT-CHAR = UPPERALPHA / LOWERALPHA / UNDERSCORE
// LAST-UNIT-CHAR = FIRST-UNIT-CHAR / DIGIT / PARENS
// MID-UNIT-CHAR = LAST-UNIT-CHAR / HYPHEN / S
//
// DIGIT = ZERO / NON-ZERO-DIGIT
// ZERO = "0"
// NON-ZERO-DIGIT = "1"-"9"
// HYPHEN = U+002D ; "-"
// CARET = U+005E ; "^"
// COLON = U+003A ; ":"
// UPPERALPHA = U+0041-005A ; "A" ... "Z"
// LOWERALPHA = U+0061-007A ; "a" ... "z"
// UNDERSCORE = U+005F ; "_"
// PARENS = U+0028 / U+0029 ; "(", ")"
// S = U+0020 ; " "
//
// This definition is further restricted as such a broad definition by the above grammar does not make sense in the
// context of virtual disk capacity.
//
// We do not allow for negative values, division operations, floating-point numbers, negative exponents, nor the use of
// multiple units. Furthermore, we limit the allowed decimal and binary prefixes. This gives us:
//
// programmatic-unit =
//
//	   number
//		/ [prefix] base-unit
//		/ number *( *S unit-operator *S number) *S unit-operator *S [prefix] base-unit
//		/ [prefix] base-unit *( *S unit-operator *S number)
//		/ number *( *S unit-operator *S number) *S unit-operator *S [prefix] base-unit *( *S unit-operator *S number)
//
// unit-operator = "*"
// number = integer-number / exponent-number
// exponent-number = base CARET exponent
// base = integer-number
// exponent = integer-number
// integer-number = NON-ZERO-DIGIT *( DIGIT )
// prefix = decimal-prefix / binary-prefix
//
// decimal-prefix =
//
//	  "kilo" ; 10^3
//	/ "mega" ; 10^6
//	/ "giga" ; 10^9
//
// binary-prefix =
//
//	  "kibi" ; 2^10
//	/ "mebi" ; 2^20
//	/ "gibi" ; 2^30
//
// This function and the regular expression validCapacityRegexp are used to verify that the string we are parsing follows
// our above restricted grammar
func validCapacityString(s string) bool {
	// Integer followed by a trailing '*' is not handled by the regular expression and so is explicitly checked
	return validCapacityRegexp.MatchString(s) && !strings.HasSuffix(s, "*")
}
