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

package transform

/**
* Binding class showing the transformpolicy that can be bound to transformpolicylabel.
 */
type Transformpolicylabeltransformpolicybinding struct {
	/**
	* Name of the URL Transformation policy to bind to the policy label.
	 */
	Policyname string `json:"policyname,omitempty"`
	/**
	* Specifies the priority of the policy.
	 */
	Priority int `json:"priority,omitempty"`
	/**
	* Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
	 */
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	/**
	* If the current policy evaluates to TRUE, terminate evaluation of policies bound to the current policy label, and then forward the request to the specified virtual server or evaluate the specified policy label.
	 */
	Invoke bool `json:"invoke,omitempty"`
	/**
	* Type of invocation. Available settings function as follows:
		* reqvserver - Forward the request to the specified request virtual server.
		* policylabel - Invoke the specified policy label.
	*/
	Labeltype string `json:"labeltype,omitempty"`
	/**
	* Name of the policy label.
	 */
	Invokelabelname string `json:"invoke_labelname,omitempty"`
	/**
	* Name of the URL Transformation policy label to which to bind the policy.
	 */
	Labelname string `json:"labelname,omitempty"`
}
