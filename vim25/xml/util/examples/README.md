# Sharing VIM objects between Golang and Python using XML

This directory includes several examples of how to share VIM objects between multiple languages, such as Golang and Python, by marshaling and unmarshaling those objects to and from XML.

# Requirements

* Golang 1.16+
* Python 3.9+
* PyVmomi (use `pip3 install --user pyvmomi` to make sure you have the latest version)

# Marshal a ConfigSpec to XML

The following examples demonstrate how to marshal a simple `ConfigSpec` object to XML to stdout in Golang and Python:

## Golang

```bash
$ go run main.go
<obj xmlns:vim25="urn:vim25" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:type="vim25:VirtualMachineConfigSpec">
  <name>go-vm</name>
  <numCPUs>2</numCPUs>
  <memoryMB>2048</memoryMB>
</obj>
```

## Python

```bash
$ python3 main.py
<?xml version="1.0" ?>
<object xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns="urn:vim25" xsi:type="VirtualMachineConfigSpec">
  <name>python-vm</name>
  <numCPUs>2</numCPUs>
  <memoryMB>2048</memoryMB>
</object>
```

# Unmarshal XML to a ConfigSpec

The following examples demonstrate how to unmarshal XML to a `ConfigSpec` and back to stdout in Golang and Python:

## Golang

Please note the output will reflect the name from the object created in Python, `python-vm`, which indicates the resource was:

1. Marshaled to XML in Python
1. Unmarshaled to XML in Go
1. Marshaled _back_ to XML in Go

```bash
$ python3 main.py | go run main.go -
<obj xmlns:vim25="urn:vim25" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:type="vim25:VirtualMachineConfigSpec">
  <name>python-vm</name>
  <numCPUs>2</numCPUs>
  <memoryMB>2048</memoryMB>
</obj>
```

## Python

Please note the output will reflect the name from the object created in Go, `go-vm`, which indicates the resource was:

1. Marshaled to XML in Go
1. Unmarshaled to XML in Python
1. Marshaled _back_ to XML in Python

```bash
$ go run main.go | python3 main.py -
<?xml version="1.0" ?>
<object xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns="urn:vim25" xsi:type="VirtualMachineConfigSpec">
  <name>go-vm</name>
  <numCPUs>2</numCPUs>
  <memoryMB>2048</memoryMB>
</object>
```
