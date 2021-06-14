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

package cs

/**
* Statistics for CS virtual server resource.
*/

type Csvserverstats struct {
	/**
	* Name of the content switching virtual server for which to display statistics. To display statistics for all configured Content Switching virtual servers, do not specify a value for this parameter.
	*/
	Name string `json:"name,omitempty"`
	/**
	* Clear the statsistics / counters
	*/
	Clearstats string `json:"clearstats,omitempty"`
	Avgcltttlb uint32 `json:"avgcltttlb,omitempty"`
	Cltresponsetimeapdex float64 `json:"cltresponsetimeapdex,omitempty"`
	Establishedconn uint32 `json:"establishedconn,omitempty"`
	Primaryipaddress string `json:"primaryipaddress,omitempty"`
	Primaryport int32 `json:"primaryport,omitempty"`
	Type string `json:"type,omitempty"`
	State string `json:"state,omitempty"`
	Tothits uint64 `json:"tothits,omitempty"`
	/**
	* Total vserver hits
	*/
	Hitsrate int32 `json:"hitsrate,omitempty"`
	Totalrequests uint64 `json:"totalrequests,omitempty"`
	/**
	* Total number of requests received on this service or virtual server. (This applies to HTTP/SSL services and servers.)
	*/
	Requestsrate int32 `json:"requestsrate,omitempty"`
	Totalresponses uint64 `json:"totalresponses,omitempty"`
	/**
	* Number of responses received on this service or virtual server. (This applies to HTTP/SSL services and servers.)
	*/
	Responsesrate int32 `json:"responsesrate,omitempty"`
	Totalrequestbytes uint64 `json:"totalrequestbytes,omitempty"`
	/**
	* Total number of request bytes received on this service or virtual server.
	*/
	Requestbytesrate int32 `json:"requestbytesrate,omitempty"`
	Totalresponsebytes uint64 `json:"totalresponsebytes,omitempty"`
	/**
	* Number of response bytes received by this service or virtual server.
	*/
	Responsebytesrate int32 `json:"responsebytesrate,omitempty"`
	Totalpktsrecvd uint64 `json:"totalpktsrecvd,omitempty"`
	/**
	* Total number of packets received by this service or virtual server.
	*/
	Pktsrecvdrate int32 `json:"pktsrecvdrate,omitempty"`
	Totalpktssent uint64 `json:"totalpktssent,omitempty"`
	/**
	* Total number of packets sent.
	*/
	Pktssentrate int32 `json:"pktssentrate,omitempty"`
	Curclntconnections uint32 `json:"curclntconnections,omitempty"`
	Cursrvrconnections uint32 `json:"cursrvrconnections,omitempty"`
	Curpersistencesessions uint64 `json:"curpersistencesessions,omitempty"`
	Curbackuppersistencesessions uint64 `json:"curbackuppersistencesessions,omitempty"`
	Sothreshold uint32 `json:"sothreshold,omitempty"`
	Totspillovers uint32 `json:"totspillovers,omitempty"`
	Labelledconn uint32 `json:"labelledconn,omitempty"`
	Pushlabel uint32 `json:"pushlabel,omitempty"`
	Deferredreq uint64 `json:"deferredreq,omitempty"`
	/**
	* Number of deferred request on this vserver
	*/
	Deferredreqrate int32 `json:"deferredreqrate,omitempty"`
	Invalidrequestresponse uint64 `json:"invalidrequestresponse,omitempty"`
	Invalidrequestresponsedropped uint64 `json:"invalidrequestresponsedropped,omitempty"`
	Totvserverdownbackuphits uint32 `json:"totvserverdownbackuphits,omitempty"`
	Curmptcpsessions uint32 `json:"curmptcpsessions,omitempty"`
	Cursubflowconn uint32 `json:"cursubflowconn,omitempty"`
	Httpmaxhdrszpkts uint64 `json:"httpmaxhdrszpkts,omitempty"`
	Tcpmaxooopkts uint64 `json:"tcpmaxooopkts,omitempty"`
	Totcltttlbtransactions uint64 `json:"totcltttlbtransactions,omitempty"`
	Toleratingttlbtransactions uint32 `json:"toleratingttlbtransactions,omitempty"`
	Frustratingttlbtransactions uint32 `json:"frustratingttlbtransactions,omitempty"`

}
