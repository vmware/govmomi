/*
Copyright (c) 2022 VMware, Inc. All Rights Reserved.

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

package staged

// Severity defines the severity of the issues fixed in the update.
type Severity string

const (
	// CRITICAL Vulnerabilities that can be exploited by an unauthenticated attacker
	// from the Internet or those that break the guest/host Operating System isolation.
	CRITICAL Severity = "CRITICAL"

	// IMPORTANT Vulnerabilities that are not rated critical but whose exploitation
	// results in the complete compromise of confidentiality and/or integrity of user data and/or processing resources through user assistance or by authenticated attackers.
	IMPORTANT Severity = "IMPORTANT"

	// MODERATE Vulnerabilities where the ability to exploit is mitigated to a
	// significant degree by configuration or difficulty of exploitation,
	//but in certain deployment scenarios could still lead to the compromise of
	//confidentiality, integrity, or availability of user data and/or processing resources.
	MODERATE Severity = "MODERATE"

	// LOW_SEVERITY All other issues that have a security impact. Vulnerabilities
	// where exploitation is believed to be extremely difficult, or where successful exploitation would have minimal impact
	LOW_SEVERITY Severity = "LOW"
)

// QuestionType defines representation of field fields in GUI or CLI
type QuestionType string

const (
	// PLAIN_TEXT  plain text answer
	PLAIN_TEXT QuestionType = "PLAIN_TEXT"

	// BOOLEAN : Yes/No,On/Off,Checkbox answer
	BOOLEAN QuestionType = "BOOLEAN"

	// PASSWORD : Password (masked) answer
	PASSWORD QuestionType = "PASSWORD"
)

// Priority defines the update installation priority recommendations.
type Priority string

const (
	//HIGH  Install ASAP
	HIGH Priority = "HIGH"

	// MEDIUM Install at the earliest convenience
	MEDIUM Priority = "MEDIUM"

	//LOW Install at your discretion
	LOW Priority = "LOW"
)

// UpdateType Updatetype defines update type
type UpdateType string

const (
	// SECURITY Fixes vulnerabilities, doesn’t change functionality
	SECURITY UpdateType = "SECURITY"

	// FIX Fixes bugs/vulnerabilities, doesn’t change functionality
	FIX UpdateType = "FIX"

	//UPDATE Changes product functionality
	UPDATE UpdateType = "UPDATE"

	// UPGRADE Introduces new features, significantly changes product functionality
	UPGRADE UpdateType = "UPGRADE"
)
