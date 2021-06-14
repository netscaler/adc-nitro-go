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
* Binding class showing the farmname that can be bound to wisite.
*/
type Wisitefarmnamebinding struct {
	/**
	* Name for the logical representation of a XenApp or XenDesktop farm to be bound to the Web Interface site. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters.
	*/
	Farmname string `json:"farmname,omitempty"`
	/**
	* Comma-separated IP addresses or host names of XenApp or XenDesktop servers providing XML services.
	*/
	Xmlserveraddresses string `json:"xmlserveraddresses,omitempty"`
	/**
	* Port number at which to contact the XML service.
	*/
	Xmlport uint32 `json:"xmlport,omitempty"`
	/**
	* Transport protocol to use for transferring data, related to the Web Interface site, between the Citrix ADC and the XML service.
	*/
	Transport string `json:"transport,omitempty"`
	/**
	* TCP port at which the XenApp or XenDesktop servers listenfor SSL Relay traffic from the Citrix ADC. This parameter is required if you have set SSL Relay as the transport protocol. 
		Web Interface uses root certificates when authenticating a server running SSL Relay. Make sure that all the servers running SSL Relay are configured to listen on the same port.
	*/
	Sslrelayport uint32 `json:"sslrelayport,omitempty"`
	/**
	* Use all the XML servers (load balancing mode) or only one server (failover mode).
	*/
	Loadbalance string `json:"loadbalance,omitempty"`
	/**
	* Active Directory groups that are permitted to enumerate resources from server farms. Including a setting for this parameter activates the user roaming feature. A maximum of 512 user groups can be specified for each farm defined with the Farm<n> parameter.  The groups must be comma separated.
	*/
	Groups string `json:"groups,omitempty"`
	/**
	* Binded farm is set as a recovery farm.
	*/
	Recoveryfarm string `json:"recoveryfarm,omitempty"`
	/**
	* Path to the Web Interface site.
	*/
	Sitepath string `json:"sitepath,omitempty"`


}