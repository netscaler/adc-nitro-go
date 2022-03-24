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

type Lsnnat64stats struct {
	/**
	* Clear the statsistics / counters
	 */
	Clearstats          string `json:"clearstats,omitempty"`
	Lsncurnat64sessions int    `json:"lsncurnat64sessions,omitempty"`
	/**
	* Current number of LSN NAT64 sessions.
	 */
	Lsncurnat64sessionsrate float64 `json:"lsncurnat64sessionsrate,omitempty"`
	Lsnnat64cursubscribers  int     `json:"lsnnat64cursubscribers,omitempty"`
	/**
	* Current number of LSN NAT64 subscribers.
	 */
	Lsnnat64cursubscribersrate float64 `json:"lsnnat64cursubscribersrate,omitempty"`
	Lsntotnat64rxpkts          int     `json:"lsntotnat64rxpkts,omitempty"`
	/**
	* Total number of LSN NAT64 rx pkts.
	 */
	Lsnnat64rxpktsrate float64 `json:"lsnnat64rxpktsrate,omitempty"`
	Lsntotnat64rxbytes int     `json:"lsntotnat64rxbytes,omitempty"`
	/**
	* Total number of LSN NAT64 rx bytes.
	 */
	Lsnnat64rxbytesrate float64 `json:"lsnnat64rxbytesrate,omitempty"`
	Lsntotnat64txpkts   int     `json:"lsntotnat64txpkts,omitempty"`
	/**
	* Total number of LSN NAT64 tx pkts.
	 */
	Lsnnat64txpktsrate float64 `json:"lsnnat64txpktsrate,omitempty"`
	Lsntotnat64txbytes int     `json:"lsntotnat64txbytes,omitempty"`
	/**
	* Total number of LSN NAT64 tx bytes.
	 */
	Lsnnat64txbytesrate    float64 `json:"lsnnat64txbytesrate,omitempty"`
	Lsncurnat64tcpsessions int     `json:"lsncurnat64tcpsessions,omitempty"`
	/**
	* Number of LSN NAT64 TCP Current Sessions.
	 */
	Lsncurnat64tcpsessionsrate float64 `json:"lsncurnat64tcpsessionsrate,omitempty"`
	Lsntotnat64tcprxpkts       int     `json:"lsntotnat64tcprxpkts,omitempty"`
	/**
	* Number of LSN NAT64 TCP Received packets.
	 */
	Lsnnat64tcprxpktsrate float64 `json:"lsnnat64tcprxpktsrate,omitempty"`
	Lsntotnat64tcprxbytes int     `json:"lsntotnat64tcprxbytes,omitempty"`
	/**
	* Number of LSN NAT64 TCP Received bytes.
	 */
	Lsnnat64tcprxbytesrate float64 `json:"lsnnat64tcprxbytesrate,omitempty"`
	Lsntotnat64tcptxpkts   int     `json:"lsntotnat64tcptxpkts,omitempty"`
	/**
	* Number of LSN NAT64 TCP Transmitted packets.
	 */
	Lsnnat64tcptxpktsrate float64 `json:"lsnnat64tcptxpktsrate,omitempty"`
	Lsntotnat64tcptxbytes int     `json:"lsntotnat64tcptxbytes,omitempty"`
	/**
	* Number of LSN NAT64 TCP Transmitted bytes.
	 */
	Lsnnat64tcptxbytesrate float64 `json:"lsnnat64tcptxbytesrate,omitempty"`
	Lsntotnat64tcpdrppkts  int     `json:"lsntotnat64tcpdrppkts,omitempty"`
	/**
	* Number of LSN NAT64 TCP Dropped packets.
	 */
	Lsnnat64tcpdrppktsrate float64 `json:"lsnnat64tcpdrppktsrate,omitempty"`
	Lsncurnat64udpsessions int     `json:"lsncurnat64udpsessions,omitempty"`
	/**
	* Number of LSN NAT64 UDP Current Sessions.
	 */
	Lsncurnat64udpsessionsrate float64 `json:"lsncurnat64udpsessionsrate,omitempty"`
	Lsntotnat64udprxpkts       int     `json:"lsntotnat64udprxpkts,omitempty"`
	/**
	* Number of LSN NAT64 UDP Received packets.
	 */
	Lsnnat64udprxpktsrate float64 `json:"lsnnat64udprxpktsrate,omitempty"`
	Lsntotnat64udprxbytes int     `json:"lsntotnat64udprxbytes,omitempty"`
	/**
	* Number of LSN NAT64 UDP Received bytes.
	 */
	Lsnnat64udprxbytesrate float64 `json:"lsnnat64udprxbytesrate,omitempty"`
	Lsntotnat64udptxpkts   int     `json:"lsntotnat64udptxpkts,omitempty"`
	/**
	* Number of LSN NAT64 UDP Transmitted packets.
	 */
	Lsnnat64udptxpktsrate float64 `json:"lsnnat64udptxpktsrate,omitempty"`
	Lsntotnat64udptxbytes int     `json:"lsntotnat64udptxbytes,omitempty"`
	/**
	* Number of LSN NAT64 UDP Transmitted bytes.
	 */
	Lsnnat64udptxbytesrate float64 `json:"lsnnat64udptxbytesrate,omitempty"`
	Lsntotnat64udpdrppkts  int     `json:"lsntotnat64udpdrppkts,omitempty"`
	/**
	* Number of LSN NAT64 UDP Dropped packets.
	 */
	Lsnnat64udpdrppktsrate  float64 `json:"lsnnat64udpdrppktsrate,omitempty"`
	Lsncurnat64icmpsessions int     `json:"lsncurnat64icmpsessions,omitempty"`
	/**
	* Number of LSN NAT64 ICMP Current Sessions.
	 */
	Lsncurnat64icmpsessionsrate float64 `json:"lsncurnat64icmpsessionsrate,omitempty"`
	Lsntotnat64icmprxpkts       int     `json:"lsntotnat64icmprxpkts,omitempty"`
	/**
	* Number of LSN NAT64 ICMP Received packets.
	 */
	Lsnnat64icmprxpktsrate float64 `json:"lsnnat64icmprxpktsrate,omitempty"`
	Lsntotnat64icmprxbytes int     `json:"lsntotnat64icmprxbytes,omitempty"`
	/**
	* Number of LSN NAT64 ICMP Received bytes.
	 */
	Lsnnat64icmprxbytesrate float64 `json:"lsnnat64icmprxbytesrate,omitempty"`
	Lsntotnat64icmptxpkts   int     `json:"lsntotnat64icmptxpkts,omitempty"`
	/**
	* Number of LSN NAT64 ICMP Transmitted packets.
	 */
	Lsnnat64icmptxpktsrate float64 `json:"lsnnat64icmptxpktsrate,omitempty"`
	Lsntotnat64icmptxbytes int     `json:"lsntotnat64icmptxbytes,omitempty"`
	/**
	* Number of LSN NAT64 ICMP Transmitted bytes.
	 */
	Lsnnat64icmptxbytesrate float64 `json:"lsnnat64icmptxbytesrate,omitempty"`
	Lsntotnat64icmpdrppkts  int     `json:"lsntotnat64icmpdrppkts,omitempty"`
	/**
	* Number of LSN NAT64 ICMP Dropped packets.
	 */
	Lsnnat64icmpdrppktsrate float64 `json:"lsnnat64icmpdrppktsrate,omitempty"`
}
