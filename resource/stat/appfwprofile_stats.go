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
* Statistics for application firewall profile resource.
 */

type Appfwprofilestats struct {
	/**
	* Name of the application firewall profile.
	 */
	Name string `json:"name,omitempty"`
	/**
	* Clear the statsistics / counters
	 */
	Clearstats                    string `json:"clearstats,omitempty"`
	Appfirewallrequestsperprofile int    `json:"appfirewallrequestsperprofile,omitempty"`
	/**
	* HTTP/HTTPS requests sent to your protected web servers via the Application Firewall.
	 */
	Appfirewallrequestsperprofilerate float64 `json:"appfirewallrequestsperprofilerate,omitempty"`
	Appfirewallreqbytesperprofile     int     `json:"appfirewallreqbytesperprofile,omitempty"`
	/**
	* Number of bytes transfered for requests
	 */
	Appfirewallreqbytesperprofilerate float64 `json:"appfirewallreqbytesperprofilerate,omitempty"`
	Appfirewallresponsesperprofile    int     `json:"appfirewallresponsesperprofile,omitempty"`
	/**
	* HTTP/HTTPS responses sent by your protected web servers via the Application Firewall.
	 */
	Appfirewallresponsesperprofilerate float64 `json:"appfirewallresponsesperprofilerate,omitempty"`
	Appfirewallresbytesperprofile      int     `json:"appfirewallresbytesperprofile,omitempty"`
	/**
	* Number of bytes transfered for responses
	 */
	Appfirewallresbytesperprofilerate float64 `json:"appfirewallresbytesperprofilerate,omitempty"`
	Appfirewallabortsperprofile       int     `json:"appfirewallabortsperprofile,omitempty"`
	/**
	* Incomplete HTTP/HTTPS requests aborted by the client before the Application Firewall could finish processing them.
	 */
	Appfirewallabortsperprofilerate float64 `json:"appfirewallabortsperprofilerate,omitempty"`
	Appfirewallredirectsperprofile  int     `json:"appfirewallredirectsperprofile,omitempty"`
	/**
	* HTTP/HTTPS requests redirected by the Application Firewall to a different Web page or web server. (HTTP 302)
	 */
	Appfirewallredirectsperprofilerate    float64 `json:"appfirewallredirectsperprofilerate,omitempty"`
	Appfirewalllongavgresptimeperprofile  int     `json:"appfirewalllongavgresptimeperprofile,omitempty"`
	Appfirewallshortavgresptimeperprofile int     `json:"appfirewallshortavgresptimeperprofile,omitempty"`
	Appfirewallviolstarturlperprofile     int     `json:"appfirewallviolstarturlperprofile,omitempty"`
	/**
	* Number of Start URL security check violations seen by the Application Firewall.
	 */
	Appfirewallviolstarturlperprofilerate float64 `json:"appfirewallviolstarturlperprofilerate,omitempty"`
	Appfirewallvioldenyurlperprofile      int     `json:"appfirewallvioldenyurlperprofile,omitempty"`
	/**
	* Number of Deny URL security check violations seen by the Application Firewall.
	 */
	Appfirewallvioldenyurlperprofilerate   float64 `json:"appfirewallvioldenyurlperprofilerate,omitempty"`
	Appfirewallviolrefererheaderperprofile int     `json:"appfirewallviolrefererheaderperprofile,omitempty"`
	/**
	* Number of Referer Header security check violations seen by the Application Firewall.
	 */
	Appfirewallviolrefererheaderperprofilerate float64 `json:"appfirewallviolrefererheaderperprofilerate,omitempty"`
	Appfirewallviolbufferoverflowperprofile    int     `json:"appfirewallviolbufferoverflowperprofile,omitempty"`
	/**
	* Number of Buffer Overflow security check violations seen by the Application Firewall.
	 */
	Appfirewallviolbufferoverflowperprofilerate  float64 `json:"appfirewallviolbufferoverflowperprofilerate,omitempty"`
	Appfirewallpostbodylimitviolationsperprofile int     `json:"appfirewallpostbodylimitviolationsperprofile,omitempty"`
	/**
	* Number of Post Body Limit security check violations seen by the Application Firewall.
	 */
	Appfirewallpostbodylimitviolationsperprofilerate float64 `json:"appfirewallpostbodylimitviolationsperprofilerate,omitempty"`
	Appfirewallviolcookieperprofile                  int     `json:"appfirewallviolcookieperprofile,omitempty"`
	/**
	* Number of Cookie Consistency security check violations seen by the Application Firewall.
	 */
	Appfirewallviolcookieperprofilerate   float64 `json:"appfirewallviolcookieperprofilerate,omitempty"`
	Appfirewallviolcookiehijackperprofile int     `json:"appfirewallviolcookiehijackperprofile,omitempty"`
	/**
	* Number of Cookie Hijacking security violations seen by the Application Firewall.
	 */
	Appfirewallviolcookiehijackperprofilerate float64 `json:"appfirewallviolcookiehijackperprofilerate,omitempty"`
	Appfirewallviolcsrftagperprofile          int     `json:"appfirewallviolcsrftagperprofile,omitempty"`
	/**
	* Number of Cross Site Request Forgery form tag security check violations seen by the Application Firewall.
	 */
	Appfirewallviolcsrftagperprofilerate float64 `json:"appfirewallviolcsrftagperprofilerate,omitempty"`
	Appfirewallviolxssperprofile         int     `json:"appfirewallviolxssperprofile,omitempty"`
	/**
	* Number of HTML Cross-Site Scripting security check violations seen by the Application Firewall.
	 */
	Appfirewallviolxssperprofilerate float64 `json:"appfirewallviolxssperprofilerate,omitempty"`
	Appfirewallviolsqlperprofile     int     `json:"appfirewallviolsqlperprofile,omitempty"`
	/**
	* Number of HTML SQL Injection security check violations seen by the Application Firewall.
	 */
	Appfirewallviolsqlperprofilerate     float64 `json:"appfirewallviolsqlperprofilerate,omitempty"`
	Appfirewallviolfieldformatperprofile int     `json:"appfirewallviolfieldformatperprofile,omitempty"`
	/**
	* Number of Field Format security check violations seen by the Application Firewall.
	 */
	Appfirewallviolfieldformatperprofilerate  float64 `json:"appfirewallviolfieldformatperprofilerate,omitempty"`
	Appfirewallviolfieldconsistencyperprofile int     `json:"appfirewallviolfieldconsistencyperprofile,omitempty"`
	/**
	* Number of Field Consistency security check violations seen by the Application Firewall.
	 */
	Appfirewallviolfieldconsistencyperprofilerate float64 `json:"appfirewallviolfieldconsistencyperprofilerate,omitempty"`
	Appfirewallviolcreditcardperprofile           int     `json:"appfirewallviolcreditcardperprofile,omitempty"`
	/**
	* Number of Credit Card security check violations seen by the Application Firewall.
	 */
	Appfirewallviolcreditcardperprofilerate float64 `json:"appfirewallviolcreditcardperprofilerate,omitempty"`
	Appfirewallviolsafeobjectperprofile     int     `json:"appfirewallviolsafeobjectperprofile,omitempty"`
	/**
	* Number of Safe Object security check violations seen by the Application Firewall.
	 */
	Appfirewallviolsafeobjectperprofilerate float64 `json:"appfirewallviolsafeobjectperprofilerate,omitempty"`
	Appfirewallviolsignatureperprofile      int     `json:"appfirewallviolsignatureperprofile,omitempty"`
	/**
	* Number of Signature violations seen by the Application Firewall.
	 */
	Appfirewallviolsignatureperprofilerate float64 `json:"appfirewallviolsignatureperprofilerate,omitempty"`
	Appfirewallviolcontenttypeperprofile   int     `json:"appfirewallviolcontenttypeperprofile,omitempty"`
	/**
	* Number of Content Type security check violations seen by the Application Firewall.
	 */
	Appfirewallviolcontenttypeperprofilerate float64 `json:"appfirewallviolcontenttypeperprofilerate,omitempty"`
	Appfirewallvioljsondosperprofile         int     `json:"appfirewallvioljsondosperprofile,omitempty"`
	/**
	* Number of JSON Denial-of-Service security check violations seen by the Application Firewall.
	 */
	Appfirewallvioljsondosperprofilerate float64 `json:"appfirewallvioljsondosperprofilerate,omitempty"`
	Appfirewallvioljsonsqlperprofile     int     `json:"appfirewallvioljsonsqlperprofile,omitempty"`
	/**
	* Number of JSON SQL Injection security check violations seen by the Application Firewall.
	 */
	Appfirewallvioljsonsqlperprofilerate float64 `json:"appfirewallvioljsonsqlperprofilerate,omitempty"`
	Appfirewallvioljsonxssperprofile     int     `json:"appfirewallvioljsonxssperprofile,omitempty"`
	/**
	* Number of JSON Cross-Site Scripting (XSS) security check violations seen by the Application Firewall.
	 */
	Appfirewallvioljsonxssperprofilerate float64 `json:"appfirewallvioljsonxssperprofilerate,omitempty"`
	Appfirewallvioljsoncmdperprofile     int     `json:"appfirewallvioljsoncmdperprofile,omitempty"`
	/**
	* Number of JSON Command Injection security check violations seen by the Application Firewall.
	 */
	Appfirewallvioljsoncmdperprofilerate     float64 `json:"appfirewallvioljsoncmdperprofilerate,omitempty"`
	Appfirewallviolfileuploadtypesperprofile int     `json:"appfirewallviolfileuploadtypesperprofile,omitempty"`
	/**
	* Number of Field Upload Types security check violations seen by the Application Firewall.
	 */
	Appfirewallviolfileuploadtypesperprofilerate       float64 `json:"appfirewallviolfileuploadtypesperprofilerate,omitempty"`
	Appfirewallxmlpayloadcontenttypemismatchperprofile int     `json:"appfirewallxmlpayloadcontenttypemismatchperprofile,omitempty"`
	/**
	* Number of Mismatched Content-Type in request with XML Payload security check violations seen by the Application Firewall.
	 */
	Appfirewallxmlpayloadcontenttypemismatchperprofilerate float64 `json:"appfirewallxmlpayloadcontenttypemismatchperprofilerate,omitempty"`
	Appfirewallviolcmdperprofile                           int     `json:"appfirewallviolcmdperprofile,omitempty"`
	/**
	* Number of HTML Command Injection security check violations seen by the Application Firewall.
	 */
	Appfirewallviolcmdperprofilerate                  float64 `json:"appfirewallviolcmdperprofilerate,omitempty"`
	Appfirewallviolwellformednessviolationsperprofile int     `json:"appfirewallviolwellformednessviolationsperprofile,omitempty"`
	/**
	* Number of XML Format security check violations seen by the Application Firewall.
	 */
	Appfirewallviolwellformednessviolationsperprofilerate float64 `json:"appfirewallviolwellformednessviolationsperprofilerate,omitempty"`
	Appfirewallviolxdosviolationsperprofile               int     `json:"appfirewallviolxdosviolationsperprofile,omitempty"`
	/**
	* Number of XML Denial-of-Service security check violations seen by the Application Firewall.
	 */
	Appfirewallviolxdosviolationsperprofilerate float64 `json:"appfirewallviolxdosviolationsperprofilerate,omitempty"`
	Appfirewallviolmsgvalviolationsperprofile   int     `json:"appfirewallviolmsgvalviolationsperprofile,omitempty"`
	/**
	* Number of XML Message Validation security check violations seen by the Application Firewall.
	 */
	Appfirewallviolmsgvalviolationsperprofilerate float64 `json:"appfirewallviolmsgvalviolationsperprofilerate,omitempty"`
	Appfirewallviolwsiviolationsperprofile        int     `json:"appfirewallviolwsiviolationsperprofile,omitempty"`
	/**
	* Number of Web Services Interoperability (WS-I) security check violations seen by the Application Firewall.
	 */
	Appfirewallviolwsiviolationsperprofilerate float64 `json:"appfirewallviolwsiviolationsperprofilerate,omitempty"`
	Appfirewallviolxmlsqlviolationsperprofile  int     `json:"appfirewallviolxmlsqlviolationsperprofile,omitempty"`
	/**
	* Number of XML SQL Injection security check violations seen by the Application Firewall.
	 */
	Appfirewallviolxmlsqlviolationsperprofilerate float64 `json:"appfirewallviolxmlsqlviolationsperprofilerate,omitempty"`
	Appfirewallviolxmlxssviolationsperprofile     int     `json:"appfirewallviolxmlxssviolationsperprofile,omitempty"`
	/**
	* Number of XML Cross-Site Scripting (XSS) security check violations seen by the Application Firewall.
	 */
	Appfirewallviolxmlxssviolationsperprofilerate    float64 `json:"appfirewallviolxmlxssviolationsperprofilerate,omitempty"`
	Appfirewallviolxmlattachmentviolationsperprofile int     `json:"appfirewallviolxmlattachmentviolationsperprofile,omitempty"`
	/**
	* Number of XML Attachment security check violations seen by the Application Firewall.
	 */
	Appfirewallviolxmlattachmentviolationsperprofilerate float64 `json:"appfirewallviolxmlattachmentviolationsperprofilerate,omitempty"`
	Appfirewallviolxmlsoapfaultviolationsperprofile      int     `json:"appfirewallviolxmlsoapfaultviolationsperprofile,omitempty"`
	/**
	* Number of requests returning soap:fault from the backend server
	 */
	Appfirewallviolxmlsoapfaultviolationsperprofilerate float64 `json:"appfirewallviolxmlsoapfaultviolationsperprofilerate,omitempty"`
	Appfirewallviolxmlgenericviolationsperprofile       int     `json:"appfirewallviolxmlgenericviolationsperprofile,omitempty"`
	/**
	* Number of requests returning XML generic violation from the backend server
	 */
	Appfirewallviolxmlgenericviolationsperprofilerate float64 `json:"appfirewallviolxmlgenericviolationsperprofilerate,omitempty"`
	Appfirewallviolsqlgramperprofile                  int     `json:"appfirewallviolsqlgramperprofile,omitempty"`
	/**
	* Number of HTML SQL Injection security check violations (reported by SQL grammar) seen by the Application Firewall.
	 */
	Appfirewallviolsqlgramperprofilerate float64 `json:"appfirewallviolsqlgramperprofilerate,omitempty"`
	Appfirewallvioljsonsqlgramperprofile int     `json:"appfirewallvioljsonsqlgramperprofile,omitempty"`
	/**
	* Number of JSON SQL Injection security check violations (reported by SQL grammar) seen by the Application Firewall.
	 */
	Appfirewallvioljsonsqlgramperprofilerate float64 `json:"appfirewallvioljsonsqlgramperprofilerate,omitempty"`
	Appfirewalltotalviolperprofile           int     `json:"appfirewalltotalviolperprofile,omitempty"`
	/**
	* Number of violations seen by the application firewall on per profile basis
	 */
	Appfirewallviolperprofilerate    float64 `json:"appfirewallviolperprofilerate,omitempty"`
	Appfirewalllogstarturlperprofile int     `json:"appfirewalllogstarturlperprofile,omitempty"`
	/**
	* Number of Start URL security check log messages generated by the Application Firewall.
	 */
	Appfirewalllogstarturlperprofilerate float64 `json:"appfirewalllogstarturlperprofilerate,omitempty"`
	Appfirewalllogdenyurlperprofile      int     `json:"appfirewalllogdenyurlperprofile,omitempty"`
	/**
	* Number of Deny URL security check log messages generated by the Application Firewall.
	 */
	Appfirewalllogdenyurlperprofilerate   float64 `json:"appfirewalllogdenyurlperprofilerate,omitempty"`
	Appfirewalllogrefererheaderperprofile int     `json:"appfirewalllogrefererheaderperprofile,omitempty"`
	/**
	* Number of Referer Header security check log messages generated by the Application Firewall.
	 */
	Appfirewalllogrefererheaderperprofilerate float64 `json:"appfirewalllogrefererheaderperprofilerate,omitempty"`
	Appfirewalllogbufferoverflowperprofile    int     `json:"appfirewalllogbufferoverflowperprofile,omitempty"`
	/**
	* Number of Buffer Overflow security check log messages generated by the Application Firewall.
	 */
	Appfirewalllogbufferoverflowperprofilerate float64 `json:"appfirewalllogbufferoverflowperprofilerate,omitempty"`
	Appfirewallpostbodylimitlogsperprofile     int     `json:"appfirewallpostbodylimitlogsperprofile,omitempty"`
	/**
	* Number of Post Body Limit security check logs seen by the Application Firewall.
	 */
	Appfirewallpostbodylimitlogsperprofilerate float64 `json:"appfirewallpostbodylimitlogsperprofilerate,omitempty"`
	Appfirewalllogcookieperprofile             int     `json:"appfirewalllogcookieperprofile,omitempty"`
	/**
	* Number of Cookie Consistency security check log messages generated by the Application Firewall.
	 */
	Appfirewalllogcookieperprofilerate   float64 `json:"appfirewalllogcookieperprofilerate,omitempty"`
	Appfirewalllogcookiehijackperprofile int     `json:"appfirewalllogcookiehijackperprofile,omitempty"`
	/**
	* Number of Cookie Hijacking security violation log messages generated by the Application Firewall.
	 */
	Appfirewalllogcookiehijackperprofilerate float64 `json:"appfirewalllogcookiehijackperprofilerate,omitempty"`
	Appfirewalllogcsrftagperprofile          int     `json:"appfirewalllogcsrftagperprofile,omitempty"`
	/**
	* Number of Cross Site Request Forgery form tag security check log messages generated by the Application Firewall.
	 */
	Appfirewalllogcsrftagperprofilerate float64 `json:"appfirewalllogcsrftagperprofilerate,omitempty"`
	Appfirewalllogxssperprofile         int     `json:"appfirewalllogxssperprofile,omitempty"`
	/**
	* Number of HTML Cross-Site Scripting security check log messages generated by the Application Firewall.
	 */
	Appfirewalllogxssperprofilerate  float64 `json:"appfirewalllogxssperprofilerate,omitempty"`
	Appfirewalllogxformxssperprofile int     `json:"appfirewalllogxformxssperprofile,omitempty"`
	/**
	* Number of HTML Cross-Site Scripting security check transform log messages generated by the Application Firewall.
	 */
	Appfirewalllogxformxssperprofilerate float64 `json:"appfirewalllogxformxssperprofilerate,omitempty"`
	Appfirewalllogsqlperprofile          int     `json:"appfirewalllogsqlperprofile,omitempty"`
	/**
	* Number of HTML SQL Injection security check log messages generated by the Application Firewall.
	 */
	Appfirewalllogsqlperprofilerate  float64 `json:"appfirewalllogsqlperprofilerate,omitempty"`
	Appfirewalllogxformsqlperprofile int     `json:"appfirewalllogxformsqlperprofile,omitempty"`
	/**
	* Number of HTML SQL Injection security check transform log messages generated by the Application Firewall.
	 */
	Appfirewalllogxformsqlperprofilerate float64 `json:"appfirewalllogxformsqlperprofilerate,omitempty"`
	Appfirewalllogfieldformatperprofile  int     `json:"appfirewalllogfieldformatperprofile,omitempty"`
	/**
	* Number of Field Format security check log messages generated by the Application Firewall.
	 */
	Appfirewalllogfieldformatperprofilerate  float64 `json:"appfirewalllogfieldformatperprofilerate,omitempty"`
	Appfirewalllogfieldconsistencyperprofile int     `json:"appfirewalllogfieldconsistencyperprofile,omitempty"`
	/**
	* Number of Field Consistency security check log messages generated by the Application Firewall.
	 */
	Appfirewalllogfieldconsistencyperprofilerate float64 `json:"appfirewalllogfieldconsistencyperprofilerate,omitempty"`
	Appfirewalllogcreditcardperprofile           int     `json:"appfirewalllogcreditcardperprofile,omitempty"`
	/**
	* Number of Credit Card security check log messages generated by the Application Firewall.
	 */
	Appfirewalllogcreditcardperprofilerate  float64 `json:"appfirewalllogcreditcardperprofilerate,omitempty"`
	Appfirewallxformlogcreditcardperprofile int     `json:"appfirewallxformlogcreditcardperprofile,omitempty"`
	/**
	* Number of Credit Card security check transform log messages generated by the Application Firewall.
	 */
	Appfirewallxformlogcreditcardperprofilerate float64 `json:"appfirewallxformlogcreditcardperprofilerate,omitempty"`
	Appfirewalllogsafeobjectperprofile          int     `json:"appfirewalllogsafeobjectperprofile,omitempty"`
	/**
	* Number of Safe Object security check log messages generated by the Application Firewall.
	 */
	Appfirewalllogsafeobjectperprofilerate float64 `json:"appfirewalllogsafeobjectperprofilerate,omitempty"`
	Appfirewalllogcontenttypeperprofile    int     `json:"appfirewalllogcontenttypeperprofile,omitempty"`
	/**
	* Number of Content type security check log messages generated by the Application Firewall.
	 */
	Appfirewalllogcontenttypeperprofilerate float64 `json:"appfirewalllogcontenttypeperprofilerate,omitempty"`
	Appfirewalllogsjsondosperprofile        int     `json:"appfirewalllogsjsondosperprofile,omitempty"`
	/**
	* Number of JSON Denial-of-Service security check log messages generated by the Application Firewall.
	 */
	Appfirewalllogsjsondosperprofilerate float64 `json:"appfirewalllogsjsondosperprofilerate,omitempty"`
	Appfirewalllogsjsonsqlperprofile     int     `json:"appfirewalllogsjsonsqlperprofile,omitempty"`
	/**
	* Number of JSON SQL Injection security check log messages generated by the Application Firewall.
	 */
	Appfirewalllogsjsonsqlperprofilerate float64 `json:"appfirewalllogsjsonsqlperprofilerate,omitempty"`
	Appfirewalllogsjsonxssperprofile     int     `json:"appfirewalllogsjsonxssperprofile,omitempty"`
	/**
	* Number of JSON Cross-Site Scripting (XSS) security check log messages generated by the Application Firewall.
	 */
	Appfirewalllogsjsonxssperprofilerate float64 `json:"appfirewalllogsjsonxssperprofilerate,omitempty"`
	Appfirewalllogsjsoncmdperprofile     int     `json:"appfirewalllogsjsoncmdperprofile,omitempty"`
	/**
	* Number of JSON Command Injection security check log messages generated by the Application Firewall.
	 */
	Appfirewalllogsjsoncmdperprofilerate    float64 `json:"appfirewalllogsjsoncmdperprofilerate,omitempty"`
	Appfirewalllogfileuploadtypesperprofile int     `json:"appfirewalllogfileuploadtypesperprofile,omitempty"`
	/**
	* Number of File Upload Types security check log messages generated by the Application Firewall.
	 */
	Appfirewalllogfileuploadtypesperprofilerate        float64 `json:"appfirewalllogfileuploadtypesperprofilerate,omitempty"`
	Appfirewallloginfercontenttypexmlpayloadperprofile int     `json:"appfirewallloginfercontenttypexmlpayloadperprofile,omitempty"`
	/**
	* Number of Mismatched Content-Type in request with XML Payload security check logs seen by the Application Firewall.
	 */
	Appfirewallloginfercontenttypexmlpayloadperprofilerate float64 `json:"appfirewallloginfercontenttypexmlpayloadperprofilerate,omitempty"`
	Appfirewalllogcmdperprofile                            int     `json:"appfirewalllogcmdperprofile,omitempty"`
	/**
	* Number of HTML Command Injection security check log messages generated by the Application Firewall.
	 */
	Appfirewalllogcmdperprofilerate         float64 `json:"appfirewalllogcmdperprofilerate,omitempty"`
	Appfirewallwellformednesslogsperprofile int     `json:"appfirewallwellformednesslogsperprofile,omitempty"`
	/**
	* Number of XML Format security check log messages generated by the Application Firewall.
	 */
	Appfirewallwellformednesslogsperprofilerate float64 `json:"appfirewallwellformednesslogsperprofilerate,omitempty"`
	Appfirewallxdoslogsperprofile               int     `json:"appfirewallxdoslogsperprofile,omitempty"`
	/**
	* Number of XML Denial-of-Service security check log messages generated by the Application Firewall.
	 */
	Appfirewallxdoslogsperprofilerate float64 `json:"appfirewallxdoslogsperprofilerate,omitempty"`
	Appfirewallmsgvallogsperprofile   int     `json:"appfirewallmsgvallogsperprofile,omitempty"`
	/**
	* Number of XML Message Validation security check log messages generated by the Application Firewall.
	 */
	Appfirewallmsgvallogsperprofilerate float64 `json:"appfirewallmsgvallogsperprofilerate,omitempty"`
	Appfirewallwsilogsperprofile        int     `json:"appfirewallwsilogsperprofile,omitempty"`
	/**
	* Number of Web Services Interoperability (WS-I) security check log messages generated by the Application Firewall.
	 */
	Appfirewallwsilogsperprofilerate float64 `json:"appfirewallwsilogsperprofilerate,omitempty"`
	Appfirewallxmlsqllogsperprofile  int     `json:"appfirewallxmlsqllogsperprofile,omitempty"`
	/**
	* Number of XML SQL Injection security check log messages generated by the Application Firewall.
	 */
	Appfirewallxmlsqllogsperprofilerate float64 `json:"appfirewallxmlsqllogsperprofilerate,omitempty"`
	Appfirewallxmlxsslogsperprofile     int     `json:"appfirewallxmlxsslogsperprofile,omitempty"`
	/**
	* Number of XML Cross-Site Scripting (XSS) security check log messages generated by the Application Firewall.
	 */
	Appfirewallxmlxsslogsperprofilerate    float64 `json:"appfirewallxmlxsslogsperprofilerate,omitempty"`
	Appfirewallxmlattachmentlogsperprofile int     `json:"appfirewallxmlattachmentlogsperprofile,omitempty"`
	/**
	* Number of XML Attachment security check log messages generated by the Application Firewall.
	 */
	Appfirewallxmlattachmentlogsperprofilerate float64 `json:"appfirewallxmlattachmentlogsperprofilerate,omitempty"`
	Appfirewallxmlsoapfaultlogsperprofile      int     `json:"appfirewallxmlsoapfaultlogsperprofile,omitempty"`
	/**
	* Number of requests generating soap:fault log messages
	 */
	Appfirewallxmlsoapfaultlogsperprofilerate float64 `json:"appfirewallxmlsoapfaultlogsperprofilerate,omitempty"`
	Appfirewallxmlgenericlogsperprofile       int     `json:"appfirewallxmlgenericlogsperprofile,omitempty"`
	/**
	* Number of requests generating XML Generic log messages
	 */
	Appfirewallxmlgenericlogsperprofilerate float64 `json:"appfirewallxmlgenericlogsperprofilerate,omitempty"`
	Appfirewalllogsqlgramperprofile         int     `json:"appfirewalllogsqlgramperprofile,omitempty"`
	/**
	* Number of HTML SQL Injection security check log messages (reported by SQL grammar) generated by the Application Firewall.
	 */
	Appfirewalllogsqlgramperprofilerate  float64 `json:"appfirewalllogsqlgramperprofilerate,omitempty"`
	Appfirewalllogsjsonsqlgramperprofile int     `json:"appfirewalllogsjsonsqlgramperprofile,omitempty"`
	/**
	* Number of JSON SQL Injection security check log messages (reported by SQL grammar) generated by the Application Firewall.
	 */
	Appfirewalllogsjsonsqlgramperprofilerate float64 `json:"appfirewalllogsjsonsqlgramperprofilerate,omitempty"`
	Appfirewalltotallogperprofile            int     `json:"appfirewalltotallogperprofile,omitempty"`
	/**
	* Number of log messages generated by the application firewall on per profile basis
	 */
	Appfirewalllogperprofilerate float64 `json:"appfirewalllogperprofilerate,omitempty"`
	Appfirewallret4xxperprofile  int     `json:"appfirewallret4xxperprofile,omitempty"`
	/**
	* Number of requests returning HTTP 4xx from the backend server
	 */
	Appfirewallret4xxperprofilerate float64 `json:"appfirewallret4xxperprofilerate,omitempty"`
	Appfirewallret5xxperprofile     int     `json:"appfirewallret5xxperprofile,omitempty"`
	/**
	* Number of requests returning HTTP 5xx from the backend server
	 */
	Appfirewallret5xxperprofilerate float64 `json:"appfirewallret5xxperprofilerate,omitempty"`
}
