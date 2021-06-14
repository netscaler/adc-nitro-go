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
	Aaaauthsuccess uint64 `json:"aaaauthsuccess,omitempty"`
	/**
	* Count of authentication successes.
	*/
	Aaaauthsuccessrate int32 `json:"aaaauthsuccessrate,omitempty"`
	Aaaauthfail uint64 `json:"aaaauthfail,omitempty"`
	/**
	* Count of authentication failures.
	*/
	Aaaauthfailrate int32 `json:"aaaauthfailrate,omitempty"`
	Aaaauthonlyhttpsuccess uint64 `json:"aaaauthonlyhttpsuccess,omitempty"`
	/**
	* Count of HTTP connections that succeeded authorization.
	*/
	Aaaauthonlyhttpsuccessrate int32 `json:"aaaauthonlyhttpsuccessrate,omitempty"`
	Aaaauthonlyhttpfail uint64 `json:"aaaauthonlyhttpfail,omitempty"`
	/**
	* Count of HTTP connections that failed authorization.
	*/
	Aaaauthonlyhttpfailrate int32 `json:"aaaauthonlyhttpfailrate,omitempty"`
	Aaaauthnonhttpsuccess uint64 `json:"aaaauthnonhttpsuccess,omitempty"`
	/**
	* Count of non HTTP connections that succeeded authorization.
	*/
	Aaaauthnonhttpsuccessrate int32 `json:"aaaauthnonhttpsuccessrate,omitempty"`
	Aaaauthnonhttpfail uint64 `json:"aaaauthnonhttpfail,omitempty"`
	/**
	* Count of non HTTP connections that failed authorization.
	*/
	Aaaauthnonhttpfailrate int32 `json:"aaaauthnonhttpfailrate,omitempty"`
	Aaacursessions uint64 `json:"aaacursessions,omitempty"`
	/**
	* Count of current SmartAccess AAA sessions.
	*/
	Aaacursessionsrate int32 `json:"aaacursessionsrate,omitempty"`
	Aaatotsessions uint64 `json:"aaatotsessions,omitempty"`
	/**
	* Count of all SmartAccess AAA sessions.
	*/
	Aaasessionsrate int32 `json:"aaasessionsrate,omitempty"`
	Aaatotsessiontimeout uint64 `json:"aaatotsessiontimeout,omitempty"`
	/**
	* Count of AAA sessions that have timed out.
	*/
	Aaasessiontimeoutrate int32 `json:"aaasessiontimeoutrate,omitempty"`
	Aaacuricasessions uint64 `json:"aaacuricasessions,omitempty"`
	/**
	* Count of current Basic ICA only sessions.
	*/
	Aaacuricasessionsrate int32 `json:"aaacuricasessionsrate,omitempty"`
	Aaacuricaonlyconn uint64 `json:"aaacuricaonlyconn,omitempty"`
	/**
	* Count of current Basic ICA only connections.
	*/
	Aaacuricaonlyconnrate int32 `json:"aaacuricaonlyconnrate,omitempty"`
	Aaacuricaconn uint64 `json:"aaacuricaconn,omitempty"`
	/**
	* Count of current SmartAccess ICA connections.
	*/
	Aaacuricaconnrate int32 `json:"aaacuricaconnrate,omitempty"`
	Aaacurtmsessions uint64 `json:"aaacurtmsessions,omitempty"`
	/**
	* Count of current AAATM sessions.
	*/
	Aaacurtmsessionsrate int32 `json:"aaacurtmsessionsrate,omitempty"`
	Aaatottmsessions uint64 `json:"aaatottmsessions,omitempty"`
	/**
	* Count of all AAATM sessions.
	*/
	Aaatmsessionsrate int32 `json:"aaatmsessionsrate,omitempty"`

}
