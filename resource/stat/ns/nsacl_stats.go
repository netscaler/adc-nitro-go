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
	Acltotpktsbridged int `json:"acltotpktsbridged,omitempty"`
	/**
	* Packets matching a bridge ACL, which is in transparent mode and bypasses service processing.
	*/
	Aclpktsbridgedrate float64 `json:"aclpktsbridgedrate,omitempty"`
	Acltotpktsdenied int `json:"acltotpktsdenied,omitempty"`
	/**
	* Packets dropped because they match ACLs with processing mode set to DENY.
	*/
	Aclpktsdeniedrate float64 `json:"aclpktsdeniedrate,omitempty"`
	Acltotpktsallowed int `json:"acltotpktsallowed,omitempty"`
	/**
	* Packets matching ACLs with processing mode set to ALLOW. Citrix ADC processes these packets.
	*/
	Aclpktsallowedrate float64 `json:"aclpktsallowedrate,omitempty"`
	Acltotpktsnat int `json:"acltotpktsnat,omitempty"`
	/**
	* Packets matching a NAT ACL, resulting in a NAT session.
	*/
	Aclpktsnatrate float64 `json:"aclpktsnatrate,omitempty"`
	Acltothits int `json:"acltothits,omitempty"`
	/**
	* Packets matching an ACL.
	*/
	Aclhitsrate float64 `json:"aclhitsrate,omitempty"`
	Acltotmisses int `json:"acltotmisses,omitempty"`
	/**
	* Packets not matching any ACL.
	*/
	Aclmissesrate float64 `json:"aclmissesrate,omitempty"`
	Acltotcount int `json:"acltotcount,omitempty"`
	Effectiveacltotcount int `json:"effectiveacltotcount,omitempty"`
	Dfdacltothits int `json:"dfdacltothits,omitempty"`
	/**
	* Packets matching an dfd ACL.
	*/
	Dfdaclhitsrate float64 `json:"dfdaclhitsrate,omitempty"`
	Dfdacltotmisses int `json:"dfdacltotmisses,omitempty"`
	/**
	* Packets not matching any DFD ACL.
	*/
	Dfdaclmissesrate float64 `json:"dfdaclmissesrate,omitempty"`
	Dfdacltotcount int `json:"dfdacltotcount,omitempty"`
	Aclperhits int `json:"aclperhits,omitempty"`
	/**
	* Number of times the acl was hit
	*/
	Aclperhitsrate float64 `json:"aclperhitsrate,omitempty"`

}
