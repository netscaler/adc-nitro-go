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

package ns

/**
* Statistics for simple ACL6 resource.
 */

type Nssimpleacl6stats struct {
	/**
	* Clear the statsistics / counters
	 */
	Clearstats   string `json:"clearstats,omitempty"`
	Sacl6tothits int    `json:"sacl6tothits,omitempty"`
	/**
	* Packets matching a SimpleACL6.
	 */
	Sacl6hitsrate  float64 `json:"sacl6hitsrate,omitempty"`
	Sacl6totmisses int     `json:"sacl6totmisses,omitempty"`
	/**
	* Packets not matching any SimpleACL6.
	 */
	Sacl6missesrate     float64 `json:"sacl6missesrate,omitempty"`
	Sacl6scount         int     `json:"sacl6scount,omitempty"`
	Sacl6totpktsallowed int     `json:"sacl6totpktsallowed,omitempty"`
	/**
	* Total packets that matched a SimpleACL6 with action ALLOW and got consumed by Citrix ADC.
	 */
	Sacl6pktsallowedrate float64 `json:"sacl6pktsallowedrate,omitempty"`
	Sacl6totpktsbridged  int     `json:"sacl6totpktsbridged,omitempty"`
	/**
	* Total packets that matched a SimpleACL6 with action BRIDGE and got bridged by Citrix ADC.
	 */
	Sacl6pktsbridgedrate float64 `json:"sacl6pktsbridgedrate,omitempty"`
	Sacl6totpktsdenied   int     `json:"sacl6totpktsdenied,omitempty"`
	/**
	* Packets dropped because they match SimpleACL6 with processing mode set to DENY.
	 */
	Sacl6pktsdeniedrate float64 `json:"sacl6pktsdeniedrate,omitempty"`
}
