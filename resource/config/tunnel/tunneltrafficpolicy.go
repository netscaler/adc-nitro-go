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

package tunnel

/**
* Configuration for tunnel policy resource.
 */
type Tunneltrafficpolicy struct {
	/**
	* Name for the tunnel traffic policy.
		Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the policy is created.
		The following requirement applies only to the Citrix ADC CLI:
		If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my policy" or 'my policy)'.
	*/
	Name string `json:"name,omitempty"`
	/**
	* Expression, against which traffic is evaluated.
		The following requirements apply only to the Citrix ADC CLI:
		*  If the expression includes blank spaces, the entire expression must be enclosed in double quotation marks.
		*  If the expression itself includes double quotation marks, you must escape the quotations by using the \ character.
		*  Alternatively, you can use single quotation marks to enclose the rule, in which case you do not have to escape the double quotation marks.
	*/
	Rule string `json:"rule,omitempty"`
	/**
	* Name of the built-in compression action to associate with the policy.
	 */
	Action string `json:"action,omitempty"`
	/**
	* Any comments to preserve information about this policy.
	 */
	Comment string `json:"comment,omitempty"`
	/**
	* Name of the messagelog action to use for requests that match this policy.
	 */
	Logaction string `json:"logaction,omitempty"`
	/**
	* New name for the tunnel traffic policy. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), e
		quals (=), and hyphen (-) characters.
		Choose a name that reflects the function that the policy performs.
		The following requirement applies only to the Citrix ADC CLI:
		If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my tunnel policy" or 'my tunnel policy').
	*/
	Newname string `json:"newname,omitempty"`

	//------- Read only Parameter ---------;

	Expressiontype     string `json:"expressiontype,omitempty"`
	Hits               string `json:"hits,omitempty"`
	Undefhits          string `json:"undefhits,omitempty"`
	Txbytes            string `json:"txbytes,omitempty"`
	Rxbytes            string `json:"rxbytes,omitempty"`
	Clientttlb         string `json:"clientttlb,omitempty"`
	Clienttransactions string `json:"clienttransactions,omitempty"`
	Serverttlb         string `json:"serverttlb,omitempty"`
	Servertransactions string `json:"servertransactions,omitempty"`
	Isdefault          string `json:"isdefault,omitempty"`
	Builtin            string `json:"builtin,omitempty"`
	Feature            string `json:"feature,omitempty"`
}
