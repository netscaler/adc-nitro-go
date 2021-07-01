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
* Statistics for GSLB site resource.
*/

type Gslbsitestats struct {
	/**
	* Name of the GSLB site for which to display detailed statistics. If a name is not specified, basic information about all GSLB sites is displayed.
	*/
	Sitename string `json:"sitename,omitempty"`
	/**
	* Clear the statsistics / counters
	*/
	Clearstats string `json:"clearstats,omitempty"`
	Sitepublicip string `json:"sitepublicip,omitempty"`
	Siteip string `json:"siteip,omitempty"`
	Sitemepstatus string `json:"sitemepstatus,omitempty"`
	Persexchange string `json:"persexchange,omitempty"`
	Nwmetricexchange string `json:"nwmetricexchange,omitempty"`
	Sitemetricexchange string `json:"sitemetricexchange,omitempty"`
	Sitetype string `json:"sitetype,omitempty"`
	Siteipstr string `json:"siteipstr,omitempty"`
	Sitepublicipstr string `json:"sitepublicipstr,omitempty"`
	Sitemetricmepstatus string `json:"sitemetricmepstatus,omitempty"`
	Nwmetricmepstatus string `json:"nwmetricmepstatus,omitempty"`
	Sitetotalrequestbytes int `json:"sitetotalrequestbytes,omitempty"`
	/**
	* Total number of request bytes received by the virtual servers represented by all GSLB services associated with this GSLB site.
	*/
	Siterequestbytesrate float64 `json:"siterequestbytesrate,omitempty"`
	Sitetotalresponsebytes int `json:"sitetotalresponsebytes,omitempty"`
	/**
	* Number of response bytes received by the virtual servers represented by all GSLB services associated with this GSLB site.
	*/
	Siteresponsebytesrate float64 `json:"siteresponsebytesrate,omitempty"`
	Sitetotalrequests int `json:"sitetotalrequests,omitempty"`
	/**
	* Total number of requests received by the virtual servers represented by all GSLB services associated with this GSLB site.
	*/
	Siterequestsrate float64 `json:"siterequestsrate,omitempty"`
	Sitetotalresponses int `json:"sitetotalresponses,omitempty"`
	/**
	* Number of responses received by the virtual servers represented by all GSLB services associated with this GSLB site.
	*/
	Siteresponsesrate float64 `json:"siteresponsesrate,omitempty"`
	Sitecurclntconnections int `json:"sitecurclntconnections,omitempty"`
	Sitecursrvrconnections int `json:"sitecursrvrconnections,omitempty"`

}
