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
* Statistics for Integrated Cache content group resource.
*/

type Cachecontentgroupstats struct {
	/**
	* Name of the cache contentgroup for which to display statistics. If you do not set this parameter, statistics are shown for all cache contentgroups.
	*/
	Name string `json:"name,omitempty"`
	/**
	* Clear the statsistics / counters
	*/
	Clearstats string `json:"clearstats,omitempty"`
	Groupnon304hit uint32 `json:"groupnon304hit,omitempty"`
	Group304hit uint64 `json:"group304hit,omitempty"`
	Totcell uint32 `json:"totcell,omitempty"`
	Totmarkercell uint32 `json:"totmarkercell,omitempty"`
	Timesflushed uint32 `json:"timesflushed,omitempty"`
	Totmemory uint32 `json:"totmemory,omitempty"`
	Maxmemory uint32 `json:"maxmemory,omitempty"`

}
