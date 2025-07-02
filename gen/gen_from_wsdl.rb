# © Broadcom. All Rights Reserved.
# The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
# SPDX-License-Identifier: Apache-2.0

$:.unshift(File.expand_path(File.dirname(__FILE__)))

require "vim_wsdl"

require "yaml"

if !File.directory?(ARGV.first)
  raise "first argument not a directory"
end

target = ARGV[1]

# Load the vijson yaml to fetch vijson schemas.
vijson_path = File.join(File.expand_path("../sdk", __FILE__), target+".yaml")
vijson = nil
if File.exists?(vijson_path)
  vijson = YAML::load(File.open(vijson_path))["components"]["schemas"]
end

wsdl = WSDL.new(WSDL.read(target+".wsdl"), vijson)
wsdl.validate_assumptions!
wsdl.peek()

ifs = Peek.types.keys.select { |name| Peek.base?(name) }.size()
puts "%d classes, %d interfaces" % [Peek.types.size(), ifs]

File.open(File.join(ARGV.first, "types/enum.go"), "w") do |io|
  io.print WSDL.header("types")

  wsdl.
    types.
    sort_by { |x| x.name }.
    uniq { |x| x.name }.
    select { |t| t.is_enum? }.
    each { |e| e.dump(io); e.dump_init(io) }
end

File.open(File.join(ARGV.first, "types/types.go"), "w") do |io|
  io.print WSDL.header("types")
  if target != "vim"
    io.print <<EOF
import (
        "context"
        "github.com/vmware/govmomi/vim25/types"
)
EOF
  end

  wsdl.
    types.
    sort_by { |x| x.name }.
    uniq { |x| x.name }.
    select { |t| !t.is_enum? }.
    each { |e| e.dump(io); e.dump_init(io) }
end

File.open(File.join(ARGV.first, "types/if.go"), "w") do |io|
  io.print WSDL.header("types")

  Peek.dump_interfaces(io)
end

File.open(File.join(ARGV.first, "methods/methods.go"), "w") do |io|
  io.print WSDL.header("methods")
  if target == "vim"
    target += "25"
  end

  io.print <<EOF
import (
        "context"
        "github.com/vmware/govmomi/#{target}/types"
        "github.com/vmware/govmomi/vim25/soap"
)
EOF

  wsdl.
    operations.
    sort_by { |x| x.name }.
    each { |e| e.dump(io) }
end

exit(0)
