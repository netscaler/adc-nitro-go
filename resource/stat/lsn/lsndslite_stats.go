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


type Lsndslitestats struct {
	/**
	* Clear the statsistics / counters
	*/
	Clearstats string `json:"clearstats,omitempty"`
	Lsntotdsliterxpkts uint64 `json:"lsntotdsliterxpkts,omitempty"`
	/**
	* Total number of LSN DS-Lite rx pkts.
	*/
	Lsndsliterxpktsrate float64 `json:"lsndsliterxpktsrate,omitempty"`
	Lsntotdsliterxbytes uint64 `json:"lsntotdsliterxbytes,omitempty"`
	/**
	* Total number of LSN DS-Lite rx bytes.
	*/
	Lsndsliterxbytesrate float64 `json:"lsndsliterxbytesrate,omitempty"`
	Lsntotdslitetxpkts uint64 `json:"lsntotdslitetxpkts,omitempty"`
	/**
	* Total number of LSN DS-Lite tx pkts.
	*/
	Lsndslitetxpktsrate float64 `json:"lsndslitetxpktsrate,omitempty"`
	Lsntotdslitetxbytes uint64 `json:"lsntotdslitetxbytes,omitempty"`
	/**
	* Total number of LSN DS-Lite tx bytes.
	*/
	Lsndslitetxbytesrate float64 `json:"lsndslitetxbytesrate,omitempty"`
	Lsntotdslitetcprxpkts uint64 `json:"lsntotdslitetcprxpkts,omitempty"`
	/**
	* Number of LSN DS-Lite TCP Received packets.
	*/
	Lsndslitetcprxpktsrate float64 `json:"lsndslitetcprxpktsrate,omitempty"`
	Lsntotdslitetcprxbytes uint64 `json:"lsntotdslitetcprxbytes,omitempty"`
	/**
	* Number of LSN DS-Lite TCP Received bytes.
	*/
	Lsndslitetcprxbytesrate float64 `json:"lsndslitetcprxbytesrate,omitempty"`
	Lsntotdslitetcptxpkts uint64 `json:"lsntotdslitetcptxpkts,omitempty"`
	/**
	* Number of LSN DS-Lite TCP Transmitted packets.
	*/
	Lsndslitetcptxpktsrate float64 `json:"lsndslitetcptxpktsrate,omitempty"`
	Lsntotdslitetcptxbytes uint64 `json:"lsntotdslitetcptxbytes,omitempty"`
	/**
	* Number of LSN DS-Lite TCP Transmitted bytes.
	*/
	Lsndslitetcptxbytesrate float64 `json:"lsndslitetcptxbytesrate,omitempty"`
	Lsntotdslitetcpdrppkts uint64 `json:"lsntotdslitetcpdrppkts,omitempty"`
	/**
	* Number of LSN DS-Lite TCP Dropped packets.
	*/
	Lsndslitetcpdrppktsrate float64 `json:"lsndslitetcpdrppktsrate,omitempty"`
	Lsncurdslitetcpsessions uint64 `json:"lsncurdslitetcpsessions,omitempty"`
	/**
	* Number of LSN DS-Lite TCP Current Sessions.
	*/
	Lsncurdslitetcpsessionsrate float64 `json:"lsncurdslitetcpsessionsrate,omitempty"`
	Lsntotdsliteudprxpkts uint64 `json:"lsntotdsliteudprxpkts,omitempty"`
	/**
	* Number of LSN DS-Lite UDP Received packets.
	*/
	Lsndsliteudprxpktsrate float64 `json:"lsndsliteudprxpktsrate,omitempty"`
	Lsntotdsliteudprxbytes uint64 `json:"lsntotdsliteudprxbytes,omitempty"`
	/**
	* Number of LSN DS-Lite UDP Received bytes.
	*/
	Lsndsliteudprxbytesrate float64 `json:"lsndsliteudprxbytesrate,omitempty"`
	Lsntotdsliteudptxpkts uint64 `json:"lsntotdsliteudptxpkts,omitempty"`
	/**
	* Number of LSN DS-Lite UDP Transmitted packets.
	*/
	Lsndsliteudptxpktsrate float64 `json:"lsndsliteudptxpktsrate,omitempty"`
	Lsntotdsliteudptxbytes uint64 `json:"lsntotdsliteudptxbytes,omitempty"`
	/**
	* Number of LSN DS-Lite UDP Transmitted bytes.
	*/
	Lsndsliteudptxbytesrate float64 `json:"lsndsliteudptxbytesrate,omitempty"`
	Lsntotdsliteudpdrppkts uint64 `json:"lsntotdsliteudpdrppkts,omitempty"`
	/**
	* Number of LSN DS-Lite UDP Dropped packets.
	*/
	Lsndsliteudpdrppktsrate float64 `json:"lsndsliteudpdrppktsrate,omitempty"`
	Lsncurdsliteudpsessions uint64 `json:"lsncurdsliteudpsessions,omitempty"`
	/**
	* Number of LSN DS-Lite UDP Current Sessions.
	*/
	Lsncurdsliteudpsessionsrate float64 `json:"lsncurdsliteudpsessionsrate,omitempty"`
	Lsntotdsliteicmprxpkts uint64 `json:"lsntotdsliteicmprxpkts,omitempty"`
	/**
	* Number of LSN DS-Lite ICMP Received packets.
	*/
	Lsndsliteicmprxpktsrate float64 `json:"lsndsliteicmprxpktsrate,omitempty"`
	Lsntotdsliteicmprxbytes uint64 `json:"lsntotdsliteicmprxbytes,omitempty"`
	/**
	* Number of LSN DS-Lite ICMP Received bytes.
	*/
	Lsndsliteicmprxbytesrate float64 `json:"lsndsliteicmprxbytesrate,omitempty"`
	Lsntotdsliteicmptxpkts uint64 `json:"lsntotdsliteicmptxpkts,omitempty"`
	/**
	* Number of LSN DS-Lite ICMP Transmitted packets.
	*/
	Lsndsliteicmptxpktsrate float64 `json:"lsndsliteicmptxpktsrate,omitempty"`
	Lsntotdsliteicmptxbytes uint64 `json:"lsntotdsliteicmptxbytes,omitempty"`
	/**
	* Number of LSN DS-Lite ICMP Transmitted bytes.
	*/
	Lsndsliteicmptxbytesrate float64 `json:"lsndsliteicmptxbytesrate,omitempty"`
	Lsntotdsliteicmpdrppkts uint64 `json:"lsntotdsliteicmpdrppkts,omitempty"`
	/**
	* Number of LSN DS-Lite ICMP Dropped packets.
	*/
	Lsndsliteicmpdrppktsrate float64 `json:"lsndsliteicmpdrppktsrate,omitempty"`
	Lsncurdsliteicmpsessions uint64 `json:"lsncurdsliteicmpsessions,omitempty"`
	/**
	* Number of LSN DS-Lite ICMP Current Sessions.
	*/
	Lsncurdsliteicmpsessionsrate float64 `json:"lsncurdsliteicmpsessionsrate,omitempty"`
	Lsncurdslitesessions uint64 `json:"lsncurdslitesessions,omitempty"`
	/**
	* Current number of LSN DS-Lite sessions.
	*/
	Lsncurdslitesessionsrate float64 `json:"lsncurdslitesessionsrate,omitempty"`
	Lsndslitecursubscribers uint64 `json:"lsndslitecursubscribers,omitempty"`
	/**
	* Current number of LSN DS-Lite subscribers.
	*/
	Lsndslitecursubscribersrate float64 `json:"lsndslitecursubscribersrate,omitempty"`

}
