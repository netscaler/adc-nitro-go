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
	Clearstats  string `json:"clearstats,omitempty"`
	Voctpdvideo int    `json:"voctpdvideo,omitempty"`
	/**
	* Total number of progressive download HTTP videos
	 */
	Voctpdvideorate float64 `json:"voctpdvideorate,omitempty"`
	Voctabvideo     int     `json:"voctabvideo,omitempty"`
	/**
	* Total number of adaptive bitrate HTTP videos
	 */
	Voctabvideorate float64 `json:"voctabvideorate,omitempty"`
	Voeabvideo      int     `json:"voeabvideo,omitempty"`
	/**
	* Total number of adaptive bitrate HTTPS videos
	 */
	Voeabvideorate float64 `json:"voeabvideorate,omitempty"`
	Voquicvideo    int     `json:"voquicvideo,omitempty"`
	/**
	* Total number of QUIC videos
	 */
	Voquicvideorate float64 `json:"voquicvideorate,omitempty"`
	Vovideooptother int     `json:"vovideooptother,omitempty"`
	/**
	* Non video traffic
	 */
	Vovideoopherrate float64 `json:"vovideoopherrate,omitempty"`
	Voctabvideoses   int     `json:"voctabvideoses,omitempty"`
	/**
	* Total number of adaptive bitrate HTTP video sessions
	 */
	Voctabvideosesrate float64 `json:"voctabvideosesrate,omitempty"`
	Voeabvideoses      int     `json:"voeabvideoses,omitempty"`
	/**
	* Total number of adaptive bitrate HTTPS videos
	 */
	Voeabvideosesrate float64 `json:"voeabvideosesrate,omitempty"`
	Voquicvideoses    int     `json:"voquicvideoses,omitempty"`
	/**
	* Total number of sessions with QUIC videos
	 */
	Voquicvideosesrate float64 `json:"voquicvideosesrate,omitempty"`
	Voctpdvideobytes   int     `json:"voctpdvideobytes,omitempty"`
	/**
	* Total number of bytes for progressive download HTTP videos
	 */
	Voctpdvideobytesrate float64 `json:"voctpdvideobytesrate,omitempty"`
	Voctabvideobytes     int     `json:"voctabvideobytes,omitempty"`
	/**
	* Total number of bytes for adaptive bitrate HTTP videos
	 */
	Voctabvideobytesrate float64 `json:"voctabvideobytesrate,omitempty"`
	Voeabvideobytes      int     `json:"voeabvideobytes,omitempty"`
	/**
	* Total number of bytes for adaptive bitrate HTTPS videos
	 */
	Voeabvideobytesrate float64 `json:"voeabvideobytesrate,omitempty"`
	Voquicvideobytes    int     `json:"voquicvideobytes,omitempty"`
	/**
	* Total number of bytes for QUIC videos
	 */
	Voquicvideobytesrate float64 `json:"voquicvideobytesrate,omitempty"`
	Vovideooptotherbytes int     `json:"vovideooptotherbytes,omitempty"`
	/**
	* Total number of bytes for non video traffic
	 */
	Vovideoopherbytesrate float64 `json:"vovideoopherbytesrate,omitempty"`
}
