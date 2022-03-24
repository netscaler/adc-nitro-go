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

type Pqstats struct {
	/**
	* Clear the statsistics / counters
	 */
	Clearstats           string `json:"clearstats,omitempty"`
	Pqtotalpolicymatches int    `json:"pqtotalpolicymatches,omitempty"`
	/**
	* Number of times the Citrix ADC matched an incoming request using any priority queuing policy.
	 */
	Pqpolicymatchesrate    float64 `json:"pqpolicymatchesrate,omitempty"`
	Pqtotalthresholdfailed int     `json:"pqtotalthresholdfailed,omitempty"`
	/**
	* Number of times the Citrix ADC failed to match an incoming request to any of priority queing policy.
	 */
	Pqthresholdfailedrate float64 `json:"pqthresholdfailedrate,omitempty"`
	Pqpriority1requests   int     `json:"pqpriority1requests,omitempty"`
	/**
	* Number of priority 1 requests that the Citrix ADC received.
	 */
	Pqpriority1requestsrate float64 `json:"pqpriority1requestsrate,omitempty"`
	Pqpriority2requests     int     `json:"pqpriority2requests,omitempty"`
	/**
	* Number of priority 2 requests that the Citrix ADC received.
	 */
	Pqpriority2requestsrate float64 `json:"pqpriority2requestsrate,omitempty"`
	Pqpriority3requests     int     `json:"pqpriority3requests,omitempty"`
	/**
	* Number of priority 3 requests that the Citrix ADC received.
	 */
	Pqpriority3requestsrate float64 `json:"pqpriority3requestsrate,omitempty"`
}
