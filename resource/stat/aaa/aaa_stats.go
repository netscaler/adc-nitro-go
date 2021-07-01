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

package aaa


type Aaastats struct {
	/**
	* Clear the statsistics / counters
	*/
	Clearstats string `json:"clearstats,omitempty"`
	Aaaauthsuccess int `json:"aaaauthsuccess,omitempty"`
	/**
	* Count of authentication successes.
	*/
	Aaaauthsuccessrate float64 `json:"aaaauthsuccessrate,omitempty"`
	Aaaauthfail int `json:"aaaauthfail,omitempty"`
	/**
	* Count of authentication failures.
	*/
	Aaaauthfailrate float64 `json:"aaaauthfailrate,omitempty"`
	Aaaauthonlyhttpsuccess int `json:"aaaauthonlyhttpsuccess,omitempty"`
	/**
	* Count of HTTP connections that succeeded authorization.
	*/
	Aaaauthonlyhttpsuccessrate float64 `json:"aaaauthonlyhttpsuccessrate,omitempty"`
	Aaaauthonlyhttpfail int `json:"aaaauthonlyhttpfail,omitempty"`
	/**
	* Count of HTTP connections that failed authorization.
	*/
	Aaaauthonlyhttpfailrate float64 `json:"aaaauthonlyhttpfailrate,omitempty"`
	Aaaauthnonhttpsuccess int `json:"aaaauthnonhttpsuccess,omitempty"`
	/**
	* Count of non HTTP connections that succeeded authorization.
	*/
	Aaaauthnonhttpsuccessrate float64 `json:"aaaauthnonhttpsuccessrate,omitempty"`
	Aaaauthnonhttpfail int `json:"aaaauthnonhttpfail,omitempty"`
	/**
	* Count of non HTTP connections that failed authorization.
	*/
	Aaaauthnonhttpfailrate float64 `json:"aaaauthnonhttpfailrate,omitempty"`
	Aaacursessions int `json:"aaacursessions,omitempty"`
	/**
	* Count of current SmartAccess AAA sessions.
	*/
	Aaacursessionsrate float64 `json:"aaacursessionsrate,omitempty"`
	Aaatotsessions int `json:"aaatotsessions,omitempty"`
	/**
	* Count of all SmartAccess AAA sessions.
	*/
	Aaasessionsrate float64 `json:"aaasessionsrate,omitempty"`
	Aaatotsessiontimeout int `json:"aaatotsessiontimeout,omitempty"`
	/**
	* Count of AAA sessions that have timed out.
	*/
	Aaasessiontimeoutrate float64 `json:"aaasessiontimeoutrate,omitempty"`
	Aaacuricasessions int `json:"aaacuricasessions,omitempty"`
	/**
	* Count of current Basic ICA only sessions.
	*/
	Aaacuricasessionsrate float64 `json:"aaacuricasessionsrate,omitempty"`
	Aaacuricaonlyconn int `json:"aaacuricaonlyconn,omitempty"`
	/**
	* Count of current Basic ICA only connections.
	*/
	Aaacuricaonlyconnrate float64 `json:"aaacuricaonlyconnrate,omitempty"`
	Aaacuricaconn int `json:"aaacuricaconn,omitempty"`
	/**
	* Count of current SmartAccess ICA connections.
	*/
	Aaacuricaconnrate float64 `json:"aaacuricaconnrate,omitempty"`
	Aaacurtmsessions int `json:"aaacurtmsessions,omitempty"`
	/**
	* Count of current AAATM sessions.
	*/
	Aaacurtmsessionsrate float64 `json:"aaacurtmsessionsrate,omitempty"`
	Aaatottmsessions int `json:"aaatottmsessions,omitempty"`
	/**
	* Count of all AAATM sessions.
	*/
	Aaatmsessionsrate float64 `json:"aaatmsessionsrate,omitempty"`

}
