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

package gslb

/**
* Statistics for Global Server Load Balancing Virtual Server resource.
*/

type Gslbvserverstats struct {
	/**
	* Name of the GSLB virtual server for which to display statistics. If you do not specify a name, statistics are displayed for all GSLB virtual servers.
	*/
	Name string `json:"name,omitempty"`
	/**
	* Clear the statsistics / counters
	*/
	Clearstats string `json:"clearstats,omitempty"`
	Establishedconn uint32 `json:"establishedconn,omitempty"`
	Inactsvcs uint64 `json:"inactsvcs,omitempty"`
	Vslbhealth uint32 `json:"vslbhealth,omitempty"`
	Type string `json:"type,omitempty"`
	State string `json:"state,omitempty"`
	Actsvcs uint64 `json:"actsvcs,omitempty"`
	Tothits uint64 `json:"tothits,omitempty"`
	/**
	* Total vserver hits
	*/
	Hitsrate float64 `json:"hitsrate,omitempty"`
	Vsvrtotbkplbhits int32 `json:"vsvrtotbkplbhits,omitempty"`
	Vsvrtotbkplbfail int32 `json:"vsvrtotbkplbfail,omitempty"`
	Curpersistencesessions uint64 `json:"curpersistencesessions,omitempty"`
	Vsvrtotpersistencehits uint64 `json:"vsvrtotpersistencehits,omitempty"`
	Totalrequestbytes uint64 `json:"totalrequestbytes,omitempty"`
	/**
	* Total number of request bytes received on this service or virtual server.
	*/
	Requestbytesrate float64 `json:"requestbytesrate,omitempty"`
	Totalresponsebytes uint64 `json:"totalresponsebytes,omitempty"`
	/**
	* Number of response bytes received by this service or virtual server.
	*/
	Responsebytesrate float64 `json:"responsebytesrate,omitempty"`
	Sothreshold uint32 `json:"sothreshold,omitempty"`
	Totspillovers uint32 `json:"totspillovers,omitempty"`
	Totvserverdownbackuphits uint32 `json:"totvserverdownbackuphits,omitempty"`
	Totalrequests uint64 `json:"totalrequests,omitempty"`
	/**
	* Total number of requests received on this service or virtual server. (This applies to HTTP/SSL services and servers.)
	*/
	Requestsrate float64 `json:"requestsrate,omitempty"`
	Totalresponses uint64 `json:"totalresponses,omitempty"`
	/**
	* Number of responses received on this service or virtual server. (This applies to HTTP/SSL services and servers.)
	*/
	Responsesrate float64 `json:"responsesrate,omitempty"`
	Curclntconnections uint32 `json:"curclntconnections,omitempty"`
	Cursrvrconnections uint32 `json:"cursrvrconnections,omitempty"`

}
