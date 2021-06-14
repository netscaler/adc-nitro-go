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
* Statistics for ACL entry resource.
*/

type Nsaclstats struct {
	/**
	* Name of the extended ACL rule whose statistics you want the Citrix ADC to display.
	*/
	Aclname string `json:"aclname,omitempty"`
	/**
	* Clear the statsistics / counters
	*/
	Clearstats string `json:"clearstats,omitempty"`
	Acltotpktsbridged uint64 `json:"acltotpktsbridged,omitempty"`
	/**
	* Packets matching a bridge ACL, which is in transparent mode and bypasses service processing.
	*/
	Aclpktsbridgedrate int32 `json:"aclpktsbridgedrate,omitempty"`
	Acltotpktsdenied uint64 `json:"acltotpktsdenied,omitempty"`
	/**
	* Packets dropped because they match ACLs with processing mode set to DENY.
	*/
	Aclpktsdeniedrate int32 `json:"aclpktsdeniedrate,omitempty"`
	Acltotpktsallowed uint64 `json:"acltotpktsallowed,omitempty"`
	/**
	* Packets matching ACLs with processing mode set to ALLOW. Citrix ADC processes these packets.
	*/
	Aclpktsallowedrate int32 `json:"aclpktsallowedrate,omitempty"`
	Acltotpktsnat uint64 `json:"acltotpktsnat,omitempty"`
	/**
	* Packets matching a NAT ACL, resulting in a NAT session.
	*/
	Aclpktsnatrate int32 `json:"aclpktsnatrate,omitempty"`
	Acltothits uint64 `json:"acltothits,omitempty"`
	/**
	* Packets matching an ACL.
	*/
	Aclhitsrate int32 `json:"aclhitsrate,omitempty"`
	Acltotmisses uint64 `json:"acltotmisses,omitempty"`
	/**
	* Packets not matching any ACL.
	*/
	Aclmissesrate int32 `json:"aclmissesrate,omitempty"`
	Acltotcount uint64 `json:"acltotcount,omitempty"`
	Effectiveacltotcount uint64 `json:"effectiveacltotcount,omitempty"`
	Dfdacltothits uint64 `json:"dfdacltothits,omitempty"`
	/**
	* Packets matching an dfd ACL.
	*/
	Dfdaclhitsrate int32 `json:"dfdaclhitsrate,omitempty"`
	Dfdacltotmisses uint64 `json:"dfdacltotmisses,omitempty"`
	/**
	* Packets not matching any DFD ACL.
	*/
	Dfdaclmissesrate int32 `json:"dfdaclmissesrate,omitempty"`
	Dfdacltotcount uint64 `json:"dfdacltotcount,omitempty"`
	Aclperhits uint64 `json:"aclperhits,omitempty"`
	/**
	* Number of times the acl was hit
	*/
	Aclperhitsrate int32 `json:"aclperhitsrate,omitempty"`

}
