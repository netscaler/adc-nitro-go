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

package ssl


type Sslstats struct {
	/**
	* Clear the statsistics / counters
	*/
	Clearstats string `json:"clearstats,omitempty"`
	Ssltothwdecbesecondary uint64 `json:"ssltothwdecbesecondary,omitempty"`
	/**
	* Number of bytes decrypted on the back-end in hardware on secondary card.
	*/
	Sslhwdecbesecondaryrate int32 `json:"sslhwdecbesecondaryrate,omitempty"`
	Ssltothwdecfesecondary uint64 `json:"ssltothwdecfesecondary,omitempty"`
	/**
	* Number of bytes decrypted on the front-end in hardware on secondary card.
	*/
	Sslhwdecfesecondaryrate int32 `json:"sslhwdecfesecondaryrate,omitempty"`
	Ssltothwencbesecondary uint64 `json:"ssltothwencbesecondary,omitempty"`
	/**
	* Number of bytes encrypted on the back-end in hardware on secondary card.
	*/
	Sslhwencbesecondaryrate int32 `json:"sslhwencbesecondaryrate,omitempty"`
	Ssltothwencfesecondary uint64 `json:"ssltothwencfesecondary,omitempty"`
	/**
	* Number of bytes encrypted on the front-end in hardware on secondary card.
	*/
	Sslhwencfesecondaryrate int32 `json:"sslhwencfesecondaryrate,omitempty"`
	Ssltotdechwsecondary uint64 `json:"ssltotdechwsecondary,omitempty"`
	/**
	* Number of bytes decrypted in hardware on secondary card.
	*/
	Ssldechwsecondaryrate int32 `json:"ssldechwsecondaryrate,omitempty"`
	Ssltotenchwsecondary uint64 `json:"ssltotenchwsecondary,omitempty"`
	/**
	* Number of bytes encrypted in hardware on secondary card.
	*/
	Sslenchwsecondaryrate int32 `json:"sslenchwsecondaryrate,omitempty"`
	Sslcryptoutilizationsymmstat uint32 `json:"sslcryptoutilizationsymmstat,omitempty"`
	Sslcryptoutilizationasymstat uint32 `json:"sslcryptoutilizationasymstat,omitempty"`
	Sslcryptoutilizationstat2nd float64 `json:"sslcryptoutilizationstat2nd,omitempty"`
	Sslcryptoutilizationstat float64 `json:"sslcryptoutilizationstat,omitempty"`
	Sslnumcardsupsecondary uint32 `json:"sslnumcardsupsecondary,omitempty"`
	Sslcardssecondary uint32 `json:"sslcardssecondary,omitempty"`
	Sslcardstatus uint64 `json:"sslcardstatus,omitempty"`
	Sslcards uint32 `json:"sslcards,omitempty"`
	Sslnumcardsup uint32 `json:"sslnumcardsup,omitempty"`
	Sslenginestatus uint32 `json:"sslenginestatus,omitempty"`
	Ssltotsessions uint64 `json:"ssltotsessions,omitempty"`
	/**
	* Number of SSL sessions on the Citrix ADC.
	*/
	Sslsessionsrate int32 `json:"sslsessionsrate,omitempty"`
	Ssltottransactions uint64 `json:"ssltottransactions,omitempty"`
	/**
	* Number of SSL transactions on the Citrix ADC
	*/
	Ssltransactionsrate int32 `json:"ssltransactionsrate,omitempty"`
	Ssltotsslv2transactions uint64 `json:"ssltotsslv2transactions,omitempty"`
	/**
	* Number of SSLv2 transactions on the Citrix ADC.
	*/
	Sslsslv2transactionsrate int32 `json:"sslsslv2transactionsrate,omitempty"`
	Ssltotsslv3transactions uint64 `json:"ssltotsslv3transactions,omitempty"`
	/**
	* Total number of SSLv3 transactions on the Citrix ADC.
	*/
	Sslsslv3transactionsrate int32 `json:"sslsslv3transactionsrate,omitempty"`
	Ssltottlsv1transactions uint64 `json:"ssltottlsv1transactions,omitempty"`
	/**
	* Number of TLSv1 transactions on the Citrix ADC.
	*/
	Ssltlsv1transactionsrate int32 `json:"ssltlsv1transactionsrate,omitempty"`
	Ssltottlsv11transactions uint64 `json:"ssltottlsv11transactions,omitempty"`
	/**
	* Number of TLSv1.1 transactions on the Citrix ADC.
	*/
	Ssltlsv11transactionsrate int32 `json:"ssltlsv11transactionsrate,omitempty"`
	Ssltottlsv12transactions uint64 `json:"ssltottlsv12transactions,omitempty"`
	/**
	* Number of TLSv1.2 transactions on the Citrix ADC.
	*/
	Ssltlsv12transactionsrate int32 `json:"ssltlsv12transactionsrate,omitempty"`
	Ssltottlsv13transactions uint64 `json:"ssltottlsv13transactions,omitempty"`
	/**
	* Number of TLSv1.3 transactions on the Citrix ADC.
	*/
	Ssltlsv13transactionsrate int32 `json:"ssltlsv13transactionsrate,omitempty"`
	Ssltotdtlsv1transactions uint64 `json:"ssltotdtlsv1transactions,omitempty"`
	/**
	* Number of DTLSv1 transactions on the Citrix ADC.
	*/
	Ssldtlsv1transactionsrate int32 `json:"ssldtlsv1transactionsrate,omitempty"`
	Ssltotdtlsv12transactions uint64 `json:"ssltotdtlsv12transactions,omitempty"`
	/**
	* Number of DTLSv1.2 transactions on the Citrix ADC.
	*/
	Ssldtlsv12transactionsrate int32 `json:"ssldtlsv12transactionsrate,omitempty"`
	Ssltotsslv2sessions uint64 `json:"ssltotsslv2sessions,omitempty"`
	/**
	* Number of SSLv2 sessions on the Citrix ADC.
	*/
	Sslsslv2sessionsrate int32 `json:"sslsslv2sessionsrate,omitempty"`
	Ssltotsslv3sessions uint64 `json:"ssltotsslv3sessions,omitempty"`
	/**
	* Number of SSLv3 sessions on the Citrix ADC.
	*/
	Sslsslv3sessionsrate int32 `json:"sslsslv3sessionsrate,omitempty"`
	Ssltottlsv1sessions uint64 `json:"ssltottlsv1sessions,omitempty"`
	/**
	* Number of TLSv1 sessions on the Citrix ADC.
	*/
	Ssltlsv1sessionsrate int32 `json:"ssltlsv1sessionsrate,omitempty"`
	Ssltottlsv11sessions uint64 `json:"ssltottlsv11sessions,omitempty"`
	/**
	* Number of TLSv1.1 sessions on the Citrix ADC.
	*/
	Ssltlsv11sessionsrate int32 `json:"ssltlsv11sessionsrate,omitempty"`
	Ssltottlsv12sessions uint64 `json:"ssltottlsv12sessions,omitempty"`
	/**
	* Number of TLSv1.2 sessions on the Citrix ADC.
	*/
	Ssltlsv12sessionsrate int32 `json:"ssltlsv12sessionsrate,omitempty"`
	Ssltottlsv13sessions uint64 `json:"ssltottlsv13sessions,omitempty"`
	/**
	* Number of TLSv1.3 sessions on the Citrix ADC.
	*/
	Ssltlsv13sessionsrate int32 `json:"ssltlsv13sessionsrate,omitempty"`
	Ssltotdtlsv1sessions uint64 `json:"ssltotdtlsv1sessions,omitempty"`
	/**
	* Number of DTLSv1 sessions on the Citrix ADC.
	*/
	Ssldtlsv1sessionsrate int32 `json:"ssldtlsv1sessionsrate,omitempty"`
	Ssltotdtlsv12sessions uint64 `json:"ssltotdtlsv12sessions,omitempty"`
	/**
	* Number of DTLSv1.2 sessions on the Citrix ADC.
	*/
	Ssldtlsv12sessionsrate int32 `json:"ssldtlsv12sessionsrate,omitempty"`
	Ssltotnewsessions uint64 `json:"ssltotnewsessions,omitempty"`
	/**
	* Number of new SSL sessions created on the Citrix ADC.
	*/
	Sslnewsessionsrate int32 `json:"sslnewsessionsrate,omitempty"`
	Ssltotsessionmiss uint64 `json:"ssltotsessionmiss,omitempty"`
	/**
	* Number of SSL session reuse misses on the Citrix ADC.
	*/
	Sslsessionmissrate int32 `json:"sslsessionmissrate,omitempty"`
	Ssltotsessionhits uint64 `json:"ssltotsessionhits,omitempty"`
	/**
	* Number of SSL session reuse hits on the Citrix ADC.
	*/
	Sslsessionhitsrate int32 `json:"sslsessionhitsrate,omitempty"`
	Sslbetotsessions uint64 `json:"sslbetotsessions,omitempty"`
	/**
	* Number of back-end SSL sessions on the Citrix ADC.
	*/
	Sslbesessionsrate int32 `json:"sslbesessionsrate,omitempty"`
	Sslbetotsslv3sessions uint64 `json:"sslbetotsslv3sessions,omitempty"`
	/**
	* Number of back-end SSLv3 sessions on the Citrix ADC.
	*/
	Sslbesslv3sessionsrate int32 `json:"sslbesslv3sessionsrate,omitempty"`
	Sslbetottlsv1sessions uint64 `json:"sslbetottlsv1sessions,omitempty"`
	/**
	* Number of back-end TLSv1 sessions on the Citrix ADC.
	*/
	Sslbetlsv1sessionsrate int32 `json:"sslbetlsv1sessionsrate,omitempty"`
	Sslbetottlsv11sessions uint64 `json:"sslbetottlsv11sessions,omitempty"`
	/**
	* Number of back-end TLSv1.1 sessions on the Citrix ADC.
	*/
	Sslbetlsv11sessionsrate int32 `json:"sslbetlsv11sessionsrate,omitempty"`
	Sslbetottlsv12sessions uint64 `json:"sslbetottlsv12sessions,omitempty"`
	/**
	* Number of back-end TLSv1.2 sessions on the Citrix ADC.
	*/
	Sslbetlsv12sessionsrate int32 `json:"sslbetlsv12sessionsrate,omitempty"`
	Sslbetotdtlsv1sessions uint64 `json:"sslbetotdtlsv1sessions,omitempty"`
	/**
	* Number of back-end DTLSv1 sessions on the Citrix ADC.
	*/
	Sslbedtlsv1sessionsrate int32 `json:"sslbedtlsv1sessionsrate,omitempty"`
	Sslbetotsessionmultiplexattempts uint64 `json:"sslbetotsessionmultiplexattempts,omitempty"`
	/**
	* Number of back-end SSL session multiplex attempts on the Citrix ADC.
	*/
	Sslbesessionmultiplexattemptsrate int32 `json:"sslbesessionmultiplexattemptsrate,omitempty"`
	Sslbetotsessionmultiplexattemptsuccess uint64 `json:"sslbetotsessionmultiplexattemptsuccess,omitempty"`
	/**
	* Number of back-end SSL session multiplex successes on the Citrix ADC.
	*/
	Sslbesessionmultiplexattemptsuccessrate int32 `json:"sslbesessionmultiplexattemptsuccessrate,omitempty"`
	Sslbetotsessionmultiplexattemptfails uint64 `json:"sslbetotsessionmultiplexattemptfails,omitempty"`
	/**
	* Number of back-end SSL session multiplex failures on the Citrix ADC.
	*/
	Sslbesessionmultiplexattemptfailsrate int32 `json:"sslbesessionmultiplexattemptfailsrate,omitempty"`
	Ssltotenc uint64 `json:"ssltotenc,omitempty"`
	/**
	* Number of bytes encrypted on the Citrix ADC.
	*/
	Sslencrate int32 `json:"sslencrate,omitempty"`
	Ssltotdec uint64 `json:"ssltotdec,omitempty"`
	/**
	* Number of bytes decrypted on the Citrix ADC.
	*/
	Ssldecrate int32 `json:"ssldecrate,omitempty"`
	Ssltotsslserverinrecords uint64 `json:"ssltotsslserverinrecords,omitempty"`
	/**
	* Number server in record on the Citrix ADC.
	*/
	Sslsslserverinrecordsrate int32 `json:"sslsslserverinrecordsrate,omitempty"`
	Ssltotrenegsessions uint64 `json:"ssltotrenegsessions,omitempty"`
	/**
	* Number of SSL session renegotiations on the Citrix ADC.
	*/
	Sslrenegsessionsrate int32 `json:"sslrenegsessionsrate,omitempty"`
	Ssltotsslv3renegsessions uint64 `json:"ssltotsslv3renegsessions,omitempty"`
	/**
	* Number of session renegotiations done on SSLv3.
	*/
	Sslsslv3renegsessionsrate int32 `json:"sslsslv3renegsessionsrate,omitempty"`
	Ssltottlsv1renegsessions uint64 `json:"ssltottlsv1renegsessions,omitempty"`
	/**
	* Number of SSL session renegotiations done on TLSv1.
	*/
	Ssltlsv1renegsessionsrate int32 `json:"ssltlsv1renegsessionsrate,omitempty"`
	Ssltottlsv11renegsessions uint64 `json:"ssltottlsv11renegsessions,omitempty"`
	/**
	* Number of SSL session renegotiations done on TLSv1.1.
	*/
	Ssltlsv11renegsessionsrate int32 `json:"ssltlsv11renegsessionsrate,omitempty"`
	Ssltottlsv12renegsessions uint64 `json:"ssltottlsv12renegsessions,omitempty"`
	/**
	* Number of SSL session renegotiations done on TLSv1.2.
	*/
	Ssltlsv12renegsessionsrate int32 `json:"ssltlsv12renegsessionsrate,omitempty"`
	Ssltotdtlsv1renegsessions uint64 `json:"ssltotdtlsv1renegsessions,omitempty"`
	/**
	* Number of SSL session renegotiations done on DTLSv1.
	*/
	Ssldtlsv1renegsessionsrate int32 `json:"ssldtlsv1renegsessionsrate,omitempty"`
	Ssltotdtlsv12renegsessions uint64 `json:"ssltotdtlsv12renegsessions,omitempty"`
	/**
	* Number of SSL session renegotiations done on DTLSv1.2.
	*/
	Ssldtlsv12renegsessionsrate int32 `json:"ssldtlsv12renegsessionsrate,omitempty"`
	Ssltotrsa512keyexchanges uint64 `json:"ssltotrsa512keyexchanges,omitempty"`
	/**
	* Number of RSA 512-bit key exchanges on the Citrix ADC.
	*/
	Sslrsa512keyexchangesrate int32 `json:"sslrsa512keyexchangesrate,omitempty"`
	Ssltotrsa1024keyexchanges uint64 `json:"ssltotrsa1024keyexchanges,omitempty"`
	/**
	* Number of RSA 1024-bit key exchanges on the Citrix ADC.
	*/
	Sslrsa1024keyexchangesrate int32 `json:"sslrsa1024keyexchangesrate,omitempty"`
	Ssltotrsa2048keyexchanges uint64 `json:"ssltotrsa2048keyexchanges,omitempty"`
	/**
	* Number of RSA 2048-bit key exchanges on the Citrix ADC.
	*/
	Sslrsa2048keyexchangesrate int32 `json:"sslrsa2048keyexchangesrate,omitempty"`
	Ssltotrsa3072keyexchanges uint64 `json:"ssltotrsa3072keyexchanges,omitempty"`
	/**
	* Number of RSA 3072-bit key exchanges on the Citrix ADC.
	*/
	Sslrsa3072keyexchangesrate int32 `json:"sslrsa3072keyexchangesrate,omitempty"`
	Ssltotrsa4096keyexchanges uint64 `json:"ssltotrsa4096keyexchanges,omitempty"`
	/**
	* Number of RSA 4096-bit key exchanges on the Citrix ADC.
	*/
	Sslrsa4096keyexchangesrate int32 `json:"sslrsa4096keyexchangesrate,omitempty"`
	Ssltotdh512keyexchanges uint64 `json:"ssltotdh512keyexchanges,omitempty"`
	/**
	* Number of Diffie-Helman 512-bit key exchanges on the Citrix ADC.
	*/
	Ssldh512keyexchangesrate int32 `json:"ssldh512keyexchangesrate,omitempty"`
	Ssltotdh1024keyexchanges uint64 `json:"ssltotdh1024keyexchanges,omitempty"`
	/**
	* Number of Diffie-Helman 1024-bit key exchanges on the Citrix ADC.
	*/
	Ssldh1024keyexchangesrate int32 `json:"ssldh1024keyexchangesrate,omitempty"`
	Ssltotdh2048keyexchanges uint64 `json:"ssltotdh2048keyexchanges,omitempty"`
	/**
	* Number of Diffie-Helman 2048-bit key exchanges on the Citrix ADC.
	*/
	Ssldh2048keyexchangesrate int32 `json:"ssldh2048keyexchangesrate,omitempty"`
	Ssltotdh4096keyexchanges uint64 `json:"ssltotdh4096keyexchanges,omitempty"`
	/**
	* Number of Diffie-Helman 4096-bit key exchanges on the Citrix ADC.
	*/
	Ssldh4096keyexchangesrate int32 `json:"ssldh4096keyexchangesrate,omitempty"`
	Ssltotecdhe521keyexchanges uint64 `json:"ssltotecdhe521keyexchanges,omitempty"`
	/**
	* Number of 521 Elliptical Curve Diffie-Helman on the Citrix ADC.
	*/
	Sslecdhe521keyexchangesrate int32 `json:"sslecdhe521keyexchangesrate,omitempty"`
	Ssltotecdhe384keyexchanges uint64 `json:"ssltotecdhe384keyexchanges,omitempty"`
	/**
	* Number of 384 Elliptical Curve Diffie-Helman on the Citrix ADC.
	*/
	Sslecdhe384keyexchangesrate int32 `json:"sslecdhe384keyexchangesrate,omitempty"`
	Ssltotecdhe256keyexchanges uint64 `json:"ssltotecdhe256keyexchanges,omitempty"`
	/**
	* Number of 256 Elliptical Curve Diffie-Helman on the Citrix ADC.
	*/
	Sslecdhe256keyexchangesrate int32 `json:"sslecdhe256keyexchangesrate,omitempty"`
	Ssltotecdhe224keyexchanges uint64 `json:"ssltotecdhe224keyexchanges,omitempty"`
	/**
	* Number of 224 Elliptical Curve Diffie-Helman on the Citrix ADC.
	*/
	Sslecdhe224keyexchangesrate int32 `json:"sslecdhe224keyexchangesrate,omitempty"`
	Ssltotecdhetransactions uint64 `json:"ssltotecdhetransactions,omitempty"`
	/**
	* Total ECDHE Transactions on Citrix ADC.
	*/
	Sslecdhetransactionsrate int32 `json:"sslecdhetransactionsrate,omitempty"`
	Ssltot40bitrc4ciphers uint64 `json:"ssltot40bitrc4ciphers,omitempty"`
	/**
	* Number of RC4 40-bit cipher encryptions on the Citrix ADC.
	*/
	Ssl40bitrc4ciphersrate int32 `json:"ssl40bitrc4ciphersrate,omitempty"`
	Ssltot56bitrc4ciphers uint64 `json:"ssltot56bitrc4ciphers,omitempty"`
	/**
	* Number of RC4 56-bit cipher encryptions on the Citrix ADC.
	*/
	Ssl56bitrc4ciphersrate int32 `json:"ssl56bitrc4ciphersrate,omitempty"`
	Ssltot64bitrc4ciphers uint64 `json:"ssltot64bitrc4ciphers,omitempty"`
	/**
	* Number of RC4 64-bit cipher encryptions on the Citrix ADC.
	*/
	Ssl64bitrc4ciphersrate int32 `json:"ssl64bitrc4ciphersrate,omitempty"`
	Ssltot128bitrc4ciphers uint64 `json:"ssltot128bitrc4ciphers,omitempty"`
	/**
	* Number of RC4 128-bit cipher encryptions on the Citrix ADC.
	*/
	Ssl128bitrc4ciphersrate int32 `json:"ssl128bitrc4ciphersrate,omitempty"`
	Ssltot40bitdesciphers uint64 `json:"ssltot40bitdesciphers,omitempty"`
	/**
	* Number of DES 40-bit cipher encryptions on the Citrix ADC.
	*/
	Ssl40bitdesciphersrate int32 `json:"ssl40bitdesciphersrate,omitempty"`
	Ssltot56bitdesciphers uint64 `json:"ssltot56bitdesciphers,omitempty"`
	/**
	* Number of DES 56-bit cipher encryptions on the Citrix ADC.
	*/
	Ssl56bitdesciphersrate int32 `json:"ssl56bitdesciphersrate,omitempty"`
	Ssltot168bit3desciphers uint64 `json:"ssltot168bit3desciphers,omitempty"`
	/**
	* Number of DES 168-bit cipher encryptions on the Citrix ADC.
	*/
	Ssl168bit3desciphersrate int32 `json:"ssl168bit3desciphersrate,omitempty"`
	Ssltotcipheraes128 uint64 `json:"ssltotcipheraes128,omitempty"`
	/**
	* Number of AES 128-bit cipher encryptions on the Citrix ADC.
	*/
	Sslcipheraes128rate int32 `json:"sslcipheraes128rate,omitempty"`
	Ssltotcipheraes256 uint64 `json:"ssltotcipheraes256,omitempty"`
	/**
	* Number of AES 256-bit cipher encryptions on the Citrix ADC.
	*/
	Sslcipheraes256rate int32 `json:"sslcipheraes256rate,omitempty"`
	Ssltot40bitrc2ciphers uint64 `json:"ssltot40bitrc2ciphers,omitempty"`
	/**
	* Number of RC2 40-bit cipher encryptions on the Citrix ADC.
	*/
	Ssl40bitrc2ciphersrate int32 `json:"ssl40bitrc2ciphersrate,omitempty"`
	Ssltot56bitrc2ciphers uint64 `json:"ssltot56bitrc2ciphers,omitempty"`
	/**
	* Number of RC2 56-bit cipher encryptions on the Citrix ADC.
	*/
	Ssl56bitrc2ciphersrate int32 `json:"ssl56bitrc2ciphersrate,omitempty"`
	Ssltot128bitrc2ciphers uint64 `json:"ssltot128bitrc2ciphers,omitempty"`
	/**
	* Number of RC2 128-bit cipher encryptions on the Citrix ADC.
	*/
	Ssl128bitrc2ciphersrate int32 `json:"ssl128bitrc2ciphersrate,omitempty"`
	Ssltot128bitaesgcmciphers uint64 `json:"ssltot128bitaesgcmciphers,omitempty"`
	/**
	* Number of AEC-GCM 128-bit cipher encryptions on the Citrix ADC.
	*/
	Ssl128bitaesgcmciphersrate int32 `json:"ssl128bitaesgcmciphersrate,omitempty"`
	Ssltot256bitaesgcmciphers uint64 `json:"ssltot256bitaesgcmciphers,omitempty"`
	/**
	* Number of AEC-GCM 256-bit cipher encryptions on the Citrix ADC.
	*/
	Ssl256bitaesgcmciphersrate int32 `json:"ssl256bitaesgcmciphersrate,omitempty"`
	Ssltotnullciphers uint64 `json:"ssltotnullciphers,omitempty"`
	/**
	* Number of Null cipher encryptions on the Citrix ADC.
	*/
	Sslnullciphersrate int32 `json:"sslnullciphersrate,omitempty"`
	Ssltotmd5mac uint64 `json:"ssltotmd5mac,omitempty"`
	/**
	* Number of MD5 hashes on the Citrix ADC.
	*/
	Sslmd5macrate int32 `json:"sslmd5macrate,omitempty"`
	Ssltotshamac uint64 `json:"ssltotshamac,omitempty"`
	/**
	* Number of SHA hashes on the Citrix ADC.
	*/
	Sslshamacrate int32 `json:"sslshamacrate,omitempty"`
	Ssltotsha256mac uint64 `json:"ssltotsha256mac,omitempty"`
	/**
	* Number of SHA256 hashes on the Citrix ADC.
	*/
	Sslsha256macrate int32 `json:"sslsha256macrate,omitempty"`
	Ssltotsha384mac uint64 `json:"ssltotsha384mac,omitempty"`
	/**
	* Number of SHA384 hashes on the Citrix ADC.
	*/
	Sslsha384macrate int32 `json:"sslsha384macrate,omitempty"`
	Ssltotsslv2handshakes uint64 `json:"ssltotsslv2handshakes,omitempty"`
	/**
	* Number of handshakes on SSLv2 on the Citrix ADC.
	*/
	Sslsslv2handshakesrate int32 `json:"sslsslv2handshakesrate,omitempty"`
	Ssltotsslv3handshakes uint64 `json:"ssltotsslv3handshakes,omitempty"`
	/**
	* Number of handshakes on SSLv3 on the Citrix ADC.
	*/
	Sslsslv3handshakesrate int32 `json:"sslsslv3handshakesrate,omitempty"`
	Ssltottlsv1handshakes uint64 `json:"ssltottlsv1handshakes,omitempty"`
	/**
	* Number of SSL handshakes on TLSv1 on the Citrix ADC.
	*/
	Ssltlsv1handshakesrate int32 `json:"ssltlsv1handshakesrate,omitempty"`
	Ssltottlsv11handshakes uint64 `json:"ssltottlsv11handshakes,omitempty"`
	/**
	* Number of SSL handshakes on TLSv1.1 on the Citrix ADC.
	*/
	Ssltlsv11handshakesrate int32 `json:"ssltlsv11handshakesrate,omitempty"`
	Ssltottlsv12handshakes uint64 `json:"ssltottlsv12handshakes,omitempty"`
	/**
	* Number of SSL handshakes on TLSv1.2 on the Citrix ADC.
	*/
	Ssltlsv12handshakesrate int32 `json:"ssltlsv12handshakesrate,omitempty"`
	Ssltottlsv13handshakes uint64 `json:"ssltottlsv13handshakes,omitempty"`
	/**
	* Number of SSL handshakes on TLSv1.3 on the Citrix ADC.
	*/
	Ssltlsv13handshakesrate int32 `json:"ssltlsv13handshakesrate,omitempty"`
	Ssltotdtlsv1handshakes uint64 `json:"ssltotdtlsv1handshakes,omitempty"`
	/**
	* Number of SSL handshakes on DTLSv1 on the Citrix ADC.
	*/
	Ssldtlsv1handshakesrate int32 `json:"ssldtlsv1handshakesrate,omitempty"`
	Ssltotdtlsv12handshakes uint64 `json:"ssltotdtlsv12handshakes,omitempty"`
	/**
	* Number of SSL handshakes on DTLSv1.2 on the Citrix ADC.
	*/
	Ssldtlsv12handshakesrate int32 `json:"ssldtlsv12handshakesrate,omitempty"`
	Ssltotsslv2clientauthentications uint64 `json:"ssltotsslv2clientauthentications,omitempty"`
	/**
	* Number of client authentications done on SSLv2.
	*/
	Sslsslv2clientauthenticationsrate int32 `json:"sslsslv2clientauthenticationsrate,omitempty"`
	Ssltotsslv3clientauthentications uint64 `json:"ssltotsslv3clientauthentications,omitempty"`
	/**
	* Number of client authentications done on SSLv3.
	*/
	Sslsslv3clientauthenticationsrate int32 `json:"sslsslv3clientauthenticationsrate,omitempty"`
	Ssltottlsv1clientauthentications uint64 `json:"ssltottlsv1clientauthentications,omitempty"`
	/**
	* Number of client authentications done on TLSv1.
	*/
	Ssltlsv1clientauthenticationsrate int32 `json:"ssltlsv1clientauthenticationsrate,omitempty"`
	Ssltottlsv11clientauthentications uint64 `json:"ssltottlsv11clientauthentications,omitempty"`
	/**
	* Number of client authentications done on TLSv1.1.
	*/
	Ssltlsv11clientauthenticationsrate int32 `json:"ssltlsv11clientauthenticationsrate,omitempty"`
	Ssltottlsv12clientauthentications uint64 `json:"ssltottlsv12clientauthentications,omitempty"`
	/**
	* Number of client authentications done on TLSv1.2.
	*/
	Ssltlsv12clientauthenticationsrate int32 `json:"ssltlsv12clientauthenticationsrate,omitempty"`
	Ssltottlsv13clientauthentications uint64 `json:"ssltottlsv13clientauthentications,omitempty"`
	/**
	* Number of client authentications done on TLSv1.3.
	*/
	Ssltlsv13clientauthenticationsrate int32 `json:"ssltlsv13clientauthenticationsrate,omitempty"`
	Ssltotdtlsv1clientauthentications uint64 `json:"ssltotdtlsv1clientauthentications,omitempty"`
	/**
	* Number of client authentications done on DTLSv1.
	*/
	Ssldtlsv1clientauthenticationsrate int32 `json:"ssldtlsv1clientauthenticationsrate,omitempty"`
	Ssltotdtlsv12clientauthentications uint64 `json:"ssltotdtlsv12clientauthentications,omitempty"`
	/**
	* Number of client authentications done on DTLSv1.2.
	*/
	Ssldtlsv12clientauthenticationsrate int32 `json:"ssldtlsv12clientauthenticationsrate,omitempty"`
	Ssltotrsaauthorizations uint64 `json:"ssltotrsaauthorizations,omitempty"`
	/**
	* Number of RSA authentications on the Citrix ADC.
	*/
	Sslrsaauthorizationsrate int32 `json:"sslrsaauthorizationsrate,omitempty"`
	Ssltotdhauthorizations uint64 `json:"ssltotdhauthorizations,omitempty"`
	/**
	* Number of Diffie-Helman authentications on the Citrix ADC.
	*/
	Ssldhauthorizationsrate int32 `json:"ssldhauthorizationsrate,omitempty"`
	Ssltotdssauthorizations uint64 `json:"ssltotdssauthorizations,omitempty"`
	/**
	* Total number of times DSS authorization is used on the Citrix ADC.
	*/
	Ssldssauthorizationsrate int32 `json:"ssldssauthorizationsrate,omitempty"`
	Ssltotecdsaauthorizations uint64 `json:"ssltotecdsaauthorizations,omitempty"`
	/**
	* Total number of times ECDSA authorization is used on the Citrix ADC.
	*/
	Sslecdsaauthorizationsrate int32 `json:"sslecdsaauthorizationsrate,omitempty"`
	Ssltotnullauthorizations uint64 `json:"ssltotnullauthorizations,omitempty"`
	/**
	* Number of Null authentications on the Citrix ADC.
	*/
	Sslnullauthorizationsrate int32 `json:"sslnullauthorizationsrate,omitempty"`
	Ssltotbkendsessionrenegotiate uint64 `json:"ssltotbkendsessionrenegotiate,omitempty"`
	/**
	* Number of back-end SSL session renegotiations on the Citrix ADC.
	*/
	Sslbkendsessionrenegotiaterate int32 `json:"sslbkendsessionrenegotiaterate,omitempty"`
	Ssltotbkendsslv3renego uint64 `json:"ssltotbkendsslv3renego,omitempty"`
	/**
	* Number of back-end SSLv3 session renegotiations on the Citrix ADC.
	*/
	Sslbkendsslv3renegorate int32 `json:"sslbkendsslv3renegorate,omitempty"`
	Ssltotbkendtlsvlrenego uint64 `json:"ssltotbkendtlsvlrenego,omitempty"`
	/**
	* Number of back-end TLSv1 session renegotiations on the Citrix ADC.
	*/
	Sslbkendtlsvlrenegorate int32 `json:"sslbkendtlsvlrenegorate,omitempty"`
	Ssltotbkendtlsv11renego uint64 `json:"ssltotbkendtlsv11renego,omitempty"`
	/**
	* Number of back-end TLSv1.1 session renegotiations on the Citrix ADC.
	*/
	Sslbkendtlsv11renegorate int32 `json:"sslbkendtlsv11renegorate,omitempty"`
	Ssltotbkendtlsv12renego uint64 `json:"ssltotbkendtlsv12renego,omitempty"`
	/**
	* Number of back-end TLSv1.2 session renegotiations on the Citrix ADC.
	*/
	Sslbkendtlsv12renegorate int32 `json:"sslbkendtlsv12renegorate,omitempty"`
	Ssltotbkenddtlsvlrenego uint64 `json:"ssltotbkenddtlsvlrenego,omitempty"`
	/**
	* Number of back-end DTLSv1 session renegotiations on the Citrix ADC.
	*/
	Sslbkenddtlsvlrenegorate int32 `json:"sslbkenddtlsvlrenegorate,omitempty"`
	Sslbetotrsa512keyexchanges uint64 `json:"sslbetotrsa512keyexchanges,omitempty"`
	/**
	* Number of back-end RSA 512-bit key exchanges on the Citrix ADC.
	*/
	Sslbersa512keyexchangesrate int32 `json:"sslbersa512keyexchangesrate,omitempty"`
	Sslbetotrsa1024keyexchanges uint64 `json:"sslbetotrsa1024keyexchanges,omitempty"`
	/**
	* Number of back-end RSA 1024-bit key exchanges on the Citrix ADC.
	*/
	Sslbersa1024keyexchangesrate int32 `json:"sslbersa1024keyexchangesrate,omitempty"`
	Sslbetotrsa2048keyexchanges uint64 `json:"sslbetotrsa2048keyexchanges,omitempty"`
	/**
	* Number of back-end RSA 2048-bit key exchanges on the Citrix ADC.
	*/
	Sslbersa2048keyexchangesrate int32 `json:"sslbersa2048keyexchangesrate,omitempty"`
	Sslbetotrsa3072keyexchanges uint64 `json:"sslbetotrsa3072keyexchanges,omitempty"`
	/**
	* Number of back-end RSA 3072-bit key exchanges on the Citrix ADC.
	*/
	Sslbersa3072keyexchangesrate int32 `json:"sslbersa3072keyexchangesrate,omitempty"`
	Sslbetotrsa4096keyexchanges uint64 `json:"sslbetotrsa4096keyexchanges,omitempty"`
	/**
	* Number of back-end RSA 4096-bit key exchanges on the Citrix ADC.
	*/
	Sslbersa4096keyexchangesrate int32 `json:"sslbersa4096keyexchangesrate,omitempty"`
	Sslbetotdh512keyexchanges uint64 `json:"sslbetotdh512keyexchanges,omitempty"`
	/**
	* Number of back-end DH 512-bit key exchanges on the Citrix ADC.
	*/
	Sslbedh512keyexchangesrate int32 `json:"sslbedh512keyexchangesrate,omitempty"`
	Sslbetotdh1024keyexchanges uint64 `json:"sslbetotdh1024keyexchanges,omitempty"`
	/**
	* Number of back-end DH 1024-bit key exchanges on the Citrix ADC.
	*/
	Sslbedh1024keyexchangesrate int32 `json:"sslbedh1024keyexchangesrate,omitempty"`
	Sslbetotdh2048keyexchanges uint64 `json:"sslbetotdh2048keyexchanges,omitempty"`
	/**
	* Number of back-end DH 2048-bit key exchanges on the Citrix ADC.
	*/
	Sslbedh2048keyexchangesrate int32 `json:"sslbedh2048keyexchangesrate,omitempty"`
	Sslbetotdh4096keyexchanges uint64 `json:"sslbetotdh4096keyexchanges,omitempty"`
	/**
	* Number of back-end DH 4096-bit key exchanges on the Citrix ADC.
	*/
	Sslbedh4096keyexchangesrate int32 `json:"sslbedh4096keyexchangesrate,omitempty"`
	Sslbetotecdhecurve521 uint64 `json:"sslbetotecdhecurve521,omitempty"`
	/**
	* Number of back-end ECDHE 521 curve Key exchanges  on the Citrix ADC.
	*/
	Sslbeecdhecurve521rate int32 `json:"sslbeecdhecurve521rate,omitempty"`
	Sslbetotecdhecurve384 uint64 `json:"sslbetotecdhecurve384,omitempty"`
	/**
	* Number of back-end ECDHE 384 curve Key exchanges  on the Citrix ADC.
	*/
	Sslbeecdhecurve384rate int32 `json:"sslbeecdhecurve384rate,omitempty"`
	Sslbetotecdhecurve256 uint64 `json:"sslbetotecdhecurve256,omitempty"`
	/**
	* Number of back-end ECDHE 256 curve Key exchanges  on the Citrix ADC.
	*/
	Sslbeecdhecurve256rate int32 `json:"sslbeecdhecurve256rate,omitempty"`
	Sslbetotecdhecurve224 uint64 `json:"sslbetotecdhecurve224,omitempty"`
	/**
	* Number of back-end ECDHE 224 curve Key exchanges  on the Citrix ADC.
	*/
	Sslbeecdhecurve224rate int32 `json:"sslbeecdhecurve224rate,omitempty"`
	Sslbetot40bitrc4ciphers uint64 `json:"sslbetot40bitrc4ciphers,omitempty"`
	/**
	* Number of back-end RC4 40-bit cipher encryptions on the Citrix ADC.
	*/
	Sslbe40bitrc4ciphersrate int32 `json:"sslbe40bitrc4ciphersrate,omitempty"`
	Sslbetot56bitrc4ciphers uint64 `json:"sslbetot56bitrc4ciphers,omitempty"`
	/**
	* Number of back-end RC4 56-bit cipher encryptions on the Citrix ADC.
	*/
	Sslbe56bitrc4ciphersrate int32 `json:"sslbe56bitrc4ciphersrate,omitempty"`
	Sslbetot64bitrc4ciphers uint64 `json:"sslbetot64bitrc4ciphers,omitempty"`
	/**
	* Number of back-end RC4 64-bit cipher encryptions on the Citrix ADC.
	*/
	Sslbe64bitrc4ciphersrate int32 `json:"sslbe64bitrc4ciphersrate,omitempty"`
	Sslbetot128bitrc4ciphers uint64 `json:"sslbetot128bitrc4ciphers,omitempty"`
	/**
	* Number of back-end RC4 128-bit cipher encryptions on the Citrix ADC.
	*/
	Sslbe128bitrc4ciphersrate int32 `json:"sslbe128bitrc4ciphersrate,omitempty"`
	Sslbetot40bitdesciphers uint64 `json:"sslbetot40bitdesciphers,omitempty"`
	/**
	* Number of back-end DES 40-bit cipher encryptions on the Citrix ADC.
	*/
	Sslbe40bitdesciphersrate int32 `json:"sslbe40bitdesciphersrate,omitempty"`
	Sslbetot56bitdesciphers uint64 `json:"sslbetot56bitdesciphers,omitempty"`
	/**
	* Number of back-end DES 56-bit cipher encryptions on the Citrix ADC.
	*/
	Sslbe56bitdesciphersrate int32 `json:"sslbe56bitdesciphersrate,omitempty"`
	Sslbetot168bit3desciphers uint64 `json:"sslbetot168bit3desciphers,omitempty"`
	/**
	* Number of back-end 3DES 168-bit cipher encryptions on the Citrix ADC.
	*/
	Sslbe168bit3desciphersrate int32 `json:"sslbe168bit3desciphersrate,omitempty"`
	Ssltotbkendcipheraes128 uint64 `json:"ssltotbkendcipheraes128,omitempty"`
	/**
	* Back-end AES 128-bit cipher encryptions on the Citrix ADC.
	*/
	Sslbkendcipheraes128rate int32 `json:"sslbkendcipheraes128rate,omitempty"`
	Ssltotbkendcipheraes256 uint64 `json:"ssltotbkendcipheraes256,omitempty"`
	/**
	* Back-end AES 256-bit cipher encryptions on the Citrix ADC.
	*/
	Sslbkendcipheraes256rate int32 `json:"sslbkendcipheraes256rate,omitempty"`
	Sslbetot40bitrc2ciphers uint64 `json:"sslbetot40bitrc2ciphers,omitempty"`
	/**
	* Number of back-end RC2 40-bit cipher encryptions on the Citrix ADC.
	*/
	Sslbe40bitrc2ciphersrate int32 `json:"sslbe40bitrc2ciphersrate,omitempty"`
	Sslbetot56bitrc2ciphers uint64 `json:"sslbetot56bitrc2ciphers,omitempty"`
	/**
	* Number of back-end RC2 56-bit cipher encryptions on the Citrix ADC.
	*/
	Sslbe56bitrc2ciphersrate int32 `json:"sslbe56bitrc2ciphersrate,omitempty"`
	Sslbetot128bitrc2ciphers uint64 `json:"sslbetot128bitrc2ciphers,omitempty"`
	/**
	* Number of back-end RC2 128-bit cipher encryptions on the Citrix ADC.
	*/
	Sslbe128bitrc2ciphersrate int32 `json:"sslbe128bitrc2ciphersrate,omitempty"`
	Ssltotbkendcipheraesgcm128 uint64 `json:"ssltotbkendcipheraesgcm128,omitempty"`
	/**
	* Back-end AES-GCM 128-bit cipher encryptions on the Citrix ADC.
	*/
	Sslbkendcipheraesgcm128rate int32 `json:"sslbkendcipheraesgcm128rate,omitempty"`
	Ssltotbkendcipheraesgcm256 uint64 `json:"ssltotbkendcipheraesgcm256,omitempty"`
	/**
	* Back-end AES-GCM 256-bit cipher encryptions on the Citrix ADC .
	*/
	Sslbkendcipheraesgcm256rate int32 `json:"sslbkendcipheraesgcm256rate,omitempty"`
	Sslbetotnullciphers uint64 `json:"sslbetotnullciphers,omitempty"`
	/**
	* Number of back-end null cipher encryptions on the Citrix ADC.
	*/
	Sslbenullciphersrate int32 `json:"sslbenullciphersrate,omitempty"`
	Sslbetotmd5mac uint64 `json:"sslbetotmd5mac,omitempty"`
	/**
	* Number of back-end MD5 hashes on the Citrix ADC.
	*/
	Sslbemd5macrate int32 `json:"sslbemd5macrate,omitempty"`
	Sslbetotshamac uint64 `json:"sslbetotshamac,omitempty"`
	/**
	* Number of back-end SHA hashes on the Citrix ADC.
	*/
	Sslbeshamacrate int32 `json:"sslbeshamacrate,omitempty"`
	Sslbetotsha256mac uint64 `json:"sslbetotsha256mac,omitempty"`
	/**
	* Number of back-end SHA256 hashes on the Citrix ADC.
	*/
	Sslbesha256macrate int32 `json:"sslbesha256macrate,omitempty"`
	Sslbetotsha384mac uint64 `json:"sslbetotsha384mac,omitempty"`
	/**
	* Number of back-end SHA384 hashes on the Citrix ADC.
	*/
	Sslbesha384macrate int32 `json:"sslbesha384macrate,omitempty"`
	Sslbetotsslv3handshakes uint64 `json:"sslbetotsslv3handshakes,omitempty"`
	/**
	* Number of back-end SSLv3 handshakes on the Citrix ADC.
	*/
	Sslbesslv3handshakesrate int32 `json:"sslbesslv3handshakesrate,omitempty"`
	Sslbetottlsv1handshakes uint64 `json:"sslbetottlsv1handshakes,omitempty"`
	/**
	* Number of back-end TLSv1 handshakes on the Citrix ADC.
	*/
	Sslbetlsv1handshakesrate int32 `json:"sslbetlsv1handshakesrate,omitempty"`
	Sslbetottlsv11handshakes uint64 `json:"sslbetottlsv11handshakes,omitempty"`
	/**
	* Number of back-end TLSv1.1 handshakes on the Citrix ADC.
	*/
	Sslbetlsv11handshakesrate int32 `json:"sslbetlsv11handshakesrate,omitempty"`
	Sslbetottlsv12handshakes uint64 `json:"sslbetottlsv12handshakes,omitempty"`
	/**
	* Number of back-end TLSv1.2 handshakes on the Citrix ADC.
	*/
	Sslbetlsv12handshakesrate int32 `json:"sslbetlsv12handshakesrate,omitempty"`
	Sslbetotdtlsv1handshakes uint64 `json:"sslbetotdtlsv1handshakes,omitempty"`
	/**
	* Number of back-end DTLSv1 handshakes on the Citrix ADC.
	*/
	Sslbedtlsv1handshakesrate int32 `json:"sslbedtlsv1handshakesrate,omitempty"`
	Sslbetotsslv3clientauthentications uint64 `json:"sslbetotsslv3clientauthentications,omitempty"`
	/**
	* Number of back-end SSLv3 client authentications on the Citrix ADC.
	*/
	Sslbesslv3clientauthenticationsrate int32 `json:"sslbesslv3clientauthenticationsrate,omitempty"`
	Sslbetottlsv1clientauthentications uint64 `json:"sslbetottlsv1clientauthentications,omitempty"`
	/**
	* Number of back-end TLSv1 client authentications on the Citrix ADC.
	*/
	Sslbetlsv1clientauthenticationsrate int32 `json:"sslbetlsv1clientauthenticationsrate,omitempty"`
	Sslbetottlsv11clientauthentications uint64 `json:"sslbetottlsv11clientauthentications,omitempty"`
	/**
	* Number of back-end TLSv1.1 client authentications on the Citrix ADC.
	*/
	Sslbetlsv11clientauthenticationsrate int32 `json:"sslbetlsv11clientauthenticationsrate,omitempty"`
	Sslbetottlsv12clientauthentications uint64 `json:"sslbetottlsv12clientauthentications,omitempty"`
	/**
	* Number of back-end TLSv1.2 client authentications on the Citrix ADC.
	*/
	Sslbetlsv12clientauthenticationsrate int32 `json:"sslbetlsv12clientauthenticationsrate,omitempty"`
	Sslbetotdtlsv1clientauthentications uint64 `json:"sslbetotdtlsv1clientauthentications,omitempty"`
	/**
	* Number of back-end DTLSv1 client authentications on the Citrix ADC.
	*/
	Sslbedtlsv1clientauthenticationsrate int32 `json:"sslbedtlsv1clientauthenticationsrate,omitempty"`
	Sslbetotrsaauthorizations uint64 `json:"sslbetotrsaauthorizations,omitempty"`
	/**
	* Number of back-end RSA authentications on the Citrix ADC.
	*/
	Sslbersaauthorizationsrate int32 `json:"sslbersaauthorizationsrate,omitempty"`
	Sslbetotdhauthorizations uint64 `json:"sslbetotdhauthorizations,omitempty"`
	/**
	* Number of back-end DH authentications on the Citrix ADC.
	*/
	Sslbedhauthorizationsrate int32 `json:"sslbedhauthorizationsrate,omitempty"`
	Sslbetotdssauthorizations uint64 `json:"sslbetotdssauthorizations,omitempty"`
	/**
	* Number of back-end DSS authentications on the Citrix ADC.
	*/
	Sslbedssauthorizationsrate int32 `json:"sslbedssauthorizationsrate,omitempty"`
	Sslbetotecdsaauthorizations uint64 `json:"sslbetotecdsaauthorizations,omitempty"`
	/**
	* Number of back-end ECDSA authentications on the Citrix ADC.
	*/
	Sslbeecdsaauthorizationsrate int32 `json:"sslbeecdsaauthorizationsrate,omitempty"`
	Sslbetotnullauthorizations uint64 `json:"sslbetotnullauthorizations,omitempty"`
	/**
	* Number of back-end null authentications on the Citrix ADC.
	*/
	Sslbenullauthorizationsrate int32 `json:"sslbenullauthorizationsrate,omitempty"`
	Ssltotoffloadrsakeyexchanges uint64 `json:"ssltotoffloadrsakeyexchanges,omitempty"`
	/**
	* Number of RSA key exchanges offloaded to the cryptography card.
	*/
	Ssloffloadrsakeyexchangesrate int32 `json:"ssloffloadrsakeyexchangesrate,omitempty"`
	Ssltotoffloadsignrsa uint64 `json:"ssltotoffloadsignrsa,omitempty"`
	/**
	* Number of RSA sign operations offloaded to the cryptography card.
	*/
	Ssloffloadsignrsarate int32 `json:"ssloffloadsignrsarate,omitempty"`
	Ssltotoffloaddhkeyexchanges uint64 `json:"ssltotoffloaddhkeyexchanges,omitempty"`
	/**
	* Number of DH key exchanges offloaded to the cryptography card.
	*/
	Ssloffloaddhkeyexchangesrate int32 `json:"ssloffloaddhkeyexchangesrate,omitempty"`
	Ssltotoffloadbulkrc4 uint64 `json:"ssltotoffloadbulkrc4,omitempty"`
	/**
	* Number of RC4 encryptions offloaded to the cryptography card.
	*/
	Ssloffloadbulkrc4rate int32 `json:"ssloffloadbulkrc4rate,omitempty"`
	Ssltotoffloadbulkdes uint64 `json:"ssltotoffloadbulkdes,omitempty"`
	/**
	* Number of DES encryptions offloaded to the cryptography card.
	*/
	Ssloffloadbulkdesrate int32 `json:"ssloffloadbulkdesrate,omitempty"`
	Ssltotoffloadbulkaes uint64 `json:"ssltotoffloadbulkaes,omitempty"`
	/**
	* Number of AES encryptions offloaded to the cryptography card.
	*/
	Ssloffloadbulkaesrate int32 `json:"ssloffloadbulkaesrate,omitempty"`
	Ssltotoffloadbulkaesgcm128 uint64 `json:"ssltotoffloadbulkaesgcm128,omitempty"`
	/**
	* Number of AES-GCM 128-bit encryptions offloaded to the cryptography card.
	*/
	Ssloffloadbulkaesgcm128rate int32 `json:"ssloffloadbulkaesgcm128rate,omitempty"`
	Ssltotoffloadbulkaesgcm256 uint64 `json:"ssltotoffloadbulkaesgcm256,omitempty"`
	/**
	* Number of AES-GCM 256-bit encryptions offloaded to the cryptography card.
	*/
	Ssloffloadbulkaesgcm256rate int32 `json:"ssloffloadbulkaesgcm256rate,omitempty"`
	Ssltotenchw uint64 `json:"ssltotenchw,omitempty"`
	/**
	* Number of bytes encrypted in hardware.
	*/
	Sslenchwrate int32 `json:"sslenchwrate,omitempty"`
	Ssltotencsw uint64 `json:"ssltotencsw,omitempty"`
	/**
	* Number of bytes encrypted in software.
	*/
	Sslencswrate int32 `json:"sslencswrate,omitempty"`
	Ssltotencfe uint64 `json:"ssltotencfe,omitempty"`
	/**
	* Number of bytes encrypted on the front-end.
	*/
	Sslencferate int32 `json:"sslencferate,omitempty"`
	Ssltothwencfe uint64 `json:"ssltothwencfe,omitempty"`
	/**
	* Number of bytes encrypted in hardware on the front-end.
	*/
	Sslhwencferate int32 `json:"sslhwencferate,omitempty"`
	Ssltotswencfe uint64 `json:"ssltotswencfe,omitempty"`
	/**
	* Number of bytes encrypted in software on the front-end.
	*/
	Sslswencferate int32 `json:"sslswencferate,omitempty"`
	Ssltotencbe uint64 `json:"ssltotencbe,omitempty"`
	/**
	* Number of bytes encrypted on the back-end.
	*/
	Sslencberate int32 `json:"sslencberate,omitempty"`
	Ssltothwencbe uint64 `json:"ssltothwencbe,omitempty"`
	/**
	* Number of bytes encrypted in hardware on the back-end.
	*/
	Sslhwencberate int32 `json:"sslhwencberate,omitempty"`
	Ssltotswencbe uint64 `json:"ssltotswencbe,omitempty"`
	/**
	* Number of bytes encrypted in software on the back-end.
	*/
	Sslswencberate int32 `json:"sslswencberate,omitempty"`
	Ssltotdechw uint64 `json:"ssltotdechw,omitempty"`
	/**
	* Number of bytes decrypted in hardware.
	*/
	Ssldechwrate int32 `json:"ssldechwrate,omitempty"`
	Ssltotdecsw uint64 `json:"ssltotdecsw,omitempty"`
	/**
	* Number of bytes decrypted in software.
	*/
	Ssldecswrate int32 `json:"ssldecswrate,omitempty"`
	Ssltotdecfe uint64 `json:"ssltotdecfe,omitempty"`
	/**
	* Number of bytes decrypted on the front-end.
	*/
	Ssldecferate int32 `json:"ssldecferate,omitempty"`
	Ssltothwdecfe uint64 `json:"ssltothwdecfe,omitempty"`
	/**
	* Number of bytes decrypted in hardware on the front-end.
	*/
	Sslhwdecferate int32 `json:"sslhwdecferate,omitempty"`
	Ssltotswdecfe uint64 `json:"ssltotswdecfe,omitempty"`
	/**
	* Number of bytes decrypted in software on the front-end.
	*/
	Sslswdecferate int32 `json:"sslswdecferate,omitempty"`
	Ssltotdecbe uint64 `json:"ssltotdecbe,omitempty"`
	/**
	* Number of bytes decrypted on the back-end.
	*/
	Ssldecberate int32 `json:"ssldecberate,omitempty"`
	Ssltothwdecbe uint64 `json:"ssltothwdecbe,omitempty"`
	/**
	* Number of bytes decrypted in hardware on the back-end.
	*/
	Sslhwdecberate int32 `json:"sslhwdecberate,omitempty"`
	Ssltotswdecbe uint64 `json:"ssltotswdecbe,omitempty"`
	/**
	* Number of bytes decrypted in software on the back-end
	*/
	Sslswdecberate int32 `json:"sslswdecberate,omitempty"`
	Sslcursslinfospcbinusecount uint32 `json:"sslcursslinfospcbinusecount,omitempty"`
	/**
	* Number of SPCB in use.
	*/
	Sslcursslinfospcbinusecountrate int32 `json:"sslcursslinfospcbinusecountrate,omitempty"`
	Sslcursessions uint32 `json:"sslcursessions,omitempty"`
	/**
	* Number of active SSL sessions on the Citrix ADC.
	*/
	Sslcursessionsrate int32 `json:"sslcursessionsrate,omitempty"`
	Sslcurqsize uint32 `json:"sslcurqsize,omitempty"`
	/**
	* Current queue size
	*/
	Sslcurqsizerate int32 `json:"sslcurqsizerate,omitempty"`
	Sslcursslinfonscardinqcount uint32 `json:"sslcursslinfonscardinqcount,omitempty"`
	/**
	* Number of current SSL card InQ count.
	*/
	Sslcursslinfonscardinqcountrate int32 `json:"sslcursslinfonscardinqcountrate,omitempty"`
	Sslcursslinfocardinblkq uint32 `json:"sslcursslinfocardinblkq,omitempty"`
	/**
	* Number of current SSL card In BulkQ count.
	*/
	Sslcursslinfocardinblkqrate int32 `json:"sslcursslinfocardinblkqrate,omitempty"`
	Sslcursslinfocardinkeyq uint32 `json:"sslcursslinfocardinkeyq,omitempty"`
	/**
	* Number of current SSL card In KeyQ count.
	*/
	Sslcursslinfocardinkeyqrate int32 `json:"sslcursslinfocardinkeyqrate,omitempty"`
	Sslbemaxmultiplexedsessions uint64 `json:"sslbemaxmultiplexedsessions,omitempty"`
	/**
	* Number of back-end SSL sessions reused on the Citrix ADC.
	*/
	Sslbemaxmultiplexedsessionsrate int32 `json:"sslbemaxmultiplexedsessionsrate,omitempty"`
	Ssltot128bitideaciphers uint64 `json:"ssltot128bitideaciphers,omitempty"`
	/**
	* Number of IDEA 128-bit cipher encryptions on the Citrix ADC.
	*/
	Ssl128bitideaciphersrate int32 `json:"ssl128bitideaciphersrate,omitempty"`
	Sslbetot128bitideaciphers uint64 `json:"sslbetot128bitideaciphers,omitempty"`
	/**
	* Number of back-end IDEA 128-bit cipher encryptions on the Citrix ADC.
	*/
	Sslbe128bitideaciphersrate int32 `json:"sslbe128bitideaciphersrate,omitempty"`

}
