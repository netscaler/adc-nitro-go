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

package config

/**
* Configuration for WF site resource.
 */
type Wfsite struct {
	/**
	* Name of the WebFront site being created on the Citrix ADC.
	 */
	Sitename string `json:"sitename,omitempty"`
	/**
	* FQDN or IP of Windows StoreFront server where the store is configured.
	 */
	Storefronturl string `json:"storefronturl,omitempty"`
	/**
	* Name of the store present in StoreFront
	 */
	Storename string `json:"storename,omitempty"`
	/**
	* Specifies whether or not to use HTML5 receiver for launching apps for the WF site.
	 */
	Html5receiver string `json:"html5receiver,omitempty"`
	/**
	* Specifies whether of not to use workspace control for the WF site.
	 */
	Workspacecontrol string `json:"workspacecontrol,omitempty"`
	/**
	* Specifies whether or not to display the accounts selection screen during First Time Use of Receiver
	 */
	Displayroamingaccounts string `json:"displayroamingaccounts,omitempty"`
	/**
	* The value to be sent in the X-Frame-Options header. WARNING: Setting this option to ALLOW could leave the end user vulnerable to Click Jacking attacks.
	 */
	Xframeoptions string `json:"xframeoptions,omitempty"`

	//------- Read only Parameter ---------;

	State string `json:"state,omitempty"`
}
