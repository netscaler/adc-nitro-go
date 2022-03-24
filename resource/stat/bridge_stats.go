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
* Statistics for bridge resource.
 */

type Bridgestats struct {
	/**
	* Clear the statsistics / counters
	 */
	Clearstats        string `json:"clearstats,omitempty"`
	Tcptotbdgmacmoved int    `json:"tcptotbdgmacmoved,omitempty"`
	/**
	* The number of times bridging registered MAC moved
	 */
	Tcpbdgmacmovedrate  float64 `json:"tcpbdgmacmovedrate,omitempty"`
	Tcptotbdgcollisions int     `json:"tcptotbdgcollisions,omitempty"`
	/**
	* The number of bridging table collisions
	 */
	Tcpbdgcollisionsrate float64 `json:"tcpbdgcollisionsrate,omitempty"`
	Tcperrbdgmuted       int     `json:"tcperrbdgmuted,omitempty"`
	/**
	* The number of bridging related interface mutes
	 */
	Tcperrbdgmutedrate float64 `json:"tcperrbdgmutedrate,omitempty"`
	Totbdgpkts         int     `json:"totbdgpkts,omitempty"`
	/**
	* The total number of bridged packets
	 */
	Bdgpktsrate float64 `json:"bdgpktsrate,omitempty"`
	Totbdgmbits int     `json:"totbdgmbits,omitempty"`
	/**
	* The total number of bridged Mbits
	 */
	Bdgmbitsrate float64 `json:"bdgmbitsrate,omitempty"`
}
