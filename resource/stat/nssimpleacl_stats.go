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

package stat

/**
* Statistics for simple ACL resource.
 */

type Nssimpleaclstats struct {
	/**
	* Clear the statsistics / counters
	 */
	Clearstats  string `json:"clearstats,omitempty"`
	Sacltothits int    `json:"sacltothits,omitempty"`
	/**
	* Packets matching a SimpleACL.
	 */
	Saclhitsrate  float64 `json:"saclhitsrate,omitempty"`
	Sacltotmisses int     `json:"sacltotmisses,omitempty"`
	/**
	* Packets not matching any SimpleACL.
	 */
	Saclmissesrate     float64 `json:"saclmissesrate,omitempty"`
	Saclscount         int     `json:"saclscount,omitempty"`
	Sacltotpktsallowed int     `json:"sacltotpktsallowed,omitempty"`
	/**
	* Total packets that matched a SimpleACL with action ALLOW and got consumed by Citrix ADC.
	 */
	Saclpktsallowedrate float64 `json:"saclpktsallowedrate,omitempty"`
	Sacltotpktsbridged  int     `json:"sacltotpktsbridged,omitempty"`
	/**
	* Total packets that matched a SimpleACL with action BRIDGE and got bridged by Citrix ADC.
	 */
	Saclpktsbridgedrate float64 `json:"saclpktsbridgedrate,omitempty"`
	Sacltotpktsdenied   int     `json:"sacltotpktsdenied,omitempty"`
	/**
	* Packets dropped because they match SimpleACL (Access Control List) with processing mode set to DENY.
	 */
	Saclpktsdeniedrate float64 `json:"saclpktsdeniedrate,omitempty"`
}
