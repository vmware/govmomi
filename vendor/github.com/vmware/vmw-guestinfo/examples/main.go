// Copyright 2016 VMware, Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"log"

	"github.com/vmware/vmw-guestinfo/rpcvmx"
	"github.com/vmware/vmw-guestinfo/vmcheck"
)

func main() {
	isVM, err := vmcheck.IsVirtualWorld()
	if err != nil {
		log.Fatalf("error: %s", err.Error())
	}

	if !isVM {
		log.Fatal("not in a virtual world... :(")
	}

	config := rpcvmx.NewConfig()

	fmt.Println(config.SetString("foo", "bar"))
	fmt.Println(config.String("foo", "foo"))

	fmt.Println(config.SetInt("foo", 3))
	fmt.Println(config.Int("foo", 0))

	fmt.Println(config.SetBool("foo", false))
	fmt.Println(config.Bool("foo", true))

}
