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
	Clearstats         string `json:"clearstats,omitempty"`
	Lsntotdsliterxpkts int    `json:"lsntotdsliterxpkts,omitempty"`
	/**
	* Total number of LSN DS-Lite rx pkts.
	 */
	Lsndsliterxpktsrate float64 `json:"lsndsliterxpktsrate,omitempty"`
	Lsntotdsliterxbytes int     `json:"lsntotdsliterxbytes,omitempty"`
	/**
	* Total number of LSN DS-Lite rx bytes.
	 */
	Lsndsliterxbytesrate float64 `json:"lsndsliterxbytesrate,omitempty"`
	Lsntotdslitetxpkts   int     `json:"lsntotdslitetxpkts,omitempty"`
	/**
	* Total number of LSN DS-Lite tx pkts.
	 */
	Lsndslitetxpktsrate float64 `json:"lsndslitetxpktsrate,omitempty"`
	Lsntotdslitetxbytes int     `json:"lsntotdslitetxbytes,omitempty"`
	/**
	* Total number of LSN DS-Lite tx bytes.
	 */
	Lsndslitetxbytesrate  float64 `json:"lsndslitetxbytesrate,omitempty"`
	Lsntotdslitetcprxpkts int     `json:"lsntotdslitetcprxpkts,omitempty"`
	/**
	* Number of LSN DS-Lite TCP Received packets.
	 */
	Lsndslitetcprxpktsrate float64 `json:"lsndslitetcprxpktsrate,omitempty"`
	Lsntotdslitetcprxbytes int     `json:"lsntotdslitetcprxbytes,omitempty"`
	/**
	* Number of LSN DS-Lite TCP Received bytes.
	 */
	Lsndslitetcprxbytesrate float64 `json:"lsndslitetcprxbytesrate,omitempty"`
	Lsntotdslitetcptxpkts   int     `json:"lsntotdslitetcptxpkts,omitempty"`
	/**
	* Number of LSN DS-Lite TCP Transmitted packets.
	 */
	Lsndslitetcptxpktsrate float64 `json:"lsndslitetcptxpktsrate,omitempty"`
	Lsntotdslitetcptxbytes int     `json:"lsntotdslitetcptxbytes,omitempty"`
	/**
	* Number of LSN DS-Lite TCP Transmitted bytes.
	 */
	Lsndslitetcptxbytesrate float64 `json:"lsndslitetcptxbytesrate,omitempty"`
	Lsntotdslitetcpdrppkts  int     `json:"lsntotdslitetcpdrppkts,omitempty"`
	/**
	* Number of LSN DS-Lite TCP Dropped packets.
	 */
	Lsndslitetcpdrppktsrate float64 `json:"lsndslitetcpdrppktsrate,omitempty"`
	Lsncurdslitetcpsessions int     `json:"lsncurdslitetcpsessions,omitempty"`
	/**
	* Number of LSN DS-Lite TCP Current Sessions.
	 */
	Lsncurdslitetcpsessionsrate float64 `json:"lsncurdslitetcpsessionsrate,omitempty"`
	Lsntotdsliteudprxpkts       int     `json:"lsntotdsliteudprxpkts,omitempty"`
	/**
	* Number of LSN DS-Lite UDP Received packets.
	 */
	Lsndsliteudprxpktsrate float64 `json:"lsndsliteudprxpktsrate,omitempty"`
	Lsntotdsliteudprxbytes int     `json:"lsntotdsliteudprxbytes,omitempty"`
	/**
	* Number of LSN DS-Lite UDP Received bytes.
	 */
	Lsndsliteudprxbytesrate float64 `json:"lsndsliteudprxbytesrate,omitempty"`
	Lsntotdsliteudptxpkts   int     `json:"lsntotdsliteudptxpkts,omitempty"`
	/**
	* Number of LSN DS-Lite UDP Transmitted packets.
	 */
	Lsndsliteudptxpktsrate float64 `json:"lsndsliteudptxpktsrate,omitempty"`
	Lsntotdsliteudptxbytes int     `json:"lsntotdsliteudptxbytes,omitempty"`
	/**
	* Number of LSN DS-Lite UDP Transmitted bytes.
	 */
	Lsndsliteudptxbytesrate float64 `json:"lsndsliteudptxbytesrate,omitempty"`
	Lsntotdsliteudpdrppkts  int     `json:"lsntotdsliteudpdrppkts,omitempty"`
	/**
	* Number of LSN DS-Lite UDP Dropped packets.
	 */
	Lsndsliteudpdrppktsrate float64 `json:"lsndsliteudpdrppktsrate,omitempty"`
	Lsncurdsliteudpsessions int     `json:"lsncurdsliteudpsessions,omitempty"`
	/**
	* Number of LSN DS-Lite UDP Current Sessions.
	 */
	Lsncurdsliteudpsessionsrate float64 `json:"lsncurdsliteudpsessionsrate,omitempty"`
	Lsntotdsliteicmprxpkts      int     `json:"lsntotdsliteicmprxpkts,omitempty"`
	/**
	* Number of LSN DS-Lite ICMP Received packets.
	 */
	Lsndsliteicmprxpktsrate float64 `json:"lsndsliteicmprxpktsrate,omitempty"`
	Lsntotdsliteicmprxbytes int     `json:"lsntotdsliteicmprxbytes,omitempty"`
	/**
	* Number of LSN DS-Lite ICMP Received bytes.
	 */
	Lsndsliteicmprxbytesrate float64 `json:"lsndsliteicmprxbytesrate,omitempty"`
	Lsntotdsliteicmptxpkts   int     `json:"lsntotdsliteicmptxpkts,omitempty"`
	/**
	* Number of LSN DS-Lite ICMP Transmitted packets.
	 */
	Lsndsliteicmptxpktsrate float64 `json:"lsndsliteicmptxpktsrate,omitempty"`
	Lsntotdsliteicmptxbytes int     `json:"lsntotdsliteicmptxbytes,omitempty"`
	/**
	* Number of LSN DS-Lite ICMP Transmitted bytes.
	 */
	Lsndsliteicmptxbytesrate float64 `json:"lsndsliteicmptxbytesrate,omitempty"`
	Lsntotdsliteicmpdrppkts  int     `json:"lsntotdsliteicmpdrppkts,omitempty"`
	/**
	* Number of LSN DS-Lite ICMP Dropped packets.
	 */
	Lsndsliteicmpdrppktsrate float64 `json:"lsndsliteicmpdrppktsrate,omitempty"`
	Lsncurdsliteicmpsessions int     `json:"lsncurdsliteicmpsessions,omitempty"`
	/**
	* Number of LSN DS-Lite ICMP Current Sessions.
	 */
	Lsncurdsliteicmpsessionsrate float64 `json:"lsncurdsliteicmpsessionsrate,omitempty"`
	Lsncurdslitesessions         int     `json:"lsncurdslitesessions,omitempty"`
	/**
	* Current number of LSN DS-Lite sessions.
	 */
	Lsncurdslitesessionsrate float64 `json:"lsncurdslitesessionsrate,omitempty"`
	Lsndslitecursubscribers  int     `json:"lsndslitecursubscribers,omitempty"`
	/**
	* Current number of LSN DS-Lite subscribers.
	 */
	Lsndslitecursubscribersrate float64 `json:"lsndslitecursubscribersrate,omitempty"`
}
