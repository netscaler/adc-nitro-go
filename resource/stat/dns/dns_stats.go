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

package dns


type Dnsstats struct {
	/**
	* Clear the statsistics / counters
	*/
	Clearstats string `json:"clearstats,omitempty"`
	Dnstotqueries int `json:"dnstotqueries,omitempty"`
	/**
	* Total number of DNS queries received.
	*/
	Dnsqueriesrate float64 `json:"dnsqueriesrate,omitempty"`
	Dnstotmultiquery int `json:"dnstotmultiquery,omitempty"`
	Dnstotanswers int `json:"dnstotanswers,omitempty"`
	/**
	* Total number of DNS responses received.
	*/
	Dnsanswersrate float64 `json:"dnsanswersrate,omitempty"`
	Dnstotserverresponse int `json:"dnstotserverresponse,omitempty"`
	/**
	* Total number of Server responses received.
	*/
	Dnsserverresponserate float64 `json:"dnsserverresponserate,omitempty"`
	Dnstotrecupdate int `json:"dnstotrecupdate,omitempty"`
	Dnscurcachesize int `json:"dnscurcachesize,omitempty"`
	Dnscurnegcachesize int `json:"dnscurnegcachesize,omitempty"`
	Dnstotjumboqueries int `json:"dnstotjumboqueries,omitempty"`
	/**
	* Total number of Jumbo DNS queries received over UDP.
	*/
	Dnsjumboqueriesrate float64 `json:"dnsjumboqueriesrate,omitempty"`
	Dnstotjumboanswers int `json:"dnstotjumboanswers,omitempty"`
	/**
	* Total number of Jumbo DNS responses sent over UDP.
	*/
	Dnsjumboanswersrate float64 `json:"dnsjumboanswersrate,omitempty"`
	Dnstotjumboserverresponses int `json:"dnstotjumboserverresponses,omitempty"`
	/**
	* Total number of Jumbo DNS responses received over UDP.
	*/
	Dnsjumboserverresponsesrate float64 `json:"dnsjumboserverresponsesrate,omitempty"`
	Dnstotauthans int `json:"dnstotauthans,omitempty"`
	Dnstotserverquery int `json:"dnstotserverquery,omitempty"`
	/**
	* Total number of Server queries sent.
	*/
	Dnsserverqueryrate float64 `json:"dnsserverqueryrate,omitempty"`
	Dnstotcacheflush int `json:"dnstotcacheflush,omitempty"`
	Dnstotcacheentriesflush int `json:"dnstotcacheentriesflush,omitempty"`
	Dnscurnoauthentries int `json:"dnscurnoauthentries,omitempty"`
	Dnscurauthentries int `json:"dnscurauthentries,omitempty"`
	Dnstotauthnonames int `json:"dnstotauthnonames,omitempty"`
	Dnstotunsupportedresponseclass int `json:"dnstotunsupportedresponseclass,omitempty"`
	Dnstotinvalidqueryformat int `json:"dnstotinvalidqueryformat,omitempty"`
	Dnstotstrayanswer int `json:"dnstotstrayanswer,omitempty"`
	Dnstotresponsebadlen int `json:"dnstotresponsebadlen,omitempty"`
	Dnstotreqrefusals int `json:"dnstotreqrefusals,omitempty"`
	Dnserrnullattack int `json:"dnserrnullattack,omitempty"`
	Dnstotunsupportedresponsetype int `json:"dnstotunsupportedresponsetype,omitempty"`
	Dnstotunsupportedqueryclass int `json:"dnstotunsupportedqueryclass,omitempty"`
	Dnstotnonauthnodatas int `json:"dnstotnonauthnodatas,omitempty"`
	Dnstotnodataresps int `json:"dnstotnodataresps,omitempty"`
	Dnstotmultiquerydisableerror int `json:"dnstotmultiquerydisableerror,omitempty"`
	Dnstotothererrors int `json:"dnstotothererrors,omitempty"`
	Dns64totqueries int `json:"dns64totqueries,omitempty"`
	/**
	* Total number of DNS64 queries recieved.
	*/
	Dns64queriesrate float64 `json:"dns64queriesrate,omitempty"`
	Dns64totanswers int `json:"dns64totanswers,omitempty"`
	/**
	* Total number of DNS64 answers served.
	*/
	Dns64answersrate float64 `json:"dns64answersrate,omitempty"`
	Dns64totrwanswers int `json:"dns64totrwanswers,omitempty"`
	/**
	* Total number of DNS64 answers served after rewriting the response.
	*/
	Dns64rwanswersrate float64 `json:"dns64rwanswersrate,omitempty"`
	Dns64totresponses int `json:"dns64totresponses,omitempty"`
	/**
	* Total number of responses recieved from backend in DNS64 context.
	*/
	Dns64responsesrate float64 `json:"dns64responsesrate,omitempty"`
	Dns64totgslbqueries int `json:"dns64totgslbqueries,omitempty"`
	/**
	* Total number of DNS64 queries for GSLB domain
	*/
	Dns64gslbqueriesrate float64 `json:"dns64gslbqueriesrate,omitempty"`
	Dns64totgslbanswers int `json:"dns64totgslbanswers,omitempty"`
	/**
	* Total number of DNS64 queries served.
	*/
	Dns64gslbanswersrate float64 `json:"dns64gslbanswersrate,omitempty"`
	Dns64tottcanswers int `json:"dns64tottcanswers,omitempty"`
	/**
	* Total number of Answers served with TC bit set in DNS64 context.
	*/
	Dns64tcanswersrate float64 `json:"dns64tcanswersrate,omitempty"`
	Dns64totsvraqueries int `json:"dns64totsvraqueries,omitempty"`
	/**
	* Total number of Queries sent by DNS64 module to backend.
	*/
	Dns64svraqueriesrate float64 `json:"dns64svraqueriesrate,omitempty"`
	Dns64totaaaabypass int `json:"dns64totaaaabypass,omitempty"`
	/**
	* Total number of times AAAA query has been bypassed in DNS64 trnsaction.
	*/
	Dns64aaaabypassrate float64 `json:"dns64aaaabypassrate,omitempty"`
	Dns64tottcpqueries int `json:"dns64tottcpqueries,omitempty"`
	/**
	* Total number of dns64 queries over TCP
	*/
	Dns64tcpqueriesrate float64 `json:"dns64tcpqueriesrate,omitempty"`
	Dns64activepolicies int `json:"dns64activepolicies,omitempty"`
	Dns64totnodataresp int `json:"dns64totnodataresp,omitempty"`
	/**
	* Total number of responses recieved from backend with ancount 0
	*/
	Dns64nodataresprate float64 `json:"dns64nodataresprate,omitempty"`
	Dnstotnsrecqueries int `json:"dnstotnsrecqueries,omitempty"`
	/**
	* Total number of NS queries received.
	*/
	Dnsnsrecqueriesrate float64 `json:"dnsnsrecqueriesrate,omitempty"`
	Dnstotsoarecqueries int `json:"dnstotsoarecqueries,omitempty"`
	/**
	* Total number of SOA queries received.
	*/
	Dnssoarecqueriesrate float64 `json:"dnssoarecqueriesrate,omitempty"`
	Dnstotptrrecqueries int `json:"dnstotptrrecqueries,omitempty"`
	/**
	* Total number of PTR queries received.
	*/
	Dnsptrrecqueriesrate float64 `json:"dnsptrrecqueriesrate,omitempty"`
	Dnstotsrvrecqueries int `json:"dnstotsrvrecqueries,omitempty"`
	/**
	* Total number of SRV queries received.
	*/
	Dnssrvrecqueriesrate float64 `json:"dnssrvrecqueriesrate,omitempty"`
	Dnstotaresponse int `json:"dnstotaresponse,omitempty"`
	/**
	* Total number of A responses received.
	*/
	Dnsaresponserate float64 `json:"dnsaresponserate,omitempty"`
	Dnstotcnameresponse int `json:"dnstotcnameresponse,omitempty"`
	/**
	* Total number of CNAME responses received.
	*/
	Dnscnameresponserate float64 `json:"dnscnameresponserate,omitempty"`
	Dnstotmxresponse int `json:"dnstotmxresponse,omitempty"`
	/**
	* Total number of MX responses received.
	*/
	Dnsmxresponserate float64 `json:"dnsmxresponserate,omitempty"`
	Dnstotanyresponse int `json:"dnstotanyresponse,omitempty"`
	/**
	* Total number of ANY responses received.
	*/
	Dnsanyresponserate float64 `json:"dnsanyresponserate,omitempty"`
	Dnstotnsrecupdate int `json:"dnstotnsrecupdate,omitempty"`
	Dnstotsoarecupdate int `json:"dnstotsoarecupdate,omitempty"`
	Dnstotptrrecupdate int `json:"dnstotptrrecupdate,omitempty"`
	Dnstotsrvrecupdate int `json:"dnstotsrvrecupdate,omitempty"`
	Dnstotaaaarecqueries int `json:"dnstotaaaarecqueries,omitempty"`
	/**
	* Total number of AAAA queries received.
	*/
	Dnsaaaarecqueriesrate float64 `json:"dnsaaaarecqueriesrate,omitempty"`
	Dnstotarecqueries int `json:"dnstotarecqueries,omitempty"`
	/**
	* Total number of A queries received.
	*/
	Dnsarecqueriesrate float64 `json:"dnsarecqueriesrate,omitempty"`
	Dnstotcnamerecqueries int `json:"dnstotcnamerecqueries,omitempty"`
	/**
	* Total number of CNAME queries received.
	*/
	Dnscnamerecqueriesrate float64 `json:"dnscnamerecqueriesrate,omitempty"`
	Dnstotmxrecqueries int `json:"dnstotmxrecqueries,omitempty"`
	/**
	* Total number of MX queries received.
	*/
	Dnsmxrecqueriesrate float64 `json:"dnsmxrecqueriesrate,omitempty"`
	Dnstotanyqueries int `json:"dnstotanyqueries,omitempty"`
	/**
	* Total number of ANY queries received.
	*/
	Dnsanyqueriesrate float64 `json:"dnsanyqueriesrate,omitempty"`
	Dnstotaaaaresponse int `json:"dnstotaaaaresponse,omitempty"`
	/**
	* Total number of AAAA responses received.
	*/
	Dnsaaaaresponserate float64 `json:"dnsaaaaresponserate,omitempty"`
	Dnstotnsresponse int `json:"dnstotnsresponse,omitempty"`
	/**
	* Total number of NS responses received.
	*/
	Dnsnsresponserate float64 `json:"dnsnsresponserate,omitempty"`
	Dnstotsoaresponse int `json:"dnstotsoaresponse,omitempty"`
	/**
	* Total number of SOA responses received.
	*/
	Dnssoaresponserate float64 `json:"dnssoaresponserate,omitempty"`
	Dnstotptrresponse int `json:"dnstotptrresponse,omitempty"`
	/**
	* Total number of PTR responses received.
	*/
	Dnsptrresponserate float64 `json:"dnsptrresponserate,omitempty"`
	Dnstotsrvresponse int `json:"dnstotsrvresponse,omitempty"`
	/**
	* Total number of SRV responses received.
	*/
	Dnssrvresponserate float64 `json:"dnssrvresponserate,omitempty"`
	Dnstotaaaarecupdate int `json:"dnstotaaaarecupdate,omitempty"`
	Dnstotarecupdate int `json:"dnstotarecupdate,omitempty"`
	Dnstotmxrecupdate int `json:"dnstotmxrecupdate,omitempty"`
	Dnstotcnamerecupdate int `json:"dnstotcnamerecupdate,omitempty"`
	Dnscuraaaarecord int `json:"dnscuraaaarecord,omitempty"`
	Dnscurarecord int `json:"dnscurarecord,omitempty"`
	Dnscurmxrecord int `json:"dnscurmxrecord,omitempty"`
	Dnscurcnamerecord int `json:"dnscurcnamerecord,omitempty"`
	Dnscurnsrecord int `json:"dnscurnsrecord,omitempty"`
	Dnscursoarecord int `json:"dnscursoarecord,omitempty"`
	Dnscurptrrecord int `json:"dnscurptrrecord,omitempty"`
	Dnscursrvrecord int `json:"dnscursrvrecord,omitempty"`
	Dnstotaaaarecfailed int `json:"dnstotaaaarecfailed,omitempty"`
	Dnstotarecfailed int `json:"dnstotarecfailed,omitempty"`
	Dnstotmxrecfailed int `json:"dnstotmxrecfailed,omitempty"`
	Dnstotptrrecfailed int `json:"dnstotptrrecfailed,omitempty"`
	Dnstotnsrecfailed int `json:"dnstotnsrecfailed,omitempty"`
	Dnstotcnamerecfailed int `json:"dnstotcnamerecfailed,omitempty"`
	Dnstotsoarecfailed int `json:"dnstotsoarecfailed,omitempty"`
	Dnstotsrvrecfailed int `json:"dnstotsrvrecfailed,omitempty"`
	Dnstotanyrecfailed int `json:"dnstotanyrecfailed,omitempty"`
	Dnstotunsupportedqueries int `json:"dnstotunsupportedqueries,omitempty"`

}
