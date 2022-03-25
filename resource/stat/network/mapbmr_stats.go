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

package network

/**
* Statistics for MAP-T Basic Mapping rule resource.
 */

type Mapbmrstats struct {
	/**
	* Basic Mapping Rule name.
	 */
	Name string `json:"name,omitempty"`
	/**
	* Clear the statsistics / counters
	 */
	Clearstats     string `json:"clearstats,omitempty"`
	Bmrtotv6rxpkts int    `json:"bmrtotv6rxpkts,omitempty"`
	/**
	* Total number of MAP-T BMR V6 Recieved packets.
	 */
	Bmrv6rxpktsrate float64 `json:"bmrv6rxpktsrate,omitempty"`
	Bmrtotv6txpkts  int     `json:"bmrtotv6txpkts,omitempty"`
	/**
	* Total number of MAP-T BMR V6 Transmitted packets.
	 */
	Bmrv6txpktsrate float64 `json:"bmrv6txpktsrate,omitempty"`
	Bmrtotv4rxpkts  int     `json:"bmrtotv4rxpkts,omitempty"`
	/**
	* Total number of MAP-T BMR V4 Recieved packets.
	 */
	Bmrv4rxpktsrate float64 `json:"bmrv4rxpktsrate,omitempty"`
	Bmrtotv4txpkts  int     `json:"bmrtotv4txpkts,omitempty"`
	/**
	* Total number of MAP-T BMR V4 Transmitted packets.
	 */
	Bmrv4txpktsrate   float64 `json:"bmrv4txpktsrate,omitempty"`
	Bmrtotv6rxpktstcp int     `json:"bmrtotv6rxpktstcp,omitempty"`
	/**
	* Total number of MAP-T BMR V6 TCP Recieved packets.
	 */
	Bmrv6rxpktstcprate float64 `json:"bmrv6rxpktstcprate,omitempty"`
	Bmrtotv6txpktstcp  int     `json:"bmrtotv6txpktstcp,omitempty"`
	/**
	* Total number of MAP-T BMR V6 TCP Transmitted packets.
	 */
	Bmrv6txpktstcprate float64 `json:"bmrv6txpktstcprate,omitempty"`
	Bmrtotv4rxpktstcp  int     `json:"bmrtotv4rxpktstcp,omitempty"`
	/**
	* Total number of MAP-T BMR V4 TCP Recieved packets.
	 */
	Bmrv4rxpktstcprate float64 `json:"bmrv4rxpktstcprate,omitempty"`
	Bmrtotv4txpktstcp  int     `json:"bmrtotv4txpktstcp,omitempty"`
	/**
	* Total number of MAP-T BMR V4 TCP Transmitted packets.
	 */
	Bmrv4txpktstcprate float64 `json:"bmrv4txpktstcprate,omitempty"`
	Bmrtotv6rxpktsudp  int     `json:"bmrtotv6rxpktsudp,omitempty"`
	/**
	* Total number of MAP-T BMR V6 UDP Recieved packets.
	 */
	Bmrv6rxpktsudprate float64 `json:"bmrv6rxpktsudprate,omitempty"`
	Bmrtotv6txpktsudp  int     `json:"bmrtotv6txpktsudp,omitempty"`
	/**
	* Total number of MAP-T BMR V6 UDP Transmitted packets.
	 */
	Bmrv6txpktsudprate float64 `json:"bmrv6txpktsudprate,omitempty"`
	Bmrtotv4rxpktsudp  int     `json:"bmrtotv4rxpktsudp,omitempty"`
	/**
	* Total number of MAP-T BMR V4 UDP Recieved packets.
	 */
	Bmrv4rxpktsudprate float64 `json:"bmrv4rxpktsudprate,omitempty"`
	Bmrtotv4txpktsudp  int     `json:"bmrtotv4txpktsudp,omitempty"`
	/**
	* Total number of MAP-T BMR V4 UDP Transmitted packets.
	 */
	Bmrv4txpktsudprate float64 `json:"bmrv4txpktsudprate,omitempty"`
	Bmrtotv6rxpktsicmp int     `json:"bmrtotv6rxpktsicmp,omitempty"`
	/**
	* Total number of MAP-T BMR V6 ICMP Recieved packets.
	 */
	Bmrv6rxpktsicmprate float64 `json:"bmrv6rxpktsicmprate,omitempty"`
	Bmrtotv6txpktsicmp  int     `json:"bmrtotv6txpktsicmp,omitempty"`
	/**
	* Total number of MAP-T BMR V6 ICMP Transmitted packets.
	 */
	Bmrv6txpktsicmprate float64 `json:"bmrv6txpktsicmprate,omitempty"`
	Bmrtotv4rxpktsicmp  int     `json:"bmrtotv4rxpktsicmp,omitempty"`
	/**
	* Total number of MAP-T BMR V4 ICMP Recieved packets.
	 */
	Bmrv4rxpktsicmprate float64 `json:"bmrv4rxpktsicmprate,omitempty"`
	Bmrtotv4txpktsicmp  int     `json:"bmrtotv4txpktsicmp,omitempty"`
	/**
	* Total number of MAP-T BMR V4 ICMP Transmitted packets.
	 */
	Bmrv4txpktsicmprate float64 `json:"bmrv4txpktsicmprate,omitempty"`
}
