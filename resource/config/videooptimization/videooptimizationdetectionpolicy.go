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

package videooptimization

/**
* Configuration for videooptimization detectionpolicy resource.
 */
type Videooptimizationdetectionpolicy struct {
	/**
	* Name for the videooptimization detection policy. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters.Can be modified, removed or renamed.
	 */
	Name string `json:"name,omitempty"`
	/**
	* Expression that determines which request or response match the video optimization detection policy.
		The following requirements apply only to the Citrix ADC CLI:
		* If the expression includes one or more spaces, enclose the entire expression in double quotation marks.
		* If the expression itself includes double quotation marks, escape the quotations by using the \ character.
		* Alternatively, you can use single quotation marks to enclose the rule, in which case you do not have to escape the double quotation marks.
	*/
	Rule string `json:"rule,omitempty"`
	/**
	* Name of the videooptimization detection action to perform if the request matches this videooptimization detection policy. Built-in actions should be used. These are:
		* DETECT_CLEARTEXT_PD - Cleartext PD is detected and increment related counters.
		* DETECT_CLEARTEXT_ABR - Cleartext ABR is detected and increment related counters.
		* DETECT_ENCRYPTED_ABR - Encrypted ABR is detected and increment related counters.
		* TRIGGER_ENC_ABR_DETECTION - This is potentially encrypted ABR. Internal traffic heuristics algorithms will further process traffic to confirm detection.
		* TRIGGER_CT_ABR_BODY_DETECTION -  This is potentially cleartext ABR. Internal traffic heuristics algorithms will further process traffic to confirm detection.
		* RESET - Reset the client connection by closing it.
		* DROP - Drop the connection without sending a response.
	*/
	Action string `json:"action,omitempty"`
	/**
	* Action to perform if the result of policy evaluation is undefined (UNDEF). An UNDEF event indicates an internal error condition. Only the above built-in actions can be used.
	 */
	Undefaction string `json:"undefaction,omitempty"`
	/**
	* Any type of information about this videooptimization detection policy.
	 */
	Comment string `json:"comment,omitempty"`
	/**
	* Name of the messagelog action to use for requests that match this policy.
	 */
	Logaction string `json:"logaction,omitempty"`
	/**
	* New name for the videooptimization detection policy. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters.
	 */
	Newname string `json:"newname,omitempty"`

	//------- Read only Parameter ---------;

	Hits      string `json:"hits,omitempty"`
	Undefhits string `json:"undefhits,omitempty"`
	Builtin   string `json:"builtin,omitempty"`
	Feature   string `json:"feature,omitempty"`
}
