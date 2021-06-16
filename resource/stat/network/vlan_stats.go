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

package network

/**
* Statistics for "VLAN" resource.
*/

type Vlanstats struct {
	/**
	* An integer specifying the VLAN identification number (VID). Possible values: 1 through 4094.
	*/
	Id uint32 `json:"id,omitempty"`
	/**
	* Clear the statsistics / counters
	*/
	Clearstats string `json:"clearstats,omitempty"`
	Vlantotrxpkts uint64 `json:"vlantotrxpkts,omitempty"`
	/**
	* Packets received on the VLAN.
	*/
	Vlanrxpktsrate float64 `json:"vlanrxpktsrate,omitempty"`
	Vlantotrxbytes uint64 `json:"vlantotrxbytes,omitempty"`
	/**
	* Bytes of data received on the VLAN.
	*/
	Vlanrxbytesrate float64 `json:"vlanrxbytesrate,omitempty"`
	Vlantottxpkts uint64 `json:"vlantottxpkts,omitempty"`
	/**
	* Packets transmitted on the VLAN.
	*/
	Vlantxpktsrate float64 `json:"vlantxpktsrate,omitempty"`
	Vlantottxbytes uint64 `json:"vlantottxbytes,omitempty"`
	/**
	* Bytes of data transmitted on the VLAN.
	*/
	Vlantxbytesrate float64 `json:"vlantxbytesrate,omitempty"`
	Vlantotdroppedpkts uint64 `json:"vlantotdroppedpkts,omitempty"`
	Vlantotbroadcastpkts uint64 `json:"vlantotbroadcastpkts,omitempty"`

}
