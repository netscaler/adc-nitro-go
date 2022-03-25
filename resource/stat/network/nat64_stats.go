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
* Statistics for nat64 config resource.
 */

type Nat64stats struct {
	/**
	* Clear the statsistics / counters
	 */
	Clearstats          string `json:"clearstats,omitempty"`
	Nat64tottcpsessions int    `json:"nat64tottcpsessions,omitempty"`
	/**
	* Total number of TCP sessions created by NAT64.
	 */
	Nat64tcpsessionsrate float64 `json:"nat64tcpsessionsrate,omitempty"`
	Nat64totudpsessions  int     `json:"nat64totudpsessions,omitempty"`
	/**
	* Total number of UDP sessions created by NAT64.
	 */
	Nat64udpsessionsrate float64 `json:"nat64udpsessionsrate,omitempty"`
	Nat64toticmpsessions int     `json:"nat64toticmpsessions,omitempty"`
	/**
	* Total number of ICMP sessions created by NAT64.
	 */
	Nat64icmpsessionsrate float64 `json:"nat64icmpsessionsrate,omitempty"`
	Nat64totsessions      int     `json:"nat64totsessions,omitempty"`
	/**
	* Total number of sessions created by NAT64.
	 */
	Nat64sessionsrate float64 `json:"nat64sessionsrate,omitempty"`
}
