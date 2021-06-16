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

package ns


type Nsstats struct {
	/**
	* Clear the statsistics / counters
	*/
	Clearstats string `json:"clearstats,omitempty"`
	Rescpuusagepcnt float64 `json:"rescpuusagepcnt,omitempty"`
	Cpuusagepcnt float64 `json:"cpuusagepcnt,omitempty"`
	Cachemaxmemorykb uint32 `json:"cachemaxmemorykb,omitempty"`
	Delcmpratio float64 `json:"delcmpratio,omitempty"`
	Rescpuusage uint32 `json:"rescpuusage,omitempty"`
	Cpuusage uint32 `json:"cpuusage,omitempty"`
	Resmemusage uint32 `json:"resmemusage,omitempty"`
	Comptotaldatacompressionratio float64 `json:"comptotaldatacompressionratio,omitempty"`
	Compratio float64 `json:"compratio,omitempty"`
	Cacheutilizedmemorykb uint32 `json:"cacheutilizedmemorykb,omitempty"`
	Cachemaxmemoryactivekb uint64 `json:"cachemaxmemoryactivekb,omitempty"`
	Cache64maxmemorykb uint64 `json:"cache64maxmemorykb,omitempty"`
	Cachepercentoriginbandwidthsaved uint32 `json:"cachepercentoriginbandwidthsaved,omitempty"`
	Cachetotmisses uint64 `json:"cachetotmisses,omitempty"`
	/**
	* Intercepted HTTP requests requiring fetches from origin server.
	*/
	Cachemissesrate float64 `json:"cachemissesrate,omitempty"`
	Cachetothits uint64 `json:"cachetothits,omitempty"`
	/**
	* Responses served from the integrated cache. These responses match a policy with a CACHE action.
	*/
	Cachehitsrate float64 `json:"cachehitsrate,omitempty"`
	Memusagepcnt float64 `json:"memusagepcnt,omitempty"`
	Memuseinmb uint32 `json:"memuseinmb,omitempty"`
	Mgmtcpuusagepcnt float64 `json:"mgmtcpuusagepcnt,omitempty"`
	Pktcpuusagepcnt float64 `json:"pktcpuusagepcnt,omitempty"`
	Starttimelocal string `json:"starttimelocal,omitempty"`
	Starttime string `json:"starttime,omitempty"`
	Transtime string `json:"transtime,omitempty"`
	Hacurstate string `json:"hacurstate,omitempty"`
	Hacurmasterstate string `json:"hacurmasterstate,omitempty"`
	Sslnumcardsup uint32 `json:"sslnumcardsup,omitempty"`
	Sslcards uint32 `json:"sslcards,omitempty"`
	Disk0perusage int32 `json:"disk0perusage,omitempty"`
	Disk1perusage int32 `json:"disk1perusage,omitempty"`
	Disk0avail int32 `json:"disk0avail,omitempty"`
	Disk1avail int32 `json:"disk1avail,omitempty"`
	Totrxmbits uint64 `json:"totrxmbits,omitempty"`
	/**
	* Number of megabytes received by the Citrix ADC.
	*/
	Rxmbitsrate float64 `json:"rxmbitsrate,omitempty"`
	Tottxmbits uint64 `json:"tottxmbits,omitempty"`
	/**
	* Number of megabytes transmitted by the Citrix ADC.
	*/
	Txmbitsrate float64 `json:"txmbitsrate,omitempty"`
	Tcpcurclientconn uint32 `json:"tcpcurclientconn,omitempty"`
	Tcpcurclientconnestablished uint32 `json:"tcpcurclientconnestablished,omitempty"`
	Tcpcurserverconn uint32 `json:"tcpcurserverconn,omitempty"`
	Tcpcurserverconnestablished uint32 `json:"tcpcurserverconnestablished,omitempty"`
	Httptotrequests uint64 `json:"httptotrequests,omitempty"`
	/**
	* Total number of HTTP requests received.
	*/
	Httprequestsrate float64 `json:"httprequestsrate,omitempty"`
	Httptotresponses uint64 `json:"httptotresponses,omitempty"`
	/**
	* Total number of HTTP responses sent.
	*/
	Httpresponsesrate float64 `json:"httpresponsesrate,omitempty"`
	Httptotrxrequestbytes uint64 `json:"httptotrxrequestbytes,omitempty"`
	/**
	* Total number of bytes of HTTP request data received.
	*/
	Httprxrequestbytesrate float64 `json:"httprxrequestbytesrate,omitempty"`
	Httptotrxresponsebytes uint64 `json:"httptotrxresponsebytes,omitempty"`
	/**
	* Total number of bytes of HTTP response data received.
	*/
	Httprxresponsebytesrate float64 `json:"httprxresponsebytesrate,omitempty"`
	Ssltottransactions uint64 `json:"ssltottransactions,omitempty"`
	/**
	* Number of SSL transactions on the Citrix ADC
	*/
	Ssltransactionsrate float64 `json:"ssltransactionsrate,omitempty"`
	Ssltotsessionhits uint64 `json:"ssltotsessionhits,omitempty"`
	/**
	* Number of SSL session reuse hits on the Citrix ADC.
	*/
	Sslsessionhitsrate float64 `json:"sslsessionhitsrate,omitempty"`
	Appfirewallrequests uint64 `json:"appfirewallrequests,omitempty"`
	/**
	* HTTP/HTTPS requests sent to your protected web servers via the Application Firewall.
	*/
	Appfirewallrequestsrate float64 `json:"appfirewallrequestsrate,omitempty"`
	Appfirewallresponses uint64 `json:"appfirewallresponses,omitempty"`
	/**
	* HTTP/HTTPS responses sent by your protected web servers via the Application Firewall.
	*/
	Appfirewallresponsesrate float64 `json:"appfirewallresponsesrate,omitempty"`
	Appfirewallaborts uint64 `json:"appfirewallaborts,omitempty"`
	/**
	* Incomplete HTTP/HTTPS requests aborted by the client before the Application Firewall could finish processing them.
	*/
	Appfirewallabortsrate float64 `json:"appfirewallabortsrate,omitempty"`
	Appfirewallredirects uint64 `json:"appfirewallredirects,omitempty"`
	/**
	* HTTP/HTTPS requests redirected by the Application Firewall to a different Web page or web server. (HTTP 302)
	*/
	Appfirewallredirectsrate float64 `json:"appfirewallredirectsrate,omitempty"`
	Misccounter0 int32 `json:"misccounter0,omitempty"`
	Misccounter1 int32 `json:"misccounter1,omitempty"`
	Numcpus uint32 `json:"numcpus,omitempty"`

}
