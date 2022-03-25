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
	Clearstats     string `json:"clearstats,omitempty"`
	Groupnon304hit int    `json:"groupnon304hit,omitempty"`
	Group304hit    int    `json:"group304hit,omitempty"`
	Totcell        int    `json:"totcell,omitempty"`
	Totmarkercell  int    `json:"totmarkercell,omitempty"`
	Timesflushed   int    `json:"timesflushed,omitempty"`
	Totmemory      int    `json:"totmemory,omitempty"`
	Maxmemory      int    `json:"maxmemory,omitempty"`
}
