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
* Statistics for stateful INAT resource.
 */

type Inatsessionstats struct {
	/**
	* INAT name
	 */
	Name string `json:"name,omitempty"`
	/**
	* Clear the statsistics / counters
	 */
	Clearstats     string `json:"clearstats,omitempty"`
	Globalinathits int    `json:"globalinathits,omitempty"`
	/**
	* Total INAT Session hits
	 */
	Globalinathitsrate    float64 `json:"globalinathitsrate,omitempty"`
	Globalinatcursessions int     `json:"globalinatcursessions,omitempty"`
	/**
	* Total INAT current sessions
	 */
	Globalinatcursessionsrate float64 `json:"globalinatcursessionsrate,omitempty"`
	Globalinatreceivebytes    int     `json:"globalinatreceivebytes,omitempty"`
	/**
	* Total INAT Received Bytes
	 */
	Globalinatreceivebytesrate float64 `json:"globalinatreceivebytesrate,omitempty"`
	Globalinattotsentbytes     int     `json:"globalinattotsentbytes,omitempty"`
	/**
	* Total INAT Sent Bytes
	 */
	Globalinatsentbytesrate float64 `json:"globalinatsentbytesrate,omitempty"`
	Globalinatpktreceived   int     `json:"globalinatpktreceived,omitempty"`
	/**
	* Total INAT Packets Received
	 */
	Globalinatpktreceivedrate float64 `json:"globalinatpktreceivedrate,omitempty"`
	Globalinatpktsent         int     `json:"globalinatpktsent,omitempty"`
	/**
	* Total INAT Packets Sent
	 */
	Globalinatpktsentrate float64 `json:"globalinatpktsentrate,omitempty"`
	Inattothits           int     `json:"inattothits,omitempty"`
	/**
	* INAT total sessions
	 */
	Inathitsrate    float64 `json:"inathitsrate,omitempty"`
	Inatcursessions int     `json:"inatcursessions,omitempty"`
	/**
	* INAT current sessions
	 */
	Inatcursessionsrate float64 `json:"inatcursessionsrate,omitempty"`
	Inattotreceivebytes int     `json:"inattotreceivebytes,omitempty"`
	/**
	* INAT total Received Bytes
	 */
	Inatreceivebytesrate float64 `json:"inatreceivebytesrate,omitempty"`
	Inattotsentbytes     int     `json:"inattotsentbytes,omitempty"`
	/**
	* INAT total Sent Bytes
	 */
	Inatsentbytesrate  float64 `json:"inatsentbytesrate,omitempty"`
	Inattotpktreceived int     `json:"inattotpktreceived,omitempty"`
	/**
	* INAT total Packets Received
	 */
	Inatpktreceivedrate float64 `json:"inatpktreceivedrate,omitempty"`
	Inattotpktsent      int     `json:"inattotpktsent,omitempty"`
	/**
	* INAT total Packets Sent
	 */
	Inatpktsentrate float64 `json:"inatpktsentrate,omitempty"`
}
