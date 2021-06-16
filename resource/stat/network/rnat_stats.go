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
* Statistics for RNAT configured route resource.
*/

type Rnatstats struct {
	/**
	* Clear the statsistics / counters
	*/
	Clearstats string `json:"clearstats,omitempty"`
	Rnattotrxbytes uint64 `json:"rnattotrxbytes,omitempty"`
	/**
	* Bytes received during RNAT sessions.
	*/
	Rnatrxbytesrate float64 `json:"rnatrxbytesrate,omitempty"`
	Rnattottxbytes uint64 `json:"rnattottxbytes,omitempty"`
	/**
	* Bytes sent during RNAT sessions.
	*/
	Rnattxbytesrate float64 `json:"rnattxbytesrate,omitempty"`
	Rnattotrxpkts uint64 `json:"rnattotrxpkts,omitempty"`
	/**
	* Packets received during RNAT sessions.
	*/
	Rnatrxpktsrate float64 `json:"rnatrxpktsrate,omitempty"`
	Rnattottxpkts uint64 `json:"rnattottxpkts,omitempty"`
	/**
	* Packets sent during RNAT sessions.
	*/
	Rnattxpktsrate float64 `json:"rnattxpktsrate,omitempty"`
	Rnattottxsyn uint64 `json:"rnattottxsyn,omitempty"`
	/**
	* Requests for connections sent during RNAT sessions.
	*/
	Rnattxsynrate float64 `json:"rnattxsynrate,omitempty"`
	Rnatcursessions uint32 `json:"rnatcursessions,omitempty"`

}
