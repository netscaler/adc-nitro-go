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
* Statistics for GSLB service group entity resource.
 */

type Gslbservicegroupmemberstats struct {
	/**
	* Displays statistics for the specified GSLB service group.Name of the GSLB service group. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at sign (@), equal sign (=), and hyphen (-) characters.CLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my servicegroup" or 'my servicegroup').
	 */
	Servicegroupname string `json:"servicegroupname,omitempty"`
	/**
	* IP address of the GSLB service group. Mutually exclusive with the server name parameter.
	 */
	Ip string `json:"ip,omitempty"`
	/**
	* Name of the server. Mutually exclusive with the IP address parameter.
	 */
	Servername string `json:"servername,omitempty"`
	/**
	* Port number of the service group member.
	 */
	Port int `json:"port,omitempty"`
	/**
	* Clear the statsistics / counters
	 */
	Clearstats        string `json:"clearstats,omitempty"`
	Establishedconn   int    `json:"establishedconn,omitempty"`
	Primaryipaddress  string `json:"primaryipaddress,omitempty"`
	Primaryport       int    `json:"primaryport,omitempty"`
	Servicetype       string `json:"servicetype,omitempty"`
	State             string `json:"state,omitempty"`
	Totalrequestbytes int    `json:"totalrequestbytes,omitempty"`
	/**
	* Total number of request bytes received on this service or virtual server.
	 */
	Requestbytesrate   float64 `json:"requestbytesrate,omitempty"`
	Totalresponsebytes int     `json:"totalresponsebytes,omitempty"`
	/**
	* Number of response bytes received by this service or virtual server.
	 */
	Responsebytesrate float64 `json:"responsebytesrate,omitempty"`
	Curload           int     `json:"curload,omitempty"`
	Totalrequests     int     `json:"totalrequests,omitempty"`
	/**
	* Total number of requests received on this service or virtual server. (This applies to HTTP/SSL services and servers.)
	 */
	Requestsrate   float64 `json:"requestsrate,omitempty"`
	Totalresponses int     `json:"totalresponses,omitempty"`
	/**
	* Number of responses received on this service or virtual server. (This applies to HTTP/SSL services and servers.)
	 */
	Responsesrate      float64 `json:"responsesrate,omitempty"`
	Curclntconnections int     `json:"curclntconnections,omitempty"`
	Cursrvrconnections int     `json:"cursrvrconnections,omitempty"`
}
