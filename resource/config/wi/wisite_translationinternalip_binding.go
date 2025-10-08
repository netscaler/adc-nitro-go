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
* Binding class showing the translationinternalip that can be bound to wisite.
*/
type Wisitetranslationinternalipbinding struct {
	/**
	* IP address of the server for which you want to associate an external IP address. (Clients access the server through the associated external address and port.)
	*/
	Translationinternalip string `json:"translationinternalip,omitempty"`
	/**
	* Type of access to the XenApp or XenDesktop server. 
		Available settings function as follows:
		* User Device - Clients can use the translated address of the mapping entry to connect to the XenApp or XenDesktop server.
		* Gateway - Access Gateway can use the translated address of the mapping entry to connect to the XenApp or XenDesktop server.
		* User Device and Gateway - Both clients and Access Gateway can use the translated address of the mapping entry to connect to the XenApp or XenDesktop server.
	*/
	Accesstype string `json:"accesstype,omitempty"`
	/**
	* Port number of the server for which you want to associate an external port.  (Clients access the server through the associated external address and port.)
	*/
	Translationinternalport *int `json:"translationinternalport,omitempty"`
	/**
	* External IP address associated with server's IP address.
	*/
	Translationexternalip string `json:"translationexternalip,omitempty"`
	/**
	* External port number associated with the server's port number.
	*/
	Translationexternalport *int `json:"translationexternalport,omitempty"`
	/**
	* Path to the Web Interface site.
	*/
	Sitepath string `json:"sitepath,omitempty"`


}