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
	Clearstats           string `json:"clearstats,omitempty"`
	Mptcptotmpcapsession int    `json:"mptcptotmpcapsession,omitempty"`
	/**
	* MPTCP total sessions created
	 */
	Mptcpmpcapsessionrate float64 `json:"mptcpmpcapsessionrate,omitempty"`
	Mptcptotsfconn        int     `json:"mptcptotsfconn,omitempty"`
	/**
	* MPTCP total Subflow connections created
	 */
	Mptcpsfconnrate           float64 `json:"mptcpsfconnrate,omitempty"`
	Mptcpcurmpcapablesessions int     `json:"mptcpcurmpcapablesessions,omitempty"`
	Mptcpcursfconnections     int     `json:"mptcpcursfconnections,omitempty"`
	Mptcpcurpendingjoin       int     `json:"mptcpcurpendingjoin,omitempty"`
	Mptcpcursesswithoutsfs    int     `json:"mptcpcursesswithoutsfs,omitempty"`
	Mptcptotmpcapsyn          int     `json:"mptcptotmpcapsyn,omitempty"`
	/**
	* MPTCP total MP_CAPABLE received
	 */
	Mptcpmpcapsynrate    float64 `json:"mptcpmpcapsynrate,omitempty"`
	Mptcptotmpcapsteered int     `json:"mptcptotmpcapsteered,omitempty"`
	/**
	* Total MP_CAPABLE sessions steered
	 */
	Mptcpmpcapsteeredrate float64 `json:"mptcpmpcapsteeredrate,omitempty"`
	Mptcptotconnest       int     `json:"mptcptotconnest,omitempty"`
	/**
	* Total MP_CAPABLE sessions created.
	 */
	Mptcpconnestrate        float64 `json:"mptcpconnestrate,omitempty"`
	Mptcptotmpcapsynacksent int     `json:"mptcptotmpcapsynacksent,omitempty"`
	/**
	* Total number of MP_CAPABLE SYN/ACKs sent.
	 */
	Mptcpmpcapsynacksentrate float64 `json:"mptcpmpcapsynacksentrate,omitempty"`
	Mptcptotmpcapfackrecvd   int     `json:"mptcptotmpcapfackrecvd,omitempty"`
	/**
	* Total number of MP_CAPABLE Final ACKs received.
	 */
	Mptcpmpcapfackrecvdrate float64 `json:"mptcpmpcapfackrecvdrate,omitempty"`
	Mptcptotmpjoinsyn       int     `json:"mptcptotmpjoinsyn,omitempty"`
	/**
	* MPTCP total MP_JOIN syn received
	 */
	Mptcpmpjoinsynrate    float64 `json:"mptcpmpjoinsynrate,omitempty"`
	Mptcptotmpjoinsteered int     `json:"mptcptotmpjoinsteered,omitempty"`
	/**
	* Total MP_JOIN subflows steered
	 */
	Mptcpmpjoinsteeredrate   float64 `json:"mptcpmpjoinsteeredrate,omitempty"`
	Mptcptotmpjoinsynacksent int     `json:"mptcptotmpjoinsynacksent,omitempty"`
	/**
	* Total MP_JOIN SYN/ACKs sent.
	 */
	Mptcpmpjoinsynacksentrate float64 `json:"mptcpmpjoinsynacksentrate,omitempty"`
	Mptcptotmpjoinfackrecvd   int     `json:"mptcptotmpjoinfackrecvd,omitempty"`
	/**
	* Total number of MP_JOIN Final ACKs
	 */
	Mptcpmpjoinfackrecvdrate float64 `json:"mptcpmpjoinfackrecvdrate,omitempty"`
	Mptcptotmpjoin4thacksent int     `json:"mptcptotmpjoin4thacksent,omitempty"`
	/**
	* Total number of Subflow final ACK from peer in 3 way handshake validated with 4th ACK.
	 */
	Mptcpmpjoin4thacksentrate float64 `json:"mptcpmpjoin4thacksentrate,omitempty"`
	Mptcptotestsfreplaced     int     `json:"mptcptotestsfreplaced,omitempty"`
	/**
	* MPTCP Total established subflows replaced due to new MP_JOIN.
	 */
	Mptcpestsfreplacedrate float64 `json:"mptcpestsfreplacedrate,omitempty"`
	Mptcptotpendsfreplaced int     `json:"mptcptotpendsfreplaced,omitempty"`
	/**
	* MPTCP Total pending subflows replaced due to new MP_JOIN.
	 */
	Mptcppendsfreplacedrate float64 `json:"mptcppendsfreplacedrate,omitempty"`
	Mptcptotfreshackfrwd    int     `json:"mptcptotfreshackfrwd,omitempty"`
	/**
	* Fresh ACK recieved on a subflow
	 */
	Mptcpfreshackfrwdrate float64 `json:"mptcpfreshackfrwdrate,omitempty"`
	Mptcpplainackfallback int     `json:"mptcpplainackfallback,omitempty"`
	/**
	* MPTCP Fallback to regular tcp on receiving plain ACK for DSS.
	 */
	Mptcpplainackfallbackrate float64 `json:"mptcpplainackfallbackrate,omitempty"`
	Mptcpinfinitemaprecvd     int     `json:"mptcpinfinitemaprecvd,omitempty"`
	/**
	* MPTCP Received and set infinite map and fallen back to regular TCP.
	 */
	Mptcpinfinitemaprecvdrate float64 `json:"mptcpinfinitemaprecvdrate,omitempty"`
	Mptcptotaddrremoved       int     `json:"mptcptotaddrremoved,omitempty"`
	/**
	* Total number of addresses removed from MPTCP connection with REMOVE_ADDR option
	 */
	Mptcpaddrremovedrate float64 `json:"mptcpaddrremovedrate,omitempty"`
	Mptcptotdss          int     `json:"mptcptotdss,omitempty"`
	/**
	* Total number of Data Sequence Signal packets.
	 */
	Mptcpdssrate  float64 `json:"mptcpdssrate,omitempty"`
	Mptcptotrxdss int     `json:"mptcptotrxdss,omitempty"`
	/**
	* MPTCP Total Data Sequence Signal packets received.
	 */
	Mptcprxdssrate float64 `json:"mptcprxdssrate,omitempty"`
	Mptcptottxdss  int     `json:"mptcptottxdss,omitempty"`
	/**
	* MMPTCP Total Data Sequence Signal packets sent
	 */
	Mptcptxdssrate float64 `json:"mptcptxdssrate,omitempty"`
	Mptcptotdssa   int     `json:"mptcptotdssa,omitempty"`
	/**
	* Total Data Sequence Signal packets during data transfer with DATA_ACK
	 */
	Mptcpdssarate       float64 `json:"mptcpdssarate,omitempty"`
	Mptcptotdssfreshack int     `json:"mptcptotdssfreshack,omitempty"`
	/**
	* MPTCP total Data Sequence Signal packets during  data transfer with fresh ACK
	 */
	Mptcpdssfreshackrate float64 `json:"mptcpdssfreshackrate,omitempty"`
	Mptcptotdssm         int     `json:"mptcptotdssm,omitempty"`
	/**
	* MPTCP total data Sequence Signal packets with Data Sequence Mapping and checksum
	 */
	Mptcpdssmrate           float64 `json:"mptcpdssmrate,omitempty"`
	Mptcptotinfinitemapfrwd int     `json:"mptcptotinfinitemapfrwd,omitempty"`
	/**
	* MPTCP received Data Sequence Signal with  infinite map flag (Fallback to regular TCP).
	 */
	Mptcpinfinitemapfrwdrate    float64 `json:"mptcpinfinitemapfrwdrate,omitempty"`
	Mptcptotdatalessthandatalen int     `json:"mptcptotdatalessthandatalen,omitempty"`
	/**
	* MPTCP Map amount of data not yet received.
	 */
	Mptcpdatalessthandatalenrate float64 `json:"mptcpdatalessthandatalenrate,omitempty"`
	Mptcppriobackuprx            int     `json:"mptcppriobackuprx,omitempty"`
	/**
	* MPTCP Subflow used as backup path.
	 */
	Mptcppriobackuprxrate  float64 `json:"mptcppriobackuprxrate,omitempty"`
	Mptcpprioclearbackuprx int     `json:"mptcpprioclearbackuprx,omitempty"`
	/**
	* Subflow earlier used only as a backup subflow, changes to regular subflow with MP_PRIO option
	 */
	Mptcpprioclearbackuprxrate float64 `json:"mptcpprioclearbackuprxrate,omitempty"`
	Mptcptottxdatafin          int     `json:"mptcptottxdatafin,omitempty"`
	/**
	* Total MPTCP connection close requests sent
	 */
	Mptcptxdatafinrate float64 `json:"mptcptxdatafinrate,omitempty"`
	Mptcptotrxdatafin  int     `json:"mptcptotrxdatafin,omitempty"`
	/**
	* Total MPTCP connection close(DATA_FIN) requests received.
	 */
	Mptcprxdatafinrate float64 `json:"mptcprxdatafinrate,omitempty"`
	Mptcptottxsffin    int     `json:"mptcptottxsffin,omitempty"`
	/**
	* MPTCP total subflow close requests.
	 */
	Mptcptxsffinrate    float64 `json:"mptcptxsffinrate,omitempty"`
	Mptcperrinvalcookie int     `json:"mptcperrinvalcookie,omitempty"`
	/**
	* MPTCP invalid cookie received on MP_CAPABLE final ACK.
	 */
	Mptcperrinvalcookierate float64 `json:"mptcperrinvalcookierate,omitempty"`
	Mptcperrextnflagset     int     `json:"mptcperrextnflagset,omitempty"`
	/**
	* Extension flag is set on MP_CAPABLE request.
	 */
	Mptcperrextnflagsetrate float64 `json:"mptcperrextnflagsetrate,omitempty"`
	Mptcperrresflagset      int     `json:"mptcperrresflagset,omitempty"`
	/**
	* MPTCP One or more reserved bits are set on MP_CAPABLE request.
	 */
	Mptcperrresflagsetrate float64 `json:"mptcperrresflagsetrate,omitempty"`
	Mptcperrunknowntoken   int     `json:"mptcperrunknowntoken,omitempty"`
	/**
	* MPTCP invalid token received on MP_JOIN request.
	 */
	Mptcperrunknowntokenrate float64 `json:"mptcperrunknowntokenrate,omitempty"`
	Mptcperraddridexist      int     `json:"mptcperraddridexist,omitempty"`
	/**
	* MPTCP MP_JOIN request on existing address id.
	 */
	Mptcperraddridexistrate float64 `json:"mptcperraddridexistrate,omitempty"`
	Mptcperraddrid0         int     `json:"mptcperraddrid0,omitempty"`
	/**
	* MPTCP MP_JOIN request on address id 0.
	 */
	Mptcperraddrid0rate float64 `json:"mptcperraddrid0rate,omitempty"`
	Mptcperrmaxsf       int     `json:"mptcperrmaxsf,omitempty"`
	/**
	* MPTCP new MP_JOIN request after maximum configured subflows are established.
	 */
	Mptcperrmaxsfrate     float64 `json:"mptcperrmaxsfrate,omitempty"`
	Mptcperrjointhreshold int     `json:"mptcperrjointhreshold,omitempty"`
	/**
	* MPTCP Global pending MP_JOIN threshold limit is reached, new MP_JOIN request will be dropped sending RST
	 */
	Mptcperrjointhresholdrate float64 `json:"mptcperrjointhresholdrate,omitempty"`
	Mptcperrjoinafterfallback int     `json:"mptcperrjoinafterfallback,omitempty"`
	/**
	* MPTCP New MP_JOIN request received after fallback to regular tcp.
	 */
	Mptcperrjoinafterfallbackrate float64 `json:"mptcperrjoinafterfallbackrate,omitempty"`
	Mptcperrinvalmac              int     `json:"mptcperrinvalmac,omitempty"`
	/**
	* MPTCP invalid MAC on MP_JOIN final ACK.
	 */
	Mptcperrinvalmacrate float64 `json:"mptcperrinvalmacrate,omitempty"`
	Mptcperrinvalopts    int     `json:"mptcperrinvalopts,omitempty"`
	/**
	* MPTCP invalid mptcp option is received and is dropped.
	 */
	Mptcperrinvaloptsrate   float64 `json:"mptcperrinvaloptsrate,omitempty"`
	Mptcperroptiondiscarded int     `json:"mptcperroptiondiscarded,omitempty"`
	/**
	* Invalid subtype in MPTCP option field and hence discarded.
	 */
	Mptcperroptiondiscardedrate float64 `json:"mptcperroptiondiscardedrate,omitempty"`
	Mptcperroptsnosession       int     `json:"mptcperroptsnosession,omitempty"`
	/**
	* MPTCP options sent on non existing connection/subflow PCBs.
	 */
	Mptcperroptsnosessionrate float64 `json:"mptcperroptsnosessionrate,omitempty"`
	Mptcperrinvalremaddr      int     `json:"mptcperrinvalremaddr,omitempty"`
	/**
	* MPTCP remove address request received on invalid/unknown address id.
	 */
	Mptcperrinvalremaddrrate float64 `json:"mptcperrinvalremaddrrate,omitempty"`
	Mptcperroptssendrst      int     `json:"mptcperroptssendrst,omitempty"`
	/**
	* MPTCP sent RST on receiving improper option field.
	 */
	Mptcperroptssendrstrate float64 `json:"mptcperroptssendrstrate,omitempty"`
	Mptcperrremaddrself     int     `json:"mptcperrremaddrself,omitempty"`
	/**
	* MPTCP remove address request for self address.
	 */
	Mptcperrremaddrselfrate float64 `json:"mptcperrremaddrselfrate,omitempty"`
	Mptcperrrssffail        int     `json:"mptcperrrssffail,omitempty"`
	/**
	* Add RSS filter to steer traffic to right node on established MPTCP session failed.
	 */
	Mptcperrrssffailrate    float64 `json:"mptcperrrssffailrate,omitempty"`
	Mptcperrnopayloadlenpkt int     `json:"mptcperrnopayloadlenpkt,omitempty"`
	/**
	* MPTCP Payload length not specified in packet.
	 */
	Mptcperrnopayloadlenpktrate      float64 `json:"mptcperrnopayloadlenpktrate,omitempty"`
	Mptcperrunsupportedmssnegotiated int     `json:"mptcperrunsupportedmssnegotiated,omitempty"`
	/**
	* MPTCP Unsupported MSS negotiated error.
	 */
	Mptcperrunsupportedmssnegotiatedrate float64 `json:"mptcperrunsupportedmssnegotiatedrate,omitempty"`
	Mptcperrbadcksum                     int     `json:"mptcperrbadcksum,omitempty"`
	/**
	* MPTCP checksum failed. Connection will fallback to regular tcp.
	 */
	Mptcperrbadcksumrate       float64 `json:"mptcperrbadcksumrate,omitempty"`
	Mptcperrcryptonotsupported int     `json:"mptcperrcryptonotsupported,omitempty"`
	/**
	* MPTCP client crypto algorithm not supported.
	 */
	Mptcperrcryptonotsupportedrate float64 `json:"mptcperrcryptonotsupportedrate,omitempty"`
	Mptcperrversionnotsupported    int     `json:"mptcperrversionnotsupported,omitempty"`
	/**
	* MPTCP MP_CAPABLE request from unsupported mptcp client.
	 */
	Mptcperrversionnotsupportedrate float64 `json:"mptcperrversionnotsupportedrate,omitempty"`
	Mptcpplainackrst                int     `json:"mptcpplainackrst,omitempty"`
	/**
	* MPTCP Sent RST on receiving plain ACK for DSS.
	 */
	Mptcpplainackrstrate   float64 `json:"mptcpplainackrstrate,omitempty"`
	Mptcperrdatafinpassive int     `json:"mptcperrdatafinpassive,omitempty"`
	/**
	* MPTCP Data FIN received on passive subflow
	 */
	Mptcperrdatafinpassiverate float64 `json:"mptcperrdatafinpassiverate,omitempty"`
	Mptcperrfastclose          int     `json:"mptcperrfastclose,omitempty"`
	/**
	* MPTCP FAST CLOSE sent.
	 */
	Mptcperrfastcloserate    float64 `json:"mptcperrfastcloserate,omitempty"`
	Mptcperrfastclosepassive int     `json:"mptcperrfastclosepassive,omitempty"`
	/**
	* MPTCP Fast close received on passive subflow.
	 */
	Mptcperrfastclosepassiverate float64 `json:"mptcperrfastclosepassiverate,omitempty"`
	Mptcperrfastclosekey         int     `json:"mptcperrfastclosekey,omitempty"`
	/**
	* MPTCP FAST_CLOSE received with invalid key and the packet is dropped.
	 */
	Mptcperrfastclosekeyrate float64 `json:"mptcperrfastclosekeyrate,omitempty"`
	Mptcpmpfailsent          int     `json:"mptcpmpfailsent,omitempty"`
	/**
	* MPTCP Total MP_FAIL sent due to checksum failure.
	 */
	Mptcpmpfailsentrate float64 `json:"mptcpmpfailsentrate,omitempty"`
	Mptcpmpfailrecvd    int     `json:"mptcpmpfailrecvd,omitempty"`
	/**
	* MPTCP Total MP_FAIL received and fallback to regular TCP.
	 */
	Mptcpmpfailrecvdrate float64 `json:"mptcpmpfailrecvdrate,omitempty"`
	Mptcperrnomappktrcvd int     `json:"mptcperrnomappktrcvd,omitempty"`
	/**
	* MPTCP Packet received with no Data Sequence Mapping.
	 */
	Mptcperrnomappktrcvdrate float64 `json:"mptcperrnomappktrcvdrate,omitempty"`
	Mptcptotmoredatarcvd     int     `json:"mptcptotmoredatarcvd,omitempty"`
	/**
	* MPTCP More data received than the available Data Sequence Mapping.
	 */
	Mptcpmoredatarcvdrate  float64 `json:"mptcpmoredatarcvdrate,omitempty"`
	Mptcperrbadmapconndrop int     `json:"mptcperrbadmapconndrop,omitempty"`
	/**
	* MPTCP Drop the session incase of invalid Data Sequence map.
	 */
	Mptcperrbadmapconndroprate float64 `json:"mptcperrbadmapconndroprate,omitempty"`
	Mptcperrdupmaprecvd        int     `json:"mptcperrdupmaprecvd,omitempty"`
	/**
	* MPTCP Duplicate maps in Data Sequence map table.
	 */
	Mptcperrdupmaprecvdrate float64 `json:"mptcperrdupmaprecvdrate,omitempty"`
	Mptcperrinvalidsfn      int     `json:"mptcperrinvalidsfn,omitempty"`
	/**
	* MPTCP subflow map doesn't exactly match MPTCP session mapping.
	 */
	Mptcperrinvalidsfnrate float64 `json:"mptcperrinvalidsfnrate,omitempty"`
	Mptcperrmapexists      int     `json:"mptcperrmapexists,omitempty"`
	/**
	* MPTCP sequence map already exists.
	 */
	Mptcperrmapexistsrate float64 `json:"mptcperrmapexistsrate,omitempty"`
	Mptcperrretxpktrcvd   int     `json:"mptcperrretxpktrcvd,omitempty"`
	/**
	* Retransmitted Data Recevied on MPTCP session.
	 */
	Mptcperrretxpktrcvdrate    float64 `json:"mptcperrretxpktrcvdrate,omitempty"`
	Mptcperrsfsessionallocfail int     `json:"mptcperrsfsessionallocfail,omitempty"`
	/**
	* Attaching the subflow to MPTCP session failed due to failure in allocating memory to subflow map table.
	 */
	Mptcperrsfsessionallocfailrate float64 `json:"mptcperrsfsessionallocfailrate,omitempty"`
	Mptcperrmpcapsessionallocfail  int     `json:"mptcperrmpcapsessionallocfail,omitempty"`
	/**
	* Creating a MPTCP connection failed due to failure in allocating memory to MPTCP connection management structure.
	 */
	Mptcperrmpcapsessionallocfailrate float64 `json:"mptcperrmpcapsessionallocfailrate,omitempty"`
	Mptcptotmpcapsfpcballoc           int     `json:"mptcptotmpcapsfpcballoc,omitempty"`
	/**
	* Allocating memory to TCP protocol control block(PCB) for subflow failed.
	 */
	Mptcpmpcapsfpcballocrate float64 `json:"mptcpmpcapsfpcballocrate,omitempty"`
	Mptcptotmpcballocfailed  int     `json:"mptcptotmpcballocfailed,omitempty"`
	/**
	* Allocating memory to MPTCP protocol control block failed.
	 */
	Mptcpmpcballocfailedrate float64 `json:"mptcpmpcballocfailedrate,omitempty"`
	Mptcperrnsballocfailed   int     `json:"mptcperrnsballocfailed,omitempty"`
	/**
	* Failed to allocate memory to output MPTCP packet.
	 */
	Mptcperrnsballocfailedrate float64 `json:"mptcperrnsballocfailedrate,omitempty"`
	Mptcperrnosffreensb        int     `json:"mptcperrnosffreensb,omitempty"`
	/**
	* MPTCP output a packet without any subflow PCB.
	 */
	Mptcperrnosffreensbrate float64 `json:"mptcperrnosffreensbrate,omitempty"`
}
