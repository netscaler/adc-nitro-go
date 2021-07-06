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
	Nodeid int `json:"nodeid,omitempty"`
	/**
	* Clear the statsistics / counters
	*/
	Clearstats string `json:"clearstats,omitempty"`
	Clsyncstate string `json:"clsyncstate,omitempty"`
	Clnodeeffectivehealth string `json:"clnodeeffectivehealth,omitempty"`
	Clnodeip string `json:"clnodeip,omitempty"`
	Clmasterstate string `json:"clmasterstate,omitempty"`
	Cltothbtx int `json:"cltothbtx,omitempty"`
	Cltothbrx int `json:"cltothbrx,omitempty"`
	Nnmcurconn int `json:"nnmcurconn,omitempty"`
	Nnmtotconntx int `json:"nnmtotconntx,omitempty"`
	Nnmtotconnrx int `json:"nnmtotconnrx,omitempty"`
	Clptpstate string `json:"clptpstate,omitempty"`
	Clptptx int `json:"clptptx,omitempty"`
	Clptprx int `json:"clptprx,omitempty"`
	Nnmerrmsend int `json:"nnmerrmsend,omitempty"`

}
