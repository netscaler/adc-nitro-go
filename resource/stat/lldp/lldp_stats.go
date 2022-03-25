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

package lldp

type Lldpstats struct {
	/**
	* LLDP Statistics per interfaces
	 */
	Ifnum string `json:"ifnum,omitempty"`
	/**
	* Clear the statsistics / counters
	 */
	Clearstats        string `json:"clearstats,omitempty"`
	Rxportframestotal int    `json:"rxportframestotal,omitempty"`
	/**
	* Total LLDP Packets received.
	 */
	Rxportframesrate float64 `json:"rxportframesrate,omitempty"`
	Rxportbytestotal int     `json:"rxportbytestotal,omitempty"`
	/**
	* Total LLDP bytes received
	 */
	Rxportbytesrate   float64 `json:"rxportbytesrate,omitempty"`
	Txportframestotal int     `json:"txportframestotal,omitempty"`
	/**
	* Total LLDP Packets transmitted
	 */
	Txportframesrate float64 `json:"txportframesrate,omitempty"`
	Txportbytestotal int     `json:"txportbytestotal,omitempty"`
	/**
	* Total LLDP bytes transmitted.
	 */
	Txportbytesrate    float64 `json:"txportbytesrate,omitempty"`
	Rxportframeserrors int     `json:"rxportframeserrors,omitempty"`
	/**
	* Total errors in LLDP packets.
	 */
	Rxportframeserrorsrate   float64 `json:"rxportframeserrorsrate,omitempty"`
	Rxporttlvsdiscardedtotal int     `json:"rxporttlvsdiscardedtotal,omitempty"`
	/**
	* Total discarded LLDP packets.
	 */
	Rxporttlvsdiscardedrate     float64 `json:"rxporttlvsdiscardedrate,omitempty"`
	Rxporttlvsunrecognizedtotal int     `json:"rxporttlvsunrecognizedtotal,omitempty"`
	/**
	* Total TLVs not Recognised.
	 */
	Rxporttlvsunrecognizedrate float64 `json:"rxporttlvsunrecognizedrate,omitempty"`
}
