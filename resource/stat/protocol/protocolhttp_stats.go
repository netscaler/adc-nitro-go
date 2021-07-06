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
	Spdytotstreams int `json:"spdytotstreams,omitempty"`
	/**
	* Total number of requests received over SPDYv2 and SPDYv3
	*/
	Spdystreamsrate float64 `json:"spdystreamsrate,omitempty"`
	Httptotrequests int `json:"httptotrequests,omitempty"`
	/**
	* Total number of HTTP requests received.
	*/
	Httprequestsrate float64 `json:"httprequestsrate,omitempty"`
	Httptotresponses int `json:"httptotresponses,omitempty"`
	/**
	* Total number of HTTP responses sent.
	*/
	Httpresponsesrate float64 `json:"httpresponsesrate,omitempty"`
	Httptotrxrequestbytes int `json:"httptotrxrequestbytes,omitempty"`
	/**
	* Total number of bytes of HTTP request data received.
	*/
	Httprxrequestbytesrate float64 `json:"httprxrequestbytesrate,omitempty"`
	Httptotrxresponsebytes int `json:"httptotrxresponsebytes,omitempty"`
	/**
	* Total number of bytes of HTTP response data received.
	*/
	Httprxresponsebytesrate float64 `json:"httprxresponsebytesrate,omitempty"`
	Httptotgets int `json:"httptotgets,omitempty"`
	/**
	* Total number of HTTP requests received with the GET method.
	*/
	Httpgetsrate float64 `json:"httpgetsrate,omitempty"`
	Httptotposts int `json:"httptotposts,omitempty"`
	/**
	* Total number of HTTP requests received with the POST method.
	*/
	Httppostsrate float64 `json:"httppostsrate,omitempty"`
	Httptotothers int `json:"httptotothers,omitempty"`
	/**
	* Total number of HTTP requests received with methods other than GET and POST. Some of the other well-defined HTTP methods are HEAD, PUT, DELETE, OPTIONS, and TRACE. User-defined methods are also allowed.
	*/
	Httptohersrate float64 `json:"httptohersrate,omitempty"`
	Httptot10requests int `json:"httptot10requests,omitempty"`
	/**
	* Total number of HTTP/1.0 requests received.
	*/
	Http10requestsrate float64 `json:"http10requestsrate,omitempty"`
	Httptot11requests int `json:"httptot11requests,omitempty"`
	/**
	* Total number of HTTP/1.1 requests received.
	*/
	Http11requestsrate float64 `json:"http11requestsrate,omitempty"`
	Httptotclenrequests int `json:"httptotclenrequests,omitempty"`
	/**
	* Total number of HTTP requests in which the Content-length field of the HTTP header has been set. Content-length specifies the length of the content, in bytes, in the associated HTTP body.
	*/
	Httpclenrequestsrate float64 `json:"httpclenrequestsrate,omitempty"`
	Httptotchunkedrequests int `json:"httptotchunkedrequests,omitempty"`
	/**
	* Total number of HTTP requests in which the Transfer-Encoding field of the HTTP header has been set to chunked.
	*/
	Httpchunkedrequestsrate float64 `json:"httpchunkedrequestsrate,omitempty"`
	Httptottxrequestbytes int `json:"httptottxrequestbytes,omitempty"`
	/**
	* Total number of bytes of HTTP request data transmitted.
	*/
	Httptxrequestbytesrate float64 `json:"httptxrequestbytesrate,omitempty"`
	Httptot10responses int `json:"httptot10responses,omitempty"`
	/**
	* Total number of HTTP/1.0 responses sent.
	*/
	Http10responsesrate float64 `json:"http10responsesrate,omitempty"`
	Httptot11responses int `json:"httptot11responses,omitempty"`
	/**
	* Total number of HTTP/1.1 responses sent.
	*/
	Http11responsesrate float64 `json:"http11responsesrate,omitempty"`
	Httptotclenresponses int `json:"httptotclenresponses,omitempty"`
	/**
	* Total number of HTTP responses sent in which the Content-length field of the HTTP header has been set. Content-length specifies the length of the content, in bytes, in the associated HTTP body.
	*/
	Httpclenresponsesrate float64 `json:"httpclenresponsesrate,omitempty"`
	Httptotchunkedresponses int `json:"httptotchunkedresponses,omitempty"`
	/**
	* Total number of HTTP responses sent in which the Transfer-Encoding field of the HTTP header has been set to chunked. This setting is used when the server wants to start sending the response before knowing its total length. The server breaks the response into chunks and sends them in sequence, inserting the length of each chunk before the actual data. The message ends with a chunk of size zero.
	*/
	Httpchunkedresponsesrate float64 `json:"httpchunkedresponsesrate,omitempty"`
	Httperrnoreusemultipart int `json:"httperrnoreusemultipart,omitempty"`
	/**
	* Total number of HTTP multi-part responses sent. In multi-part responses, one or more entities are encapsulated within the body of a single message.
	*/
	Httperrnoreusemultipartrate float64 `json:"httperrnoreusemultipartrate,omitempty"`
	Httptotnoclenchunkresponses int `json:"httptotnoclenchunkresponses,omitempty"`
	/**
	* Total number of FIN-terminated responses sent. In FIN-terminated responses, the server finishes sending the data and closes the connection.
	*/
	Httpnoclenchunkresponsesrate float64 `json:"httpnoclenchunkresponsesrate,omitempty"`
	Httptottxresponsebytes int `json:"httptottxresponsebytes,omitempty"`
	/**
	* Total number of bytes of HTTP response data transmitted.
	*/
	Httptxresponsebytesrate float64 `json:"httptxresponsebytesrate,omitempty"`
	Httperrincompleteheaders int `json:"httperrincompleteheaders,omitempty"`
	Httperrincompleterequests int `json:"httperrincompleterequests,omitempty"`
	/**
	* Total number of HTTP requests received in which the header spans more than one packet.
	*/
	Httperrincompleterequestsrate float64 `json:"httperrincompleterequestsrate,omitempty"`
	Httperrincompleteresponses int `json:"httperrincompleteresponses,omitempty"`
	/**
	* Total number of HTTP responses received in which the header spans more than one packet.
	*/
	Httperrincompleteresponsesrate float64 `json:"httperrincompleteresponsesrate,omitempty"`
	Httperrserverbusy int `json:"httperrserverbusy,omitempty"`
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
	Httperrlargecontent int `json:"httperrlargecontent,omitempty"`
	Httperrlargechunk int `json:"httperrlargechunk,omitempty"`
	Httperrlargectlen int `json:"httperrlargectlen,omitempty"`
	Httperrrfc7230desynctlente int `json:"httperrrfc7230desynctlente,omitempty"`
	Httperrrfc7230desynmultictlen int `json:"httperrrfc7230desynmultictlen,omitempty"`
	Httperrrfc7230desynidenticalctlen int `json:"httperrrfc7230desynidenticalctlen,omitempty"`
	Spdyv2totstreams int `json:"spdyv2totstreams,omitempty"`
	/**
	* Total number of requests received over SPDYv2
	*/
	Spdyv2streamsrate float64 `json:"spdyv2streamsrate,omitempty"`
	Spdyv3totstreams int `json:"spdyv3totstreams,omitempty"`
	/**
	* Total number of requests received over SPDYv3
	*/
	Spdyv3streamsrate float64 `json:"spdyv3streamsrate,omitempty"`

}
