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
	Acl6totpktsbridged uint64 `json:"acl6totpktsbridged,omitempty"`
	/**
	* Packets matching a bridge IPv6 ACL, which is in transparent mode and bypasses service processing.
	*/
	Acl6pktsbridgedrate int32 `json:"acl6pktsbridgedrate,omitempty"`
	Acl6totpktsdenied uint64 `json:"acl6totpktsdenied,omitempty"`
	/**
	* Packets dropped because they match IPv6 ACLs with processing mode set to DENY.
	*/
	Acl6pktsdeniedrate int32 `json:"acl6pktsdeniedrate,omitempty"`
	Acl6totpktsallowed uint64 `json:"acl6totpktsallowed,omitempty"`
	/**
	* Packets matching IPv6 ACLs with processing mode set to ALLOW. Citrix ADC processes these packets.
	*/
	Acl6pktsallowedrate int32 `json:"acl6pktsallowedrate,omitempty"`
	Acl6totpktsnat uint64 `json:"acl6totpktsnat,omitempty"`
	/**
	* Packets matching a NAT ACL6, resulting in a NAT session.
	*/
	Acl6pktsnatrate int32 `json:"acl6pktsnatrate,omitempty"`
	Acl6tothits uint64 `json:"acl6tothits,omitempty"`
	/**
	* Packets matching an IPv6 ACL.
	*/
	Acl6hitsrate int32 `json:"acl6hitsrate,omitempty"`
	Acl6totmisses uint64 `json:"acl6totmisses,omitempty"`
	/**
	* Packets not matching any IPv6 ACL.
	*/
	Acl6missesrate int32 `json:"acl6missesrate,omitempty"`
	Acl6totpktsnat64 uint64 `json:"acl6totpktsnat64,omitempty"`
	/**
	* Packets matching a NAT64 ACL6, resulting in a NAT64 translation.
	*/
	Acl6pktsnat64rate int32 `json:"acl6pktsnat64rate,omitempty"`
	Acl6totcount uint64 `json:"acl6totcount,omitempty"`
	Dfdacl6totcount uint64 `json:"dfdacl6totcount,omitempty"`
	Dfdacl6tothits uint64 `json:"dfdacl6tothits,omitempty"`
	/**
	* Packets matching an dfd ACL6.
	*/
	Dfdacl6hitsrate int32 `json:"dfdacl6hitsrate,omitempty"`
	Dfdacl6totmisses uint64 `json:"dfdacl6totmisses,omitempty"`
	/**
	* Packets not matching any DFD ACL6.
	*/
	Dfdacl6missesrate int32 `json:"dfdacl6missesrate,omitempty"`
	Acl6perhits uint64 `json:"acl6perhits,omitempty"`
	/**
	* Number of times the acl6 was hit
	*/
	Acl6perhitsrate int32 `json:"acl6perhitsrate,omitempty"`

}
