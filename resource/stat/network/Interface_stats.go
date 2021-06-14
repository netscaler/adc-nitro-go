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
* Statistics for interface resource.
*/

type Interfacestats struct {
	/**
	* Interface number, in C/U format, where C can take one of the following values:
		* 0 - Indicates a management interface.
		* 1 - Indicates a 1 Gbps port.
		* 10 - Indicates a 10 Gbps port.
		* LA - Indicates a link aggregation port.
		* LO - Indicates a loop back port.
		U is a unique integer for representing an interface in a particular port group.
	*/
	Id string `json:"id,omitempty"`
	/**
	* Clear the statsistics / counters
	*/
	Clearstats string `json:"clearstats,omitempty"`
	Curintfstate string `json:"curintfstate,omitempty"`
	Curlinkuptime string `json:"curlinkuptime,omitempty"`
	Curlinkdowntime string `json:"curlinkdowntime,omitempty"`
	Totrxbytes uint64 `json:"totrxbytes,omitempty"`
	/**
	* Number of bytes received by an interface since the Citrix ADC was started or the interface statistics were cleared.
	*/
	Rxbytesrate int32 `json:"rxbytesrate,omitempty"`
	Tottxbytes uint64 `json:"tottxbytes,omitempty"`
	/**
	* Number of bytes transmitted by an interface since the Citrix ADC was started or the interface statistics were cleared.
	*/
	Txbytesrate int32 `json:"txbytesrate,omitempty"`
	Totrxpkts uint64 `json:"totrxpkts,omitempty"`
	/**
	* Number of packets received by an interface since the Citrix ADC was started or the interface statistics were cleared.
	*/
	Rxpktsrate int32 `json:"rxpktsrate,omitempty"`
	Tottxpkts uint64 `json:"tottxpkts,omitempty"`
	/**
	* Number of packets transmitted by an interface since the Citrix ADC was started or the interface statistics were cleared.
	*/
	Txpktsrate int32 `json:"txpktsrate,omitempty"`
	Jumbopktsreceived uint64 `json:"jumbopktsreceived,omitempty"`
	/**
	* Number of Jumbo Packets received on this interface.
	*/
	Jumbopktsreceivedrate int32 `json:"jumbopktsreceivedrate,omitempty"`
	Jumbopktstransmitted uint64 `json:"jumbopktstransmitted,omitempty"`
	/**
	* Number of Jumbo packets transmitted on this interface by upper layer, with TSO enabled actual trasmission size could be non Jumbo.
	*/
	Jumbopktstransmittedrate int32 `json:"jumbopktstransmittedrate,omitempty"`
	Trunkpktsreceived uint64 `json:"trunkpktsreceived,omitempty"`
	/**
	* Number of Tagged Packets received on this Trunk interface through Allowed VLan List.
	*/
	Trunkpktsreceivedrate int32 `json:"trunkpktsreceivedrate,omitempty"`
	Trunkpktstransmitted uint64 `json:"trunkpktstransmitted,omitempty"`
	/**
	* Number of Tagged Packets transmitted on this Trunk interface through Allowed VLan List.
	*/
	Trunkpktstransmittedrate int32 `json:"trunkpktstransmittedrate,omitempty"`
	Nictotmulticastpkts uint64 `json:"nictotmulticastpkts,omitempty"`
	/**
	* Number of multicast packets received by the specified interface since the Citrix ADC was started or the interface statistics were cleared.
	*/
	Nicmulticastpktsrate int32 `json:"nicmulticastpktsrate,omitempty"`
	Totnetscalerpkts uint64 `json:"totnetscalerpkts,omitempty"`
	/**
	* Number of packets, destined to the Citrix ADC, received by an interface since the Citrix ADC was started or the interface statistics were cleared. The packets destined to Citrix ADC are those that have the same MAC address as that of an interface or a VMAC address owned by the Citrix ADC.
	*/
	Netscalerpktsrate int32 `json:"netscalerpktsrate,omitempty"`
	Rxlacpdu uint64 `json:"rxlacpdu,omitempty"`
	/**
	* Number of Link Aggregation Control Protocol Data Units(LACPDUs) received by the specified interface since the Citrix ADC was started or the interface statistics were cleared.
	*/
	Rxlacpdurate int32 `json:"rxlacpdurate,omitempty"`
	Txlacpdu uint64 `json:"txlacpdu,omitempty"`
	/**
	* Number of Link Aggregation Control Protocol Data Units(LACPDUs) transmitted by the specified interface since the Citrix ADC was started or the interface statistics were cleared.
	*/
	Txlacpdurate int32 `json:"txlacpdurate,omitempty"`
	Errpktrx uint64 `json:"errpktrx,omitempty"`
	/**
	* Number of inbound packets dropped by the hardware on a specified interface once the Citrix ADC starts or the interface statistics are cleared. This happens due to following reasons:
		1)	The hardware receives packets at a rate higher rate than that at which the software is processing packets. In this case, the hardware FIFO overruns and starts dropping the packets .
		2)	The specified interface fails to receive inbound packets from the appliance because of insufficient memory.
		3)	The specified interface receives packets with CRC errors (Alignment or Frame Check Sequence).
		4)	The specified interface receives overly long packets.
		5)	The specified interface receives packets with alignment errors.
		6)	The software does less buffering because it is running out of available memory. When hardware detects that there is no space into which to push newly arrived packets, it starts dropping them.
		7)	The specified interface receives packets with Frame Check Sequence (FCS) errors.
		8)	The specified interface receives packets smaller than 64 bytes.
		9)	The specified interface discards error-free inbound packets because of insufficient resources. For example: NIC buffers.
		10)	Packets are missed because of collision detection, link lost, physical decoding error, or MAC abort.
	*/
	Errpktrxrate int32 `json:"errpktrxrate,omitempty"`
	Errpkttx uint64 `json:"errpkttx,omitempty"`
	/**
	* Number of outbound packets dropped by the hardware on a specified interface since the Citrix ADC was started or the interface statistics were cleared. This could happen due to length (undersize or oversize) errors and lack of resources. This statistic is available only for: 
		(1) Loop back interface (LO) of all platforms.
		(2) All data ports on the Citrix ADC 12000 platform.
		(3) Management ports on the MPX 15000 and 17000 platforms.
	*/
	Errpkttxrate int32 `json:"errpkttxrate,omitempty"`
	Errifindiscards uint64 `json:"errifindiscards,omitempty"`
	/**
	* Number of error-free inbound packets discarded by the specified interface due to a lack of resources, for example, insufficient receive buffers.
	*/
	Errifindiscardsrate int32 `json:"errifindiscardsrate,omitempty"`
	Nicerrifoutdiscards uint64 `json:"nicerrifoutdiscards,omitempty"`
	/**
	* Number of error-free outbound packets discarded by the specified interface due to a lack of resources. This statistic is not available on:
		(1) 10G ports of Citrix ADC MPX 12500/12500/15500-10G  platforms. 
		(2) 10G data ports on Citrix ADC MPX 17500/19500/21500 platforms.
	*/
	Nicerrifoutdiscardsrate int32 `json:"nicerrifoutdiscardsrate,omitempty"`
	Errdroppedrxpkts uint64 `json:"errdroppedrxpkts,omitempty"`
	/**
	* Number of inbound packets dropped by the specified interface. Commonly dropped packets are multicast frames, spanning tree BPDUs, packets destined to a MAC not owned by the Citrix ADC when L2 mode is disabled, or packets tagged for a VLAN that is not bound to the interface.  This statistic will increment in most healthy networks at a steady rate regardless of traffic load.  If a sharp spike in dropped packets occurs, it generally indicates an issue with connected L2 switches, such as a forwarding database overflow resulting in packets being broadcast on all ports.
	*/
	Errdroppedrxpktsrate int32 `json:"errdroppedrxpktsrate,omitempty"`
	Errdroppedtxpkts uint64 `json:"errdroppedtxpkts,omitempty"`
	/**
	* Number of packets dropped in transmission by the specified interface due to one of the following reasons.
		(1) VLAN mismatch.
		(2) Oversized packets.
		(3) Interface congestion.  
		(4) Loopback packets sent on non loop back interface.
	*/
	Errdroppedtxpktsrate int32 `json:"errdroppedtxpktsrate,omitempty"`
	Errlinkhangs uint32 `json:"errlinkhangs,omitempty"`
	Nicstsstalls uint32 `json:"nicstsstalls,omitempty"`
	Nictxstalls uint32 `json:"nictxstalls,omitempty"`
	Nicrxstalls uint32 `json:"nicrxstalls,omitempty"`
	Nicerrdisables uint32 `json:"nicerrdisables,omitempty"`
	Linkreinits uint32 `json:"linkreinits,omitempty"`
	Totmacmoved uint64 `json:"totmacmoved,omitempty"`
	/**
	* Number of MAC moves between ports. If a high rate of MAC moves is observed, it is likely that there is a bridge loop between two interfaces.
	*/
	Macmovedrate int32 `json:"macmovedrate,omitempty"`
	Errnicmuted uint32 `json:"errnicmuted,omitempty"`
	Rxcrcerrors uint64 `json:"rxcrcerrors,omitempty"`
	/**
	* Number of packets received with the wrong checksum by the specified interface since the Citrix ADC was started or the interface statistics were cleared. This indicates the number of Jabber frames received instead of CRC errors on the 10G data ports of Citrix ADC 12000-10G platform and the data ports of Citrix ADC MPX 15000 and 17000 platforms.
	*/
	Rxcrcerrorsrate int32 `json:"rxcrcerrorsrate,omitempty"`
	Interfacealias string `json:"interfacealias,omitempty"`
	Curlinkstate string `json:"curlinkstate,omitempty"`

}
