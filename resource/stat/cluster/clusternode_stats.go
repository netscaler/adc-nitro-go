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

package cluster

/**
* Statistics for cluster node resource.
*/

type Clusternodestats struct {
	/**
	* ID of the cluster node for which to display statistics. If an ID is not provided, statistics are shown for all nodes.
	*/
	Nodeid uint32 `json:"nodeid,omitempty"`
	/**
	* Clear the statsistics / counters
	*/
	Clearstats string `json:"clearstats,omitempty"`
	Clsyncstate string `json:"clsyncstate,omitempty"`
	Clnodeeffectivehealth string `json:"clnodeeffectivehealth,omitempty"`
	Clnodeip string `json:"clnodeip,omitempty"`
	Clmasterstate string `json:"clmasterstate,omitempty"`
	Cltothbtx uint64 `json:"cltothbtx,omitempty"`
	Cltothbrx uint64 `json:"cltothbrx,omitempty"`
	Nnmcurconn uint32 `json:"nnmcurconn,omitempty"`
	Nnmtotconntx uint64 `json:"nnmtotconntx,omitempty"`
	Nnmtotconnrx uint64 `json:"nnmtotconnrx,omitempty"`
	Clptpstate string `json:"clptpstate,omitempty"`
	Clptptx uint64 `json:"clptptx,omitempty"`
	Clptprx uint64 `json:"clptprx,omitempty"`
	Nnmerrmsend uint64 `json:"nnmerrmsend,omitempty"`

}
