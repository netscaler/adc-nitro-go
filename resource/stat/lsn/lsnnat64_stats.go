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
	Clearstats string `json:"clearstats,omitempty"`
	Lsncurnat64sessions uint64 `json:"lsncurnat64sessions,omitempty"`
	/**
	* Current number of LSN NAT64 sessions.
	*/
	Lsncurnat64sessionsrate int32 `json:"lsncurnat64sessionsrate,omitempty"`
	Lsnnat64cursubscribers uint64 `json:"lsnnat64cursubscribers,omitempty"`
	/**
	* Current number of LSN NAT64 subscribers.
	*/
	Lsnnat64cursubscribersrate int32 `json:"lsnnat64cursubscribersrate,omitempty"`
	Lsntotnat64rxpkts uint64 `json:"lsntotnat64rxpkts,omitempty"`
	/**
	* Total number of LSN NAT64 rx pkts.
	*/
	Lsnnat64rxpktsrate int32 `json:"lsnnat64rxpktsrate,omitempty"`
	Lsntotnat64rxbytes uint64 `json:"lsntotnat64rxbytes,omitempty"`
	/**
	* Total number of LSN NAT64 rx bytes.
	*/
	Lsnnat64rxbytesrate int32 `json:"lsnnat64rxbytesrate,omitempty"`
	Lsntotnat64txpkts uint64 `json:"lsntotnat64txpkts,omitempty"`
	/**
	* Total number of LSN NAT64 tx pkts.
	*/
	Lsnnat64txpktsrate int32 `json:"lsnnat64txpktsrate,omitempty"`
	Lsntotnat64txbytes uint64 `json:"lsntotnat64txbytes,omitempty"`
	/**
	* Total number of LSN NAT64 tx bytes.
	*/
	Lsnnat64txbytesrate int32 `json:"lsnnat64txbytesrate,omitempty"`
	Lsncurnat64tcpsessions uint64 `json:"lsncurnat64tcpsessions,omitempty"`
	/**
	* Number of LSN NAT64 TCP Current Sessions.
	*/
	Lsncurnat64tcpsessionsrate int32 `json:"lsncurnat64tcpsessionsrate,omitempty"`
	Lsntotnat64tcprxpkts uint64 `json:"lsntotnat64tcprxpkts,omitempty"`
	/**
	* Number of LSN NAT64 TCP Received packets.
	*/
	Lsnnat64tcprxpktsrate int32 `json:"lsnnat64tcprxpktsrate,omitempty"`
	Lsntotnat64tcprxbytes uint64 `json:"lsntotnat64tcprxbytes,omitempty"`
	/**
	* Number of LSN NAT64 TCP Received bytes.
	*/
	Lsnnat64tcprxbytesrate int32 `json:"lsnnat64tcprxbytesrate,omitempty"`
	Lsntotnat64tcptxpkts uint64 `json:"lsntotnat64tcptxpkts,omitempty"`
	/**
	* Number of LSN NAT64 TCP Transmitted packets.
	*/
	Lsnnat64tcptxpktsrate int32 `json:"lsnnat64tcptxpktsrate,omitempty"`
	Lsntotnat64tcptxbytes uint64 `json:"lsntotnat64tcptxbytes,omitempty"`
	/**
	* Number of LSN NAT64 TCP Transmitted bytes.
	*/
	Lsnnat64tcptxbytesrate int32 `json:"lsnnat64tcptxbytesrate,omitempty"`
	Lsntotnat64tcpdrppkts uint64 `json:"lsntotnat64tcpdrppkts,omitempty"`
	/**
	* Number of LSN NAT64 TCP Dropped packets.
	*/
	Lsnnat64tcpdrppktsrate int32 `json:"lsnnat64tcpdrppktsrate,omitempty"`
	Lsncurnat64udpsessions uint64 `json:"lsncurnat64udpsessions,omitempty"`
	/**
	* Number of LSN NAT64 UDP Current Sessions.
	*/
	Lsncurnat64udpsessionsrate int32 `json:"lsncurnat64udpsessionsrate,omitempty"`
	Lsntotnat64udprxpkts uint64 `json:"lsntotnat64udprxpkts,omitempty"`
	/**
	* Number of LSN NAT64 UDP Received packets.
	*/
	Lsnnat64udprxpktsrate int32 `json:"lsnnat64udprxpktsrate,omitempty"`
	Lsntotnat64udprxbytes uint64 `json:"lsntotnat64udprxbytes,omitempty"`
	/**
	* Number of LSN NAT64 UDP Received bytes.
	*/
	Lsnnat64udprxbytesrate int32 `json:"lsnnat64udprxbytesrate,omitempty"`
	Lsntotnat64udptxpkts uint64 `json:"lsntotnat64udptxpkts,omitempty"`
	/**
	* Number of LSN NAT64 UDP Transmitted packets.
	*/
	Lsnnat64udptxpktsrate int32 `json:"lsnnat64udptxpktsrate,omitempty"`
	Lsntotnat64udptxbytes uint64 `json:"lsntotnat64udptxbytes,omitempty"`
	/**
	* Number of LSN NAT64 UDP Transmitted bytes.
	*/
	Lsnnat64udptxbytesrate int32 `json:"lsnnat64udptxbytesrate,omitempty"`
	Lsntotnat64udpdrppkts uint64 `json:"lsntotnat64udpdrppkts,omitempty"`
	/**
	* Number of LSN NAT64 UDP Dropped packets.
	*/
	Lsnnat64udpdrppktsrate int32 `json:"lsnnat64udpdrppktsrate,omitempty"`
	Lsncurnat64icmpsessions uint64 `json:"lsncurnat64icmpsessions,omitempty"`
	/**
	* Number of LSN NAT64 ICMP Current Sessions.
	*/
	Lsncurnat64icmpsessionsrate int32 `json:"lsncurnat64icmpsessionsrate,omitempty"`
	Lsntotnat64icmprxpkts uint64 `json:"lsntotnat64icmprxpkts,omitempty"`
	/**
	* Number of LSN NAT64 ICMP Received packets.
	*/
	Lsnnat64icmprxpktsrate int32 `json:"lsnnat64icmprxpktsrate,omitempty"`
	Lsntotnat64icmprxbytes uint64 `json:"lsntotnat64icmprxbytes,omitempty"`
	/**
	* Number of LSN NAT64 ICMP Received bytes.
	*/
	Lsnnat64icmprxbytesrate int32 `json:"lsnnat64icmprxbytesrate,omitempty"`
	Lsntotnat64icmptxpkts uint64 `json:"lsntotnat64icmptxpkts,omitempty"`
	/**
	* Number of LSN NAT64 ICMP Transmitted packets.
	*/
	Lsnnat64icmptxpktsrate int32 `json:"lsnnat64icmptxpktsrate,omitempty"`
	Lsntotnat64icmptxbytes uint64 `json:"lsntotnat64icmptxbytes,omitempty"`
	/**
	* Number of LSN NAT64 ICMP Transmitted bytes.
	*/
	Lsnnat64icmptxbytesrate int32 `json:"lsnnat64icmptxbytesrate,omitempty"`
	Lsntotnat64icmpdrppkts uint64 `json:"lsntotnat64icmpdrppkts,omitempty"`
	/**
	* Number of LSN NAT64 ICMP Dropped packets.
	*/
	Lsnnat64icmpdrppktsrate int32 `json:"lsnnat64icmpdrppktsrate,omitempty"`

}
