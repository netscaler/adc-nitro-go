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
* Configuration for PQ bindings resource.
 */
type Pqbinding struct {
	/**
	* Name of the load balancing virtual server for which to display the priority queuing policy.
	 */
	Vservername string `json:"vservername,omitempty"`

	//------- Read only Parameter ---------;

	Policyname string `json:"policyname,omitempty"`
	Rule       string `json:"rule,omitempty"`
	Priority   string `json:"priority,omitempty"`
	Weight     string `json:"weight,omitempty"`
	Qdepth     string `json:"qdepth,omitempty"`
	Polqdepth  string `json:"polqdepth,omitempty"`
	Hits       string `json:"hits,omitempty"`
}
