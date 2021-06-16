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
* Statistics for tcp protocol resource.
*/

type Protocoltcpstats struct {
	/**
	* Clear the statsistics / counters
	*/
	Clearstats string `json:"clearstats,omitempty"`
	Tcpactiveserverconn uint32 `json:"tcpactiveserverconn,omitempty"`
	Tcpcurserverconnopening uint32 `json:"tcpcurserverconnopening,omitempty"`
	Tcpcurclientconnopening uint32 `json:"tcpcurclientconnopening,omitempty"`
	Tcpcurclientconnestablished uint32 `json:"tcpcurclientconnestablished,omitempty"`
	Tcpcurserverconnestablished uint32 `json:"tcpcurserverconnestablished,omitempty"`
	Tcptotrxpkts uint64 `json:"tcptotrxpkts,omitempty"`
	/**
	* TCP packets received.
	*/
	Tcprxpktsrate float64 `json:"tcprxpktsrate,omitempty"`
	Tcptotrxbytes uint64 `json:"tcptotrxbytes,omitempty"`
	/**
	* Bytes of TCP data received.
	*/
	Tcprxbytesrate float64 `json:"tcprxbytesrate,omitempty"`
	Tcptottxpkts uint64 `json:"tcptottxpkts,omitempty"`
	/**
	* TCP packets transmitted.
	*/
	Tcptxpktsrate float64 `json:"tcptxpktsrate,omitempty"`
	Tcptottxbytes uint64 `json:"tcptottxbytes,omitempty"`
	/**
	* Bytes of TCP data transmitted.
	*/
	Tcptxbytesrate float64 `json:"tcptxbytesrate,omitempty"`
	Tcpcurclientconn uint32 `json:"tcpcurclientconn,omitempty"`
	Tcpcurclientconnclosing uint32 `json:"tcpcurclientconnclosing,omitempty"`
	Tcptotclientconnopened uint64 `json:"tcptotclientconnopened,omitempty"`
	/**
	* Client connections opened by the Citrix ADC since startup (after three-way handshake). This counter is reset when the Citrix ADC is restarted.
	*/
	Tcpclientconnopenedrate float64 `json:"tcpclientconnopenedrate,omitempty"`
	Tcpcurserverconn uint32 `json:"tcpcurserverconn,omitempty"`
	Tcpcurserverconnclosing uint32 `json:"tcpcurserverconnclosing,omitempty"`
	Tcptotserverconnopened uint64 `json:"tcptotserverconnopened,omitempty"`
	/**
	* Server connections initiated by the Citrix ADC since startup. This counter is reset when the Citrix ADC is restarted.
	*/
	Tcpserverconnopenedrate float64 `json:"tcpserverconnopenedrate,omitempty"`
	Tcpsurgequeuelen uint32 `json:"tcpsurgequeuelen,omitempty"`
	Tcpspareconn uint32 `json:"tcpspareconn,omitempty"`
	Tcptotzombiecltconnflushed uint64 `json:"tcptotzombiecltconnflushed,omitempty"`
	/**
	* Client connections that are flushed because the client has been idle for some time.
	*/
	Tcpzombiecltconnflushedrate float64 `json:"tcpzombiecltconnflushedrate,omitempty"`
	Tcptotzombiehalfopencltconnflushed uint64 `json:"tcptotzombiehalfopencltconnflushed,omitempty"`
	/**
	* Half-opened client connections that are flushed because the three-way handshakes are not complete.
	*/
	Tcpzombiehalfopencltconnflushedrate float64 `json:"tcpzombiehalfopencltconnflushedrate,omitempty"`
	Tcptotzombieactivehalfclosecltconnflushed uint64 `json:"tcptotzombieactivehalfclosecltconnflushed,omitempty"`
	/**
	* Active half-closed client connections that are flushed because the client has closed the connection and there has been no activity on the connection.
	*/
	Tcpzombieactivehalfclosecltconnflushedrate float64 `json:"tcpzombieactivehalfclosecltconnflushedrate,omitempty"`
	Tcptotzombiepassivehalfclosecltconnflushed uint64 `json:"tcptotzombiepassivehalfclosecltconnflushed,omitempty"`
	/**
	* Passive half-closed client connections that are flushed because the Citrix ADC has closed the connection and there has been no activity on the connection.
	*/
	Tcpzombiepassivehalfclosecltconnflushedrate float64 `json:"tcpzombiepassivehalfclosecltconnflushedrate,omitempty"`
	Tcptotzombiesvrconnflushed uint64 `json:"tcptotzombiesvrconnflushed,omitempty"`
	/**
	* Server connections that are flushed because there have been no client requests in the queue for some time.
	*/
	Tcpzombiesvrconnflushedrate float64 `json:"tcpzombiesvrconnflushedrate,omitempty"`
	Tcptotzombiehalfopensvrconnflushed uint64 `json:"tcptotzombiehalfopensvrconnflushed,omitempty"`
	/**
	* Half-opened server connections that are flushed because the three-way handshakes are not complete.
	*/
	Tcpzombiehalfopensvrconnflushedrate float64 `json:"tcpzombiehalfopensvrconnflushedrate,omitempty"`
	Tcptotzombieactivehalfclosesvrconnflushed uint64 `json:"tcptotzombieactivehalfclosesvrconnflushed,omitempty"`
	/**
	* Active half-closed server connections that are flushed because the server has closed the connection and there has been no activity on the connection.
	*/
	Tcpzombieactivehalfclosesvrconnflushedrate float64 `json:"tcpzombieactivehalfclosesvrconnflushedrate,omitempty"`
	Tcptotzombiepassivehalfclosesrvconnflushed uint64 `json:"tcptotzombiepassivehalfclosesrvconnflushed,omitempty"`
	/**
	* Passive half-closed server connections that are flushed because the Citrix ADC has closed the connection and there has been no activity on the connection.
	*/
	Tcpzombiepassivehalfclosesrvconnflushedrate float64 `json:"tcpzombiepassivehalfclosesrvconnflushedrate,omitempty"`
	Pcbtotzombiecall uint64 `json:"pcbtotzombiecall,omitempty"`
	/**
	* Times the Zombie cleanup function is called. Every time a connection is flushed, it is marked for cleanup. The Zombie cleanup function clears all these connections at predefined intervals.
	*/
	Pcbzombiecallrate float64 `json:"pcbzombiecallrate,omitempty"`
	Tcptotsyn uint64 `json:"tcptotsyn,omitempty"`
	/**
	* SYN packets received
	*/
	Tcpsynrate float64 `json:"tcpsynrate,omitempty"`
	Tcptotsynprobe uint64 `json:"tcptotsynprobe,omitempty"`
	/**
	* Probes from the Citrix ADC to a server. The Citrix ADC sends a SYN packet to the server to check its availability and expects a SYN_ACK packet from the server before a specified response timeout.
	*/
	Tcpsynproberate float64 `json:"tcpsynproberate,omitempty"`
	Tcptotsvrfin uint64 `json:"tcptotsvrfin,omitempty"`
	/**
	* FIN packets received from the server.
	*/
	Tcpsvrfinrate float64 `json:"tcpsvrfinrate,omitempty"`
	Tcptotcltfin uint64 `json:"tcptotcltfin,omitempty"`
	/**
	* FIN packets received from the clients.
	*/
	Tcpcltfinrate float64 `json:"tcpcltfinrate,omitempty"`
	Tcpwaittosyn uint64 `json:"tcpwaittosyn,omitempty"`
	/**
	* SYN packets (packets used to initiate a TCP connection) received on connections that are in the TIME_WAIT state. Packets cannot be transferred on a connection in this state.
	*/
	Tcpwaittosynrate float64 `json:"tcpwaittosynrate,omitempty"`
	Tcpwaittodata uint64 `json:"tcpwaittodata,omitempty"`
	/**
	* Bytes of data received on connections that are in the TIME_WAIT state. Data cannot be transferred on a connection that is in this state.
	*/
	Tcpwaittodatarate float64 `json:"tcpwaittodatarate,omitempty"`
	Tcptotsynheld uint64 `json:"tcptotsynheld,omitempty"`
	/**
	* SYN packets held on the Citrix ADC that are waiting for a server connection.
	*/
	Tcpsynheldrate float64 `json:"tcpsynheldrate,omitempty"`
	Tcptotsynflush uint64 `json:"tcptotsynflush,omitempty"`
	/**
	* SYN packets flushed on the Citrix ADC because of no response from the server for three or more seconds.
	*/
	Tcpsynflushrate float64 `json:"tcpsynflushrate,omitempty"`
	Tcptotfinwaitclosed uint64 `json:"tcptotfinwaitclosed,omitempty"`
	/**
	* Connections closed on the Citrix ADC because the number of connections in the TIME_WAIT state has exceeded the default value of 7000.
	*/
	Tcpfinwaitclosedrate float64 `json:"tcpfinwaitclosedrate,omitempty"`
	Tcperrbadchecksum uint64 `json:"tcperrbadchecksum,omitempty"`
	/**
	* Packets received with a TCP checksum error.
	*/
	Tcperrbadchecksumrate float64 `json:"tcperrbadchecksumrate,omitempty"`
	Tcperrdataafterfin uint64 `json:"tcperrdataafterfin,omitempty"`
	/**
	* Packets received following a connection termination request. This error is usually caused by a reordering of packets during transmission.
	*/
	Tcperrdataafterfinrate float64 `json:"tcperrdataafterfinrate,omitempty"`
	Tcperrsyninsynrcvd uint64 `json:"tcperrsyninsynrcvd,omitempty"`
	/**
	* SYN packets received on a connection that is in the SYN_RCVD state. A connection goes into the SYN_RCVD state after receiving a SYN packet.
	*/
	Tcperrsyninsynrcvdrate float64 `json:"tcperrsyninsynrcvdrate,omitempty"`
	Tcperrsyninest uint64 `json:"tcperrsyninest,omitempty"`
	/**
	* SYN packets received on a connection that is in the ESTABLISHED state. A SYN packet is not expected on an ESTABLISHED connection.
	*/
	Tcperrsyninestrate float64 `json:"tcperrsyninestrate,omitempty"`
	Tcperrsynsentbadack uint64 `json:"tcperrsynsentbadack,omitempty"`
	/**
	* Incorrect ACK packets received on a connection that is in the SYN_SENT state. An incorrect ACK packet is the third packet in the three-way handshake that has an incorrect sequence number.
	*/
	Tcperrsynsentbadackrate float64 `json:"tcperrsynsentbadackrate,omitempty"`
	Tcperrrst uint64 `json:"tcperrrst,omitempty"`
	/**
	* Reset packets received from a client or a server.
	*/
	Tcperrrstrate float64 `json:"tcperrrstrate,omitempty"`
	Tcperrrstnonest uint64 `json:"tcperrrstnonest,omitempty"`
	/**
	* Reset packets received on a connection that is not in the ESTABLISHED state.
	*/
	Tcperrrstnonestrate float64 `json:"tcperrrstnonestrate,omitempty"`
	Tcperrrstoutofwindow uint64 `json:"tcperrrstoutofwindow,omitempty"`
	/**
	* Reset packets received on a connection that is out of the current TCP window.
	*/
	Tcperrrstoutofwindowrate float64 `json:"tcperrrstoutofwindowrate,omitempty"`
	Tcperrrstintimewait uint64 `json:"tcperrrstintimewait,omitempty"`
	/**
	* Reset packets received on a connection that is in the TIME_WAIT state. Packets cannot be transferred on a connection in the TIME_WAIT state.
	*/
	Tcperrrstintimewaitrate float64 `json:"tcperrrstintimewaitrate,omitempty"`
	Tcperrsvroutoforder uint64 `json:"tcperrsvroutoforder,omitempty"`
	/**
	* Out of order TCP packets received from a server. 
	*/
	Tcperrsvroutoforderrate float64 `json:"tcperrsvroutoforderrate,omitempty"`
	Tcperrcltoutoforder uint64 `json:"tcperrcltoutoforder,omitempty"`
	/**
	* Out of order TCP packets received from a client.
	*/
	Tcperrcltoutoforderrate float64 `json:"tcperrcltoutoforderrate,omitempty"`
	Tcperrclthole uint64 `json:"tcperrclthole,omitempty"`
	/**
	* TCP holes created on a client connection. When out of order packets are received from a client, a hole is created on the Citrix ADC for each group of missing packets.
	*/
	Tcperrcltholerate float64 `json:"tcperrcltholerate,omitempty"`
	Tcperrsvrhole uint64 `json:"tcperrsvrhole,omitempty"`
	/**
	* TCP holes created on a server connection. When out of order packets are received from a server, a hole is created on the Citrix ADC for each group of missing packets.
	*/
	Tcperrsvrholerate float64 `json:"tcperrsvrholerate,omitempty"`
	Tcperrcookiepktseqreject uint64 `json:"tcperrcookiepktseqreject,omitempty"`
	/**
	* SYN cookie packets rejected because they contain an incorrect sequence number.
	*/
	Tcperrcookiepktseqrejectrate float64 `json:"tcperrcookiepktseqrejectrate,omitempty"`
	Tcperrcookiepktsigreject uint64 `json:"tcperrcookiepktsigreject,omitempty"`
	/**
	* SYN cookie packets rejected because they contain an incorrect signature.
	*/
	Tcperrcookiepktsigrejectrate float64 `json:"tcperrcookiepktsigrejectrate,omitempty"`
	Tcperrcookiepktseqdrop uint64 `json:"tcperrcookiepktseqdrop,omitempty"`
	/**
	* SYN cookie packets dropped because the sequence number specified in the packets is outside the current window.
	*/
	Tcperrcookiepktseqdroprate float64 `json:"tcperrcookiepktseqdroprate,omitempty"`
	Tcperrcookiepktmssreject uint64 `json:"tcperrcookiepktmssreject,omitempty"`
	/**
	* SYN cookie packets rejected because the maximum segment size (MSS) specified in the packets is incorrect.
	*/
	Tcperrcookiepktmssrejectrate float64 `json:"tcperrcookiepktmssrejectrate,omitempty"`
	Tcperranyportfail uint64 `json:"tcperranyportfail,omitempty"`
	/**
	* Port allocations that have failed on a mapped IP address because the maximum limit of 65536 has been exceeded.
	*/
	Tcperranyportfailrate float64 `json:"tcperranyportfailrate,omitempty"`
	Tcperripportfail uint64 `json:"tcperripportfail,omitempty"`
	/**
	* Port allocations that have failed on a subnet IP address or vserver IP address because the maximum limit of 65536 has been exceeded.
	*/
	Tcperripportfailrate float64 `json:"tcperripportfailrate,omitempty"`
	Tcperrstraypkt uint64 `json:"tcperrstraypkt,omitempty"`
	/**
	* Number of stray or misrouted packets.
	*/
	Tcperrstraypktrate float64 `json:"tcperrstraypktrate,omitempty"`
	Tcperrsentrst uint64 `json:"tcperrsentrst,omitempty"`
	/**
	* Reset packets sent to a client or a server.
	*/
	Tcperrsentrstrate float64 `json:"tcperrsentrstrate,omitempty"`
	Tcperrbadstateconn uint64 `json:"tcperrbadstateconn,omitempty"`
	/**
	* Connections that are not in a valid TCP state.
	*/
	Tcperrbadstateconnrate float64 `json:"tcperrbadstateconnrate,omitempty"`
	Tcperrrstthreshold uint64 `json:"tcperrrstthreshold,omitempty"`
	/**
	* Reset packets dropped because the default threshold of 100 resets per 10 milliseconds has been exceeded. This is a configurable value using the set rateControl command.
	*/
	Tcperrrstthresholdrate float64 `json:"tcperrrstthresholdrate,omitempty"`
	Tcperroutofwindowpkts uint64 `json:"tcperroutofwindowpkts,omitempty"`
	/**
	* Packets received that are out of the current advertised window.
	*/
	Tcperroutofwindowpktsrate float64 `json:"tcperroutofwindowpktsrate,omitempty"`
	Tcperrsyndroppedcongestion uint64 `json:"tcperrsyndroppedcongestion,omitempty"`
	/**
	* SYN packets dropped because of network congestion.
	*/
	Tcperrsyndroppedcongestionrate float64 `json:"tcperrsyndroppedcongestionrate,omitempty"`
	Tcperrcltretrasmit uint64 `json:"tcperrcltretrasmit,omitempty"`
	/**
	* Packets retransmitted by a client. This usually occurs because the acknowledgement from the Citrix ADC has not reached the client.
	*/
	Tcperrcltretrasmitrate float64 `json:"tcperrcltretrasmitrate,omitempty"`
	Tcperrfullretrasmit uint64 `json:"tcperrfullretrasmit,omitempty"`
	/**
	* Full packets retransmitted by the client or the server.
	*/
	Tcperrfullretrasmitrate float64 `json:"tcperrfullretrasmitrate,omitempty"`
	Tcperrsynretry uint64 `json:"tcperrsynretry,omitempty"`
	/**
	* SYN packets resent to a server.
	*/
	Tcperrsynretryrate float64 `json:"tcperrsynretryrate,omitempty"`
	Tcperrsyngiveup uint64 `json:"tcperrsyngiveup,omitempty"`
	/**
	* Attempts to establish a connection on the Citrix ADC that timed out.
	*/
	Tcperrsyngiveuprate float64 `json:"tcperrsyngiveuprate,omitempty"`
	Tcperrretransmit uint64 `json:"tcperrretransmit,omitempty"`
	/**
	* TCP packets retransmitted. The Citrix ADC attempts to retransmit the packet up to seven times, after which it resets the other half of the TCP connection.
	*/
	Tcperrretransmitrate float64 `json:"tcperrretransmitrate,omitempty"`
	Tcperrfirstretransmissions uint64 `json:"tcperrfirstretransmissions,omitempty"`
	/**
	* Packets retransmitted once by the Citrix ADC.
	*/
	Tcperrfirstretransmissionsrate float64 `json:"tcperrfirstretransmissionsrate,omitempty"`
	Tcperrthirdretransmissions uint64 `json:"tcperrthirdretransmissions,omitempty"`
	/**
	* Packets retransmitted three times by the Citrix ADC.
	*/
	Tcperrthirdretransmissionsrate float64 `json:"tcperrthirdretransmissionsrate,omitempty"`
	Tcperrfifthretransmissions uint64 `json:"tcperrfifthretransmissions,omitempty"`
	/**
	* Packets retransmitted five times by the Citrix ADC.
	*/
	Tcperrfifthretransmissionsrate float64 `json:"tcperrfifthretransmissionsrate,omitempty"`
	Tcperrseventhretransmissions uint64 `json:"tcperrseventhretransmissions,omitempty"`
	/**
	* Packets retransmitted seven times by the Citrix ADC. If this fails, the Citrix ADC terminates the connection.
	*/
	Tcperrseventhretransmissionsrate float64 `json:"tcperrseventhretransmissionsrate,omitempty"`
	Tcperrfastretransmissions uint64 `json:"tcperrfastretransmissions,omitempty"`
	/**
	* TCP packets on which the Citrix ADC performs a fast retransmission in response to three duplicate acknowledgements or a partial acknowledgement.  The Citrix ADC assumes that the packet is lost and retransmits the packet before its time-out.
	*/
	Tcperrfastretransmissionsrate float64 `json:"tcperrfastretransmissionsrate,omitempty"`
	Tcperrsvrretrasmit uint64 `json:"tcperrsvrretrasmit,omitempty"`
	/**
	* Packets retransmitted by a server. This usually occurs because the acknowledgement from the Citrix ADC has not reached the server.
	*/
	Tcperrsvrretrasmitrate float64 `json:"tcperrsvrretrasmitrate,omitempty"`
	Tcperrpartialretrasmit uint64 `json:"tcperrpartialretrasmit,omitempty"`
	/**
	* Partial packet retransmits by a client or server due to congestion on the connection. This usually occurs because the window advertised by the Citrix ADC is not big enough to hold the full packet.
	*/
	Tcperrpartialretrasmitrate float64 `json:"tcperrpartialretrasmitrate,omitempty"`
	Tcperrfinretry uint64 `json:"tcperrfinretry,omitempty"`
	/**
	* FIN packets resent to a server or a client.
	*/
	Tcperrfinretryrate float64 `json:"tcperrfinretryrate,omitempty"`
	Tcperrfingiveup uint64 `json:"tcperrfingiveup,omitempty"`
	/**
	* Connections that were timed out by the Citrix ADC because of not receiving the ACK packet after retransmitting the FIN packet four times.
	*/
	Tcperrfingiveuprate float64 `json:"tcperrfingiveuprate,omitempty"`
	Tcperrsecondretransmissions uint64 `json:"tcperrsecondretransmissions,omitempty"`
	/**
	* Packets retransmitted twice by the Citrix ADC.
	*/
	Tcperrsecondretransmissionsrate float64 `json:"tcperrsecondretransmissionsrate,omitempty"`
	Tcperrforthretransmissions uint64 `json:"tcperrforthretransmissions,omitempty"`
	/**
	* Packets retransmitted four times by the Citrix ADC.
	*/
	Tcperrforthretransmissionsrate float64 `json:"tcperrforthretransmissionsrate,omitempty"`
	Tcperrsixthretransmissions uint64 `json:"tcperrsixthretransmissions,omitempty"`
	/**
	* Packets retransmitted six times by the Citrix ADC.
	*/
	Tcperrsixthretransmissionsrate float64 `json:"tcperrsixthretransmissionsrate,omitempty"`
	Tcperrretransmitgiveup uint64 `json:"tcperrretransmitgiveup,omitempty"`
	/**
	* Number of times Citrix ADC terminates a connection after retransmitting the packet seven times on that connection.Retrasnmission happens when recieving end doesn't acknowledges the packet.
	*/
	Tcperrretransmitgiveuprate float64 `json:"tcperrretransmitgiveuprate,omitempty"`
	Tcperrcipalloc uint64 `json:"tcperrcipalloc,omitempty"`
	/**
	* Number of times TCP level client header insertion failure
	*/
	Tcperrcipallocrate float64 `json:"tcperrcipallocrate,omitempty"`

}
