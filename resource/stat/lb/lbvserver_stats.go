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
	Sortorder            string  `json:"sortorder,omitempty"`
	Avgcltttlb           int     `json:"avgcltttlb,omitempty"`
	Cltresponsetimeapdex float64 `json:"cltresponsetimeapdex,omitempty"`
	Vsvrsurgecount       int     `json:"vsvrsurgecount,omitempty"`
	Establishedconn      int     `json:"establishedconn,omitempty"`
	Inactsvcs            int     `json:"inactsvcs,omitempty"`
	Vslbhealth           int     `json:"vslbhealth,omitempty"`
	Primaryipaddress     string  `json:"primaryipaddress,omitempty"`
	Primaryport          int     `json:"primaryport,omitempty"`
	Type                 string  `json:"type,omitempty"`
	State                string  `json:"state,omitempty"`
	Actsvcs              int     `json:"actsvcs,omitempty"`
	Cpuusagepm           int     `json:"cpuusagepm,omitempty"`
	Tothits              int     `json:"tothits,omitempty"`
	/**
	* Total vserver hits
	 */
	Hitsrate      float64 `json:"hitsrate,omitempty"`
	Totalrequests int     `json:"totalrequests,omitempty"`
	/**
	* Total number of requests received on this service or virtual server. (This applies to HTTP/SSL services and servers.)
	 */
	Requestsrate   float64 `json:"requestsrate,omitempty"`
	Totalresponses int     `json:"totalresponses,omitempty"`
	/**
	* Number of responses received on this service or virtual server. (This applies to HTTP/SSL services and servers.)
	 */
	Responsesrate     float64 `json:"responsesrate,omitempty"`
	Totalrequestbytes int     `json:"totalrequestbytes,omitempty"`
	/**
	* Total number of request bytes received on this service or virtual server.
	 */
	Requestbytesrate   float64 `json:"requestbytesrate,omitempty"`
	Totalresponsebytes int     `json:"totalresponsebytes,omitempty"`
	/**
	* Number of response bytes received by this service or virtual server.
	 */
	Responsebytesrate float64 `json:"responsebytesrate,omitempty"`
	Totalh2requests   int     `json:"totalh2requests,omitempty"`
	/**
	* Total number of Http2 requests received on this service or virtual server. (This applies to HTTP/SSL services and servers.)
	 */
	H2requestsrate   float64 `json:"h2requestsrate,omitempty"`
	Totalh2responses int     `json:"totalh2responses,omitempty"`
	/**
	* Number of Http2 responses received on this service or virtual server. (This applies to HTTP/SSL services and servers.)
	 */
	H2responsesrate float64 `json:"h2responsesrate,omitempty"`
	Totalpktsrecvd  int     `json:"totalpktsrecvd,omitempty"`
	/**
	* Total number of packets received by this service or virtual server.
	 */
	Pktsrecvdrate float64 `json:"pktsrecvdrate,omitempty"`
	Totalpktssent int     `json:"totalpktssent,omitempty"`
	/**
	* Total number of packets sent.
	 */
	Pktssentrate                 float64 `json:"pktssentrate,omitempty"`
	Curclntconnections           int     `json:"curclntconnections,omitempty"`
	Cursrvrconnections           int     `json:"cursrvrconnections,omitempty"`
	Curpersistencesessions       int     `json:"curpersistencesessions,omitempty"`
	Curbackuppersistencesessions int     `json:"curbackuppersistencesessions,omitempty"`
	Surgecount                   int     `json:"surgecount,omitempty"`
	Svcsurgecount                int     `json:"svcsurgecount,omitempty"`
	Sothreshold                  int     `json:"sothreshold,omitempty"`
	Totspillovers                int     `json:"totspillovers,omitempty"`
	Labelledconn                 int     `json:"labelledconn,omitempty"`
	Pushlabel                    int     `json:"pushlabel,omitempty"`
	Deferredreq                  int     `json:"deferredreq,omitempty"`
	/**
	* Number of deferred request on this vserver
	 */
	Deferredreqrate               float64 `json:"deferredreqrate,omitempty"`
	Invalidrequestresponse        int     `json:"invalidrequestresponse,omitempty"`
	Invalidrequestresponsedropped int     `json:"invalidrequestresponsedropped,omitempty"`
	Totvserverdownbackuphits      int     `json:"totvserverdownbackuphits,omitempty"`
	Curmptcpsessions              int     `json:"curmptcpsessions,omitempty"`
	Cursubflowconn                int     `json:"cursubflowconn,omitempty"`
	Totalconnreassemblyqueue75    int     `json:"totalconnreassemblyqueue75,omitempty"`
	Totalconnreassemblyqueueflush int     `json:"totalconnreassemblyqueueflush,omitempty"`
	Totalsvrbusyerr               int     `json:"totalsvrbusyerr,omitempty"`
	/**
	* Total no of server busy error
	 */
	Svrbusyerrrate              float64 `json:"svrbusyerrrate,omitempty"`
	Reqretrycount               int     `json:"reqretrycount,omitempty"`
	Reqretrycountexceeded       int     `json:"reqretrycountexceeded,omitempty"`
	Httpmaxhdrszpkts            int     `json:"httpmaxhdrszpkts,omitempty"`
	Tcpmaxooopkts               int     `json:"tcpmaxooopkts,omitempty"`
	Totcltttlbtransactions      int     `json:"totcltttlbtransactions,omitempty"`
	Toleratingttlbtransactions  int     `json:"toleratingttlbtransactions,omitempty"`
	Frustratingttlbtransactions int     `json:"frustratingttlbtransactions,omitempty"`
}
