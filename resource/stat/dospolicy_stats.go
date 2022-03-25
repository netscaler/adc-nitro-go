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
* Statistics for DoS policy resource.
 */

type Dospolicystats struct {
	/**
	* The name of the DoS protection policy whose statistics must be displayed. If a name is not provided, statistics of all the DoS protection policies are displayed.
	 */
	Name string `json:"name,omitempty"`
	/**
	* Clear the statsistics / counters
	 */
	Clearstats              string  `json:"clearstats,omitempty"`
	Doscurrentcltdetectrate float64 `json:"doscurrentcltdetectrate,omitempty"`
	Dosphysicalserviceip    string  `json:"dosphysicalserviceip,omitempty"`
	Dosphysicalserviceport  int     `json:"dosphysicalserviceport,omitempty"`
	Doscurrentqueuesize     int     `json:"doscurrentqueuesize,omitempty"`
	/**
	* Current queue size of the server to which this policy is bound.
	 */
	Doscurrentqueuesizerate float64 `json:"doscurrentqueuesizerate,omitempty"`
	Dostotjssent            int     `json:"dostotjssent,omitempty"`
	/**
	* Total number of DoS JavaScript transactions performed for this policy.
	 */
	Dosjssentrate   float64 `json:"dosjssentrate,omitempty"`
	Dostotjsrefused int     `json:"dostotjsrefused,omitempty"`
	/**
	* Number of times the DoS JavaScript was not sent because the set JavaScript rate was not met for this policy.
	 */
	Dosjsrefusedrate   float64 `json:"dosjsrefusedrate,omitempty"`
	Dostotvalidclients int     `json:"dostotvalidclients,omitempty"`
	/**
	* Total number of valid DoS cookies received for this policy.
	 */
	Dosvalidclientsrate float64 `json:"dosvalidclientsrate,omitempty"`
	Dostotjsbytessent   int     `json:"dostotjsbytessent,omitempty"`
	/**
	* Total number of DoS JavaScript bytes sent for this policy.
	 */
	Dosjsbytessentrate       float64 `json:"dosjsbytessentrate,omitempty"`
	Dostotnongetpostrequests int     `json:"dostotnongetpostrequests,omitempty"`
	/**
	* Number of non-GET and non-POST requests for which DOS JavaScript was sent.
	 */
	Dosnongetpostrequestsrate float64 `json:"dosnongetpostrequestsrate,omitempty"`
}
