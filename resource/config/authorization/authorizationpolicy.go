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

package authorization

/**
* Configuration for authorization policy resource.
 */
type Authorizationpolicy struct {
	/**
	* Name for the new authorization policy.
		Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after the authorization policy is added.
		The following requirement applies only to the Citrix ADC CLI:
		If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my authorization policy" or 'my authorization policy').
	*/
	Name string `json:"name,omitempty"`
	/**
	* Name of the Citrix ADC named rule, or an expression, that the policy uses to perform the authentication.
	 */
	Rule string `json:"rule,omitempty"`
	/**
	* Action to perform if the policy matches: either allow or deny the request.
	 */
	Action string `json:"action,omitempty"`
	/**
	* The new name of the author policy.
	 */
	Newname string `json:"newname,omitempty"`

	//------- Read only Parameter ---------;

	Activepolicy   string `json:"activepolicy,omitempty"`
	Expressiontype string `json:"expressiontype,omitempty"`
	Hits           string `json:"hits,omitempty"`
}
