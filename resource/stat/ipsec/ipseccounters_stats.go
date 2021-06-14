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
	Ipsectotrxbytes uint64 `json:"ipsectotrxbytes,omitempty"`
	/**
	* Bytes received during IPsec sessions.
	*/
	Ipsecrxbytesrate int32 `json:"ipsecrxbytesrate,omitempty"`
	Ipsectottxbytes uint64 `json:"ipsectottxbytes,omitempty"`
	/**
	* Bytes sent during IPsec sessions.
	*/
	Ipsectxbytesrate int32 `json:"ipsectxbytesrate,omitempty"`
	Ipsectotrxpkts uint64 `json:"ipsectotrxpkts,omitempty"`
	/**
	* Packets received during IPsec sessions.
	*/
	Ipsecrxpktsrate int32 `json:"ipsecrxpktsrate,omitempty"`
	Ipsectottxpkts uint64 `json:"ipsectottxpkts,omitempty"`
	/**
	* Packets sent during IPsec sessions.
	*/
	Ipsectxpktsrate int32 `json:"ipsectxpktsrate,omitempty"`

}
