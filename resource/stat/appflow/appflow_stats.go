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

package appflow


type Appflowstats struct {
	/**
	* Clear the statsistics / counters
	*/
	Clearstats string `json:"clearstats,omitempty"`
	Appflowtotaltxoctets uint64 `json:"appflowtotaltxoctets,omitempty"`
	/**
	* The total number of AppFlow (IPFIX) octets that the Citrix ADC has transmitted.
	*/
	Appflowtxoctetsrate float64 `json:"appflowtxoctetsrate,omitempty"`
	Appflowtotalflows uint64 `json:"appflowtotalflows,omitempty"`
	/**
	* The total number of AppFlow (IPFIX) flows that the Citrix ADC has transmitted.
	*/
	Appflowflowsrate float64 `json:"appflowflowsrate,omitempty"`
	Appflowtotaltxmessagess uint64 `json:"appflowtotaltxmessagess,omitempty"`
	/**
	* The total number of AppFlow (IPFIX) messages that the Citrix ADC has transmitted.
	*/
	Appflowtxmessagessrate float64 `json:"appflowtxmessagessrate,omitempty"`
	Appflowtotalignoctets uint64 `json:"appflowtotalignoctets,omitempty"`
	/**
	* The total number of octets that the Citrix ADC has ignored for AppFlow (IPFIX).
	*/
	Appflowignoctetsrate float64 `json:"appflowignoctetsrate,omitempty"`
	Appflowtotalignpacketss uint64 `json:"appflowtotalignpacketss,omitempty"`
	/**
	* The total number of packets that the Citrix ADC has ignored for AppFlow (IPFIX).
	*/
	Appflowignpacketssrate float64 `json:"appflowignpacketssrate,omitempty"`
	Appflowtotalnotxoctets uint64 `json:"appflowtotalnotxoctets,omitempty"`
	/**
	* The total number of AppFlow (IPFIX) octets that the Citrix ADC has not transmitted.
	*/
	Appflownotxoctetsrate float64 `json:"appflownotxoctetsrate,omitempty"`
	Appflowtotalnotxflows uint64 `json:"appflowtotalnotxflows,omitempty"`
	/**
	* The total number of AppFlow (IPFIX) flows that the Citrix ADC has not transmitted.
	*/
	Appflownotxflowsrate float64 `json:"appflownotxflowsrate,omitempty"`
	Appflowtotalnotxpackets uint64 `json:"appflowtotalnotxpackets,omitempty"`
	/**
	* The total number of AppFlow (IPFIX) packets that the Citrix ADC has not transmitted.
	*/
	Appflownotxpacketsrate float64 `json:"appflownotxpacketsrate,omitempty"`

}
