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
	Inlinerequestssent int `json:"inlinerequestssent,omitempty"`
	Inlineresponsessent int `json:"inlineresponsessent,omitempty"`
	Inlinereqbytessent int `json:"inlinereqbytessent,omitempty"`
	Inlinereqbytesrecv int `json:"inlinereqbytesrecv,omitempty"`
	Inlinerespbytessent int `json:"inlinerespbytessent,omitempty"`
	Inlinerespbytesrecv int `json:"inlinerespbytesrecv,omitempty"`
	Inlineserverdownreset int `json:"inlineserverdownreset,omitempty"`
	Inlineserverdowndrop int `json:"inlineserverdowndrop,omitempty"`
	Inlineserverdownbypass int `json:"inlineserverdownbypass,omitempty"`
	Inlinegeneratedresponses int `json:"inlinegeneratedresponses,omitempty"`
	Mirrorrequestssent int `json:"mirrorrequestssent,omitempty"`
	Mirrorresponsessent int `json:"mirrorresponsessent,omitempty"`
	Mirrorreqbytessent int `json:"mirrorreqbytessent,omitempty"`
	Mirrorrespbytessent int `json:"mirrorrespbytessent,omitempty"`
	Mirrorserverdownreset int `json:"mirrorserverdownreset,omitempty"`
	Mirrorserverdowndrop int `json:"mirrorserverdowndrop,omitempty"`
	Mirrorserverdownbypass int `json:"mirrorserverdownbypass,omitempty"`
	Icapreqmodrequests int `json:"icapreqmodrequests,omitempty"`
	Icaprespmodrequests int `json:"icaprespmodrequests,omitempty"`
	Icappreviewenabledrequests int `json:"icappreviewenabledrequests,omitempty"`
	Icap204enabledrequests int `json:"icap204enabledrequests,omitempty"`
	Icap100contrecv int `json:"icap100contrecv,omitempty"`
	Icap204nocontentrecv int `json:"icap204nocontentrecv,omitempty"`
	Icapadaptiverequests int `json:"icapadaptiverequests,omitempty"`
	Icapadaptiveresponses int `json:"icapadaptiveresponses,omitempty"`
	Icapcalloutinitiated int `json:"icapcalloutinitiated,omitempty"`
	Icapcalloutcompleted int `json:"icapcalloutcompleted,omitempty"`
	Icaperrorshandled int `json:"icaperrorshandled,omitempty"`
	Icapserverdownreset int `json:"icapserverdownreset,omitempty"`
	Icapserverdowndrop int `json:"icapserverdowndrop,omitempty"`
	Icapserverdownbypass int `json:"icapserverdownbypass,omitempty"`

}
