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

package ipsecalg

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
	Clearstats string `json:"clearstats,omitempty"`
	Ipsecalgtotsessions uint64 `json:"ipsecalgtotsessions,omitempty"`
	/**
	* Total session for IPSec ALG.
	*/
	Ipsecalgsessionsrate int32 `json:"ipsecalgsessionsrate,omitempty"`
	Ipsecalgtotdrsessions uint64 `json:"ipsecalgtotdrsessions,omitempty"`
	/**
	* Total Dropped IPsec ALG sessions.
	*/
	Ipsecalgdrsessionsrate int32 `json:"ipsecalgdrsessionsrate,omitempty"`
	Ipsecalgcuractsessions uint64 `json:"ipsecalgcuractsessions,omitempty"`
	/**
	* Currently active IPsec ALG sessions.
	*/
	Ipsecalgcuractsessionsrate int32 `json:"ipsecalgcuractsessionsrate,omitempty"`
	Ipsecalgcurblksessions uint64 `json:"ipsecalgcurblksessions,omitempty"`
	/**
	* Currently blocked sessions on ESP Gate.
	*/
	Ipsecalgcurblksessionsrate int32 `json:"ipsecalgcurblksessionsrate,omitempty"`
	Ipsecalgtotrxpkts uint64 `json:"ipsecalgtotrxpkts,omitempty"`
	/**
	* Total Packets received during IPsec ALG sessions.
	*/
	Ipsecalgrxpktsrate int32 `json:"ipsecalgrxpktsrate,omitempty"`
	Ipsecalgtottxpkts uint64 `json:"ipsecalgtottxpkts,omitempty"`
	/**
	* Total Packets sent during IPsec ALG sessions.
	*/
	Ipsecalgtxpktsrate int32 `json:"ipsecalgtxpktsrate,omitempty"`

}
