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
* Statistics for QUIC protocol resource.
*/

type Protocolquicstats struct {
	/**
	* Clear the statsistics / counters
	*/
	Clearstats string `json:"clearstats,omitempty"`
	Quicclientdgrmrcvd uint64 `json:"quicclientdgrmrcvd,omitempty"`
	/**
	* Total QUIC client UDP datagrams received
	*/
	Quicclientdgrmrcvdrate float64 `json:"quicclientdgrmrcvdrate,omitempty"`
	Quicserverdgrmrcvd uint64 `json:"quicserverdgrmrcvd,omitempty"`
	/**
	* Total QUIC server UDP datagrams received
	*/
	Quicserverdgrmrcvdrate float64 `json:"quicserverdgrmrcvdrate,omitempty"`
	Quicclientdgrmsent uint64 `json:"quicclientdgrmsent,omitempty"`
	/**
	* Total QUIC client UDP datagrams sent
	*/
	Quicclientdgrmsentrate float64 `json:"quicclientdgrmsentrate,omitempty"`
	Quicserverdgrmsent uint64 `json:"quicserverdgrmsent,omitempty"`
	/**
	* Total QUIC server UDP datagrams sent
	*/
	Quicserverdgrmsentrate float64 `json:"quicserverdgrmsentrate,omitempty"`
	Quiccurclientconn uint64 `json:"quiccurclientconn,omitempty"`
	/**
	* Current QUIC client connections
	*/
	Quiccurclientconnrate float64 `json:"quiccurclientconnrate,omitempty"`
	Quiccurserverconn uint64 `json:"quiccurserverconn,omitempty"`
	/**
	* Current QUIC server connections
	*/
	Quiccurserverconnrate float64 `json:"quiccurserverconnrate,omitempty"`
	Quiclocalconnid uint64 `json:"quiclocalconnid,omitempty"`
	/**
	* Current QUIC local connection IDs allocated
	*/
	Quiclocalconnidrate float64 `json:"quiclocalconnidrate,omitempty"`
	Quictotclientconn uint64 `json:"quictotclientconn,omitempty"`
	/**
	* Total QUIC client connections
	*/
	Quicclientconnrate float64 `json:"quicclientconnrate,omitempty"`
	Quictotserverconn uint64 `json:"quictotserverconn,omitempty"`
	/**
	* Total QUIC server connections
	*/
	Quicserverconnrate float64 `json:"quicserverconnrate,omitempty"`
	Quicmigratedconn uint64 `json:"quicmigratedconn,omitempty"`
	/**
	* Total number of migrated QUIC connections
	*/
	Quicmigratedconnrate float64 `json:"quicmigratedconnrate,omitempty"`
	Quicjumboframesrcvd uint64 `json:"quicjumboframesrcvd,omitempty"`
	/**
	* Total number of QUIC jumbo frames received
	*/
	Quicjumboframesrcvdrate float64 `json:"quicjumboframesrcvdrate,omitempty"`
	Quicretrypktsent uint64 `json:"quicretrypktsent,omitempty"`
	/**
	* Number of QUIC Retry packets sent
	*/
	Quicretrypktsentrate float64 `json:"quicretrypktsentrate,omitempty"`
	Quichandshakecmpltd uint64 `json:"quichandshakecmpltd,omitempty"`
	/**
	* Number of QUIC handshake messages completed
	*/
	Quichandshakecmpltdrate float64 `json:"quichandshakecmpltdrate,omitempty"`
	Quictransptconnclosepktsent uint64 `json:"quictransptconnclosepktsent,omitempty"`
	/**
	* Number of QUIC transport no-error Connection Close packets sent
	*/
	Quictransptconnclosepktsentrate float64 `json:"quictransptconnclosepktsentrate,omitempty"`
	Quicappconnclosepktsent uint64 `json:"quicappconnclosepktsent,omitempty"`
	/**
	* Number of QUIC application no-error Connection Close packets sent
	*/
	Quicappconnclosepktsentrate float64 `json:"quicappconnclosepktsentrate,omitempty"`
	Quicconninfoalcfail uint64 `json:"quicconninfoalcfail,omitempty"`
	/**
	* Quic session allocations failed
	*/
	Quicconninfoalcfailrate float64 `json:"quicconninfoalcfailrate,omitempty"`
	Quicnsbalcfail uint64 `json:"quicnsbalcfail,omitempty"`
	/**
	* Quic NSB allocations failed
	*/
	Quicnsbalcfailrate float64 `json:"quicnsbalcfailrate,omitempty"`
	Quictlsalertsent uint64 `json:"quictlsalertsent,omitempty"`
	/**
	* Total QUIC TLS 1.3 transport errors sent
	*/
	Quictlsalertsentrate float64 `json:"quictlsalertsentrate,omitempty"`
	Quicstlessconnclosepktsent uint64 `json:"quicstlessconnclosepktsent,omitempty"`
	/**
	* Number of QUIC stateless Connection Close packets sent
	*/
	Quicstlessconnclosepktsentrate float64 `json:"quicstlessconnclosepktsentrate,omitempty"`
	Quicvernegpktsent uint64 `json:"quicvernegpktsent,omitempty"`
	/**
	* Number of QUIC Version Negotiation packets sent
	*/
	Quicvernegpktsentrate float64 `json:"quicvernegpktsentrate,omitempty"`
	Quictransptconnclosepktfail uint64 `json:"quictransptconnclosepktfail,omitempty"`
	/**
	* Number of QUIC transport error Connection Close packets sent
	*/
	Quictransptconnclosepktfailrate float64 `json:"quictransptconnclosepktfailrate,omitempty"`
	Quicappconnclosepktfail uint64 `json:"quicappconnclosepktfail,omitempty"`
	/**
	* Number of QUIC application error Connection Close packets sent
	*/
	Quicappconnclosepktfailrate float64 `json:"quicappconnclosepktfailrate,omitempty"`
	Quicretrytokenverfail uint64 `json:"quicretrytokenverfail,omitempty"`
	/**
	* Number of times QUIC Retry token verification failed
	*/
	Quicretrytokenverfailrate float64 `json:"quicretrytokenverfailrate,omitempty"`
	Quicnewtokenverfail uint64 `json:"quicnewtokenverfail,omitempty"`
	/**
	* Number of times QUIC NEW_TOKEN token verification failed
	*/
	Quicnewtokenverfailrate float64 `json:"quicnewtokenverfailrate,omitempty"`

}
