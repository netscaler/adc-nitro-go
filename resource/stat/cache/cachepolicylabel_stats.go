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

package cache

/**
* Statistics for cache policy label resource.
*/

type Cachepolicylabelstats struct {
	/**
	* Name of the cache-policy label for which to display statistics. If you do not set this parameter statistics are shown for all cache-policy labels.
	*/
	Labelname string `json:"labelname,omitempty"`
	/**
	* Clear the statsistics / counters
	*/
	Clearstats string `json:"clearstats,omitempty"`
	Pipolicylabelhits uint64 `json:"pipolicylabelhits,omitempty"`
	/**
	* Number of times policy label was invoked. 
	*/
	Pipolicylabelhitsrate int32 `json:"pipolicylabelhitsrate,omitempty"`

}
