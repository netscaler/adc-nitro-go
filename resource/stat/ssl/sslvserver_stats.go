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
	Vslbhealth uint32 `json:"vslbhealth,omitempty"`
	Primaryipaddress string `json:"primaryipaddress,omitempty"`
	Primaryport int32 `json:"primaryport,omitempty"`
	Type string `json:"type,omitempty"`
	State string `json:"state,omitempty"`
	Actsvcs uint64 `json:"actsvcs,omitempty"`
	Ssltotclientauthsuccess uint32 `json:"ssltotclientauthsuccess,omitempty"`
	/**
	* Number of successful client authentication on this vserver
	*/
	Sslclientauthsuccessrate int32 `json:"sslclientauthsuccessrate,omitempty"`
	Ssltotclientauthfailure uint32 `json:"ssltotclientauthfailure,omitempty"`
	/**
	* Number of failure client authentication on this vserver
	*/
	Sslclientauthfailurerate int32 `json:"sslclientauthfailurerate,omitempty"`
	Sslctxtotencbytes uint32 `json:"sslctxtotencbytes,omitempty"`
	/**
	* Number of encrypted bytes per SSL vserver
	*/
	Sslctxencbytesrate int32 `json:"sslctxencbytesrate,omitempty"`
	Sslctxtotdecbytes uint32 `json:"sslctxtotdecbytes,omitempty"`
	/**
	* Number of decrypted bytes per SSL vserver
	*/
	Sslctxdecbytesrate int32 `json:"sslctxdecbytesrate,omitempty"`
	Sslctxtothwencbytes uint32 `json:"sslctxtothwencbytes,omitempty"`
	/**
	* Number of hardware encrypted bytes per SSL vserver
	*/
	Sslctxhwencbytesrate int32 `json:"sslctxhwencbytesrate,omitempty"`
	Sslctxtothwdecbytes uint32 `json:"sslctxtothwdec_bytes,omitempty"`
	/**
	* Number of hw decrypted bytes per SSL vserver
	*/
	Sslctxhwdecbytesrate int32 `json:"sslctxhwdec_bytesrate,omitempty"`
	Sslctxtotsessionnew uint32 `json:"sslctxtotsessionnew,omitempty"`
	/**
	* Number of new sessions created
	*/
	Sslctxsessionnewrate int32 `json:"sslctxsessionnewrate,omitempty"`
	Sslctxtotsessionhits uint32 `json:"sslctxtotsessionhits,omitempty"`
	/**
	* Number of session hits
	*/
	Sslctxsessionhitsrate int32 `json:"sslctxsessionhitsrate,omitempty"`

}
