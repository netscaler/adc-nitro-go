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
	Clearstats string `json:"clearstats,omitempty"`
	Rxportframestotal uint64 `json:"rxportframestotal,omitempty"`
	/**
	* Total LLDP Packets received.
	*/
	Rxportframesrate int32 `json:"rxportframesrate,omitempty"`
	Rxportbytestotal uint64 `json:"rxportbytestotal,omitempty"`
	/**
	* Total LLDP bytes received
	*/
	Rxportbytesrate int32 `json:"rxportbytesrate,omitempty"`
	Txportframestotal uint64 `json:"txportframestotal,omitempty"`
	/**
	* Total LLDP Packets transmitted
	*/
	Txportframesrate int32 `json:"txportframesrate,omitempty"`
	Txportbytestotal uint64 `json:"txportbytestotal,omitempty"`
	/**
	* Total LLDP bytes transmitted.
	*/
	Txportbytesrate int32 `json:"txportbytesrate,omitempty"`
	Rxportframeserrors uint64 `json:"rxportframeserrors,omitempty"`
	/**
	* Total errors in LLDP packets.
	*/
	Rxportframeserrorsrate int32 `json:"rxportframeserrorsrate,omitempty"`
	Rxporttlvsdiscardedtotal uint64 `json:"rxporttlvsdiscardedtotal,omitempty"`
	/**
	* Total discarded LLDP packets.
	*/
	Rxporttlvsdiscardedrate int32 `json:"rxporttlvsdiscardedrate,omitempty"`
	Rxporttlvsunrecognizedtotal uint64 `json:"rxporttlvsunrecognizedtotal,omitempty"`
	/**
	* Total TLVs not Recognised.
	*/
	Rxporttlvsunrecognizedrate int32 `json:"rxporttlvsunrecognizedrate,omitempty"`

}
