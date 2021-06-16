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
* Statistics for mptcp protocol resource.
*/

type Protocolmptcpstats struct {
	/**
	* Clear the statsistics / counters
	*/
	Clearstats string `json:"clearstats,omitempty"`
	Mptcptotmpcapsession uint64 `json:"mptcptotmpcapsession,omitempty"`
	/**
	* MPTCP total sessions created
	*/
	Mptcpmpcapsessionrate float64 `json:"mptcpmpcapsessionrate,omitempty"`
	Mptcptotsfconn uint64 `json:"mptcptotsfconn,omitempty"`
	/**
	* MPTCP total Subflow connections created
	*/
	Mptcpsfconnrate float64 `json:"mptcpsfconnrate,omitempty"`
	Mptcpcurmpcapablesessions uint32 `json:"mptcpcurmpcapablesessions,omitempty"`
	Mptcpcursfconnections uint32 `json:"mptcpcursfconnections,omitempty"`
	Mptcpcurpendingjoin uint32 `json:"mptcpcurpendingjoin,omitempty"`
	Mptcpcursesswithoutsfs uint32 `json:"mptcpcursesswithoutsfs,omitempty"`
	Mptcptotmpcapsyn uint64 `json:"mptcptotmpcapsyn,omitempty"`
	/**
	* MPTCP total MP_CAPABLE received
	*/
	Mptcpmpcapsynrate float64 `json:"mptcpmpcapsynrate,omitempty"`
	Mptcptotmpcapsteered uint64 `json:"mptcptotmpcapsteered,omitempty"`
	/**
	* Total MP_CAPABLE sessions steered
	*/
	Mptcpmpcapsteeredrate float64 `json:"mptcpmpcapsteeredrate,omitempty"`
	Mptcptotconnest uint64 `json:"mptcptotconnest,omitempty"`
	/**
	* Total MP_CAPABLE sessions created.
	*/
	Mptcpconnestrate float64 `json:"mptcpconnestrate,omitempty"`
	Mptcptotmpcapsynacksent uint64 `json:"mptcptotmpcapsynacksent,omitempty"`
	/**
	* Total number of MP_CAPABLE SYN/ACKs sent.
	*/
	Mptcpmpcapsynacksentrate float64 `json:"mptcpmpcapsynacksentrate,omitempty"`
	Mptcptotmpcapfackrecvd uint64 `json:"mptcptotmpcapfackrecvd,omitempty"`
	/**
	* Total number of MP_CAPABLE Final ACKs received.
	*/
	Mptcpmpcapfackrecvdrate float64 `json:"mptcpmpcapfackrecvdrate,omitempty"`
	Mptcptotmpjoinsyn uint64 `json:"mptcptotmpjoinsyn,omitempty"`
	/**
	* MPTCP total MP_JOIN syn received
	*/
	Mptcpmpjoinsynrate float64 `json:"mptcpmpjoinsynrate,omitempty"`
	Mptcptotmpjoinsteered uint64 `json:"mptcptotmpjoinsteered,omitempty"`
	/**
	* Total MP_JOIN subflows steered
	*/
	Mptcpmpjoinsteeredrate float64 `json:"mptcpmpjoinsteeredrate,omitempty"`
	Mptcptotmpjoinsynacksent uint64 `json:"mptcptotmpjoinsynacksent,omitempty"`
	/**
	* Total MP_JOIN SYN/ACKs sent.
	*/
	Mptcpmpjoinsynacksentrate float64 `json:"mptcpmpjoinsynacksentrate,omitempty"`
	Mptcptotmpjoinfackrecvd uint64 `json:"mptcptotmpjoinfackrecvd,omitempty"`
	/**
	* Total number of MP_JOIN Final ACKs
	*/
	Mptcpmpjoinfackrecvdrate float64 `json:"mptcpmpjoinfackrecvdrate,omitempty"`
	Mptcptotmpjoin4thacksent uint64 `json:"mptcptotmpjoin4thacksent,omitempty"`
	/**
	* Total number of Subflow final ACK from peer in 3 way handshake validated with 4th ACK.
	*/
	Mptcpmpjoin4thacksentrate float64 `json:"mptcpmpjoin4thacksentrate,omitempty"`
	Mptcptotestsfreplaced uint64 `json:"mptcptotestsfreplaced,omitempty"`
	/**
	* MPTCP Total established subflows replaced due to new MP_JOIN.
	*/
	Mptcpestsfreplacedrate float64 `json:"mptcpestsfreplacedrate,omitempty"`
	Mptcptotpendsfreplaced uint64 `json:"mptcptotpendsfreplaced,omitempty"`
	/**
	* MPTCP Total pending subflows replaced due to new MP_JOIN.
	*/
	Mptcppendsfreplacedrate float64 `json:"mptcppendsfreplacedrate,omitempty"`
	Mptcptotfreshackfrwd uint64 `json:"mptcptotfreshackfrwd,omitempty"`
	/**
	* Fresh ACK recieved on a subflow
	*/
	Mptcpfreshackfrwdrate float64 `json:"mptcpfreshackfrwdrate,omitempty"`
	Mptcpplainackfallback uint64 `json:"mptcpplainackfallback,omitempty"`
	/**
	* MPTCP Fallback to regular tcp on receiving plain ACK for DSS.
	*/
	Mptcpplainackfallbackrate float64 `json:"mptcpplainackfallbackrate,omitempty"`
	Mptcpinfinitemaprecvd uint64 `json:"mptcpinfinitemaprecvd,omitempty"`
	/**
	* MPTCP Received and set infinite map and fallen back to regular TCP.
	*/
	Mptcpinfinitemaprecvdrate float64 `json:"mptcpinfinitemaprecvdrate,omitempty"`
	Mptcptotaddrremoved uint64 `json:"mptcptotaddrremoved,omitempty"`
	/**
	* Total number of addresses removed from MPTCP connection with REMOVE_ADDR option
	*/
	Mptcpaddrremovedrate float64 `json:"mptcpaddrremovedrate,omitempty"`
	Mptcptotdss uint64 `json:"mptcptotdss,omitempty"`
	/**
	* Total number of Data Sequence Signal packets.
	*/
	Mptcpdssrate float64 `json:"mptcpdssrate,omitempty"`
	Mptcptotrxdss uint64 `json:"mptcptotrxdss,omitempty"`
	/**
	* MPTCP Total Data Sequence Signal packets received.
	*/
	Mptcprxdssrate float64 `json:"mptcprxdssrate,omitempty"`
	Mptcptottxdss uint64 `json:"mptcptottxdss,omitempty"`
	/**
	* MMPTCP Total Data Sequence Signal packets sent
	*/
	Mptcptxdssrate float64 `json:"mptcptxdssrate,omitempty"`
	Mptcptotdssa uint64 `json:"mptcptotdssa,omitempty"`
	/**
	* Total Data Sequence Signal packets during data transfer with DATA_ACK
	*/
	Mptcpdssarate float64 `json:"mptcpdssarate,omitempty"`
	Mptcptotdssfreshack uint64 `json:"mptcptotdssfreshack,omitempty"`
	/**
	* MPTCP total Data Sequence Signal packets during  data transfer with fresh ACK
	*/
	Mptcpdssfreshackrate float64 `json:"mptcpdssfreshackrate,omitempty"`
	Mptcptotdssm uint64 `json:"mptcptotdssm,omitempty"`
	/**
	* MPTCP total data Sequence Signal packets with Data Sequence Mapping and checksum
	*/
	Mptcpdssmrate float64 `json:"mptcpdssmrate,omitempty"`
	Mptcptotinfinitemapfrwd uint64 `json:"mptcptotinfinitemapfrwd,omitempty"`
	/**
	* MPTCP received Data Sequence Signal with  infinite map flag (Fallback to regular TCP).
	*/
	Mptcpinfinitemapfrwdrate float64 `json:"mptcpinfinitemapfrwdrate,omitempty"`
	Mptcptotdatalessthandatalen uint64 `json:"mptcptotdatalessthandatalen,omitempty"`
	/**
	* MPTCP Map amount of data not yet received.
	*/
	Mptcpdatalessthandatalenrate float64 `json:"mptcpdatalessthandatalenrate,omitempty"`
	Mptcppriobackuprx uint64 `json:"mptcppriobackuprx,omitempty"`
	/**
	* MPTCP Subflow used as backup path.
	*/
	Mptcppriobackuprxrate float64 `json:"mptcppriobackuprxrate,omitempty"`
	Mptcpprioclearbackuprx uint64 `json:"mptcpprioclearbackuprx,omitempty"`
	/**
	* Subflow earlier used only as a backup subflow, changes to regular subflow with MP_PRIO option
	*/
	Mptcpprioclearbackuprxrate float64 `json:"mptcpprioclearbackuprxrate,omitempty"`
	Mptcptottxdatafin uint64 `json:"mptcptottxdatafin,omitempty"`
	/**
	* Total MPTCP connection close requests sent
	*/
	Mptcptxdatafinrate float64 `json:"mptcptxdatafinrate,omitempty"`
	Mptcptotrxdatafin uint64 `json:"mptcptotrxdatafin,omitempty"`
	/**
	* Total MPTCP connection close(DATA_FIN) requests received.
	*/
	Mptcprxdatafinrate float64 `json:"mptcprxdatafinrate,omitempty"`
	Mptcptottxsffin uint64 `json:"mptcptottxsffin,omitempty"`
	/**
	* MPTCP total subflow close requests.
	*/
	Mptcptxsffinrate float64 `json:"mptcptxsffinrate,omitempty"`
	Mptcperrinvalcookie uint64 `json:"mptcperrinvalcookie,omitempty"`
	/**
	* MPTCP invalid cookie received on MP_CAPABLE final ACK.
	*/
	Mptcperrinvalcookierate float64 `json:"mptcperrinvalcookierate,omitempty"`
	Mptcperrextnflagset uint64 `json:"mptcperrextnflagset,omitempty"`
	/**
	* Extension flag is set on MP_CAPABLE request.
	*/
	Mptcperrextnflagsetrate float64 `json:"mptcperrextnflagsetrate,omitempty"`
	Mptcperrresflagset uint64 `json:"mptcperrresflagset,omitempty"`
	/**
	* MPTCP One or more reserved bits are set on MP_CAPABLE request.
	*/
	Mptcperrresflagsetrate float64 `json:"mptcperrresflagsetrate,omitempty"`
	Mptcperrunknowntoken uint64 `json:"mptcperrunknowntoken,omitempty"`
	/**
	* MPTCP invalid token received on MP_JOIN request.
	*/
	Mptcperrunknowntokenrate float64 `json:"mptcperrunknowntokenrate,omitempty"`
	Mptcperraddridexist uint64 `json:"mptcperraddridexist,omitempty"`
	/**
	* MPTCP MP_JOIN request on existing address id.
	*/
	Mptcperraddridexistrate float64 `json:"mptcperraddridexistrate,omitempty"`
	Mptcperraddrid0 uint64 `json:"mptcperraddrid0,omitempty"`
	/**
	* MPTCP MP_JOIN request on address id 0.
	*/
	Mptcperraddrid0rate float64 `json:"mptcperraddrid0rate,omitempty"`
	Mptcperrmaxsf uint64 `json:"mptcperrmaxsf,omitempty"`
	/**
	* MPTCP new MP_JOIN request after maximum configured subflows are established.
	*/
	Mptcperrmaxsfrate float64 `json:"mptcperrmaxsfrate,omitempty"`
	Mptcperrjointhreshold uint64 `json:"mptcperrjointhreshold,omitempty"`
	/**
	* MPTCP Global pending MP_JOIN threshold limit is reached, new MP_JOIN request will be dropped sending RST
	*/
	Mptcperrjointhresholdrate float64 `json:"mptcperrjointhresholdrate,omitempty"`
	Mptcperrjoinafterfallback uint64 `json:"mptcperrjoinafterfallback,omitempty"`
	/**
	* MPTCP New MP_JOIN request received after fallback to regular tcp.
	*/
	Mptcperrjoinafterfallbackrate float64 `json:"mptcperrjoinafterfallbackrate,omitempty"`
	Mptcperrinvalmac uint64 `json:"mptcperrinvalmac,omitempty"`
	/**
	* MPTCP invalid MAC on MP_JOIN final ACK.
	*/
	Mptcperrinvalmacrate float64 `json:"mptcperrinvalmacrate,omitempty"`
	Mptcperrinvalopts uint64 `json:"mptcperrinvalopts,omitempty"`
	/**
	* MPTCP invalid mptcp option is received and is dropped.
	*/
	Mptcperrinvaloptsrate float64 `json:"mptcperrinvaloptsrate,omitempty"`
	Mptcperroptiondiscarded uint64 `json:"mptcperroptiondiscarded,omitempty"`
	/**
	* Invalid subtype in MPTCP option field and hence discarded.
	*/
	Mptcperroptiondiscardedrate float64 `json:"mptcperroptiondiscardedrate,omitempty"`
	Mptcperroptsnosession uint64 `json:"mptcperroptsnosession,omitempty"`
	/**
	* MPTCP options sent on non existing connection/subflow PCBs.
	*/
	Mptcperroptsnosessionrate float64 `json:"mptcperroptsnosessionrate,omitempty"`
	Mptcperrinvalremaddr uint64 `json:"mptcperrinvalremaddr,omitempty"`
	/**
	* MPTCP remove address request received on invalid/unknown address id.
	*/
	Mptcperrinvalremaddrrate float64 `json:"mptcperrinvalremaddrrate,omitempty"`
	Mptcperroptssendrst uint64 `json:"mptcperroptssendrst,omitempty"`
	/**
	* MPTCP sent RST on receiving improper option field.
	*/
	Mptcperroptssendrstrate float64 `json:"mptcperroptssendrstrate,omitempty"`
	Mptcperrremaddrself uint64 `json:"mptcperrremaddrself,omitempty"`
	/**
	* MPTCP remove address request for self address.
	*/
	Mptcperrremaddrselfrate float64 `json:"mptcperrremaddrselfrate,omitempty"`
	Mptcperrrssffail uint64 `json:"mptcperrrssffail,omitempty"`
	/**
	* Add RSS filter to steer traffic to right node on established MPTCP session failed.
	*/
	Mptcperrrssffailrate float64 `json:"mptcperrrssffailrate,omitempty"`
	Mptcperrnopayloadlenpkt uint64 `json:"mptcperrnopayloadlenpkt,omitempty"`
	/**
	* MPTCP Payload length not specified in packet.
	*/
	Mptcperrnopayloadlenpktrate float64 `json:"mptcperrnopayloadlenpktrate,omitempty"`
	Mptcperrunsupportedmssnegotiated uint64 `json:"mptcperrunsupportedmssnegotiated,omitempty"`
	/**
	* MPTCP Unsupported MSS negotiated error.
	*/
	Mptcperrunsupportedmssnegotiatedrate float64 `json:"mptcperrunsupportedmssnegotiatedrate,omitempty"`
	Mptcperrbadcksum uint64 `json:"mptcperrbadcksum,omitempty"`
	/**
	* MPTCP checksum failed. Connection will fallback to regular tcp.
	*/
	Mptcperrbadcksumrate float64 `json:"mptcperrbadcksumrate,omitempty"`
	Mptcperrcryptonotsupported uint64 `json:"mptcperrcryptonotsupported,omitempty"`
	/**
	* MPTCP client crypto algorithm not supported.
	*/
	Mptcperrcryptonotsupportedrate float64 `json:"mptcperrcryptonotsupportedrate,omitempty"`
	Mptcperrversionnotsupported uint64 `json:"mptcperrversionnotsupported,omitempty"`
	/**
	* MPTCP MP_CAPABLE request from unsupported mptcp client.
	*/
	Mptcperrversionnotsupportedrate float64 `json:"mptcperrversionnotsupportedrate,omitempty"`
	Mptcpplainackrst uint64 `json:"mptcpplainackrst,omitempty"`
	/**
	* MPTCP Sent RST on receiving plain ACK for DSS.
	*/
	Mptcpplainackrstrate float64 `json:"mptcpplainackrstrate,omitempty"`
	Mptcperrdatafinpassive uint64 `json:"mptcperrdatafinpassive,omitempty"`
	/**
	* MPTCP Data FIN received on passive subflow
	*/
	Mptcperrdatafinpassiverate float64 `json:"mptcperrdatafinpassiverate,omitempty"`
	Mptcperrfastclose uint64 `json:"mptcperrfastclose,omitempty"`
	/**
	* MPTCP FAST CLOSE sent.
	*/
	Mptcperrfastcloserate float64 `json:"mptcperrfastcloserate,omitempty"`
	Mptcperrfastclosepassive uint64 `json:"mptcperrfastclosepassive,omitempty"`
	/**
	* MPTCP Fast close received on passive subflow.
	*/
	Mptcperrfastclosepassiverate float64 `json:"mptcperrfastclosepassiverate,omitempty"`
	Mptcperrfastclosekey uint64 `json:"mptcperrfastclosekey,omitempty"`
	/**
	* MPTCP FAST_CLOSE received with invalid key and the packet is dropped.
	*/
	Mptcperrfastclosekeyrate float64 `json:"mptcperrfastclosekeyrate,omitempty"`
	Mptcpmpfailsent uint64 `json:"mptcpmpfailsent,omitempty"`
	/**
	* MPTCP Total MP_FAIL sent due to checksum failure.
	*/
	Mptcpmpfailsentrate float64 `json:"mptcpmpfailsentrate,omitempty"`
	Mptcpmpfailrecvd uint64 `json:"mptcpmpfailrecvd,omitempty"`
	/**
	* MPTCP Total MP_FAIL received and fallback to regular TCP.
	*/
	Mptcpmpfailrecvdrate float64 `json:"mptcpmpfailrecvdrate,omitempty"`
	Mptcperrnomappktrcvd uint64 `json:"mptcperrnomappktrcvd,omitempty"`
	/**
	* MPTCP Packet received with no Data Sequence Mapping.
	*/
	Mptcperrnomappktrcvdrate float64 `json:"mptcperrnomappktrcvdrate,omitempty"`
	Mptcptotmoredatarcvd uint64 `json:"mptcptotmoredatarcvd,omitempty"`
	/**
	* MPTCP More data received than the available Data Sequence Mapping.
	*/
	Mptcpmoredatarcvdrate float64 `json:"mptcpmoredatarcvdrate,omitempty"`
	Mptcperrbadmapconndrop uint64 `json:"mptcperrbadmapconndrop,omitempty"`
	/**
	* MPTCP Drop the session incase of invalid Data Sequence map.
	*/
	Mptcperrbadmapconndroprate float64 `json:"mptcperrbadmapconndroprate,omitempty"`
	Mptcperrdupmaprecvd uint64 `json:"mptcperrdupmaprecvd,omitempty"`
	/**
	* MPTCP Duplicate maps in Data Sequence map table.
	*/
	Mptcperrdupmaprecvdrate float64 `json:"mptcperrdupmaprecvdrate,omitempty"`
	Mptcperrinvalidsfn uint64 `json:"mptcperrinvalidsfn,omitempty"`
	/**
	* MPTCP subflow map doesn't exactly match MPTCP session mapping.
	*/
	Mptcperrinvalidsfnrate float64 `json:"mptcperrinvalidsfnrate,omitempty"`
	Mptcperrmapexists uint64 `json:"mptcperrmapexists,omitempty"`
	/**
	* MPTCP sequence map already exists.
	*/
	Mptcperrmapexistsrate float64 `json:"mptcperrmapexistsrate,omitempty"`
	Mptcperrretxpktrcvd uint64 `json:"mptcperrretxpktrcvd,omitempty"`
	/**
	* Retransmitted Data Recevied on MPTCP session.
	*/
	Mptcperrretxpktrcvdrate float64 `json:"mptcperrretxpktrcvdrate,omitempty"`
	Mptcperrsfsessionallocfail uint64 `json:"mptcperrsfsessionallocfail,omitempty"`
	/**
	* Attaching the subflow to MPTCP session failed due to failure in allocating memory to subflow map table.
	*/
	Mptcperrsfsessionallocfailrate float64 `json:"mptcperrsfsessionallocfailrate,omitempty"`
	Mptcperrmpcapsessionallocfail uint64 `json:"mptcperrmpcapsessionallocfail,omitempty"`
	/**
	* Creating a MPTCP connection failed due to failure in allocating memory to MPTCP connection management structure.
	*/
	Mptcperrmpcapsessionallocfailrate float64 `json:"mptcperrmpcapsessionallocfailrate,omitempty"`
	Mptcptotmpcapsfpcballoc uint64 `json:"mptcptotmpcapsfpcballoc,omitempty"`
	/**
	* Allocating memory to TCP protocol control block(PCB) for subflow failed.
	*/
	Mptcpmpcapsfpcballocrate float64 `json:"mptcpmpcapsfpcballocrate,omitempty"`
	Mptcptotmpcballocfailed uint64 `json:"mptcptotmpcballocfailed,omitempty"`
	/**
	* Allocating memory to MPTCP protocol control block failed.
	*/
	Mptcpmpcballocfailedrate float64 `json:"mptcpmpcballocfailedrate,omitempty"`
	Mptcperrnsballocfailed uint64 `json:"mptcperrnsballocfailed,omitempty"`
	/**
	* Failed to allocate memory to output MPTCP packet.
	*/
	Mptcperrnsballocfailedrate float64 `json:"mptcperrnsballocfailedrate,omitempty"`
	Mptcperrnosffreensb uint64 `json:"mptcperrnosffreensb,omitempty"`
	/**
	* MPTCP output a packet without any subflow PCB.
	*/
	Mptcperrnosffreensbrate float64 `json:"mptcperrnosffreensbrate,omitempty"`

}
