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

package system

/**
* Statistics for Global system memory stats resource.
 */

type Systemmemorystats struct {
	/**
	* Clear the statsistics / counters
	 */
	Clearstats            string  `json:"clearstats,omitempty"`
	Cacmemmaxmemlimitpcnt float64 `json:"cacmemmaxmemlimitpcnt,omitempty"`
	Cacmemmaxmemlimit     int     `json:"cacmemmaxmemlimit,omitempty"`
	Shmemerrallocfailed   int     `json:"shmemerrallocfailed,omitempty"`
	Shmemallocpcnt        float64 `json:"shmemallocpcnt,omitempty"`
	Shmemallocinmb        int     `json:"shmemallocinmb,omitempty"`
	Shmemtotinmb          int     `json:"shmemtotinmb,omitempty"`
	Memtotallocfailed     int     `json:"memtotallocfailed,omitempty"`
	Memtotfree            int     `json:"memtotfree,omitempty"`
	Memusagepcnt          float64 `json:"memusagepcnt,omitempty"`
	Memtotuseinmb         int     `json:"memtotuseinmb,omitempty"`
	Memtotallocpcnt       float64 `json:"memtotallocpcnt,omitempty"`
	Memtotallocmb         int     `json:"memtotallocmb,omitempty"`
	Memtotinmb            int     `json:"memtotinmb,omitempty"`
	Memtotavail           int     `json:"memtotavail,omitempty"`
	Cacmemmaxsyslimitmb   int     `json:"cacmemmaxsyslimitmb,omitempty"`
}
