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

package feo


type Feostats struct {
	/**
	* Clear the statsistics / counters
	*/
	Clearstats string `json:"clearstats,omitempty"`
	Optcacheobjects uint32 `json:"optcacheobjects,omitempty"`
	/**
	* Total number of optimized cache objects ready to be served.
	*/
	Optcacheobjectsrate int32 `json:"optcacheobjectsrate,omitempty"`
	Origcacheobjects uint32 `json:"origcacheobjects,omitempty"`
	/**
	* Total number of original cache objects ready to be served.
	*/
	Origcacheobjectsrate int32 `json:"origcacheobjectsrate,omitempty"`
	Totalimgsdomainsharded uint32 `json:"totalimgsdomainsharded,omitempty"`
	/**
	* Total no of images whose domain has been set from shards.
	*/
	Imgsdomainshardedrate int32 `json:"imgsdomainshardedrate,omitempty"`
	Totalimgslazyloaded uint32 `json:"totalimgslazyloaded,omitempty"`
	/**
	* Total number of images modified for lazy loading.
	*/
	Imgslazyloadedrate int32 `json:"imgslazyloadedrate,omitempty"`
	Toturireplaced uint32 `json:"toturireplaced,omitempty"`
	/**
	* Total number of URI replaced.
	*/
	Urireplacedrate int32 `json:"urireplacedrate,omitempty"`
	Totalimgsinlinedincss uint32 `json:"totalimgsinlinedincss,omitempty"`
	/**
	* Total number of images inlined in CSS.
	*/
	Imgsinlinedincssrate int32 `json:"imgsinlinedincssrate,omitempty"`
	Totalinlinedjs uint32 `json:"totalinlinedjs,omitempty"`
	/**
	* Total number of inlined JS files.
	*/
	Inlinedjsrate int32 `json:"inlinedjsrate,omitempty"`
	Totalinlinedcss uint32 `json:"totalinlinedcss,omitempty"`
	/**
	* Total number of inlined CSS files.
	*/
	Inlinedcssrate int32 `json:"inlinedcssrate,omitempty"`
	Totalinlinedimgs uint32 `json:"totalinlinedimgs,omitempty"`
	/**
	* Total number of inlined images in HTML.
	*/
	Inlinedimgsrate int32 `json:"inlinedimgsrate,omitempty"`
	Totaldatasavings uint32 `json:"totaldatasavings,omitempty"`
	/**
	* Total data savings in bytes.
	*/
	Datasavingsrate int32 `json:"datasavingsrate,omitempty"`
	Htmlcommentsremoved uint32 `json:"htmlcommentsremoved,omitempty"`
	/**
	* The total number of HTML comments removed.
	*/
	Htmlcommentsremovedrate int32 `json:"htmlcommentsremovedrate,omitempty"`
	Totalcacheextended uint32 `json:"totalcacheextended,omitempty"`
	/**
	* The total number of objects cache extended.
	*/
	Cacheextendedrate int32 `json:"cacheextendedrate,omitempty"`
	Totalcsscombined uint32 `json:"totalcsscombined,omitempty"`
	/**
	* The total number of CSS combined.
	*/
	Csscombinedrate int32 `json:"csscombinedrate,omitempty"`
	Totalimporttolink uint32 `json:"totalimporttolink,omitempty"`
	/**
	* The total number of CSS imports converted to links
	*/
	Importtolinkrate int32 `json:"importtolinkrate,omitempty"`
	Totaljsmoved uint32 `json:"totaljsmoved,omitempty"`
	/**
	* The total number of JS moved to end.
	*/
	Jsmovedrate int32 `json:"jsmovedrate,omitempty"`
	Totalcssmoved uint32 `json:"totalcssmoved,omitempty"`
	/**
	* The total number of CSS moved to head.
	*/
	Cssmovedrate int32 `json:"cssmovedrate,omitempty"`
	Totaljsmin uint32 `json:"totaljsmin,omitempty"`
	/**
	* The total number of JS files minified.
	*/
	Jsminrate int32 `json:"jsminrate,omitempty"`
	Totalcssmin uint32 `json:"totalcssmin,omitempty"`
	/**
	* The total number of CSS files minified.
	*/
	Cssminrate int32 `json:"cssminrate,omitempty"`
	Totalimgstojxr uint32 `json:"totalimgstojxr,omitempty"`
	/**
	* The total number of images converted to JPEGXR format.
	*/
	Imgstojxrrate int32 `json:"imgstojxrrate,omitempty"`
	Totalimgstowebp uint32 `json:"totalimgstowebp,omitempty"`
	/**
	* The total number of images converted to WEBP format.
	*/
	Imgstowebprate int32 `json:"imgstowebprate,omitempty"`
	Totaljpegsoptimized uint32 `json:"totaljpegsoptimized,omitempty"`
	/**
	* The total number of JPEG format images optimized.
	*/
	Jpegsoptimizedrate int32 `json:"jpegsoptimizedrate,omitempty"`
	Totalgifstopng uint32 `json:"totalgifstopng,omitempty"`
	/**
	* The total number of images converted from GIF to PNG format.
	*/
	Gifstopngrate int32 `json:"gifstopngrate,omitempty"`
	Totalimgsresized uint32 `json:"totalimgsresized,omitempty"`
	/**
	* The total number of images resized to dimensions in the <img> tag.
	*/
	Imgsresizedrate int32 `json:"imgsresizedrate,omitempty"`

}
