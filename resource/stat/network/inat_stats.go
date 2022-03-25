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
* Statistics for inbound nat resource.
 */

type Inatstats struct {
	/**
	* The INAT.
	 */
	Name string `json:"name,omitempty"`
	/**
	* Clear the statsistics / counters
	 */
	Clearstats    string `json:"clearstats,omitempty"`
	Nat46tottcp46 int    `json:"nat46tottcp46,omitempty"`
	/**
	* Total TCP packets translated (V4->v6).
	 */
	Nat46tcp46rate float64 `json:"nat46tcp46rate,omitempty"`
	Nat46totudp46  int     `json:"nat46totudp46,omitempty"`
	/**
	* Total UDP packets translated (V4->v6).
	 */
	Nat46udp46rate float64 `json:"nat46udp46rate,omitempty"`
	Nat46toticmp46 int     `json:"nat46toticmp46,omitempty"`
	/**
	* Total ICMP packets translated (V4->v6).
	 */
	Nat46icmp46rate float64 `json:"nat46icmp46rate,omitempty"`
	Nat46totdrop46  int     `json:"nat46totdrop46,omitempty"`
	/**
	* Total IPV4 packets dropped.
	 */
	Nat46drop46rate float64 `json:"nat46drop46rate,omitempty"`
	Nat46tottcp64   int     `json:"nat46tottcp64,omitempty"`
	/**
	* Total TCP packets translated (V6->v4).
	 */
	Nat46tcp64rate float64 `json:"nat46tcp64rate,omitempty"`
	Nat46totudp64  int     `json:"nat46totudp64,omitempty"`
	/**
	* Total UDP packets translated (V6->v4).
	 */
	Nat46udp64rate float64 `json:"nat46udp64rate,omitempty"`
	Nat46toticmp64 int     `json:"nat46toticmp64,omitempty"`
	/**
	* Total ICMP packets translated (V6->v4).
	 */
	Nat46icmp64rate float64 `json:"nat46icmp64rate,omitempty"`
	Nat46totdrop64  int     `json:"nat46totdrop64,omitempty"`
	/**
	* Total IPV6 packets dropped.
	 */
	Nat46drop64rate float64 `json:"nat46drop64rate,omitempty"`
	Inatnat46tcp46  int     `json:"inatnat46tcp46,omitempty"`
	/**
	* TCP packets translated (V4->v6).
	 */
	Inatnat46tcp46rate float64 `json:"inatnat46tcp46rate,omitempty"`
	Inatnat46udp46     int     `json:"inatnat46udp46,omitempty"`
	/**
	* UDP packets translated (V4->v6).
	 */
	Inatnat46udp46rate float64 `json:"inatnat46udp46rate,omitempty"`
	Inatnat46icmp46    int     `json:"inatnat46icmp46,omitempty"`
	/**
	* ICMP packets translated (V4->v6).
	 */
	Inatnat46icmp46rate float64 `json:"inatnat46icmp46rate,omitempty"`
	Inatnat46drop46     int     `json:"inatnat46drop46,omitempty"`
	/**
	* IPV4 packets dropped.
	 */
	Inatnat46drop46rate float64 `json:"inatnat46drop46rate,omitempty"`
	Inatnat46tcp64      int     `json:"inatnat46tcp64,omitempty"`
	/**
	* TCP packets translated (V6->v4).
	 */
	Inatnat46tcp64rate float64 `json:"inatnat46tcp64rate,omitempty"`
	Inatnat46udp64     int     `json:"inatnat46udp64,omitempty"`
	/**
	* UDP packets translated (V6->v4).
	 */
	Inatnat46udp64rate float64 `json:"inatnat46udp64rate,omitempty"`
	Inatnat46icmp64    int     `json:"inatnat46icmp64,omitempty"`
	/**
	* ICMP packets translated (V6->v4).
	 */
	Inatnat46icmp64rate float64 `json:"inatnat46icmp64rate,omitempty"`
	Inatnat46drop64     int     `json:"inatnat46drop64,omitempty"`
	/**
	* IPV6 packets dropped.
	 */
	Inatnat46drop64rate float64 `json:"inatnat46drop64rate,omitempty"`
}
