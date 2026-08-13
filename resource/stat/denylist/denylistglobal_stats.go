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

package denylist

/**
* Statistics for globally active denylist binding resource.
*/

type Denylistglobalstats struct {
	/**
	* Name of the denylist label type.
	*/
	Type string `json:"type,omitempty"`
	/**
	* Clear the statsistics / counters
	*/
	Clearstats string `json:"clearstats,omitempty"`
	Lasthittimelocal string `json:"lasthittimelocal,omitempty"`
	Evals int `json:"evals,omitempty"`
	/**
	* Number of times the denylist label is invoked
	*/
	Evalsrate float64 `json:"evalsrate,omitempty"`
	Rulehits int `json:"rulehits,omitempty"`
	/**
	* Number of times the rule matched on the denylist label
	*/
	Rulehitsrate float64 `json:"rulehitsrate,omitempty"`
	Ruleundefhits int `json:"ruleundefhits,omitempty"`
	/**
	* Number of times the rule evaluated to undef on the denylist label
	*/
	Ruleundefhitsrate float64 `json:"ruleundefhitsrate,omitempty"`
	Logs int `json:"logs,omitempty"`
	/**
	* Number of times the log generated on the denylist label
	*/
	Logsrate float64 `json:"logsrate,omitempty"`
	Blocks int `json:"blocks,omitempty"`
	/**
	* Number of times the block generated on the denylist label
	*/
	Blocksrate float64 `json:"blocksrate,omitempty"`
	Lasthittime string `json:"lasthittime,omitempty"`

}
