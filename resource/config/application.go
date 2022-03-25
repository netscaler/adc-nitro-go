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
* Configuration for application resource.
 */
type Application struct {
	/**
	* Name of the AppExpert application template file.
	 */
	Apptemplatefilename string `json:"apptemplatefilename,omitempty"`
	/**
	* Name to assign to the application on the Citrix ADC. If you do not provide a name, the appliance assigns the application the name of the template file.
	 */
	Appname string `json:"appname,omitempty"`
	/**
	* Name of the deployment file.
	 */
	Deploymentfilename string `json:"deploymentfilename,omitempty"`
}
