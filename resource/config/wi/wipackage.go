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
* Configuration for Web Interface resource.
 */
type Wipackage struct {
	/**
	* Complete path to the JRE tar file.
		You can use OpenJDK7 package for FreeBSD 8.x/amd64.The Java package can be downloaded from http://ftp-archive.freebsd.org/pub/FreeBSD-Archive/old-releases/amd64/amd64/8.4-RELEASE/packages/java/openjdk-7.17.02_2.tbz
	*/
	Jre string `json:"jre,omitempty"`
	/**
	* Complete path to the Web Interface tar file for installing the Web Interface on the Citrix ADC. This file includes Apache Tomcat Web server. The file name has the following format: nswi-<version number>.tgz (for example, nswi-1.5.tgz).
	 */
	Wi string `json:"wi,omitempty"`
	/**
	* Maximum number of Web Interface sites that can be created on the Citrix ADC; changes the amount of RAM reserved for Web Interface usage; changing its value results in restart of Tomcat server and invalidates any existing Web Interface sessions.
	 */
	Maxsites string `json:"maxsites,omitempty"`
}
