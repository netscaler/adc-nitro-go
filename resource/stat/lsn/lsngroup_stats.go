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

package lsn

/**
* Statistics for LSN group resource.
*/

type Lsngroupstats struct {
	/**
	* Name of the LSN Group.
	*/
	Groupname string `json:"groupname,omitempty"`
	/**
	* Clear the statsistics / counters
	*/
	Clearstats string `json:"clearstats,omitempty"`
	Lsngrpcursessions uint64 `json:"lsngrpcursessions,omitempty"`
	/**
	* Number of Current Sessions for LSN group
	*/
	Lsngrpcursessionsrate int32 `json:"lsngrpcursessionsrate,omitempty"`
	Lsngrptottranslbytes uint64 `json:"lsngrptottranslbytes,omitempty"`
	/**
	* Number of Translated Bytes for LSN group
	*/
	Lsngrptranslbytesrate int32 `json:"lsngrptranslbytesrate,omitempty"`
	Lsngrptottranslpkts uint64 `json:"lsngrptottranslpkts,omitempty"`
	/**
	* Number of Translated Pkts for LSN group
	*/
	Lsngrptranslpktsrate int32 `json:"lsngrptranslpktsrate,omitempty"`
	Lsngrptottcptranslpkts uint64 `json:"lsngrptottcptranslpkts,omitempty"`
	/**
	* Number of TCP Translated Pkts for LSN group
	*/
	Lsngrptcptranslpktsrate int32 `json:"lsngrptcptranslpktsrate,omitempty"`
	Lsngrptottcptranslbytes uint64 `json:"lsngrptottcptranslbytes,omitempty"`
	/**
	* Number of TCP Translated Bytes for LSN group
	*/
	Lsngrptcptranslbytesrate int32 `json:"lsngrptcptranslbytesrate,omitempty"`
	Lsngrptottcpdrppkts uint64 `json:"lsngrptottcpdrppkts,omitempty"`
	/**
	* Number of TCP Dropped Pkts for LSN group
	*/
	Lsngrptcpdrppktsrate int32 `json:"lsngrptcpdrppktsrate,omitempty"`
	Lsngrpcurtcpsessions uint64 `json:"lsngrpcurtcpsessions,omitempty"`
	/**
	* Number of TCP Current Sessions for LSN group
	*/
	Lsngrpcurtcpsessionsrate int32 `json:"lsngrpcurtcpsessionsrate,omitempty"`
	Lsngrptotudptranslpkts uint64 `json:"lsngrptotudptranslpkts,omitempty"`
	/**
	* Number of UDP Translated Pkts for LSN group
	*/
	Lsngrpudptranslpktsrate int32 `json:"lsngrpudptranslpktsrate,omitempty"`
	Lsngrptotudptranslbytes uint64 `json:"lsngrptotudptranslbytes,omitempty"`
	/**
	* Number of UDP Translated Bytes for LSN group
	*/
	Lsngrpudptranslbytesrate int32 `json:"lsngrpudptranslbytesrate,omitempty"`
	Lsngrptotudpdrppkts uint64 `json:"lsngrptotudpdrppkts,omitempty"`
	/**
	* Number of UDP Dropped Pkts for LSN group
	*/
	Lsngrpudpdrppktsrate int32 `json:"lsngrpudpdrppktsrate,omitempty"`
	Lsngrpcurudpsessions uint64 `json:"lsngrpcurudpsessions,omitempty"`
	/**
	* Number of UDP Current Sessions for LSN group
	*/
	Lsngrpcurudpsessionsrate int32 `json:"lsngrpcurudpsessionsrate,omitempty"`
	Lsngrptoticmptranslpkts uint64 `json:"lsngrptoticmptranslpkts,omitempty"`
	/**
	* Number of ICMP Translated Pkts for LSN group
	*/
	Lsngrpicmptranslpktsrate int32 `json:"lsngrpicmptranslpktsrate,omitempty"`
	Lsngrptoticmptranslbytes uint64 `json:"lsngrptoticmptranslbytes,omitempty"`
	/**
	* Number of ICMP Translated Bytes for LSN group
	*/
	Lsngrpicmptranslbytesrate int32 `json:"lsngrpicmptranslbytesrate,omitempty"`
	Lsngrptoticmpdrppkts uint64 `json:"lsngrptoticmpdrppkts,omitempty"`
	/**
	* Number of ICMP Dropped Pkts for LSN group
	*/
	Lsngrpicmpdrppktsrate int32 `json:"lsngrpicmpdrppktsrate,omitempty"`
	Lsngrpcuricmpsessions uint64 `json:"lsngrpcuricmpsessions,omitempty"`
	/**
	* Number of ICMP Current Sessions for LSN group
	*/
	Lsngrpcuricmpsessionsrate int32 `json:"lsngrpcuricmpsessionsrate,omitempty"`
	Lsngrpcursubscribers uint64 `json:"lsngrpcursubscribers,omitempty"`
	/**
	* Number of ICMP Current Sessions for LSN group
	*/
	Lsngrpcursubscribersrate int32 `json:"lsngrpcursubscribersrate,omitempty"`

}
