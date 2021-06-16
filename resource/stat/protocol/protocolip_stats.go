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

package protocol

/**
* Statistics for protocolip resource.
*/

type Protocolipstats struct {
	/**
	* Clear the statsistics / counters
	*/
	Clearstats string `json:"clearstats,omitempty"`
	Iptotrxpkts uint64 `json:"iptotrxpkts,omitempty"`
	/**
	* IP packets received.
	*/
	Iprxpktsrate float64 `json:"iprxpktsrate,omitempty"`
	Iptotrxbytes uint64 `json:"iptotrxbytes,omitempty"`
	/**
	* Bytes of IP data received.
	*/
	Iprxbytesrate float64 `json:"iprxbytesrate,omitempty"`
	Iptottxpkts uint64 `json:"iptottxpkts,omitempty"`
	/**
	* IP packets transmitted.
	*/
	Iptxpktsrate float64 `json:"iptxpktsrate,omitempty"`
	Iptottxbytes uint64 `json:"iptottxbytes,omitempty"`
	/**
	* Bytes of IP data transmitted.
	*/
	Iptxbytesrate float64 `json:"iptxbytesrate,omitempty"`
	Iptotrxmbits uint64 `json:"iptotrxmbits,omitempty"`
	/**
	* Megabits of IP data received.
	*/
	Iprxmbitsrate float64 `json:"iprxmbitsrate,omitempty"`
	Iptottxmbits uint64 `json:"iptottxmbits,omitempty"`
	/**
	* Megabits of IP data transmitted.
	*/
	Iptxmbitsrate float64 `json:"iptxmbitsrate,omitempty"`
	Iptotroutedpkts uint64 `json:"iptotroutedpkts,omitempty"`
	/**
	* Total routed packets.
	*/
	Iproutedpktsrate float64 `json:"iproutedpktsrate,omitempty"`
	Iptotroutedmbits uint64 `json:"iptotroutedmbits,omitempty"`
	/**
	* Total routed Mbits
	*/
	Iproutedmbitsrate float64 `json:"iproutedmbitsrate,omitempty"`
	Iptotfragments uint64 `json:"iptotfragments,omitempty"`
	Iptotsuccreassembly uint64 `json:"iptotsuccreassembly,omitempty"`
	Iptotreassemblyattempt uint64 `json:"iptotreassemblyattempt,omitempty"`
	Iptotaddrlookup uint64 `json:"iptotaddrlookup,omitempty"`
	Iptotaddrlookupfail uint64 `json:"iptotaddrlookupfail,omitempty"`
	Iptotudpfragmentsfwd uint64 `json:"iptotudpfragmentsfwd,omitempty"`
	Iptottcpfragmentsfwd uint64 `json:"iptottcpfragmentsfwd,omitempty"`
	Iptotfragpktsgen uint64 `json:"iptotfragpktsgen,omitempty"`
	Iptotbadchecksums uint64 `json:"iptotbadchecksums,omitempty"`
	Iptotunsuccreassembly uint64 `json:"iptotunsuccreassembly,omitempty"`
	Iptottoobig uint64 `json:"iptottoobig,omitempty"`
	Iptotzerofragmentlen uint64 `json:"iptotzerofragmentlen,omitempty"`
	Iptotdupfragments uint64 `json:"iptotdupfragments,omitempty"`
	Iptotoutoforderfrag uint64 `json:"iptotoutoforderfrag,omitempty"`
	Iptotunknowndstrcvd uint64 `json:"iptotunknowndstrcvd,omitempty"`
	Iptotbadtransport uint64 `json:"iptotbadtransport,omitempty"`
	Iptotvipdown uint64 `json:"iptotvipdown,omitempty"`
	Iptotfixheaderfail uint64 `json:"iptotfixheaderfail,omitempty"`
	Iptotttlexpired uint64 `json:"iptotttlexpired,omitempty"`
	Iptotmaxclients uint64 `json:"iptotmaxclients,omitempty"`
	Iptotunknownsvcs uint64 `json:"iptotunknownsvcs,omitempty"`
	Iptotlandattacks uint64 `json:"iptotlandattacks,omitempty"`
	Iptotinvalidheadersz uint64 `json:"iptotinvalidheadersz,omitempty"`
	Iptotinvalidpacketsize uint64 `json:"iptotinvalidpacketsize,omitempty"`
	Iptottruncatedpackets uint64 `json:"iptottruncatedpackets,omitempty"`
	Noniptottruncatedpackets uint64 `json:"noniptottruncatedpackets,omitempty"`
	Iptotzeronexthop uint64 `json:"iptotzeronexthop,omitempty"`
	Iptotbadlens uint64 `json:"iptotbadlens,omitempty"`
	Iptotbadmacaddrs uint64 `json:"iptotbadmacaddrs,omitempty"`

}
