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

package appqoe

/**
* Statistics for AppQoS policy resource.
*/

type Appqoepolicystats struct {
	/**
	* policyName
	*/
	Name string `json:"name,omitempty"`
	/**
	* Clear the statsistics / counters
	*/
	Clearstats string `json:"clearstats,omitempty"`
	Qosavgserverttfb uint64 `json:"qosavgserverttfb,omitempty"`
	/**
	* Average Server Time-To-First-Byte in milliseconds calculated for this AppQoE policy.
	*/
	Qosavgserverttfbrate float64 `json:"qosavgserverttfbrate,omitempty"`
	Qosavgserverttlb uint64 `json:"qosavgserverttlb,omitempty"`
	/**
	* Average Server Time-To-Last-Byte in milliseconds calculated for this AppQoE policy.
	*/
	Qosavgserverttlbrate float64 `json:"qosavgserverttlbrate,omitempty"`
	Qosavgclientttlb uint64 `json:"qosavgclientttlb,omitempty"`
	/**
	* Average Client Time-To-Last-Byte in milliseconds calculated for this AppQoE policy.
	*/
	Qosavgclientttlbrate float64 `json:"qosavgclientttlbrate,omitempty"`
	Totthroughput uint64 `json:"totthroughput,omitempty"`
	/**
	* Throughput in Kbps calculated on this AppQoE policy
	*/
	Throughputrate float64 `json:"throughputrate,omitempty"`
	Totsvrlinkedto uint64 `json:"totsvrlinkedto,omitempty"`
	/**
	* Total number of server connections that were established through this AppQoE Policy
	*/
	Svrlinkedtorate float64 `json:"svrlinkedtorate,omitempty"`
	Totcltrequests uint64 `json:"totcltrequests,omitempty"`
	/**
	* Total number of client connections that were requested through this AppQoE Policy
	*/
	Cltrequestsrate float64 `json:"cltrequestsrate,omitempty"`
	Totrequests uint64 `json:"totrequests,omitempty"`
	/**
	* Total number of requests that were requested through this AppQoE policy
	*/
	Requestsrate float64 `json:"requestsrate,omitempty"`
	Totrequestbytes []uint64 `json:"totrequestbytes,omitempty"`
	/**
	* Total number of requests bytes that were requested through this AppQoE policy
	*/
	Requestbytesrate float64 `json:"requestbytesrate,omitempty"`
	Totresponse uint64 `json:"totresponse,omitempty"`
	/**
	* Total number of responses received by this AppQoE policy
	*/
	Responserate float64 `json:"responserate,omitempty"`
	Totresponsebytes []uint64 `json:"totresponsebytes,omitempty"`
	/**
	* Total number of response bytes received by this AppQoE policy
	*/
	Responsebytesrate float64 `json:"responsebytesrate,omitempty"`
	Totjssent uint64 `json:"totjssent,omitempty"`
	/**
	* Total number of in-memory responses sent instead of expected responses through this AppQoE policy
	*/
	Jssentrate float64 `json:"jssentrate,omitempty"`
	Totjsbytessent uint64 `json:"totjsbytessent,omitempty"`
	/**
	* Total bytes of in-memory responses sent through this AppQoE policy
	*/
	Jsbytessentrate float64 `json:"jsbytessentrate,omitempty"`
	Pipolicyhits uint64 `json:"pipolicyhits,omitempty"`
	/**
	* Number of hits on the policy
	*/
	Pipolicyhitsrate float64 `json:"pipolicyhitsrate,omitempty"`

}
