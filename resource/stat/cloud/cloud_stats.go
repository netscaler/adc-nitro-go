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

package cloud


type Cloudstats struct {
	/**
	* Clear the statsistics / counters
	*/
	Clearstats string `json:"clearstats,omitempty"`
	Ngsshieldle1secclxmtpconnesttime int `json:"ngsshieldle1secclxmtpconnesttime,omitempty"`
	/**
	* Total Connection with CLXMTP connect time less than 1 sec
	*/
	Ngsshieldle1secclxmtpconnesttimerate float64 `json:"ngsshieldle1secclxmtpconnesttimerate,omitempty"`
	Ngsshieldle2secclxmtpconnesttime int `json:"ngsshieldle2secclxmtpconnesttime,omitempty"`
	/**
	* Total Connection with CLXMTP connect time less than 2 sec
	*/
	Ngsshieldle2secclxmtpconnesttimerate float64 `json:"ngsshieldle2secclxmtpconnesttimerate,omitempty"`
	Ngsshieldle3secclxmtpconnesttime int `json:"ngsshieldle3secclxmtpconnesttime,omitempty"`
	/**
	* Total Connection with CLXMTP connect time less than 3 sec
	*/
	Ngsshieldle3secclxmtpconnesttimerate float64 `json:"ngsshieldle3secclxmtpconnesttimerate,omitempty"`
	Ngsshieldle5secclxmtpconnesttime int `json:"ngsshieldle5secclxmtpconnesttime,omitempty"`
	/**
	* Total Connection with CLXMTP connect time less than 5 sec
	*/
	Ngsshieldle5secclxmtpconnesttimerate float64 `json:"ngsshieldle5secclxmtpconnesttimerate,omitempty"`
	Ngsshieldle8secclxmtpconnesttime int `json:"ngsshieldle8secclxmtpconnesttime,omitempty"`
	/**
	* Total Connection with CLXMPT connect time less than 8 sec
	*/
	Ngsshieldle8secclxmtpconnesttimerate float64 `json:"ngsshieldle8secclxmtpconnesttimerate,omitempty"`
	Ngsshieldle13secclxmtpconnesttime int `json:"ngsshieldle13secclxmtpconnesttime,omitempty"`
	/**
	* Total Connection with CLXMPT connect time less than 13 sec
	*/
	Ngsshieldle13secclxmtpconnesttimerate float64 `json:"ngsshieldle13secclxmtpconnesttimerate,omitempty"`
	Ngsshieldle21secclxmtpconnesttime int `json:"ngsshieldle21secclxmtpconnesttime,omitempty"`
	/**
	* Total Connection with CLXMPT connect time less than 21 sec
	*/
	Ngsshieldle21secclxmtpconnesttimerate float64 `json:"ngsshieldle21secclxmtpconnesttimerate,omitempty"`
	Ngsshieldgt21secclxmtpconnesttime int `json:"ngsshieldgt21secclxmtpconnesttime,omitempty"`
	/**
	* Total Connection with CLXMPT connect time  greater than 21 sec
	*/
	Ngsshieldgt21secclxmtpconnesttimerate float64 `json:"ngsshieldgt21secclxmtpconnesttimerate,omitempty"`
	Ngsgctle1secconnesttime int `json:"ngsgctle1secconnesttime,omitempty"`
	/**
	* Total Connection with GCT connect time  less than 1 sec
	*/
	Ngsgctle1secconnesttimerate float64 `json:"ngsgctle1secconnesttimerate,omitempty"`
	Ngsgctle2secconnesttime int `json:"ngsgctle2secconnesttime,omitempty"`
	/**
	* Total Connection with GCT connect time  less than 2 sec
	*/
	Ngsgctle2secconnesttimerate float64 `json:"ngsgctle2secconnesttimerate,omitempty"`
	Ngsgctle3secconnesttime int `json:"ngsgctle3secconnesttime,omitempty"`
	/**
	* Total Connection with GCT connect time  less than 3 sec
	*/
	Ngsgctle3secconnesttimerate float64 `json:"ngsgctle3secconnesttimerate,omitempty"`
	Ngsgctle5secconnesttime int `json:"ngsgctle5secconnesttime,omitempty"`
	/**
	* Total Connection with GCT connect time  less than 5 sec
	*/
	Ngsgctle5secconnesttimerate float64 `json:"ngsgctle5secconnesttimerate,omitempty"`
	Ngsgctle8secconnesttime int `json:"ngsgctle8secconnesttime,omitempty"`
	/**
	* Total Connection with GCT connect time  less than 8 sec
	*/
	Ngsgctle8secconnesttimerate float64 `json:"ngsgctle8secconnesttimerate,omitempty"`
	Ngsgctgt8secconnesttime int `json:"ngsgctgt8secconnesttime,omitempty"`
	/**
	* Total Connection with GCT connect time greater than 8 sec
	*/
	Ngsgctgt8secconnesttimerate float64 `json:"ngsgctgt8secconnesttimerate,omitempty"`
	Ngsshieldclxmtpconnestattempted int `json:"ngsshieldclxmtpconnestattempted,omitempty"`
	/**
	* Total attempted CLXMTP connection
	*/
	Ngsshieldclxmtpconnestattemptedrate float64 `json:"ngsshieldclxmtpconnestattemptedrate,omitempty"`
	Ngsshieldclxmtpconnestsuccess int `json:"ngsshieldclxmtpconnestsuccess,omitempty"`
	/**
	* Total successful CLXMTP connection
	*/
	Ngsshieldclxmtpconnestsuccessrate float64 `json:"ngsshieldclxmtpconnestsuccessrate,omitempty"`
	Ngsshieldclxmtpsoftdeny int `json:"ngsshieldclxmtpsoftdeny,omitempty"`
	/**
	* Total number of softdeny received by SN
	*/
	Ngsshieldclxmtpsoftdenyrate float64 `json:"ngsshieldclxmtpsoftdenyrate,omitempty"`
	Ngsshieldclxmtpdeny int `json:"ngsshieldclxmtpdeny,omitempty"`
	/**
	* Total number of deny received by SN
	*/
	Ngsshieldclxmtpdenyrate float64 `json:"ngsshieldclxmtpdenyrate,omitempty"`
	Ngsshieldclxmtpallowtarget int `json:"ngsshieldclxmtpallowtarget,omitempty"`
	/**
	* Total number of allowtarget received by SN
	*/
	Ngsshieldclxmtpallowtargetrate float64 `json:"ngsshieldclxmtpallowtargetrate,omitempty"`
	Ngsshieldclxmtpredirecttarget int `json:"ngsshieldclxmtpredirecttarget,omitempty"`
	/**
	* Total number of redirecttarget received by SN
	*/
	Ngsshieldclxmtpredirecttargetrate float64 `json:"ngsshieldclxmtpredirecttargetrate,omitempty"`
	Ngsshieldclxmtpconnectorfanouttimeout int `json:"ngsshieldclxmtpconnectorfanouttimeout,omitempty"`
	/**
	* Total number of fanout timeout
	*/
	Ngsshieldclxmtpconnectorfanouttimeoutrate float64 `json:"ngsshieldclxmtpconnectorfanouttimeoutrate,omitempty"`
	Ngsshieldclxmtpglobalfanouttimeout int `json:"ngsshieldclxmtpglobalfanouttimeout,omitempty"`
	/**
	* Total number of global fanout timeout
	*/
	Ngsshieldclxmtpglobalfanouttimeoutrate float64 `json:"ngsshieldclxmtpglobalfanouttimeoutrate,omitempty"`
	Ngsshieldfanout int `json:"ngsshieldfanout,omitempty"`
	/**
	* Total number of fanout
	*/
	Ngsshieldfanoutrate float64 `json:"ngsshieldfanoutrate,omitempty"`
	Ngsshieldclvalidationfailure int `json:"ngsshieldclvalidationfailure,omitempty"`
	/**
	* Total Connection lease validation failure
	*/
	Ngsshieldclvalidationfailurerate float64 `json:"ngsshieldclvalidationfailurerate,omitempty"`
	Ngsshieldcldecryptionfailure int `json:"ngsshieldcldecryptionfailure,omitempty"`
	/**
	* Total Connection lease decryption failure
	*/
	Ngsshieldcldecryptionfailurerate float64 `json:"ngsshieldcldecryptionfailurerate,omitempty"`
	Ngsshieldclxmtpcomplete int `json:"ngsshieldclxmtpcomplete,omitempty"`
	/**
	* Total CLXMTP completed
	*/
	Ngsshieldclxmtpcompleterate float64 `json:"ngsshieldclxmtpcompleterate,omitempty"`
	Ngsshieldclxmtptrustestfailure int `json:"ngsshieldclxmtptrustestfailure,omitempty"`
	/**
	* Total CLXMTP trust establishment failure
	*/
	Ngsshieldclxmtptrustestfailurerate float64 `json:"ngsshieldclxmtptrustestfailurerate,omitempty"`
	Ngsshieldclxmtppubkeyvalfailure int `json:"ngsshieldclxmtppubkeyvalfailure,omitempty"`
	/**
	* Total CLXMTP pub key validation failure
	*/
	Ngsshieldclxmtppubkeyvalfailurerate float64 `json:"ngsshieldclxmtppubkeyvalfailurerate,omitempty"`
	Ngsshieldclxmtpconnvalidationfailure int `json:"ngsshieldclxmtpconnvalidationfailure,omitempty"`
	/**
	* Total CLXMTP connection lease validation failure
	*/
	Ngsshieldclxmtpconnvalidationfailurerate float64 `json:"ngsshieldclxmtpconnvalidationfailurerate,omitempty"`
	Ngsshieldclxmtpauthorizeconnreqfailure int `json:"ngsshieldclxmtpauthorizeconnreqfailure,omitempty"`
	/**
	* Total CLXMTP authorize connection request failure
	*/
	Ngsshieldclxmtpauthorizeconnreqfailurerate float64 `json:"ngsshieldclxmtpauthorizeconnreqfailurerate,omitempty"`
	Ngsshieldkeyfetchfailure int `json:"ngsshieldkeyfetchfailure,omitempty"`
	/**
	* Total Shield key fetching failure
	*/
	Ngsshieldkeyfetchfailurerate float64 `json:"ngsshieldkeyfetchfailurerate,omitempty"`
	Ngsshielddemandedrtkeynotfound int `json:"ngsshielddemandedrtkeynotfound,omitempty"`
	/**
	* Total root of trust keys not found failure
	*/
	Ngsshielddemandedrtkeynotfoundrate float64 `json:"ngsshielddemandedrtkeynotfoundrate,omitempty"`
	Ngsshielddemandedcliskeynotfound int `json:"ngsshielddemandedcliskeynotfound,omitempty"`
	/**
	* Total CLIS keys not found failure
	*/
	Ngsshielddemandedcliskeynotfoundrate float64 `json:"ngsshielddemandedcliskeynotfoundrate,omitempty"`
	Ngsshielddemandedgwkeynotfound int `json:"ngsshielddemandedgwkeynotfound,omitempty"`
	/**
	* Total gateway keys not found failure
	*/
	Ngsshielddemandedgwkeynotfoundrate float64 `json:"ngsshielddemandedgwkeynotfoundrate,omitempty"`
	Ngsshieldkeycacheempty int `json:"ngsshieldkeycacheempty,omitempty"`
	/**
	* Total shield key cache list empty
	*/
	Ngsshieldkeycacheemptyrate float64 `json:"ngsshieldkeycacheemptyrate,omitempty"`
	Ngsshieldakvgetsecretlistfailure int `json:"ngsshieldakvgetsecretlistfailure,omitempty"`
	/**
	* Total failures to get akv secret list
	*/
	Ngsshieldakvgetsecretlistfailurerate float64 `json:"ngsshieldakvgetsecretlistfailurerate,omitempty"`
	Ngsshieldakvgetsecretfailure int `json:"ngsshieldakvgetsecretfailure,omitempty"`
	/**
	* Total failures to get akv secrets
	*/
	Ngsshieldakvgetsecretfailurerate float64 `json:"ngsshieldakvgetsecretfailurerate,omitempty"`
	Ngsshieldakvtokenfailure int `json:"ngsshieldakvtokenfailure,omitempty"`
	/**
	* Total akv token failures
	*/
	Ngsshieldakvtokenfailurerate float64 `json:"ngsshieldakvtokenfailurerate,omitempty"`
	Ngsshieldawssmgetmastersecretfailure int `json:"ngsshieldawssmgetmastersecretfailure,omitempty"`
	/**
	* Total failures to get aws master secret list
	*/
	Ngsshieldawssmgetmastersecretfailurerate float64 `json:"ngsshieldawssmgetmastersecretfailurerate,omitempty"`
	Ngsshieldawssmgetsecretfailure int `json:"ngsshieldawssmgetsecretfailure,omitempty"`
	/**
	* Total failures to get aws secrets
	*/
	Ngsshieldawssmgetsecretfailurerate float64 `json:"ngsshieldawssmgetsecretfailurerate,omitempty"`
	Ngsshieldawssmtokenfailure int `json:"ngsshieldawssmtokenfailure,omitempty"`
	/**
	* Total aws secret manager token failures
	*/
	Ngsshieldawssmtokenfailurerate float64 `json:"ngsshieldawssmtokenfailurerate,omitempty"`
	Ngsrendezvousv1launchsuccess int `json:"ngsrendezvousv1launchsuccess,omitempty"`
	/**
	* Total rendezvous v1 launch success
	*/
	Ngsrendezvousv1launchsuccessrate float64 `json:"ngsrendezvousv1launchsuccessrate,omitempty"`
	Ngsrendezvousv1launchfailure int `json:"ngsrendezvousv1launchfailure,omitempty"`
	/**
	* Total rendezvous v1 launch failures
	*/
	Ngsrendezvousv1launchfailurerate float64 `json:"ngsrendezvousv1launchfailurerate,omitempty"`
	Ngsrendezvousv2launchsuccess int `json:"ngsrendezvousv2launchsuccess,omitempty"`
	/**
	* Total rendezvous v2 launch success
	*/
	Ngsrendezvousv2launchsuccessrate float64 `json:"ngsrendezvousv2launchsuccessrate,omitempty"`
	Ngsrendezvousv2launchfailure int `json:"ngsrendezvousv2launchfailure,omitempty"`
	/**
	* Total rendezvous v2 launch failures
	*/
	Ngsrendezvousv2launchfailurerate float64 `json:"ngsrendezvousv2launchfailurerate,omitempty"`
	Ngsrendezvousv1sntcntcmdfailure int `json:"ngsrendezvousv1sntcntcmdfailure,omitempty"`
	/**
	* Total rendezvous v1 sent connect command failure
	*/
	Ngsrendezvousv1sntcntcmdfailurerate float64 `json:"ngsrendezvousv1sntcntcmdfailurerate,omitempty"`
	Ngsrendezvousv2sntcntcmdfailure int `json:"ngsrendezvousv2sntcntcmdfailure,omitempty"`
	/**
	* Total rendezvous v2 sent connect command failure
	*/
	Ngsrendezvousv2sntcntcmdfailurerate float64 `json:"ngsrendezvousv2sntcntcmdfailurerate,omitempty"`
	Ngsrendezvousv2fallbacktov1 int `json:"ngsrendezvousv2fallbacktov1,omitempty"`
	/**
	* Total rendezvous v2 to v1 fallback
	*/
	Ngsrendezvousv2fallbacktov1rate float64 `json:"ngsrendezvousv2fallbacktov1rate,omitempty"`
	Ngsrendezvousv2delayedresponse int `json:"ngsrendezvousv2delayedresponse,omitempty"`
	/**
	* Total rendezvous v2 delayed response
	*/
	Ngsrendezvousv2delayedresponserate float64 `json:"ngsrendezvousv2delayedresponserate,omitempty"`
	Ngsrendezvousv2invalidcapability int `json:"ngsrendezvousv2invalidcapability,omitempty"`
	/**
	* Total invalid rendezvous capability received
	*/
	Ngsrendezvousv2invalidcapabilityrate float64 `json:"ngsrendezvousv2invalidcapabilityrate,omitempty"`
	Ngsle1secconnesttime int `json:"ngsle1secconnesttime,omitempty"`
	/**
	* Total connection with NGS connect less than 1 sec
	*/
	Ngsle1secconnesttimerate float64 `json:"ngsle1secconnesttimerate,omitempty"`
	Ngsle2secconnesttime int `json:"ngsle2secconnesttime,omitempty"`
	/**
	* Total connection with NGS connect less than 2 sec
	*/
	Ngsle2secconnesttimerate float64 `json:"ngsle2secconnesttimerate,omitempty"`
	Ngsle3secconnesttime int `json:"ngsle3secconnesttime,omitempty"`
	/**
	* Total connection with NGS connect less than 3 sec
	*/
	Ngsle3secconnesttimerate float64 `json:"ngsle3secconnesttimerate,omitempty"`
	Ngsle5secconnesttime int `json:"ngsle5secconnesttime,omitempty"`
	/**
	* Total connection with NGS connect less than 5 sec
	*/
	Ngsle5secconnesttimerate float64 `json:"ngsle5secconnesttimerate,omitempty"`
	Ngsle8secconnesttime int `json:"ngsle8secconnesttime,omitempty"`
	/**
	* Total connection with NGS connect less than 8 sec
	*/
	Ngsle8secconnesttimerate float64 `json:"ngsle8secconnesttimerate,omitempty"`
	Ngsgt8secconnesttime int `json:"ngsgt8secconnesttime,omitempty"`
	/**
	* Total connection with NGS connect greater than 8 sec
	*/
	Ngsgt8secconnesttimerate float64 `json:"ngsgt8secconnesttimerate,omitempty"`
	Ngsapplaunchattempted int `json:"ngsapplaunchattempted,omitempty"`
	/**
	* Total app launch attempted
	*/
	Ngsapplaunchattemptedrate float64 `json:"ngsapplaunchattemptedrate,omitempty"`
	Ngsapplaunchsuccess int `json:"ngsapplaunchsuccess,omitempty"`
	/**
	* Total app launch success
	*/
	Ngsapplaunchsuccessrate float64 `json:"ngsapplaunchsuccessrate,omitempty"`
	Ngsapplaunchfailure int `json:"ngsapplaunchfailure,omitempty"`
	/**
	* Total app launch failures
	*/
	Ngsapplaunchfailurerate float64 `json:"ngsapplaunchfailurerate,omitempty"`
	Ngsstavalfailure int `json:"ngsstavalfailure,omitempty"`
	/**
	* Total STA validation failures
	*/
	Ngsstavalfailurerate float64 `json:"ngsstavalfailurerate,omitempty"`
	Ngsstavalsuccess int `json:"ngsstavalsuccess,omitempty"`
	/**
	* Total STA validation success
	*/
	Ngsstavalsuccessrate float64 `json:"ngsstavalsuccessrate,omitempty"`
	Ngsticketsvcdown int `json:"ngsticketsvcdown,omitempty"`
	/**
	* Total Ticket service down
	*/
	Ngsticketsvcdownrate float64 `json:"ngsticketsvcdownrate,omitempty"`
	Ngsdelayedconnectorresponse int `json:"ngsdelayedconnectorresponse,omitempty"`
	/**
	* Total delayed response received from connector
	*/
	Ngsdelayedconnectorresponserate float64 `json:"ngsdelayedconnectorresponserate,omitempty"`
	Ngstunnelclosedbyclient int `json:"ngstunnelclosedbyclient,omitempty"`
	/**
	* Total number of tunnels closed by client
	*/
	Ngstunnelclosedbyclientrate float64 `json:"ngstunnelclosedbyclientrate,omitempty"`
	Ngstunnelclosedbyvda int `json:"ngstunnelclosedbyvda,omitempty"`
	/**
	* Total number of tunnels closed by VDA
	*/
	Ngstunnelclosedbyvdarate float64 `json:"ngstunnelclosedbyvdarate,omitempty"`
	Ngsconncmdfailbyfr int `json:"ngsconncmdfailbyfr,omitempty"`
	/**
	* Total number fo coonect command failed by FR
	*/
	Ngsconncmdfailbyfrrate float64 `json:"ngsconncmdfailbyfrrate,omitempty"`
	Ngsconncmdfailclientconnnotfound int `json:"ngsconncmdfailclientconnnotfound,omitempty"`
	/**
	* Totla connect command failed with client connection not found
	*/
	Ngsconncmdfailclientconnnotfoundrate float64 `json:"ngsconncmdfailclientconnnotfoundrate,omitempty"`
	Ngsicastart int `json:"ngsicastart,omitempty"`
	/**
	* Total number of ICA start
	*/
	Ngsicastartrate float64 `json:"ngsicastartrate,omitempty"`
	Ngsicaend int `json:"ngsicaend,omitempty"`
	/**
	* Total number of ICA end
	*/
	Ngsicaendrate float64 `json:"ngsicaendrate,omitempty"`
	Ngssocksapplaunchattempted int `json:"ngssocksapplaunchattempted,omitempty"`
	/**
	* Total number of SOCKS app launch attempted
	*/
	Ngssocksapplaunchattemptedrate float64 `json:"ngssocksapplaunchattemptedrate,omitempty"`
	Ngssocksapplaunchsuccess int `json:"ngssocksapplaunchsuccess,omitempty"`
	/**
	* Total number of SOCKS app launch success
	*/
	Ngssocksapplaunchsuccessrate float64 `json:"ngssocksapplaunchsuccessrate,omitempty"`
	Ngscgpapplaunchattempted int `json:"ngscgpapplaunchattempted,omitempty"`
	/**
	* Total number of CGP app launch attempted
	*/
	Ngscgpapplaunchattemptedrate float64 `json:"ngscgpapplaunchattemptedrate,omitempty"`
	Ngscgpapplaunchsuccess int `json:"ngscgpapplaunchsuccess,omitempty"`
	/**
	* Total number of CGP app launch success
	*/
	Ngscgpapplaunchsuccessrate float64 `json:"ngscgpapplaunchsuccessrate,omitempty"`
	Ngsgctapplaunchattempted int `json:"ngsgctapplaunchattempted,omitempty"`
	/**
	* Total number of GCT app launch attempted
	*/
	Ngsgctapplaunchattemptedrate float64 `json:"ngsgctapplaunchattemptedrate,omitempty"`
	Ngsgctapplaunchsuccess int `json:"ngsgctapplaunchsuccess,omitempty"`
	/**
	* Total number of GCT app launch success
	*/
	Ngsgctapplaunchsuccessrate float64 `json:"ngsgctapplaunchsuccessrate,omitempty"`
	Ngsstaapplaunchattempted int `json:"ngsstaapplaunchattempted,omitempty"`
	/**
	* Total number of STA app launch attempted
	*/
	Ngsstaapplaunchattemptedrate float64 `json:"ngsstaapplaunchattemptedrate,omitempty"`
	Ngsstaapplaunchsuccess int `json:"ngsstaapplaunchsuccess,omitempty"`
	/**
	* Total number of STA app launch success
	*/
	Ngsstaapplaunchsuccessrate float64 `json:"ngsstaapplaunchsuccessrate,omitempty"`
	Ngsedtle1secconnesttime int `json:"ngsedtle1secconnesttime,omitempty"`
	/**
	* Total EDT connection with connect time less than 1 sec
	*/
	Ngsedtle1secconnesttimerate float64 `json:"ngsedtle1secconnesttimerate,omitempty"`
	Ngsedtle2secconnesttime int `json:"ngsedtle2secconnesttime,omitempty"`
	/**
	* Total EDT connection with connect time less than 2 sec
	*/
	Ngsedtle2secconnesttimerate float64 `json:"ngsedtle2secconnesttimerate,omitempty"`
	Ngsedtle3secconnesttime int `json:"ngsedtle3secconnesttime,omitempty"`
	/**
	* Total EDT connection with connect time less than 3 sec
	*/
	Ngsedtle3secconnesttimerate float64 `json:"ngsedtle3secconnesttimerate,omitempty"`
	Ngsedtle5secconnesttime int `json:"ngsedtle5secconnesttime,omitempty"`
	/**
	* Total EDT connection with connect time less than 5 sec
	*/
	Ngsedtle5secconnesttimerate float64 `json:"ngsedtle5secconnesttimerate,omitempty"`
	Ngsedtle8secconnesttime int `json:"ngsedtle8secconnesttime,omitempty"`
	/**
	* Total EDT connection with connect time less than 8 sec
	*/
	Ngsedtle8secconnesttimerate float64 `json:"ngsedtle8secconnesttimerate,omitempty"`
	Ngsedtgt8secconnesttime int `json:"ngsedtgt8secconnesttime,omitempty"`
	/**
	* Total EDT connection with connect time greater than 8 sec
	*/
	Ngsedtgt8secconnesttimerate float64 `json:"ngsedtgt8secconnesttimerate,omitempty"`
	Ngsedtle1secclienthandshakeesttime int `json:"ngsedtle1secclienthandshakeesttime,omitempty"`
	/**
	* Total EDT handshake establishment time with client less than 1sec
	*/
	Ngsedtle1secclienthandshakeesttimerate float64 `json:"ngsedtle1secclienthandshakeesttimerate,omitempty"`
	Ngsedtle2secclienthandshakeesttime int `json:"ngsedtle2secclienthandshakeesttime,omitempty"`
	/**
	* Total EDT handshake establishment time with client less than 2sec
	*/
	Ngsedtle2secclienthandshakeesttimerate float64 `json:"ngsedtle2secclienthandshakeesttimerate,omitempty"`
	Ngsedtle3secclienthandshakeesttime int `json:"ngsedtle3secclienthandshakeesttime,omitempty"`
	/**
	* Total EDT handshake establishment time with client less than 3sec
	*/
	Ngsedtle3secclienthandshakeesttimerate float64 `json:"ngsedtle3secclienthandshakeesttimerate,omitempty"`
	Ngsedtgt3secclienthandshakeesttime int `json:"ngsedtgt3secclienthandshakeesttime,omitempty"`
	/**
	* Total EDT handshake establishment time with client greater than 3sec
	*/
	Ngsedtgt3secclienthandshakeesttimerate float64 `json:"ngsedtgt3secclienthandshakeesttimerate,omitempty"`
	Ngsedtle1secvdahandshakeesttime int `json:"ngsedtle1secvdahandshakeesttime,omitempty"`
	/**
	* Total EDT handshake establishment time with vda less than 1sec
	*/
	Ngsedtle1secvdahandshakeesttimerate float64 `json:"ngsedtle1secvdahandshakeesttimerate,omitempty"`
	Ngsedtle2secvdahandshakeesttime int `json:"ngsedtle2secvdahandshakeesttime,omitempty"`
	/**
	* Total EDT handshake establishment time with vda less than 2sec
	*/
	Ngsedtle2secvdahandshakeesttimerate float64 `json:"ngsedtle2secvdahandshakeesttimerate,omitempty"`
	Ngsedtle3secvdahandshakeesttime int `json:"ngsedtle3secvdahandshakeesttime,omitempty"`
	/**
	* Total EDT handshake establishment time with vda less than 3sec
	*/
	Ngsedtle3secvdahandshakeesttimerate float64 `json:"ngsedtle3secvdahandshakeesttimerate,omitempty"`
	Ngsedtgt3secvdahandshakeesttime int `json:"ngsedtgt3secvdahandshakeesttime,omitempty"`
	/**
	* Total EDT handshake establishment time with vda greater than 3sec
	*/
	Ngsedtgt3secvdahandshakeesttimerate float64 `json:"ngsedtgt3secvdahandshakeesttimerate,omitempty"`
	Ngsedtconnrequest int `json:"ngsedtconnrequest,omitempty"`
	/**
	* Total EDT connection request
	*/
	Ngsedtconnrequestrate float64 `json:"ngsedtconnrequestrate,omitempty"`
	Ngsedtclientconnest int `json:"ngsedtclientconnest,omitempty"`
	/**
	* Total client EDT connection establishment
	*/
	Ngsedtclientconnestrate float64 `json:"ngsedtclientconnestrate,omitempty"`
	Ngsedtnonrendzconnreqdrop int `json:"ngsedtnonrendzconnreqdrop,omitempty"`
	/**
	* Total connect request drop due to rendezvous not enabled
	*/
	Ngsedtnonrendzconnreqdroprate float64 `json:"ngsedtnonrendzconnreqdroprate,omitempty"`
	Ngsedtctrlsendconnfailure int `json:"ngsedtctrlsendconnfailure,omitempty"`
	/**
	* Total send connect to controller failed
	*/
	Ngsedtctrlsendconnfailurerate float64 `json:"ngsedtctrlsendconnfailurerate,omitempty"`
	Ngsedtapplaunchsuccess int `json:"ngsedtapplaunchsuccess,omitempty"`
	/**
	* Total EDT app launch success
	*/
	Ngsedtapplaunchsuccessrate float64 `json:"ngsedtapplaunchsuccessrate,omitempty"`
	Ngsedtapplaunchfailure int `json:"ngsedtapplaunchfailure,omitempty"`
	/**
	* Total EDT app launch failures
	*/
	Ngsedtapplaunchfailurerate float64 `json:"ngsedtapplaunchfailurerate,omitempty"`
	Ngsedtstavalfailure int `json:"ngsedtstavalfailure,omitempty"`
	/**
	* Total STA validation failures
	*/
	Ngsedtstavalfailurerate float64 `json:"ngsedtstavalfailurerate,omitempty"`
	Ngsedtstavalsuccess int `json:"ngsedtstavalsuccess,omitempty"`
	/**
	* Total STA validation success
	*/
	Ngsedtstavalsuccessrate float64 `json:"ngsedtstavalsuccessrate,omitempty"`
	Ngsedtfrvdaconn int `json:"ngsedtfrvdaconn,omitempty"`
	/**
	* Total FR VDA connection
	*/
	Ngsedtfrvdaconnrate float64 `json:"ngsedtfrvdaconnrate,omitempty"`
	Ngsedtfrsnconn int `json:"ngsedtfrsnconn,omitempty"`
	/**
	* Total FR SN connection
	*/
	Ngsedtfrsnconnrate float64 `json:"ngsedtfrsnconnrate,omitempty"`
	Ngsedtvdaconnest int `json:"ngsedtvdaconnest,omitempty"`
	/**
	* Total VDA EDT connection establishment
	*/
	Ngsedtvdaconnestrate float64 `json:"ngsedtvdaconnestrate,omitempty"`
	Ngsedtvdarendezvoustimeout int `json:"ngsedtvdarendezvoustimeout,omitempty"`
	/**
	* Total VDA rendezvous timeout
	*/
	Ngsedtvdarendezvoustimeoutrate float64 `json:"ngsedtvdarendezvoustimeoutrate,omitempty"`
	Ngsedtnonlossytunnel int `json:"ngsedtnonlossytunnel,omitempty"`
	/**
	* Total EDT non lossy tunnels
	*/
	Ngsedtnonlossytunnelrate float64 `json:"ngsedtnonlossytunnelrate,omitempty"`
	Ngsedtlossytunnel int `json:"ngsedtlossytunnel,omitempty"`
	/**
	* Total EDT lossy tunnels
	*/
	Ngsedtlossytunnelrate float64 `json:"ngsedtlossytunnelrate,omitempty"`
	Ngsedtgetrequestsenttoredis int `json:"ngsedtgetrequestsenttoredis,omitempty"`
	/**
	* Total Redis GET Request
	*/
	Ngsedtgetrequestsenttoredisrate float64 `json:"ngsedtgetrequestsenttoredisrate,omitempty"`
	Ngsedtredislookupsuccess int `json:"ngsedtredislookupsuccess,omitempty"`
	/**
	* Total Redis lookup success
	*/
	Ngsedtredislookupsuccessrate float64 `json:"ngsedtredislookupsuccessrate,omitempty"`
	Ngsedtredislookupfailure int `json:"ngsedtredislookupfailure,omitempty"`
	/**
	* Total Redis lookup failure
	*/
	Ngsedtredislookupfailurerate float64 `json:"ngsedtredislookupfailurerate,omitempty"`
	Ngsedtredisentryaddsucces int `json:"ngsedtredisentryaddsucces,omitempty"`
	/**
	* Entries Added To Redis Successfully
	*/
	Ngsedtredisentryaddsuccesrate float64 `json:"ngsedtredisentryaddsuccesrate,omitempty"`
	Ngsedtredisentryaddfailure int `json:"ngsedtredisentryaddfailure,omitempty"`
	/**
	* Entries Addition to Redis failure
	*/
	Ngsedtredisentryaddfailurerate float64 `json:"ngsedtredisentryaddfailurerate,omitempty"`
	Ngsedtredisentrydeletesuccess int `json:"ngsedtredisentrydeletesuccess,omitempty"`
	/**
	* Entry Deleted Successfully In Redis
	*/
	Ngsedtredisentrydeletesuccessrate float64 `json:"ngsedtredisentrydeletesuccessrate,omitempty"`
	Ngsedtredisentrydeletefailure int `json:"ngsedtredisentrydeletefailure,omitempty"`
	/**
	* Entry Deletion Failed in Redis
	*/
	Ngsedtredisentrydeletefailurerate float64 `json:"ngsedtredisentrydeletefailurerate,omitempty"`
	Ngsedttotsessionsreroutingdatapkt int `json:"ngsedttotsessionsreroutingdatapkt,omitempty"`
	/**
	* Total EDT Sessions Re-Routing Data Pkt
	*/
	Ngsedtsessionsreroutingdatapktrate float64 `json:"ngsedtsessionsreroutingdatapktrate,omitempty"`
	Ngsedttotsessionsreceivingredirectedpkt int `json:"ngsedttotsessionsreceivingredirectedpkt,omitempty"`
	/**
	* Total EDT Sessions Receiving Redirected Pkt
	*/
	Ngsedtsessionsreceivingredirectedpktrate float64 `json:"ngsedtsessionsreceivingredirectedpktrate,omitempty"`
	Ngsedtcurrentsessionsreroutingdatapkt int `json:"ngsedtcurrentsessionsreroutingdatapkt,omitempty"`
	/**
	* Current EDT Sessions Re-Routing Data Pkt
	*/
	Ngsedtcurrentsessionsreroutingdatapktrate float64 `json:"ngsedtcurrentsessionsreroutingdatapktrate,omitempty"`
	Ngsedtcurrentsessionsreceivingredirectedpkt int `json:"ngsedtcurrentsessionsreceivingredirectedpkt,omitempty"`
	/**
	* Current EDT Sessions Receiving Redirected Pkt
	*/
	Ngsedtcurrentsessionsreceivingredirectedpktrate float64 `json:"ngsedtcurrentsessionsreceivingredirectedpktrate,omitempty"`
	Ngsagentaccessle1secconnesttime int `json:"ngsagentaccessle1secconnesttime,omitempty"`
	/**
	* Total Agent Access based Connection with connect time less than 1 sec
	*/
	Ngsagentaccessle1secconnesttimerate float64 `json:"ngsagentaccessle1secconnesttimerate,omitempty"`
	Ngsagentaccessle2secconnesttime int `json:"ngsagentaccessle2secconnesttime,omitempty"`
	/**
	* Total Agent Access based Connection with connect time less than 2 sec
	*/
	Ngsagentaccessle2secconnesttimerate float64 `json:"ngsagentaccessle2secconnesttimerate,omitempty"`
	Ngsagentaccessle3secconnesttime int `json:"ngsagentaccessle3secconnesttime,omitempty"`
	/**
	* Total Agent Access based Connection with connect time less than 3 sec
	*/
	Ngsagentaccessle3secconnesttimerate float64 `json:"ngsagentaccessle3secconnesttimerate,omitempty"`
	Ngsagentaccessle5secconnesttime int `json:"ngsagentaccessle5secconnesttime,omitempty"`
	/**
	* Total Agent Access based Connection with connect time less than 5 sec
	*/
	Ngsagentaccessle5secconnesttimerate float64 `json:"ngsagentaccessle5secconnesttimerate,omitempty"`
	Ngsagentaccessle8secconnesttime int `json:"ngsagentaccessle8secconnesttime,omitempty"`
	/**
	* Total Agent Access based Connection with connect time less than 8 sec
	*/
	Ngsagentaccessle8secconnesttimerate float64 `json:"ngsagentaccessle8secconnesttimerate,omitempty"`
	Ngsagentaccessle13secconnesttime int `json:"ngsagentaccessle13secconnesttime,omitempty"`
	/**
	* Total Agent Access based Connection with connect time less than 13 sec
	*/
	Ngsagentaccessle13secconnesttimerate float64 `json:"ngsagentaccessle13secconnesttimerate,omitempty"`
	Ngsagentaccessle21secconnesttime int `json:"ngsagentaccessle21secconnesttime,omitempty"`
	/**
	* Total Agent Access based Connection with connect time less than 21 sec
	*/
	Ngsagentaccessle21secconnesttimerate float64 `json:"ngsagentaccessle21secconnesttimerate,omitempty"`
	Ngsagentaccessgt21secconnesttime int `json:"ngsagentaccessgt21secconnesttime,omitempty"`
	/**
	* Total Agent Access based Connection with connect time  greater than 21 sec
	*/
	Ngsagentaccessgt21secconnesttimerate float64 `json:"ngsagentaccessgt21secconnesttimerate,omitempty"`

}
