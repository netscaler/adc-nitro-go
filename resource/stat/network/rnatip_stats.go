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
* Statistics for RNAT ipaddress resource.
*/

type Rnatipstats struct {
	/**
	* Specifies the NAT IP address of the configured RNAT entry for which you want to see the statistics. If you do not specify an IP address, this displays the statistics for all the configured RNAT entries.
	*/
	Rnatip string `json:"Rnatip,omitempty"`
	/**
	* Clear the statsistics / counters
	*/
	Clearstats string `json:"clearstats,omitempty"`
	Iptd uint32 `json:"iptd,omitempty"`
	Iprnattotrxbytes uint64 `json:"iprnattotrxbytes,omitempty"`
	/**
	* Bytes received on this IP address during RNAT sessions.
	*/
	Iprnatrxbytesrate int32 `json:"iprnatrxbytesrate,omitempty"`
	Iprnattottxbytes uint64 `json:"iprnattottxbytes,omitempty"`
	/**
	* Bytes sent from this IP address during RNAT sessions.
	*/
	Iprnattxbytesrate int32 `json:"iprnattxbytesrate,omitempty"`
	Iprnattotrxpkts uint64 `json:"iprnattotrxpkts,omitempty"`
	/**
	* Packets received on this IP address during RNAT sessions.
	*/
	Iprnatrxpktsrate int32 `json:"iprnatrxpktsrate,omitempty"`
	Iprnattottxpkts uint64 `json:"iprnattottxpkts,omitempty"`
	/**
	* Packets sent from this IP address during RNAT sessions.
	*/
	Iprnattxpktsrate int32 `json:"iprnattxpktsrate,omitempty"`
	Iprnattottxsyn uint64 `json:"iprnattottxsyn,omitempty"`
	/**
	* Requests for connections sent from this IP address during RNAT sessions.
	*/
	Iprnattxsynrate int32 `json:"iprnattxsynrate,omitempty"`
	Iprnatcursessions uint32 `json:"iprnatcursessions,omitempty"`

}
