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

/**
* Statistics for MAP-T Map Domain resource.
 */

type Mapdomainstats struct {
	/**
	* An integer specifying the VLAN identification number (VID). Possible values: 1 through 4094.
	 */
	Name string `json:"name,omitempty"`
	/**
	* Clear the statsistics / counters
	 */
	Clearstats        string `json:"clearstats,omitempty"`
	Maptotv6rxpktstcp int    `json:"maptotv6rxpktstcp,omitempty"`
	/**
	* Total number of MAP-T IPv6 TCP Recieved packets.
	 */
	Mapv6rxpktstcprate float64 `json:"mapv6rxpktstcprate,omitempty"`
	Maptotv6txpktstcp  int     `json:"maptotv6txpktstcp,omitempty"`
	/**
	* Total number of MAP-T IPv6 TCP Transmitted packets.
	 */
	Mapv6txpktstcprate float64 `json:"mapv6txpktstcprate,omitempty"`
	Maptotv6rxbytestcp int     `json:"maptotv6rxbytestcp,omitempty"`
	/**
	* Total number of MAP-T IPv6 TCP Recieved Bytes.
	 */
	Mapv6rxbytestcprate float64 `json:"mapv6rxbytestcprate,omitempty"`
	Maptotv6txbytestcp  int     `json:"maptotv6txbytestcp,omitempty"`
	/**
	* Total number of MAP-T IPv6 TCP Transmitted Bytes.
	 */
	Mapv6txbytestcprate float64 `json:"mapv6txbytestcprate,omitempty"`
	Maptotv4rxpktstcp   int     `json:"maptotv4rxpktstcp,omitempty"`
	/**
	* Total number of MAP-T IPv4 TCP Recieved packets.
	 */
	Mapv4rxpktstcprate float64 `json:"mapv4rxpktstcprate,omitempty"`
	Maptotv4txpktstcp  int     `json:"maptotv4txpktstcp,omitempty"`
	/**
	* Total number of MAP-T IPv4 TCP Transmitted packets.
	 */
	Mapv4txpktstcprate float64 `json:"mapv4txpktstcprate,omitempty"`
	Maptotv4rxbytestcp int     `json:"maptotv4rxbytestcp,omitempty"`
	/**
	* Total number of MAP-T IPv4 TCP Recieved Bytes.
	 */
	Mapv4rxbytestcprate float64 `json:"mapv4rxbytestcprate,omitempty"`
	Maptotv4txbytestcp  int     `json:"maptotv4txbytestcp,omitempty"`
	/**
	* Total number of MAP-T IPv4 TCP Transmitted Bytes.
	 */
	Mapv4txbytestcprate float64 `json:"mapv4txbytestcprate,omitempty"`
	Maptotv6rxpktsudp   int     `json:"maptotv6rxpktsudp,omitempty"`
	/**
	* Total number of MAP-T IPv6 UDP Recieved packets.
	 */
	Mapv6rxpktsudprate float64 `json:"mapv6rxpktsudprate,omitempty"`
	Maptotv6txpktsudp  int     `json:"maptotv6txpktsudp,omitempty"`
	/**
	* Total number of MAP-T IPv6 UDP Transmitted packets.
	 */
	Mapv6txpktsudprate float64 `json:"mapv6txpktsudprate,omitempty"`
	Maptotv6rxbytesudp int     `json:"maptotv6rxbytesudp,omitempty"`
	/**
	* Total number of MAP-T IPv6 UDP Recieved Bytes.
	 */
	Mapv6rxbytesudprate float64 `json:"mapv6rxbytesudprate,omitempty"`
	Maptotv6txbytesudp  int     `json:"maptotv6txbytesudp,omitempty"`
	/**
	* Total number of MAP-T IPv6 UDP Transmitted Bytes.
	 */
	Mapv6txbytesudprate float64 `json:"mapv6txbytesudprate,omitempty"`
	Maptotv4rxpktsudp   int     `json:"maptotv4rxpktsudp,omitempty"`
	/**
	* Total number of MAP-T IPv4 UDP Recieved packets.
	 */
	Mapv4rxpktsudprate float64 `json:"mapv4rxpktsudprate,omitempty"`
	Maptotv4txpktsudp  int     `json:"maptotv4txpktsudp,omitempty"`
	/**
	* Total number of MAP-T IPv4 UDP Transmitted packets.
	 */
	Mapv4txpktsudprate float64 `json:"mapv4txpktsudprate,omitempty"`
	Maptotv4rxbytesudp int     `json:"maptotv4rxbytesudp,omitempty"`
	/**
	* Total number of MAP-T IPv4 UDP Recieved Bytes.
	 */
	Mapv4rxbytesudprate float64 `json:"mapv4rxbytesudprate,omitempty"`
	Maptotv4txbytesudp  int     `json:"maptotv4txbytesudp,omitempty"`
	/**
	* Total number of MAP-T IPv4 UDP Transmitted Bytes.
	 */
	Mapv4txbytesudprate float64 `json:"mapv4txbytesudprate,omitempty"`
	Maptotv6rxpktsicmp  int     `json:"maptotv6rxpktsicmp,omitempty"`
	/**
	* Total number of MAP-T IPv6 ICMP Recieved packets.
	 */
	Mapv6rxpktsicmprate float64 `json:"mapv6rxpktsicmprate,omitempty"`
	Maptotv6txpktsicmp  int     `json:"maptotv6txpktsicmp,omitempty"`
	/**
	* Total number of MAP-T IPv6 ICMP Transmitted packets.
	 */
	Mapv6txpktsicmprate float64 `json:"mapv6txpktsicmprate,omitempty"`
	Maptotv6rxbytesicmp int     `json:"maptotv6rxbytesicmp,omitempty"`
	/**
	* Total number of MAP-T IPv6 ICMP Recieved Bytes.
	 */
	Mapv6rxbytesicmprate float64 `json:"mapv6rxbytesicmprate,omitempty"`
	Maptotv6txbytesicmp  int     `json:"maptotv6txbytesicmp,omitempty"`
	/**
	* Total number of MAP-T IPv6 ICMP Transmitted Bytes.
	 */
	Mapv6txbytesicmprate float64 `json:"mapv6txbytesicmprate,omitempty"`
	Maptotv4rxpktsicmp   int     `json:"maptotv4rxpktsicmp,omitempty"`
	/**
	* Total number of MAP-T IPv4 ICMP Recieved packets.
	 */
	Mapv4rxpktsicmprate float64 `json:"mapv4rxpktsicmprate,omitempty"`
	Maptotv4txpktsicmp  int     `json:"maptotv4txpktsicmp,omitempty"`
	/**
	* Total number of MAP-T IPv4 ICMP Transmitted packets.
	 */
	Mapv4txpktsicmprate float64 `json:"mapv4txpktsicmprate,omitempty"`
	Maptotv4rxbytesicmp int     `json:"maptotv4rxbytesicmp,omitempty"`
	/**
	* Total number of MAP-T IPv4 ICMP Recieved Bytes.
	 */
	Mapv4rxbytesicmprate float64 `json:"mapv4rxbytesicmprate,omitempty"`
	Maptotv4txbytesicmp  int     `json:"maptotv4txbytesicmp,omitempty"`
	/**
	* Total number of MAP-T IPv4 ICMP Transmitted Bytes.
	 */
	Mapv4txbytesicmprate float64 `json:"mapv4txbytesicmprate,omitempty"`
	Maptotv6rxpkts       int     `json:"maptotv6rxpkts,omitempty"`
	/**
	* Total number of MAP-T V6 Recieved packets.
	 */
	Mapv6rxpktsrate float64 `json:"mapv6rxpktsrate,omitempty"`
	Maptotv6txpkts  int     `json:"maptotv6txpkts,omitempty"`
	/**
	* Total number of MAP-T V6 Transmitted packets.
	 */
	Mapv6txpktsrate float64 `json:"mapv6txpktsrate,omitempty"`
	Maptotv4rxpkts  int     `json:"maptotv4rxpkts,omitempty"`
	/**
	* Total number of MAP-T V4 Recieved packets.
	 */
	Mapv4rxpktsrate float64 `json:"mapv4rxpktsrate,omitempty"`
	Maptotv4txpkts  int     `json:"maptotv4txpkts,omitempty"`
	/**
	* Total number of MAP-T V4 Transmitted packets.
	 */
	Mapv4txpktsrate float64 `json:"mapv4txpktsrate,omitempty"`
}
