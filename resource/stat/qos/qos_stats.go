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
	Snmpqosqospacketsreceived uint64 `json:"snmpqosqos_packets_received,omitempty"`
	/**
	* Receive direction packets processed by QoS
	*/
	Snmpqosqospacketsreceivedrate int32 `json:"snmpqosqos_packets_receivedrate,omitempty"`
	Snmpqosqospacketssent uint64 `json:"snmpqosqos_packets_sent,omitempty"`
	/**
	* Send direction packets processed by QoS
	*/
	Snmpqosqospacketssentrate int32 `json:"snmpqosqos_packets_sentrate,omitempty"`
	Snmpqosqospacketsbypassed uint64 `json:"snmpqosqos_packets_bypassed,omitempty"`
	/**
	* Packets bypassing QoS
	*/
	Snmpqosqospacketsbypassedrate int32 `json:"snmpqosqos_packets_bypassedrate,omitempty"`
	Snmpqosqospacketsdropped uint64 `json:"snmpqosqos_packets_dropped,omitempty"`
	/**
	* Total packets dropped
	*/
	Snmpqosqospacketsdroppedrate int32 `json:"snmpqosqos_packets_droppedrate,omitempty"`
	Snmpqosqosbytesrx uint64 `json:"snmpqosqos_bytes_rx,omitempty"`
	/**
	* Received bytes processed by QoS
	*/
	Snmpqosqosbytesrxrate int32 `json:"snmpqosqos_bytes_rxrate,omitempty"`
	Snmpqosqosbytestx uint64 `json:"snmpqosqos_bytes_tx,omitempty"`
	/**
	* Sent bytes processed by QoS
	*/
	Snmpqosqosbytestxrate int32 `json:"snmpqosqos_bytes_txrate,omitempty"`
	Snmpqosqoslazybytes uint64 `json:"snmpqosqos_lazy_bytes,omitempty"`
	/**
	* QoS lazy byte optimization rate
	*/
	Snmpqosqoslazybytesrate int32 `json:"snmpqosqos_lazy_bytesrate,omitempty"`
	Snmpqosqosrealbytes uint64 `json:"snmpqosqos_real_bytes,omitempty"`
	/**
	* QoS actual bytes scheduled
	*/
	Snmpqosqosrealbytesrate int32 `json:"snmpqosqos_real_bytesrate,omitempty"`
	Snmpqosqospacketsfiltered uint64 `json:"snmpqosqos_packets_filtered,omitempty"`
	/**
	* Total packets filtered by QoS
	*/
	Snmpqosqospacketsfilteredrate int32 `json:"snmpqosqos_packets_filteredrate,omitempty"`
	Snmpqosqospacketsclassified uint64 `json:"snmpqosqos_packets_classified,omitempty"`
	/**
	* Total packets classified by QoS
	*/
	Snmpqosqospacketsclassifiedrate int32 `json:"snmpqosqos_packets_classifiedrate,omitempty"`
	Snmpqosqosflows uint64 `json:"snmpqosqos_flows,omitempty"`
	/**
	* New QoS flows
	*/
	Snmpqosqosflowsrate int32 `json:"snmpqosqos_flowsrate,omitempty"`
	Snmpqosqosflowrecycles uint64 `json:"snmpqosqos_flow_recycles,omitempty"`
	/**
	* Recycled QoS flows
	*/
	Snmpqosqosflowrecyclesrate int32 `json:"snmpqosqos_flow_recyclesrate,omitempty"`
	Snmpqosqossessionrecyclefailure uint64 `json:"snmpqosqos_session_recycle_failure,omitempty"`
	/**
	* QoS Flow Recycle failures
	*/
	Snmpqosqossessionrecyclefailurerate int32 `json:"snmpqosqos_session_recycle_failurerate,omitempty"`
	Snmpqosqossessionsignored uint64 `json:"snmpqosqos_sessions_ignored,omitempty"`
	/**
	* Sessions manually ignored
	*/
	Snmpqosqossessionsignoredrate int32 `json:"snmpqosqos_sessions_ignoredrate,omitempty"`
	Snmpqosqossessionsconsumed uint64 `json:"snmpqosqos_sessions_consumed,omitempty"`
	/**
	* sessions manually consumed
	*/
	Snmpqosqossessionsconsumedrate int32 `json:"snmpqosqos_sessions_consumedrate,omitempty"`
	Snmpqosqosactionscreated uint64 `json:"snmpqosqos_actions_created,omitempty"`
	/**
	* Uneque qos action objects created
	*/
	Snmpqosqosactionscreatedrate int32 `json:"snmpqosqos_actions_createdrate,omitempty"`
	Snmpqosqospolicyreeval uint64 `json:"snmpqosqos_policy_reeval,omitempty"`
	/**
	* Policies re-evaluated due to cli change
	*/
	Snmpqosqospolicyreevalrate int32 `json:"snmpqosqos_policy_reevalrate,omitempty"`
	Snmpqosqoscfytcpunknown uint64 `json:"snmpqosqos_cfy_tcp_unknown,omitempty"`
	/**
	* Connections unable to be classified beyond TCP
	*/
	Snmpqosqoscfytcpunknownrate int32 `json:"snmpqosqos_cfy_tcp_unknownrate,omitempty"`
	Snmpqosqoscfyudpunknown uint64 `json:"snmpqosqos_cfy_udp_unknown,omitempty"`
	/**
	* Connections unable to be classified beyond UDP
	*/
	Snmpqosqoscfyudpunknownrate int32 `json:"snmpqosqos_cfy_udp_unknownrate,omitempty"`
	Snmpqosqosschleafs uint64 `json:"snmpqosqos_sch_leafs,omitempty"`
	/**
	* Scheduler leaf nodes constructed
	*/
	Snmpqosqosschleafsrate int32 `json:"snmpqosqos_sch_leafsrate,omitempty"`
	Snmpqosqossessionmem uint64 `json:"snmpqosqos_session_mem,omitempty"`
	/**
	* Session memory allocated
	*/
	Snmpqosqossessionmemrate int32 `json:"snmpqosqos_session_memrate,omitempty"`
	Snmpqosqosschvirtualpackets uint64 `json:"snmpqosqos_sch_virtual_packets,omitempty"`
	/**
	* Scheduler virtual packets constructed
	*/
	Snmpqosqosschvirtualpacketsrate int32 `json:"snmpqosqos_sch_virtual_packetsrate,omitempty"`
	Snmpqosqosschvirtualbytesaccepted uint64 `json:"snmpqosqos_sch_virtual_bytes_accepted,omitempty"`
	/**
	* Scheduler bytes accepted
	*/
	Snmpqosqosschvirtualbytesacceptedrate int32 `json:"snmpqosqos_sch_virtual_bytes_acceptedrate,omitempty"`
	Snmpqosqosschleafrecyclefailures uint64 `json:"snmpqosqos_sch_leaf_recycle_failures,omitempty"`
	/**
	* Scheduler Failures to recycle QoS flows
	*/
	Snmpqosqosschleafrecyclefailuresrate int32 `json:"snmpqosqos_sch_leaf_recycle_failuresrate,omitempty"`
	Snmpqosqosschnoderegulatedcount uint64 `json:"snmpqosqos_sch_node_regulated_count,omitempty"`
	/**
	* Scheduler Regulated node count
	*/
	Snmpqosqosschnoderegulatedcountrate int32 `json:"snmpqosqos_sch_node_regulated_countrate,omitempty"`
	Snmpqosqosschsessionscreated uint64 `json:"snmpqosqos_sch_sessions_created,omitempty"`
	/**
	* Scheduler session classes constructed
	*/
	Snmpqosqosschsessionscreatedrate int32 `json:"snmpqosqos_sch_sessions_createdrate,omitempty"`
	Snmpqosqosschsessionsdeleted uint64 `json:"snmpqosqos_sch_sessions_deleted,omitempty"`
	/**
	* Scheduler session classes constructed
	*/
	Snmpqosqosschsessionsdeletedrate int32 `json:"snmpqosqos_sch_sessions_deletedrate,omitempty"`
	Snmpqosqosschsdrrnodes uint64 `json:"snmpqosqos_sch_sdrr_nodes,omitempty"`
	/**
	* Scheduler sdrr nodes constructed
	*/
	Snmpqosqosschsdrrnodesrate int32 `json:"snmpqosqos_sch_sdrr_nodesrate,omitempty"`
	Snmpqosqosschsessionconns uint64 `json:"snmpqosqos_sch_session_conns,omitempty"`
	/**
	* Scheduler session connections created
	*/
	Snmpqosqosschsessionconnsrate int32 `json:"snmpqosqos_sch_session_connsrate,omitempty"`
	Snmpqosqosschsessionconnsremoved uint64 `json:"snmpqosqos_sch_session_conns_removed,omitempty"`
	/**
	* Scheduler session connections removed
	*/
	Snmpqosqosschsessionconnsremovedrate int32 `json:"snmpqosqos_sch_session_conns_removedrate,omitempty"`
	Snmpqosqosschsessionsregulatedcount uint64 `json:"snmpqosqos_sch_sessions_regulated_count,omitempty"`
	/**
	* Scheduler regulated sessions count
	*/
	Snmpqosqosschsessionsregulatedcountrate int32 `json:"snmpqosqos_sch_sessions_regulated_countrate,omitempty"`
	Snmpqosqosschsessionsbytecount uint64 `json:"snmpqosqos_sch_sessions_byte_count,omitempty"`
	/**
	* Scheduler session bytes total
	*/
	Snmpqosqosschsessionsbytecountrate int32 `json:"snmpqosqos_sch_sessions_byte_countrate,omitempty"`
	Snmpqosqosschregulatedcount uint64 `json:"snmpqosqos_sch_regulated_count,omitempty"`
	/**
	* Scheduler regulated node count
	*/
	Snmpqosqosschregulatedcountrate int32 `json:"snmpqosqos_sch_regulated_countrate,omitempty"`
	Snmpqosqosschlinkscreated uint64 `json:"snmpqosqos_sch_links_created,omitempty"`
	/**
	* Scheduler links created
	*/
	Snmpqosqosschlinkscreatedrate int32 `json:"snmpqosqos_sch_links_createdrate,omitempty"`
	Snmpqosqosschlinksdeleted uint64 `json:"snmpqosqos_sch_links_deleted,omitempty"`
	/**
	* Scheduler links deleted
	*/
	Snmpqosqosschlinksdeletedrate int32 `json:"snmpqosqos_sch_links_deletedrate,omitempty"`
	Snmpqosqosschlinksupdated uint64 `json:"snmpqosqos_sch_links_updated,omitempty"`
	/**
	* Scheduler links updated
	*/
	Snmpqosqosschlinksupdatedrate int32 `json:"snmpqosqos_sch_links_updatedrate,omitempty"`
	Snmpqosqosschpollcount uint64 `json:"snmpqosqos_sch_poll_count,omitempty"`
	/**
	* Scheduler calls to poll_libqos
	*/
	Snmpqosqosschpollcountrate int32 `json:"snmpqosqos_sch_poll_countrate,omitempty"`
	Snmpqosqosschpeermsgs uint64 `json:"snmpqosqos_sch_peer_msgs,omitempty"`
	/**
	* Scheduler peer messages received
	*/
	Snmpqosqosschpeermsgsrate int32 `json:"snmpqosqos_sch_peer_msgsrate,omitempty"`
	Snmpqosqoserroripc uint64 `json:"snmpqosqos_error_ipc,omitempty"`
	/**
	* IPC failed for QoS messages.
	*/
	Snmpqosqoserroripcrate int32 `json:"snmpqosqos_error_ipcrate,omitempty"`
	Snmpqosqosflowmem uint64 `json:"snmpqosqos_flow_mem,omitempty"`
	/**
	* Flow memory allocated
	*/
	Snmpqosqosflowmemrate int32 `json:"snmpqosqos_flow_memrate,omitempty"`
	Snmpqosqosflowsavailable uint64 `json:"snmpqosqos_flows_available,omitempty"`
	/**
	* Flows free list size
	*/
	Snmpqosqosflowsavailablerate int32 `json:"snmpqosqos_flows_availablerate,omitempty"`
	Snmpqosqosrecyclefailedbacklog uint64 `json:"snmpqosqos_recycle_failed_backlog,omitempty"`
	/**
	* Recycle failed due to backlog
	*/
	Snmpqosqosrecyclefailedbacklograte int32 `json:"snmpqosqos_recycle_failed_backlograte,omitempty"`
	Snmpqosqosrecyclefailedsession uint64 `json:"snmpqosqos_recycle_failed_session,omitempty"`
	/**
	* Recycle failed due to session attachment
	*/
	Snmpqosqosrecyclefailedsessionrate int32 `json:"snmpqosqos_recycle_failed_sessionrate,omitempty"`
	Snmpqosqoserrorcreateactionfailed uint64 `json:"snmpqosqos_error_create_action_failed,omitempty"`
	/**
	* Failed attempts to create actions
	*/
	Snmpqosqoserrorcreateactionfailedrate int32 `json:"snmpqosqos_error_create_action_failedrate,omitempty"`
	Snmpqosqoserrormodifyactionfailed uint64 `json:"snmpqosqos_error_modify_action_failed,omitempty"`
	/**
	* Failed attempts to modify actions
	*/
	Snmpqosqoserrormodifyactionfailedrate int32 `json:"snmpqosqos_error_modify_action_failedrate,omitempty"`
	Snmpqosqoserrorremoveactionfailed uint64 `json:"snmpqosqos_error_remove_action_failed,omitempty"`
	/**
	* Failed attempts to remove actions
	*/
	Snmpqosqoserrorremoveactionfailedrate int32 `json:"snmpqosqos_error_remove_action_failedrate,omitempty"`
	Snmpqosqoserrorcliunknown uint64 `json:"snmpqosqos_error_cli_unknown,omitempty"`
	/**
	* Internal CLI error
	*/
	Snmpqosqoserrorcliunknownrate int32 `json:"snmpqosqos_error_cli_unknownrate,omitempty"`
	Snmpqosqoserrorrenamenotimplemented uint64 `json:"snmpqosqos_error_rename_not_implemented,omitempty"`
	/**
	* qos action rename not yet implemented
	*/
	Snmpqosqoserrorrenamenotimplementedrate int32 `json:"snmpqosqos_error_rename_not_implementedrate,omitempty"`
	Snmpqosqoserrorremovepolicyfailed uint64 `json:"snmpqosqos_error_remove_policy_failed,omitempty"`
	/**
	* Failed attempts to remove qos policy
	*/
	Snmpqosqoserrorremovepolicyfailedrate int32 `json:"snmpqosqos_error_remove_policy_failedrate,omitempty"`
	Snmpqosqoserrorcreatepolicyfailed uint64 `json:"snmpqosqos_error_create_policy_failed,omitempty"`
	/**
	* Failed attempts to create qos policy
	*/
	Snmpqosqoserrorcreatepolicyfailedrate int32 `json:"snmpqosqos_error_create_policy_failedrate,omitempty"`
	Snmpqosqoserrorlibqosapifailures uint64 `json:"snmpqosqos_error_libqos_api_failures,omitempty"`
	/**
	* Libqos api failures
	*/
	Snmpqosqoserrorlibqosapifailuresrate int32 `json:"snmpqosqos_error_libqos_api_failuresrate,omitempty"`
	Snmpqosqoserrorapisesinvalidpcb uint64 `json:"snmpqosqos_error_api_ses_invalidpcb,omitempty"`
	/**
	* Libqos api qos_session_add_pcb/natpcb() failed for reason QS_EINVALIDPCB
	*/
	Snmpqosqoserrorapisesinvalidpcbrate int32 `json:"snmpqosqos_error_api_ses_invalidpcbrate,omitempty"`
	Snmpqosqoserrorapisesnotready uint64 `json:"snmpqosqos_error_api_ses_notready,omitempty"`
	/**
	* Libqos api qos_session_add_pcb/natpcb() failed for reason QS_ENOTREADY
	*/
	Snmpqosqoserrorapisesnotreadyrate int32 `json:"snmpqosqos_error_api_ses_notreadyrate,omitempty"`
	Snmpqosqoserrorapisesaddinsession uint64 `json:"snmpqosqos_error_api_ses_add_insession,omitempty"`
	/**
	* Libqos api qos_session_add_pcb/natpcb() failed for reason QS_EINSESSION
	*/
	Snmpqosqoserrorapisesaddinsessionrate int32 `json:"snmpqosqos_error_api_ses_add_insessionrate,omitempty"`
	Snmpqosqoserrorapisesaddother uint64 `json:"snmpqosqos_error_api_ses_add_other,omitempty"`
	/**
	* Libqos api qos_session_add_pcb/natpcb() failed
	*/
	Snmpqosqoserrorapisesaddotherrate int32 `json:"snmpqosqos_error_api_ses_add_otherrate,omitempty"`
	Snmpqosqoserrorapisesremnotinsession uint64 `json:"snmpqosqos_error_api_ses_rem_notinsession,omitempty"`
	/**
	* Libqos api qos_session_rem_pcb/natpcb() failed for reason QS_ENOTINSESSION
	*/
	Snmpqosqoserrorapisesremnotinsessionrate int32 `json:"snmpqosqos_error_api_ses_rem_notinsessionrate,omitempty"`
	Snmpqosqoserrorapisesremother uint64 `json:"snmpqosqos_error_api_ses_rem_other,omitempty"`
	/**
	* Libqos api qos_session_rem_pcb/natpcb() failed
	*/
	Snmpqosqoserrorapisesremotherrate int32 `json:"snmpqosqos_error_api_ses_rem_otherrate,omitempty"`
	Snmpqosqoserrorapisesdel uint64 `json:"snmpqosqos_error_api_ses_del,omitempty"`
	/**
	* Libqos api qos_session_delete faled
	*/
	Snmpqosqoserrorapisesdelrate int32 `json:"snmpqosqos_error_api_ses_delrate,omitempty"`
	Snmpqosqoserrornoflows uint64 `json:"snmpqosqos_error_no_flows,omitempty"`
	/**
	* Libqos out of flow memory
	*/
	Snmpqosqoserrornoflowsrate int32 `json:"snmpqosqos_error_no_flowsrate,omitempty"`
	Snmpqosqoserrornosessions uint64 `json:"snmpqosqos_error_no_sessions,omitempty"`
	/**
	* Libqos out of session memory
	*/
	Snmpqosqoserrornosessionsrate int32 `json:"snmpqosqos_error_no_sessionsrate,omitempty"`

}
