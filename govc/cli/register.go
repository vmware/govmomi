/*
Copyright (c) 2014-2015 VMware, Inc. All Rights Reserved.

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

package cli

import "os"

var commands = map[string]Command{}

var aliases = map[string]string{}

// hideUnreleased allows commands to be compiled into the govc binary without being registered by default.
// Unreleased commands are omitted from 'govc -h' help text and the generated govc/USAGE.md
// Setting the env var GOVC_SHOW_UNRELEASED=true enables any commands registered as unreleased.
var hideUnreleased = os.Getenv("GOVC_SHOW_UNRELEASED") != "true"

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
