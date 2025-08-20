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

package api

/**
* Statistics for API specification resource.
*/

type Apispecstats struct {
	/**
	* Name of the api spec for which to display stats.
	*/
	Name string `json:"name,omitempty"`
	/**
	* Clear the statsistics / counters
	*/
	Clearstats string `json:"clearstats,omitempty"`
	Apispechits int `json:"apispechits,omitempty"`
	/**
	* Number of received requests for the API specification.
	*/
	Apispechitsrate float64 `json:"apispechitsrate,omitempty"`
	Apispecsuccessfullyvalidated int `json:"apispecsuccessfullyvalidated,omitempty"`
	/**
	* Number of successfully validated requests for the API specification.
	*/
	Apispecsuccessfullyvalidatedrate float64 `json:"apispecsuccessfullyvalidatedrate,omitempty"`
	Apispecunmatched int `json:"apispecunmatched,omitempty"`
	/**
	* Number of requests non matching this API specification.
	*/
	Apispecunmatchedrate float64 `json:"apispecunmatchedrate,omitempty"`

}
