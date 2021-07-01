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

package authorization

/**
* Statistics for authorization policy label resource.
*/

type Authorizationpolicylabelstats struct {
	/**
	* Name of the authorization policy label.
	*/
	Labelname string `json:"labelname,omitempty"`
	/**
	* Clear the statsistics / counters
	*/
	Clearstats string `json:"clearstats,omitempty"`
	Pipolicylabelhits int `json:"pipolicylabelhits,omitempty"`
	/**
	* Number of times policy label was invoked. 
	*/
	Pipolicylabelhitsrate float64 `json:"pipolicylabelhitsrate,omitempty"`

}
