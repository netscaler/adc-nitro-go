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

package ha

/**
* Statistics for node resource.
 */

type Hanodestats struct {
	/**
	* Clear the statsistics / counters
	 */
	Clearstats       string `json:"clearstats,omitempty"`
	Hacurstatus      string `json:"hacurstatus,omitempty"`
	Hacurstate       string `json:"hacurstate,omitempty"`
	Hacurmasterstate string `json:"hacurmasterstate,omitempty"`
	Transtime        string `json:"transtime,omitempty"`
	Hatotpktrx       int    `json:"hatotpktrx,omitempty"`
	/**
	* Number of heartbeat packets received from the peer node. Heartbeats are sent at regular intervals (default is 200 milliseconds) to determine the state of the peer node.
	 */
	Hapktrxrate float64 `json:"hapktrxrate,omitempty"`
	Hatotpkttx  int     `json:"hatotpkttx,omitempty"`
	/**
	* Number of heartbeat packets sent to the peer node. Heartbeats are sent at regular intervals (default is 200 milliseconds) to determine the state of the peer node.
	 */
	Hapkttxrate      float64 `json:"hapkttxrate,omitempty"`
	Haerrproptimeout int     `json:"haerrproptimeout,omitempty"`
	Haerrsyncfailure int     `json:"haerrsyncfailure,omitempty"`
}
