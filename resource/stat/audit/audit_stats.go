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

package audit


type Auditstats struct {
	/**
	* Clear the statsistics / counters
	*/
	Clearstats string `json:"clearstats,omitempty"`
	Auditsyslogmsgsent uint64 `json:"auditsyslogmsgsent,omitempty"`
	/**
	* Syslog messages sent to the syslog server(s) over UDP.
	*/
	Auditsyslogmsgsentrate int32 `json:"auditsyslogmsgsentrate,omitempty"`
	Auditsyslogmsggen uint64 `json:"auditsyslogmsggen,omitempty"`
	/**
	* Syslog messages about to be sent to the syslog server.
	*/
	Auditsyslogmsggenrate int32 `json:"auditsyslogmsggenrate,omitempty"`
	Auditsyslogmsgsenttcp uint64 `json:"auditsyslogmsgsenttcp,omitempty"`
	/**
	* Syslog messages sent to the syslog server(s) over TCP.
	*/
	Auditsyslogmsgsenttcprate int32 `json:"auditsyslogmsgsenttcprate,omitempty"`
	Auditnsballocfail uint32 `json:"auditnsballocfail,omitempty"`
	/**
	* NAT allocation failed.
	*/
	Auditnsballocfailrate int32 `json:"auditnsballocfailrate,omitempty"`
	Auditlog32errsyslogallocnsbfail uint32 `json:"auditlog32errsyslogallocnsbfail,omitempty"`
	/**
	* Nsb allocation failed.
	*/
	Auditlog32errsyslogallocnsbfailrate int32 `json:"auditlog32errsyslogallocnsbfailrate,omitempty"`
	Auditmemallocfail uint32 `json:"auditmemallocfail,omitempty"`
	/**
	* Failures in allocation of Access Gateway context structure. When an Access Gateway session is established, the Citrix ADC creates an internal context structure , which identifies the user and the IP address from which the user has logged in.
	*/
	Auditmemallocfailrate int32 `json:"auditmemallocfailrate,omitempty"`
	Auditportallocfail uint32 `json:"auditportallocfail,omitempty"`
	/**
	* Number of times the Citrix ADC failed to allocate a port when sending a syslog message to the syslog server(s).
	*/
	Auditportallocfailrate int32 `json:"auditportallocfailrate,omitempty"`
	Auditcontextnotfound uint32 `json:"auditcontextnotfound,omitempty"`
	/**
	* Failures in finding the context structure for an Access Gateway session during attempts to send session-specific audit messages.
		During an Access Gateway session, audit messages related to the session are queued up in the auditlog buffer for transmission to the audit log server(s). If the session is killed before the messages are sent, the context structure allocated at session creation is removed. This structure is needed for sending the queued auditlog messages. If it is not found, this counter is incremented.
	*/
	Auditcontextnotfoundrate int32 `json:"auditcontextnotfoundrate,omitempty"`
	Nsbchainallocfail uint32 `json:"nsbchainallocfail,omitempty"`
	/**
	* Nsb Chain allocation failed.
	*/
	Nsbchainallocfailrate int32 `json:"nsbchainallocfailrate,omitempty"`
	Clientconnfail uint32 `json:"clientconnfail,omitempty"`
	/**
	* Failures in establishment of a connection between the Citrix ADC and the auditserver tool (the Citrix ADC's custom logging tool).
	*/
	Clientconnfailrate int32 `json:"clientconnfailrate,omitempty"`
	Flushcmdcnt uint32 `json:"flushcmdcnt,omitempty"`
	/**
	* Auditlog buffer flushes. In a multiprocessor Citrix ADC, both the main processor and the co-processor can generate auditlog messages and fill up the auditlog buffers. But only the primary processor can free up the buffers by sending auditlog messages to the auditlog server(s). The number of auditlog buffers is fixed. If the co-processor detects that all the auditlog buffers are full, it issues a flush command to the main processor.
	*/
	Flushcmdcntrate int32 `json:"flushcmdcntrate,omitempty"`
	Systcpconnfail uint32 `json:"systcpconnfail,omitempty"`
	/**
	* Failures in establishment of a connection between the Citrix ADC and the syslog server.
	*/
	Systcpconnfailrate int32 `json:"systcpconnfailrate,omitempty"`
	Logunsentlbsys uint32 `json:"logunsentlbsys,omitempty"`
	/**
	* Total auditlog messages which are not delivered to load balanced syslog servers
	*/
	Logunsentlbsysrate int32 `json:"logunsentlbsysrate,omitempty"`
	Logsdropped uint64 `json:"logsdropped,omitempty"`
	/**
	* Total number of log messages dropped by Citrix ADC after max hold limit is reached
	*/
	Logsdroppedrate int32 `json:"logsdroppedrate,omitempty"`
	Logsdroppedtxminnsbs uint64 `json:"logsdroppedtxminnsbs,omitempty"`
	/**
	* Total number of log messages dropped by Citrix ADC when NSBQ length is less than TX min NSBs
	*/
	Logsdroppedtxminnsbsrate int32 `json:"logsdroppedtxminnsbsrate,omitempty"`

}
