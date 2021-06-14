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

package contentinspection


type Contentinspectionstats struct {
	/**
	* Clear the statsistics / counters
	*/
	Clearstats string `json:"clearstats,omitempty"`
	Inlinerequestssent uint64 `json:"inlinerequestssent,omitempty"`
	Inlineresponsessent uint64 `json:"inlineresponsessent,omitempty"`
	Inlinereqbytessent uint64 `json:"inlinereqbytessent,omitempty"`
	Inlinereqbytesrecv uint64 `json:"inlinereqbytesrecv,omitempty"`
	Inlinerespbytessent uint64 `json:"inlinerespbytessent,omitempty"`
	Inlinerespbytesrecv uint64 `json:"inlinerespbytesrecv,omitempty"`
	Inlineserverdownreset uint64 `json:"inlineserverdownreset,omitempty"`
	Inlineserverdowndrop uint64 `json:"inlineserverdowndrop,omitempty"`
	Inlineserverdownbypass uint64 `json:"inlineserverdownbypass,omitempty"`
	Inlinegeneratedresponses uint64 `json:"inlinegeneratedresponses,omitempty"`
	Mirrorrequestssent uint64 `json:"mirrorrequestssent,omitempty"`
	Mirrorresponsessent uint64 `json:"mirrorresponsessent,omitempty"`
	Mirrorreqbytessent uint64 `json:"mirrorreqbytessent,omitempty"`
	Mirrorrespbytessent uint64 `json:"mirrorrespbytessent,omitempty"`
	Mirrorserverdownreset uint64 `json:"mirrorserverdownreset,omitempty"`
	Mirrorserverdowndrop uint64 `json:"mirrorserverdowndrop,omitempty"`
	Mirrorserverdownbypass uint64 `json:"mirrorserverdownbypass,omitempty"`
	Icapreqmodrequests uint64 `json:"icapreqmodrequests,omitempty"`
	Icaprespmodrequests uint64 `json:"icaprespmodrequests,omitempty"`
	Icappreviewenabledrequests uint64 `json:"icappreviewenabledrequests,omitempty"`
	Icap204enabledrequests uint64 `json:"icap204enabledrequests,omitempty"`
	Icap100contrecv uint64 `json:"icap100contrecv,omitempty"`
	Icap204nocontentrecv uint64 `json:"icap204nocontentrecv,omitempty"`
	Icapadaptiverequests uint64 `json:"icapadaptiverequests,omitempty"`
	Icapadaptiveresponses uint64 `json:"icapadaptiveresponses,omitempty"`
	Icapcalloutinitiated uint64 `json:"icapcalloutinitiated,omitempty"`
	Icapcalloutcompleted uint64 `json:"icapcalloutcompleted,omitempty"`
	Icaperrorshandled uint64 `json:"icaperrorshandled,omitempty"`
	Icapserverdownreset uint64 `json:"icapserverdownreset,omitempty"`
	Icapserverdowndrop uint64 `json:"icapserverdowndrop,omitempty"`
	Icapserverdownbypass uint64 `json:"icapserverdownbypass,omitempty"`

}
