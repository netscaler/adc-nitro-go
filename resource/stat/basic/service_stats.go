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

package basic

/**
* Statistics for service resource.
 */

type Servicestats struct {
	/**
	* Name of the service.
	 */
	Name string `json:"name,omitempty"`
	/**
	* Clear the statsistics / counters
	 */
	Clearstats string `json:"clearstats,omitempty"`
	Throughput int    `json:"throughput,omitempty"`
	/**
	* Number of bytes received or sent by this service (Mbps).
	 */
	Throughputrate   float64 `json:"throughputrate,omitempty"`
	Avgsvrttfb       int     `json:"avgsvrttfb,omitempty"`
	Primaryipaddress string  `json:"primaryipaddress,omitempty"`
	Primaryport      int     `json:"primaryport,omitempty"`
	Servicetype      string  `json:"servicetype,omitempty"`
	State            string  `json:"state,omitempty"`
	Totalrequests    int     `json:"totalrequests,omitempty"`
	/**
	* Total number of requests received on this service or virtual server. (This applies to HTTP/SSL services and servers.)
	 */
	Requestsrate   float64 `json:"requestsrate,omitempty"`
	Totalresponses int     `json:"totalresponses,omitempty"`
	/**
	* Number of responses received on this service or virtual server. (This applies to HTTP/SSL services and servers.)
	 */
	Responsesrate     float64 `json:"responsesrate,omitempty"`
	Totalrequestbytes []int   `json:"totalrequestbytes,omitempty"`
	/**
	* Total number of request bytes received on this service or virtual server.
	 */
	Requestbytesrate   float64 `json:"requestbytesrate,omitempty"`
	Totalresponsebytes []int   `json:"totalresponsebytes,omitempty"`
	/**
	* Number of response bytes received by this service or virtual server.
	 */
	Responsebytesrate             float64 `json:"responsebytesrate,omitempty"`
	Curclntconnections            int     `json:"curclntconnections,omitempty"`
	Surgecount                    int     `json:"surgecount,omitempty"`
	Cursrvrconnections            int     `json:"cursrvrconnections,omitempty"`
	Svrestablishedconn            int     `json:"svrestablishedconn,omitempty"`
	Curreusepool                  int     `json:"curreusepool,omitempty"`
	Maxclients                    int     `json:"maxclients,omitempty"`
	Curload                       int     `json:"curload,omitempty"`
	Totalconnreassemblyqueue75    int     `json:"totalconnreassemblyqueue75,omitempty"`
	Totalconnreassemblyqueueflush int     `json:"totalconnreassemblyqueueflush,omitempty"`
	Httpmaxhdrszpkts              int     `json:"httpmaxhdrszpkts,omitempty"`
	Tcpmaxooopkts                 int     `json:"tcpmaxooopkts,omitempty"`
	Curtflags                     int     `json:"curtflags,omitempty"`
	Totsvrttlbtransactions        int     `json:"totsvrttlbtransactions,omitempty"`
	Toleratingttlbtransactions    int     `json:"toleratingttlbtransactions,omitempty"`
	Frustratingttlbtransactions   int     `json:"frustratingttlbtransactions,omitempty"`
	Vsvrservicehits               int     `json:"vsvrservicehits,omitempty"`
	/**
	* Number of times that the service has been provided.
	 */
	Vsvrservicehitsrate float64 `json:"vsvrservicehitsrate,omitempty"`
	Activetransactions  int     `json:"activetransactions,omitempty"`
}
