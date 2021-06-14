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

package vpn


type Vpnstats struct {
	/**
	* Clear the statsistics / counters
	*/
	Clearstats string `json:"clearstats,omitempty"`
	Indexhtmlhit uint64 `json:"indexhtmlhit,omitempty"`
	Indexhtmlnoserved uint64 `json:"indexhtmlnoserved,omitempty"`
	Cfghtmlserved uint64 `json:"cfghtmlserved,omitempty"`
	/**
	* Number of client configuration requests received by VPN server.
	*/
	Cfghtmlservedrate int32 `json:"cfghtmlservedrate,omitempty"`
	Dnsreqhit uint64 `json:"dnsreqhit,omitempty"`
	/**
	* Number of DNS queries resolved by VPN server.
	*/
	Dnsreqhitrate int32 `json:"dnsreqhitrate,omitempty"`
	Winsrequesthit uint64 `json:"winsrequesthit,omitempty"`
	/**
	* Number of WINS queries resolved by VPN server.
	*/
	Winsrequesthitrate int32 `json:"winsrequesthitrate,omitempty"`
	Csrequesthit uint64 `json:"csrequesthit,omitempty"`
	/**
	* Number of SSL VPN tunnels formed between VPN server and client.
	*/
	Csrequesthitrate int32 `json:"csrequesthitrate,omitempty"`
	Csnonhttpprobehit uint64 `json:"csnonhttpprobehit,omitempty"`
	/**
	* Number of probes from VPN to back-end non-HTTP servers that have been accessed by the VPN client.
	*/
	Csnonhttpprobehitrate int32 `json:"csnonhttpprobehitrate,omitempty"`
	Cshttpprobehit uint64 `json:"cshttpprobehit,omitempty"`
	/**
	* Number of probes from VPN to back-end HTTP servers that have been accessed by the VPN client.
	*/
	Cshttpprobehitrate int32 `json:"cshttpprobehitrate,omitempty"`
	Totalcsconnsucc uint64 `json:"totalcsconnsucc,omitempty"`
	/**
	* Number of successful probes to all back-end servers.
	*/
	Csconnsuccrate int32 `json:"csconnsuccrate,omitempty"`
	Totalfsrequest uint64 `json:"totalfsrequest,omitempty"`
	/**
	* Number of file system requests received by VPN server.
	*/
	Fsrequestrate int32 `json:"fsrequestrate,omitempty"`
	Vpnlicensefail uint64 `json:"vpnlicensefail,omitempty"`
	Iipdisabledmipused uint64 `json:"iipdisabledmipused,omitempty"`
	/**
	* Number of times SNIP is used as IIP is disabled.
	*/
	Iipdisabledmipusedrate int32 `json:"iipdisabledmipusedrate,omitempty"`
	Iipfailedmipused uint64 `json:"iipfailedmipused,omitempty"`
	/**
	* Number of times SNIP is used as IIP assignment failed.
	*/
	Iipfailedmipusedrate int32 `json:"iipfailedmipusedrate,omitempty"`
	Iipspillovermipused uint64 `json:"iipspillovermipused,omitempty"`
	/**
	* Number of times SNIP is used on IIP Spillover.
	*/
	Iipspillovermipusedrate int32 `json:"iipspillovermipusedrate,omitempty"`
	Iipdisabledmipdisabled uint64 `json:"iipdisabledmipdisabled,omitempty"`
	/**
	* Both IIP and SNIP is disabled.
	*/
	Iipdisabledmipdisabledrate int32 `json:"iipdisabledmipdisabledrate,omitempty"`
	Iipfailedmipdisabled uint64 `json:"iipfailedmipdisabled,omitempty"`
	/**
	* Number of times IIP assignment failed and SNIP is disabled.
	*/
	Iipfailedmipdisabledrate int32 `json:"iipfailedmipdisabledrate,omitempty"`
	Csgtotalconnectedusers uint64 `json:"csgtotalconnectedusers,omitempty"`
	/**
	* Number of total users connected across VPN vservers
	*/
	Csgconnectedusersrate int32 `json:"csgconnectedusersrate,omitempty"`
	Socksmethreqrcvd uint64 `json:"socksmethreqrcvd,omitempty"`
	/**
	* Number of received SOCKS method request.
	*/
	Socksmethreqrcvdrate int32 `json:"socksmethreqrcvdrate,omitempty"`
	Socksmethreqsent uint64 `json:"socksmethreqsent,omitempty"`
	/**
	* Number of sent SOCKS method request.
	*/
	Socksmethreqsentrate int32 `json:"socksmethreqsentrate,omitempty"`
	Socksmethresprcvd uint64 `json:"socksmethresprcvd,omitempty"`
	/**
	* Number of received SOCKS method response.
	*/
	Socksmethresprcvdrate int32 `json:"socksmethresprcvdrate,omitempty"`
	Socksmethrespsent uint64 `json:"socksmethrespsent,omitempty"`
	/**
	* Number of sent SOCKS method response.
	*/
	Socksmethrespsentrate int32 `json:"socksmethrespsentrate,omitempty"`
	Socksconnreqrcvd uint64 `json:"socksconnreqrcvd,omitempty"`
	/**
	* Number of received SOCKS connect request.
	*/
	Socksconnreqrcvdrate int32 `json:"socksconnreqrcvdrate,omitempty"`
	Socksconnreqsent uint64 `json:"socksconnreqsent,omitempty"`
	/**
	* Number of sent SOCKS connect request.
	*/
	Socksconnreqsentrate int32 `json:"socksconnreqsentrate,omitempty"`
	Socksconnresprcvd uint64 `json:"socksconnresprcvd,omitempty"`
	/**
	* Number of received SOCKS connect response.
	*/
	Socksconnresprcvdrate int32 `json:"socksconnresprcvdrate,omitempty"`
	Socksconnrespsent uint64 `json:"socksconnrespsent,omitempty"`
	/**
	* Number of sent SOCKS connect response.
	*/
	Socksconnrespsentrate int32 `json:"socksconnrespsentrate,omitempty"`
	Socksservererror uint64 `json:"socksservererror,omitempty"`
	/**
	* Number of SOCKS server error.
	*/
	Socksservererrorrate int32 `json:"socksservererrorrate,omitempty"`
	Socksclienterror uint64 `json:"socksclienterror,omitempty"`
	/**
	* Number of SOCKS client error.
	*/
	Socksclienterrorrate int32 `json:"socksclienterrorrate,omitempty"`
	Staconnsuccess uint64 `json:"staconnsuccess,omitempty"`
	/**
	* Number of STA connection success.
	*/
	Staconnsuccessrate int32 `json:"staconnsuccessrate,omitempty"`
	Staconnfailure uint64 `json:"staconnfailure,omitempty"`
	/**
	* Number of STA connection failure.
	*/
	Staconnfailurerate int32 `json:"staconnfailurerate,omitempty"`
	Cpsconnsuccess uint64 `json:"cpsconnsuccess,omitempty"`
	/**
	* Number of CPS connection success.
	*/
	Cpsconnsuccessrate int32 `json:"cpsconnsuccessrate,omitempty"`
	Cpsconnfailure uint64 `json:"cpsconnfailure,omitempty"`
	/**
	* Number of CPS connection failure.
	*/
	Cpsconnfailurerate int32 `json:"cpsconnfailurerate,omitempty"`
	Starequestsent uint64 `json:"starequestsent,omitempty"`
	/**
	* Number of STA request sent.
	*/
	Starequestsentrate int32 `json:"starequestsentrate,omitempty"`
	Staresponserecvd uint64 `json:"staresponserecvd,omitempty"`
	/**
	* Number of STA response received.
	*/
	Staresponserecvdrate int32 `json:"staresponserecvdrate,omitempty"`
	Icalicensefailure uint64 `json:"icalicensefailure,omitempty"`
	/**
	* Number of ICA license failure.
	*/
	Icalicensefailurerate int32 `json:"icalicensefailurerate,omitempty"`
	Stamonsent uint64 `json:"stamonsent,omitempty"`
	/**
	* Number of STA monitor requests sent.
	*/
	Stamonsentrate int32 `json:"stamonsentrate,omitempty"`
	Stamonrcvd uint64 `json:"stamonrcvd,omitempty"`
	/**
	* Number of STA monitor responses recieved.
	*/
	Stamonrcvdrate int32 `json:"stamonrcvdrate,omitempty"`
	Stamonsucc uint64 `json:"stamonsucc,omitempty"`
	/**
	* Number of STA monitor successful responses.
	*/
	Stamonsuccrate int32 `json:"stamonsuccrate,omitempty"`
	Stamonfail uint64 `json:"stamonfail,omitempty"`
	/**
	* Number of STA monitor failed responses.
	*/
	Stamonfailrate int32 `json:"stamonfailrate,omitempty"`
	Csgptktvalidatenotstarted uint64 `json:"csgptktvalidatenotstarted,omitempty"`
	/**
	* Total number of STA server lookup failures for auth-id in primary ticket
	*/
	Csgptktvalidatenotstartedrate int32 `json:"csgptktvalidatenotstartedrate,omitempty"`
	Csgrtktvalidatenotstarted uint64 `json:"csgrtktvalidatenotstarted,omitempty"`
	/**
	* Total number of STA server lookup failures for auth-id in redundant ticket
	*/
	Csgrtktvalidatenotstartedrate int32 `json:"csgrtktvalidatenotstartedrate,omitempty"`

}
