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

package ns

/**
* Statistics for admin partition resource.
*/

type Nspartitionstats struct {
	/**
	* Name of the partition.
	*/
	Partitionname string `json:"partitionname,omitempty"`
	/**
	* Clear the statsistics / counters
	*/
	Clearstats string `json:"clearstats,omitempty"`
	Maxbandwidth uint32 `json:"maxbandwidth,omitempty"`
	Maxconnection uint32 `json:"maxconnection,omitempty"`
	Maxmemory uint32 `json:"maxmemory,omitempty"`
	Currentbandwidth uint32 `json:"currentbandwidth,omitempty"`
	Currentconnections uint32 `json:"currentconnections,omitempty"`
	Memoryusagepcnt uint32 `json:"memoryusagepcnt,omitempty"`
	Totaldrops uint32 `json:"totaldrops,omitempty"`
	/**
	* Total packet drops for the partition.
	*/
	Dropsrate float64 `json:"dropsrate,omitempty"`
	Totaltokendrops uint32 `json:"totaltokendrops,omitempty"`
	/**
	* Total drops(KB) for the partition.
	*/
	Tokendropsrate float64 `json:"tokendropsrate,omitempty"`
	Totalconnectiondrops uint32 `json:"totalconnectiondrops,omitempty"`
	/**
	* Total connection drops for the partition.
	*/
	Connectiondropsrate float64 `json:"connectiondropsrate,omitempty"`

}
