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

package dos


type Dosstats struct {
	/**
	* Clear the statsistics / counters
	*/
	Clearstats string `json:"clearstats,omitempty"`
	Dostotconditiontriggered uint64 `json:"dostotconditiontriggered,omitempty"`
	/**
	* Number of times the Citrix ADC triggered the DOS JavaScript due to a condition match.
	*/
	Dosconditiontriggeredrate int32 `json:"dosconditiontriggeredrate,omitempty"`
	Dostotvalidcookies uint64 `json:"dostotvalidcookies,omitempty"`
	/**
	* Number of clients from whom the Citrix ADC received a valid DOS cookie.
	*/
	Dosvalidcookiesrate int32 `json:"dosvalidcookiesrate,omitempty"`
	Dostotdospriorityclients uint64 `json:"dostotdospriorityclients,omitempty"`
	/**
	* Number of valid clients that were given DOS priority.
	*/
	Dosdospriorityclientsrate int32 `json:"dosdospriorityclientsrate,omitempty"`

}
