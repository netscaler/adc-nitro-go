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

package pcp

/**
* Statistics for server resource.
 */

type Pcpserverstats struct {
	/**
	* PCP Statistics per Server
	 */
	Name string `json:"name,omitempty"`
	/**
	* Clear the statsistics / counters
	 */
	Clearstats     string `json:"clearstats,omitempty"`
	Pcptotrequests int    `json:"pcptotrequests,omitempty"`
	/**
	* PCP Request received.
	 */
	Pcprequestsrate float64 `json:"pcprequestsrate,omitempty"`
	Pcptotresponses int     `json:"pcptotresponses,omitempty"`
	/**
	* Number PCP responces sent.
	 */
	Pcpresponsesrate  float64 `json:"pcpresponsesrate,omitempty"`
	Pcptotmaprequests int     `json:"pcptotmaprequests,omitempty"`
	/**
	* PCP MAP Requests received.
	 */
	Pcpmaprequestsrate float64 `json:"pcpmaprequestsrate,omitempty"`
	Pcptotpeerrequests int     `json:"pcptotpeerrequests,omitempty"`
	/**
	* PCP PEER Requests received.
	 */
	Pcppeerrequestsrate float64 `json:"pcppeerrequestsrate,omitempty"`
	Pcptoterrinrquest   int     `json:"pcptoterrinrquest,omitempty"`
	/**
	* total PCP request with error.
	 */
	Pcperrinrquestrate float64 `json:"pcperrinrquestrate,omitempty"`
	Pcptoterrinres     int     `json:"pcptoterrinres,omitempty"`
	/**
	* Total PCP responses with errors.
	 */
	Pcperrinresrate  float64 `json:"pcperrinresrate,omitempty"`
	Pcptotunsuppvers int     `json:"pcptotunsuppvers,omitempty"`
	/**
	* PCP request with unsupported version.
	 */
	Pcpunsuppversrate  float64 `json:"pcpunsuppversrate,omitempty"`
	Pcptotmalformedreq int     `json:"pcptotmalformedreq,omitempty"`
	/**
	* total PCP request having malformed PCP packets.
	 */
	Pcpmalformedreqrate float64 `json:"pcpmalformedreqrate,omitempty"`
	Pcptotunsuppopcode  int     `json:"pcptotunsuppopcode,omitempty"`
	/**
	* total Unsupproted OPCODES received Requests.
	 */
	Pcpunsuppopcoderate float64 `json:"pcpunsuppopcoderate,omitempty"`
	Pcptotunsuppoption  int     `json:"pcptotunsuppoption,omitempty"`
	/**
	* total Unsupproted OPTIONS received in requests.
	 */
	Pcpunsuppoptionrate   float64 `json:"pcpunsuppoptionrate,omitempty"`
	Pcptotmalformedoption int     `json:"pcptotmalformedoption,omitempty"`
	/**
	* total malformed OPTIONS received in requests.
	 */
	Pcpmalformedoptionrate float64 `json:"pcpmalformedoptionrate,omitempty"`
	Pcptotnetfailure       int     `json:"pcptotnetfailure,omitempty"`
	/**
	* Total Network Failures.
	 */
	Pcpnetfailurerate float64 `json:"pcpnetfailurerate,omitempty"`
	Pcptotnoresources int     `json:"pcptotnoresources,omitempty"`
	/**
	* no resources
	 */
	Pcpnoresourcesrate        float64 `json:"pcpnoresourcesrate,omitempty"`
	Pcptotunsupportedprotocol int     `json:"pcptotunsupportedprotocol,omitempty"`
	/**
	* Total Unsupported Protocols requests.
	 */
	Pcpunsupportedprotocolrate float64 `json:"pcpunsupportedprotocolrate,omitempty"`
	Pcptotuserexqouta          int     `json:"pcptotuserexqouta,omitempty"`
	/**
	* Total user ex quota.
	 */
	Pcpuserexqoutarate float64 `json:"pcpuserexqoutarate,omitempty"`
	Pcptotnoexternal   int     `json:"pcptotnoexternal,omitempty"`
	/**
	* Total responses with opcode can not provide external.
	 */
	Pcpnoexternalrate  float64 `json:"pcpnoexternalrate,omitempty"`
	Pcptotaddrmismatch int     `json:"pcptotaddrmismatch,omitempty"`
	/**
	* Total address mismatch.
	 */
	Pcpaddrmismatchrate float64 `json:"pcpaddrmismatchrate,omitempty"`
	Pcptotexcesspeers   int     `json:"pcptotexcesspeers,omitempty"`
	/**
	* Total responses with opcode excessive remote peers.
	 */
	Pcpexcesspeersrate float64 `json:"pcpexcesspeersrate,omitempty"`
}
