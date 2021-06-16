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
* Statistics for UDP Protocol resource.
*/

type Protocoludpstats struct {
	/**
	* Clear the statsistics / counters
	*/
	Clearstats string `json:"clearstats,omitempty"`
	Udptotrxpkts uint64 `json:"udptotrxpkts,omitempty"`
	/**
	* Total number of UDP packets received.
	*/
	Udprxpktsrate float64 `json:"udprxpktsrate,omitempty"`
	Udptotrxbytes uint64 `json:"udptotrxbytes,omitempty"`
	/**
	* Total number of UDP data received in bytes.
	*/
	Udprxbytesrate float64 `json:"udprxbytesrate,omitempty"`
	Udptottxpkts uint64 `json:"udptottxpkts,omitempty"`
	/**
	* Total number of UDP packets transmitted.
	*/
	Udptxpktsrate float64 `json:"udptxpktsrate,omitempty"`
	Udptottxbytes uint64 `json:"udptottxbytes,omitempty"`
	/**
	* Total number of UDP data transmitted in bytes.
	*/
	Udptxbytesrate float64 `json:"udptxbytesrate,omitempty"`
	Udpcurratethreshold uint32 `json:"udpcurratethreshold,omitempty"`
	Udptotunknownsvcpkts uint64 `json:"udptotunknownsvcpkts,omitempty"`
	Udpbadchecksum uint64 `json:"udpbadchecksum,omitempty"`
	Udpcurratethresholdexceeds uint64 `json:"udpcurratethresholdexceeds,omitempty"`

}
