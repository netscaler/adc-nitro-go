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
	Clearstats            string `json:"clearstats,omitempty"`
	Pqtotavgqueuewaittime int    `json:"pqtotavgqueuewaittime,omitempty"`
	/**
	* Average wait time for clients for this priority queuing policy.
	 */
	Pqavgqueuewaittimerate       float64 `json:"pqavgqueuewaittimerate,omitempty"`
	Pqavgclienttransactiontimems int     `json:"pqavgclienttransactiontimems,omitempty"`
	/**
	* Average time taken by a priority queuing client to complete its transaction for this  priority queuing policy.
	 */
	Pqavgclienttransactiontimemsrate float64 `json:"pqavgclienttransactiontimemsrate,omitempty"`
	Pqvserverip                      string  `json:"pqvserverip,omitempty"`
	Pqvserverport                    int     `json:"pqvserverport,omitempty"`
	Pqqdepth                         int     `json:"pqqdepth,omitempty"`
	/**
	* Number of clients waiting currently for this priority queuing policy.
	 */
	Pqqdepthrate               float64 `json:"pqqdepthrate,omitempty"`
	Pqcurrentclientconnections int     `json:"pqcurrentclientconnections,omitempty"`
	/**
	* Current number of server connections established for serving clients for this priority queuing policy.
	 */
	Pqcurrentclientconnectionsrate float64 `json:"pqcurrentclientconnectionsrate,omitempty"`
	Pqtotclientconnections         []int   `json:"pqtotclientconnections,omitempty"`
	/**
	* Total number of server connections established for serving clients for this priority queuing policy.
	 */
	Pqclientconnectionsrate float64 `json:"pqclientconnectionsrate,omitempty"`
	Pqdropped               int     `json:"pqdropped,omitempty"`
	/**
	* Total number of dropped transactions for this priority queuing policy.
	 */
	Pqdroppedrate         float64 `json:"pqdroppedrate,omitempty"`
	Totclienttransactions int     `json:"totclienttransactions,omitempty"`
	/**
	* Total number of client transactions for this priority queuing policy.
	 */
	Clienttransactionsrate float64 `json:"clienttransactionsrate,omitempty"`
	Pqtotqueuedepth        int     `json:"pqtotqueuedepth,omitempty"`
	/**
	* Total number of waiting clients for this priority queuing policy.
	 */
	Pqqueuedepthrate float64 `json:"pqqueuedepthrate,omitempty"`
}
