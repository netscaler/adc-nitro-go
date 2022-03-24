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

package stat

type Feostats struct {
	/**
	* Clear the statsistics / counters
	 */
	Clearstats      string `json:"clearstats,omitempty"`
	Optcacheobjects int    `json:"optcacheobjects,omitempty"`
	/**
	* Total number of optimized cache objects ready to be served.
	 */
	Optcacheobjectsrate float64 `json:"optcacheobjectsrate,omitempty"`
	Origcacheobjects    int     `json:"origcacheobjects,omitempty"`
	/**
	* Total number of original cache objects ready to be served.
	 */
	Origcacheobjectsrate   float64 `json:"origcacheobjectsrate,omitempty"`
	Totalimgsdomainsharded int     `json:"totalimgsdomainsharded,omitempty"`
	/**
	* Total no of images whose domain has been set from shards.
	 */
	Imgsdomainshardedrate float64 `json:"imgsdomainshardedrate,omitempty"`
	Totalimgslazyloaded   int     `json:"totalimgslazyloaded,omitempty"`
	/**
	* Total number of images modified for lazy loading.
	 */
	Imgslazyloadedrate float64 `json:"imgslazyloadedrate,omitempty"`
	Toturireplaced     int     `json:"toturireplaced,omitempty"`
	/**
	* Total number of URI replaced.
	 */
	Urireplacedrate       float64 `json:"urireplacedrate,omitempty"`
	Totalimgsinlinedincss int     `json:"totalimgsinlinedincss,omitempty"`
	/**
	* Total number of images inlined in CSS.
	 */
	Imgsinlinedincssrate float64 `json:"imgsinlinedincssrate,omitempty"`
	Totalinlinedjs       int     `json:"totalinlinedjs,omitempty"`
	/**
	* Total number of inlined JS files.
	 */
	Inlinedjsrate   float64 `json:"inlinedjsrate,omitempty"`
	Totalinlinedcss int     `json:"totalinlinedcss,omitempty"`
	/**
	* Total number of inlined CSS files.
	 */
	Inlinedcssrate   float64 `json:"inlinedcssrate,omitempty"`
	Totalinlinedimgs int     `json:"totalinlinedimgs,omitempty"`
	/**
	* Total number of inlined images in HTML.
	 */
	Inlinedimgsrate  float64 `json:"inlinedimgsrate,omitempty"`
	Totaldatasavings int     `json:"totaldatasavings,omitempty"`
	/**
	* Total data savings in bytes.
	 */
	Datasavingsrate     float64 `json:"datasavingsrate,omitempty"`
	Htmlcommentsremoved int     `json:"htmlcommentsremoved,omitempty"`
	/**
	* The total number of HTML comments removed.
	 */
	Htmlcommentsremovedrate float64 `json:"htmlcommentsremovedrate,omitempty"`
	Totalcacheextended      int     `json:"totalcacheextended,omitempty"`
	/**
	* The total number of objects cache extended.
	 */
	Cacheextendedrate float64 `json:"cacheextendedrate,omitempty"`
	Totalcsscombined  int     `json:"totalcsscombined,omitempty"`
	/**
	* The total number of CSS combined.
	 */
	Csscombinedrate   float64 `json:"csscombinedrate,omitempty"`
	Totalimporttolink int     `json:"totalimporttolink,omitempty"`
	/**
	* The total number of CSS imports converted to links
	 */
	Importtolinkrate float64 `json:"importtolinkrate,omitempty"`
	Totaljsmoved     int     `json:"totaljsmoved,omitempty"`
	/**
	* The total number of JS moved to end.
	 */
	Jsmovedrate   float64 `json:"jsmovedrate,omitempty"`
	Totalcssmoved int     `json:"totalcssmoved,omitempty"`
	/**
	* The total number of CSS moved to head.
	 */
	Cssmovedrate float64 `json:"cssmovedrate,omitempty"`
	Totaljsmin   int     `json:"totaljsmin,omitempty"`
	/**
	* The total number of JS files minified.
	 */
	Jsminrate   float64 `json:"jsminrate,omitempty"`
	Totalcssmin int     `json:"totalcssmin,omitempty"`
	/**
	* The total number of CSS files minified.
	 */
	Cssminrate     float64 `json:"cssminrate,omitempty"`
	Totalimgstojxr int     `json:"totalimgstojxr,omitempty"`
	/**
	* The total number of images converted to JPEGXR format.
	 */
	Imgstojxrrate   float64 `json:"imgstojxrrate,omitempty"`
	Totalimgstowebp int     `json:"totalimgstowebp,omitempty"`
	/**
	* The total number of images converted to WEBP format.
	 */
	Imgstowebprate      float64 `json:"imgstowebprate,omitempty"`
	Totaljpegsoptimized int     `json:"totaljpegsoptimized,omitempty"`
	/**
	* The total number of JPEG format images optimized.
	 */
	Jpegsoptimizedrate float64 `json:"jpegsoptimizedrate,omitempty"`
	Totalgifstopng     int     `json:"totalgifstopng,omitempty"`
	/**
	* The total number of images converted from GIF to PNG format.
	 */
	Gifstopngrate    float64 `json:"gifstopngrate,omitempty"`
	Totalimgsresized int     `json:"totalimgsresized,omitempty"`
	/**
	* The total number of images resized to dimensions in the <img> tag.
	 */
	Imgsresizedrate float64 `json:"imgsresizedrate,omitempty"`
}
