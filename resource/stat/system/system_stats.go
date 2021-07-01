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

package system


type Systemstats struct {
	/**
	* Clear the statsistics / counters
	*/
	Clearstats string `json:"clearstats,omitempty"`
	Voltagev12n float64 `json:"voltagev12n,omitempty"`
	Voltagev5n float64 `json:"voltagev5n,omitempty"`
	Cpuusage int `json:"cpuusage,omitempty"`
	Rescpuusage int `json:"rescpuusage,omitempty"`
	Slavecpuusage int `json:"slavecpuusage,omitempty"`
	Mastercpuusage int `json:"mastercpuusage,omitempty"`
	Auxvolt7 float64 `json:"auxvolt7,omitempty"`
	Auxvolt6 float64 `json:"auxvolt6,omitempty"`
	Auxvolt5 float64 `json:"auxvolt5,omitempty"`
	Auxvolt4 float64 `json:"auxvolt4,omitempty"`
	Auxvolt3 float64 `json:"auxvolt3,omitempty"`
	Auxvolt2 float64 `json:"auxvolt2,omitempty"`
	Auxvolt1 float64 `json:"auxvolt1,omitempty"`
	Auxvolt0 float64 `json:"auxvolt0,omitempty"`
	Voltagevsen2 float64 `json:"voltagevsen2,omitempty"`
	Voltagev5sb float64 `json:"voltagev5sb,omitempty"`
	Voltagevtt float64 `json:"voltagevtt,omitempty"`
	Voltagevbat float64 `json:"voltagevbat,omitempty"`
	Voltagev12p float64 `json:"voltagev12p,omitempty"`
	Voltagev5p float64 `json:"voltagev5p,omitempty"`
	Voltagev33stby float64 `json:"voltagev33stby,omitempty"`
	Voltagev33main float64 `json:"voltagev33main,omitempty"`
	Voltagevcc1 float64 `json:"voltagevcc1,omitempty"`
	Voltagevcc0 float64 `json:"voltagevcc0,omitempty"`
	Numcpus int `json:"numcpus,omitempty"`
	Memusagepcnt float64 `json:"memusagepcnt,omitempty"`
	Memuseinmb int `json:"memuseinmb,omitempty"`
	Addimgmtcpuusagepcnt float64 `json:"addimgmtcpuusagepcnt,omitempty"`
	Mgmtcpu0usagepcnt float64 `json:"mgmtcpu0usagepcnt,omitempty"`
	Mgmtcpuusagepcnt float64 `json:"mgmtcpuusagepcnt,omitempty"`
	Pktcpuusagepcnt float64 `json:"pktcpuusagepcnt,omitempty"`
	Cpuusagepcnt float64 `json:"cpuusagepcnt,omitempty"`
	Rescpuusagepcnt float64 `json:"rescpuusagepcnt,omitempty"`
	Starttimelocal string `json:"starttimelocal,omitempty"`
	Starttime string `json:"starttime,omitempty"`
	Disk0perusage int `json:"disk0perusage,omitempty"`
	Disk1perusage int `json:"disk1perusage,omitempty"`
	Cpufan0speed int `json:"cpufan0speed,omitempty"`
	Cpufan1speed int `json:"cpufan1speed,omitempty"`
	Systemfanspeed int `json:"systemfanspeed,omitempty"`
	Fan0speed int `json:"fan0speed,omitempty"`
	Fanspeed int `json:"fanspeed,omitempty"`
	Cpu0temp int `json:"cpu0temp,omitempty"`
	Cpu1temp int `json:"cpu1temp,omitempty"`
	Internaltemp int `json:"internaltemp,omitempty"`
	Powersupply1status string `json:"powersupply1status,omitempty"`
	Powersupply2status string `json:"powersupply2status,omitempty"`
	Powersupply3status string `json:"powersupply3status,omitempty"`
	Powersupply4status string `json:"powersupply4status,omitempty"`
	Disk0size int `json:"disk0size,omitempty"`
	Disk0used int `json:"disk0used,omitempty"`
	Disk0avail int `json:"disk0avail,omitempty"`
	Disk1size int `json:"disk1size,omitempty"`
	Disk1used int `json:"disk1used,omitempty"`
	Disk1avail int `json:"disk1avail,omitempty"`
	Fan2speed int `json:"fan2speed,omitempty"`
	Fan3speed int `json:"fan3speed,omitempty"`
	Fan4speed int `json:"fan4speed,omitempty"`
	Fan5speed int `json:"fan5speed,omitempty"`
	Auxtemp0 int `json:"auxtemp0,omitempty"`
	Auxtemp1 int `json:"auxtemp1,omitempty"`
	Auxtemp2 int `json:"auxtemp2,omitempty"`
	Auxtemp3 int `json:"auxtemp3,omitempty"`
	Timesincestart string `json:"timesincestart,omitempty"`
	Memsizemb int `json:"memsizemb,omitempty"`

}
