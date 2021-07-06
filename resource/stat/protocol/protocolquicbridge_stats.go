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
* Statistics for QUIC Bridge protocol resource.
*/

type Protocolquicbridgestats struct {
	/**
	* Clear the statsistics / counters
	*/
	Clearstats string `json:"clearstats,omitempty"`
	Quicbridgeconn int `json:"quicbridgeconn,omitempty"`
	/**
	* Total number of QUIC Bridge connections
	*/
	Quicbridgeconnrate float64 `json:"quicbridgeconnrate,omitempty"`
	Quicbridgemigratedconn int `json:"quicbridgemigratedconn,omitempty"`
	/**
	* Total number of migrated QUIC Bridge connections
	*/
	Quicbridgemigratedconnrate float64 `json:"quicbridgemigratedconnrate,omitempty"`
	Quicbridgeqci int `json:"quicbridgeqci,omitempty"`
	/**
	* Current number of QUIC Bridge connection infos
	*/
	Quicbridgeqcirate float64 `json:"quicbridgeqcirate,omitempty"`
	Quicbridgeqpi int `json:"quicbridgeqpi,omitempty"`
	/**
	* Current number of QUIC Bridge peer infos
	*/
	Quicbridgeqpirate float64 `json:"quicbridgeqpirate,omitempty"`
	Quicbridgeqpialcfail int `json:"quicbridgeqpialcfail,omitempty"`
	/**
	* Number of QUIC Bridge peer info allocation failures
	*/
	Quicbridgeqpialcfailrate float64 `json:"quicbridgeqpialcfailrate,omitempty"`
	Quicbridgeqcialcfail int `json:"quicbridgeqcialcfail,omitempty"`
	/**
	* Number of QUIC Bridge connection info allocation failures
	*/
	Quicbridgeqcialcfailrate float64 `json:"quicbridgeqcialcfailrate,omitempty"`

}
