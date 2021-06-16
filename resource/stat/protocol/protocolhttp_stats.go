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
* Statistics for HTTP resource.
*/

type Protocolhttpstats struct {
	/**
	* Clear the statsistics / counters
	*/
	Clearstats string `json:"clearstats,omitempty"`
	Spdytotstreams uint64 `json:"spdytotstreams,omitempty"`
	/**
	* Total number of requests received over SPDYv2 and SPDYv3
	*/
	Spdystreamsrate float64 `json:"spdystreamsrate,omitempty"`
	Httptotrequests uint64 `json:"httptotrequests,omitempty"`
	/**
	* Total number of HTTP requests received.
	*/
	Httprequestsrate float64 `json:"httprequestsrate,omitempty"`
	Httptotresponses uint64 `json:"httptotresponses,omitempty"`
	/**
	* Total number of HTTP responses sent.
	*/
	Httpresponsesrate float64 `json:"httpresponsesrate,omitempty"`
	Httptotrxrequestbytes uint64 `json:"httptotrxrequestbytes,omitempty"`
	/**
	* Total number of bytes of HTTP request data received.
	*/
	Httprxrequestbytesrate float64 `json:"httprxrequestbytesrate,omitempty"`
	Httptotrxresponsebytes uint64 `json:"httptotrxresponsebytes,omitempty"`
	/**
	* Total number of bytes of HTTP response data received.
	*/
	Httprxresponsebytesrate float64 `json:"httprxresponsebytesrate,omitempty"`
	Httptotgets uint64 `json:"httptotgets,omitempty"`
	/**
	* Total number of HTTP requests received with the GET method.
	*/
	Httpgetsrate float64 `json:"httpgetsrate,omitempty"`
	Httptotposts uint64 `json:"httptotposts,omitempty"`
	/**
	* Total number of HTTP requests received with the POST method.
	*/
	Httppostsrate float64 `json:"httppostsrate,omitempty"`
	Httptotothers uint64 `json:"httptotothers,omitempty"`
	/**
	* Total number of HTTP requests received with methods other than GET and POST. Some of the other well-defined HTTP methods are HEAD, PUT, DELETE, OPTIONS, and TRACE. User-defined methods are also allowed.
	*/
	Httptohersrate float64 `json:"httptohersrate,omitempty"`
	Httptot10requests uint64 `json:"httptot10requests,omitempty"`
	/**
	* Total number of HTTP/1.0 requests received.
	*/
	Http10requestsrate float64 `json:"http10requestsrate,omitempty"`
	Httptot11requests uint64 `json:"httptot11requests,omitempty"`
	/**
	* Total number of HTTP/1.1 requests received.
	*/
	Http11requestsrate float64 `json:"http11requestsrate,omitempty"`
	Httptotclenrequests uint64 `json:"httptotclenrequests,omitempty"`
	/**
	* Total number of HTTP requests in which the Content-length field of the HTTP header has been set. Content-length specifies the length of the content, in bytes, in the associated HTTP body.
	*/
	Httpclenrequestsrate float64 `json:"httpclenrequestsrate,omitempty"`
	Httptotchunkedrequests uint64 `json:"httptotchunkedrequests,omitempty"`
	/**
	* Total number of HTTP requests in which the Transfer-Encoding field of the HTTP header has been set to chunked.
	*/
	Httpchunkedrequestsrate float64 `json:"httpchunkedrequestsrate,omitempty"`
	Httptottxrequestbytes uint64 `json:"httptottxrequestbytes,omitempty"`
	/**
	* Total number of bytes of HTTP request data transmitted.
	*/
	Httptxrequestbytesrate float64 `json:"httptxrequestbytesrate,omitempty"`
	Httptot10responses uint64 `json:"httptot10responses,omitempty"`
	/**
	* Total number of HTTP/1.0 responses sent.
	*/
	Http10responsesrate float64 `json:"http10responsesrate,omitempty"`
	Httptot11responses uint64 `json:"httptot11responses,omitempty"`
	/**
	* Total number of HTTP/1.1 responses sent.
	*/
	Http11responsesrate float64 `json:"http11responsesrate,omitempty"`
	Httptotclenresponses uint64 `json:"httptotclenresponses,omitempty"`
	/**
	* Total number of HTTP responses sent in which the Content-length field of the HTTP header has been set. Content-length specifies the length of the content, in bytes, in the associated HTTP body.
	*/
	Httpclenresponsesrate float64 `json:"httpclenresponsesrate,omitempty"`
	Httptotchunkedresponses uint64 `json:"httptotchunkedresponses,omitempty"`
	/**
	* Total number of HTTP responses sent in which the Transfer-Encoding field of the HTTP header has been set to chunked. This setting is used when the server wants to start sending the response before knowing its total length. The server breaks the response into chunks and sends them in sequence, inserting the length of each chunk before the actual data. The message ends with a chunk of size zero.
	*/
	Httpchunkedresponsesrate float64 `json:"httpchunkedresponsesrate,omitempty"`
	Httperrnoreusemultipart uint64 `json:"httperrnoreusemultipart,omitempty"`
	/**
	* Total number of HTTP multi-part responses sent. In multi-part responses, one or more entities are encapsulated within the body of a single message.
	*/
	Httperrnoreusemultipartrate float64 `json:"httperrnoreusemultipartrate,omitempty"`
	Httptotnoclenchunkresponses uint64 `json:"httptotnoclenchunkresponses,omitempty"`
	/**
	* Total number of FIN-terminated responses sent. In FIN-terminated responses, the server finishes sending the data and closes the connection.
	*/
	Httpnoclenchunkresponsesrate float64 `json:"httpnoclenchunkresponsesrate,omitempty"`
	Httptottxresponsebytes uint64 `json:"httptottxresponsebytes,omitempty"`
	/**
	* Total number of bytes of HTTP response data transmitted.
	*/
	Httptxresponsebytesrate float64 `json:"httptxresponsebytesrate,omitempty"`
	Httperrincompleteheaders uint64 `json:"httperrincompleteheaders,omitempty"`
	Httperrincompleterequests uint64 `json:"httperrincompleterequests,omitempty"`
	/**
	* Total number of HTTP requests received in which the header spans more than one packet.
	*/
	Httperrincompleterequestsrate float64 `json:"httperrincompleterequestsrate,omitempty"`
	Httperrincompleteresponses uint64 `json:"httperrincompleteresponses,omitempty"`
	/**
	* Total number of HTTP responses received in which the header spans more than one packet.
	*/
	Httperrincompleteresponsesrate float64 `json:"httperrincompleteresponsesrate,omitempty"`
	Httperrserverbusy uint64 `json:"httperrserverbusy,omitempty"`
	/**
	* Total number of HTTP error responses received. Some of the error responses are: 
		500 	Internal Server Error
		501 	Not Implemented
		502 	Bad Gateway
		503 	Service Unavailable
		504 	Gateway Timeout
		505 	HTTP Version Not Supported
	*/
	Httperrserverbusyrate float64 `json:"httperrserverbusyrate,omitempty"`
	Httperrlargecontent uint64 `json:"httperrlargecontent,omitempty"`
	Httperrlargechunk uint64 `json:"httperrlargechunk,omitempty"`
	Httperrlargectlen uint64 `json:"httperrlargectlen,omitempty"`
	Httperrrfc7230desynctlente uint64 `json:"httperrrfc7230desynctlente,omitempty"`
	Httperrrfc7230desynmultictlen uint64 `json:"httperrrfc7230desynmultictlen,omitempty"`
	Httperrrfc7230desynidenticalctlen uint64 `json:"httperrrfc7230desynidenticalctlen,omitempty"`
	Spdyv2totstreams uint64 `json:"spdyv2totstreams,omitempty"`
	/**
	* Total number of requests received over SPDYv2
	*/
	Spdyv2streamsrate float64 `json:"spdyv2streamsrate,omitempty"`
	Spdyv3totstreams uint64 `json:"spdyv3totstreams,omitempty"`
	/**
	* Total number of requests received over SPDYv3
	*/
	Spdyv3streamsrate float64 `json:"spdyv3streamsrate,omitempty"`

}
