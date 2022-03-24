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

package wi

/**
* Binding object which returns the resources bound to wisite_binding.
 */
type Wisitebinding struct {
	/**
	* Path of a Web Interface site whose details you want the Citrix ADC to display.<br/>Minimum value =
	 */
	Sitepath string `json:"sitepath,omitempty"`
}
