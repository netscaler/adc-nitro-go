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

type Cachestats struct {
	/**
	* Clear the statsistics / counters
	 */
	Clearstats                               string `json:"clearstats,omitempty"`
	Cachemaxmemorykb                         int    `json:"cachemaxmemorykb,omitempty"`
	Cacherecentpercentsuccessfulrevalidation int    `json:"cacherecentpercentsuccessfulrevalidation,omitempty"`
	Cacherecentpercentstoreablemiss          int    `json:"cacherecentpercentstoreablemiss,omitempty"`
	Cacherecentpercentparameterizedhits      int    `json:"cacherecentpercentparameterizedhits,omitempty"`
	Cacherecentpercentoriginbandwidthsaved   int    `json:"cacherecentpercentoriginbandwidthsaved,omitempty"`
	Cacherecentpercenthit                    int    `json:"cacherecentpercenthit,omitempty"`
	Cacherecentpercentbytehit                int    `json:"cacherecentpercentbytehit,omitempty"`
	Cacherecentpercent304hits                int    `json:"cacherecentpercent304hits,omitempty"`
	Cacheutilizedmemorykb                    int    `json:"cacheutilizedmemorykb,omitempty"`
	Cachemaxmemoryactivekb                   int    `json:"cachemaxmemoryactivekb,omitempty"`
	Cache64maxmemorykb                       int    `json:"cache64maxmemorykb,omitempty"`
	Cachepercentpethits                      int    `json:"cachepercentpethits,omitempty"`
	Cachetotpethits                          int    `json:"cachetotpethits,omitempty"`
	/**
	* Number of times a cache hit was found during a search of a content group that has Poll Every Time enabled.
	 */
	Cachepethitsrate                 float64 `json:"cachepethitsrate,omitempty"`
	Cachepercentparameterized304hits int     `json:"cachepercentparameterized304hits,omitempty"`
	Cachetotparameterizedhits        int     `json:"cachetotparameterizedhits,omitempty"`
	/**
	* Parameterized requests resulting in either a 304 or non-304 hit.
	 */
	Cacheparameterizedhitsrate         float64 `json:"cacheparameterizedhitsrate,omitempty"`
	Cachepercentsuccessfulrevalidation int     `json:"cachepercentsuccessfulrevalidation,omitempty"`
	Cachepercentstoreablemiss          int     `json:"cachepercentstoreablemiss,omitempty"`
	Cachetotfulltoconditionalrequest   int     `json:"cachetotfulltoconditionalrequest,omitempty"`
	/**
	* Number of user-agent requests for a cached  Poll Every Time (PET) response that were sent to the origin server as conditional requests.
	 */
	Cachefulltoconditionalrequestrate float64 `json:"cachefulltoconditionalrequestrate,omitempty"`
	Cachetotsuccessfulrevalidation    int     `json:"cachetotsuccessfulrevalidation,omitempty"`
	/**
	* Total number of times stored content was successfully revalidated by a 304 Not Modified response from the origin.
	 */
	Cachesuccessfulrevalidationrate float64 `json:"cachesuccessfulrevalidationrate,omitempty"`
	Cachetotrevalidationmiss        int     `json:"cachetotrevalidationmiss,omitempty"`
	/**
	* Responses that an intervening cache revalidated with the integrated cache before serving, as determined by a Cache-Control: Max-Age header configurable in the integrated cache
	 */
	Cacherevalidationmissrate  float64 `json:"cacherevalidationmissrate,omitempty"`
	Cachetotnonstoreablemisses int     `json:"cachetotnonstoreablemisses,omitempty"`
	/**
	* Cache misses for which the fetched response is not stored in the cache. These responses match policies with a NOCACHE action or are affected by Poll Every Time.
	 */
	Cachenonstoreablemissesrate float64 `json:"cachenonstoreablemissesrate,omitempty"`
	Cachetotstoreablemisses     int     `json:"cachetotstoreablemisses,omitempty"`
	/**
	* Cache misses for which the fetched response is stored in the cache before serving it to the client. Storable misses conform to a built-in or user-defined caching policy that contains a CACHE action.
	 */
	Cachestoreablemissesrate   float64 `json:"cachestoreablemissesrate,omitempty"`
	Cachecompressedbytesserved int     `json:"cachecompressedbytesserved,omitempty"`
	/**
	* Number of compressed bytes served from the cache
	 */
	Cachecompressedbytesservedrate float64 `json:"cachecompressedbytesservedrate,omitempty"`
	Cachepercentbytehit            int     `json:"cachepercentbytehit,omitempty"`
	Cachebytesserved               int     `json:"cachebytesserved,omitempty"`
	/**
	* Total number of bytes served from the integrated cache
	 */
	Cachebytesservedrate  float64 `json:"cachebytesservedrate,omitempty"`
	Cachetotresponsebytes int     `json:"cachetotresponsebytes,omitempty"`
	/**
	* Total number of HTTP response bytes served by Citrix ADC from both the origin and the cache
	 */
	Cacheresponsebytesrate           float64 `json:"cacheresponsebytesrate,omitempty"`
	Cachepercent304hits              int     `json:"cachepercent304hits,omitempty"`
	Cachenummarker                   int     `json:"cachenummarker,omitempty"`
	Cachepercentoriginbandwidthsaved int     `json:"cachepercentoriginbandwidthsaved,omitempty"`
	Cachepercenthit                  int     `json:"cachepercenthit,omitempty"`
	Cachetotmisses                   int     `json:"cachetotmisses,omitempty"`
	/**
	* Intercepted HTTP requests requiring fetches from origin server.
	 */
	Cachemissesrate float64 `json:"cachemissesrate,omitempty"`
	Cachetothits    int     `json:"cachetothits,omitempty"`
	/**
	* Responses served from the integrated cache. These responses match a policy with a CACHE action.
	 */
	Cachehitsrate    float64 `json:"cachehitsrate,omitempty"`
	Cachetotrequests []int   `json:"cachetotrequests,omitempty"`
	/**
	* Total cache hits plus total cache misses.
	 */
	Cacherequestsrate  float64 `json:"cacherequestsrate,omitempty"`
	Cachenumcached     int     `json:"cachenumcached,omitempty"`
	Cachecurhits       int     `json:"cachecurhits,omitempty"`
	Cachecurmisses     int     `json:"cachecurmisses,omitempty"`
	Cachetotnon304hits []int   `json:"cachetotnon304hits,omitempty"`
	/**
	* Total number of full (non-304) responses served from the cache. A 304 status code indicates that a response has not been modified since the last time it was served
	 */
	Cachenon304hitsrate float64 `json:"cachenon304hitsrate,omitempty"`
	Cachetot304hits     []int   `json:"cachetot304hits,omitempty"`
	/**
	* Object not modified responses served from the cache. (Status code 304 served instead of the full response.)
	 */
	Cache304hitsrate float64 `json:"cache304hitsrate,omitempty"`
	Cachetotsqlhits  []int   `json:"cachetotsqlhits,omitempty"`
	/**
	* sql response served from cache
	 */
	Cachesqlhitsrate         float64 `json:"cachesqlhitsrate,omitempty"`
	Cachetotexpireatlastbyte int     `json:"cachetotexpireatlastbyte,omitempty"`
	/**
	* Instances of content expiring immediately after receiving the last body byte due to the Expire at Last Byte setting for the content group.
	 */
	Cacheexpireatlastbyterate float64 `json:"cacheexpireatlastbyterate,omitempty"`
	Cachetotflashcachemisses  int     `json:"cachetotflashcachemisses,omitempty"`
	/**
	* Number of requests to a content group with flash cache enabled that were cache misses. Flash cache distributes the response to all the clients in aqueue.
	 */
	Cacheflashcachemissesrate float64 `json:"cacheflashcachemissesrate,omitempty"`
	Cachetotflashcachehits    int     `json:"cachetotflashcachehits,omitempty"`
	/**
	* Number of requests to a content group with flash cache enabled that were cache hits. The flash cache setting queues requests that arrive simultaneously and distributes the response to all the clients in the queue.
	 */
	Cacheflashcachehitsrate                   float64 `json:"cacheflashcachehitsrate,omitempty"`
	Cachetotparameterizedinvalidationrequests int     `json:"cachetotparameterizedinvalidationrequests,omitempty"`
	/**
	* Requests matching a policy with an invalidation (INVAL) action and a content group that uses an invalidation selector or parameters.
	 */
	Cacheparameterizedinvalidationrequestsrate   float64 `json:"cacheparameterizedinvalidationrequestsrate,omitempty"`
	Cachetotnonparameterizedinvalidationrequests int     `json:"cachetotnonparameterizedinvalidationrequests,omitempty"`
	/**
	* Requests that match an invalidation policy where the invalGroups parameter is configured and expires one or more content groups.
	 */
	Cachenonparameterizedinvalidationrequestsrate float64 `json:"cachenonparameterizedinvalidationrequestsrate,omitempty"`
	Cachetotinvalidationrequests                  int     `json:"cachetotinvalidationrequests,omitempty"`
	/**
	* Requests that match an invalidation policy and result in expiration of specific cached responses or entire content groups.
	 */
	Cacheinvalidationrequestsrate float64 `json:"cacheinvalidationrequestsrate,omitempty"`
	Cachetotparameterizedrequests int     `json:"cachetotparameterizedrequests,omitempty"`
	/**
	* Total number of requests where the content group has hit and invalidation parameters or selectors.
	 */
	Cacheparameterizedrequestsrate  float64 `json:"cacheparameterizedrequestsrate,omitempty"`
	Cachetotparameterizednon304hits []int   `json:"cachetotparameterizednon304hits,omitempty"`
	/**
	* Parameterized requests resulting in a full response (not status code 304: Object Not Updated) served from the cache.
	 */
	Cacheparameterizednon304hitsrate float64 `json:"cacheparameterizednon304hitsrate,omitempty"`
	Cachetotparameterized304hits     []int   `json:"cachetotparameterized304hits,omitempty"`
	/**
	* Parameterized requests resulting in an object not modified (status code 304) response.
	 */
	Cacheparameterized304hitsrate float64 `json:"cacheparameterized304hitsrate,omitempty"`
	Cachetotpetrequests           int     `json:"cachetotpetrequests,omitempty"`
	/**
	* Requests that triggered a search of a content group that has Poll Every Time (PET) enabled (always consult the origin server before serving cached data).
	 */
	Cachepetrequestsrate         float64 `json:"cachepetrequestsrate,omitempty"`
	Cacheerrmemalloc             int     `json:"cacheerrmemalloc,omitempty"`
	Cachelargestresponsereceived int     `json:"cachelargestresponsereceived,omitempty"`
}
