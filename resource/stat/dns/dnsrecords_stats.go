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

package dns


type Dnsrecordsstats struct {
	/**
	* Display statistics for the specified DNS record or query type or, if a record or query type is not specified, statistics for all record types supported on the Citrix ADC.
	*/
	Dnsrecordtype string `json:"dnsrecordtype,omitempty"`
	/**
	* Clear the statsistics / counters
	*/
	Clearstats string `json:"clearstats,omitempty"`
	Dnstotentries uint64 `json:"dnstotentries,omitempty"`
	Dnstotupdates uint64 `json:"dnstotupdates,omitempty"`
	Dnstotresponses uint64 `json:"dnstotresponses,omitempty"`
	Dnstotrequests uint64 `json:"dnstotrequests,omitempty"`
	Dnscurentries uint32 `json:"dnscurentries,omitempty"`
	Dnstotqueryforexpiredrecord uint64 `json:"dnstotqueryforexpiredrecord,omitempty"`
	Dnstoterrlimits uint64 `json:"dnstoterrlimits,omitempty"`
	Dnstoterrrespform uint64 `json:"dnstoterrrespform,omitempty"`
	Dnstoterraliasex uint64 `json:"dnstoterraliasex,omitempty"`
	Dnstoterrnodomains uint64 `json:"dnstoterrnodomains,omitempty"`
	Dnscurrecords uint32 `json:"dnscurrecords,omitempty"`

}
