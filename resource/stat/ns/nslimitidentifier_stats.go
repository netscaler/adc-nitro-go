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
* Statistics for limit Indetifier resource.
 */

type Nslimitidentifierstats struct {
	/**
	* The name of the identifier.
	 */
	Name string `json:"name,omitempty"`
	/**
	* Pattern for the selector field, ? means field is required, * means field value does not matter, anything else is a regular pattern
	 */
	Pattern []string `json:"pattern,omitempty"`
	/**
	* Clear the statsistics / counters
	 */
	Clearstats string `json:"clearstats,omitempty"`
	/**
	* use this argument to sort by specific key
	 */
	Sortby string `json:"sortby,omitempty"`
	/**
	* use this argument to specify sort order
	 */
	Sortorder             string `json:"sortorder,omitempty"`
	Ratelmtobjhits        int    `json:"ratelmtobjhits,omitempty"`
	Ratelmtobjdrops       int    `json:"ratelmtobjdrops,omitempty"`
	Ratelmtsessionobjhits int    `json:"ratelmtsessionobjhits,omitempty"`
}
