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
	Ssltothwdecbesecondary int `json:"ssltothwdecbesecondary,omitempty"`
	/**
	* Number of bytes decrypted on the back-end in hardware on secondary card.
	*/
	Sslhwdecbesecondaryrate float64 `json:"sslhwdecbesecondaryrate,omitempty"`
	Ssltothwdecfesecondary int `json:"ssltothwdecfesecondary,omitempty"`
	/**
	* Number of bytes decrypted on the front-end in hardware on secondary card.
	*/
	Sslhwdecfesecondaryrate float64 `json:"sslhwdecfesecondaryrate,omitempty"`
	Ssltothwencbesecondary int `json:"ssltothwencbesecondary,omitempty"`
	/**
	* Number of bytes encrypted on the back-end in hardware on secondary card.
	*/
	Sslhwencbesecondaryrate float64 `json:"sslhwencbesecondaryrate,omitempty"`
	Ssltothwencfesecondary int `json:"ssltothwencfesecondary,omitempty"`
	/**
	* Number of bytes encrypted on the front-end in hardware on secondary card.
	*/
	Sslhwencfesecondaryrate float64 `json:"sslhwencfesecondaryrate,omitempty"`
	Sslecdsap521cryptoutilizationstat float64 `json:"sslecdsap521cryptoutilizationstat,omitempty"`
	Sslecdsap384cryptoutilizationstat float64 `json:"sslecdsap384cryptoutilizationstat,omitempty"`
	Sslecdsap256cryptoutilizationstat float64 `json:"sslecdsap256cryptoutilizationstat,omitempty"`
	Sslecdsap224cryptoutilizationstat float64 `json:"sslecdsap224cryptoutilizationstat,omitempty"`
	Sslasymecdsacryptoutilizationstat float64 `json:"sslasymecdsacryptoutilizationstat,omitempty"`
	Sslecdhx25519cryptoutilizationstat float64 `json:"sslecdhx25519cryptoutilizationstat,omitempty"`
	Sslecdhp521cryptoutilizationstat float64 `json:"sslecdhp521cryptoutilizationstat,omitempty"`
	Sslecdhp384cryptoutilizationstat float64 `json:"sslecdhp384cryptoutilizationstat,omitempty"`
	Sslecdhp256cryptoutilizationstat float64 `json:"sslecdhp256cryptoutilizationstat,omitempty"`
	Sslecdhp224cryptoutilizationstat float64 `json:"sslecdhp224cryptoutilizationstat,omitempty"`
	Sslasymecdhcryptoutilizationstat float64 `json:"sslasymecdhcryptoutilizationstat,omitempty"`
	Sslasymdhcryptoutilizationstat float64 `json:"sslasymdhcryptoutilizationstat,omitempty"`
	Sslasymrsaotherscryptoutilizationstat float64 `json:"sslasymrsaotherscryptoutilizationstat,omitempty"`
	Sslasymrsa1kcryptoutilizationstat float64 `json:"sslasymrsa1kcryptoutilizationstat,omitempty"`
	Sslasymrsa2kcryptoutilizationstat float64 `json:"sslasymrsa2kcryptoutilizationstat,omitempty"`
	Sslasymrsa4kcryptoutilizationstat float64 `json:"sslasymrsa4kcryptoutilizationstat,omitempty"`
	Sslasymrsacryptoutilizationstat float64 `json:"sslasymrsacryptoutilizationstat,omitempty"`
	Ssltotdechwsecondary int `json:"ssltotdechwsecondary,omitempty"`
	/**
	* Number of bytes decrypted in hardware on secondary card.
	*/
	Ssldechwsecondaryrate float64 `json:"ssldechwsecondaryrate,omitempty"`
	Ssltotenchwsecondary int `json:"ssltotenchwsecondary,omitempty"`
	/**
	* Number of bytes encrypted in hardware on secondary card.
	*/
	Sslenchwsecondaryrate float64 `json:"sslenchwsecondaryrate,omitempty"`
	Sslsymcryptoutilizationstat float64 `json:"sslsymcryptoutilizationstat,omitempty"`
	Sslasymcryptoutilizationstat float64 `json:"sslasymcryptoutilizationstat,omitempty"`
	Sslcryptoutilizationsymmstat int `json:"sslcryptoutilizationsymmstat,omitempty"`
	Sslcryptoutilizationasymstat int `json:"sslcryptoutilizationasymstat,omitempty"`
	Sslcryptoutilizationstat2nd float64 `json:"sslcryptoutilizationstat2nd,omitempty"`
	Sslcryptoutilizationstat float64 `json:"sslcryptoutilizationstat,omitempty"`
	Sslnumcardsupsecondary int `json:"sslnumcardsupsecondary,omitempty"`
	Sslcardssecondary int `json:"sslcardssecondary,omitempty"`
	Sslcardstatus int `json:"sslcardstatus,omitempty"`
	Sslcards int `json:"sslcards,omitempty"`
	Sslnumcardsup int `json:"sslnumcardsup,omitempty"`
	Sslenginestatus int `json:"sslenginestatus,omitempty"`
	Ssltotsessions int `json:"ssltotsessions,omitempty"`
	/**
	* Number of SSL sessions on the Citrix ADC.
	*/
	Sslsessionsrate float64 `json:"sslsessionsrate,omitempty"`
	Ssltottransactions int `json:"ssltottransactions,omitempty"`
	/**
	* Number of SSL transactions on the Citrix ADC
	*/
	Ssltransactionsrate float64 `json:"ssltransactionsrate,omitempty"`
	Ssltotsslv3transactions int `json:"ssltotsslv3transactions,omitempty"`
	/**
	* Total number of SSLv3 transactions on the Citrix ADC.
	*/
	Sslsslv3transactionsrate float64 `json:"sslsslv3transactionsrate,omitempty"`
	Ssltottlsv1transactions int `json:"ssltottlsv1transactions,omitempty"`
	/**
	* Number of TLSv1 transactions on the Citrix ADC.
	*/
	Ssltlsv1transactionsrate float64 `json:"ssltlsv1transactionsrate,omitempty"`
	Ssltottlsv11transactions int `json:"ssltottlsv11transactions,omitempty"`
	/**
	* Number of TLSv1.1 transactions on the Citrix ADC.
	*/
	Ssltlsv11transactionsrate float64 `json:"ssltlsv11transactionsrate,omitempty"`
	Ssltottlsv12transactions int `json:"ssltottlsv12transactions,omitempty"`
	/**
	* Number of TLSv1.2 transactions on the Citrix ADC.
	*/
	Ssltlsv12transactionsrate float64 `json:"ssltlsv12transactionsrate,omitempty"`
	Ssltottlsv13transactions int `json:"ssltottlsv13transactions,omitempty"`
	/**
	* Number of TLSv1.3 transactions on the Citrix ADC.
	*/
	Ssltlsv13transactionsrate float64 `json:"ssltlsv13transactionsrate,omitempty"`
	Ssltotdtlsv1transactions int `json:"ssltotdtlsv1transactions,omitempty"`
	/**
	* Number of DTLSv1 transactions on the Citrix ADC.
	*/
	Ssldtlsv1transactionsrate float64 `json:"ssldtlsv1transactionsrate,omitempty"`
	Ssltotdtlsv12transactions int `json:"ssltotdtlsv12transactions,omitempty"`
	/**
	* Number of DTLSv1.2 transactions on the Citrix ADC.
	*/
	Ssldtlsv12transactionsrate float64 `json:"ssldtlsv12transactionsrate,omitempty"`
	Ssltotsslv3sessions int `json:"ssltotsslv3sessions,omitempty"`
	/**
	* Number of SSLv3 sessions on the Citrix ADC.
	*/
	Sslsslv3sessionsrate float64 `json:"sslsslv3sessionsrate,omitempty"`
	Ssltottlsv1sessions int `json:"ssltottlsv1sessions,omitempty"`
	/**
	* Number of TLSv1 sessions on the Citrix ADC.
	*/
	Ssltlsv1sessionsrate float64 `json:"ssltlsv1sessionsrate,omitempty"`
	Ssltottlsv11sessions int `json:"ssltottlsv11sessions,omitempty"`
	/**
	* Number of TLSv1.1 sessions on the Citrix ADC.
	*/
	Ssltlsv11sessionsrate float64 `json:"ssltlsv11sessionsrate,omitempty"`
	Ssltottlsv12sessions int `json:"ssltottlsv12sessions,omitempty"`
	/**
	* Number of TLSv1.2 sessions on the Citrix ADC.
	*/
	Ssltlsv12sessionsrate float64 `json:"ssltlsv12sessionsrate,omitempty"`
	Ssltottlsv13sessions int `json:"ssltottlsv13sessions,omitempty"`
	/**
	* Number of TLSv1.3 sessions on the Citrix ADC.
	*/
	Ssltlsv13sessionsrate float64 `json:"ssltlsv13sessionsrate,omitempty"`
	Ssltotdtlsv1sessions int `json:"ssltotdtlsv1sessions,omitempty"`
	/**
	* Number of DTLSv1 sessions on the Citrix ADC.
	*/
	Ssldtlsv1sessionsrate float64 `json:"ssldtlsv1sessionsrate,omitempty"`
	Ssltotdtlsv12sessions int `json:"ssltotdtlsv12sessions,omitempty"`
	/**
	* Number of DTLSv1.2 sessions on the Citrix ADC.
	*/
	Ssldtlsv12sessionsrate float64 `json:"ssldtlsv12sessionsrate,omitempty"`
	Ssltotnewsessions int `json:"ssltotnewsessions,omitempty"`
	/**
	* Number of new SSL sessions created on the Citrix ADC.
	*/
	Sslnewsessionsrate float64 `json:"sslnewsessionsrate,omitempty"`
	Ssltotsessionmiss int `json:"ssltotsessionmiss,omitempty"`
	/**
	* Number of SSL session reuse misses on the Citrix ADC.
	*/
	Sslsessionmissrate float64 `json:"sslsessionmissrate,omitempty"`
	Ssltotsessionhits int `json:"ssltotsessionhits,omitempty"`
	/**
	* Number of SSL session reuse hits on the Citrix ADC.
	*/
	Sslsessionhitsrate float64 `json:"sslsessionhitsrate,omitempty"`
	Sslbetotsessions int `json:"sslbetotsessions,omitempty"`
	/**
	* Number of back-end SSL sessions on the Citrix ADC.
	*/
	Sslbesessionsrate float64 `json:"sslbesessionsrate,omitempty"`
	Sslbetotsslv3sessions int `json:"sslbetotsslv3sessions,omitempty"`
	/**
	* Number of back-end SSLv3 sessions on the Citrix ADC.
	*/
	Sslbesslv3sessionsrate float64 `json:"sslbesslv3sessionsrate,omitempty"`
	Sslbetottlsv1sessions int `json:"sslbetottlsv1sessions,omitempty"`
	/**
	* Number of back-end TLSv1 sessions on the Citrix ADC.
	*/
	Sslbetlsv1sessionsrate float64 `json:"sslbetlsv1sessionsrate,omitempty"`
	Sslbetottlsv11sessions int `json:"sslbetottlsv11sessions,omitempty"`
	/**
	* Number of back-end TLSv1.1 sessions on the Citrix ADC.
	*/
	Sslbetlsv11sessionsrate float64 `json:"sslbetlsv11sessionsrate,omitempty"`
	Sslbetottlsv12sessions int `json:"sslbetottlsv12sessions,omitempty"`
	/**
	* Number of back-end TLSv1.2 sessions on the Citrix ADC.
	*/
	Sslbetlsv12sessionsrate float64 `json:"sslbetlsv12sessionsrate,omitempty"`
	Sslbetottlsv13sessions int `json:"sslbetottlsv13sessions,omitempty"`
	/**
	* Number of back-end TLSv1.3 sessions on the Citrix ADC.
	*/
	Sslbetlsv13sessionsrate float64 `json:"sslbetlsv13sessionsrate,omitempty"`
	Sslbetotdtlsv1sessions int `json:"sslbetotdtlsv1sessions,omitempty"`
	/**
	* Number of back-end DTLSv1 sessions on the Citrix ADC.
	*/
	Sslbedtlsv1sessionsrate float64 `json:"sslbedtlsv1sessionsrate,omitempty"`
	Sslbetotdtlsv12sessions int `json:"sslbetotdtlsv12sessions,omitempty"`
	/**
	* Number of back-end DTLSv1.2 sessions on the Citrix ADC.
	*/
	Sslbedtlsv12sessionsrate float64 `json:"sslbedtlsv12sessionsrate,omitempty"`
	Sslbetotsessionmultiplexattempts int `json:"sslbetotsessionmultiplexattempts,omitempty"`
	/**
	* Number of back-end SSL session multiplex attempts on the Citrix ADC.
	*/
	Sslbesessionmultiplexattemptsrate float64 `json:"sslbesessionmultiplexattemptsrate,omitempty"`
	Sslbetotsessionmultiplexattemptsuccess int `json:"sslbetotsessionmultiplexattemptsuccess,omitempty"`
	/**
	* Number of back-end SSL session multiplex successes on the Citrix ADC.
	*/
	Sslbesessionmultiplexattemptsuccessrate float64 `json:"sslbesessionmultiplexattemptsuccessrate,omitempty"`
	Sslbetotsessionmultiplexattemptfails int `json:"sslbetotsessionmultiplexattemptfails,omitempty"`
	/**
	* Number of back-end SSL session multiplex failures on the Citrix ADC.
	*/
	Sslbesessionmultiplexattemptfailsrate float64 `json:"sslbesessionmultiplexattemptfailsrate,omitempty"`
	Ssltotenc int `json:"ssltotenc,omitempty"`
	/**
	* Number of bytes encrypted on the Citrix ADC.
	*/
	Sslencrate float64 `json:"sslencrate,omitempty"`
	Ssltotdec int `json:"ssltotdec,omitempty"`
	/**
	* Number of bytes decrypted on the Citrix ADC.
	*/
	Ssldecrate float64 `json:"ssldecrate,omitempty"`
	Ssltotsslserverinrecords int `json:"ssltotsslserverinrecords,omitempty"`
	/**
	* Number server in record on the Citrix ADC.
	*/
	Sslsslserverinrecordsrate float64 `json:"sslsslserverinrecordsrate,omitempty"`
	Ssltotrenegsessions int `json:"ssltotrenegsessions,omitempty"`
	/**
	* Number of SSL session renegotiations on the Citrix ADC.
	*/
	Sslrenegsessionsrate float64 `json:"sslrenegsessionsrate,omitempty"`
	Ssltotsslv3renegsessions int `json:"ssltotsslv3renegsessions,omitempty"`
	/**
	* Number of session renegotiations done on SSLv3.
	*/
	Sslsslv3renegsessionsrate float64 `json:"sslsslv3renegsessionsrate,omitempty"`
	Ssltottlsv1renegsessions int `json:"ssltottlsv1renegsessions,omitempty"`
	/**
	* Number of SSL session renegotiations done on TLSv1.
	*/
	Ssltlsv1renegsessionsrate float64 `json:"ssltlsv1renegsessionsrate,omitempty"`
	Ssltottlsv11renegsessions int `json:"ssltottlsv11renegsessions,omitempty"`
	/**
	* Number of SSL session renegotiations done on TLSv1.1.
	*/
	Ssltlsv11renegsessionsrate float64 `json:"ssltlsv11renegsessionsrate,omitempty"`
	Ssltottlsv12renegsessions int `json:"ssltottlsv12renegsessions,omitempty"`
	/**
	* Number of SSL session renegotiations done on TLSv1.2.
	*/
	Ssltlsv12renegsessionsrate float64 `json:"ssltlsv12renegsessionsrate,omitempty"`
	Ssltotdtlsv1renegsessions int `json:"ssltotdtlsv1renegsessions,omitempty"`
	/**
	* Number of SSL session renegotiations done on DTLSv1.
	*/
	Ssldtlsv1renegsessionsrate float64 `json:"ssldtlsv1renegsessionsrate,omitempty"`
	Ssltotdtlsv12renegsessions int `json:"ssltotdtlsv12renegsessions,omitempty"`
	/**
	* Number of SSL session renegotiations done on DTLSv1.2.
	*/
	Ssldtlsv12renegsessionsrate float64 `json:"ssldtlsv12renegsessionsrate,omitempty"`
	Ssltotrsa512keyexchanges int `json:"ssltotrsa512keyexchanges,omitempty"`
	/**
	* Number of RSA 512-bit key exchanges on the Citrix ADC.
	*/
	Sslrsa512keyexchangesrate float64 `json:"sslrsa512keyexchangesrate,omitempty"`
	Ssltotrsa1024keyexchanges int `json:"ssltotrsa1024keyexchanges,omitempty"`
	/**
	* Number of RSA 1024-bit key exchanges on the Citrix ADC.
	*/
	Sslrsa1024keyexchangesrate float64 `json:"sslrsa1024keyexchangesrate,omitempty"`
	Ssltotrsa2048keyexchanges int `json:"ssltotrsa2048keyexchanges,omitempty"`
	/**
	* Number of RSA 2048-bit key exchanges on the Citrix ADC.
	*/
	Sslrsa2048keyexchangesrate float64 `json:"sslrsa2048keyexchangesrate,omitempty"`
	Ssltotrsa3072keyexchanges int `json:"ssltotrsa3072keyexchanges,omitempty"`
	/**
	* Number of RSA 3072-bit key exchanges on the Citrix ADC.
	*/
	Sslrsa3072keyexchangesrate float64 `json:"sslrsa3072keyexchangesrate,omitempty"`
	Ssltotrsa4096keyexchanges int `json:"ssltotrsa4096keyexchanges,omitempty"`
	/**
	* Number of RSA 4096-bit key exchanges on the Citrix ADC.
	*/
	Sslrsa4096keyexchangesrate float64 `json:"sslrsa4096keyexchangesrate,omitempty"`
	Ssltotdh512keyexchanges int `json:"ssltotdh512keyexchanges,omitempty"`
	/**
	* Number of Diffie-Helman 512-bit key exchanges on the Citrix ADC.
	*/
	Ssldh512keyexchangesrate float64 `json:"ssldh512keyexchangesrate,omitempty"`
	Ssltotdh1024keyexchanges int `json:"ssltotdh1024keyexchanges,omitempty"`
	/**
	* Number of Diffie-Helman 1024-bit key exchanges on the Citrix ADC.
	*/
	Ssldh1024keyexchangesrate float64 `json:"ssldh1024keyexchangesrate,omitempty"`
	Ssltotdh2048keyexchanges int `json:"ssltotdh2048keyexchanges,omitempty"`
	/**
	* Number of Diffie-Helman 2048-bit key exchanges on the Citrix ADC.
	*/
	Ssldh2048keyexchangesrate float64 `json:"ssldh2048keyexchangesrate,omitempty"`
	Ssltotdh3072keyexchanges int `json:"ssltotdh3072keyexchanges,omitempty"`
	/**
	* Number of Diffie-Helman 3072-bit key exchanges on the Citrix ADC.
	*/
	Ssldh3072keyexchangesrate float64 `json:"ssldh3072keyexchangesrate,omitempty"`
	Ssltotdh4096keyexchanges int `json:"ssltotdh4096keyexchanges,omitempty"`
	/**
	* Number of Diffie-Helman 4096-bit key exchanges on the Citrix ADC.
	*/
	Ssldh4096keyexchangesrate float64 `json:"ssldh4096keyexchangesrate,omitempty"`
	Ssltotpqcx25519mlkem768keyexchanges int `json:"ssltotpqcx25519mlkem768keyexchanges,omitempty"`
	/**
	* Number of Post Quantum Crypto Key Exchanges Using X25519MLKEM768 Group on the Citrix ADC.
	*/
	Sslpqcx25519mlkem768keyexchangesrate float64 `json:"sslpqcx25519mlkem768keyexchangesrate,omitempty"`
	Ssltotecdhe521keyexchanges int `json:"ssltotecdhe521keyexchanges,omitempty"`
	/**
	* Number of 521 Elliptical Curve Diffie-Helman on the Citrix ADC.
	*/
	Sslecdhe521keyexchangesrate float64 `json:"sslecdhe521keyexchangesrate,omitempty"`
	Ssltotecdhe384keyexchanges int `json:"ssltotecdhe384keyexchanges,omitempty"`
	/**
	* Number of 384 Elliptical Curve Diffie-Helman on the Citrix ADC.
	*/
	Sslecdhe384keyexchangesrate float64 `json:"sslecdhe384keyexchangesrate,omitempty"`
	Ssltotecdhe256keyexchanges int `json:"ssltotecdhe256keyexchanges,omitempty"`
	/**
	* Number of 256 Elliptical Curve Diffie-Helman on the Citrix ADC.
	*/
	Sslecdhe256keyexchangesrate float64 `json:"sslecdhe256keyexchangesrate,omitempty"`
	Ssltotecdhe224keyexchanges int `json:"ssltotecdhe224keyexchanges,omitempty"`
	/**
	* Number of 224 Elliptical Curve Diffie-Helman on the Citrix ADC.
	*/
	Sslecdhe224keyexchangesrate float64 `json:"sslecdhe224keyexchangesrate,omitempty"`
	Ssltotecdhe25519keyexchanges int `json:"ssltotecdhe25519keyexchanges,omitempty"`
	/**
	* Number of 25519 Elliptical Curve Diffie-Helman on the Citrix ADC.
	*/
	Sslecdhe25519keyexchangesrate float64 `json:"sslecdhe25519keyexchangesrate,omitempty"`
	Ssltotecdhetransactions int `json:"ssltotecdhetransactions,omitempty"`
	/**
	* Total ECDHE Transactions on Citrix ADC.
	*/
	Sslecdhetransactionsrate float64 `json:"sslecdhetransactionsrate,omitempty"`
	Ssltot40bitrc4ciphers int `json:"ssltot40bitrc4ciphers,omitempty"`
	/**
	* Number of RC4 40-bit cipher encryptions on the Citrix ADC.
	*/
	Ssl40bitrc4ciphersrate float64 `json:"ssl40bitrc4ciphersrate,omitempty"`
	Ssltot56bitrc4ciphers int `json:"ssltot56bitrc4ciphers,omitempty"`
	/**
	* Number of RC4 56-bit cipher encryptions on the Citrix ADC.
	*/
	Ssl56bitrc4ciphersrate float64 `json:"ssl56bitrc4ciphersrate,omitempty"`
	Ssltot64bitrc4ciphers int `json:"ssltot64bitrc4ciphers,omitempty"`
	/**
	* Number of RC4 64-bit cipher encryptions on the Citrix ADC.
	*/
	Ssl64bitrc4ciphersrate float64 `json:"ssl64bitrc4ciphersrate,omitempty"`
	Ssltot128bitrc4ciphers int `json:"ssltot128bitrc4ciphers,omitempty"`
	/**
	* Number of RC4 128-bit cipher encryptions on the Citrix ADC.
	*/
	Ssl128bitrc4ciphersrate float64 `json:"ssl128bitrc4ciphersrate,omitempty"`
	Ssltot40bitdesciphers int `json:"ssltot40bitdesciphers,omitempty"`
	/**
	* Number of DES 40-bit cipher encryptions on the Citrix ADC.
	*/
	Ssl40bitdesciphersrate float64 `json:"ssl40bitdesciphersrate,omitempty"`
	Ssltot56bitdesciphers int `json:"ssltot56bitdesciphers,omitempty"`
	/**
	* Number of DES 56-bit cipher encryptions on the Citrix ADC.
	*/
	Ssl56bitdesciphersrate float64 `json:"ssl56bitdesciphersrate,omitempty"`
	Ssltot168bit3desciphers int `json:"ssltot168bit3desciphers,omitempty"`
	/**
	* Number of DES 168-bit cipher encryptions on the Citrix ADC.
	*/
	Ssl168bit3desciphersrate float64 `json:"ssl168bit3desciphersrate,omitempty"`
	Ssltotcipheraes128 int `json:"ssltotcipheraes128,omitempty"`
	/**
	* Number of AES 128-bit cipher encryptions on the Citrix ADC.
	*/
	Sslcipheraes128rate float64 `json:"sslcipheraes128rate,omitempty"`
	Ssltotcipheraes256 int `json:"ssltotcipheraes256,omitempty"`
	/**
	* Number of AES 256-bit cipher encryptions on the Citrix ADC.
	*/
	Sslcipheraes256rate float64 `json:"sslcipheraes256rate,omitempty"`
	Ssltot40bitrc2ciphers int `json:"ssltot40bitrc2ciphers,omitempty"`
	/**
	* Number of RC2 40-bit cipher encryptions on the Citrix ADC.
	*/
	Ssl40bitrc2ciphersrate float64 `json:"ssl40bitrc2ciphersrate,omitempty"`
	Ssltot56bitrc2ciphers int `json:"ssltot56bitrc2ciphers,omitempty"`
	/**
	* Number of RC2 56-bit cipher encryptions on the Citrix ADC.
	*/
	Ssl56bitrc2ciphersrate float64 `json:"ssl56bitrc2ciphersrate,omitempty"`
	Ssltot128bitrc2ciphers int `json:"ssltot128bitrc2ciphers,omitempty"`
	/**
	* Number of RC2 128-bit cipher encryptions on the Citrix ADC.
	*/
	Ssl128bitrc2ciphersrate float64 `json:"ssl128bitrc2ciphersrate,omitempty"`
	Ssltot128bitaesgcmciphers int `json:"ssltot128bitaesgcmciphers,omitempty"`
	/**
	* Number of AEC-GCM 128-bit cipher encryptions on the Citrix ADC.
	*/
	Ssl128bitaesgcmciphersrate float64 `json:"ssl128bitaesgcmciphersrate,omitempty"`
	Ssltot256bitaesgcmciphers int `json:"ssltot256bitaesgcmciphers,omitempty"`
	/**
	* Number of AEC-GCM 256-bit cipher encryptions on the Citrix ADC.
	*/
	Ssl256bitaesgcmciphersrate float64 `json:"ssl256bitaesgcmciphersrate,omitempty"`
	Ssltotnullciphers int `json:"ssltotnullciphers,omitempty"`
	/**
	* Number of Null cipher encryptions on the Citrix ADC.
	*/
	Sslnullciphersrate float64 `json:"sslnullciphersrate,omitempty"`
	Ssltotmd5mac int `json:"ssltotmd5mac,omitempty"`
	/**
	* Number of MD5 hashes on the Citrix ADC.
	*/
	Sslmd5macrate float64 `json:"sslmd5macrate,omitempty"`
	Ssltotshamac int `json:"ssltotshamac,omitempty"`
	/**
	* Number of SHA hashes on the Citrix ADC.
	*/
	Sslshamacrate float64 `json:"sslshamacrate,omitempty"`
	Ssltotsha256mac int `json:"ssltotsha256mac,omitempty"`
	/**
	* Number of SHA256 hashes on the Citrix ADC.
	*/
	Sslsha256macrate float64 `json:"sslsha256macrate,omitempty"`
	Ssltotsha384mac int `json:"ssltotsha384mac,omitempty"`
	/**
	* Number of SHA384 hashes on the Citrix ADC.
	*/
	Sslsha384macrate float64 `json:"sslsha384macrate,omitempty"`
	Ssltotsslv3handshakes int `json:"ssltotsslv3handshakes,omitempty"`
	/**
	* Number of handshakes on SSLv3 on the Citrix ADC.
	*/
	Sslsslv3handshakesrate float64 `json:"sslsslv3handshakesrate,omitempty"`
	Ssltottlsv1handshakes int `json:"ssltottlsv1handshakes,omitempty"`
	/**
	* Number of SSL handshakes on TLSv1 on the Citrix ADC.
	*/
	Ssltlsv1handshakesrate float64 `json:"ssltlsv1handshakesrate,omitempty"`
	Ssltottlsv11handshakes int `json:"ssltottlsv11handshakes,omitempty"`
	/**
	* Number of SSL handshakes on TLSv1.1 on the Citrix ADC.
	*/
	Ssltlsv11handshakesrate float64 `json:"ssltlsv11handshakesrate,omitempty"`
	Ssltottlsv12handshakes int `json:"ssltottlsv12handshakes,omitempty"`
	/**
	* Number of SSL handshakes on TLSv1.2 on the Citrix ADC.
	*/
	Ssltlsv12handshakesrate float64 `json:"ssltlsv12handshakesrate,omitempty"`
	Ssltottlsv13handshakes int `json:"ssltottlsv13handshakes,omitempty"`
	/**
	* Number of SSL handshakes on TLSv1.3 on the Citrix ADC.
	*/
	Ssltlsv13handshakesrate float64 `json:"ssltlsv13handshakesrate,omitempty"`
	Ssltotdtlsv1handshakes int `json:"ssltotdtlsv1handshakes,omitempty"`
	/**
	* Number of SSL handshakes on DTLSv1 on the Citrix ADC.
	*/
	Ssldtlsv1handshakesrate float64 `json:"ssldtlsv1handshakesrate,omitempty"`
	Ssltotdtlsv12handshakes int `json:"ssltotdtlsv12handshakes,omitempty"`
	/**
	* Number of SSL handshakes on DTLSv1.2 on the Citrix ADC.
	*/
	Ssldtlsv12handshakesrate float64 `json:"ssldtlsv12handshakesrate,omitempty"`
	Ssltotsslv3clientauthentications int `json:"ssltotsslv3clientauthentications,omitempty"`
	/**
	* Number of client authentications done on SSLv3.
	*/
	Sslsslv3clientauthenticationsrate float64 `json:"sslsslv3clientauthenticationsrate,omitempty"`
	Ssltottlsv1clientauthentications int `json:"ssltottlsv1clientauthentications,omitempty"`
	/**
	* Number of client authentications done on TLSv1.
	*/
	Ssltlsv1clientauthenticationsrate float64 `json:"ssltlsv1clientauthenticationsrate,omitempty"`
	Ssltottlsv11clientauthentications int `json:"ssltottlsv11clientauthentications,omitempty"`
	/**
	* Number of client authentications done on TLSv1.1.
	*/
	Ssltlsv11clientauthenticationsrate float64 `json:"ssltlsv11clientauthenticationsrate,omitempty"`
	Ssltottlsv12clientauthentications int `json:"ssltottlsv12clientauthentications,omitempty"`
	/**
	* Number of client authentications done on TLSv1.2.
	*/
	Ssltlsv12clientauthenticationsrate float64 `json:"ssltlsv12clientauthenticationsrate,omitempty"`
	Ssltottlsv13clientauthentications int `json:"ssltottlsv13clientauthentications,omitempty"`
	/**
	* Number of client authentications done on TLSv1.3.
	*/
	Ssltlsv13clientauthenticationsrate float64 `json:"ssltlsv13clientauthenticationsrate,omitempty"`
	Ssltotdtlsv1clientauthentications int `json:"ssltotdtlsv1clientauthentications,omitempty"`
	/**
	* Number of client authentications done on DTLSv1.
	*/
	Ssldtlsv1clientauthenticationsrate float64 `json:"ssldtlsv1clientauthenticationsrate,omitempty"`
	Ssltotdtlsv12clientauthentications int `json:"ssltotdtlsv12clientauthentications,omitempty"`
	/**
	* Number of client authentications done on DTLSv1.2.
	*/
	Ssldtlsv12clientauthenticationsrate float64 `json:"ssldtlsv12clientauthenticationsrate,omitempty"`
	Ssltotrsaauthorizations int `json:"ssltotrsaauthorizations,omitempty"`
	/**
	* Number of RSA authentications on the Citrix ADC.
	*/
	Sslrsaauthorizationsrate float64 `json:"sslrsaauthorizationsrate,omitempty"`
	Ssltotdhauthorizations int `json:"ssltotdhauthorizations,omitempty"`
	/**
	* Number of Diffie-Helman authentications on the Citrix ADC.
	*/
	Ssldhauthorizationsrate float64 `json:"ssldhauthorizationsrate,omitempty"`
	Ssltotdssauthorizations int `json:"ssltotdssauthorizations,omitempty"`
	/**
	* Total number of times DSS authorization is used on the Citrix ADC.
	*/
	Ssldssauthorizationsrate float64 `json:"ssldssauthorizationsrate,omitempty"`
	Ssltotecdsaauthorizations int `json:"ssltotecdsaauthorizations,omitempty"`
	/**
	* Total number of times ECDSA authorization is used on the Citrix ADC.
	*/
	Sslecdsaauthorizationsrate float64 `json:"sslecdsaauthorizationsrate,omitempty"`
	Ssltotnullauthorizations int `json:"ssltotnullauthorizations,omitempty"`
	/**
	* Number of Null authentications on the Citrix ADC.
	*/
	Sslnullauthorizationsrate float64 `json:"sslnullauthorizationsrate,omitempty"`
	Ssltotbkendsessionrenegotiate int `json:"ssltotbkendsessionrenegotiate,omitempty"`
	/**
	* Number of back-end SSL session renegotiations on the Citrix ADC.
	*/
	Sslbkendsessionrenegotiaterate float64 `json:"sslbkendsessionrenegotiaterate,omitempty"`
	Ssltotbkendsslv3renego int `json:"ssltotbkendsslv3renego,omitempty"`
	/**
	* Number of back-end SSLv3 session renegotiations on the Citrix ADC.
	*/
	Sslbkendsslv3renegorate float64 `json:"sslbkendsslv3renegorate,omitempty"`
	Ssltotbkendtlsvlrenego int `json:"ssltotbkendtlsvlrenego,omitempty"`
	/**
	* Number of back-end TLSv1 session renegotiations on the Citrix ADC.
	*/
	Sslbkendtlsvlrenegorate float64 `json:"sslbkendtlsvlrenegorate,omitempty"`
	Ssltotbkendtlsv11renego int `json:"ssltotbkendtlsv11renego,omitempty"`
	/**
	* Number of back-end TLSv1.1 session renegotiations on the Citrix ADC.
	*/
	Sslbkendtlsv11renegorate float64 `json:"sslbkendtlsv11renegorate,omitempty"`
	Ssltotbkendtlsv12renego int `json:"ssltotbkendtlsv12renego,omitempty"`
	/**
	* Number of back-end TLSv1.2 session renegotiations on the Citrix ADC.
	*/
	Sslbkendtlsv12renegorate float64 `json:"sslbkendtlsv12renegorate,omitempty"`
	Ssltotbkenddtlsvlrenego int `json:"ssltotbkenddtlsvlrenego,omitempty"`
	/**
	* Number of back-end DTLSv1 session renegotiations on the Citrix ADC.
	*/
	Sslbkenddtlsvlrenegorate float64 `json:"sslbkenddtlsvlrenegorate,omitempty"`
	Ssltotbkenddtlsvl2renego int `json:"ssltotbkenddtlsvl2renego,omitempty"`
	/**
	* Number of back-end DTLSv1.2 session renegotiations on the Citrix ADC.
	*/
	Sslbkenddtlsvl2renegorate float64 `json:"sslbkenddtlsvl2renegorate,omitempty"`
	Sslbetotrsa512keyexchanges int `json:"sslbetotrsa512keyexchanges,omitempty"`
	/**
	* Number of back-end RSA 512-bit key exchanges on the Citrix ADC.
	*/
	Sslbersa512keyexchangesrate float64 `json:"sslbersa512keyexchangesrate,omitempty"`
	Sslbetotrsa1024keyexchanges int `json:"sslbetotrsa1024keyexchanges,omitempty"`
	/**
	* Number of back-end RSA 1024-bit key exchanges on the Citrix ADC.
	*/
	Sslbersa1024keyexchangesrate float64 `json:"sslbersa1024keyexchangesrate,omitempty"`
	Sslbetotrsa2048keyexchanges int `json:"sslbetotrsa2048keyexchanges,omitempty"`
	/**
	* Number of back-end RSA 2048-bit key exchanges on the Citrix ADC.
	*/
	Sslbersa2048keyexchangesrate float64 `json:"sslbersa2048keyexchangesrate,omitempty"`
	Sslbetotrsa3072keyexchanges int `json:"sslbetotrsa3072keyexchanges,omitempty"`
	/**
	* Number of back-end RSA 3072-bit key exchanges on the Citrix ADC.
	*/
	Sslbersa3072keyexchangesrate float64 `json:"sslbersa3072keyexchangesrate,omitempty"`
	Sslbetotrsa4096keyexchanges int `json:"sslbetotrsa4096keyexchanges,omitempty"`
	/**
	* Number of back-end RSA 4096-bit key exchanges on the Citrix ADC.
	*/
	Sslbersa4096keyexchangesrate float64 `json:"sslbersa4096keyexchangesrate,omitempty"`
	Sslbetotdh512keyexchanges int `json:"sslbetotdh512keyexchanges,omitempty"`
	/**
	* Number of back-end DH 512-bit key exchanges on the Citrix ADC.
	*/
	Sslbedh512keyexchangesrate float64 `json:"sslbedh512keyexchangesrate,omitempty"`
	Sslbetotdh1024keyexchanges int `json:"sslbetotdh1024keyexchanges,omitempty"`
	/**
	* Number of back-end DH 1024-bit key exchanges on the Citrix ADC.
	*/
	Sslbedh1024keyexchangesrate float64 `json:"sslbedh1024keyexchangesrate,omitempty"`
	Sslbetotdh2048keyexchanges int `json:"sslbetotdh2048keyexchanges,omitempty"`
	/**
	* Number of back-end DH 2048-bit key exchanges on the Citrix ADC.
	*/
	Sslbedh2048keyexchangesrate float64 `json:"sslbedh2048keyexchangesrate,omitempty"`
	Sslbetotdh3072keyexchanges int `json:"sslbetotdh3072keyexchanges,omitempty"`
	/**
	* Number of back-end DH 3072-bit key exchanges on the Citrix ADC.
	*/
	Sslbedh3072keyexchangesrate float64 `json:"sslbedh3072keyexchangesrate,omitempty"`
	Sslbetotdh4096keyexchanges int `json:"sslbetotdh4096keyexchanges,omitempty"`
	/**
	* Number of back-end DH 4096-bit key exchanges on the Citrix ADC.
	*/
	Sslbedh4096keyexchangesrate float64 `json:"sslbedh4096keyexchangesrate,omitempty"`
	Sslbetotecdhecurve521 int `json:"sslbetotecdhecurve521,omitempty"`
	/**
	* Number of back-end ECDHE 521 curve Key exchanges  on the Citrix ADC.
	*/
	Sslbeecdhecurve521rate float64 `json:"sslbeecdhecurve521rate,omitempty"`
	Sslbetotecdhecurve384 int `json:"sslbetotecdhecurve384,omitempty"`
	/**
	* Number of back-end ECDHE 384 curve Key exchanges  on the Citrix ADC.
	*/
	Sslbeecdhecurve384rate float64 `json:"sslbeecdhecurve384rate,omitempty"`
	Sslbetotecdhecurve256 int `json:"sslbetotecdhecurve256,omitempty"`
	/**
	* Number of back-end ECDHE 256 curve Key exchanges  on the Citrix ADC.
	*/
	Sslbeecdhecurve256rate float64 `json:"sslbeecdhecurve256rate,omitempty"`
	Sslbetotecdhecurve224 int `json:"sslbetotecdhecurve224,omitempty"`
	/**
	* Number of back-end ECDHE 224 curve Key exchanges  on the Citrix ADC.
	*/
	Sslbeecdhecurve224rate float64 `json:"sslbeecdhecurve224rate,omitempty"`
	Sslbetotecdhecurve25519 int `json:"sslbetotecdhecurve25519,omitempty"`
	/**
	* Number of back-end ECDHE 25519 curve Key exchanges  on the Citrix ADC.
	*/
	Sslbeecdhecurve25519rate float64 `json:"sslbeecdhecurve25519rate,omitempty"`
	Sslbetot40bitrc4ciphers int `json:"sslbetot40bitrc4ciphers,omitempty"`
	/**
	* Number of back-end RC4 40-bit cipher encryptions on the Citrix ADC.
	*/
	Sslbe40bitrc4ciphersrate float64 `json:"sslbe40bitrc4ciphersrate,omitempty"`
	Sslbetot56bitrc4ciphers int `json:"sslbetot56bitrc4ciphers,omitempty"`
	/**
	* Number of back-end RC4 56-bit cipher encryptions on the Citrix ADC.
	*/
	Sslbe56bitrc4ciphersrate float64 `json:"sslbe56bitrc4ciphersrate,omitempty"`
	Sslbetot64bitrc4ciphers int `json:"sslbetot64bitrc4ciphers,omitempty"`
	/**
	* Number of back-end RC4 64-bit cipher encryptions on the Citrix ADC.
	*/
	Sslbe64bitrc4ciphersrate float64 `json:"sslbe64bitrc4ciphersrate,omitempty"`
	Sslbetot128bitrc4ciphers int `json:"sslbetot128bitrc4ciphers,omitempty"`
	/**
	* Number of back-end RC4 128-bit cipher encryptions on the Citrix ADC.
	*/
	Sslbe128bitrc4ciphersrate float64 `json:"sslbe128bitrc4ciphersrate,omitempty"`
	Sslbetot40bitdesciphers int `json:"sslbetot40bitdesciphers,omitempty"`
	/**
	* Number of back-end DES 40-bit cipher encryptions on the Citrix ADC.
	*/
	Sslbe40bitdesciphersrate float64 `json:"sslbe40bitdesciphersrate,omitempty"`
	Sslbetot56bitdesciphers int `json:"sslbetot56bitdesciphers,omitempty"`
	/**
	* Number of back-end DES 56-bit cipher encryptions on the Citrix ADC.
	*/
	Sslbe56bitdesciphersrate float64 `json:"sslbe56bitdesciphersrate,omitempty"`
	Sslbetot168bit3desciphers int `json:"sslbetot168bit3desciphers,omitempty"`
	/**
	* Number of back-end 3DES 168-bit cipher encryptions on the Citrix ADC.
	*/
	Sslbe168bit3desciphersrate float64 `json:"sslbe168bit3desciphersrate,omitempty"`
	Ssltotbkendcipheraes128 int `json:"ssltotbkendcipheraes128,omitempty"`
	/**
	* Back-end AES 128-bit cipher encryptions on the Citrix ADC.
	*/
	Sslbkendcipheraes128rate float64 `json:"sslbkendcipheraes128rate,omitempty"`
	Ssltotbkendcipheraes256 int `json:"ssltotbkendcipheraes256,omitempty"`
	/**
	* Back-end AES 256-bit cipher encryptions on the Citrix ADC.
	*/
	Sslbkendcipheraes256rate float64 `json:"sslbkendcipheraes256rate,omitempty"`
	Sslbetot40bitrc2ciphers int `json:"sslbetot40bitrc2ciphers,omitempty"`
	/**
	* Number of back-end RC2 40-bit cipher encryptions on the Citrix ADC.
	*/
	Sslbe40bitrc2ciphersrate float64 `json:"sslbe40bitrc2ciphersrate,omitempty"`
	Sslbetot56bitrc2ciphers int `json:"sslbetot56bitrc2ciphers,omitempty"`
	/**
	* Number of back-end RC2 56-bit cipher encryptions on the Citrix ADC.
	*/
	Sslbe56bitrc2ciphersrate float64 `json:"sslbe56bitrc2ciphersrate,omitempty"`
	Sslbetot128bitrc2ciphers int `json:"sslbetot128bitrc2ciphers,omitempty"`
	/**
	* Number of back-end RC2 128-bit cipher encryptions on the Citrix ADC.
	*/
	Sslbe128bitrc2ciphersrate float64 `json:"sslbe128bitrc2ciphersrate,omitempty"`
	Ssltotbkendcipheraesgcm128 int `json:"ssltotbkendcipheraesgcm128,omitempty"`
	/**
	* Back-end AES-GCM 128-bit cipher encryptions on the Citrix ADC.
	*/
	Sslbkendcipheraesgcm128rate float64 `json:"sslbkendcipheraesgcm128rate,omitempty"`
	Ssltotbkendcipheraesgcm256 int `json:"ssltotbkendcipheraesgcm256,omitempty"`
	/**
	* Back-end AES-GCM 256-bit cipher encryptions on the Citrix ADC .
	*/
	Sslbkendcipheraesgcm256rate float64 `json:"sslbkendcipheraesgcm256rate,omitempty"`
	Sslbetotnullciphers int `json:"sslbetotnullciphers,omitempty"`
	/**
	* Number of back-end null cipher encryptions on the Citrix ADC.
	*/
	Sslbenullciphersrate float64 `json:"sslbenullciphersrate,omitempty"`
	Sslbetotmd5mac int `json:"sslbetotmd5mac,omitempty"`
	/**
	* Number of back-end MD5 hashes on the Citrix ADC.
	*/
	Sslbemd5macrate float64 `json:"sslbemd5macrate,omitempty"`
	Sslbetotshamac int `json:"sslbetotshamac,omitempty"`
	/**
	* Number of back-end SHA hashes on the Citrix ADC.
	*/
	Sslbeshamacrate float64 `json:"sslbeshamacrate,omitempty"`
	Sslbetotsha256mac int `json:"sslbetotsha256mac,omitempty"`
	/**
	* Number of back-end SHA256 hashes on the Citrix ADC.
	*/
	Sslbesha256macrate float64 `json:"sslbesha256macrate,omitempty"`
	Sslbetotsha384mac int `json:"sslbetotsha384mac,omitempty"`
	/**
	* Number of back-end SHA384 hashes on the Citrix ADC.
	*/
	Sslbesha384macrate float64 `json:"sslbesha384macrate,omitempty"`
	Sslbetotsslv3handshakes int `json:"sslbetotsslv3handshakes,omitempty"`
	/**
	* Number of back-end SSLv3 handshakes on the Citrix ADC.
	*/
	Sslbesslv3handshakesrate float64 `json:"sslbesslv3handshakesrate,omitempty"`
	Sslbetottlsv1handshakes int `json:"sslbetottlsv1handshakes,omitempty"`
	/**
	* Number of back-end TLSv1 handshakes on the Citrix ADC.
	*/
	Sslbetlsv1handshakesrate float64 `json:"sslbetlsv1handshakesrate,omitempty"`
	Sslbetottlsv11handshakes int `json:"sslbetottlsv11handshakes,omitempty"`
	/**
	* Number of back-end TLSv1.1 handshakes on the Citrix ADC.
	*/
	Sslbetlsv11handshakesrate float64 `json:"sslbetlsv11handshakesrate,omitempty"`
	Sslbetottlsv12handshakes int `json:"sslbetottlsv12handshakes,omitempty"`
	/**
	* Number of back-end TLSv1.2 handshakes on the Citrix ADC.
	*/
	Sslbetlsv12handshakesrate float64 `json:"sslbetlsv12handshakesrate,omitempty"`
	Sslbetottlsv13handshakes int `json:"sslbetottlsv13handshakes,omitempty"`
	/**
	* Number of back-end TLSv1.3 handshakes on the Citrix ADC.
	*/
	Sslbetlsv13handshakesrate float64 `json:"sslbetlsv13handshakesrate,omitempty"`
	Sslbetotdtlsv1handshakes int `json:"sslbetotdtlsv1handshakes,omitempty"`
	/**
	* Number of back-end DTLSv1 handshakes on the Citrix ADC.
	*/
	Sslbedtlsv1handshakesrate float64 `json:"sslbedtlsv1handshakesrate,omitempty"`
	Sslbetotdtlsv12handshakes int `json:"sslbetotdtlsv12handshakes,omitempty"`
	/**
	* Number of back-end DTLSv1.2 handshakes on the Citrix ADC.
	*/
	Sslbedtlsv12handshakesrate float64 `json:"sslbedtlsv12handshakesrate,omitempty"`
	Sslbetotsslv3clientauthentications int `json:"sslbetotsslv3clientauthentications,omitempty"`
	/**
	* Number of back-end SSLv3 client authentications on the Citrix ADC.
	*/
	Sslbesslv3clientauthenticationsrate float64 `json:"sslbesslv3clientauthenticationsrate,omitempty"`
	Sslbetottlsv1clientauthentications int `json:"sslbetottlsv1clientauthentications,omitempty"`
	/**
	* Number of back-end TLSv1 client authentications on the Citrix ADC.
	*/
	Sslbetlsv1clientauthenticationsrate float64 `json:"sslbetlsv1clientauthenticationsrate,omitempty"`
	Sslbetottlsv11clientauthentications int `json:"sslbetottlsv11clientauthentications,omitempty"`
	/**
	* Number of back-end TLSv1.1 client authentications on the Citrix ADC.
	*/
	Sslbetlsv11clientauthenticationsrate float64 `json:"sslbetlsv11clientauthenticationsrate,omitempty"`
	Sslbetottlsv12clientauthentications int `json:"sslbetottlsv12clientauthentications,omitempty"`
	/**
	* Number of back-end TLSv1.2 client authentications on the Citrix ADC.
	*/
	Sslbetlsv12clientauthenticationsrate float64 `json:"sslbetlsv12clientauthenticationsrate,omitempty"`
	Sslbetottlsv13clientauthentications int `json:"sslbetottlsv13clientauthentications,omitempty"`
	/**
	* Number of back-end TLSv1.3 client authentications on the Citrix ADC.
	*/
	Sslbetlsv13clientauthenticationsrate float64 `json:"sslbetlsv13clientauthenticationsrate,omitempty"`
	Sslbetotdtlsv1clientauthentications int `json:"sslbetotdtlsv1clientauthentications,omitempty"`
	/**
	* Number of back-end DTLSv1 client authentications on the Citrix ADC.
	*/
	Sslbedtlsv1clientauthenticationsrate float64 `json:"sslbedtlsv1clientauthenticationsrate,omitempty"`
	Sslbetotdtlsv12clientauthentications int `json:"sslbetotdtlsv12clientauthentications,omitempty"`
	/**
	* Number of back-end DTLSv1.2 client authentications on the Citrix ADC.
	*/
	Sslbedtlsv12clientauthenticationsrate float64 `json:"sslbedtlsv12clientauthenticationsrate,omitempty"`
	Sslbetotrsaauthorizations int `json:"sslbetotrsaauthorizations,omitempty"`
	/**
	* Number of back-end RSA authentications on the Citrix ADC.
	*/
	Sslbersaauthorizationsrate float64 `json:"sslbersaauthorizationsrate,omitempty"`
	Sslbetotdhauthorizations int `json:"sslbetotdhauthorizations,omitempty"`
	/**
	* Number of back-end DH authentications on the Citrix ADC.
	*/
	Sslbedhauthorizationsrate float64 `json:"sslbedhauthorizationsrate,omitempty"`
	Sslbetotdssauthorizations int `json:"sslbetotdssauthorizations,omitempty"`
	/**
	* Number of back-end DSS authentications on the Citrix ADC.
	*/
	Sslbedssauthorizationsrate float64 `json:"sslbedssauthorizationsrate,omitempty"`
	Sslbetotecdsaauthorizations int `json:"sslbetotecdsaauthorizations,omitempty"`
	/**
	* Number of back-end ECDSA authentications on the Citrix ADC.
	*/
	Sslbeecdsaauthorizationsrate float64 `json:"sslbeecdsaauthorizationsrate,omitempty"`
	Sslbetotnullauthorizations int `json:"sslbetotnullauthorizations,omitempty"`
	/**
	* Number of back-end null authentications on the Citrix ADC.
	*/
	Sslbenullauthorizationsrate float64 `json:"sslbenullauthorizationsrate,omitempty"`
	Ssltotoffloadrsakeyexchanges int `json:"ssltotoffloadrsakeyexchanges,omitempty"`
	/**
	* Number of RSA key exchanges offloaded to the cryptography card.
	*/
	Ssloffloadrsakeyexchangesrate float64 `json:"ssloffloadrsakeyexchangesrate,omitempty"`
	Ssltotoffloadsignrsa int `json:"ssltotoffloadsignrsa,omitempty"`
	/**
	* Number of RSA sign operations offloaded to the cryptography card.
	*/
	Ssloffloadsignrsarate float64 `json:"ssloffloadsignrsarate,omitempty"`
	Ssltotoffloaddhkeyexchanges int `json:"ssltotoffloaddhkeyexchanges,omitempty"`
	/**
	* Number of DH key exchanges offloaded to the cryptography card.
	*/
	Ssloffloaddhkeyexchangesrate float64 `json:"ssloffloaddhkeyexchangesrate,omitempty"`
	Ssltotoffloadbulkrc4 int `json:"ssltotoffloadbulkrc4,omitempty"`
	/**
	* Number of RC4 encryptions offloaded to the cryptography card.
	*/
	Ssloffloadbulkrc4rate float64 `json:"ssloffloadbulkrc4rate,omitempty"`
	Ssltotoffloadbulkdes int `json:"ssltotoffloadbulkdes,omitempty"`
	/**
	* Number of DES encryptions offloaded to the cryptography card.
	*/
	Ssloffloadbulkdesrate float64 `json:"ssloffloadbulkdesrate,omitempty"`
	Ssltotoffloadbulkaes int `json:"ssltotoffloadbulkaes,omitempty"`
	/**
	* Number of AES encryptions offloaded to the cryptography card.
	*/
	Ssloffloadbulkaesrate float64 `json:"ssloffloadbulkaesrate,omitempty"`
	Ssltotoffloadbulkaesgcm128 int `json:"ssltotoffloadbulkaesgcm128,omitempty"`
	/**
	* Number of AES-GCM 128-bit encryptions offloaded to the cryptography card.
	*/
	Ssloffloadbulkaesgcm128rate float64 `json:"ssloffloadbulkaesgcm128rate,omitempty"`
	Ssltotoffloadbulkaesgcm256 int `json:"ssltotoffloadbulkaesgcm256,omitempty"`
	/**
	* Number of AES-GCM 256-bit encryptions offloaded to the cryptography card.
	*/
	Ssloffloadbulkaesgcm256rate float64 `json:"ssloffloadbulkaesgcm256rate,omitempty"`
	Ssltotenchw int `json:"ssltotenchw,omitempty"`
	/**
	* Number of bytes encrypted in hardware.
	*/
	Sslenchwrate float64 `json:"sslenchwrate,omitempty"`
	Ssltotencsw int `json:"ssltotencsw,omitempty"`
	/**
	* Number of bytes encrypted in software.
	*/
	Sslencswrate float64 `json:"sslencswrate,omitempty"`
	Ssltotencfe int `json:"ssltotencfe,omitempty"`
	/**
	* Number of bytes encrypted on the front-end.
	*/
	Sslencferate float64 `json:"sslencferate,omitempty"`
	Ssltothwencfe int `json:"ssltothwencfe,omitempty"`
	/**
	* Number of bytes encrypted in hardware on the front-end.
	*/
	Sslhwencferate float64 `json:"sslhwencferate,omitempty"`
	Ssltotswencfe int `json:"ssltotswencfe,omitempty"`
	/**
	* Number of bytes encrypted in software on the front-end.
	*/
	Sslswencferate float64 `json:"sslswencferate,omitempty"`
	Ssltotencbe int `json:"ssltotencbe,omitempty"`
	/**
	* Number of bytes encrypted on the back-end.
	*/
	Sslencberate float64 `json:"sslencberate,omitempty"`
	Ssltothwencbe int `json:"ssltothwencbe,omitempty"`
	/**
	* Number of bytes encrypted in hardware on the back-end.
	*/
	Sslhwencberate float64 `json:"sslhwencberate,omitempty"`
	Ssltotswencbe int `json:"ssltotswencbe,omitempty"`
	/**
	* Number of bytes encrypted in software on the back-end.
	*/
	Sslswencberate float64 `json:"sslswencberate,omitempty"`
	Ssltotdechw int `json:"ssltotdechw,omitempty"`
	/**
	* Number of bytes decrypted in hardware.
	*/
	Ssldechwrate float64 `json:"ssldechwrate,omitempty"`
	Ssltotdecsw int `json:"ssltotdecsw,omitempty"`
	/**
	* Number of bytes decrypted in software.
	*/
	Ssldecswrate float64 `json:"ssldecswrate,omitempty"`
	Ssltotdecfe int `json:"ssltotdecfe,omitempty"`
	/**
	* Number of bytes decrypted on the front-end.
	*/
	Ssldecferate float64 `json:"ssldecferate,omitempty"`
	Ssltothwdecfe int `json:"ssltothwdecfe,omitempty"`
	/**
	* Number of bytes decrypted in hardware on the front-end.
	*/
	Sslhwdecferate float64 `json:"sslhwdecferate,omitempty"`
	Ssltotswdecfe int `json:"ssltotswdecfe,omitempty"`
	/**
	* Number of bytes decrypted in software on the front-end.
	*/
	Sslswdecferate float64 `json:"sslswdecferate,omitempty"`
	Ssltotdecbe int `json:"ssltotdecbe,omitempty"`
	/**
	* Number of bytes decrypted on the back-end.
	*/
	Ssldecberate float64 `json:"ssldecberate,omitempty"`
	Ssltothwdecbe int `json:"ssltothwdecbe,omitempty"`
	/**
	* Number of bytes decrypted in hardware on the back-end.
	*/
	Sslhwdecberate float64 `json:"sslhwdecberate,omitempty"`
	Ssltotswdecbe int `json:"ssltotswdecbe,omitempty"`
	/**
	* Number of bytes decrypted in software on the back-end
	*/
	Sslswdecberate float64 `json:"sslswdecberate,omitempty"`
	Sslcursslinfospcbinusecount int `json:"sslcursslinfospcbinusecount,omitempty"`
	/**
	* Number of SPCB in use.
	*/
	Sslcursslinfospcbinusecountrate float64 `json:"sslcursslinfospcbinusecountrate,omitempty"`
	Sslcursessions int `json:"sslcursessions,omitempty"`
	/**
	* Number of active SSL sessions on the Citrix ADC.
	*/
	Sslcursessionsrate float64 `json:"sslcursessionsrate,omitempty"`
	Sslcurqsize int `json:"sslcurqsize,omitempty"`
	/**
	* Current queue size
	*/
	Sslcurqsizerate float64 `json:"sslcurqsizerate,omitempty"`
	Sslcursslinfonscardinqcount int `json:"sslcursslinfonscardinqcount,omitempty"`
	/**
	* Number of current SSL card InQ count.
	*/
	Sslcursslinfonscardinqcountrate float64 `json:"sslcursslinfonscardinqcountrate,omitempty"`
	Sslcursslinfocardinblkq int `json:"sslcursslinfocardinblkq,omitempty"`
	/**
	* Number of current SSL card In BulkQ count.
	*/
	Sslcursslinfocardinblkqrate float64 `json:"sslcursslinfocardinblkqrate,omitempty"`
	Sslcursslinfocardinkeyq int `json:"sslcursslinfocardinkeyq,omitempty"`
	/**
	* Number of current SSL card In KeyQ count.
	*/
	Sslcursslinfocardinkeyqrate float64 `json:"sslcursslinfocardinkeyqrate,omitempty"`
	Sslbemaxmultiplexedsessions int `json:"sslbemaxmultiplexedsessions,omitempty"`
	/**
	* Number of back-end SSL sessions reused on the Citrix ADC.
	*/
	Sslbemaxmultiplexedsessionsrate float64 `json:"sslbemaxmultiplexedsessionsrate,omitempty"`
	Ssltot128bitideaciphers int `json:"ssltot128bitideaciphers,omitempty"`
	/**
	* Number of IDEA 128-bit cipher encryptions on the Citrix ADC.
	*/
	Ssl128bitideaciphersrate float64 `json:"ssl128bitideaciphersrate,omitempty"`
	Sslbetot128bitideaciphers int `json:"sslbetot128bitideaciphers,omitempty"`
	/**
	* Number of back-end IDEA 128-bit cipher encryptions on the Citrix ADC.
	*/
	Sslbe128bitideaciphersrate float64 `json:"sslbe128bitideaciphersrate,omitempty"`

}
