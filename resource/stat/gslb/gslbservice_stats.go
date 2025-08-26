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
* Statistics for GSLB service resource.
*/

type Gslbservicestats struct {
	/**
	* Name of the GSLB service.
	*/
	Servicename string `json:"servicename,omitempty"`
	/**
	* Clear the statsistics / counters
	*/
	Clearstats string `json:"clearstats,omitempty"`
	Establishedconn int `json:"establishedconn,omitempty"`
	Primaryipaddress string `json:"primaryipaddress,omitempty"`
	Primaryport int `json:"primaryport,omitempty"`
	Servicetype string `json:"servicetype,omitempty"`
	State string `json:"state,omitempty"`
	Totalrequestbytes int `json:"totalrequestbytes,omitempty"`
	/**
	* Total number of request bytes received on this service or virtual server.
	*/
	Requestbytesrate float64 `json:"requestbytesrate,omitempty"`
	Totalresponsebytes int `json:"totalresponsebytes,omitempty"`
	/**
	* Number of response bytes received by this service or virtual server.
	*/
	Responsebytesrate float64 `json:"responsebytesrate,omitempty"`
	Curload int `json:"curload,omitempty"`
	Totalrequests int `json:"totalrequests,omitempty"`
	/**
	* Total number of requests received on this service or virtual server. (This applies to HTTP/SSL services and servers.)
	*/
	Requestsrate float64 `json:"requestsrate,omitempty"`
	Totalresponses int `json:"totalresponses,omitempty"`
	/**
	* Number of responses received on this service or virtual server. (This applies to HTTP/SSL services and servers.)
	*/
	Responsesrate float64 `json:"responsesrate,omitempty"`
	Curclntconnections int `json:"curclntconnections,omitempty"`
	Cursrvrconnections int `json:"cursrvrconnections,omitempty"`
	Vsvrservicehits int `json:"vsvrservicehits,omitempty"`
	/**
	* Number of times that the service has been provided.
	*/
	Vsvrservicehitsrate float64 `json:"vsvrservicehitsrate,omitempty"`
	Serviceorder int `json:"serviceorder,omitempty"`

}
