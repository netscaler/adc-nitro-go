/*
Copyright 2021 Citrix Systems, Inc

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
package main

import (
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config"
	"github.com/citrix/adc-nitro-go/service"
)

func main() {
	client, _ := service.NewNitroClientFromEnv()
	route := config.Route{
		Network: "192.168.15.0",
		Netmask: "255.255.255.0",
		Gateway: "172.17.0.2",
	}
	// adding route
	_, err := client.AddResource(service.Route.Type(), "anyRoute", &route)
	if err == nil {
		client.SaveConfig()
	}
	// deleting route
	//var argsBundle = []string{"network:192.168.15.0", "netmask:255.255.255.0", "gateway:172.17.0.2"}
	var argsBundle = map[string]string{"network": "192.168.15.0", "netmask": "255.255.255.0", "gateway": "172.17.0.2"}
	//err2 := client.DeleteResourceWithArgs(service.Route.Type(), "", argsBundle)
	err2 := client.DeleteResourceWithArgsMap(service.Route.Type(), "", argsBundle)
	if err2 != nil {
		fmt.Println(err2)
	}
}
