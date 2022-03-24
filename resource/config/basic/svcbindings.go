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

package basic

/**
* Configuration for service bindings resource.
 */
type Svcbindings struct {
	/**
	* The name of the service.
	 */
	Servicename string `json:"servicename,omitempty"`

	//------- Read only Parameter ---------;

	Ipaddress   string `json:"ipaddress,omitempty"`
	Port        string `json:"port,omitempty"`
	Svrstate    string `json:"svrstate,omitempty"`
	Vservername string `json:"vservername,omitempty"`
}
