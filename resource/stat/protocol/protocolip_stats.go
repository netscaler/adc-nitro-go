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
	Iptotrxpkts int `json:"iptotrxpkts,omitempty"`
	/**
	* IP packets received.
	*/
	Iprxpktsrate float64 `json:"iprxpktsrate,omitempty"`
	Iptotrxbytes int `json:"iptotrxbytes,omitempty"`
	/**
	* Bytes of IP data received.
	*/
	Iprxbytesrate float64 `json:"iprxbytesrate,omitempty"`
	Iptottxpkts int `json:"iptottxpkts,omitempty"`
	/**
	* IP packets transmitted.
	*/
	Iptxpktsrate float64 `json:"iptxpktsrate,omitempty"`
	Iptottxbytes int `json:"iptottxbytes,omitempty"`
	/**
	* Bytes of IP data transmitted.
	*/
	Iptxbytesrate float64 `json:"iptxbytesrate,omitempty"`
	Iptotrxmbits int `json:"iptotrxmbits,omitempty"`
	/**
	* Megabits of IP data received.
	*/
	Iprxmbitsrate float64 `json:"iprxmbitsrate,omitempty"`
	Iptottxmbits int `json:"iptottxmbits,omitempty"`
	/**
	* Megabits of IP data transmitted.
	*/
	Iptxmbitsrate float64 `json:"iptxmbitsrate,omitempty"`
	Iptotroutedpkts int `json:"iptotroutedpkts,omitempty"`
	/**
	* Total routed packets.
	*/
	Iproutedpktsrate float64 `json:"iproutedpktsrate,omitempty"`
	Iptotroutedmbits int `json:"iptotroutedmbits,omitempty"`
	/**
	* Total routed Mbits
	*/
	Iproutedmbitsrate float64 `json:"iproutedmbitsrate,omitempty"`
	Iptotfragments int `json:"iptotfragments,omitempty"`
	Iptotsuccreassembly int `json:"iptotsuccreassembly,omitempty"`
	Iptotreassemblyattempt int `json:"iptotreassemblyattempt,omitempty"`
	Iptotaddrlookup int `json:"iptotaddrlookup,omitempty"`
	Iptotaddrlookupfail int `json:"iptotaddrlookupfail,omitempty"`
	Iptotudpfragmentsfwd int `json:"iptotudpfragmentsfwd,omitempty"`
	Iptottcpfragmentsfwd int `json:"iptottcpfragmentsfwd,omitempty"`
	Iptotfragpktsgen int `json:"iptotfragpktsgen,omitempty"`
	Iptotbadchecksums int `json:"iptotbadchecksums,omitempty"`
	Iptotunsuccreassembly int `json:"iptotunsuccreassembly,omitempty"`
	Iptottoobig int `json:"iptottoobig,omitempty"`
	Iptotzerofragmentlen int `json:"iptotzerofragmentlen,omitempty"`
	Iptotdupfragments int `json:"iptotdupfragments,omitempty"`
	Iptotoutoforderfrag int `json:"iptotoutoforderfrag,omitempty"`
	Iptotunknowndstrcvd int `json:"iptotunknowndstrcvd,omitempty"`
	Iptotbadtransport int `json:"iptotbadtransport,omitempty"`
	Iptotvipdown int `json:"iptotvipdown,omitempty"`
	Iptotfixheaderfail int `json:"iptotfixheaderfail,omitempty"`
	Iptotttlexpired int `json:"iptotttlexpired,omitempty"`
	Iptotmaxclients int `json:"iptotmaxclients,omitempty"`
	Iptotunknownsvcs int `json:"iptotunknownsvcs,omitempty"`
	Iptotlandattacks int `json:"iptotlandattacks,omitempty"`
	Iptotinvalidheadersz int `json:"iptotinvalidheadersz,omitempty"`
	Iptotinvalidpacketsize int `json:"iptotinvalidpacketsize,omitempty"`
	Iptottruncatedpackets int `json:"iptottruncatedpackets,omitempty"`
	Noniptottruncatedpackets int `json:"noniptottruncatedpackets,omitempty"`
	Iptotzeronexthop int `json:"iptotzeronexthop,omitempty"`
	Iptotbadlens int `json:"iptotbadlens,omitempty"`
	Iptotbadmacaddrs int `json:"iptotbadmacaddrs,omitempty"`

}
