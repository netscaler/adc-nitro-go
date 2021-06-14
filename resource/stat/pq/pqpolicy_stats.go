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

package pq

/**
* Statistics for PQ policy resource.
*/

type Pqpolicystats struct {
	/**
	* Name of the priority queuing policy whose statistics must be displayed. If a name is not provided, statistics of all priority queuing policies are shown.
	*/
	Policyname string `json:"policyname,omitempty"`
	/**
	* Clear the statsistics / counters
	*/
	Clearstats string `json:"clearstats,omitempty"`
	Pqtotavgqueuewaittime uint64 `json:"pqtotavgqueuewaittime,omitempty"`
	/**
	* Average wait time for clients for this priority queuing policy.
	*/
	Pqavgqueuewaittimerate int32 `json:"pqavgqueuewaittimerate,omitempty"`
	Pqavgclienttransactiontimems uint64 `json:"pqavgclienttransactiontimems,omitempty"`
	/**
	* Average time taken by a priority queuing client to complete its transaction for this  priority queuing policy.
	*/
	Pqavgclienttransactiontimemsrate int32 `json:"pqavgclienttransactiontimemsrate,omitempty"`
	Pqvserverip string `json:"pqvserverip,omitempty"`
	Pqvserverport int32 `json:"pqvserverport,omitempty"`
	Pqqdepth uint32 `json:"pqqdepth,omitempty"`
	/**
	* Number of clients waiting currently for this priority queuing policy.
	*/
	Pqqdepthrate int32 `json:"pqqdepthrate,omitempty"`
	Pqcurrentclientconnections uint32 `json:"pqcurrentclientconnections,omitempty"`
	/**
	* Current number of server connections established for serving clients for this priority queuing policy.
	*/
	Pqcurrentclientconnectionsrate int32 `json:"pqcurrentclientconnectionsrate,omitempty"`
	Pqtotclientconnections []uint64 `json:"pqtotclientconnections,omitempty"`
	/**
	* Total number of server connections established for serving clients for this priority queuing policy.
	*/
	Pqclientconnectionsrate int32 `json:"pqclientconnectionsrate,omitempty"`
	Pqdropped uint64 `json:"pqdropped,omitempty"`
	/**
	* Total number of dropped transactions for this priority queuing policy.
	*/
	Pqdroppedrate int32 `json:"pqdroppedrate,omitempty"`
	Totclienttransactions uint64 `json:"totclienttransactions,omitempty"`
	/**
	* Total number of client transactions for this priority queuing policy.
	*/
	Clienttransactionsrate int32 `json:"clienttransactionsrate,omitempty"`
	Pqtotqueuedepth uint64 `json:"pqtotqueuedepth,omitempty"`
	/**
	* Total number of waiting clients for this priority queuing policy.
	*/
	Pqqueuedepthrate int32 `json:"pqqueuedepthrate,omitempty"`

}
