# © Broadcom. All Rights Reserved.
# The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
# SPDX-License-Identifier: Apache-2.0

# Provide a simple shell extension to save and load govc
# environments to disk. No more running `export GOVC_ABC=xyz`
# in different shells over and over again. Loading the right
# govc environment variables is now only one short and
# autocompleted command away!
#
# Usage:
# * Source this file from your `~/.bashrc` or running shell.
# * Execute `govc-env` to print GOVC_* variables.
# * Execute `govc-env --save <name>` to save GOVC_* variables.
# * Execute `govc-env <name>` to load GOVC_* variables.
#

_govc_env_dir=$HOME/.govmomi/env
mkdir -p "${_govc_env_dir}"

_govc-env-complete() {
  local w="${COMP_WORDS[COMP_CWORD]}"
  local c="$(find ${_govc_env_dir} -mindepth 1 -maxdepth 1 -type f  | sort | xargs -r -L1 basename | xargs echo)"

  # Only allow completion if preceding argument if the function itself
  if [ "$3" == "govc-env" ]; then
    COMPREPLY=( $(compgen -W "${c}" -- "${w}") )
  fi
}

govc-env() {
  # Print current environment
  if [ -z "$1" ]; then
    for VAR in $(env | grep ^GOVC_ | cut -d= -f1); do
      echo "export ${VAR}='${!VAR}'"
    done

    return
  fi

  # Save current environment
  if [ "$1" == "--save" ]; then
    if [ ! -z "$2" ]; then
    	govc-env > ${_govc_env_dir}/$2 && echo govc env has been saved to ${_govc_env_dir}/$2
    else
	echo Usage: govc-env --save configname
    fi
    return
  fi

  # Load specified environment
  source ${_govc_env_dir}/$1
}

complete -F _govc-env-complete govc-env

