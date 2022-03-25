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

package cluster

/**
* Binding class showing the identifier that can be bound to clusternodegroup.
 */
type Clusternodegroupidentifierbinding struct {
	/**
	* stream identifier  and rate limit identifier that need to be bound to this nodegroup.
	 */
	Identifiername string `json:"identifiername,omitempty"`
	/**
	* Name of the nodegroup to which you want to bind a cluster node or an entity.
	 */
	Name string `json:"name,omitempty"`
}
