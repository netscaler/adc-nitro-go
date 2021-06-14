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
* Statistics for icmp resource.
*/

type Protocolicmpstats struct {
	/**
	* Clear the statsistics / counters
	*/
	Clearstats string `json:"clearstats,omitempty"`
	Icmptotrxpkts uint64 `json:"icmptotrxpkts,omitempty"`
	/**
	* ICMP packets received.
	*/
	Icmprxpktsrate int32 `json:"icmprxpktsrate,omitempty"`
	Icmptotrxbytes uint64 `json:"icmptotrxbytes,omitempty"`
	/**
	* Bytes of ICMP data received.
	*/
	Icmprxbytesrate int32 `json:"icmprxbytesrate,omitempty"`
	Icmptottxpkts uint64 `json:"icmptottxpkts,omitempty"`
	/**
	* ICMP packets transmitted.
	*/
	Icmptxpktsrate int32 `json:"icmptxpktsrate,omitempty"`
	Icmptottxbytes uint64 `json:"icmptottxbytes,omitempty"`
	/**
	* Bytes of ICMP data transmitted.
	*/
	Icmptxbytesrate int32 `json:"icmptxbytesrate,omitempty"`
	Icmptotrxechoreply uint64 `json:"icmptotrxechoreply,omitempty"`
	/**
	* ICMP Ping echo replies received.
	*/
	Icmprxechoreplyrate int32 `json:"icmprxechoreplyrate,omitempty"`
	Icmptottxechoreply uint64 `json:"icmptottxechoreply,omitempty"`
	/**
	* ICMP Ping echo replies transmitted.
	*/
	Icmptxechoreplyrate int32 `json:"icmptxechoreplyrate,omitempty"`
	Icmptotrxecho uint64 `json:"icmptotrxecho,omitempty"`
	/**
	* ICMP Ping Echo Request and Echo Reply packets received.
	*/
	Icmprxechorate int32 `json:"icmprxechorate,omitempty"`
	Icmptotdstiplookup uint64 `json:"icmptotdstiplookup,omitempty"`
	Icmpcurratethreshold uint32 `json:"icmpcurratethreshold,omitempty"`
	Icmptotportunreachablerx uint64 `json:"icmptotportunreachablerx,omitempty"`
	Icmptotportunreachabletx uint64 `json:"icmptotportunreachabletx,omitempty"`
	Icmptotneedfragrx uint64 `json:"icmptotneedfragrx,omitempty"`
	Icmptotthresholdexceeds uint64 `json:"icmptotthresholdexceeds,omitempty"`
	Icmptotpktsdropped uint64 `json:"icmptotpktsdropped,omitempty"`
	Icmptotbadchecksum uint64 `json:"icmptotbadchecksum,omitempty"`
	Icmptotnonfirstipfrag uint64 `json:"icmptotnonfirstipfrag,omitempty"`
	Icmptotinvalidbodylen uint64 `json:"icmptotinvalidbodylen,omitempty"`
	Icmptotnotcpconn uint64 `json:"icmptotnotcpconn,omitempty"`
	Icmptotnoudpconn uint64 `json:"icmptotnoudpconn,omitempty"`
	Icmptotinvalidtcpseqno uint64 `json:"icmptotinvalidtcpseqno,omitempty"`
	Icmptotinvalidnextmtuval uint64 `json:"icmptotinvalidnextmtuval,omitempty"`
	Icmptotbignextmtu uint64 `json:"icmptotbignextmtu,omitempty"`
	Icmptotinvalidprotocol uint64 `json:"icmptotinvalidprotocol,omitempty"`
	Icmptotbadpmtuipchecksum uint64 `json:"icmptotbadpmtuipchecksum,omitempty"`
	Icmptotpmtunolink uint64 `json:"icmptotpmtunolink,omitempty"`
	Icmptotpmtudiscoverydisabled uint64 `json:"icmptotpmtudiscoverydisabled,omitempty"`

}
