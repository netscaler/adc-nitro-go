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
* Statistics for icmpv6 resource.
*/

type Protocolicmpv6stats struct {
	/**
	* Clear the statsistics / counters
	*/
	Clearstats string `json:"clearstats,omitempty"`
	Icmpv6totrxpkts uint64 `json:"icmpv6totrxpkts,omitempty"`
	/**
	* ICMPv6 packets received.
	*/
	Icmpv6rxpktsrate int32 `json:"icmpv6rxpktsrate,omitempty"`
	Icmpv6totrxbytes uint64 `json:"icmpv6totrxbytes,omitempty"`
	/**
	* Bytes of ICMPv6 data received.
	*/
	Icmpv6rxbytesrate int32 `json:"icmpv6rxbytesrate,omitempty"`
	Icmpv6tottxpkts uint64 `json:"icmpv6tottxpkts,omitempty"`
	/**
	* ICMPv6 packets transmitted.
	*/
	Icmpv6txpktsrate int32 `json:"icmpv6txpktsrate,omitempty"`
	Icmpv6tottxbytes uint64 `json:"icmpv6tottxbytes,omitempty"`
	/**
	* Bytes of ICMPv6 data transmitted.
	*/
	Icmpv6txbytesrate int32 `json:"icmpv6txbytesrate,omitempty"`
	Icmpv6totrxna uint64 `json:"icmpv6totrxna,omitempty"`
	/**
	* ICMPv6 neighbor advertisement packets received. These packets are received in response to a neighbor solicitation message sent out by this node, or if the link layer address of a neighbor has changed.
	*/
	Icmpv6rxnarate int32 `json:"icmpv6rxnarate,omitempty"`
	Icmpv6totrxns uint64 `json:"icmpv6totrxns,omitempty"`
	/**
	* ICMPv6 neighbor solicitation packets received. These packets are received if the link layer address of a neighbor has changed, or in response to a neighbor solicitation message sent out by this node.
	*/
	Icmpv6rxnsrate int32 `json:"icmpv6rxnsrate,omitempty"`
	Icmpv6totrxra uint64 `json:"icmpv6totrxra,omitempty"`
	/**
	* ICMPv6 router advertisement packets received. These are received at defined intervals or in response to a router solicitation message.
	*/
	Icmpv6rxrarate int32 `json:"icmpv6rxrarate,omitempty"`
	Icmpv6totrxrs uint64 `json:"icmpv6totrxrs,omitempty"`
	/**
	* ICMPv6 router solicitation packets received. These could be sent by a neighboring router to initiate address resolution.
	*/
	Icmpv6rxrsrate int32 `json:"icmpv6rxrsrate,omitempty"`
	Icmpv6totrxechoreq uint64 `json:"icmpv6totrxechoreq,omitempty"`
	/**
	* ICMPv6 Ping Echo Request packets received.
	*/
	Icmpv6rxechoreqrate int32 `json:"icmpv6rxechoreqrate,omitempty"`
	Icmpv6totrxechoreply uint64 `json:"icmpv6totrxechoreply,omitempty"`
	/**
	* ICMPv6 Ping Echo Reply packets received.
	*/
	Icmpv6rxechoreplyrate int32 `json:"icmpv6rxechoreplyrate,omitempty"`
	Icmpv6tottxna uint64 `json:"icmpv6tottxna,omitempty"`
	/**
	* ICMPv6 neighbor advertisement packets transmitted. These packets are sent in response to a neighbor solicitation packet, or if the link layer address of this node has changed.
	*/
	Icmpv6txnarate int32 `json:"icmpv6txnarate,omitempty"`
	Icmpv6tottxns uint64 `json:"icmpv6tottxns,omitempty"`
	/**
	* ICMPv6 neighbor solicitation packets transmitted. These packets are sent to get the link layer addresses of neighboring nodes or to confirm that they are reachable.
	*/
	Icmpv6txnsrate int32 `json:"icmpv6txnsrate,omitempty"`
	Icmpv6tottxra uint64 `json:"icmpv6tottxra,omitempty"`
	/**
	* ICMPv6 router advertisement packets transmitted. These packets are sent at regular intervals or in response to a router solicitation packet from a neighbor.
	*/
	Icmpv6txrarate int32 `json:"icmpv6txrarate,omitempty"`
	Icmpv6tottxrs uint64 `json:"icmpv6tottxrs,omitempty"`
	/**
	* ICMPv6 router solicitation packets transmitted. These packets are sent to request neighboring routers to generate router advertisements immediately rather than wait for the next defined time.
	*/
	Icmpv6txrsrate int32 `json:"icmpv6txrsrate,omitempty"`
	Icmpv6tottxechoreq uint64 `json:"icmpv6tottxechoreq,omitempty"`
	/**
	* ICMPv6 Ping Echo Request packets transmitted.
	*/
	Icmpv6txechoreqrate int32 `json:"icmpv6txechoreqrate,omitempty"`
	Icmpv6tottxechoreply uint64 `json:"icmpv6tottxechoreply,omitempty"`
	/**
	* ICMP Ping Echo Reply packets transmitted.
	*/
	Icmpv6txechoreplyrate int32 `json:"icmpv6txechoreplyrate,omitempty"`
	Icmpv6errra uint64 `json:"icmpv6errra,omitempty"`
	Icmpv6errna uint64 `json:"icmpv6errna,omitempty"`
	Icmpv6errns uint64 `json:"icmpv6errns,omitempty"`
	Icmpv6badchecksums uint64 `json:"icmpv6badchecksums,omitempty"`
	Icmpv6unspt uint64 `json:"icmpv6unspt,omitempty"`
	Icmpv6rtthsld uint64 `json:"icmpv6rtthsld,omitempty"`

}