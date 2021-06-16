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

package sc


type Scstats struct {
	/**
	* Clear the statsistics / counters
	*/
	Clearstats string `json:"clearstats,omitempty"`
	Sctotcondtriggered uint64 `json:"sctotcondtriggered,omitempty"`
	/**
	* Number of times that SureConnect conditions were triggered.
	*/
	Sccondtriggeredrate float64 `json:"sccondtriggeredrate,omitempty"`
	Scpolicyurlhits uint64 `json:"scpolicyurlhits,omitempty"`
	/**
	* Total number of incoming requests that matched configured sureconnect policies.
	*/
	Scpolicyurlhitsrate float64 `json:"scpolicyurlhitsrate,omitempty"`
	Scpopups uint64 `json:"scpopups,omitempty"`
	/**
	* Total number of in-memory java script  served which throws the pop-up window.
	*/
	Scpopupsrate float64 `json:"scpopupsrate,omitempty"`
	Sctotreissuedrequests uint64 `json:"sctotreissuedrequests,omitempty"`
	/**
	* Total number of reissued SureConnect requests.
	*/
	Screissuedrequestsrate float64 `json:"screissuedrequestsrate,omitempty"`
	Scsessionreqs uint64 `json:"scsessionreqs,omitempty"`
	/**
	* Total number of requests that were  handled in a single SureConnect session.
	*/
	Scsessionreqsrate float64 `json:"scsessionreqsrate,omitempty"`
	Scaltconturls uint64 `json:"scaltconturls,omitempty"`
	/**
	* Total number of alternate content served which throws the pop-up window.
	*/
	Scaltconturlsrate float64 `json:"scaltconturlsrate,omitempty"`
	Scpostreqs uint64 `json:"scpostreqs,omitempty"`
	/**
	* Total number of   HTTP POST requests that triggered SureConnect feature.
	*/
	Scpostreqsrate float64 `json:"scpostreqsrate,omitempty"`
	Scresetstats uint64 `json:"scresetstats,omitempty"`
	/**
	* Toal number of times that SureConnect statistics were reset.
	*/
	Scresetstatsrate float64 `json:"scresetstatsrate,omitempty"`
	Scunsupbrow uint64 `json:"scunsupbrow,omitempty"`
	/**
	* Total number of requests that came from all unsupported browsers.
	*/
	Scunsupbrowrate float64 `json:"scunsupbrowrate,omitempty"`
	Scfaultycookies uint64 `json:"scfaultycookies,omitempty"`
	/**
	* Total number of corrupted SureConnect cookies.
	*/
	Scfaultycookiesrate float64 `json:"scfaultycookiesrate,omitempty"`

}
