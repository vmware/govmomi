# Generating the types

This document describes how a VMware engineer can generate the vim types from internal builds:

## Requirements

* Docker or Podman

## Inputs

Stage the vSphere SDK files in `./gen/sdk`:

1. The public files of `wsdl.zip`.
2. `eam-types.xsd` and `eam-messagetypes.xsd`.
3. `vi-json.yaml`, once per namespace, as `vim.yaml`, `pbm.yaml`, `sms.yaml`, `vslm.yaml` and `eam.yaml`.

Set `VIM_VERSION` in `gen.sh`.

## Steps

1. Open a terminal window.
2. Run `make generate-types`.

    ---

    :warning: **Please note**

    This can take a while the first time because it has to build the image used to generate the types.

    ---

## Generator rules

`overrides.yaml` holds rules that the generator applies on top of the WSDL output, keyed by Go package:

* `fields`: the Go type and the `xml`/`json` struct tags of a field. The `vim` struct tag still comes from the inputs.
* `optionality`: for a field whose optionality `vi-json.yaml` and the WSDL disagree on, `required` or `optional`
  overrides the WSDL's `minOccurs` and `wsdl` keeps it.
* `extra_fields`: struct fields that the inputs do not contain. `import` names a package that the declaration
  refers to.
* `extra_enum_values`: enum values that the inputs do not contain.

Generation fails if a rule no longer changes the output, if a field that `vi-json.yaml` and the WSDL disagree on has
no `optionality` rule, if an extra field is also generated in its struct or in an ancestor or descendant of it, or if
`overrides.yaml` names an unknown package or section.
