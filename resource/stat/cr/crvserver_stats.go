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

package cr

/**
* Statistics for CR virtual server resource.
*/

type Crvserverstats struct {
	/**
	* Name of a specific cache redirection virtual server.
	*/
	Name string `json:"name,omitempty"`
	/**
	* Clear the statsistics / counters
	*/
	Clearstats string `json:"clearstats,omitempty"`
	Primaryipaddress string `json:"primaryipaddress,omitempty"`
	Primaryport int32 `json:"primaryport,omitempty"`
	Type string `json:"type,omitempty"`
	State string `json:"state,omitempty"`
	Tothits uint64 `json:"tothits,omitempty"`
	/**
	* Total vserver hits
	*/
	Hitsrate float64 `json:"hitsrate,omitempty"`
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
	Totalpktsrecvd uint64 `json:"totalpktsrecvd,omitempty"`
	/**
	* Total number of packets received by this service or virtual server.
	*/
	Pktsrecvdrate float64 `json:"pktsrecvdrate,omitempty"`
	Totalpktssent uint64 `json:"totalpktssent,omitempty"`
	/**
	* Total number of packets sent.
	*/
	Pktssentrate float64 `json:"pktssentrate,omitempty"`
	Invalidrequestresponse uint64 `json:"invalidrequestresponse,omitempty"`
	Invalidrequestresponsedropped uint64 `json:"invalidrequestresponsedropped,omitempty"`

}
