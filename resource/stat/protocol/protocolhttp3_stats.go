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

package protocol

/**
* Statistics for http3 resource.
*/

type Protocolhttp3stats struct {
	/**
	* Clear the statsistics / counters
	*/
	Clearstats string `json:"clearstats,omitempty"`
	Http3requestsrcvd uint64 `json:"http3requestsrcvd,omitempty"`
	/**
	* Total number of HTTP/3 requests received
	*/
	Http3requestsrcvdrate int32 `json:"http3requestsrcvdrate,omitempty"`
	Http3requestssent uint64 `json:"http3requestssent,omitempty"`
	/**
	* Total number of HTTP/3 requests sent
	*/
	Http3requestssentrate int32 `json:"http3requestssentrate,omitempty"`
	Http3responsesrcvd uint64 `json:"http3responsesrcvd,omitempty"`
	/**
	* Total number of HTTP/3 responses received
	*/
	Http3responsesrcvdrate int32 `json:"http3responsesrcvdrate,omitempty"`
	Http3responsessent uint64 `json:"http3responsessent,omitempty"`
	/**
	* Total number of HTTP/3 responses sent
	*/
	Http3responsessentrate int32 `json:"http3responsessentrate,omitempty"`
	Http3conninfalcfail uint64 `json:"http3conninfalcfail,omitempty"`
	/**
	* Number of HTTP/3 connection-info allocation failures
	*/
	Http3conninfalcfailrate int32 `json:"http3conninfalcfailrate,omitempty"`
	Http3nsbalcfail uint64 `json:"http3nsbalcfail,omitempty"`
	/**
	* Number of HTTP/3 NSB allocation failures
	*/
	Http3nsbalcfailrate int32 `json:"http3nsbalcfailrate,omitempty"`
	Http3strminfalcfail uint64 `json:"http3strminfalcfail,omitempty"`
	/**
	* Number of HTTP/3 stream-info allocation failures
	*/
	Http3strminfalcfailrate int32 `json:"http3strminfalcfailrate,omitempty"`
	Http3strmpcbalcfail uint64 `json:"http3strmpcbalcfail,omitempty"`
	/**
	* Number of HTTP/3 stream PCB allocation failures
	*/
	Http3strmpcbalcfailrate int32 `json:"http3strmpcbalcfailrate,omitempty"`

}
