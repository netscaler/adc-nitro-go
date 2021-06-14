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
* Statistics for PBR6 entry resource.
*/

type Nspbr6stats struct {
	/**
	* Name of the PBR6 whose statistics you want the Citrix ADC to display.
	*/
	Name string `json:"name,omitempty"`
	/**
	* Clear the statsistics / counters
	*/
	Clearstats string `json:"clearstats,omitempty"`
	Pbr6totpktsallowed uint64 `json:"pbr6totpktsallowed,omitempty"`
	/**
	* Total packets that matched the PBR6 with action ALLOW 
	*/
	Pbr6pktsallowedrate int32 `json:"pbr6pktsallowedrate,omitempty"`
	Pbr6totpktsdenied uint64 `json:"pbr6totpktsdenied,omitempty"`
	/**
	* Total packets that matched PBR6 with action DENY 
	*/
	Pbr6pktsdeniedrate int32 `json:"pbr6pktsdeniedrate,omitempty"`
	Pbr6tothits uint64 `json:"pbr6tothits,omitempty"`
	/**
	* Total packets that matched one of the configured PBR6
	*/
	Pbr6hitsrate int32 `json:"pbr6hitsrate,omitempty"`
	Pbr6totmisses uint64 `json:"pbr6totmisses,omitempty"`
	/**
	* Total packets that did not match any PBR6
	*/
	Pbr6missesrate int32 `json:"pbr6missesrate,omitempty"`
	Pbr6totnulldrop uint64 `json:"pbr6totnulldrop,omitempty"`
	/**
	* Total packets that are dropped due to null next hop
	*/
	Pbr6nulldroprate int32 `json:"pbr6nulldroprate,omitempty"`
	Pbr6perhits uint64 `json:"pbr6perhits,omitempty"`
	/**
	* Number of times the pbr6 was hit
	*/
	Pbr6perhitsrate int32 `json:"pbr6perhitsrate,omitempty"`

}
