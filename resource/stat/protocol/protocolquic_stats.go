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
	Quicclientdgrmrcvd int `json:"quicclientdgrmrcvd,omitempty"`
	/**
	* Total QUIC client UDP datagrams received
	*/
	Quicclientdgrmrcvdrate float64 `json:"quicclientdgrmrcvdrate,omitempty"`
	Quicserverdgrmrcvd int `json:"quicserverdgrmrcvd,omitempty"`
	/**
	* Total QUIC server UDP datagrams received
	*/
	Quicserverdgrmrcvdrate float64 `json:"quicserverdgrmrcvdrate,omitempty"`
	Quicclientdgrmsent int `json:"quicclientdgrmsent,omitempty"`
	/**
	* Total QUIC client UDP datagrams sent
	*/
	Quicclientdgrmsentrate float64 `json:"quicclientdgrmsentrate,omitempty"`
	Quicserverdgrmsent int `json:"quicserverdgrmsent,omitempty"`
	/**
	* Total QUIC server UDP datagrams sent
	*/
	Quicserverdgrmsentrate float64 `json:"quicserverdgrmsentrate,omitempty"`
	Quiccurclientconn int `json:"quiccurclientconn,omitempty"`
	/**
	* Current QUIC client connections
	*/
	Quiccurclientconnrate float64 `json:"quiccurclientconnrate,omitempty"`
	Quiccurserverconn int `json:"quiccurserverconn,omitempty"`
	/**
	* Current QUIC server connections
	*/
	Quiccurserverconnrate float64 `json:"quiccurserverconnrate,omitempty"`
	Quiclocalconnid int `json:"quiclocalconnid,omitempty"`
	/**
	* Current QUIC local connection IDs allocated
	*/
	Quiclocalconnidrate float64 `json:"quiclocalconnidrate,omitempty"`
	Quiccursavedcryptoctx int `json:"quiccursavedcryptoctx,omitempty"`
	/**
	* Current QUIC crypto contexts allocated
	*/
	Quiccursavedcryptoctxrate float64 `json:"quiccursavedcryptoctxrate,omitempty"`
	Quictotclientconn int `json:"quictotclientconn,omitempty"`
	/**
	* Total QUIC client connections
	*/
	Quicclientconnrate float64 `json:"quicclientconnrate,omitempty"`
	Quictotserverconn int `json:"quictotserverconn,omitempty"`
	/**
	* Total QUIC server connections
	*/
	Quicserverconnrate float64 `json:"quicserverconnrate,omitempty"`
	Quicmigratedconn int `json:"quicmigratedconn,omitempty"`
	/**
	* Total number of migrated QUIC connections
	*/
	Quicmigratedconnrate float64 `json:"quicmigratedconnrate,omitempty"`
	Quicjumboframesrcvd int `json:"quicjumboframesrcvd,omitempty"`
	/**
	* Total number of QUIC jumbo frames received
	*/
	Quicjumboframesrcvdrate float64 `json:"quicjumboframesrcvdrate,omitempty"`
	Quicretrypktsent int `json:"quicretrypktsent,omitempty"`
	/**
	* Number of QUIC Retry packets sent
	*/
	Quicretrypktsentrate float64 `json:"quicretrypktsentrate,omitempty"`
	Quichandshakecmpltd int `json:"quichandshakecmpltd,omitempty"`
	/**
	* Number of QUIC handshake messages completed
	*/
	Quichandshakecmpltdrate float64 `json:"quichandshakecmpltdrate,omitempty"`
	Quictransptconnclosepktsent int `json:"quictransptconnclosepktsent,omitempty"`
	/**
	* Number of QUIC transport no-error Connection Close packets sent
	*/
	Quictransptconnclosepktsentrate float64 `json:"quictransptconnclosepktsentrate,omitempty"`
	Quicappconnclosepktsent int `json:"quicappconnclosepktsent,omitempty"`
	/**
	* Number of QUIC application no-error Connection Close packets sent
	*/
	Quicappconnclosepktsentrate float64 `json:"quicappconnclosepktsentrate,omitempty"`
	Quicasyncdatagramssent int `json:"quicasyncdatagramssent,omitempty"`
	/**
	* Number of QUIC datagrams sent using async mode
	*/
	Quicasyncdatagramssentrate float64 `json:"quicasyncdatagramssentrate,omitempty"`
	Quicconninfoalcfail int `json:"quicconninfoalcfail,omitempty"`
	/**
	* Quic session allocations failed
	*/
	Quicconninfoalcfailrate float64 `json:"quicconninfoalcfailrate,omitempty"`
	Quicnsbalcfail int `json:"quicnsbalcfail,omitempty"`
	/**
	* Quic NSB allocations failed
	*/
	Quicnsbalcfailrate float64 `json:"quicnsbalcfailrate,omitempty"`
	Quictlsalertsent int `json:"quictlsalertsent,omitempty"`
	/**
	* Total QUIC TLS 1.3 transport errors sent
	*/
	Quictlsalertsentrate float64 `json:"quictlsalertsentrate,omitempty"`
	Quicstlessconnclosepktsent int `json:"quicstlessconnclosepktsent,omitempty"`
	/**
	* Number of QUIC stateless Connection Close packets sent
	*/
	Quicstlessconnclosepktsentrate float64 `json:"quicstlessconnclosepktsentrate,omitempty"`
	Quicvernegpktsent int `json:"quicvernegpktsent,omitempty"`
	/**
	* Number of QUIC Version Negotiation packets sent
	*/
	Quicvernegpktsentrate float64 `json:"quicvernegpktsentrate,omitempty"`
	Quictransptconnclosepktfail int `json:"quictransptconnclosepktfail,omitempty"`
	/**
	* Number of QUIC transport error Connection Close packets sent
	*/
	Quictransptconnclosepktfailrate float64 `json:"quictransptconnclosepktfailrate,omitempty"`
	Quicappconnclosepktfail int `json:"quicappconnclosepktfail,omitempty"`
	/**
	* Number of QUIC application error Connection Close packets sent
	*/
	Quicappconnclosepktfailrate float64 `json:"quicappconnclosepktfailrate,omitempty"`
	Quicretrytokenverfail int `json:"quicretrytokenverfail,omitempty"`
	/**
	* Number of times QUIC Retry token verification failed
	*/
	Quicretrytokenverfailrate float64 `json:"quicretrytokenverfailrate,omitempty"`
	Quicnewtokenverfail int `json:"quicnewtokenverfail,omitempty"`
	/**
	* Number of times QUIC NEW_TOKEN token verification failed
	*/
	Quicnewtokenverfailrate float64 `json:"quicnewtokenverfailrate,omitempty"`

}
