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

package cache


type Cachestats struct {
	/**
	* Clear the statsistics / counters
	*/
	Clearstats string `json:"clearstats,omitempty"`
	Cachemaxmemorykb uint32 `json:"cachemaxmemorykb,omitempty"`
	Cacherecentpercentsuccessfulrevalidation uint32 `json:"cacherecentpercentsuccessfulrevalidation,omitempty"`
	Cacherecentpercentstoreablemiss uint32 `json:"cacherecentpercentstoreablemiss,omitempty"`
	Cacherecentpercentparameterizedhits uint32 `json:"cacherecentpercentparameterizedhits,omitempty"`
	Cacherecentpercentoriginbandwidthsaved uint32 `json:"cacherecentpercentoriginbandwidthsaved,omitempty"`
	Cacherecentpercenthit uint32 `json:"cacherecentpercenthit,omitempty"`
	Cacherecentpercentbytehit uint32 `json:"cacherecentpercentbytehit,omitempty"`
	Cacherecentpercent304hits uint32 `json:"cacherecentpercent304hits,omitempty"`
	Cacheutilizedmemorykb uint32 `json:"cacheutilizedmemorykb,omitempty"`
	Cachemaxmemoryactivekb uint64 `json:"cachemaxmemoryactivekb,omitempty"`
	Cache64maxmemorykb uint64 `json:"cache64maxmemorykb,omitempty"`
	Cachepercentpethits uint32 `json:"cachepercentpethits,omitempty"`
	Cachetotpethits uint64 `json:"cachetotpethits,omitempty"`
	/**
	* Number of times a cache hit was found during a search of a content group that has Poll Every Time enabled.
	*/
	Cachepethitsrate int32 `json:"cachepethitsrate,omitempty"`
	Cachepercentparameterized304hits uint32 `json:"cachepercentparameterized304hits,omitempty"`
	Cachetotparameterizedhits uint64 `json:"cachetotparameterizedhits,omitempty"`
	/**
	* Parameterized requests resulting in either a 304 or non-304 hit.
	*/
	Cacheparameterizedhitsrate int32 `json:"cacheparameterizedhitsrate,omitempty"`
	Cachepercentsuccessfulrevalidation uint32 `json:"cachepercentsuccessfulrevalidation,omitempty"`
	Cachepercentstoreablemiss uint32 `json:"cachepercentstoreablemiss,omitempty"`
	Cachetotfulltoconditionalrequest uint64 `json:"cachetotfulltoconditionalrequest,omitempty"`
	/**
	* Number of user-agent requests for a cached  Poll Every Time (PET) response that were sent to the origin server as conditional requests. 
	*/
	Cachefulltoconditionalrequestrate int32 `json:"cachefulltoconditionalrequestrate,omitempty"`
	Cachetotsuccessfulrevalidation uint64 `json:"cachetotsuccessfulrevalidation,omitempty"`
	/**
	* Total number of times stored content was successfully revalidated by a 304 Not Modified response from the origin.
	*/
	Cachesuccessfulrevalidationrate int32 `json:"cachesuccessfulrevalidationrate,omitempty"`
	Cachetotrevalidationmiss uint64 `json:"cachetotrevalidationmiss,omitempty"`
	/**
	* Responses that an intervening cache revalidated with the integrated cache before serving, as determined by a Cache-Control: Max-Age header configurable in the integrated cache
	*/
	Cacherevalidationmissrate int32 `json:"cacherevalidationmissrate,omitempty"`
	Cachetotnonstoreablemisses uint64 `json:"cachetotnonstoreablemisses,omitempty"`
	/**
	* Cache misses for which the fetched response is not stored in the cache. These responses match policies with a NOCACHE action or are affected by Poll Every Time.
	*/
	Cachenonstoreablemissesrate int32 `json:"cachenonstoreablemissesrate,omitempty"`
	Cachetotstoreablemisses uint64 `json:"cachetotstoreablemisses,omitempty"`
	/**
	* Cache misses for which the fetched response is stored in the cache before serving it to the client. Storable misses conform to a built-in or user-defined caching policy that contains a CACHE action.
	*/
	Cachestoreablemissesrate int32 `json:"cachestoreablemissesrate,omitempty"`
	Cachecompressedbytesserved uint64 `json:"cachecompressedbytesserved,omitempty"`
	/**
	* Number of compressed bytes served from the cache
	*/
	Cachecompressedbytesservedrate int32 `json:"cachecompressedbytesservedrate,omitempty"`
	Cachepercentbytehit uint32 `json:"cachepercentbytehit,omitempty"`
	Cachebytesserved uint64 `json:"cachebytesserved,omitempty"`
	/**
	* Total number of bytes served from the integrated cache
	*/
	Cachebytesservedrate int32 `json:"cachebytesservedrate,omitempty"`
	Cachetotresponsebytes uint64 `json:"cachetotresponsebytes,omitempty"`
	/**
	* Total number of HTTP response bytes served by Citrix ADC from both the origin and the cache
	*/
	Cacheresponsebytesrate int32 `json:"cacheresponsebytesrate,omitempty"`
	Cachepercent304hits uint32 `json:"cachepercent304hits,omitempty"`
	Cachenummarker uint64 `json:"cachenummarker,omitempty"`
	Cachepercentoriginbandwidthsaved uint32 `json:"cachepercentoriginbandwidthsaved,omitempty"`
	Cachepercenthit uint32 `json:"cachepercenthit,omitempty"`
	Cachetotmisses uint64 `json:"cachetotmisses,omitempty"`
	/**
	* Intercepted HTTP requests requiring fetches from origin server.
	*/
	Cachemissesrate int32 `json:"cachemissesrate,omitempty"`
	Cachetothits uint64 `json:"cachetothits,omitempty"`
	/**
	* Responses served from the integrated cache. These responses match a policy with a CACHE action.
	*/
	Cachehitsrate int32 `json:"cachehitsrate,omitempty"`
	Cachetotrequests []uint64 `json:"cachetotrequests,omitempty"`
	/**
	* Total cache hits plus total cache misses.
	*/
	Cacherequestsrate int32 `json:"cacherequestsrate,omitempty"`
	Cachenumcached uint64 `json:"cachenumcached,omitempty"`
	Cachecurhits uint64 `json:"cachecurhits,omitempty"`
	Cachecurmisses uint64 `json:"cachecurmisses,omitempty"`
	Cachetotnon304hits []uint64 `json:"cachetotnon304hits,omitempty"`
	/**
	* Total number of full (non-304) responses served from the cache. A 304 status code indicates that a response has not been modified since the last time it was served
	*/
	Cachenon304hitsrate int32 `json:"cachenon304hitsrate,omitempty"`
	Cachetot304hits []uint64 `json:"cachetot304hits,omitempty"`
	/**
	* Object not modified responses served from the cache. (Status code 304 served instead of the full response.)
	*/
	Cache304hitsrate int32 `json:"cache304hitsrate,omitempty"`
	Cachetotsqlhits []uint64 `json:"cachetotsqlhits,omitempty"`
	/**
	* sql response served from cache
	*/
	Cachesqlhitsrate int32 `json:"cachesqlhitsrate,omitempty"`
	Cachetotexpireatlastbyte uint64 `json:"cachetotexpireatlastbyte,omitempty"`
	/**
	* Instances of content expiring immediately after receiving the last body byte due to the Expire at Last Byte setting for the content group.
	*/
	Cacheexpireatlastbyterate int32 `json:"cacheexpireatlastbyterate,omitempty"`
	Cachetotflashcachemisses uint64 `json:"cachetotflashcachemisses,omitempty"`
	/**
	* Number of requests to a content group with flash cache enabled that were cache misses. Flash cache distributes the response to all the clients in aqueue.
	*/
	Cacheflashcachemissesrate int32 `json:"cacheflashcachemissesrate,omitempty"`
	Cachetotflashcachehits uint64 `json:"cachetotflashcachehits,omitempty"`
	/**
	* Number of requests to a content group with flash cache enabled that were cache hits. The flash cache setting queues requests that arrive simultaneously and distributes the response to all the clients in the queue.
	*/
	Cacheflashcachehitsrate int32 `json:"cacheflashcachehitsrate,omitempty"`
	Cachetotparameterizedinvalidationrequests uint64 `json:"cachetotparameterizedinvalidationrequests,omitempty"`
	/**
	* Requests matching a policy with an invalidation (INVAL) action and a content group that uses an invalidation selector or parameters.
	*/
	Cacheparameterizedinvalidationrequestsrate int32 `json:"cacheparameterizedinvalidationrequestsrate,omitempty"`
	Cachetotnonparameterizedinvalidationrequests uint64 `json:"cachetotnonparameterizedinvalidationrequests,omitempty"`
	/**
	* Requests that match an invalidation policy where the invalGroups parameter is configured and expires one or more content groups.
	*/
	Cachenonparameterizedinvalidationrequestsrate int32 `json:"cachenonparameterizedinvalidationrequestsrate,omitempty"`
	Cachetotinvalidationrequests uint64 `json:"cachetotinvalidationrequests,omitempty"`
	/**
	* Requests that match an invalidation policy and result in expiration of specific cached responses or entire content groups.
	*/
	Cacheinvalidationrequestsrate int32 `json:"cacheinvalidationrequestsrate,omitempty"`
	Cachetotparameterizedrequests uint64 `json:"cachetotparameterizedrequests,omitempty"`
	/**
	* Total number of requests where the content group has hit and invalidation parameters or selectors.
	*/
	Cacheparameterizedrequestsrate int32 `json:"cacheparameterizedrequestsrate,omitempty"`
	Cachetotparameterizednon304hits []uint64 `json:"cachetotparameterizednon304hits,omitempty"`
	/**
	* Parameterized requests resulting in a full response (not status code 304: Object Not Updated) served from the cache.
	*/
	Cacheparameterizednon304hitsrate int32 `json:"cacheparameterizednon304hitsrate,omitempty"`
	Cachetotparameterized304hits []uint64 `json:"cachetotparameterized304hits,omitempty"`
	/**
	* Parameterized requests resulting in an object not modified (status code 304) response. 
	*/
	Cacheparameterized304hitsrate int32 `json:"cacheparameterized304hitsrate,omitempty"`
	Cachetotpetrequests uint64 `json:"cachetotpetrequests,omitempty"`
	/**
	* Requests that triggered a search of a content group that has Poll Every Time (PET) enabled (always consult the origin server before serving cached data).
	*/
	Cachepetrequestsrate int32 `json:"cachepetrequestsrate,omitempty"`
	Cacheerrmemalloc uint64 `json:"cacheerrmemalloc,omitempty"`
	Cachelargestresponsereceived uint64 `json:"cachelargestresponsereceived,omitempty"`

}
