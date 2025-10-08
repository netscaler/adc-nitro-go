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

package pq

/**
* Configuration for PQ policy resource.
*/
type Pqpolicy struct {
	/**
	* Name for the priority queuing policy. Must begin with a letter, number, or the underscore symbol (_). Other characters allowed, after the first character, are the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), and colon (:) characters.
	*/
	Policyname string `json:"policyname,omitempty"`
	/**
	* Expression or name of a named expression, against which the request is evaluated. The priority queuing policy is applied if the rule evaluates to true.
		Note:
		* On the command line interface, if the expression includes blank spaces, the entire expression must be enclosed in double quotation marks.
		* If the expression itself includes double quotation marks, you must escape the quotations by using the \ character. 
		* Alternatively, you can use single quotation marks to enclose the rule, in which case you will not have to escape the double quotation marks.
		* Maximum length of a string literal in the expression is 255 characters. A longer string can be split into smaller strings of up to 255 characters each, and the smaller strings concatenated with the + operator. For example, you can create a 500-character string as follows: '"<string of 255 characters>" + "<string of 245 characters>"'
	*/
	Rule string `json:"rule,omitempty"`
	/**
	* Priority for queuing the request. If server resources are not available for a request that matches the configured rule, this option specifies a priority for queuing the request until the server resources are available again. Enter the value of positive_integer as 1, 2 or 3. The highest priority level is 1 and the lowest priority value is 3.
	*/
	Priority *int `json:"priority,omitempty"`
	/**
	* Weight of the priority. Each priority is assigned a weight according to which it is served when server resources are available. The weight for a higher priority request must be set higher than that of a lower priority request.
		To prevent delays for low-priority requests across multiple priority levels, you can configure weighted queuing for serving requests. The default weights for the priorities
		are:
		* Gold - Priority 1 - Weight 3
		* Silver - Priority 2 - Weight 2
		* Bronze - Priority 3 - Weight 1
		Specify the weights as 0 through 101. A weight of 0 indicates that the particular priority level should be served only when there are no requests in any of the priority queues.
		A weight of 101 specifies a weight of infinity. This means that this priority level is served irrespective of the number of clients waiting in other priority queues.
	*/
	Weight *int `json:"weight,omitempty"`
	/**
	* Queue depth threshold value. When the queue size (number of requests in the queue) on the virtual server to which this policy is bound, increases to the specified qDepth value, subsequent requests are dropped to the lowest priority level.
	*/
	Qdepth *int `json:"qdepth,omitempty"`
	/**
	* Policy queue depth threshold value. When the policy queue size (number of requests in all the queues belonging to this policy) increases to the specified polqDepth value, subsequent requests are dropped to the lowest priority level.
	*/
	Polqdepth *int `json:"polqdepth,omitempty"`

	//------- Read only Parameter ---------;

	Hits string `json:"hits,omitempty"`

}
