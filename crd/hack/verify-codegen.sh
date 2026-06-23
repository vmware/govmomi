#!/usr/bin/env bash

# © Broadcom. All Rights Reserved.
# The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
# SPDX-License-Identifier: Apache-2.0

# Check if generated files are up to date

set -o errexit
set -o nounset
set -o pipefail

# Change directories to the parent directory of the one in which this
# script is located.
cd "$(dirname "${BASH_SOURCE[0]}")/.."

make generate || exit_code="${?}"
if [ "${exit_code:-0}" -ne 0 ]; then

  # Please note the following heredoc uses leading tabs to allow
  # the contents to be indented with the if/fi statement. For
  # more information on indenting heredocs, please see the section
  # entitled "Multi-line message, with tabs suppressed" from The
  # Linux Documentation Project (TLDP) at
  # https://tldp.org/LDP/abs/html/here-docs.html.
  cat <<-EOF
	Please investigate why the target is failing by running the
	following command locally from the root of the project:

	    make generate

	Thank you!
	EOF

  exit "${exit_code}"
fi

if git diff --exit-code; then
  printf '\nCongratulations! Generated assets are up-to-date!\n'
else
  exit_code="${?}"

  # Please note the following heredoc uses leading tabs to allow
  # the contents to be indented with the if/fi statement. For
  # more information on indenting heredocs, please see the section
  # entitled "Multi-line message, with tabs suppressed" from The
  # Linux Documentation Project (TLDP) at
  # https://tldp.org/LDP/abs/html/here-docs.html.
  cat <<-EOF
	Please update generated assets before opening a pull request or
	pushing new changes. To generate the assets, please run the
	following command from the root of the project:

	    make generate

	Thank you!
	EOF

  exit "${exit_code}"
fi

make generate-go-conversions || exit_code="${?}"
if [ "${exit_code:-0}" -ne 0 ]; then

  # Please note the following heredoc uses leading tabs to allow
  # the contents to be indented with the if/fi statement. For
  # more information on indenting heredocs, please see the section
  # entitled "Multi-line message, with tabs suppressed" from The
  # Linux Documentation Project (TLDP) at
  # https://tldp.org/LDP/abs/html/here-docs.html.
  cat <<-EOF
	Please investigate why the target is failing by running the
	following command locally from the root of the project:

	    make generate-go-conversions

	Thank you!
	EOF

  exit "${exit_code}"
fi

if git diff --exit-code; then
  printf '\nCongratulations! Generated conversions are up-to-date!\n'
else
  exit_code="${?}"

  # Please note the following heredoc uses leading tabs to allow
  # the contents to be indented with the if/fi statement. For
  # more information on indenting heredocs, please see the section
  # entitled "Multi-line message, with tabs suppressed" from The
  # Linux Documentation Project (TLDP) at
  # https://tldp.org/LDP/abs/html/here-docs.html.
  cat <<-EOF
	Please update generated conversions before opening a pull request or
	pushing new changes. To generate the conversions, please run the
	following command from the root of the project:

	    make generate-go-conversions

	Thank you!
	EOF

  exit "${exit_code}"
fi
