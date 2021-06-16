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

package sc

/**
* Statistics for SureConnect policy resource.
*/

type Scpolicystats struct {
	/**
	* Name of the policy about which to display statistics. To display statistics about all SureConnect policies, do not set this parameter.
	*/
	Name string `json:"name,omitempty"`
	/**
	* Clear the statsistics / counters
	*/
	Clearstats string `json:"clearstats,omitempty"`
	Avgservertransactiontime uint64 `json:"avgservertransactiontime,omitempty"`
	/**
	* Average server transaction time in seconds for this SureConnect Policy.
	*/
	Avgservertransactiontimerate float64 `json:"avgservertransactiontimerate,omitempty"`
	Scaverageclientttlb uint32 `json:"scaverageclientttlb,omitempty"`
	/**
	* Average value of the client Time-To-Last-Byte in seconds for this SureConnect policy.
	*/
	Scaverageclientttlbrate float64 `json:"scaverageclientttlbrate,omitempty"`
	Scphysicalserviceip string `json:"scphysicalserviceip,omitempty"`
	Scphysicalserviceport int32 `json:"scphysicalserviceport,omitempty"`
	Sccurrentclientconnections uint32 `json:"sccurrentclientconnections,omitempty"`
	/**
	* Number of clients currently  allowed a server connection by this SureConnect policy.
	*/
	Sccurrentclientconnectionsrate float64 `json:"sccurrentclientconnectionsrate,omitempty"`
	Sccurrentwaitingclients uint32 `json:"sccurrentwaitingclients,omitempty"`
	/**
	* Current number of SureConnect priority clients that are waiting for a server connection.
	*/
	Sccurrentwaitingclientsrate float64 `json:"sccurrentwaitingclientsrate,omitempty"`
	Totopenconn uint32 `json:"totopenconn,omitempty"`
	/**
	* Current number of open connections to the servers matching this policy.
	*/
	Openconnrate float64 `json:"openconnrate,omitempty"`
	Sccurrentwaitingtime uint32 `json:"sccurrentwaitingtime,omitempty"`
	/**
	* Value of the currently estimated waiting time in seconds for the configured URL.
	*/
	Sccurrentwaitingtimerate float64 `json:"sccurrentwaitingtimerate,omitempty"`
	Sctotalclientconnections uint64 `json:"sctotalclientconnections,omitempty"`
	/**
	* Total number of clients that were allowed a server connection by this SureConnect policy.
	*/
	Scclientconnectionsrate float64 `json:"scclientconnectionsrate,omitempty"`
	Sctotalserverconnections uint64 `json:"sctotalserverconnections,omitempty"`
	/**
	* Total number of server connections that were established through this SureConnect policy.
	*/
	Scserverconnectionsrate float64 `json:"scserverconnectionsrate,omitempty"`
	Totclienttransaction uint64 `json:"totclienttransaction,omitempty"`
	/**
	* Total number of client transactions processed by this SureConnect policy.
	*/
	Clienttransactionrate float64 `json:"clienttransactionrate,omitempty"`
	Sctotalservertransactions []uint64 `json:"sctotalservertransactions,omitempty"`
	/**
	* Number of 200 OK responses received from the web server by this SureConnect policy.
	*/
	Scservertransactionsrate float64 `json:"scservertransactionsrate,omitempty"`
	Sctotalrequestsreceived uint64 `json:"sctotalrequestsreceived,omitempty"`
	/**
	* Total number of requests received by this SureConnect policy.
	*/
	Screquestsreceivedrate float64 `json:"screquestsreceivedrate,omitempty"`
	Sctotalrequestbytes uint64 `json:"sctotalrequestbytes,omitempty"`
	/**
	* Total number of request bytes received by this SureConnect policy.
	*/
	Screquestbytesrate float64 `json:"screquestbytesrate,omitempty"`
	Sctotalresponsesreceived uint64 `json:"sctotalresponsesreceived,omitempty"`
	/**
	* Total number of server responses received by this SureConnect policy.
	*/
	Scresponsesreceivedrate float64 `json:"scresponsesreceivedrate,omitempty"`
	Sctotalresponsebytes uint64 `json:"sctotalresponsebytes,omitempty"`
	/**
	* Total number of response bytes received by this SureConnect policy.
	*/
	Scresponsebytesrate float64 `json:"scresponsebytesrate,omitempty"`

}
