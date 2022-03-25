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
* Statistics for  resource.
 */

type Nsmemorystats struct {
	/**
	* Feature name for which to display memory statistics.
	 */
	Pool string `json:"pool,omitempty"`
	/**
	* Clear the statsistics / counters
	 */
	Clearstats     string  `json:"clearstats,omitempty"`
	Allocf         float64 `json:"allocf,omitempty"`
	Memcurallocper float64 `json:"memcurallocper,omitempty"`
	Memcurinkb     int     `json:"memcurinkb,omitempty"`
}
