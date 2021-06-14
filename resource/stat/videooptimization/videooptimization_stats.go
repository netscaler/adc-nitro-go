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

package videooptimization


type Videooptimizationstats struct {
	/**
	* Clear the statsistics / counters
	*/
	Clearstats string `json:"clearstats,omitempty"`
	Voctpdvideo uint64 `json:"voctpdvideo,omitempty"`
	/**
	* Total number of progressive download HTTP videos
	*/
	Voctpdvideorate int32 `json:"voctpdvideorate,omitempty"`
	Voctabvideo uint64 `json:"voctabvideo,omitempty"`
	/**
	* Total number of adaptive bitrate HTTP videos
	*/
	Voctabvideorate int32 `json:"voctabvideorate,omitempty"`
	Voeabvideo uint64 `json:"voeabvideo,omitempty"`
	/**
	* Total number of adaptive bitrate HTTPS videos
	*/
	Voeabvideorate int32 `json:"voeabvideorate,omitempty"`
	Voquicvideo uint64 `json:"voquicvideo,omitempty"`
	/**
	* Total number of QUIC videos
	*/
	Voquicvideorate int32 `json:"voquicvideorate,omitempty"`
	Vovideooptother uint64 `json:"vovideooptother,omitempty"`
	/**
	* Non video traffic
	*/
	Vovideoopherrate int32 `json:"vovideoopherrate,omitempty"`
	Voctabvideoses uint64 `json:"voctabvideoses,omitempty"`
	/**
	* Total number of adaptive bitrate HTTP video sessions
	*/
	Voctabvideosesrate int32 `json:"voctabvideosesrate,omitempty"`
	Voeabvideoses uint64 `json:"voeabvideoses,omitempty"`
	/**
	* Total number of adaptive bitrate HTTPS videos
	*/
	Voeabvideosesrate int32 `json:"voeabvideosesrate,omitempty"`
	Voquicvideoses uint64 `json:"voquicvideoses,omitempty"`
	/**
	* Total number of sessions with QUIC videos
	*/
	Voquicvideosesrate int32 `json:"voquicvideosesrate,omitempty"`
	Voctpdvideobytes uint64 `json:"voctpdvideobytes,omitempty"`
	/**
	* Total number of bytes for progressive download HTTP videos
	*/
	Voctpdvideobytesrate int32 `json:"voctpdvideobytesrate,omitempty"`
	Voctabvideobytes uint64 `json:"voctabvideobytes,omitempty"`
	/**
	* Total number of bytes for adaptive bitrate HTTP videos
	*/
	Voctabvideobytesrate int32 `json:"voctabvideobytesrate,omitempty"`
	Voeabvideobytes uint64 `json:"voeabvideobytes,omitempty"`
	/**
	* Total number of bytes for adaptive bitrate HTTPS videos
	*/
	Voeabvideobytesrate int32 `json:"voeabvideobytesrate,omitempty"`
	Voquicvideobytes uint64 `json:"voquicvideobytes,omitempty"`
	/**
	* Total number of bytes for QUIC videos
	*/
	Voquicvideobytesrate int32 `json:"voquicvideobytesrate,omitempty"`
	Vovideooptotherbytes uint64 `json:"vovideooptotherbytes,omitempty"`
	/**
	* Total number of bytes for non video traffic
	*/
	Vovideoopherbytesrate int32 `json:"vovideoopherbytesrate,omitempty"`

}
