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

package stat

type Nsstats struct {
	/**
	* Clear the statsistics / counters
	 */
	Clearstats                       string  `json:"clearstats,omitempty"`
	Rescpuusagepcnt                  float64 `json:"rescpuusagepcnt,omitempty"`
	Cpuusagepcnt                     float64 `json:"cpuusagepcnt,omitempty"`
	Cachemaxmemorykb                 int     `json:"cachemaxmemorykb,omitempty"`
	Delcmpratio                      float64 `json:"delcmpratio,omitempty"`
	Rescpuusage                      int     `json:"rescpuusage,omitempty"`
	Cpuusage                         int     `json:"cpuusage,omitempty"`
	Resmemusage                      int     `json:"resmemusage,omitempty"`
	Comptotaldatacompressionratio    float64 `json:"comptotaldatacompressionratio,omitempty"`
	Compratio                        float64 `json:"compratio,omitempty"`
	Cacheutilizedmemorykb            int     `json:"cacheutilizedmemorykb,omitempty"`
	Cachemaxmemoryactivekb           int     `json:"cachemaxmemoryactivekb,omitempty"`
	Cache64maxmemorykb               int     `json:"cache64maxmemorykb,omitempty"`
	Cachepercentoriginbandwidthsaved int     `json:"cachepercentoriginbandwidthsaved,omitempty"`
	Cachetotmisses                   int     `json:"cachetotmisses,omitempty"`
	/**
	* Intercepted HTTP requests requiring fetches from origin server.
	 */
	Cachemissesrate float64 `json:"cachemissesrate,omitempty"`
	Cachetothits    int     `json:"cachetothits,omitempty"`
	/**
	* Responses served from the integrated cache. These responses match a policy with a CACHE action.
	 */
	Cachehitsrate    float64 `json:"cachehitsrate,omitempty"`
	Memusagepcnt     float64 `json:"memusagepcnt,omitempty"`
	Memuseinmb       int     `json:"memuseinmb,omitempty"`
	Mgmtcpuusagepcnt float64 `json:"mgmtcpuusagepcnt,omitempty"`
	Pktcpuusagepcnt  float64 `json:"pktcpuusagepcnt,omitempty"`
	Starttimelocal   string  `json:"starttimelocal,omitempty"`
	Starttime        string  `json:"starttime,omitempty"`
	Transtime        string  `json:"transtime,omitempty"`
	Hacurstate       string  `json:"hacurstate,omitempty"`
	Hacurmasterstate string  `json:"hacurmasterstate,omitempty"`
	Sslnumcardsup    int     `json:"sslnumcardsup,omitempty"`
	Sslcards         int     `json:"sslcards,omitempty"`
	Disk0perusage    int     `json:"disk0perusage,omitempty"`
	Disk1perusage    int     `json:"disk1perusage,omitempty"`
	Disk0avail       int     `json:"disk0avail,omitempty"`
	Disk1avail       int     `json:"disk1avail,omitempty"`
	Totrxmbits       int     `json:"totrxmbits,omitempty"`
	/**
	* Number of megabytes received by the Citrix ADC.
	 */
	Rxmbitsrate float64 `json:"rxmbitsrate,omitempty"`
	Tottxmbits  int     `json:"tottxmbits,omitempty"`
	/**
	* Number of megabytes transmitted by the Citrix ADC.
	 */
	Txmbitsrate                 float64 `json:"txmbitsrate,omitempty"`
	Tcpcurclientconn            int     `json:"tcpcurclientconn,omitempty"`
	Tcpcurclientconnestablished int     `json:"tcpcurclientconnestablished,omitempty"`
	Tcpcurserverconn            int     `json:"tcpcurserverconn,omitempty"`
	Tcpcurserverconnestablished int     `json:"tcpcurserverconnestablished,omitempty"`
	Httptotrequests             int     `json:"httptotrequests,omitempty"`
	/**
	* Total number of HTTP requests received.
	 */
	Httprequestsrate float64 `json:"httprequestsrate,omitempty"`
	Httptotresponses int     `json:"httptotresponses,omitempty"`
	/**
	* Total number of HTTP responses sent.
	 */
	Httpresponsesrate     float64 `json:"httpresponsesrate,omitempty"`
	Httptotrxrequestbytes int     `json:"httptotrxrequestbytes,omitempty"`
	/**
	* Total number of bytes of HTTP request data received.
	 */
	Httprxrequestbytesrate float64 `json:"httprxrequestbytesrate,omitempty"`
	Httptotrxresponsebytes int     `json:"httptotrxresponsebytes,omitempty"`
	/**
	* Total number of bytes of HTTP response data received.
	 */
	Httprxresponsebytesrate float64 `json:"httprxresponsebytesrate,omitempty"`
	Ssltottransactions      int     `json:"ssltottransactions,omitempty"`
	/**
	* Number of SSL transactions on the Citrix ADC
	 */
	Ssltransactionsrate float64 `json:"ssltransactionsrate,omitempty"`
	Ssltotsessionhits   int     `json:"ssltotsessionhits,omitempty"`
	/**
	* Number of SSL session reuse hits on the Citrix ADC.
	 */
	Sslsessionhitsrate  float64 `json:"sslsessionhitsrate,omitempty"`
	Appfirewallrequests int     `json:"appfirewallrequests,omitempty"`
	/**
	* HTTP/HTTPS requests sent to your protected web servers via the Application Firewall.
	 */
	Appfirewallrequestsrate float64 `json:"appfirewallrequestsrate,omitempty"`
	Appfirewallresponses    int     `json:"appfirewallresponses,omitempty"`
	/**
	* HTTP/HTTPS responses sent by your protected web servers via the Application Firewall.
	 */
	Appfirewallresponsesrate float64 `json:"appfirewallresponsesrate,omitempty"`
	Appfirewallaborts        int     `json:"appfirewallaborts,omitempty"`
	/**
	* Incomplete HTTP/HTTPS requests aborted by the client before the Application Firewall could finish processing them.
	 */
	Appfirewallabortsrate float64 `json:"appfirewallabortsrate,omitempty"`
	Appfirewallredirects  int     `json:"appfirewallredirects,omitempty"`
	/**
	* HTTP/HTTPS requests redirected by the Application Firewall to a different Web page or web server. (HTTP 302)
	 */
	Appfirewallredirectsrate float64 `json:"appfirewallredirectsrate,omitempty"`
	Misccounter0             int     `json:"misccounter0,omitempty"`
	Misccounter1             int     `json:"misccounter1,omitempty"`
	Numcpus                  int     `json:"numcpus,omitempty"`
}
