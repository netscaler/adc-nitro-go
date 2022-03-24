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
	Clearstats    string `json:"clearstats,omitempty"`
	Icmptotrxpkts int    `json:"icmptotrxpkts,omitempty"`
	/**
	* ICMP packets received.
	 */
	Icmprxpktsrate float64 `json:"icmprxpktsrate,omitempty"`
	Icmptotrxbytes int     `json:"icmptotrxbytes,omitempty"`
	/**
	* Bytes of ICMP data received.
	 */
	Icmprxbytesrate float64 `json:"icmprxbytesrate,omitempty"`
	Icmptottxpkts   int     `json:"icmptottxpkts,omitempty"`
	/**
	* ICMP packets transmitted.
	 */
	Icmptxpktsrate float64 `json:"icmptxpktsrate,omitempty"`
	Icmptottxbytes int     `json:"icmptottxbytes,omitempty"`
	/**
	* Bytes of ICMP data transmitted.
	 */
	Icmptxbytesrate    float64 `json:"icmptxbytesrate,omitempty"`
	Icmptotrxechoreply int     `json:"icmptotrxechoreply,omitempty"`
	/**
	* ICMP Ping echo replies received.
	 */
	Icmprxechoreplyrate float64 `json:"icmprxechoreplyrate,omitempty"`
	Icmptottxechoreply  int     `json:"icmptottxechoreply,omitempty"`
	/**
	* ICMP Ping echo replies transmitted.
	 */
	Icmptxechoreplyrate float64 `json:"icmptxechoreplyrate,omitempty"`
	Icmptotrxecho       int     `json:"icmptotrxecho,omitempty"`
	/**
	* ICMP Ping Echo Request and Echo Reply packets received.
	 */
	Icmprxechorate               float64 `json:"icmprxechorate,omitempty"`
	Icmptotdstiplookup           int     `json:"icmptotdstiplookup,omitempty"`
	Icmpcurratethreshold         int     `json:"icmpcurratethreshold,omitempty"`
	Icmptotportunreachablerx     int     `json:"icmptotportunreachablerx,omitempty"`
	Icmptotportunreachabletx     int     `json:"icmptotportunreachabletx,omitempty"`
	Icmptotneedfragrx            int     `json:"icmptotneedfragrx,omitempty"`
	Icmptotthresholdexceeds      int     `json:"icmptotthresholdexceeds,omitempty"`
	Icmptotpktsdropped           int     `json:"icmptotpktsdropped,omitempty"`
	Icmptotbadchecksum           int     `json:"icmptotbadchecksum,omitempty"`
	Icmptotnonfirstipfrag        int     `json:"icmptotnonfirstipfrag,omitempty"`
	Icmptotinvalidbodylen        int     `json:"icmptotinvalidbodylen,omitempty"`
	Icmptotnotcpconn             int     `json:"icmptotnotcpconn,omitempty"`
	Icmptotnoudpconn             int     `json:"icmptotnoudpconn,omitempty"`
	Icmptotinvalidtcpseqno       int     `json:"icmptotinvalidtcpseqno,omitempty"`
	Icmptotinvalidnextmtuval     int     `json:"icmptotinvalidnextmtuval,omitempty"`
	Icmptotbignextmtu            int     `json:"icmptotbignextmtu,omitempty"`
	Icmptotinvalidprotocol       int     `json:"icmptotinvalidprotocol,omitempty"`
	Icmptotbadpmtuipchecksum     int     `json:"icmptotbadpmtuipchecksum,omitempty"`
	Icmptotpmtunolink            int     `json:"icmptotpmtunolink,omitempty"`
	Icmptotpmtudiscoverydisabled int     `json:"icmptotpmtudiscoverydisabled,omitempty"`
}
