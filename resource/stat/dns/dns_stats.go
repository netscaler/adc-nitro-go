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
	Dnstotqueries uint64 `json:"dnstotqueries,omitempty"`
	/**
	* Total number of DNS queries received.
	*/
	Dnsqueriesrate float64 `json:"dnsqueriesrate,omitempty"`
	Dnstotmultiquery uint64 `json:"dnstotmultiquery,omitempty"`
	Dnstotanswers uint64 `json:"dnstotanswers,omitempty"`
	/**
	* Total number of DNS responses received.
	*/
	Dnsanswersrate float64 `json:"dnsanswersrate,omitempty"`
	Dnstotserverresponse uint64 `json:"dnstotserverresponse,omitempty"`
	/**
	* Total number of Server responses received.
	*/
	Dnsserverresponserate float64 `json:"dnsserverresponserate,omitempty"`
	Dnstotrecupdate uint64 `json:"dnstotrecupdate,omitempty"`
	Dnscurcachesize uint64 `json:"dnscurcachesize,omitempty"`
	Dnscurnegcachesize uint64 `json:"dnscurnegcachesize,omitempty"`
	Dnstotjumboqueries uint64 `json:"dnstotjumboqueries,omitempty"`
	/**
	* Total number of Jumbo DNS queries received over UDP.
	*/
	Dnsjumboqueriesrate float64 `json:"dnsjumboqueriesrate,omitempty"`
	Dnstotjumboanswers uint64 `json:"dnstotjumboanswers,omitempty"`
	/**
	* Total number of Jumbo DNS responses sent over UDP.
	*/
	Dnsjumboanswersrate float64 `json:"dnsjumboanswersrate,omitempty"`
	Dnstotjumboserverresponses uint64 `json:"dnstotjumboserverresponses,omitempty"`
	/**
	* Total number of Jumbo DNS responses received over UDP.
	*/
	Dnsjumboserverresponsesrate float64 `json:"dnsjumboserverresponsesrate,omitempty"`
	Dnstotauthans uint64 `json:"dnstotauthans,omitempty"`
	Dnstotserverquery uint64 `json:"dnstotserverquery,omitempty"`
	/**
	* Total number of Server queries sent.
	*/
	Dnsserverqueryrate float64 `json:"dnsserverqueryrate,omitempty"`
	Dnstotcacheflush uint64 `json:"dnstotcacheflush,omitempty"`
	Dnstotcacheentriesflush uint64 `json:"dnstotcacheentriesflush,omitempty"`
	Dnscurnoauthentries uint32 `json:"dnscurnoauthentries,omitempty"`
	Dnscurauthentries uint32 `json:"dnscurauthentries,omitempty"`
	Dnstotauthnonames uint64 `json:"dnstotauthnonames,omitempty"`
	Dnstotunsupportedresponseclass uint64 `json:"dnstotunsupportedresponseclass,omitempty"`
	Dnstotinvalidqueryformat uint64 `json:"dnstotinvalidqueryformat,omitempty"`
	Dnstotstrayanswer uint64 `json:"dnstotstrayanswer,omitempty"`
	Dnstotresponsebadlen uint64 `json:"dnstotresponsebadlen,omitempty"`
	Dnstotreqrefusals uint64 `json:"dnstotreqrefusals,omitempty"`
	Dnserrnullattack uint64 `json:"dnserrnullattack,omitempty"`
	Dnstotunsupportedresponsetype uint64 `json:"dnstotunsupportedresponsetype,omitempty"`
	Dnstotunsupportedqueryclass uint64 `json:"dnstotunsupportedqueryclass,omitempty"`
	Dnstotnonauthnodatas uint64 `json:"dnstotnonauthnodatas,omitempty"`
	Dnstotnodataresps uint64 `json:"dnstotnodataresps,omitempty"`
	Dnstotmultiquerydisableerror uint64 `json:"dnstotmultiquerydisableerror,omitempty"`
	Dnstotothererrors uint64 `json:"dnstotothererrors,omitempty"`
	Dns64totqueries uint32 `json:"dns64totqueries,omitempty"`
	/**
	* Total number of DNS64 queries recieved.
	*/
	Dns64queriesrate float64 `json:"dns64queriesrate,omitempty"`
	Dns64totanswers uint32 `json:"dns64totanswers,omitempty"`
	/**
	* Total number of DNS64 answers served.
	*/
	Dns64answersrate float64 `json:"dns64answersrate,omitempty"`
	Dns64totrwanswers uint32 `json:"dns64totrwanswers,omitempty"`
	/**
	* Total number of DNS64 answers served after rewriting the response.
	*/
	Dns64rwanswersrate float64 `json:"dns64rwanswersrate,omitempty"`
	Dns64totresponses uint32 `json:"dns64totresponses,omitempty"`
	/**
	* Total number of responses recieved from backend in DNS64 context.
	*/
	Dns64responsesrate float64 `json:"dns64responsesrate,omitempty"`
	Dns64totgslbqueries uint32 `json:"dns64totgslbqueries,omitempty"`
	/**
	* Total number of DNS64 queries for GSLB domain
	*/
	Dns64gslbqueriesrate float64 `json:"dns64gslbqueriesrate,omitempty"`
	Dns64totgslbanswers uint32 `json:"dns64totgslbanswers,omitempty"`
	/**
	* Total number of DNS64 queries served.
	*/
	Dns64gslbanswersrate float64 `json:"dns64gslbanswersrate,omitempty"`
	Dns64tottcanswers uint32 `json:"dns64tottcanswers,omitempty"`
	/**
	* Total number of Answers served with TC bit set in DNS64 context.
	*/
	Dns64tcanswersrate float64 `json:"dns64tcanswersrate,omitempty"`
	Dns64totsvraqueries uint32 `json:"dns64totsvraqueries,omitempty"`
	/**
	* Total number of Queries sent by DNS64 module to backend.
	*/
	Dns64svraqueriesrate float64 `json:"dns64svraqueriesrate,omitempty"`
	Dns64totaaaabypass uint32 `json:"dns64totaaaabypass,omitempty"`
	/**
	* Total number of times AAAA query has been bypassed in DNS64 trnsaction.
	*/
	Dns64aaaabypassrate float64 `json:"dns64aaaabypassrate,omitempty"`
	Dns64tottcpqueries uint32 `json:"dns64tottcpqueries,omitempty"`
	/**
	* Total number of dns64 queries over TCP
	*/
	Dns64tcpqueriesrate float64 `json:"dns64tcpqueriesrate,omitempty"`
	Dns64activepolicies uint32 `json:"dns64activepolicies,omitempty"`
	Dns64totnodataresp uint32 `json:"dns64totnodataresp,omitempty"`
	/**
	* Total number of responses recieved from backend with ancount 0
	*/
	Dns64nodataresprate float64 `json:"dns64nodataresprate,omitempty"`
	Dnstotnsrecqueries uint64 `json:"dnstotnsrecqueries,omitempty"`
	/**
	* Total number of NS queries received.
	*/
	Dnsnsrecqueriesrate float64 `json:"dnsnsrecqueriesrate,omitempty"`
	Dnstotsoarecqueries uint64 `json:"dnstotsoarecqueries,omitempty"`
	/**
	* Total number of SOA queries received.
	*/
	Dnssoarecqueriesrate float64 `json:"dnssoarecqueriesrate,omitempty"`
	Dnstotptrrecqueries uint64 `json:"dnstotptrrecqueries,omitempty"`
	/**
	* Total number of PTR queries received.
	*/
	Dnsptrrecqueriesrate float64 `json:"dnsptrrecqueriesrate,omitempty"`
	Dnstotsrvrecqueries uint64 `json:"dnstotsrvrecqueries,omitempty"`
	/**
	* Total number of SRV queries received.
	*/
	Dnssrvrecqueriesrate float64 `json:"dnssrvrecqueriesrate,omitempty"`
	Dnstotaresponse uint64 `json:"dnstotaresponse,omitempty"`
	/**
	* Total number of A responses received.
	*/
	Dnsaresponserate float64 `json:"dnsaresponserate,omitempty"`
	Dnstotcnameresponse uint64 `json:"dnstotcnameresponse,omitempty"`
	/**
	* Total number of CNAME responses received.
	*/
	Dnscnameresponserate float64 `json:"dnscnameresponserate,omitempty"`
	Dnstotmxresponse uint64 `json:"dnstotmxresponse,omitempty"`
	/**
	* Total number of MX responses received.
	*/
	Dnsmxresponserate float64 `json:"dnsmxresponserate,omitempty"`
	Dnstotanyresponse uint64 `json:"dnstotanyresponse,omitempty"`
	/**
	* Total number of ANY responses received.
	*/
	Dnsanyresponserate float64 `json:"dnsanyresponserate,omitempty"`
	Dnstotnsrecupdate uint64 `json:"dnstotnsrecupdate,omitempty"`
	Dnstotsoarecupdate uint64 `json:"dnstotsoarecupdate,omitempty"`
	Dnstotptrrecupdate uint64 `json:"dnstotptrrecupdate,omitempty"`
	Dnstotsrvrecupdate uint64 `json:"dnstotsrvrecupdate,omitempty"`
	Dnstotaaaarecqueries uint64 `json:"dnstotaaaarecqueries,omitempty"`
	/**
	* Total number of AAAA queries received.
	*/
	Dnsaaaarecqueriesrate float64 `json:"dnsaaaarecqueriesrate,omitempty"`
	Dnstotarecqueries uint64 `json:"dnstotarecqueries,omitempty"`
	/**
	* Total number of A queries received.
	*/
	Dnsarecqueriesrate float64 `json:"dnsarecqueriesrate,omitempty"`
	Dnstotcnamerecqueries uint64 `json:"dnstotcnamerecqueries,omitempty"`
	/**
	* Total number of CNAME queries received.
	*/
	Dnscnamerecqueriesrate float64 `json:"dnscnamerecqueriesrate,omitempty"`
	Dnstotmxrecqueries uint64 `json:"dnstotmxrecqueries,omitempty"`
	/**
	* Total number of MX queries received.
	*/
	Dnsmxrecqueriesrate float64 `json:"dnsmxrecqueriesrate,omitempty"`
	Dnstotanyqueries uint64 `json:"dnstotanyqueries,omitempty"`
	/**
	* Total number of ANY queries received.
	*/
	Dnsanyqueriesrate float64 `json:"dnsanyqueriesrate,omitempty"`
	Dnstotaaaaresponse uint64 `json:"dnstotaaaaresponse,omitempty"`
	/**
	* Total number of AAAA responses received.
	*/
	Dnsaaaaresponserate float64 `json:"dnsaaaaresponserate,omitempty"`
	Dnstotnsresponse uint64 `json:"dnstotnsresponse,omitempty"`
	/**
	* Total number of NS responses received.
	*/
	Dnsnsresponserate float64 `json:"dnsnsresponserate,omitempty"`
	Dnstotsoaresponse uint64 `json:"dnstotsoaresponse,omitempty"`
	/**
	* Total number of SOA responses received.
	*/
	Dnssoaresponserate float64 `json:"dnssoaresponserate,omitempty"`
	Dnstotptrresponse uint64 `json:"dnstotptrresponse,omitempty"`
	/**
	* Total number of PTR responses received.
	*/
	Dnsptrresponserate float64 `json:"dnsptrresponserate,omitempty"`
	Dnstotsrvresponse uint64 `json:"dnstotsrvresponse,omitempty"`
	/**
	* Total number of SRV responses received.
	*/
	Dnssrvresponserate float64 `json:"dnssrvresponserate,omitempty"`
	Dnstotaaaarecupdate uint64 `json:"dnstotaaaarecupdate,omitempty"`
	Dnstotarecupdate uint64 `json:"dnstotarecupdate,omitempty"`
	Dnstotmxrecupdate uint64 `json:"dnstotmxrecupdate,omitempty"`
	Dnstotcnamerecupdate uint64 `json:"dnstotcnamerecupdate,omitempty"`
	Dnscuraaaarecord uint32 `json:"dnscuraaaarecord,omitempty"`
	Dnscurarecord uint32 `json:"dnscurarecord,omitempty"`
	Dnscurmxrecord uint32 `json:"dnscurmxrecord,omitempty"`
	Dnscurcnamerecord uint32 `json:"dnscurcnamerecord,omitempty"`
	Dnscurnsrecord uint32 `json:"dnscurnsrecord,omitempty"`
	Dnscursoarecord uint32 `json:"dnscursoarecord,omitempty"`
	Dnscurptrrecord uint32 `json:"dnscurptrrecord,omitempty"`
	Dnscursrvrecord uint32 `json:"dnscursrvrecord,omitempty"`
	Dnstotaaaarecfailed uint64 `json:"dnstotaaaarecfailed,omitempty"`
	Dnstotarecfailed uint64 `json:"dnstotarecfailed,omitempty"`
	Dnstotmxrecfailed uint64 `json:"dnstotmxrecfailed,omitempty"`
	Dnstotptrrecfailed uint64 `json:"dnstotptrrecfailed,omitempty"`
	Dnstotnsrecfailed uint64 `json:"dnstotnsrecfailed,omitempty"`
	Dnstotcnamerecfailed uint64 `json:"dnstotcnamerecfailed,omitempty"`
	Dnstotsoarecfailed uint64 `json:"dnstotsoarecfailed,omitempty"`
	Dnstotsrvrecfailed uint64 `json:"dnstotsrvrecfailed,omitempty"`
	Dnstotanyrecfailed uint64 `json:"dnstotanyrecfailed,omitempty"`
	Dnstotunsupportedqueries uint64 `json:"dnstotunsupportedqueries,omitempty"`

}
