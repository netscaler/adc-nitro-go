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

package urlfiltering

/**
* Configuration for URLFILTERING paramter resource.
*/
type Urlfilteringparameter struct {
	/**
	* URL Filtering hours between DB updates.
	*/
	Hoursbetweendbupdates int `json:"hoursbetweendbupdates,omitempty"`
	/**
	* URL Filtering time of day to update DB.
	*/
	Timeofdaytoupdatedb string `json:"timeofdaytoupdatedb,omitempty"`
	/**
	* URL Filtering Local DB number of threads.
	*/
	Localdatabasethreads int `json:"localdatabasethreads,omitempty"`
	/**
	* URL Filtering Cloud host.
	*/
	Cloudhost string `json:"cloudhost,omitempty"`
	/**
	* URL Filtering Seed DB path.
	*/
	Seeddbpath string `json:"seeddbpath,omitempty"`

	//------- Read only Parameter ---------;

	Maxnumberofcloudthreads string `json:"maxnumberofcloudthreads,omitempty"`
	Cloudkeepalivetimeout string `json:"cloudkeepalivetimeout,omitempty"`
	Cloudserverconnecttimeout string `json:"cloudserverconnecttimeout,omitempty"`
	Clouddblookuptimeout string `json:"clouddblookuptimeout,omitempty"`
	Proxyhostip string `json:"proxyhostip,omitempty"`
	Proxyport string `json:"proxyport,omitempty"`
	Proxyusername string `json:"proxyusername,omitempty"`
	Proxypassword string `json:"proxypassword,omitempty"`
	Seeddbsizelevel string `json:"seeddbsizelevel,omitempty"`

}
