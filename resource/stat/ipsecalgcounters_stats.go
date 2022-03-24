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
* Statistics for counters resource.
 */

type Ipsecalgcountersstats struct {
	/**
	* Name of IPSec ALG Profile.
	 */
	Name string `json:"name,omitempty"`
	/**
	* Clear the statsistics / counters
	 */
	Clearstats          string `json:"clearstats,omitempty"`
	Ipsecalgtotsessions int    `json:"ipsecalgtotsessions,omitempty"`
	/**
	* Total session for IPSec ALG.
	 */
	Ipsecalgsessionsrate  float64 `json:"ipsecalgsessionsrate,omitempty"`
	Ipsecalgtotdrsessions int     `json:"ipsecalgtotdrsessions,omitempty"`
	/**
	* Total Dropped IPsec ALG sessions.
	 */
	Ipsecalgdrsessionsrate float64 `json:"ipsecalgdrsessionsrate,omitempty"`
	Ipsecalgcuractsessions int     `json:"ipsecalgcuractsessions,omitempty"`
	/**
	* Currently active IPsec ALG sessions.
	 */
	Ipsecalgcuractsessionsrate float64 `json:"ipsecalgcuractsessionsrate,omitempty"`
	Ipsecalgcurblksessions     int     `json:"ipsecalgcurblksessions,omitempty"`
	/**
	* Currently blocked sessions on ESP Gate.
	 */
	Ipsecalgcurblksessionsrate float64 `json:"ipsecalgcurblksessionsrate,omitempty"`
	Ipsecalgtotrxpkts          int     `json:"ipsecalgtotrxpkts,omitempty"`
	/**
	* Total Packets received during IPsec ALG sessions.
	 */
	Ipsecalgrxpktsrate float64 `json:"ipsecalgrxpktsrate,omitempty"`
	Ipsecalgtottxpkts  int     `json:"ipsecalgtottxpkts,omitempty"`
	/**
	* Total Packets sent during IPsec ALG sessions.
	 */
	Ipsecalgtxpktsrate float64 `json:"ipsecalgtxpktsrate,omitempty"`
}
