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
* Statistics for Policy Based Routing(PBR) entry resource.
*/

type Nspbrstats struct {
	/**
	* Name of the PBR whose statistics you want the Citrix ADC to display.
	*/
	Name string `json:"name,omitempty"`
	/**
	* Clear the statsistics / counters
	*/
	Clearstats string `json:"clearstats,omitempty"`
	Pbrtotpktsallowed uint64 `json:"pbrtotpktsallowed,omitempty"`
	/**
	* Total packets that matched the PBR (Policy-Based Routes) with action ALLOW 
	*/
	Pbrpktsallowedrate int32 `json:"pbrpktsallowedrate,omitempty"`
	Pbrtotpktsdenied uint64 `json:"pbrtotpktsdenied,omitempty"`
	/**
	* Total packets that matched the PBR with action DENY 
	*/
	Pbrpktsdeniedrate int32 `json:"pbrpktsdeniedrate,omitempty"`
	Pbrtothits uint64 `json:"pbrtothits,omitempty"`
	/**
	* Total packets that matched one of the configured PBR
	*/
	Pbrhitsrate int32 `json:"pbrhitsrate,omitempty"`
	Pbrtotmisses uint64 `json:"pbrtotmisses,omitempty"`
	/**
	* Total packets that did not match any PBR
	*/
	Pbrmissesrate int32 `json:"pbrmissesrate,omitempty"`
	Pbrtotnulldrop uint64 `json:"pbrtotnulldrop,omitempty"`
	/**
	* Total packets that are dropped due to null nexthop
	*/
	Pbrnulldroprate int32 `json:"pbrnulldroprate,omitempty"`
	Pbrperhits uint64 `json:"pbrperhits,omitempty"`
	/**
	* Number of times the pbr was hit
	*/
	Pbrperhitsrate int32 `json:"pbrperhitsrate,omitempty"`

}
