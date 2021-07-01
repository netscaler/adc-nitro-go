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
	Lsngrpcursessions int `json:"lsngrpcursessions,omitempty"`
	/**
	* Number of Current Sessions for LSN group
	*/
	Lsngrpcursessionsrate float64 `json:"lsngrpcursessionsrate,omitempty"`
	Lsngrptottranslbytes int `json:"lsngrptottranslbytes,omitempty"`
	/**
	* Number of Translated Bytes for LSN group
	*/
	Lsngrptranslbytesrate float64 `json:"lsngrptranslbytesrate,omitempty"`
	Lsngrptottranslpkts int `json:"lsngrptottranslpkts,omitempty"`
	/**
	* Number of Translated Pkts for LSN group
	*/
	Lsngrptranslpktsrate float64 `json:"lsngrptranslpktsrate,omitempty"`
	Lsngrptottcptranslpkts int `json:"lsngrptottcptranslpkts,omitempty"`
	/**
	* Number of TCP Translated Pkts for LSN group
	*/
	Lsngrptcptranslpktsrate float64 `json:"lsngrptcptranslpktsrate,omitempty"`
	Lsngrptottcptranslbytes int `json:"lsngrptottcptranslbytes,omitempty"`
	/**
	* Number of TCP Translated Bytes for LSN group
	*/
	Lsngrptcptranslbytesrate float64 `json:"lsngrptcptranslbytesrate,omitempty"`
	Lsngrptottcpdrppkts int `json:"lsngrptottcpdrppkts,omitempty"`
	/**
	* Number of TCP Dropped Pkts for LSN group
	*/
	Lsngrptcpdrppktsrate float64 `json:"lsngrptcpdrppktsrate,omitempty"`
	Lsngrpcurtcpsessions int `json:"lsngrpcurtcpsessions,omitempty"`
	/**
	* Number of TCP Current Sessions for LSN group
	*/
	Lsngrpcurtcpsessionsrate float64 `json:"lsngrpcurtcpsessionsrate,omitempty"`
	Lsngrptotudptranslpkts int `json:"lsngrptotudptranslpkts,omitempty"`
	/**
	* Number of UDP Translated Pkts for LSN group
	*/
	Lsngrpudptranslpktsrate float64 `json:"lsngrpudptranslpktsrate,omitempty"`
	Lsngrptotudptranslbytes int `json:"lsngrptotudptranslbytes,omitempty"`
	/**
	* Number of UDP Translated Bytes for LSN group
	*/
	Lsngrpudptranslbytesrate float64 `json:"lsngrpudptranslbytesrate,omitempty"`
	Lsngrptotudpdrppkts int `json:"lsngrptotudpdrppkts,omitempty"`
	/**
	* Number of UDP Dropped Pkts for LSN group
	*/
	Lsngrpudpdrppktsrate float64 `json:"lsngrpudpdrppktsrate,omitempty"`
	Lsngrpcurudpsessions int `json:"lsngrpcurudpsessions,omitempty"`
	/**
	* Number of UDP Current Sessions for LSN group
	*/
	Lsngrpcurudpsessionsrate float64 `json:"lsngrpcurudpsessionsrate,omitempty"`
	Lsngrptoticmptranslpkts int `json:"lsngrptoticmptranslpkts,omitempty"`
	/**
	* Number of ICMP Translated Pkts for LSN group
	*/
	Lsngrpicmptranslpktsrate float64 `json:"lsngrpicmptranslpktsrate,omitempty"`
	Lsngrptoticmptranslbytes int `json:"lsngrptoticmptranslbytes,omitempty"`
	/**
	* Number of ICMP Translated Bytes for LSN group
	*/
	Lsngrpicmptranslbytesrate float64 `json:"lsngrpicmptranslbytesrate,omitempty"`
	Lsngrptoticmpdrppkts int `json:"lsngrptoticmpdrppkts,omitempty"`
	/**
	* Number of ICMP Dropped Pkts for LSN group
	*/
	Lsngrpicmpdrppktsrate float64 `json:"lsngrpicmpdrppktsrate,omitempty"`
	Lsngrpcuricmpsessions int `json:"lsngrpcuricmpsessions,omitempty"`
	/**
	* Number of ICMP Current Sessions for LSN group
	*/
	Lsngrpcuricmpsessionsrate float64 `json:"lsngrpcuricmpsessionsrate,omitempty"`
	Lsngrpcursubscribers int `json:"lsngrpcursubscribers,omitempty"`
	/**
	* Number of ICMP Current Sessions for LSN group
	*/
	Lsngrpcursubscribersrate float64 `json:"lsngrpcursubscribersrate,omitempty"`

}
