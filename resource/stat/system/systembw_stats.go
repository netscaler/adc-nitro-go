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

package system

/**
* Statistics for bw resource.
*/

type Systembwstats struct {
	/**
	* Clear the statsistics / counters
	*/
	Clearstats string `json:"clearstats,omitempty"`
	Httpcltpoolinactive uint64 `json:"httpcltpoolinactive,omitempty"`
	Httpcltpooloutactive uint64 `json:"httpcltpooloutactive,omitempty"`
	Httpsvr200okresp uint64 `json:"httpsvr200okresp,omitempty"`
	/**
	* Number of 200 Ok response sent from the BW appliance.
	*/
	Httpsvr200okresprate int32 `json:"httpsvr200okresprate,omitempty"`
	Httpsvr404notfound uint64 `json:"httpsvr404notfound,omitempty"`
	/**
	* Number of 404 Not Found responses sent
	*/
	Httpsvr404notfoundrate int32 `json:"httpsvr404notfoundrate,omitempty"`
	Httpclterrstray uint64 `json:"httpclterrstray,omitempty"`
	/**
	* Number of stray packets received from server without HTTP request
	*/
	Httpclterrstrayrate int32 `json:"httpclterrstrayrate,omitempty"`
	Httpcltttfplwm uint64 `json:"httpcltttfplwm,omitempty"`
	/**
	* Number of Responses Falling on LWM for TTFP.
	*/
	Httpcltttfplwmrate int32 `json:"httpcltttfplwmrate,omitempty"`
	Httpcltttfp0 uint64 `json:"httpcltttfp_0,omitempty"`
	/**
	* Number of Responses Falling on Band-0 for TTFP.
	*/
	Httpcltttfp0rate int32 `json:"httpcltttfp_0rate,omitempty"`
	Httpcltttfp1 uint64 `json:"httpcltttfp_1,omitempty"`
	/**
	* Number of Responses Falling on Band-1 for TTFP.
	*/
	Httpcltttfp1rate int32 `json:"httpcltttfp_1rate,omitempty"`
	Httpcltttfp2 uint64 `json:"httpcltttfp_2,omitempty"`
	/**
	* Number of Responses Falling on Band-2 for TTFP.
	*/
	Httpcltttfp2rate int32 `json:"httpcltttfp_2rate,omitempty"`
	Httpcltttfp3 uint64 `json:"httpcltttfp_3,omitempty"`
	/**
	* Number of Responses Falling on Band-3 for TTFP.
	*/
	Httpcltttfp3rate int32 `json:"httpcltttfp_3rate,omitempty"`
	Httpcltttfp4 uint64 `json:"httpcltttfp_4,omitempty"`
	/**
	* Number of Responses Falling on Band-4 for TTFP.
	*/
	Httpcltttfp4rate int32 `json:"httpcltttfp_4rate,omitempty"`
	Httpcltttfp5 uint64 `json:"httpcltttfp_5,omitempty"`
	/**
	* Number of Responses Falling on Band-5 for TTFP.
	*/
	Httpcltttfp5rate int32 `json:"httpcltttfp_5rate,omitempty"`
	Httpcltttfp6 uint64 `json:"httpcltttfp_6,omitempty"`
	/**
	* Number of Responses Falling on Band-6 for TTFP.
	*/
	Httpcltttfp6rate int32 `json:"httpcltttfp_6rate,omitempty"`
	Httpcltttfp7 uint64 `json:"httpcltttfp_7,omitempty"`
	/**
	* Number of Responses Falling on Band-7 for TTFP.
	*/
	Httpcltttfp7rate int32 `json:"httpcltttfp_7rate,omitempty"`
	Httpcltttfphwm uint64 `json:"httpcltttfphwm,omitempty"`
	/**
	* Number of Responses Falling on HWM for TTFP.
	*/
	Httpcltttfphwmrate int32 `json:"httpcltttfphwmrate,omitempty"`
	Httpcltttfpmax uint64 `json:"httpcltttfpmax,omitempty"`
	Httpcltttlplwm uint64 `json:"httpcltttlplwm,omitempty"`
	/**
	* Number of Responses Falling on LWM for TTLP.
	*/
	Httpcltttlplwmrate int32 `json:"httpcltttlplwmrate,omitempty"`
	Httpcltttlp0 uint64 `json:"httpcltttlp_0,omitempty"`
	/**
	* Number of Responses Falling on Band-0 for TTLP.
	*/
	Httpcltttlp0rate int32 `json:"httpcltttlp_0rate,omitempty"`
	Httpcltttlp1 uint64 `json:"httpcltttlp_1,omitempty"`
	/**
	* Number of Responses Falling on Band-1 for TTLP.
	*/
	Httpcltttlp1rate int32 `json:"httpcltttlp_1rate,omitempty"`
	Httpcltttlp2 uint64 `json:"httpcltttlp_2,omitempty"`
	/**
	* Number of Responses Falling on Band-2 for TTLP.
	*/
	Httpcltttlp2rate int32 `json:"httpcltttlp_2rate,omitempty"`
	Httpcltttlp3 uint64 `json:"httpcltttlp_3,omitempty"`
	/**
	* Number of Responses Falling on Band-3 for TTLP.
	*/
	Httpcltttlp3rate int32 `json:"httpcltttlp_3rate,omitempty"`
	Httpcltttlp4 uint64 `json:"httpcltttlp_4,omitempty"`
	/**
	* Number of Responses Falling on Band-4 for TTLP.
	*/
	Httpcltttlp4rate int32 `json:"httpcltttlp_4rate,omitempty"`
	Httpcltttlp5 uint64 `json:"httpcltttlp_5,omitempty"`
	/**
	* Number of Responses Falling on Band-5 for TTLP.
	*/
	Httpcltttlp5rate int32 `json:"httpcltttlp_5rate,omitempty"`
	Httpcltttlp6 uint64 `json:"httpcltttlp_6,omitempty"`
	/**
	* Number of Responses Falling on Band-6 for TTLP.
	*/
	Httpcltttlp6rate int32 `json:"httpcltttlp_6rate,omitempty"`
	Httpcltttlp7 uint64 `json:"httpcltttlp_7,omitempty"`
	/**
	* Number of Responses Falling on Band-7 for TTLP.
	*/
	Httpcltttlp7rate int32 `json:"httpcltttlp_7rate,omitempty"`
	Httpcltttlphwm uint64 `json:"httpcltttlphwm,omitempty"`
	/**
	* Number of Responses Falling on HWM for TTLP.
	*/
	Httpcltttlphwmrate int32 `json:"httpcltttlphwmrate,omitempty"`
	Httpcltttlpmax uint64 `json:"httpcltttlpmax,omitempty"`

}
