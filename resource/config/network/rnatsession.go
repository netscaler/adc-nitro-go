/*
* Copyright (c) 2021 Citrix Systems, Inc.
*
*   Licensed under the Apache License, Version 2.0 (the "License");
*   you may not use this file except in compliance with the License.
*   You may obtain a copy of the License at
*
*       http://www.apache.org/licenses/LICENSE-2.0
*
*  Unless required by applicable law or agreed to in writing, software
*   distributed under the License is distributed on an "AS IS" BASIS,
*   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
*   See the License for the specific language governing permissions and
*   limitations under the License.
 */

package network

/**
* Configuration for RNAT session resource.
 */
type Rnatsession struct {
	/**
	* IPv4 network address on whose traffic you want the Citrix ADC to do RNAT processing.
	 */
	Network string `json:"network,omitempty"`
	/**
	* Subnet mask associated with the network address.
	 */
	Netmask string `json:"netmask,omitempty"`
	/**
	* The NAT IP address defined for the RNAT entry.
	 */
	Natip string `json:"natip,omitempty"`
	/**
	* Name of any configured extended ACL whose action is ALLOW.
	 */
	Aclname string `json:"aclname,omitempty"`
}
