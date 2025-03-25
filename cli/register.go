// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package cli

import "os"

var commands = map[string]Command{}

var aliases = map[string]string{}

// hideUnreleased allows commands to be compiled into the govc binary without being registered by default.
// Unreleased commands are omitted from 'govc -h' help text and the generated govc/USAGE.md
// Setting the env var GOVC_SHOW_UNRELEASED=true enables any commands registered as unreleased.
var hideUnreleased = os.Getenv("GOVC_SHOW_UNRELEASED") != "true"

func ShowUnreleased() bool {
	return !hideUnreleased
}

func Register(name string, c Command, unreleased ...bool) {
	if len(unreleased) != 0 && unreleased[0] && hideUnreleased {
		return
	}
	commands[name] = c
}

func Alias(name string, alias string) {
	aliases[alias] = name
}

func Commands() map[string]Command {
	return commands
}
