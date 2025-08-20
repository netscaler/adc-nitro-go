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

package endpoint

/**
* Configuration for Information resource.
*/
type Endpointinfo struct {
	/**
	* Endpoint kind. Currently, IP endpoints are supported
	*/
	Endpointkind string `json:"endpointkind,omitempty"`
	/**
	* Name of endpoint, depends on kind. For IP Endpoint - IP address.
	*/
	Endpointname string `json:"endpointname,omitempty"`
	/**
	* String of qualifiers, in dotted notation, structured metadata for an endpoint. Each qualifier is more specific than the one that precedes it, as in cluster.namespace.service. For example: cluster.default.frontend. 
		Note: A qualifier that includes a dot (.) or space ( ) must be enclosed in double quotation marks.
	*/
	Endpointmetadata string `json:"endpointmetadata,omitempty"`
	/**
	* String representing labels in json form. Maximum length 16K
	*/
	Endpointlabelsjson string `json:"endpointlabelsjson,omitempty"`

	//------- Read only Parameter ---------;

	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
