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


type Appqoestats struct {
	/**
	* Clear the statsistics / counters
	*/
	Clearstats string `json:"clearstats,omitempty"`
	Totinmemrsp uint64 `json:"totinmemrsp,omitempty"`
	/**
	* Total in-memory responses sent from NS
	*/
	Inmemrsprate float64 `json:"inmemrsprate,omitempty"`
	Totfaultycookies uint64 `json:"totfaultycookies,omitempty"`
	/**
	* Total faulty cookies received
	*/
	Faultycookiesrate float64 `json:"faultycookiesrate,omitempty"`
	Totvalidcookies uint64 `json:"totvalidcookies,omitempty"`
	/**
	* Total valid cookies received
	*/
	Validcookiesrate float64 `json:"validcookiesrate,omitempty"`
	Tothighprireq uint64 `json:"tothighprireq,omitempty"`
	/**
	* Total Requests served from higher priority queue
	*/
	Highprireqrate float64 `json:"highprireqrate,omitempty"`
	Totmediumprireq uint64 `json:"totmediumprireq,omitempty"`
	/**
	* Total Requests served from medium priority queue
	*/
	Mediumprireqrate float64 `json:"mediumprireqrate,omitempty"`
	Totlowprireq uint64 `json:"totlowprireq,omitempty"`
	/**
	* Total Requests served from low priority queue
	*/
	Lowprireqrate float64 `json:"lowprireqrate,omitempty"`
	Totlowestprireq uint64 `json:"totlowestprireq,omitempty"`
	/**
	* Total Requests served from surge priority queue
	*/
	Lowestprireqrate float64 `json:"lowestprireqrate,omitempty"`
	Totaltsvrsubfailed uint64 `json:"totaltsvrsubfailed,omitempty"`
	/**
	* Total number of times alternate server substitution failed
	*/
	Tsvrsubfailedrate float64 `json:"tsvrsubfailedrate,omitempty"`
	Totdostrig uint64 `json:"totdostrig,omitempty"`
	/**
	* Total number of times HDOS condition triggered
	*/
	Dostrigrate float64 `json:"dostrigrate,omitempty"`
	Totdosqvalidcookies uint64 `json:"totdosqvalidcookies,omitempty"`
	/**
	* Total DOSQ valid cookies received
	*/
	Dosqvalidcookiesrate float64 `json:"dosqvalidcookiesrate,omitempty"`
	Totdoshvalidcookies uint64 `json:"totdoshvalidcookies,omitempty"`
	/**
	* Total DOSH valid cookies received
	*/
	Doshvalidcookiesrate float64 `json:"doshvalidcookiesrate,omitempty"`
	Totsidvalidcookies uint64 `json:"totsidvalidcookies,omitempty"`
	/**
	* Total SID valid cookies received
	*/
	Sidvalidcookiesrate float64 `json:"sidvalidcookiesrate,omitempty"`
	Totonhvalidcookies uint64 `json:"totonhvalidcookies,omitempty"`
	/**
	* Total ONH valid cookies received
	*/
	Onhvalidcookiesrate float64 `json:"onhvalidcookiesrate,omitempty"`
	Totpriqvalidcookies uint64 `json:"totpriqvalidcookies,omitempty"`
	/**
	* Total PRIQ valid cookies received
	*/
	Priqvalidcookiesrate float64 `json:"priqvalidcookiesrate,omitempty"`
	Totdosqfaultycookies uint64 `json:"totdosqfaultycookies,omitempty"`
	/**
	* Total DOSQ faulty cookies received
	*/
	Dosqfaultycookiesrate float64 `json:"dosqfaultycookiesrate,omitempty"`
	Totdoshfaultycookies uint64 `json:"totdoshfaultycookies,omitempty"`
	/**
	* Total DOSH faulty cookies received
	*/
	Doshfaultycookiesrate float64 `json:"doshfaultycookiesrate,omitempty"`
	Totsidfaultycookies uint64 `json:"totsidfaultycookies,omitempty"`
	/**
	* Total SID faulty cookies received
	*/
	Sidfaultycookiesrate float64 `json:"sidfaultycookiesrate,omitempty"`
	Totonhfaultycookies uint64 `json:"totonhfaultycookies,omitempty"`
	/**
	* Total ONH faulty cookies received
	*/
	Onhfaultycookiesrate float64 `json:"onhfaultycookiesrate,omitempty"`
	Totpriqfaultycookies uint64 `json:"totpriqfaultycookies,omitempty"`
	/**
	* Total PRIQ faulty cookies received
	*/
	Priqfaultycookiesrate float64 `json:"priqfaultycookiesrate,omitempty"`
	Totpriembedlinks uint64 `json:"totpriembedlinks,omitempty"`
	/**
	* Total requests for valid embedded links
	*/
	Priembedlinksrate float64 `json:"priembedlinksrate,omitempty"`
	Totsessreq uint64 `json:"totsessreq,omitempty"`
	/**
	* Total valid SIDQ requests within session
	*/
	Sessreqrate float64 `json:"sessreqrate,omitempty"`
	Totaltcntreq uint64 `json:"totaltcntreq,omitempty"`
	/**
	* Total requests for alternate contents
	*/
	Tcntreqrate float64 `json:"tcntreqrate,omitempty"`
	Totgetinmemrsp uint64 `json:"totgetinmemrsp,omitempty"`
	/**
	* Total in-memory GET responses sent from NS
	*/
	Getinmemrsprate float64 `json:"getinmemrsprate,omitempty"`
	Totpostinmemrsp uint64 `json:"totpostinmemrsp,omitempty"`
	/**
	* Total in-memory POST responses sent from NS
	*/
	Postinmemrsprate float64 `json:"postinmemrsprate,omitempty"`
	Totpostinmemrspbytes uint64 `json:"totpostinmemrspbytes,omitempty"`
	/**
	* Total in-memory response bytes sent from NS
	*/
	Postinmemrspbytesrate float64 `json:"postinmemrspbytesrate,omitempty"`

}
