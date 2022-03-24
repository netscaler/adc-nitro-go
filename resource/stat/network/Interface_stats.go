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
	Clearstats      string `json:"clearstats,omitempty"`
	Curintfstate    string `json:"curintfstate,omitempty"`
	Curlinkuptime   string `json:"curlinkuptime,omitempty"`
	Curlinkdowntime string `json:"curlinkdowntime,omitempty"`
	Totrxbytes      int    `json:"totrxbytes,omitempty"`
	/**
	* Number of bytes received by an interface since the Citrix ADC was started or the interface statistics were cleared.
	 */
	Rxbytesrate float64 `json:"rxbytesrate,omitempty"`
	Tottxbytes  int     `json:"tottxbytes,omitempty"`
	/**
	* Number of bytes transmitted by an interface since the Citrix ADC was started or the interface statistics were cleared.
	 */
	Txbytesrate float64 `json:"txbytesrate,omitempty"`
	Totrxpkts   int     `json:"totrxpkts,omitempty"`
	/**
	* Number of packets received by an interface since the Citrix ADC was started or the interface statistics were cleared.
	 */
	Rxpktsrate float64 `json:"rxpktsrate,omitempty"`
	Tottxpkts  int     `json:"tottxpkts,omitempty"`
	/**
	* Number of packets transmitted by an interface since the Citrix ADC was started or the interface statistics were cleared.
	 */
	Txpktsrate        float64 `json:"txpktsrate,omitempty"`
	Jumbopktsreceived int     `json:"jumbopktsreceived,omitempty"`
	/**
	* Number of Jumbo Packets received on this interface.
	 */
	Jumbopktsreceivedrate float64 `json:"jumbopktsreceivedrate,omitempty"`
	Jumbopktstransmitted  int     `json:"jumbopktstransmitted,omitempty"`
	/**
	* Number of Jumbo packets transmitted on this interface by upper layer, with TSO enabled actual trasmission size could be non Jumbo.
	 */
	Jumbopktstransmittedrate float64 `json:"jumbopktstransmittedrate,omitempty"`
	Trunkpktsreceived        int     `json:"trunkpktsreceived,omitempty"`
	/**
	* Number of Tagged Packets received on this Trunk interface through Allowed VLan List.
	 */
	Trunkpktsreceivedrate float64 `json:"trunkpktsreceivedrate,omitempty"`
	Trunkpktstransmitted  int     `json:"trunkpktstransmitted,omitempty"`
	/**
	* Number of Tagged Packets transmitted on this Trunk interface through Allowed VLan List.
	 */
	Trunkpktstransmittedrate float64 `json:"trunkpktstransmittedrate,omitempty"`
	Nictotmulticastpkts      int     `json:"nictotmulticastpkts,omitempty"`
	/**
	* Number of multicast packets received by the specified interface since the Citrix ADC was started or the interface statistics were cleared.
	 */
	Nicmulticastpktsrate float64 `json:"nicmulticastpktsrate,omitempty"`
	Totnetscalerpkts     int     `json:"totnetscalerpkts,omitempty"`
	/**
	* Number of packets, destined to the Citrix ADC, received by an interface since the Citrix ADC was started or the interface statistics were cleared. The packets destined to Citrix ADC are those that have the same MAC address as that of an interface or a VMAC address owned by the Citrix ADC.
	 */
	Netscalerpktsrate float64 `json:"netscalerpktsrate,omitempty"`
	Rxlacpdu          int     `json:"rxlacpdu,omitempty"`
	/**
	* Number of Link Aggregation Control Protocol Data Units(LACPDUs) received by the specified interface since the Citrix ADC was started or the interface statistics were cleared.
	 */
	Rxlacpdurate float64 `json:"rxlacpdurate,omitempty"`
	Txlacpdu     int     `json:"txlacpdu,omitempty"`
	/**
	* Number of Link Aggregation Control Protocol Data Units(LACPDUs) transmitted by the specified interface since the Citrix ADC was started or the interface statistics were cleared.
	 */
	Txlacpdurate float64 `json:"txlacpdurate,omitempty"`
	Errpktrx     int     `json:"errpktrx,omitempty"`
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
	Errpktrxrate float64 `json:"errpktrxrate,omitempty"`
	Errpkttx     int     `json:"errpkttx,omitempty"`
	/**
	* Number of outbound packets dropped by the hardware on a specified interface since the Citrix ADC was started or the interface statistics were cleared. This could happen due to length (undersize or oversize) errors and lack of resources. This statistic is available only for:
		(1) Loop back interface (LO) of all platforms.
		(2) All data ports on the Citrix ADC 12000 platform.
		(3) Management ports on the MPX 15000 and 17000 platforms.
	*/
	Errpkttxrate    float64 `json:"errpkttxrate,omitempty"`
	Errifindiscards int     `json:"errifindiscards,omitempty"`
	/**
	* Number of error-free inbound packets discarded by the specified interface due to a lack of resources, for example, insufficient receive buffers.
	 */
	Errifindiscardsrate float64 `json:"errifindiscardsrate,omitempty"`
	Nicerrifoutdiscards int     `json:"nicerrifoutdiscards,omitempty"`
	/**
	* Number of error-free outbound packets discarded by the specified interface due to a lack of resources. This statistic is not available on:
		(1) 10G ports of Citrix ADC MPX 12500/12500/15500-10G  platforms.
		(2) 10G data ports on Citrix ADC MPX 17500/19500/21500 platforms.
	*/
	Nicerrifoutdiscardsrate float64 `json:"nicerrifoutdiscardsrate,omitempty"`
	Errdroppedrxpkts        int     `json:"errdroppedrxpkts,omitempty"`
	/**
	* Number of inbound packets dropped by the specified interface. Commonly dropped packets are multicast frames, spanning tree BPDUs, packets destined to a MAC not owned by the Citrix ADC when L2 mode is disabled, or packets tagged for a VLAN that is not bound to the interface.  This statistic will increment in most healthy networks at a steady rate regardless of traffic load.  If a sharp spike in dropped packets occurs, it generally indicates an issue with connected L2 switches, such as a forwarding database overflow resulting in packets being broadcast on all ports.
	 */
	Errdroppedrxpktsrate float64 `json:"errdroppedrxpktsrate,omitempty"`
	Errdroppedtxpkts     int     `json:"errdroppedtxpkts,omitempty"`
	/**
	* Number of packets dropped in transmission by the specified interface due to one of the following reasons.
		(1) VLAN mismatch.
		(2) Oversized packets.
		(3) Interface congestion.
		(4) Loopback packets sent on non loop back interface.
	*/
	Errdroppedtxpktsrate float64 `json:"errdroppedtxpktsrate,omitempty"`
	Errlinkhangs         int     `json:"errlinkhangs,omitempty"`
	Nicstsstalls         int     `json:"nicstsstalls,omitempty"`
	Nictxstalls          int     `json:"nictxstalls,omitempty"`
	Nicrxstalls          int     `json:"nicrxstalls,omitempty"`
	Nicerrdisables       int     `json:"nicerrdisables,omitempty"`
	Linkreinits          int     `json:"linkreinits,omitempty"`
	Totmacmoved          int     `json:"totmacmoved,omitempty"`
	/**
	* Number of MAC moves between ports. If a high rate of MAC moves is observed, it is likely that there is a bridge loop between two interfaces.
	 */
	Macmovedrate float64 `json:"macmovedrate,omitempty"`
	Errnicmuted  int     `json:"errnicmuted,omitempty"`
	Rxcrcerrors  int     `json:"rxcrcerrors,omitempty"`
	/**
	* Number of packets received with the wrong checksum by the specified interface since the Citrix ADC was started or the interface statistics were cleared. This indicates the number of Jabber frames received instead of CRC errors on the 10G data ports of Citrix ADC 12000-10G platform and the data ports of Citrix ADC MPX 15000 and 17000 platforms.
	 */
	Rxcrcerrorsrate float64 `json:"rxcrcerrorsrate,omitempty"`
	Interfacealias  string  `json:"interfacealias,omitempty"`
	Curlinkstate    string  `json:"curlinkstate,omitempty"`
}
