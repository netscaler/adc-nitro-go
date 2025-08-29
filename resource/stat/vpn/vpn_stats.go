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
	Indexhtmlhit int `json:"indexhtmlhit,omitempty"`
	Indexhtmlnoserved int `json:"indexhtmlnoserved,omitempty"`
	Cfghtmlserved int `json:"cfghtmlserved,omitempty"`
	/**
	* Number of client configuration requests received by VPN server.
	*/
	Cfghtmlservedrate float64 `json:"cfghtmlservedrate,omitempty"`
	Dnsreqhit int `json:"dnsreqhit,omitempty"`
	/**
	* Number of DNS queries resolved by VPN server.
	*/
	Dnsreqhitrate float64 `json:"dnsreqhitrate,omitempty"`
	Winsrequesthit int `json:"winsrequesthit,omitempty"`
	/**
	* Number of WINS queries resolved by VPN server.
	*/
	Winsrequesthitrate float64 `json:"winsrequesthitrate,omitempty"`
	Csrequesthit int `json:"csrequesthit,omitempty"`
	/**
	* Number of SSL VPN tunnels formed between VPN server and client.
	*/
	Csrequesthitrate float64 `json:"csrequesthitrate,omitempty"`
	Csnonhttpprobehit int `json:"csnonhttpprobehit,omitempty"`
	/**
	* Number of probes from VPN to back-end non-HTTP servers that have been accessed by the VPN client.
	*/
	Csnonhttpprobehitrate float64 `json:"csnonhttpprobehitrate,omitempty"`
	Cshttpprobehit int `json:"cshttpprobehit,omitempty"`
	/**
	* Number of probes from VPN to back-end HTTP servers that have been accessed by the VPN client.
	*/
	Cshttpprobehitrate float64 `json:"cshttpprobehitrate,omitempty"`
	Totalcsconnsucc int `json:"totalcsconnsucc,omitempty"`
	/**
	* Number of successful probes to all back-end servers.
	*/
	Csconnsuccrate float64 `json:"csconnsuccrate,omitempty"`
	Totalfsrequest int `json:"totalfsrequest,omitempty"`
	/**
	* Number of file system requests received by VPN server.
	*/
	Fsrequestrate float64 `json:"fsrequestrate,omitempty"`
	Vpnlicensefail int `json:"vpnlicensefail,omitempty"`
	Svpnhttpdtotalrequests int `json:"svpnhttpdtotalrequests,omitempty"`
	Svpnhttpdtotalresponses int `json:"svpnhttpdtotalresponses,omitempty"`
	Svpnhttpdtotalrequestscachehits int `json:"svpnhttpdtotalrequestscachehits,omitempty"`
	Svpnhttpdcurserverconnectionsopen int `json:"svpnhttpdcurserverconnectionsopen,omitempty"`
	Svpnreqwaftotblock int `json:"svpnreqwaftotblock,omitempty"`
	Iipdisabledmipused int `json:"iipdisabledmipused,omitempty"`
	/**
	* Number of times SNIP is used as IIP is disabled.
	*/
	Iipdisabledmipusedrate float64 `json:"iipdisabledmipusedrate,omitempty"`
	Iipfailedmipused int `json:"iipfailedmipused,omitempty"`
	/**
	* Number of times SNIP is used as IIP assignment failed.
	*/
	Iipfailedmipusedrate float64 `json:"iipfailedmipusedrate,omitempty"`
	Iipspillovermipused int `json:"iipspillovermipused,omitempty"`
	/**
	* Number of times SNIP is used on IIP Spillover.
	*/
	Iipspillovermipusedrate float64 `json:"iipspillovermipusedrate,omitempty"`
	Iipdisabledmipdisabled int `json:"iipdisabledmipdisabled,omitempty"`
	/**
	* Both IIP and SNIP is disabled.
	*/
	Iipdisabledmipdisabledrate float64 `json:"iipdisabledmipdisabledrate,omitempty"`
	Iipfailedmipdisabled int `json:"iipfailedmipdisabled,omitempty"`
	/**
	* Number of times IIP assignment failed and SNIP is disabled.
	*/
	Iipfailedmipdisabledrate float64 `json:"iipfailedmipdisabledrate,omitempty"`
	Csgtotalconnectedusers int `json:"csgtotalconnectedusers,omitempty"`
	/**
	* Number of total users connected across VPN vservers
	*/
	Csgconnectedusersrate float64 `json:"csgconnectedusersrate,omitempty"`
	Socksmethreqrcvd int `json:"socksmethreqrcvd,omitempty"`
	/**
	* Number of received SOCKS method request.
	*/
	Socksmethreqrcvdrate float64 `json:"socksmethreqrcvdrate,omitempty"`
	Socksmethreqsent int `json:"socksmethreqsent,omitempty"`
	/**
	* Number of sent SOCKS method request.
	*/
	Socksmethreqsentrate float64 `json:"socksmethreqsentrate,omitempty"`
	Socksmethresprcvd int `json:"socksmethresprcvd,omitempty"`
	/**
	* Number of received SOCKS method response.
	*/
	Socksmethresprcvdrate float64 `json:"socksmethresprcvdrate,omitempty"`
	Socksmethrespsent int `json:"socksmethrespsent,omitempty"`
	/**
	* Number of sent SOCKS method response.
	*/
	Socksmethrespsentrate float64 `json:"socksmethrespsentrate,omitempty"`
	Socksconnreqrcvd int `json:"socksconnreqrcvd,omitempty"`
	/**
	* Number of received SOCKS connect request.
	*/
	Socksconnreqrcvdrate float64 `json:"socksconnreqrcvdrate,omitempty"`
	Socksconnreqsent int `json:"socksconnreqsent,omitempty"`
	/**
	* Number of sent SOCKS connect request.
	*/
	Socksconnreqsentrate float64 `json:"socksconnreqsentrate,omitempty"`
	Socksconnresprcvd int `json:"socksconnresprcvd,omitempty"`
	/**
	* Number of received SOCKS connect response.
	*/
	Socksconnresprcvdrate float64 `json:"socksconnresprcvdrate,omitempty"`
	Socksconnrespsent int `json:"socksconnrespsent,omitempty"`
	/**
	* Number of sent SOCKS connect response.
	*/
	Socksconnrespsentrate float64 `json:"socksconnrespsentrate,omitempty"`
	Socksservererror int `json:"socksservererror,omitempty"`
	/**
	* Number of SOCKS server error.
	*/
	Socksservererrorrate float64 `json:"socksservererrorrate,omitempty"`
	Socksclienterror int `json:"socksclienterror,omitempty"`
	/**
	* Number of SOCKS client error.
	*/
	Socksclienterrorrate float64 `json:"socksclienterrorrate,omitempty"`
	Staconnsuccess int `json:"staconnsuccess,omitempty"`
	/**
	* Number of STA connection success.
	*/
	Staconnsuccessrate float64 `json:"staconnsuccessrate,omitempty"`
	Staconnfailure int `json:"staconnfailure,omitempty"`
	/**
	* Number of STA connection failure.
	*/
	Staconnfailurerate float64 `json:"staconnfailurerate,omitempty"`
	Cpsconnsuccess int `json:"cpsconnsuccess,omitempty"`
	/**
	* Number of CPS connection success.
	*/
	Cpsconnsuccessrate float64 `json:"cpsconnsuccessrate,omitempty"`
	Cpsconnfailure int `json:"cpsconnfailure,omitempty"`
	/**
	* Number of CPS connection failure.
	*/
	Cpsconnfailurerate float64 `json:"cpsconnfailurerate,omitempty"`
	Starequestsent int `json:"starequestsent,omitempty"`
	/**
	* Number of STA request sent.
	*/
	Starequestsentrate float64 `json:"starequestsentrate,omitempty"`
	Staresponserecvd int `json:"staresponserecvd,omitempty"`
	/**
	* Number of STA response received.
	*/
	Staresponserecvdrate float64 `json:"staresponserecvdrate,omitempty"`
	Icalicensefailure int `json:"icalicensefailure,omitempty"`
	/**
	* Number of ICA license failure.
	*/
	Icalicensefailurerate float64 `json:"icalicensefailurerate,omitempty"`
	Stamonsent int `json:"stamonsent,omitempty"`
	/**
	* Number of STA monitor requests sent.
	*/
	Stamonsentrate float64 `json:"stamonsentrate,omitempty"`
	Stamonrcvd int `json:"stamonrcvd,omitempty"`
	/**
	* Number of STA monitor responses recieved.
	*/
	Stamonrcvdrate float64 `json:"stamonrcvdrate,omitempty"`
	Stamonsucc int `json:"stamonsucc,omitempty"`
	/**
	* Number of STA monitor successful responses.
	*/
	Stamonsuccrate float64 `json:"stamonsuccrate,omitempty"`
	Stamonfail int `json:"stamonfail,omitempty"`
	/**
	* Number of STA monitor failed responses.
	*/
	Stamonfailrate float64 `json:"stamonfailrate,omitempty"`
	Csgptktvalidatenotstarted int `json:"csgptktvalidatenotstarted,omitempty"`
	/**
	* Total number of STA server lookup failures for auth-id in primary ticket
	*/
	Csgptktvalidatenotstartedrate float64 `json:"csgptktvalidatenotstartedrate,omitempty"`
	Csgrtktvalidatenotstarted int `json:"csgrtktvalidatenotstarted,omitempty"`
	/**
	* Total number of STA server lookup failures for auth-id in redundant ticket
	*/
	Csgrtktvalidatenotstartedrate float64 `json:"csgrtktvalidatenotstartedrate,omitempty"`
	Totalsbprequest int `json:"totalsbprequest,omitempty"`
	/**
	* Number of securebrowse requests received
	*/
	Sbprequestrate float64 `json:"sbprequestrate,omitempty"`
	Totalsbpsaasrequest int `json:"totalsbpsaasrequest,omitempty"`
	/**
	* Number of securebrowse requests received for saas apps
	*/
	Sbpsaasrequestrate float64 `json:"sbpsaasrequestrate,omitempty"`
	Totalsbpwebrequest int `json:"totalsbpwebrequest,omitempty"`
	/**
	* Number of securebrowse requests received for web apps
	*/
	Sbpwebrequestrate float64 `json:"sbpwebrequestrate,omitempty"`
	Totalsbprequestsuccess int `json:"totalsbprequestsuccess,omitempty"`
	/**
	* Number of securebrowse requests succeeded
	*/
	Sbprequestsuccessrate float64 `json:"sbprequestsuccessrate,omitempty"`
	Totalsbpsaasrequestsuccess int `json:"totalsbpsaasrequestsuccess,omitempty"`
	/**
	* Number of securebrowse requests succeeded for saas apps
	*/
	Sbpsaasrequestsuccessrate float64 `json:"sbpsaasrequestsuccessrate,omitempty"`
	Totalsbpwebrequestsuccess int `json:"totalsbpwebrequestsuccess,omitempty"`
	/**
	* Number of securebrowse requests succeeded for web apps
	*/
	Sbpwebrequestsuccessrate float64 `json:"sbpwebrequestsuccessrate,omitempty"`
	Totalsbprequestfailure int `json:"totalsbprequestfailure,omitempty"`
	/**
	* Number of securebrowse requests failed
	*/
	Sbprequestfailurerate float64 `json:"sbprequestfailurerate,omitempty"`
	Totalsbpsaasrequestfailure int `json:"totalsbpsaasrequestfailure,omitempty"`
	/**
	* Number of securebrowse requests failed for saas apps
	*/
	Sbpsaasrequestfailurerate float64 `json:"sbpsaasrequestfailurerate,omitempty"`
	Totalsbpwebrequestfailure int `json:"totalsbpwebrequestfailure,omitempty"`
	/**
	* Number of securebrowse requests failed for web apps
	*/
	Sbpwebrequestfailurerate float64 `json:"sbpwebrequestfailurerate,omitempty"`
	Totalsbpdnsrequestfailure int `json:"totalsbpdnsrequestfailure,omitempty"`
	/**
	* Number of DNS failures while using securebrowse protocol
	*/
	Sbpdnsrequestfailurerate float64 `json:"sbpdnsrequestfailurerate,omitempty"`
	Totalsbpbackendconnectionfailure int `json:"totalsbpbackendconnectionfailure,omitempty"`
	/**
	* Number of backend connection failure while using securebrowse requests
	*/
	Sbpbackendconnectionfailurerate float64 `json:"sbpbackendconnectionfailurerate,omitempty"`
	Totalsbpconnectorfailure int `json:"totalsbpconnectorfailure,omitempty"`
	/**
	* Number of times not able to reach connector using securebrowse protocol
	*/
	Sbpconnectorfailurerate float64 `json:"sbpconnectorfailurerate,omitempty"`
	Totalcvpnrequest int `json:"totalcvpnrequest,omitempty"`
	/**
	* Number of cvpn requests received
	*/
	Cvpnrequestrate float64 `json:"cvpnrequestrate,omitempty"`
	Totalcvpnsaasrequest int `json:"totalcvpnsaasrequest,omitempty"`
	/**
	* Number of cvpn requests received for saas apps
	*/
	Cvpnsaasrequestrate float64 `json:"cvpnsaasrequestrate,omitempty"`
	Totalcvpnwebrequest int `json:"totalcvpnwebrequest,omitempty"`
	/**
	* Number of cvpn requests received for web apps
	*/
	Cvpnwebrequestrate float64 `json:"cvpnwebrequestrate,omitempty"`
	Totalcvpnrequestsuccess int `json:"totalcvpnrequestsuccess,omitempty"`
	/**
	* Number of cvpn requests succeeded
	*/
	Cvpnrequestsuccessrate float64 `json:"cvpnrequestsuccessrate,omitempty"`
	Totalcvpnsaasrequestsuccess int `json:"totalcvpnsaasrequestsuccess,omitempty"`
	/**
	* Number of cvpn requests succeeded for saas apps
	*/
	Cvpnsaasrequestsuccessrate float64 `json:"cvpnsaasrequestsuccessrate,omitempty"`
	Totalcvpnwebrequestsuccess int `json:"totalcvpnwebrequestsuccess,omitempty"`
	/**
	* Number of cvpn requests succeeded for web apps
	*/
	Cvpnwebrequestsuccessrate float64 `json:"cvpnwebrequestsuccessrate,omitempty"`
	Totalcvpnrequestfailure int `json:"totalcvpnrequestfailure,omitempty"`
	/**
	* Number of cvpn requests failed
	*/
	Cvpnrequestfailurerate float64 `json:"cvpnrequestfailurerate,omitempty"`
	Totalcvpnsaasrequestfailure int `json:"totalcvpnsaasrequestfailure,omitempty"`
	/**
	* Number of cvpn requests failed for saas apps
	*/
	Cvpnsaasrequestfailurerate float64 `json:"cvpnsaasrequestfailurerate,omitempty"`
	Totalcvpnwebrequestfailure int `json:"totalcvpnwebrequestfailure,omitempty"`
	/**
	* Number of cvpn requests failed for web apps
	*/
	Cvpnwebrequestfailurerate float64 `json:"cvpnwebrequestfailurerate,omitempty"`
	Totalcvpndnsrequestfailure int `json:"totalcvpndnsrequestfailure,omitempty"`
	/**
	* Number of DNS failures while using cvpn
	*/
	Cvpndnsrequestfailurerate float64 `json:"cvpndnsrequestfailurerate,omitempty"`
	Totalcvpnbackendconnectionfailure int `json:"totalcvpnbackendconnectionfailure,omitempty"`
	/**
	* Number of backend connection failure while using cvpn
	*/
	Cvpnbackendconnectionfailurerate float64 `json:"cvpnbackendconnectionfailurerate,omitempty"`
	Totalcvpnconnectorfailure int `json:"totalcvpnconnectorfailure,omitempty"`
	/**
	* Number of times not able to reach connector using cvpn
	*/
	Cvpnconnectorfailurerate float64 `json:"cvpnconnectorfailurerate,omitempty"`
	Totaldirectaccessrequest int `json:"totaldirectaccessrequest,omitempty"`
	/**
	* Number of directaccess requests received
	*/
	Directaccessrequestrate float64 `json:"directaccessrequestrate,omitempty"`
	Totaldirectaccessrequestsuccess int `json:"totaldirectaccessrequestsuccess,omitempty"`
	/**
	* Number of directaccess requests succeeded
	*/
	Directaccessrequestsuccessrate float64 `json:"directaccessrequestsuccessrate,omitempty"`
	Totaldirectaccessrequestfailure int `json:"totaldirectaccessrequestfailure,omitempty"`
	/**
	* Number of directaccess requests failed
	*/
	Directaccessrequestfailurerate float64 `json:"directaccessrequestfailurerate,omitempty"`
	Totaldirectaccessconnectorfailure int `json:"totaldirectaccessconnectorfailure,omitempty"`
	/**
	* Number of times not able to reach connector using Directaccess
	*/
	Directaccessconnectorfailurerate float64 `json:"directaccessconnectorfailurerate,omitempty"`

}
