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

package tunnel

/**
* Statistics for tunnel policy resource.
 */

type Tunneltrafficpolicystats struct {
	/**
	* Name of the advanced tunnel traffic policy.
	 */
	Name string `json:"name,omitempty"`
	/**
	* Clear the statsistics / counters
	 */
	Clearstats   string `json:"clearstats,omitempty"`
	Pipolicyhits int    `json:"pipolicyhits,omitempty"`
	/**
	* Number of hits on the policy
	 */
	Pipolicyhitsrate  float64 `json:"pipolicyhitsrate,omitempty"`
	Pipolicyundefhits int     `json:"pipolicyundefhits,omitempty"`
	/**
	* Number of undef hits on the policy
	 */
	Pipolicyundefhitsrate float64 `json:"pipolicyundefhitsrate,omitempty"`
}
