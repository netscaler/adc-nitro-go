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
* Statistics for Traffic Domain resource.
 */

type Nstrafficdomainstats struct {
	/**
	* An integer specifying the Traffic Domain ID. Possible values: 1 through 4094.
	 */
	Td int `json:"td,omitempty"`
	/**
	* Clear the statsistics / counters
	 */
	Clearstats    string `json:"clearstats,omitempty"`
	Nstdtotrxpkts int    `json:"nstdtotrxpkts,omitempty"`
	/**
	* Packets received on this TD.
	 */
	Nstdrxpktsrate float64 `json:"nstdrxpktsrate,omitempty"`
	Nstdtottxpkts  int     `json:"nstdtottxpkts,omitempty"`
	/**
	* Packets transmitted from this TD.
	 */
	Nstdtxpktsrate     float64 `json:"nstdtxpktsrate,omitempty"`
	Nstdtotdroppedpkts int     `json:"nstdtotdroppedpkts,omitempty"`
	/**
	* Inbound packets dropped on this TD by reception.
	 */
	Nstddroppedpktsrate float64 `json:"nstddroppedpktsrate,omitempty"`
}
