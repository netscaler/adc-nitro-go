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

package metrics

/**
* Configuration for Metrics profile resource.
*/
type Metricsprofile struct {
	/**
	* Name for the metrics profile. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at
		(@), equals (=), and hyphen (-) characters.!
	*/
	Name string `json:"name,omitempty"`
	/**
	* The collector should be a HTTP/HTTPS service. 
	*/
	Collector string `json:"collector,omitempty"`
	/**
	* This option indicates the format in which metrics data is generated
	*/
	Outputmode string `json:"outputmode,omitempty"`
	/**
	* This option is used enable or disable metrics 
	*/
	Metrics string `json:"metrics,omitempty"`
	/**
	* This option is to configure metrics pull or push mode. In push mode metricscollector exports metrics to configured collector. In pull mode, metricscollector only generates the metrics which will be pulled by external agent. No collector configuration is required in pull mode and it is applicable only for output mode Prometheus
	*/
	Servemode string `json:"servemode,omitempty"`
	/**
	* This option is for configuring json schema file containing a list of counters to be exported by metricscollector. Schema file should be present under /var/metrics_conf path
	*/
	Schemafile string `json:"schemafile,omitempty"`
	/**
	* This option is for configuring the metrics export frequency in seconds, frequency value must be in [30,300] seconds range
	*/
	Metricsexportfrequency *int `json:"metricsexportfrequency,omitempty"`
	/**
	* Token for authenticating with the endpoint. If the endpoint requires the Authorization header in a particular format, specify the complete format as the value to this parameter. For eg., in case of splunk, the Authorizaiton header is required to be of the form - Splunk <auth-token>.
	*/
	Metricsauthtoken string `json:"metricsauthtoken,omitempty"`
	/**
	* The URL at which to upload the metrics data on the endpoint
	*/
	Metricsendpointurl string `json:"metricsendpointurl,omitempty"`

	//------- Read only Parameter ---------;

	Refcnt string `json:"refcnt,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
