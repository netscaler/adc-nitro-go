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

package lsn

/**
* Statistics for LSN pool resource.
*/

type Lsnpoolstats struct {
	/**
	* Name of the LSN Pool.
	*/
	Poolname string `json:"poolname,omitempty"`
	/**
	* Clear the statsistics / counters
	*/
	Clearstats string `json:"clearstats,omitempty"`
	Lsnpoolotherportusagepercen float64 `json:"lsnpoolotherportusagepercen,omitempty"`
	Lsnpooltcpportusagepercen float64 `json:"lsnpooltcpportusagepercen,omitempty"`
	Lsnpoolcurportalloctcp uint64 `json:"lsnpoolcurportalloctcp,omitempty"`
	/**
	* Current port allocation for tcp in this pool.
	*/
	Lsnpoolcurportalloctcprate int32 `json:"lsnpoolcurportalloctcprate,omitempty"`
	Lsnpoolcurportallocother uint64 `json:"lsnpoolcurportallocother,omitempty"`
	/**
	* Current port allocation for other protocols in this pool.
	*/
	Lsnpoolcurportallocotherrate int32 `json:"lsnpoolcurportallocotherrate,omitempty"`
	Lsnpooltotips uint32 `json:"lsnpooltotips,omitempty"`

}
