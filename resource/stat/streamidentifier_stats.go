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
* Statistics for identifier resource.
 */

type Streamidentifierstats struct {
	/**
	* Name of the stream identifier.
	 */
	Name string `json:"name,omitempty"`
	/**
	* Values on which grouping is performed are displayed in the output as row titles. If grouping is performed on two or more fields, their values are separated by a question mark in the row title.
		For example, consider a selector that contains the expressions HTTP.REQ.URL and CLIENT.IP.SRC (in that order), on an appliance that has accumulated records of a number of requests for two URLs, example.com/page1.html and example.com/page2.html, from two client IP addresses, 192.0.2.10 and 192.0.2.11.
		With a pattern of ? ?, the appliance performs grouping on both fields and displays statistics for the following:
		* Requests for example.com/abc.html from 192.0.2.10, with a row title of example.com/abc.html?192.0.2.10.
		* Requests for example.com/abc.html from 192.0.2.11, with a row title of example.com/abc.html?192.0.2.11.
		* Requests for example.com/def.html from 192.0.2.10, with a row title of example.com/def.html?192.0.2.10.
		* Requests for example.com/def.html from 192.0.2.11, with a row title of example.com/def.html?192.0.2.11.
		With a pattern of * ?, the appliance performs grouping on only the client IP address values and displays statistics for the following requests:
		* All requests from 192.0.2.10, with the IP address as the row title.
		* All requests from 192.0.2.11, with the IP address as the row title.
		With a pattern of ? *, the appliance performs grouping on only the URL values and displays statistics for the following requests:
		* All requests for example.com/abc.html, with the URL as the row title.
		* All requests for example.com/def.html, with the URL as the row title.
		With a pattern of * *, the appliance displays one set of collective statistics for all the requests received, with no row title.
		With a pattern of example.com/abc.html ?, the appliance displays statistics for requests for example.com/abc.html from each unique client IP address.
		With a pattern of * 192.0.2.11, the appliance displays statistics for all requests from 192.0.2.11.
	*/
	Pattern []string `json:"pattern,omitempty"`
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
	Sortorder              string `json:"sortorder,omitempty"`
	Streamobjreq           int    `json:"streamobjreq,omitempty"`
	Streamobjbandw         int    `json:"streamobjbandw,omitempty"`
	Streamobjresptime      int    `json:"streamobjresptime,omitempty"`
	Streamobjconn          int    `json:"streamobjconn,omitempty"`
	Streamobjbreachcnt     int    `json:"streamobjbreachcnt,omitempty"`
	Streamobjpktcredits    int    `json:"streamobjpktcredits,omitempty"`
	Streamobjpktspersecond int    `json:"streamobjpktspersecond,omitempty"`
	Streamobjdroppedconns  int    `json:"streamobjdroppedconns,omitempty"`
}
