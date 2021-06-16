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
	Throughput uint64 `json:"throughput,omitempty"`
	/**
	* Number of bytes received or sent by this service (Mbps).
	*/
	Throughputrate float64 `json:"throughputrate,omitempty"`
	Avgsvrttfb uint32 `json:"avgsvrttfb,omitempty"`
	Primaryipaddress string `json:"primaryipaddress,omitempty"`
	Primaryport int32 `json:"primaryport,omitempty"`
	Servicetype string `json:"servicetype,omitempty"`
	State string `json:"state,omitempty"`
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
	Totalrequestbytes []uint64 `json:"totalrequestbytes,omitempty"`
	/**
	* Total number of request bytes received on this service or virtual server.
	*/
	Requestbytesrate float64 `json:"requestbytesrate,omitempty"`
	Totalresponsebytes []uint64 `json:"totalresponsebytes,omitempty"`
	/**
	* Number of response bytes received by this service or virtual server.
	*/
	Responsebytesrate float64 `json:"responsebytesrate,omitempty"`
	Curclntconnections uint32 `json:"curclntconnections,omitempty"`
	Surgecount uint32 `json:"surgecount,omitempty"`
	Cursrvrconnections uint32 `json:"cursrvrconnections,omitempty"`
	Svrestablishedconn uint32 `json:"svrestablishedconn,omitempty"`
	Curreusepool uint32 `json:"curreusepool,omitempty"`
	Maxclients uint32 `json:"maxclients,omitempty"`
	Curload uint32 `json:"curload,omitempty"`
	Totalconnreassemblyqueue75 uint64 `json:"totalconnreassemblyqueue75,omitempty"`
	Totalconnreassemblyqueueflush uint64 `json:"totalconnreassemblyqueueflush,omitempty"`
	Httpmaxhdrszpkts uint64 `json:"httpmaxhdrszpkts,omitempty"`
	Tcpmaxooopkts uint64 `json:"tcpmaxooopkts,omitempty"`
	Curtflags uint32 `json:"curtflags,omitempty"`
	Totsvrttlbtransactions uint64 `json:"totsvrttlbtransactions,omitempty"`
	Toleratingttlbtransactions uint32 `json:"toleratingttlbtransactions,omitempty"`
	Frustratingttlbtransactions uint32 `json:"frustratingttlbtransactions,omitempty"`
	Vsvrservicehits uint64 `json:"vsvrservicehits,omitempty"`
	/**
	* Number of times that the service has been provided.
	*/
	Vsvrservicehitsrate float64 `json:"vsvrservicehitsrate,omitempty"`
	Activetransactions uint32 `json:"activetransactions,omitempty"`

}
