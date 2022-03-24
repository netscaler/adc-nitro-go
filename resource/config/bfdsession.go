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
* Configuration for BFD configuration resource.
 */
type Bfdsession struct {
	/**
	* IPV4 or IPV6 Address of Local Node
	 */
	Localip string `json:"localip,omitempty"`
	/**
	* IPV4 or IPV6 Address of Remote Node
	 */
	Remoteip string `json:"remoteip,omitempty"`

	//------- Read only Parameter ---------;

	State                             string `json:"state,omitempty"`
	Localport                         string `json:"localport,omitempty"`
	Remoteport                        string `json:"remoteport,omitempty"`
	Minimumtransmitinterval           string `json:"minimumtransmitinterval,omitempty"`
	Negotiatedminimumtransmitinterval string `json:"negotiatedminimumtransmitinterval,omitempty"`
	Minimumreceiveinterval            string `json:"minimumreceiveinterval,omitempty"`
	Negotiatedminimumreceiveinterval  string `json:"negotiatedminimumreceiveinterval,omitempty"`
	Multiplier                        string `json:"multiplier,omitempty"`
	Remotemultiplier                  string `json:"remotemultiplier,omitempty"`
	Vlan                              string `json:"vlan,omitempty"`
	Localdiagnotic                    string `json:"localdiagnotic,omitempty"`
	Localdiscriminator                string `json:"localdiscriminator,omitempty"`
	Remotediscriminator               string `json:"remotediscriminator,omitempty"`
	Passive                           string `json:"passive,omitempty"`
	Multihop                          string `json:"multihop,omitempty"`
	Admindown                         string `json:"admindown,omitempty"`
	Originalownerpe                   string `json:"originalownerpe,omitempty"`
	Currentownerpe                    string `json:"currentownerpe,omitempty"`
	Ownernode                         string `json:"ownernode,omitempty"`
}
