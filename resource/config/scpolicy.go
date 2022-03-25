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

package config

/**
* Configuration for SureConnect policy resource.
 */
type Scpolicy struct {
	/**
	* Name for the policy. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters.
	 */
	Name string `json:"name,omitempty"`
	/**
	* URL against which to match incoming client request.
	 */
	Url string `json:"url,omitempty"`
	/**
	* Expression against which the traffic is evaluated.
		Maximum length of a string literal in the expression is 255 characters. A longer string can be split into smaller strings of up to 255 characters each, and the smaller strings concatenated with the + operator. For example, you can create a 500-character string as follows: '"<string of 255 characters>" + "<string of 245 characters>"'
		The following requirements apply only to the Citrix ADC CLI:
		* If the expression includes one or more spaces, enclose the entire expression in double quotation marks.
		* If the expression itself includes double quotation marks, escape the quotations by using the  character.
		* Alternatively, you can use single quotation marks to enclose the rule, in which case you do not have to escape the double quotation marks.
	*/
	Rule string `json:"rule,omitempty"`
	/**
	* Delay threshold, in microseconds, for requests that match the policy's URL or rule. If the delay statistics gathered for the matching request exceed the specified delay, SureConnect is triggered for that request.
	 */
	Delay int `json:"delay,omitempty"`
	/**
	* Maximum number of concurrent connections that can be open for requests that match the policy's URL or rule.
	 */
	Maxconn int `json:"maxconn,omitempty"`
	/**
	* Action to be taken when the delay or maximum-connections threshold is reached. Available settings function as follows:
		ACS - Serve content from an alternative content service.
		NS - Serve alternative content from the Citrix ADC.
		NO ACTION - Serve no alternative content. However, delay statistics are still collected for the configured URLs, and, if the Maximum Client Connections parameter is set, the number of connections is limited to the value specified by that parameter. (However, alternative content is not served even if the maxConn threshold is met).
	*/
	Action string `json:"action,omitempty"`
	/**
	* Name of the alternative content service to be used in the ACS action.
	 */
	Altcontentsvcname string `json:"altcontentsvcname,omitempty"`
	/**
	* Path to the alternative content service to be used in the ACS action.
	 */
	Altcontentpath string `json:"altcontentpath,omitempty"`
}
