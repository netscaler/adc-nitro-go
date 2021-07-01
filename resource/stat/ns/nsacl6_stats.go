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

package ns

/**
* Statistics for ACL6 entry resource.
*/

type Nsacl6stats struct {
	/**
	* Name of the ACL6 rule whose statistics you want the Citrix ADC to display.
	*/
	Acl6name string `json:"acl6name,omitempty"`
	/**
	* Clear the statsistics / counters
	*/
	Clearstats string `json:"clearstats,omitempty"`
	Acl6totpktsbridged int `json:"acl6totpktsbridged,omitempty"`
	/**
	* Packets matching a bridge IPv6 ACL, which is in transparent mode and bypasses service processing.
	*/
	Acl6pktsbridgedrate float64 `json:"acl6pktsbridgedrate,omitempty"`
	Acl6totpktsdenied int `json:"acl6totpktsdenied,omitempty"`
	/**
	* Packets dropped because they match IPv6 ACLs with processing mode set to DENY.
	*/
	Acl6pktsdeniedrate float64 `json:"acl6pktsdeniedrate,omitempty"`
	Acl6totpktsallowed int `json:"acl6totpktsallowed,omitempty"`
	/**
	* Packets matching IPv6 ACLs with processing mode set to ALLOW. Citrix ADC processes these packets.
	*/
	Acl6pktsallowedrate float64 `json:"acl6pktsallowedrate,omitempty"`
	Acl6totpktsnat int `json:"acl6totpktsnat,omitempty"`
	/**
	* Packets matching a NAT ACL6, resulting in a NAT session.
	*/
	Acl6pktsnatrate float64 `json:"acl6pktsnatrate,omitempty"`
	Acl6tothits int `json:"acl6tothits,omitempty"`
	/**
	* Packets matching an IPv6 ACL.
	*/
	Acl6hitsrate float64 `json:"acl6hitsrate,omitempty"`
	Acl6totmisses int `json:"acl6totmisses,omitempty"`
	/**
	* Packets not matching any IPv6 ACL.
	*/
	Acl6missesrate float64 `json:"acl6missesrate,omitempty"`
	Acl6totpktsnat64 int `json:"acl6totpktsnat64,omitempty"`
	/**
	* Packets matching a NAT64 ACL6, resulting in a NAT64 translation.
	*/
	Acl6pktsnat64rate float64 `json:"acl6pktsnat64rate,omitempty"`
	Acl6totcount int `json:"acl6totcount,omitempty"`
	Dfdacl6totcount int `json:"dfdacl6totcount,omitempty"`
	Dfdacl6tothits int `json:"dfdacl6tothits,omitempty"`
	/**
	* Packets matching an dfd ACL6.
	*/
	Dfdacl6hitsrate float64 `json:"dfdacl6hitsrate,omitempty"`
	Dfdacl6totmisses int `json:"dfdacl6totmisses,omitempty"`
	/**
	* Packets not matching any DFD ACL6.
	*/
	Dfdacl6missesrate float64 `json:"dfdacl6missesrate,omitempty"`
	Acl6perhits int `json:"acl6perhits,omitempty"`
	/**
	* Number of times the acl6 was hit
	*/
	Acl6perhitsrate float64 `json:"acl6perhitsrate,omitempty"`

}
