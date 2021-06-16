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

package cmp


type Cmpstats struct {
	/**
	* Clear the statsistics / counters
	*/
	Clearstats string `json:"clearstats,omitempty"`
	Delbwsaving float64 `json:"delbwsaving,omitempty"`
	Delcmpratio float64 `json:"delcmpratio,omitempty"`
	Decomptcpratio float64 `json:"decomptcpratio,omitempty"`
	Decomptcpbandwidthsaving float64 `json:"decomptcpbandwidthsaving,omitempty"`
	Comptcpratio float64 `json:"comptcpratio,omitempty"`
	Comptcpbandwidthsaving float64 `json:"comptcpbandwidthsaving,omitempty"`
	Comptotaldatacompressionratio float64 `json:"comptotaldatacompressionratio,omitempty"`
	Comphttpbandwidthsaving float64 `json:"comphttpbandwidthsaving,omitempty"`
	Compratio float64 `json:"compratio,omitempty"`
	Comptotalrequests uint64 `json:"comptotalrequests,omitempty"`
	/**
	* Number of HTTP compression requests the Citrix ADC receives for which the response is successfully compressed. For example, after you enable compression and configure services, if you send HTTP requests to the Citrix ADC using a HTTP client that supports compression, and Citrix ADC compresses the corresponding response, this counter is incremented.
	*/
	Comprequestsrate float64 `json:"comprequestsrate,omitempty"`
	Comptotalrxbytes uint64 `json:"comptotalrxbytes,omitempty"`
	/**
	* Number of bytes that can be compressed, which the Citrix ADC receives from the server. This gives the content length of the response that the Citrix ADC receives from server.
	*/
	Comprxbytesrate float64 `json:"comprxbytesrate,omitempty"`
	Comptotaltxbytes uint64 `json:"comptotaltxbytes,omitempty"`
	/**
	* Number of bytes the Citrix ADC sends to the client after compressing the response from the server.
	*/
	Comptxbytesrate float64 `json:"comptxbytesrate,omitempty"`
	Comptotalrxpackets uint64 `json:"comptotalrxpackets,omitempty"`
	/**
	* Number of HTTP packets that can be compressed, which the Citrix ADC receives from the server.
	*/
	Comprxpacketsrate float64 `json:"comprxpacketsrate,omitempty"`
	Comptotaltxpackets uint64 `json:"comptotaltxpackets,omitempty"`
	/**
	* Number of HTTP packets that the Citrix ADC sends to the client after compressing the response from the server.
	*/
	Comptxpacketsrate float64 `json:"comptxpacketsrate,omitempty"`
	Comptcptotalrxbytes uint64 `json:"comptcptotalrxbytes,omitempty"`
	/**
	* Number of bytes that can be compressed, which the Citrix ADC receives from the server. This gives the content length of the response that the Citrix ADC receives from server.
	*/
	Comptcprxbytesrate float64 `json:"comptcprxbytesrate,omitempty"`
	Comptcptotalrxpackets uint64 `json:"comptcptotalrxpackets,omitempty"`
	/**
	* Total number of compressible packets received by Citrix ADC.
	*/
	Comptcprxpacketsrate float64 `json:"comptcprxpacketsrate,omitempty"`
	Comptcptotaltxbytes uint64 `json:"comptcptotaltxbytes,omitempty"`
	/**
	* Number of bytes that the Citrix ADC sends to the client after compressing the response from the server.
	*/
	Comptcptxbytesrate float64 `json:"comptcptxbytesrate,omitempty"`
	Comptcptotaltxpackets uint64 `json:"comptcptotaltxpackets,omitempty"`
	/**
	* Number of TCP packets that the Citrix ADC sends to the client after compressing the response from the server.
	*/
	Comptcptxpacketsrate float64 `json:"comptcptxpacketsrate,omitempty"`
	Comptcptotalquantum uint64 `json:"comptcptotalquantum,omitempty"`
	/**
	* Number of times the Citrix ADC compresses a quantum of data.  Citrix ADC buffers the data received from the server till it reaches the quantum size and then compresses the buffered data and transmits to the client.
	*/
	Comptcpquantumrate float64 `json:"comptcpquantumrate,omitempty"`
	Comptcptotalpush uint64 `json:"comptcptotalpush,omitempty"`
	/**
	* Number of times the Citrix ADC compresses data on receiving a TCP PUSH flag from the server. The PUSH flag ensures that data is compressed immediately without waiting for the buffered data size to reach the quantum size.
	*/
	Comptcppushrate float64 `json:"comptcppushrate,omitempty"`
	Comptcptotaleoi uint64 `json:"comptcptotaleoi,omitempty"`
	/**
	* Number of times the Citrix ADC compresses data on receiving End Of Input (FIN packet).  When the Citrix ADC  receives End Of Input (FIN packet), it compresses the buffered data immediately without waiting for the buffered data size to reach the quantum size.
	*/
	Comptcpeoirate float64 `json:"comptcpeoirate,omitempty"`
	Comptcptotaltimer uint64 `json:"comptcptotaltimer,omitempty"`
	/**
	* Number of times the Citrix ADC compresses data on expiration of data accumulation timer. The timer expires if the server response is very slow and consequently, the Citrix ADC does not receive response for a certain amount of time.  Under such a condition, the Citrix ADC compresses the buffered data immediately without waiting for the buffered data size to reach the quantum size.
	*/
	Comptcptimerrate float64 `json:"comptcptimerrate,omitempty"`
	Decomptcprxbytes uint64 `json:"decomptcprxbytes,omitempty"`
	/**
	* Total number of compressed bytes received by Citrix ADC.
	*/
	Decomptcprxbytesrate float64 `json:"decomptcprxbytesrate,omitempty"`
	Decomptcprxpackets uint64 `json:"decomptcprxpackets,omitempty"`
	/**
	* Total number of compressed packets received by Citrix ADC.
	*/
	Decomptcprxpacketsrate float64 `json:"decomptcprxpacketsrate,omitempty"`
	Decomptcptxbytes uint64 `json:"decomptcptxbytes,omitempty"`
	/**
	* Total number of decompressed bytes transmitted by Citrix ADC.
	*/
	Decomptcptxbytesrate float64 `json:"decomptcptxbytesrate,omitempty"`
	Decomptcptxpackets uint64 `json:"decomptcptxpackets,omitempty"`
	/**
	* Total number of decompressed packets transmitted by Citrix ADC.
	*/
	Decomptcptxpacketsrate float64 `json:"decomptcptxpacketsrate,omitempty"`
	Decomptcperrdata uint64 `json:"decomptcperrdata,omitempty"`
	/**
	* Number of data errors encountered while decompressing.
	*/
	Decomptcperrdatarate float64 `json:"decomptcperrdatarate,omitempty"`
	Decomptcperrlessdata uint64 `json:"decomptcperrlessdata,omitempty"`
	/**
	* Number of times Citrix ADC received less data than declared by protocol.
	*/
	Decomptcperrlessdatarate float64 `json:"decomptcperrlessdatarate,omitempty"`
	Decomptcperrmoredata uint64 `json:"decomptcperrmoredata,omitempty"`
	/**
	* Number of times Citrix ADC received more data than declared by protocol.
	*/
	Decomptcperrmoredatarate float64 `json:"decomptcperrmoredatarate,omitempty"`
	Decomptcperrmemory uint64 `json:"decomptcperrmemory,omitempty"`
	/**
	* Number of times memory failures occurred while decompressing.
	*/
	Decomptcperrmemoryrate float64 `json:"decomptcperrmemoryrate,omitempty"`
	Decomptcperrunknown uint64 `json:"decomptcperrunknown,omitempty"`
	/**
	* Number of times unknown errors occurred while decompressing.
	*/
	Decomptcperrunknownrate float64 `json:"decomptcperrunknownrate,omitempty"`
	Delcomptotalrequests uint64 `json:"delcomptotalrequests,omitempty"`
	/**
	* Total number of delta compression requests received by Citrix ADC.
	*/
	Delcomprequestsrate float64 `json:"delcomprequestsrate,omitempty"`
	Delcompdone uint64 `json:"delcompdone,omitempty"`
	/**
	* Total number of delta compressions done by Citrix ADC.
	*/
	Delcompdonerate float64 `json:"delcompdonerate,omitempty"`
	Delcomptcprxbytes uint64 `json:"delcomptcprxbytes,omitempty"`
	/**
	* Total number of delta-compressible bytes received by Citrix ADC.
	*/
	Delcomptcprxbytesrate float64 `json:"delcomptcprxbytesrate,omitempty"`
	Delcomptcptxbytes uint64 `json:"delcomptcptxbytes,omitempty"`
	/**
	* Total number of delta-compressed bytes transmitted by Citrix ADC.
	*/
	Delcomptcptxbytesrate float64 `json:"delcomptcptxbytesrate,omitempty"`
	Delcompfirstaccess uint64 `json:"delcompfirstaccess,omitempty"`
	/**
	* Total number of delta compression first accesses.
	*/
	Delcompfirstaccessrate float64 `json:"delcompfirstaccessrate,omitempty"`
	Delcomptcprxpackets uint64 `json:"delcomptcprxpackets,omitempty"`
	/**
	* Number of delta-compressible packets received.
	*/
	Delcomptcprxpacketsrate float64 `json:"delcomptcprxpacketsrate,omitempty"`
	Delcomptcptxpackets uint64 `json:"delcomptcptxpackets,omitempty"`
	/**
	* Total number of delta-compressed packets transmitted by Citrix ADC.
	*/
	Delcomptcptxpacketsrate float64 `json:"delcomptcptxpacketsrate,omitempty"`
	Delcompbaseserved uint64 `json:"delcompbaseserved,omitempty"`
	/**
	* Total number of basefile requests served by Citrix ADC.
	*/
	Delcompbaseservedrate float64 `json:"delcompbaseservedrate,omitempty"`
	Delcompbasetcptxbytes uint64 `json:"delcompbasetcptxbytes,omitempty"`
	/**
	* Number of basefile bytes transmitted by Citrix ADC.
	*/
	Delcompbasetcptxbytesrate float64 `json:"delcompbasetcptxbytesrate,omitempty"`
	Delcomperrbypassed uint64 `json:"delcomperrbypassed,omitempty"`
	/**
	* Number of times delta-compression bypassed by Citrix ADC.
	*/
	Delcomperrbypassedrate float64 `json:"delcomperrbypassedrate,omitempty"`
	Delcomperrbfilewhdrfailed uint64 `json:"delcomperrbfilewhdrfailed,omitempty"`
	/**
	* Number of times basefile could not be updated in Citrix ADC cache.
	*/
	Delcomperrbfilewhdrfailedrate float64 `json:"delcomperrbfilewhdrfailedrate,omitempty"`
	Delcomperrnostoremiss uint64 `json:"delcomperrnostoremiss,omitempty"`
	/**
	* Number of times basefile was not found in Citrix ADC cache.
	*/
	Delcomperrnostoremissrate float64 `json:"delcomperrnostoremissrate,omitempty"`
	Delcomperrreqinfotoobig uint64 `json:"delcomperrreqinfotoobig,omitempty"`
	/**
	* Number of times basefile request URL was too large.
	*/
	Delcomperrreqinfotoobigrate float64 `json:"delcomperrreqinfotoobigrate,omitempty"`
	Delcomperrreqinfoallocfail uint64 `json:"delcomperrreqinfoallocfail,omitempty"`
	/**
	* Number of times requested basefile could not be allocated.
	*/
	Delcomperrreqinfoallocfailrate float64 `json:"delcomperrreqinfoallocfailrate,omitempty"`
	Delcomperrsessallocfail uint64 `json:"delcomperrsessallocfail,omitempty"`
	/**
	* Number of times delta compression session could not be allocated.
	*/
	Delcomperrsessallocfailrate float64 `json:"delcomperrsessallocfailrate,omitempty"`

}
