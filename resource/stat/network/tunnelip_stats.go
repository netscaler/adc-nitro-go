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
* Statistics for TUNNEL Remote ipaddress resource.
 */

type Tunnelipstats struct {
	/**
	* remote IP address of the configured tunnel.
	 */
	Tunnelip string `json:"tunnelip,omitempty"`
	/**
	* Clear the statsistics / counters
	 */
	Clearstats   string `json:"clearstats,omitempty"`
	Tnltotrxpkts int    `json:"tnltotrxpkts,omitempty"`
	/**
	* Total number of packets received on the tunnel.
	 */
	Tnlrxpktsrate float64 `json:"tnlrxpktsrate,omitempty"`
	Tnltottxpkts  int     `json:"tnltottxpkts,omitempty"`
	/**
	* Total number of packets transmitted on the tunnel.
	 */
	Tnltxpktsrate float64 `json:"tnltxpktsrate,omitempty"`
	Tnltotrxbytes int     `json:"tnltotrxbytes,omitempty"`
	/**
	* Total number of bytes received on the tunnel.
	 */
	Tnlrxbytesrate float64 `json:"tnlrxbytesrate,omitempty"`
	Tnltottxbytes  int     `json:"tnltottxbytes,omitempty"`
	/**
	* Total number of bytes transmitted on the tunnel.
	 */
	Tnltxbytesrate float64 `json:"tnltxbytesrate,omitempty"`
}
