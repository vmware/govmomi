#!/usr/bin/env python3

"""
Copyright (c) 2021 VMware, Inc. All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
"""

from os import linesep;
from sys import argv, stdin;
from xml.dom import minidom;

from pyVmomi import SoapAdapter;
from pyVmomi import vim;

objToXml = None

# If the program is invoked with a "-" command line argument then decode stdin.
# Otherwise create a new vim.vm.ConfigSpec.
if len(argv) > 1 and argv[1] == "-":
    objToXml = SoapAdapter.Deserialize(stdin.buffer)
else:
    objToXml = vim.vm.ConfigSpec()
    objToXml.name = "python-vm"
    objToXml.numCPUs = 2
    objToXml.memoryMB = 2048

# Serialize the object to an XML string.
xmlStr = SoapAdapter.SerializeToUnicode(objToXml)

# Prettify the XML string.
xmlDom = minidom.parseString(xmlStr)
xmlStr = xmlDom.toprettyxml(indent="  ")
xmlStr = linesep.join([s for s in xmlStr.splitlines() if s.strip()])

# Print the XML string.
print(xmlStr)