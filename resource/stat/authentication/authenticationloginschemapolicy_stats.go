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

package authentication


type Authenticationloginschemapolicystats struct {
	/**
	* The name of the LoginSchema policy for which statistics will be displayed.  If not given statistics are shown for all policies.
	*/
	Name string `json:"name,omitempty"`
	/**
	* Clear the statsistics / counters
	*/
	Clearstats string `json:"clearstats,omitempty"`
	Pipolicyhits uint64 `json:"pipolicyhits,omitempty"`
	/**
	* Number of hits on the policy
	*/
	Pipolicyhitsrate float64 `json:"pipolicyhitsrate,omitempty"`
	Pipolicyundefhits uint64 `json:"pipolicyundefhits,omitempty"`
	/**
	* Number of undef hits on the policy
	*/
	Pipolicyundefhitsrate float64 `json:"pipolicyundefhitsrate,omitempty"`

}
