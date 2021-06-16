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


type Lsnstats struct {
	/**
	* Clear the statsistics / counters
	*/
	Clearstats string `json:"clearstats,omitempty"`
	Lsntottcprxpkts uint64 `json:"lsntottcprxpkts,omitempty"`
	/**
	* Number of LSN TCP Received packets.
	*/
	Lsntcprxpktsrate float64 `json:"lsntcprxpktsrate,omitempty"`
	Lsntottcprxbytes uint64 `json:"lsntottcprxbytes,omitempty"`
	/**
	* Number of LSN TCP Received bytes.
	*/
	Lsntcprxbytesrate float64 `json:"lsntcprxbytesrate,omitempty"`
	Lsntottcptxpkts uint64 `json:"lsntottcptxpkts,omitempty"`
	/**
	* Number of LSN TCP Transmitted packets.
	*/
	Lsntcptxpktsrate float64 `json:"lsntcptxpktsrate,omitempty"`
	Lsntottcptxbytes uint64 `json:"lsntottcptxbytes,omitempty"`
	/**
	* Number of LSN TCP Transmitted bytes.
	*/
	Lsntcptxbytesrate float64 `json:"lsntcptxbytesrate,omitempty"`
	Lsntottcpdrppkts uint64 `json:"lsntottcpdrppkts,omitempty"`
	/**
	* Number of LSN TCP Dropped packets.
	*/
	Lsntcpdrppktsrate float64 `json:"lsntcpdrppktsrate,omitempty"`
	Lsncurtcpsessions uint64 `json:"lsncurtcpsessions,omitempty"`
	/**
	* Number of LSN TCP Current Sessions.
	*/
	Lsncurtcpsessionsrate float64 `json:"lsncurtcpsessionsrate,omitempty"`
	Lsntotudprxpkts uint64 `json:"lsntotudprxpkts,omitempty"`
	/**
	* Number of LSN UDP Received packets.
	*/
	Lsnudprxpktsrate float64 `json:"lsnudprxpktsrate,omitempty"`
	Lsntotudprxbytes uint64 `json:"lsntotudprxbytes,omitempty"`
	/**
	* Number of LSN UDP Received bytes.
	*/
	Lsnudprxbytesrate float64 `json:"lsnudprxbytesrate,omitempty"`
	Lsntotudptxpkts uint64 `json:"lsntotudptxpkts,omitempty"`
	/**
	* Number of LSN UDP Transmitted packets.
	*/
	Lsnudptxpktsrate float64 `json:"lsnudptxpktsrate,omitempty"`
	Lsntotudptxbytes uint64 `json:"lsntotudptxbytes,omitempty"`
	/**
	* Number of LSN UDP Transmitted bytes.
	*/
	Lsnudptxbytesrate float64 `json:"lsnudptxbytesrate,omitempty"`
	Lsntotudpdrppkts uint64 `json:"lsntotudpdrppkts,omitempty"`
	/**
	* Number of LSN UDP Dropped packets.
	*/
	Lsnudpdrppktsrate float64 `json:"lsnudpdrppktsrate,omitempty"`
	Lsncurudpsessions uint64 `json:"lsncurudpsessions,omitempty"`
	/**
	* Number of LSN UDP Current Sessions.
	*/
	Lsncurudpsessionsrate float64 `json:"lsncurudpsessionsrate,omitempty"`
	Lsntoticmprxpkts uint64 `json:"lsntoticmprxpkts,omitempty"`
	/**
	* Number of LSN ICMP Received packets.
	*/
	Lsnicmprxpktsrate float64 `json:"lsnicmprxpktsrate,omitempty"`
	Lsntoticmprxbytes uint64 `json:"lsntoticmprxbytes,omitempty"`
	/**
	* Number of LSN ICMP Received bytes.
	*/
	Lsnicmprxbytesrate float64 `json:"lsnicmprxbytesrate,omitempty"`
	Lsntoticmptxpkts uint64 `json:"lsntoticmptxpkts,omitempty"`
	/**
	* Number of LSN ICMP Transmitted packets.
	*/
	Lsnicmptxpktsrate float64 `json:"lsnicmptxpktsrate,omitempty"`
	Lsntoticmptxbytes uint64 `json:"lsntoticmptxbytes,omitempty"`
	/**
	* Number of LSN ICMP Transmitted bytes.
	*/
	Lsnicmptxbytesrate float64 `json:"lsnicmptxbytesrate,omitempty"`
	Lsntoticmpdrppkts uint64 `json:"lsntoticmpdrppkts,omitempty"`
	/**
	* Number of LSN ICMP Dropped packets.
	*/
	Lsnicmpdrppktsrate float64 `json:"lsnicmpdrppktsrate,omitempty"`
	Lsncuricmpsessions uint64 `json:"lsncuricmpsessions,omitempty"`
	/**
	* Number of LSN ICMP Current Sessions.
	*/
	Lsncuricmpsessionsrate float64 `json:"lsncuricmpsessionsrate,omitempty"`
	Lsncursessions uint64 `json:"lsncursessions,omitempty"`
	/**
	* Current number of LSN sessions.
	*/
	Lsncursessionsrate float64 `json:"lsncursessionsrate,omitempty"`
	Lsncursubscribers uint64 `json:"lsncursubscribers,omitempty"`
	/**
	* Current number of LSN subscribers.
	*/
	Lsncursubscribersrate float64 `json:"lsncursubscribersrate,omitempty"`

}
