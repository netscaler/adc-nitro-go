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
* Statistics for ip v6 resource.
*/

type Protocolipv6stats struct {
	/**
	* Clear the statsistics / counters
	*/
	Clearstats string `json:"clearstats,omitempty"`
	Ipv6totrxpkts uint64 `json:"ipv6totrxpkts,omitempty"`
	/**
	* IPv6 packets received.
	*/
	Ipv6rxpktsrate float64 `json:"ipv6rxpktsrate,omitempty"`
	Ipv6totrxbytes uint64 `json:"ipv6totrxbytes,omitempty"`
	/**
	* Bytes of IPv6 data received.
	*/
	Ipv6rxbytesrate float64 `json:"ipv6rxbytesrate,omitempty"`
	Ipv6tottxpkts uint64 `json:"ipv6tottxpkts,omitempty"`
	/**
	* IPv6 packets transmitted
	*/
	Ipv6txpktsrate float64 `json:"ipv6txpktsrate,omitempty"`
	Ipv6tottxbytes uint64 `json:"ipv6tottxbytes,omitempty"`
	/**
	* Bytes of IPv6 data transmitted.
	*/
	Ipv6txbytesrate float64 `json:"ipv6txbytesrate,omitempty"`
	Ipv6totroutepkts uint64 `json:"ipv6totroutepkts,omitempty"`
	/**
	* IPv6 packets routed.
	*/
	Ipv6routepktsrate float64 `json:"ipv6routepktsrate,omitempty"`
	Ipv6totroutembits uint64 `json:"ipv6totroutembits,omitempty"`
	/**
	* IPv6 total Mbits.
	*/
	Ipv6routembitsrate float64 `json:"ipv6routembitsrate,omitempty"`
	Ipv6fragtotrxpkts uint64 `json:"ipv6fragtotrxpkts,omitempty"`
	/**
	* IPv6 fragments received.
	*/
	Ipv6fragrxpktsrate float64 `json:"ipv6fragrxpktsrate,omitempty"`
	Ipv6fragtottcpreassembled uint64 `json:"ipv6fragtottcpreassembled,omitempty"`
	/**
	* TCP fragments processed after reassembly.
	*/
	Ipv6fragtcpreassembledrate float64 `json:"ipv6fragtcpreassembledrate,omitempty"`
	Ipv6fragtotudpreassembled uint64 `json:"ipv6fragtotudpreassembled,omitempty"`
	/**
	* UDP fragments processed after reassembly.
	*/
	Ipv6fragudpreassembledrate float64 `json:"ipv6fragudpreassembledrate,omitempty"`
	Ipv6fragtotpktsprocessnoreass uint64 `json:"ipv6fragtotpktsprocessnoreass,omitempty"`
	/**
	* IPv6 fragments processed without reassembly.
	*/
	Ipv6fragpktsprocessnoreassrate float64 `json:"ipv6fragpktsprocessnoreassrate,omitempty"`
	Ipv6fragtotpktsforward uint64 `json:"ipv6fragtotpktsforward,omitempty"`
	/**
	* IPv6 fragments forwarded to the client or server without reassembly.
	*/
	Ipv6fragpktsforwardrate float64 `json:"ipv6fragpktsforwardrate,omitempty"`
	Ipv6errhdr uint64 `json:"ipv6errhdr,omitempty"`
	Ipv6nextheadererr uint64 `json:"ipv6nextheadererr,omitempty"`
	Ipv6landattack uint64 `json:"ipv6landattack,omitempty"`
	Ipv6fragpkttoobig uint64 `json:"ipv6fragpkttoobig,omitempty"`
	Ipv6fragzerolenpkt uint64 `json:"ipv6fragzerolenpkt,omitempty"`
	Ipv6toticmpnarx uint64 `json:"ipv6toticmpnarx,omitempty"`
	/**
	* Number of ICMPv6 NA packets received by Citrix ADC (OBSOLETE).
	*/
	Ipv6icmpnarxrate float64 `json:"ipv6icmpnarxrate,omitempty"`
	Ipv6toticmpnsrx uint64 `json:"ipv6toticmpnsrx,omitempty"`
	/**
	* Number of ICMPv6 NS packets received by Citrix ADC (OBSOLETE).
	*/
	Ipv6icmpnsrxrate float64 `json:"ipv6icmpnsrxrate,omitempty"`
	Ipv6toticmpnatx uint64 `json:"ipv6toticmpnatx,omitempty"`
	/**
	* Number of ICMPv6 NA packets transmitted by Citrix ADC (OBSOLETE).
	*/
	Ipv6icmpnatxrate float64 `json:"ipv6icmpnatxrate,omitempty"`
	Ipv6toticmpnstx uint64 `json:"ipv6toticmpnstx,omitempty"`
	/**
	* Number of ICMPv6 NS packets transmitted by Citrix ADC (OBSOLETE).
	*/
	Ipv6icmpnstxrate float64 `json:"ipv6icmpnstxrate,omitempty"`
	Ipv6toticmprarx uint64 `json:"ipv6toticmprarx,omitempty"`
	/**
	* Number of ICMPv6 RA packets received by Citrix ADC (OBSOLETE).
	*/
	Ipv6icmprarxrate float64 `json:"ipv6icmprarxrate,omitempty"`
	Ipv6toticmprstx uint64 `json:"ipv6toticmprstx,omitempty"`
	/**
	* Number of ICMPv6 RS packets transmitted by Citrix ADC (OBSOLETE).
	*/
	Ipv6icmprstxrate float64 `json:"ipv6icmprstxrate,omitempty"`
	Ipv6toticmprx uint64 `json:"ipv6toticmprx,omitempty"`
	/**
	* Number of ICMPv6 packets received by Citrix ADC (OBSOLETE).
	*/
	Ipv6icmprxrate float64 `json:"ipv6icmprxrate,omitempty"`
	Ipv6toticmptx uint64 `json:"ipv6toticmptx,omitempty"`
	/**
	* Number of ICMPv6 packets transmitted by Citrix ADC (OBSOLETE).
	*/
	Ipv6icmptxrate float64 `json:"ipv6icmptxrate,omitempty"`
	Ipv6errrxhdr uint64 `json:"ipv6errrxhdr,omitempty"`
	Ipv6errrxpkt uint64 `json:"ipv6errrxpkt,omitempty"`
	Ipv6badcksum uint64 `json:"ipv6badcksum,omitempty"`
	Icmpv6error uint64 `json:"icmpv6error,omitempty"`
	Icmpv6unspt uint64 `json:"icmpv6unspt,omitempty"`
	Icmpv6rtthsld uint64 `json:"icmpv6rtthsld,omitempty"`

}
