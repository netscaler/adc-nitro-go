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

package ssl

/**
* Statistics for SSL virtual server resource.
*/

type Sslvserverstats struct {
	/**
	* Name of the SSL virtual server for which to show detailed statistics
	*/
	Vservername string `json:"vservername,omitempty"`
	/**
	* Clear the statsistics / counters
	*/
	Clearstats string `json:"clearstats,omitempty"`
	Vslbhealth int `json:"vslbhealth,omitempty"`
	Primaryipaddress string `json:"primaryipaddress,omitempty"`
	Primaryport int `json:"primaryport,omitempty"`
	Type string `json:"type,omitempty"`
	State string `json:"state,omitempty"`
	Actsvcs int `json:"actsvcs,omitempty"`
	Ssltotclientauthsuccess int `json:"ssltotclientauthsuccess,omitempty"`
	/**
	* Number of successful client authentication on this vserver
	*/
	Sslclientauthsuccessrate float64 `json:"sslclientauthsuccessrate,omitempty"`
	Ssltotclientauthfailure int `json:"ssltotclientauthfailure,omitempty"`
	/**
	* Number of failure client authentication on this vserver
	*/
	Sslclientauthfailurerate float64 `json:"sslclientauthfailurerate,omitempty"`
	Sslctxtotencbytes int `json:"sslctxtotencbytes,omitempty"`
	/**
	* Number of encrypted bytes per SSL vserver
	*/
	Sslctxencbytesrate float64 `json:"sslctxencbytesrate,omitempty"`
	Sslctxtotdecbytes int `json:"sslctxtotdecbytes,omitempty"`
	/**
	* Number of decrypted bytes per SSL vserver
	*/
	Sslctxdecbytesrate float64 `json:"sslctxdecbytesrate,omitempty"`
	Sslctxtothwencbytes int `json:"sslctxtothwencbytes,omitempty"`
	/**
	* Number of hardware encrypted bytes per SSL vserver
	*/
	Sslctxhwencbytesrate float64 `json:"sslctxhwencbytesrate,omitempty"`
	Sslctxtothwdecbytes int `json:"sslctxtothwdec_bytes,omitempty"`
	/**
	* Number of hw decrypted bytes per SSL vserver
	*/
	Sslctxhwdecbytesrate float64 `json:"sslctxhwdec_bytesrate,omitempty"`
	Sslctxtotsessionnew int `json:"sslctxtotsessionnew,omitempty"`
	/**
	* Number of new sessions created
	*/
	Sslctxsessionnewrate float64 `json:"sslctxsessionnewrate,omitempty"`
	Sslctxtotsessionhits int `json:"sslctxtotsessionhits,omitempty"`
	/**
	* Number of session hits
	*/
	Sslctxsessionhitsrate float64 `json:"sslctxsessionhitsrate,omitempty"`

}
