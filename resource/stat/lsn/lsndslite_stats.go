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
	Lsndsliterxpktsrate int32 `json:"lsndsliterxpktsrate,omitempty"`
	Lsntotdsliterxbytes uint64 `json:"lsntotdsliterxbytes,omitempty"`
	/**
	* Total number of LSN DS-Lite rx bytes.
	*/
	Lsndsliterxbytesrate int32 `json:"lsndsliterxbytesrate,omitempty"`
	Lsntotdslitetxpkts uint64 `json:"lsntotdslitetxpkts,omitempty"`
	/**
	* Total number of LSN DS-Lite tx pkts.
	*/
	Lsndslitetxpktsrate int32 `json:"lsndslitetxpktsrate,omitempty"`
	Lsntotdslitetxbytes uint64 `json:"lsntotdslitetxbytes,omitempty"`
	/**
	* Total number of LSN DS-Lite tx bytes.
	*/
	Lsndslitetxbytesrate int32 `json:"lsndslitetxbytesrate,omitempty"`
	Lsntotdslitetcprxpkts uint64 `json:"lsntotdslitetcprxpkts,omitempty"`
	/**
	* Number of LSN DS-Lite TCP Received packets.
	*/
	Lsndslitetcprxpktsrate int32 `json:"lsndslitetcprxpktsrate,omitempty"`
	Lsntotdslitetcprxbytes uint64 `json:"lsntotdslitetcprxbytes,omitempty"`
	/**
	* Number of LSN DS-Lite TCP Received bytes.
	*/
	Lsndslitetcprxbytesrate int32 `json:"lsndslitetcprxbytesrate,omitempty"`
	Lsntotdslitetcptxpkts uint64 `json:"lsntotdslitetcptxpkts,omitempty"`
	/**
	* Number of LSN DS-Lite TCP Transmitted packets.
	*/
	Lsndslitetcptxpktsrate int32 `json:"lsndslitetcptxpktsrate,omitempty"`
	Lsntotdslitetcptxbytes uint64 `json:"lsntotdslitetcptxbytes,omitempty"`
	/**
	* Number of LSN DS-Lite TCP Transmitted bytes.
	*/
	Lsndslitetcptxbytesrate int32 `json:"lsndslitetcptxbytesrate,omitempty"`
	Lsntotdslitetcpdrppkts uint64 `json:"lsntotdslitetcpdrppkts,omitempty"`
	/**
	* Number of LSN DS-Lite TCP Dropped packets.
	*/
	Lsndslitetcpdrppktsrate int32 `json:"lsndslitetcpdrppktsrate,omitempty"`
	Lsncurdslitetcpsessions uint64 `json:"lsncurdslitetcpsessions,omitempty"`
	/**
	* Number of LSN DS-Lite TCP Current Sessions.
	*/
	Lsncurdslitetcpsessionsrate int32 `json:"lsncurdslitetcpsessionsrate,omitempty"`
	Lsntotdsliteudprxpkts uint64 `json:"lsntotdsliteudprxpkts,omitempty"`
	/**
	* Number of LSN DS-Lite UDP Received packets.
	*/
	Lsndsliteudprxpktsrate int32 `json:"lsndsliteudprxpktsrate,omitempty"`
	Lsntotdsliteudprxbytes uint64 `json:"lsntotdsliteudprxbytes,omitempty"`
	/**
	* Number of LSN DS-Lite UDP Received bytes.
	*/
	Lsndsliteudprxbytesrate int32 `json:"lsndsliteudprxbytesrate,omitempty"`
	Lsntotdsliteudptxpkts uint64 `json:"lsntotdsliteudptxpkts,omitempty"`
	/**
	* Number of LSN DS-Lite UDP Transmitted packets.
	*/
	Lsndsliteudptxpktsrate int32 `json:"lsndsliteudptxpktsrate,omitempty"`
	Lsntotdsliteudptxbytes uint64 `json:"lsntotdsliteudptxbytes,omitempty"`
	/**
	* Number of LSN DS-Lite UDP Transmitted bytes.
	*/
	Lsndsliteudptxbytesrate int32 `json:"lsndsliteudptxbytesrate,omitempty"`
	Lsntotdsliteudpdrppkts uint64 `json:"lsntotdsliteudpdrppkts,omitempty"`
	/**
	* Number of LSN DS-Lite UDP Dropped packets.
	*/
	Lsndsliteudpdrppktsrate int32 `json:"lsndsliteudpdrppktsrate,omitempty"`
	Lsncurdsliteudpsessions uint64 `json:"lsncurdsliteudpsessions,omitempty"`
	/**
	* Number of LSN DS-Lite UDP Current Sessions.
	*/
	Lsncurdsliteudpsessionsrate int32 `json:"lsncurdsliteudpsessionsrate,omitempty"`
	Lsntotdsliteicmprxpkts uint64 `json:"lsntotdsliteicmprxpkts,omitempty"`
	/**
	* Number of LSN DS-Lite ICMP Received packets.
	*/
	Lsndsliteicmprxpktsrate int32 `json:"lsndsliteicmprxpktsrate,omitempty"`
	Lsntotdsliteicmprxbytes uint64 `json:"lsntotdsliteicmprxbytes,omitempty"`
	/**
	* Number of LSN DS-Lite ICMP Received bytes.
	*/
	Lsndsliteicmprxbytesrate int32 `json:"lsndsliteicmprxbytesrate,omitempty"`
	Lsntotdsliteicmptxpkts uint64 `json:"lsntotdsliteicmptxpkts,omitempty"`
	/**
	* Number of LSN DS-Lite ICMP Transmitted packets.
	*/
	Lsndsliteicmptxpktsrate int32 `json:"lsndsliteicmptxpktsrate,omitempty"`
	Lsntotdsliteicmptxbytes uint64 `json:"lsntotdsliteicmptxbytes,omitempty"`
	/**
	* Number of LSN DS-Lite ICMP Transmitted bytes.
	*/
	Lsndsliteicmptxbytesrate int32 `json:"lsndsliteicmptxbytesrate,omitempty"`
	Lsntotdsliteicmpdrppkts uint64 `json:"lsntotdsliteicmpdrppkts,omitempty"`
	/**
	* Number of LSN DS-Lite ICMP Dropped packets.
	*/
	Lsndsliteicmpdrppktsrate int32 `json:"lsndsliteicmpdrppktsrate,omitempty"`
	Lsncurdsliteicmpsessions uint64 `json:"lsncurdsliteicmpsessions,omitempty"`
	/**
	* Number of LSN DS-Lite ICMP Current Sessions.
	*/
	Lsncurdsliteicmpsessionsrate int32 `json:"lsncurdsliteicmpsessionsrate,omitempty"`
	Lsncurdslitesessions uint64 `json:"lsncurdslitesessions,omitempty"`
	/**
	* Current number of LSN DS-Lite sessions.
	*/
	Lsncurdslitesessionsrate int32 `json:"lsncurdslitesessionsrate,omitempty"`
	Lsndslitecursubscribers uint64 `json:"lsndslitecursubscribers,omitempty"`
	/**
	* Current number of LSN DS-Lite subscribers.
	*/
	Lsndslitecursubscribersrate int32 `json:"lsndslitecursubscribersrate,omitempty"`

}
