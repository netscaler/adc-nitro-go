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

package dps

/**
* Configuration for dps parameter resource.
*/
type Dpsparameter struct {
	/**
	* Customer ID of the Citrix Cloud customer
	*/
	Customerid string `json:"customerid,omitempty"`
	/**
	* Describes if the customer is connecting to Commerical/JapanCloud/Gov  Citrix Cloud customer
	*/
	Deployment string `json:"deployment,omitempty"`
	/**
	* Service URL of the Citrix Cloud customer
	*/
	Serviceurl string `json:"serviceurl,omitempty"`

	//------- Read only Parameter ---------;

	Builtin string `json:"builtin,omitempty"`
	Feature string `json:"feature,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
