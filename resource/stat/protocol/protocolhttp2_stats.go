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
* Statistics for http2 resource.
*/

type Protocolhttp2stats struct {
	/**
	* Clear the statsistics / counters
	*/
	Clearstats string `json:"clearstats,omitempty"`
	Http2requests uint64 `json:"http2requests,omitempty"`
	/**
	* Total number of http2 requests
	*/
	Http2requestsrate int32 `json:"http2requestsrate,omitempty"`
	Http2responses uint64 `json:"http2responses,omitempty"`
	/**
	* Total number of http2 responses
	*/
	Http2responsesrate int32 `json:"http2responsesrate,omitempty"`
	Http2totgrpcrequest uint64 `json:"http2totgrpcrequest,omitempty"`
	/**
	* Total number of gRPC requests
	*/
	Http2grpcrequestrate int32 `json:"http2grpcrequestrate,omitempty"`
	Http2totgrpcresponse uint64 `json:"http2totgrpcresponse,omitempty"`
	/**
	* Total number of gRPC responses
	*/
	Http2grpcresponserate int32 `json:"http2grpcresponserate,omitempty"`
	Http2totgrpcsuccess uint64 `json:"http2totgrpcsuccess,omitempty"`
	/**
	* Total number of gRPC success
	*/
	Http2grpcsuccessrate int32 `json:"http2grpcsuccessrate,omitempty"`
	Http2totgrpcfailure uint64 `json:"http2totgrpcfailure,omitempty"`
	/**
	* Total number of gRPC failures
	*/
	Http2grpcfailurerate int32 `json:"http2grpcfailurerate,omitempty"`
	Http2direct uint64 `json:"http2direct,omitempty"`
	/**
	* Total number of http2 direct connections established
	*/
	Http2directrate int32 `json:"http2directrate,omitempty"`
	Http2serverdirect uint64 `json:"http2serverdirect,omitempty"`
	/**
	* Number of HTTP/2 server direct
	*/
	Http2serverdirectrate int32 `json:"http2serverdirectrate,omitempty"`
	Http2requpg uint64 `json:"http2requpg,omitempty"`
	/**
	* Total number of connections upgraded to HTTP2
	*/
	Http2requpgrate int32 `json:"http2requpgrate,omitempty"`
	Http2nomatcipher uint64 `json:"http2nomatcipher,omitempty"`
	/**
	* Total number of cipher mismatch failures
	*/
	Http2nomatcipherrate int32 `json:"http2nomatcipherrate,omitempty"`
	Http2serverdirectfailed uint64 `json:"http2serverdirectfailed,omitempty"`
	/**
	* Number of HTTP/2 server direct failed
	*/
	Http2serverdirectfailedrate int32 `json:"http2serverdirectfailedrate,omitempty"`
	Http2serverupgradefailed uint64 `json:"http2serverupgradefailed,omitempty"`
	/**
	* Number of HTTP/2 server upgrade failed
	*/
	Http2serverupgradefailedrate int32 `json:"http2serverupgradefailedrate,omitempty"`
	Http2requestupgradefailed uint64 `json:"http2requestupgradefailed,omitempty"`
	/**
	* Number of HTTP/2 request upgrade failed
	*/
	Http2requestupgradefailedrate int32 `json:"http2requestupgradefailedrate,omitempty"`
	Http2dataframessent uint64 `json:"http2dataframessent,omitempty"`
	/**
	* Number of HTTP/2 DATA frames sent
	*/
	Http2dataframessentrate int32 `json:"http2dataframessentrate,omitempty"`
	Http2headerframessent uint64 `json:"http2headerframessent,omitempty"`
	/**
	* Number of HTTP/2 HEADER frames sent
	*/
	Http2headerframessentrate int32 `json:"http2headerframessentrate,omitempty"`
	Http2priorityframessent uint64 `json:"http2priorityframessent,omitempty"`
	/**
	* Number of HTTP/2 PRIORITY frames sent
	*/
	Http2priorityframessentrate int32 `json:"http2priorityframessentrate,omitempty"`
	Http2rststreamframessent uint64 `json:"http2rststreamframessent,omitempty"`
	/**
	* Number of HTTP/2 RST_STREAM frames sent
	*/
	Http2rststreamframessentrate int32 `json:"http2rststreamframessentrate,omitempty"`
	Http2settingframessent uint64 `json:"http2settingframessent,omitempty"`
	/**
	* Number of HTTP/2 SETTINGS frames sent
	*/
	Http2settingframessentrate int32 `json:"http2settingframessentrate,omitempty"`
	Http2pushpromiseframessent uint64 `json:"http2pushpromiseframessent,omitempty"`
	/**
	* Number of HTTP/2 PUSH_PROMISE frames sent
	*/
	Http2pushpromiseframessentrate int32 `json:"http2pushpromiseframessentrate,omitempty"`
	Http2pingframessent uint64 `json:"http2pingframessent,omitempty"`
	/**
	* Number of HTTP/2 PING frames sent
	*/
	Http2pingframessentrate int32 `json:"http2pingframessentrate,omitempty"`
	Http2goawayframessent uint64 `json:"http2goawayframessent,omitempty"`
	/**
	* Number of HTTP/2 GOAWAY frames sent
	*/
	Http2goawayframessentrate int32 `json:"http2goawayframessentrate,omitempty"`
	Http2windowupdateframessent uint64 `json:"http2windowupdateframessent,omitempty"`
	/**
	* Number of HTTP/2 WINDOW_UPDATE frames sent
	*/
	Http2windowupdateframessentrate int32 `json:"http2windowupdateframessentrate,omitempty"`
	Http2continuationframessent uint64 `json:"http2continuationframessent,omitempty"`
	/**
	* Number of HTTP/2 CONTINUATION frames sent
	*/
	Http2continuationframessentrate int32 `json:"http2continuationframessentrate,omitempty"`
	Http2altsvcframessent uint64 `json:"http2altsvcframessent,omitempty"`
	/**
	* Number of HTTP/2 ALTSVC frames sent
	*/
	Http2altsvcframessentrate int32 `json:"http2altsvcframessentrate,omitempty"`
	Http2dataframesrcvd uint64 `json:"http2dataframesrcvd,omitempty"`
	/**
	* Number of HTTP/2 DATA frames received
	*/
	Http2dataframesrcvdrate int32 `json:"http2dataframesrcvdrate,omitempty"`
	Http2headerframesrcvd uint64 `json:"http2headerframesrcvd,omitempty"`
	/**
	* Total number of http2 header frames received
	*/
	Http2headerframesrcvdrate int32 `json:"http2headerframesrcvdrate,omitempty"`
	Http2priorityframesrcvd uint64 `json:"http2priorityframesrcvd,omitempty"`
	/**
	* Total number of http2 priority frames received
	*/
	Http2priorityframesrcvdrate int32 `json:"http2priorityframesrcvdrate,omitempty"`
	Http2rststreamframesrcvd uint64 `json:"http2rststreamframesrcvd,omitempty"`
	/**
	* Number of HTTP/2 RST_STREAM frames received
	*/
	Http2rststreamframesrcvdrate int32 `json:"http2rststreamframesrcvdrate,omitempty"`
	Http2settingframesrcvd uint64 `json:"http2settingframesrcvd,omitempty"`
	/**
	* Number of HTTP/2 SETTINGS frames received
	*/
	Http2settingframesrcvdrate int32 `json:"http2settingframesrcvdrate,omitempty"`
	Http2pushpromframesrcvd uint64 `json:"http2pushpromframesrcvd,omitempty"`
	/**
	* Number of HTTP/2 PUSH_PROMISE frames received
	*/
	Http2pushpromframesrcvdrate int32 `json:"http2pushpromframesrcvdrate,omitempty"`
	Http2pingframesrcvd uint64 `json:"http2pingframesrcvd,omitempty"`
	/**
	* Number of HTTP/2 PING frames received
	*/
	Http2pingframesrcvdrate int32 `json:"http2pingframesrcvdrate,omitempty"`
	Http2goawayframesrcvd uint64 `json:"http2goawayframesrcvd,omitempty"`
	/**
	* Number of HTTP/2 GOAWAY frames received
	*/
	Http2goawayframesrcvdrate int32 `json:"http2goawayframesrcvdrate,omitempty"`
	Http2winupdateframesrcvd uint64 `json:"http2winupdateframesrcvd,omitempty"`
	/**
	* Number of HTTP/2 WINDOW_UPDATE frames received
	*/
	Http2winupdateframesrcvdrate int32 `json:"http2winupdateframesrcvdrate,omitempty"`
	Http2continuationframesrcvd uint64 `json:"http2continuationframesrcvd,omitempty"`
	/**
	* Number of HTTP/2 CONTINUATION frames received
	*/
	Http2continuationframesrcvdrate int32 `json:"http2continuationframesrcvdrate,omitempty"`
	Http2indataframes uint64 `json:"http2indataframes,omitempty"`
	/**
	* Number of HTTP/2 DATA frames
	*/
	Http2indataframesrate int32 `json:"http2indataframesrate,omitempty"`
	Http2inheaderframes uint64 `json:"http2inheaderframes,omitempty"`
	/**
	* Number of HTTP/2 HEADER frames
	*/
	Http2inheaderframesrate int32 `json:"http2inheaderframesrate,omitempty"`
	Http2inpriorityframes uint64 `json:"http2inpriorityframes,omitempty"`
	/**
	* Number of HTTP/2 PRIORITY frames
	*/
	Http2inpriorityframesrate int32 `json:"http2inpriorityframesrate,omitempty"`
	Http2inrststreamframes uint64 `json:"http2inrststreamframes,omitempty"`
	/**
	* Number of HTTP/2 RST_STREAM frames
	*/
	Http2inrststreamframesrate int32 `json:"http2inrststreamframesrate,omitempty"`
	Http2insettingframes uint64 `json:"http2insettingframes,omitempty"`
	/**
	* Number of HTTP/2 SETTINGS frames
	*/
	Http2insettingframesrate int32 `json:"http2insettingframesrate,omitempty"`
	Http2inpushpromiseframes uint64 `json:"http2inpushpromiseframes,omitempty"`
	/**
	* Number of HTTP/2 PUSH_PROMISE frames
	*/
	Http2inpushpromiseframesrate int32 `json:"http2inpushpromiseframesrate,omitempty"`
	Http2inpingframes uint64 `json:"http2inpingframes,omitempty"`
	/**
	* Number of HTTP/2 PING frames
	*/
	Http2inpingframesrate int32 `json:"http2inpingframesrate,omitempty"`
	Http2ingoawayframes uint64 `json:"http2ingoawayframes,omitempty"`
	/**
	* Number of HTTP/2 GOAWAY frames
	*/
	Http2ingoawayframesrate int32 `json:"http2ingoawayframesrate,omitempty"`
	Http2inwindowupdateframes uint64 `json:"http2inwindowupdateframes,omitempty"`
	/**
	* Number of HTTP/2 WINDOW_UPDATE frames
	*/
	Http2inwindowupdateframesrate int32 `json:"http2inwindowupdateframesrate,omitempty"`
	Http2incontinuationframes uint64 `json:"http2incontinuationframes,omitempty"`
	/**
	* Number of HTTP/2 CONTINUATION frames
	*/
	Http2incontinuationframesrate int32 `json:"http2incontinuationframesrate,omitempty"`
	Http2frametoobig uint64 `json:"http2frametoobig,omitempty"`
	/**
	* Number of HTTP/2 frames received carrying a frame length greater than SETTINGS_MAX_FRAME_SIZE sent by NetScale 
	*/
	Http2frametoobigrate int32 `json:"http2frametoobigrate,omitempty"`
	Http2pingflood uint64 `json:"http2pingflood,omitempty"`
	/**
	* HTTP/2 number of ping frames received on connection is above rate limit
	*/
	Http2pingfloodrate int32 `json:"http2pingfloodrate,omitempty"`
	Http2errsetflood uint64 `json:"http2errsetflood,omitempty"`
	/**
	* HTTP/2 number of settings frames received on connection is above rate limit
	*/
	Http2errsetfloodrate int32 `json:"http2errsetfloodrate,omitempty"`
	Http2errresfraflood uint64 `json:"http2errresfraflood,omitempty"`
	/**
	* HTTP/2 number of reset frames received on connection is above rate limit
	*/
	Http2errresfrafloodrate int32 `json:"http2errresfrafloodrate,omitempty"`
	Http2errempfraflood uint64 `json:"http2errempfraflood,omitempty"`
	/**
	* HTTP/2 number of empty frames received on connection is above rate limit
	*/
	Http2errempfrafloodrate int32 `json:"http2errempfrafloodrate,omitempty"`

}
