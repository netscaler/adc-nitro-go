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
* Binding class showing the profile that can be bound to sslcipher.
 */
type Sslcipherprofilebinding struct {
	/**
	* Name of the profile to which cipher is attached.
	 */
	Sslprofile string `json:"sslprofile,omitempty"`
	/**
	* Cipher suite description.
	 */
	Description string `json:"description,omitempty"`
	/**
	* Name of the user-defined cipher group.
	 */
	Ciphergroupname string `json:"ciphergroupname,omitempty"`
	/**
	* The operation that is performed when adding the cipher-suite.
		Possible cipher operations are:
		ADD - Appends the given cipher-suite to the existing one configured for the virtual server.
		REM - Removes the given cipher-suite from the existing one configured for the virtual server.
		ORD - Overrides the current configured cipher-suite for the virtual server with the given cipher-suite.
	*/
	Cipheroperation string `json:"cipheroperation,omitempty"`
	/**
	* A cipher-suite can consist of an individual cipher name, the system predefined cipher-alias name, or user defined cipher-group name.
	 */
	Ciphgrpals string `json:"ciphgrpals,omitempty"`
	/**
	* Priority of the cipher to be added
	 */
	Cipherpriority uint32 `json:"cipherpriority,omitempty"`
}
