#!/bin/bash

set -e

header_dir=$(dirname $0)/headers

tmpfile=$(mktemp)
trap "rm -f ${tmpfile}" EXIT

git diff --name-status main | awk '{print $2}' | while read file; do
  years=( $(git log --format='%ai' $file | cut -d- -f1 | sort -u) )
  num_years=${#years[@]}

  if [ "${num_years}" == 0 ]; then
    export YEARS="$(date +%Y)"
  else
    yearA=${years[0]}
    yearB=${years[$((${num_years}-1))]}

    export YEARS="${yearA}-${yearB}"
  fi

  case "$file" in
    vim25/{json,xml}/*)
      # Ignore
      ;;
    *.go)
      sed -e "s/\${YEARS}/${YEARS}/" ${header_dir}/go.txt > ${tmpfile}
      last_header_line=$(grep -n '\*/' ${file} | head -1 | cut -d: -f1)
      tail -n +$((${last_header_line} + 1)) ${file} >> ${tmpfile}
      mv ${tmpfile} ${file}
      ;;
    *.rb)
      sed -e "s/\${YEARS}/${YEARS}/" ${header_dir}/rb.txt > ${tmpfile}
      last_header_line=$(grep -n '^$' ${file} | head -1 | cut -d: -f1)
      tail -n +$((${last_header_line})) ${file} >> ${tmpfile}
      mv ${tmpfile} ${file}
      ;;
    *)
      echo "Unhandled file: $file"
      ;;
  esac
done
