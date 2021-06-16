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

package lb

/**
* Statistics for Load Balancing Virtual Server resource.
*/

type Lbvserverstats struct {
	/**
	* Name of the virtual server. If no name is provided, statistical data of all configured virtual servers is displayed.
	*/
	Name string `json:"name,omitempty"`
	/**
	* Clear the statsistics / counters
	*/
	Clearstats string `json:"clearstats,omitempty"`
	/**
	* use this argument to sort by specific key
	*/
	Sortby string `json:"sortby,omitempty"`
	/**
	* use this argument to specify sort order
	*/
	Sortorder string `json:"sortorder,omitempty"`
	Avgcltttlb uint32 `json:"avgcltttlb,omitempty"`
	Cltresponsetimeapdex float64 `json:"cltresponsetimeapdex,omitempty"`
	Vsvrsurgecount uint32 `json:"vsvrsurgecount,omitempty"`
	Establishedconn uint32 `json:"establishedconn,omitempty"`
	Inactsvcs uint64 `json:"inactsvcs,omitempty"`
	Vslbhealth uint32 `json:"vslbhealth,omitempty"`
	Primaryipaddress string `json:"primaryipaddress,omitempty"`
	Primaryport int32 `json:"primaryport,omitempty"`
	Type string `json:"type,omitempty"`
	State string `json:"state,omitempty"`
	Actsvcs uint64 `json:"actsvcs,omitempty"`
	Cpuusagepm uint64 `json:"cpuusagepm,omitempty"`
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
	Totalh2requests uint64 `json:"totalh2requests,omitempty"`
	/**
	* Total number of Http2 requests received on this service or virtual server. (This applies to HTTP/SSL services and servers.)
	*/
	H2requestsrate float64 `json:"h2requestsrate,omitempty"`
	Totalh2responses uint64 `json:"totalh2responses,omitempty"`
	/**
	* Number of Http2 responses received on this service or virtual server. (This applies to HTTP/SSL services and servers.)
	*/
	H2responsesrate float64 `json:"h2responsesrate,omitempty"`
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
	Curclntconnections uint32 `json:"curclntconnections,omitempty"`
	Cursrvrconnections uint32 `json:"cursrvrconnections,omitempty"`
	Curpersistencesessions uint64 `json:"curpersistencesessions,omitempty"`
	Curbackuppersistencesessions uint64 `json:"curbackuppersistencesessions,omitempty"`
	Surgecount uint32 `json:"surgecount,omitempty"`
	Svcsurgecount uint32 `json:"svcsurgecount,omitempty"`
	Sothreshold uint32 `json:"sothreshold,omitempty"`
	Totspillovers uint32 `json:"totspillovers,omitempty"`
	Labelledconn uint32 `json:"labelledconn,omitempty"`
	Pushlabel uint32 `json:"pushlabel,omitempty"`
	Deferredreq uint64 `json:"deferredreq,omitempty"`
	/**
	* Number of deferred request on this vserver
	*/
	Deferredreqrate float64 `json:"deferredreqrate,omitempty"`
	Invalidrequestresponse uint64 `json:"invalidrequestresponse,omitempty"`
	Invalidrequestresponsedropped uint64 `json:"invalidrequestresponsedropped,omitempty"`
	Totvserverdownbackuphits uint32 `json:"totvserverdownbackuphits,omitempty"`
	Curmptcpsessions uint32 `json:"curmptcpsessions,omitempty"`
	Cursubflowconn uint32 `json:"cursubflowconn,omitempty"`
	Totalconnreassemblyqueue75 uint64 `json:"totalconnreassemblyqueue75,omitempty"`
	Totalconnreassemblyqueueflush uint64 `json:"totalconnreassemblyqueueflush,omitempty"`
	Totalsvrbusyerr uint64 `json:"totalsvrbusyerr,omitempty"`
	/**
	* Total no of server busy error
	*/
	Svrbusyerrrate float64 `json:"svrbusyerrrate,omitempty"`
	Reqretrycount uint64 `json:"reqretrycount,omitempty"`
	Reqretrycountexceeded uint64 `json:"reqretrycountexceeded,omitempty"`
	Httpmaxhdrszpkts uint64 `json:"httpmaxhdrszpkts,omitempty"`
	Tcpmaxooopkts uint64 `json:"tcpmaxooopkts,omitempty"`
	Totcltttlbtransactions uint64 `json:"totcltttlbtransactions,omitempty"`
	Toleratingttlbtransactions uint32 `json:"toleratingttlbtransactions,omitempty"`
	Frustratingttlbtransactions uint32 `json:"frustratingttlbtransactions,omitempty"`

}
