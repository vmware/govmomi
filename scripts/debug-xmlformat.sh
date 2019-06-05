#!/bin/bash -e

# pipe the most recent debug run to xmlformat
cd ${GOVC_DEBUG_PATH-"$HOME/.govmomi/debug"}
cd $(ls -t | head -1)

header() {
    printf "<!-- %s %s/%s\n%s\n-->\n" "$1" "$PWD" "$2" "$(tr -d '\r' < "$3")"
}

jqformat() {
  jq .
}

xmlformat() {
  xmlstarlet fo
}

for file in *.req.{xml,json}; do
    ext=${file##*.}
    base=$(basename "$file" ".req.$ext")
    header Request "$file" "${base}.req.headers"
    format=xmlformat
    if [ "$ext" = "json" ] ; then
      format=jqformat
    fi
    $format < "$file"
    file="${base}.res.$ext"
    if [ -e "$file" ] ; then
      header Response "$file" "${base}.res.headers"
      $format < "$file"
    fi
done
