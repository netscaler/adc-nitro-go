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
	Auditsyslogmsgsent int `json:"auditsyslogmsgsent,omitempty"`
	/**
	* Syslog messages sent to the syslog server(s) over UDP.
	*/
	Auditsyslogmsgsentrate float64 `json:"auditsyslogmsgsentrate,omitempty"`
	Auditsyslogmsggen int `json:"auditsyslogmsggen,omitempty"`
	/**
	* Syslog messages about to be sent to the syslog server.
	*/
	Auditsyslogmsggenrate float64 `json:"auditsyslogmsggenrate,omitempty"`
	Auditsyslogmsgsenttcp int `json:"auditsyslogmsgsenttcp,omitempty"`
	/**
	* Syslog messages sent to the syslog server(s) over TCP.
	*/
	Auditsyslogmsgsenttcprate float64 `json:"auditsyslogmsgsenttcprate,omitempty"`
	Auditjsonobjcreated int `json:"auditjsonobjcreated,omitempty"`
	/**
	* Total number of JSON objects created for HTTP based syslog servers
	*/
	Auditjsonobjcreatedrate float64 `json:"auditjsonobjcreatedrate,omitempty"`
	Auditjsonbufexported int `json:"auditjsonbufexported,omitempty"`
	/**
	* Total number of JSON buffers exported for HTTP based syslog servers
	*/
	Auditjsonbufexportedrate float64 `json:"auditjsonbufexportedrate,omitempty"`
	Auditjson2xxresp int `json:"auditjson2xxresp,omitempty"`
	/**
	* Total number of 2xx responses received for HTTP based syslog servers
	*/
	Auditjson2xxresprate float64 `json:"auditjson2xxresprate,omitempty"`
	Auditjsonnon2xxresp int `json:"auditjsonnon2xxresp,omitempty"`
	/**
	* Total number of non 2xx responses received for HTTP based syslog servers
	*/
	Auditjsonnon2xxresprate float64 `json:"auditjsonnon2xxresprate,omitempty"`
	Auditnsballocfail int `json:"auditnsballocfail,omitempty"`
	/**
	* NAT allocation failed.
	*/
	Auditnsballocfailrate float64 `json:"auditnsballocfailrate,omitempty"`
	Auditlog32errsyslogallocnsbfail int `json:"auditlog32errsyslogallocnsbfail,omitempty"`
	/**
	* Nsb allocation failed.
	*/
	Auditlog32errsyslogallocnsbfailrate float64 `json:"auditlog32errsyslogallocnsbfailrate,omitempty"`
	Auditmemallocfail int `json:"auditmemallocfail,omitempty"`
	/**
	* Failures in allocation of Access Gateway context structure. When an Access Gateway session is established, the Citrix ADC creates an internal context structure , which identifies the user and the IP address from which the user has logged in.
	*/
	Auditmemallocfailrate float64 `json:"auditmemallocfailrate,omitempty"`
	Auditportallocfail int `json:"auditportallocfail,omitempty"`
	/**
	* Number of times the Citrix ADC failed to allocate a port when sending a syslog message to the syslog server(s).
	*/
	Auditportallocfailrate float64 `json:"auditportallocfailrate,omitempty"`
	Auditcontextnotfound int `json:"auditcontextnotfound,omitempty"`
	/**
	* Failures in finding the context structure for an Access Gateway session during attempts to send session-specific audit messages.
		During an Access Gateway session, audit messages related to the session are queued up in the auditlog buffer for transmission to the audit log server(s). If the session is killed before the messages are sent, the context structure allocated at session creation is removed. This structure is needed for sending the queued auditlog messages. If it is not found, this counter is incremented.
	*/
	Auditcontextnotfoundrate float64 `json:"auditcontextnotfoundrate,omitempty"`
	Nsbchainallocfail int `json:"nsbchainallocfail,omitempty"`
	/**
	* Nsb Chain allocation failed.
	*/
	Nsbchainallocfailrate float64 `json:"nsbchainallocfailrate,omitempty"`
	Clientconnfail int `json:"clientconnfail,omitempty"`
	/**
	* Failures in establishment of a connection between the Citrix ADC and the auditserver tool (the Citrix ADC's custom logging tool).
	*/
	Clientconnfailrate float64 `json:"clientconnfailrate,omitempty"`
	Flushcmdcnt int `json:"flushcmdcnt,omitempty"`
	/**
	* Auditlog buffer flushes. In a multiprocessor Citrix ADC, both the main processor and the co-processor can generate auditlog messages and fill up the auditlog buffers. But only the primary processor can free up the buffers by sending auditlog messages to the auditlog server(s). The number of auditlog buffers is fixed. If the co-processor detects that all the auditlog buffers are full, it issues a flush command to the main processor.
	*/
	Flushcmdcntrate float64 `json:"flushcmdcntrate,omitempty"`
	Systcpconnfail int `json:"systcpconnfail,omitempty"`
	/**
	* Failures in establishment of a connection between the Citrix ADC and the syslog server.
	*/
	Systcpconnfailrate float64 `json:"systcpconnfailrate,omitempty"`
	Logunsentlbsys int `json:"logunsentlbsys,omitempty"`
	/**
	* Total auditlog messages which are not delivered to load balanced syslog servers
	*/
	Logunsentlbsysrate float64 `json:"logunsentlbsysrate,omitempty"`
	Logsdropped int `json:"logsdropped,omitempty"`
	/**
	* Total number of log messages dropped by Citrix ADC after max hold limit is reached
	*/
	Logsdroppedrate float64 `json:"logsdroppedrate,omitempty"`
	Logsdroppedtxminnsbs int `json:"logsdroppedtxminnsbs,omitempty"`
	/**
	* Total number of log messages dropped by Citrix ADC when NSBQ length is less than TX min NSBs
	*/
	Logsdroppedtxminnsbsrate float64 `json:"logsdroppedtxminnsbsrate,omitempty"`
	Auditjsonobjcreateerr int `json:"auditjsonobjcreateerr,omitempty"`
	/**
	* Total number of JSON objects creation error for HTTP based syslog servers
	*/
	Auditjsonobjcreateerrrate float64 `json:"auditjsonobjcreateerrrate,omitempty"`
	Auditjsonbufexporterr int `json:"auditjsonbufexporterr,omitempty"`
	/**
	* Total number of JSON buffers export error for HTTP based syslog servers
	*/
	Auditjsonbufexporterrrate float64 `json:"auditjsonbufexporterrrate,omitempty"`

}
