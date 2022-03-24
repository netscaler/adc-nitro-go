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

package stat

/**
* Statistics for bw resource.
 */

type Systembwstats struct {
	/**
	* Clear the statsistics / counters
	 */
	Clearstats           string `json:"clearstats,omitempty"`
	Httpcltpoolinactive  int    `json:"httpcltpoolinactive,omitempty"`
	Httpcltpooloutactive int    `json:"httpcltpooloutactive,omitempty"`
	Httpsvr200okresp     int    `json:"httpsvr200okresp,omitempty"`
	/**
	* Number of 200 Ok response sent from the BW appliance.
	 */
	Httpsvr200okresprate float64 `json:"httpsvr200okresprate,omitempty"`
	Httpsvr404notfound   int     `json:"httpsvr404notfound,omitempty"`
	/**
	* Number of 404 Not Found responses sent
	 */
	Httpsvr404notfoundrate float64 `json:"httpsvr404notfoundrate,omitempty"`
	Httpclterrstray        int     `json:"httpclterrstray,omitempty"`
	/**
	* Number of stray packets received from server without HTTP request
	 */
	Httpclterrstrayrate float64 `json:"httpclterrstrayrate,omitempty"`
	Httpcltttfplwm      int     `json:"httpcltttfplwm,omitempty"`
	/**
	* Number of Responses Falling on LWM for TTFP.
	 */
	Httpcltttfplwmrate float64 `json:"httpcltttfplwmrate,omitempty"`
	Httpcltttfp0       int     `json:"httpcltttfp_0,omitempty"`
	/**
	* Number of Responses Falling on Band-0 for TTFP.
	 */
	Httpcltttfp0rate float64 `json:"httpcltttfp_0rate,omitempty"`
	Httpcltttfp1     int     `json:"httpcltttfp_1,omitempty"`
	/**
	* Number of Responses Falling on Band-1 for TTFP.
	 */
	Httpcltttfp1rate float64 `json:"httpcltttfp_1rate,omitempty"`
	Httpcltttfp2     int     `json:"httpcltttfp_2,omitempty"`
	/**
	* Number of Responses Falling on Band-2 for TTFP.
	 */
	Httpcltttfp2rate float64 `json:"httpcltttfp_2rate,omitempty"`
	Httpcltttfp3     int     `json:"httpcltttfp_3,omitempty"`
	/**
	* Number of Responses Falling on Band-3 for TTFP.
	 */
	Httpcltttfp3rate float64 `json:"httpcltttfp_3rate,omitempty"`
	Httpcltttfp4     int     `json:"httpcltttfp_4,omitempty"`
	/**
	* Number of Responses Falling on Band-4 for TTFP.
	 */
	Httpcltttfp4rate float64 `json:"httpcltttfp_4rate,omitempty"`
	Httpcltttfp5     int     `json:"httpcltttfp_5,omitempty"`
	/**
	* Number of Responses Falling on Band-5 for TTFP.
	 */
	Httpcltttfp5rate float64 `json:"httpcltttfp_5rate,omitempty"`
	Httpcltttfp6     int     `json:"httpcltttfp_6,omitempty"`
	/**
	* Number of Responses Falling on Band-6 for TTFP.
	 */
	Httpcltttfp6rate float64 `json:"httpcltttfp_6rate,omitempty"`
	Httpcltttfp7     int     `json:"httpcltttfp_7,omitempty"`
	/**
	* Number of Responses Falling on Band-7 for TTFP.
	 */
	Httpcltttfp7rate float64 `json:"httpcltttfp_7rate,omitempty"`
	Httpcltttfphwm   int     `json:"httpcltttfphwm,omitempty"`
	/**
	* Number of Responses Falling on HWM for TTFP.
	 */
	Httpcltttfphwmrate float64 `json:"httpcltttfphwmrate,omitempty"`
	Httpcltttfpmax     int     `json:"httpcltttfpmax,omitempty"`
	Httpcltttlplwm     int     `json:"httpcltttlplwm,omitempty"`
	/**
	* Number of Responses Falling on LWM for TTLP.
	 */
	Httpcltttlplwmrate float64 `json:"httpcltttlplwmrate,omitempty"`
	Httpcltttlp0       int     `json:"httpcltttlp_0,omitempty"`
	/**
	* Number of Responses Falling on Band-0 for TTLP.
	 */
	Httpcltttlp0rate float64 `json:"httpcltttlp_0rate,omitempty"`
	Httpcltttlp1     int     `json:"httpcltttlp_1,omitempty"`
	/**
	* Number of Responses Falling on Band-1 for TTLP.
	 */
	Httpcltttlp1rate float64 `json:"httpcltttlp_1rate,omitempty"`
	Httpcltttlp2     int     `json:"httpcltttlp_2,omitempty"`
	/**
	* Number of Responses Falling on Band-2 for TTLP.
	 */
	Httpcltttlp2rate float64 `json:"httpcltttlp_2rate,omitempty"`
	Httpcltttlp3     int     `json:"httpcltttlp_3,omitempty"`
	/**
	* Number of Responses Falling on Band-3 for TTLP.
	 */
	Httpcltttlp3rate float64 `json:"httpcltttlp_3rate,omitempty"`
	Httpcltttlp4     int     `json:"httpcltttlp_4,omitempty"`
	/**
	* Number of Responses Falling on Band-4 for TTLP.
	 */
	Httpcltttlp4rate float64 `json:"httpcltttlp_4rate,omitempty"`
	Httpcltttlp5     int     `json:"httpcltttlp_5,omitempty"`
	/**
	* Number of Responses Falling on Band-5 for TTLP.
	 */
	Httpcltttlp5rate float64 `json:"httpcltttlp_5rate,omitempty"`
	Httpcltttlp6     int     `json:"httpcltttlp_6,omitempty"`
	/**
	* Number of Responses Falling on Band-6 for TTLP.
	 */
	Httpcltttlp6rate float64 `json:"httpcltttlp_6rate,omitempty"`
	Httpcltttlp7     int     `json:"httpcltttlp_7,omitempty"`
	/**
	* Number of Responses Falling on Band-7 for TTLP.
	 */
	Httpcltttlp7rate float64 `json:"httpcltttlp_7rate,omitempty"`
	Httpcltttlphwm   int     `json:"httpcltttlphwm,omitempty"`
	/**
	* Number of Responses Falling on HWM for TTLP.
	 */
	Httpcltttlphwmrate float64 `json:"httpcltttlphwmrate,omitempty"`
	Httpcltttlpmax     int     `json:"httpcltttlpmax,omitempty"`
}
