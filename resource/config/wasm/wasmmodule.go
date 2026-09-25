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

package wasm

/**
* Configuration for WASM Module resource.
*/
type Wasmmodule struct {
	/**
	* The name of the WASM module file.
	*/
	Name string `json:"name,omitempty"`
	/**
	* File name of the WASM module.
	*/
	Modulefile string `json:"modulefile,omitempty"`
	/**
	* The SHA256 file contains the hash value used to validate the WASM module..
	*/
	Signaturefile string `json:"signaturefile,omitempty"`
	/**
	* The WASM module filename contains module-specific configuration settings.
	*/
	Settingfile string `json:"settingfile,omitempty"`
	/**
	* Any type of information about this WASM module
	*/
	Comment string `json:"comment,omitempty"`

	//------- Read only Parameter ---------;

	Referencecount string `json:"referencecount,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
