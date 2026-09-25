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
* Configuration for WASM module related files resource.
*/
type Wasmfile struct {
	/**
	* Local path or URL (protocol, host, path, and file name) for the file from which to retrieve the imported HTML page.
		NOTE: The import fails if the object to be imported is on an HTTPS server that requires client certificate authentication for access.
	*/
	Src string `json:"src,omitempty"`
	/**
	* Name to assign to the WASM mdoule/signature page object on the Citrix ADC.
	*/
	Name string `json:"name,omitempty"`
	/**
	* WASM file type to be imported.
	*/
	Filetype string `json:"filetype,omitempty"`
	/**
	* Any comments to preserve information about the WASM page object.
	*/
	Comment string `json:"comment,omitempty"`
	/**
	* Overwrites the existing file
	*/
	Overwrite bool `json:"overwrite,omitempty"`

	//------- Read only Parameter ---------;

	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
