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

package ipsec

/**
* Statistics for counters resource.
*/

type Ipseccountersstats struct {
	/**
	* Clear the statsistics / counters
	*/
	Clearstats string `json:"clearstats,omitempty"`
	Ipsectotrxbytes int `json:"ipsectotrxbytes,omitempty"`
	/**
	* Bytes received during IPsec sessions.
	*/
	Ipsecrxbytesrate float64 `json:"ipsecrxbytesrate,omitempty"`
	Ipsectottxbytes int `json:"ipsectottxbytes,omitempty"`
	/**
	* Bytes sent during IPsec sessions.
	*/
	Ipsectxbytesrate float64 `json:"ipsectxbytesrate,omitempty"`
	Ipsectotrxpkts int `json:"ipsectotrxpkts,omitempty"`
	/**
	* Packets received during IPsec sessions.
	*/
	Ipsecrxpktsrate float64 `json:"ipsecrxpktsrate,omitempty"`
	Ipsectottxpkts int `json:"ipsectottxpkts,omitempty"`
	/**
	* Packets sent during IPsec sessions.
	*/
	Ipsectxpktsrate float64 `json:"ipsectxpktsrate,omitempty"`

}
