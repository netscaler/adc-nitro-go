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

package qos


type Qosstats struct {
	Name string `json:"name,omitempty"`
	/**
	* Clear the statsistics / counters
	*/
	Clearstats string `json:"clearstats,omitempty"`
	Snmpqosqospacketsreceived int `json:"snmpqosqos_packets_received,omitempty"`
	/**
	* Receive direction packets processed by QoS
	*/
	Snmpqosqospacketsreceivedrate float64 `json:"snmpqosqos_packets_receivedrate,omitempty"`
	Snmpqosqospacketssent int `json:"snmpqosqos_packets_sent,omitempty"`
	/**
	* Send direction packets processed by QoS
	*/
	Snmpqosqospacketssentrate float64 `json:"snmpqosqos_packets_sentrate,omitempty"`
	Snmpqosqospacketsbypassed int `json:"snmpqosqos_packets_bypassed,omitempty"`
	/**
	* Packets bypassing QoS
	*/
	Snmpqosqospacketsbypassedrate float64 `json:"snmpqosqos_packets_bypassedrate,omitempty"`
	Snmpqosqospacketsdropped int `json:"snmpqosqos_packets_dropped,omitempty"`
	/**
	* Total packets dropped
	*/
	Snmpqosqospacketsdroppedrate float64 `json:"snmpqosqos_packets_droppedrate,omitempty"`
	Snmpqosqosbytesrx int `json:"snmpqosqos_bytes_rx,omitempty"`
	/**
	* Received bytes processed by QoS
	*/
	Snmpqosqosbytesrxrate float64 `json:"snmpqosqos_bytes_rxrate,omitempty"`
	Snmpqosqosbytestx int `json:"snmpqosqos_bytes_tx,omitempty"`
	/**
	* Sent bytes processed by QoS
	*/
	Snmpqosqosbytestxrate float64 `json:"snmpqosqos_bytes_txrate,omitempty"`
	Snmpqosqoslazybytes int `json:"snmpqosqos_lazy_bytes,omitempty"`
	/**
	* QoS lazy byte optimization rate
	*/
	Snmpqosqoslazybytesrate float64 `json:"snmpqosqos_lazy_bytesrate,omitempty"`
	Snmpqosqosrealbytes int `json:"snmpqosqos_real_bytes,omitempty"`
	/**
	* QoS actual bytes scheduled
	*/
	Snmpqosqosrealbytesrate float64 `json:"snmpqosqos_real_bytesrate,omitempty"`
	Snmpqosqospacketsfiltered int `json:"snmpqosqos_packets_filtered,omitempty"`
	/**
	* Total packets filtered by QoS
	*/
	Snmpqosqospacketsfilteredrate float64 `json:"snmpqosqos_packets_filteredrate,omitempty"`
	Snmpqosqospacketsclassified int `json:"snmpqosqos_packets_classified,omitempty"`
	/**
	* Total packets classified by QoS
	*/
	Snmpqosqospacketsclassifiedrate float64 `json:"snmpqosqos_packets_classifiedrate,omitempty"`
	Snmpqosqosflows int `json:"snmpqosqos_flows,omitempty"`
	/**
	* New QoS flows
	*/
	Snmpqosqosflowsrate float64 `json:"snmpqosqos_flowsrate,omitempty"`
	Snmpqosqosflowrecycles int `json:"snmpqosqos_flow_recycles,omitempty"`
	/**
	* Recycled QoS flows
	*/
	Snmpqosqosflowrecyclesrate float64 `json:"snmpqosqos_flow_recyclesrate,omitempty"`
	Snmpqosqossessionrecyclefailure int `json:"snmpqosqos_session_recycle_failure,omitempty"`
	/**
	* QoS Flow Recycle failures
	*/
	Snmpqosqossessionrecyclefailurerate float64 `json:"snmpqosqos_session_recycle_failurerate,omitempty"`
	Snmpqosqossessionsignored int `json:"snmpqosqos_sessions_ignored,omitempty"`
	/**
	* Sessions manually ignored
	*/
	Snmpqosqossessionsignoredrate float64 `json:"snmpqosqos_sessions_ignoredrate,omitempty"`
	Snmpqosqossessionsconsumed int `json:"snmpqosqos_sessions_consumed,omitempty"`
	/**
	* sessions manually consumed
	*/
	Snmpqosqossessionsconsumedrate float64 `json:"snmpqosqos_sessions_consumedrate,omitempty"`
	Snmpqosqosactionscreated int `json:"snmpqosqos_actions_created,omitempty"`
	/**
	* Uneque qos action objects created
	*/
	Snmpqosqosactionscreatedrate float64 `json:"snmpqosqos_actions_createdrate,omitempty"`
	Snmpqosqospolicyreeval int `json:"snmpqosqos_policy_reeval,omitempty"`
	/**
	* Policies re-evaluated due to cli change
	*/
	Snmpqosqospolicyreevalrate float64 `json:"snmpqosqos_policy_reevalrate,omitempty"`
	Snmpqosqoscfytcpunknown int `json:"snmpqosqos_cfy_tcp_unknown,omitempty"`
	/**
	* Connections unable to be classified beyond TCP
	*/
	Snmpqosqoscfytcpunknownrate float64 `json:"snmpqosqos_cfy_tcp_unknownrate,omitempty"`
	Snmpqosqoscfyudpunknown int `json:"snmpqosqos_cfy_udp_unknown,omitempty"`
	/**
	* Connections unable to be classified beyond UDP
	*/
	Snmpqosqoscfyudpunknownrate float64 `json:"snmpqosqos_cfy_udp_unknownrate,omitempty"`
	Snmpqosqosschleafs int `json:"snmpqosqos_sch_leafs,omitempty"`
	/**
	* Scheduler leaf nodes constructed
	*/
	Snmpqosqosschleafsrate float64 `json:"snmpqosqos_sch_leafsrate,omitempty"`
	Snmpqosqossessionmem int `json:"snmpqosqos_session_mem,omitempty"`
	/**
	* Session memory allocated
	*/
	Snmpqosqossessionmemrate float64 `json:"snmpqosqos_session_memrate,omitempty"`
	Snmpqosqosschvirtualpackets int `json:"snmpqosqos_sch_virtual_packets,omitempty"`
	/**
	* Scheduler virtual packets constructed
	*/
	Snmpqosqosschvirtualpacketsrate float64 `json:"snmpqosqos_sch_virtual_packetsrate,omitempty"`
	Snmpqosqosschvirtualbytesaccepted int `json:"snmpqosqos_sch_virtual_bytes_accepted,omitempty"`
	/**
	* Scheduler bytes accepted
	*/
	Snmpqosqosschvirtualbytesacceptedrate float64 `json:"snmpqosqos_sch_virtual_bytes_acceptedrate,omitempty"`
	Snmpqosqosschleafrecyclefailures int `json:"snmpqosqos_sch_leaf_recycle_failures,omitempty"`
	/**
	* Scheduler Failures to recycle QoS flows
	*/
	Snmpqosqosschleafrecyclefailuresrate float64 `json:"snmpqosqos_sch_leaf_recycle_failuresrate,omitempty"`
	Snmpqosqosschnoderegulatedcount int `json:"snmpqosqos_sch_node_regulated_count,omitempty"`
	/**
	* Scheduler Regulated node count
	*/
	Snmpqosqosschnoderegulatedcountrate float64 `json:"snmpqosqos_sch_node_regulated_countrate,omitempty"`
	Snmpqosqosschsessionscreated int `json:"snmpqosqos_sch_sessions_created,omitempty"`
	/**
	* Scheduler session classes constructed
	*/
	Snmpqosqosschsessionscreatedrate float64 `json:"snmpqosqos_sch_sessions_createdrate,omitempty"`
	Snmpqosqosschsessionsdeleted int `json:"snmpqosqos_sch_sessions_deleted,omitempty"`
	/**
	* Scheduler session classes constructed
	*/
	Snmpqosqosschsessionsdeletedrate float64 `json:"snmpqosqos_sch_sessions_deletedrate,omitempty"`
	Snmpqosqosschsdrrnodes int `json:"snmpqosqos_sch_sdrr_nodes,omitempty"`
	/**
	* Scheduler sdrr nodes constructed
	*/
	Snmpqosqosschsdrrnodesrate float64 `json:"snmpqosqos_sch_sdrr_nodesrate,omitempty"`
	Snmpqosqosschsessionconns int `json:"snmpqosqos_sch_session_conns,omitempty"`
	/**
	* Scheduler session connections created
	*/
	Snmpqosqosschsessionconnsrate float64 `json:"snmpqosqos_sch_session_connsrate,omitempty"`
	Snmpqosqosschsessionconnsremoved int `json:"snmpqosqos_sch_session_conns_removed,omitempty"`
	/**
	* Scheduler session connections removed
	*/
	Snmpqosqosschsessionconnsremovedrate float64 `json:"snmpqosqos_sch_session_conns_removedrate,omitempty"`
	Snmpqosqosschsessionsregulatedcount int `json:"snmpqosqos_sch_sessions_regulated_count,omitempty"`
	/**
	* Scheduler regulated sessions count
	*/
	Snmpqosqosschsessionsregulatedcountrate float64 `json:"snmpqosqos_sch_sessions_regulated_countrate,omitempty"`
	Snmpqosqosschsessionsbytecount int `json:"snmpqosqos_sch_sessions_byte_count,omitempty"`
	/**
	* Scheduler session bytes total
	*/
	Snmpqosqosschsessionsbytecountrate float64 `json:"snmpqosqos_sch_sessions_byte_countrate,omitempty"`
	Snmpqosqosschregulatedcount int `json:"snmpqosqos_sch_regulated_count,omitempty"`
	/**
	* Scheduler regulated node count
	*/
	Snmpqosqosschregulatedcountrate float64 `json:"snmpqosqos_sch_regulated_countrate,omitempty"`
	Snmpqosqosschlinkscreated int `json:"snmpqosqos_sch_links_created,omitempty"`
	/**
	* Scheduler links created
	*/
	Snmpqosqosschlinkscreatedrate float64 `json:"snmpqosqos_sch_links_createdrate,omitempty"`
	Snmpqosqosschlinksdeleted int `json:"snmpqosqos_sch_links_deleted,omitempty"`
	/**
	* Scheduler links deleted
	*/
	Snmpqosqosschlinksdeletedrate float64 `json:"snmpqosqos_sch_links_deletedrate,omitempty"`
	Snmpqosqosschlinksupdated int `json:"snmpqosqos_sch_links_updated,omitempty"`
	/**
	* Scheduler links updated
	*/
	Snmpqosqosschlinksupdatedrate float64 `json:"snmpqosqos_sch_links_updatedrate,omitempty"`
	Snmpqosqosschpollcount int `json:"snmpqosqos_sch_poll_count,omitempty"`
	/**
	* Scheduler calls to poll_libqos
	*/
	Snmpqosqosschpollcountrate float64 `json:"snmpqosqos_sch_poll_countrate,omitempty"`
	Snmpqosqosschpeermsgs int `json:"snmpqosqos_sch_peer_msgs,omitempty"`
	/**
	* Scheduler peer messages received
	*/
	Snmpqosqosschpeermsgsrate float64 `json:"snmpqosqos_sch_peer_msgsrate,omitempty"`
	Snmpqosqoserroripc int `json:"snmpqosqos_error_ipc,omitempty"`
	/**
	* IPC failed for QoS messages.
	*/
	Snmpqosqoserroripcrate float64 `json:"snmpqosqos_error_ipcrate,omitempty"`
	Snmpqosqosflowmem int `json:"snmpqosqos_flow_mem,omitempty"`
	/**
	* Flow memory allocated
	*/
	Snmpqosqosflowmemrate float64 `json:"snmpqosqos_flow_memrate,omitempty"`
	Snmpqosqosflowsavailable int `json:"snmpqosqos_flows_available,omitempty"`
	/**
	* Flows free list size
	*/
	Snmpqosqosflowsavailablerate float64 `json:"snmpqosqos_flows_availablerate,omitempty"`
	Snmpqosqosrecyclefailedbacklog int `json:"snmpqosqos_recycle_failed_backlog,omitempty"`
	/**
	* Recycle failed due to backlog
	*/
	Snmpqosqosrecyclefailedbacklograte float64 `json:"snmpqosqos_recycle_failed_backlograte,omitempty"`
	Snmpqosqosrecyclefailedsession int `json:"snmpqosqos_recycle_failed_session,omitempty"`
	/**
	* Recycle failed due to session attachment
	*/
	Snmpqosqosrecyclefailedsessionrate float64 `json:"snmpqosqos_recycle_failed_sessionrate,omitempty"`
	Snmpqosqoserrorcreateactionfailed int `json:"snmpqosqos_error_create_action_failed,omitempty"`
	/**
	* Failed attempts to create actions
	*/
	Snmpqosqoserrorcreateactionfailedrate float64 `json:"snmpqosqos_error_create_action_failedrate,omitempty"`
	Snmpqosqoserrormodifyactionfailed int `json:"snmpqosqos_error_modify_action_failed,omitempty"`
	/**
	* Failed attempts to modify actions
	*/
	Snmpqosqoserrormodifyactionfailedrate float64 `json:"snmpqosqos_error_modify_action_failedrate,omitempty"`
	Snmpqosqoserrorremoveactionfailed int `json:"snmpqosqos_error_remove_action_failed,omitempty"`
	/**
	* Failed attempts to remove actions
	*/
	Snmpqosqoserrorremoveactionfailedrate float64 `json:"snmpqosqos_error_remove_action_failedrate,omitempty"`
	Snmpqosqoserrorcliunknown int `json:"snmpqosqos_error_cli_unknown,omitempty"`
	/**
	* Internal CLI error
	*/
	Snmpqosqoserrorcliunknownrate float64 `json:"snmpqosqos_error_cli_unknownrate,omitempty"`
	Snmpqosqoserrorrenamenotimplemented int `json:"snmpqosqos_error_rename_not_implemented,omitempty"`
	/**
	* qos action rename not yet implemented
	*/
	Snmpqosqoserrorrenamenotimplementedrate float64 `json:"snmpqosqos_error_rename_not_implementedrate,omitempty"`
	Snmpqosqoserrorremovepolicyfailed int `json:"snmpqosqos_error_remove_policy_failed,omitempty"`
	/**
	* Failed attempts to remove qos policy
	*/
	Snmpqosqoserrorremovepolicyfailedrate float64 `json:"snmpqosqos_error_remove_policy_failedrate,omitempty"`
	Snmpqosqoserrorcreatepolicyfailed int `json:"snmpqosqos_error_create_policy_failed,omitempty"`
	/**
	* Failed attempts to create qos policy
	*/
	Snmpqosqoserrorcreatepolicyfailedrate float64 `json:"snmpqosqos_error_create_policy_failedrate,omitempty"`
	Snmpqosqoserrorlibqosapifailures int `json:"snmpqosqos_error_libqos_api_failures,omitempty"`
	/**
	* Libqos api failures
	*/
	Snmpqosqoserrorlibqosapifailuresrate float64 `json:"snmpqosqos_error_libqos_api_failuresrate,omitempty"`
	Snmpqosqoserrorapisesinvalidpcb int `json:"snmpqosqos_error_api_ses_invalidpcb,omitempty"`
	/**
	* Libqos api qos_session_add_pcb/natpcb() failed for reason QS_EINVALIDPCB
	*/
	Snmpqosqoserrorapisesinvalidpcbrate float64 `json:"snmpqosqos_error_api_ses_invalidpcbrate,omitempty"`
	Snmpqosqoserrorapisesnotready int `json:"snmpqosqos_error_api_ses_notready,omitempty"`
	/**
	* Libqos api qos_session_add_pcb/natpcb() failed for reason QS_ENOTREADY
	*/
	Snmpqosqoserrorapisesnotreadyrate float64 `json:"snmpqosqos_error_api_ses_notreadyrate,omitempty"`
	Snmpqosqoserrorapisesaddinsession int `json:"snmpqosqos_error_api_ses_add_insession,omitempty"`
	/**
	* Libqos api qos_session_add_pcb/natpcb() failed for reason QS_EINSESSION
	*/
	Snmpqosqoserrorapisesaddinsessionrate float64 `json:"snmpqosqos_error_api_ses_add_insessionrate,omitempty"`
	Snmpqosqoserrorapisesaddother int `json:"snmpqosqos_error_api_ses_add_other,omitempty"`
	/**
	* Libqos api qos_session_add_pcb/natpcb() failed
	*/
	Snmpqosqoserrorapisesaddotherrate float64 `json:"snmpqosqos_error_api_ses_add_otherrate,omitempty"`
	Snmpqosqoserrorapisesremnotinsession int `json:"snmpqosqos_error_api_ses_rem_notinsession,omitempty"`
	/**
	* Libqos api qos_session_rem_pcb/natpcb() failed for reason QS_ENOTINSESSION
	*/
	Snmpqosqoserrorapisesremnotinsessionrate float64 `json:"snmpqosqos_error_api_ses_rem_notinsessionrate,omitempty"`
	Snmpqosqoserrorapisesremother int `json:"snmpqosqos_error_api_ses_rem_other,omitempty"`
	/**
	* Libqos api qos_session_rem_pcb/natpcb() failed
	*/
	Snmpqosqoserrorapisesremotherrate float64 `json:"snmpqosqos_error_api_ses_rem_otherrate,omitempty"`
	Snmpqosqoserrorapisesdel int `json:"snmpqosqos_error_api_ses_del,omitempty"`
	/**
	* Libqos api qos_session_delete faled
	*/
	Snmpqosqoserrorapisesdelrate float64 `json:"snmpqosqos_error_api_ses_delrate,omitempty"`
	Snmpqosqoserrornoflows int `json:"snmpqosqos_error_no_flows,omitempty"`
	/**
	* Libqos out of flow memory
	*/
	Snmpqosqoserrornoflowsrate float64 `json:"snmpqosqos_error_no_flowsrate,omitempty"`
	Snmpqosqoserrornosessions int `json:"snmpqosqos_error_no_sessions,omitempty"`
	/**
	* Libqos out of session memory
	*/
	Snmpqosqoserrornosessionsrate float64 `json:"snmpqosqos_error_no_sessionsrate,omitempty"`

}
