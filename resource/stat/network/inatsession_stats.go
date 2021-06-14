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

package network

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
	Clearstats string `json:"clearstats,omitempty"`
	Globalinathits uint64 `json:"globalinathits,omitempty"`
	/**
	* Total INAT Session hits
	*/
	Globalinathitsrate int32 `json:"globalinathitsrate,omitempty"`
	Globalinatcursessions uint64 `json:"globalinatcursessions,omitempty"`
	/**
	* Total INAT current sessions
	*/
	Globalinatcursessionsrate int32 `json:"globalinatcursessionsrate,omitempty"`
	Globalinatreceivebytes uint64 `json:"globalinatreceivebytes,omitempty"`
	/**
	* Total INAT Received Bytes
	*/
	Globalinatreceivebytesrate int32 `json:"globalinatreceivebytesrate,omitempty"`
	Globalinattotsentbytes uint64 `json:"globalinattotsentbytes,omitempty"`
	/**
	* Total INAT Sent Bytes
	*/
	Globalinatsentbytesrate int32 `json:"globalinatsentbytesrate,omitempty"`
	Globalinatpktreceived uint64 `json:"globalinatpktreceived,omitempty"`
	/**
	* Total INAT Packets Received
	*/
	Globalinatpktreceivedrate int32 `json:"globalinatpktreceivedrate,omitempty"`
	Globalinatpktsent uint64 `json:"globalinatpktsent,omitempty"`
	/**
	* Total INAT Packets Sent
	*/
	Globalinatpktsentrate int32 `json:"globalinatpktsentrate,omitempty"`
	Inattothits uint64 `json:"inattothits,omitempty"`
	/**
	* INAT total sessions
	*/
	Inathitsrate int32 `json:"inathitsrate,omitempty"`
	Inatcursessions uint64 `json:"inatcursessions,omitempty"`
	/**
	* INAT current sessions
	*/
	Inatcursessionsrate int32 `json:"inatcursessionsrate,omitempty"`
	Inattotreceivebytes uint64 `json:"inattotreceivebytes,omitempty"`
	/**
	* INAT total Received Bytes
	*/
	Inatreceivebytesrate int32 `json:"inatreceivebytesrate,omitempty"`
	Inattotsentbytes uint64 `json:"inattotsentbytes,omitempty"`
	/**
	* INAT total Sent Bytes
	*/
	Inatsentbytesrate int32 `json:"inatsentbytesrate,omitempty"`
	Inattotpktreceived uint64 `json:"inattotpktreceived,omitempty"`
	/**
	* INAT total Packets Received
	*/
	Inatpktreceivedrate int32 `json:"inatpktreceivedrate,omitempty"`
	Inattotpktsent uint64 `json:"inattotpktsent,omitempty"`
	/**
	* INAT total Packets Sent
	*/
	Inatpktsentrate int32 `json:"inatpktsentrate,omitempty"`

}
