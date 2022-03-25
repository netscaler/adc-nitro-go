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

package stat

type Lsnstats struct {
	/**
	* Clear the statsistics / counters
	 */
	Clearstats      string `json:"clearstats,omitempty"`
	Lsntottcprxpkts int    `json:"lsntottcprxpkts,omitempty"`
	/**
	* Number of LSN TCP Received packets.
	 */
	Lsntcprxpktsrate float64 `json:"lsntcprxpktsrate,omitempty"`
	Lsntottcprxbytes int     `json:"lsntottcprxbytes,omitempty"`
	/**
	* Number of LSN TCP Received bytes.
	 */
	Lsntcprxbytesrate float64 `json:"lsntcprxbytesrate,omitempty"`
	Lsntottcptxpkts   int     `json:"lsntottcptxpkts,omitempty"`
	/**
	* Number of LSN TCP Transmitted packets.
	 */
	Lsntcptxpktsrate float64 `json:"lsntcptxpktsrate,omitempty"`
	Lsntottcptxbytes int     `json:"lsntottcptxbytes,omitempty"`
	/**
	* Number of LSN TCP Transmitted bytes.
	 */
	Lsntcptxbytesrate float64 `json:"lsntcptxbytesrate,omitempty"`
	Lsntottcpdrppkts  int     `json:"lsntottcpdrppkts,omitempty"`
	/**
	* Number of LSN TCP Dropped packets.
	 */
	Lsntcpdrppktsrate float64 `json:"lsntcpdrppktsrate,omitempty"`
	Lsncurtcpsessions int     `json:"lsncurtcpsessions,omitempty"`
	/**
	* Number of LSN TCP Current Sessions.
	 */
	Lsncurtcpsessionsrate float64 `json:"lsncurtcpsessionsrate,omitempty"`
	Lsntotudprxpkts       int     `json:"lsntotudprxpkts,omitempty"`
	/**
	* Number of LSN UDP Received packets.
	 */
	Lsnudprxpktsrate float64 `json:"lsnudprxpktsrate,omitempty"`
	Lsntotudprxbytes int     `json:"lsntotudprxbytes,omitempty"`
	/**
	* Number of LSN UDP Received bytes.
	 */
	Lsnudprxbytesrate float64 `json:"lsnudprxbytesrate,omitempty"`
	Lsntotudptxpkts   int     `json:"lsntotudptxpkts,omitempty"`
	/**
	* Number of LSN UDP Transmitted packets.
	 */
	Lsnudptxpktsrate float64 `json:"lsnudptxpktsrate,omitempty"`
	Lsntotudptxbytes int     `json:"lsntotudptxbytes,omitempty"`
	/**
	* Number of LSN UDP Transmitted bytes.
	 */
	Lsnudptxbytesrate float64 `json:"lsnudptxbytesrate,omitempty"`
	Lsntotudpdrppkts  int     `json:"lsntotudpdrppkts,omitempty"`
	/**
	* Number of LSN UDP Dropped packets.
	 */
	Lsnudpdrppktsrate float64 `json:"lsnudpdrppktsrate,omitempty"`
	Lsncurudpsessions int     `json:"lsncurudpsessions,omitempty"`
	/**
	* Number of LSN UDP Current Sessions.
	 */
	Lsncurudpsessionsrate float64 `json:"lsncurudpsessionsrate,omitempty"`
	Lsntoticmprxpkts      int     `json:"lsntoticmprxpkts,omitempty"`
	/**
	* Number of LSN ICMP Received packets.
	 */
	Lsnicmprxpktsrate float64 `json:"lsnicmprxpktsrate,omitempty"`
	Lsntoticmprxbytes int     `json:"lsntoticmprxbytes,omitempty"`
	/**
	* Number of LSN ICMP Received bytes.
	 */
	Lsnicmprxbytesrate float64 `json:"lsnicmprxbytesrate,omitempty"`
	Lsntoticmptxpkts   int     `json:"lsntoticmptxpkts,omitempty"`
	/**
	* Number of LSN ICMP Transmitted packets.
	 */
	Lsnicmptxpktsrate float64 `json:"lsnicmptxpktsrate,omitempty"`
	Lsntoticmptxbytes int     `json:"lsntoticmptxbytes,omitempty"`
	/**
	* Number of LSN ICMP Transmitted bytes.
	 */
	Lsnicmptxbytesrate float64 `json:"lsnicmptxbytesrate,omitempty"`
	Lsntoticmpdrppkts  int     `json:"lsntoticmpdrppkts,omitempty"`
	/**
	* Number of LSN ICMP Dropped packets.
	 */
	Lsnicmpdrppktsrate float64 `json:"lsnicmpdrppktsrate,omitempty"`
	Lsncuricmpsessions int     `json:"lsncuricmpsessions,omitempty"`
	/**
	* Number of LSN ICMP Current Sessions.
	 */
	Lsncuricmpsessionsrate float64 `json:"lsncuricmpsessionsrate,omitempty"`
	Lsncursessions         int     `json:"lsncursessions,omitempty"`
	/**
	* Current number of LSN sessions.
	 */
	Lsncursessionsrate float64 `json:"lsncursessionsrate,omitempty"`
	Lsncursubscribers  int     `json:"lsncursubscribers,omitempty"`
	/**
	* Current number of LSN subscribers.
	 */
	Lsncursubscribersrate float64 `json:"lsncursubscribersrate,omitempty"`
}
