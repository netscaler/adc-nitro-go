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

package wi

/**
* Binding class showing the accessmethod that can be bound to wisite.
 */
type Wisiteaccessmethodbinding struct {
	/**
	* Secure access method to be applied to the IPv4 or network address of the client specified by the Client IP Address parameter.
		Depending on whether the Web Interface site is configured to use an HTTP or HTTPS virtual server or to use access gateway, you can send clients or access gateway the IP address, or the alternate address, of a XenApp or XenDesktop server. Or, you can send the IP address translated from a mapping entry, which defines mapping of an internal address and port to an external address and port.
	*/
	Accessmethod string `json:"accessmethod,omitempty"`
	/**
	* IPv4 or network address of the client for which you want to associate an access method.
	 */
	Clientipaddress string `json:"clientipaddress,omitempty"`
	/**
	* Subnet mask associated with the IPv4 or network address specified by the Client IP Address parameter.
	 */
	Clientnetmask string `json:"clientnetmask,omitempty"`
	/**
	* Path to the Web Interface site.
	 */
	Sitepath string `json:"sitepath,omitempty"`
}
