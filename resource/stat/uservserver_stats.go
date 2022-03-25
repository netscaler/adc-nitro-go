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

/**
* Statistics for virtual server resource.
 */

type Uservserverstats struct {
	/**
	* Name of the user defined virtual server. If no name is provided, statistical data of all configured user defined virtual servers is displayed.
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
	Sortorder        string `json:"sortorder,omitempty"`
	Establishedconn  int    `json:"establishedconn,omitempty"`
	Inactsvcs        int    `json:"inactsvcs,omitempty"`
	Vslbhealth       int    `json:"vslbhealth,omitempty"`
	Primaryipaddress string `json:"primaryipaddress,omitempty"`
	Primaryport      int    `json:"primaryport,omitempty"`
	Protocolname     string `json:"protocolname,omitempty"`
	State            string `json:"state,omitempty"`
	Actsvcs          int    `json:"actsvcs,omitempty"`
	Tothits          int    `json:"tothits,omitempty"`
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
	Totalpktsrecvd    int     `json:"totalpktsrecvd,omitempty"`
	/**
	* Total number of packets received by this service or virtual server.
	 */
	Pktsrecvdrate float64 `json:"pktsrecvdrate,omitempty"`
	Totalpktssent int     `json:"totalpktssent,omitempty"`
	/**
	* Total number of packets sent.
	 */
	Pktssentrate                  float64 `json:"pktssentrate,omitempty"`
	Curclntconnections            int     `json:"curclntconnections,omitempty"`
	Cursrvrconnections            int     `json:"cursrvrconnections,omitempty"`
	Invalidrequestresponse        int     `json:"invalidrequestresponse,omitempty"`
	Invalidrequestresponsedropped int     `json:"invalidrequestresponsedropped,omitempty"`
}
