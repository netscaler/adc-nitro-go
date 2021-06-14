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

package bot

/**
* Statistics for Bot profile resource.
*/

type Botprofilestats struct {
	/**
	* Name of the bot profile.
	*/
	Name string `json:"name,omitempty"`
	/**
	* Clear the statsistics / counters
	*/
	Clearstats string `json:"clearstats,omitempty"`
	Botrequestsperprofile uint64 `json:"botrequestsperprofile,omitempty"`
	/**
	* HTTP/HTTPS requests sent to your protected web servers via the Bot profile.
	*/
	Botrequestsperprofilerate int32 `json:"botrequestsperprofilerate,omitempty"`
	Botreqbytesperprofile uint64 `json:"botreqbytesperprofile,omitempty"`
	/**
	* Number of bytes transfered for requests
	*/
	Botreqbytesperprofilerate int32 `json:"botreqbytesperprofilerate,omitempty"`
	Botresponsesperprofile uint64 `json:"botresponsesperprofile,omitempty"`
	/**
	* HTTP/HTTPS responses sent by your protected web servers via the Bot profile.
	*/
	Botresponsesperprofilerate int32 `json:"botresponsesperprofilerate,omitempty"`
	Botresbytesperprofile uint64 `json:"botresbytesperprofile,omitempty"`
	/**
	* Number of bytes transfered for responses
	*/
	Botresbytesperprofilerate int32 `json:"botresbytesperprofilerate,omitempty"`
	Bottotallogprofile uint64 `json:"bottotallogprofile,omitempty"`
	/**
	* Total number of logs by the Bot profile.
	*/
	Botlogprofilerate int32 `json:"botlogprofilerate,omitempty"`
	Bottotaldropprofile uint64 `json:"bottotaldropprofile,omitempty"`
	/**
	* Total number of drops by the Bot profile.
	*/
	Botdropprofilerate int32 `json:"botdropprofilerate,omitempty"`
	Bottotalredirectprofile uint64 `json:"bottotalredirectprofile,omitempty"`
	/**
	* Total number of redirects by the Bot profile.
	*/
	Botredirectprofilerate int32 `json:"botredirectprofilerate,omitempty"`
	Bottotalresetprofile uint64 `json:"bottotalresetprofile,omitempty"`
	/**
	* Total number of resets by the Bot profile.
	*/
	Botresetprofilerate int32 `json:"botresetprofilerate,omitempty"`
	Botvioldevicefingerprintprofile uint64 `json:"botvioldevicefingerprintprofile,omitempty"`
	/**
	* Number of device fingerprint violations seen by the Bot profile.
	*/
	Botvioldevicefingerprintprofilerate int32 `json:"botvioldevicefingerprintprofilerate,omitempty"`
	Botvioldevicefingerprintlogprofile uint64 `json:"botvioldevicefingerprintlogprofile,omitempty"`
	/**
	* Number of device fingerprint violations logged by the Bot profile.
	*/
	Botvioldevicefingerprintlogprofilerate int32 `json:"botvioldevicefingerprintlogprofilerate,omitempty"`
	Botvioldevicefingerprintdropprofile uint64 `json:"botvioldevicefingerprintdropprofile,omitempty"`
	/**
	* Number of device fingerprint violations dropped by the Bot profile.
	*/
	Botvioldevicefingerprintdropprofilerate int32 `json:"botvioldevicefingerprintdropprofilerate,omitempty"`
	Botvioldevicefingerprintredirectprofile uint64 `json:"botvioldevicefingerprintredirectprofile,omitempty"`
	/**
	* Number of device fingerprint violations requests redirected by the Bot profile to a different Web page or web server.
	*/
	Botvioldevicefingerprintredirectprofilerate int32 `json:"botvioldevicefingerprintredirectprofilerate,omitempty"`
	Botvioldevicefingerprintcaptchaprofile uint64 `json:"botvioldevicefingerprintcaptchaprofile,omitempty"`
	/**
	* Number of device fingerprint violation requests for which CAPTCHA challenge was sent due to Bot profile.
	*/
	Botvioldevicefingerprintcaptchaprofilerate int32 `json:"botvioldevicefingerprintcaptchaprofilerate,omitempty"`
	Botvioldevicefingerprintresetprofile uint64 `json:"botvioldevicefingerprintresetprofile,omitempty"`
	/**
	* Number of device fingerprint violations reset by the Bot profile.
	*/
	Botvioldevicefingerprintresetprofilerate int32 `json:"botvioldevicefingerprintresetprofilerate,omitempty"`
	Botviolipreputationprofile uint64 `json:"botviolipreputationprofile,omitempty"`
	/**
	* Number of ip reputation violations seen by the Bot profile.
	*/
	Botviolipreputationprofilerate int32 `json:"botviolipreputationprofilerate,omitempty"`
	Botviolipreputationlogprofile uint64 `json:"botviolipreputationlogprofile,omitempty"`
	/**
	* Number of ip reputation violations logged by the Bot Profile.
	*/
	Botviolipreputationlogprofilerate int32 `json:"botviolipreputationlogprofilerate,omitempty"`
	Botviolipreputationdropprofile uint64 `json:"botviolipreputationdropprofile,omitempty"`
	/**
	* Number of ip reputation violations dropped by the Bot profile.
	*/
	Botviolipreputationdropprofilerate int32 `json:"botviolipreputationdropprofilerate,omitempty"`
	Botviolipreputationredirectprofile uint64 `json:"botviolipreputationredirectprofile,omitempty"`
	/**
	* Number of ip reputation violations requests redirected by the Bot profile to a different Web page or web server.
	*/
	Botviolipreputationredirectprofilerate int32 `json:"botviolipreputationredirectprofilerate,omitempty"`
	Botviolipreputationcaptchaprofile uint64 `json:"botviolipreputationcaptchaprofile,omitempty"`
	/**
	* Number of ip reputation violation requests for which CAPTCHA challenge was sent due to Bot profile.
	*/
	Botviolipreputationcaptchaprofilerate int32 `json:"botviolipreputationcaptchaprofilerate,omitempty"`
	Botviolipreputationresetprofile uint64 `json:"botviolipreputationresetprofile,omitempty"`
	/**
	* Number of ip reputation violations reset by the Bot profile.
	*/
	Botviolipreputationresetprofilerate int32 `json:"botviolipreputationresetprofilerate,omitempty"`
	Botviolwhitelistprofile uint64 `json:"botviolwhitelistprofile,omitempty"`
	/**
	* Number of white list violations seen by the Bot profile.
	*/
	Botviolwhitelistprofilerate int32 `json:"botviolwhitelistprofilerate,omitempty"`
	Botviolwhitelistlogprofile uint64 `json:"botviolwhitelistlogprofile,omitempty"`
	/**
	* Number of white list violations logged by the Bot profile.
	*/
	Botviolwhitelistlogprofilerate int32 `json:"botviolwhitelistlogprofilerate,omitempty"`
	Botviolblacklistprofile uint64 `json:"botviolblacklistprofile,omitempty"`
	/**
	* Number of black list violations seen by the Bot profile.
	*/
	Botviolblacklistprofilerate int32 `json:"botviolblacklistprofilerate,omitempty"`
	Botviolblacklistlogprofile uint64 `json:"botviolblacklistlogprofile,omitempty"`
	/**
	* Number of black list violations logged by the Bot profile.
	*/
	Botviolblacklistlogprofilerate int32 `json:"botviolblacklistlogprofilerate,omitempty"`
	Botviolblacklistdropprofile uint64 `json:"botviolblacklistdropprofile,omitempty"`
	/**
	* Number of black list violations dropped by the Bot profile.
	*/
	Botviolblacklistdropprofilerate int32 `json:"botviolblacklistdropprofilerate,omitempty"`
	Botviolblacklistresetprofile uint64 `json:"botviolblacklistresetprofile,omitempty"`
	/**
	* Number of black list violations reset by the Bot profile.
	*/
	Botviolblacklistresetprofilerate int32 `json:"botviolblacklistresetprofilerate,omitempty"`
	Botviolblacklistredirectprofile uint64 `json:"botviolblacklistredirectprofile,omitempty"`
	/**
	* Number of black list violations redirected by the Bot profile to a different Web page or web server.
	*/
	Botviolblacklistredirectprofilerate int32 `json:"botviolblacklistredirectprofilerate,omitempty"`
	Botviolratelimitprofile uint64 `json:"botviolratelimitprofile,omitempty"`
	/**
	* Number of rate limiting violations seen by the Bot profile.
	*/
	Botviolratelimitprofilerate int32 `json:"botviolratelimitprofilerate,omitempty"`
	Botviolratelimitlogprofile uint64 `json:"botviolratelimitlogprofile,omitempty"`
	/**
	* Number of rate limiting violations logged by the Bot profile.
	*/
	Botviolratelimitlogprofilerate int32 `json:"botviolratelimitlogprofilerate,omitempty"`
	Botviolratelimitdropprofile uint64 `json:"botviolratelimitdropprofile,omitempty"`
	/**
	* Number of rate limiting violations dropped by the Bot profile.
	*/
	Botviolratelimitdropprofilerate int32 `json:"botviolratelimitdropprofilerate,omitempty"`
	Botviolratelimitredirectprofile uint64 `json:"botviolratelimitredirectprofile,omitempty"`
	/**
	* Number of rate limiting violations requests redirected by the Bot profile to a different Web page or web server.
	*/
	Botviolratelimitredirectprofilerate int32 `json:"botviolratelimitredirectprofilerate,omitempty"`
	Botviolratelimitresetprofile uint64 `json:"botviolratelimitresetprofile,omitempty"`
	/**
	* Number of rate limiting violations reset by the Bot profile.
	*/
	Botviolratelimitresetprofilerate int32 `json:"botviolratelimitresetprofilerate,omitempty"`
	Botviolstaticsignatureprofile uint64 `json:"botviolstaticsignatureprofile,omitempty"`
	/**
	* Number of static signatutre violations seen by the Bot profile.
	*/
	Botviolstaticsignatureprofilerate int32 `json:"botviolstaticsignatureprofilerate,omitempty"`
	Botviolstaticsignaturelogprofile uint64 `json:"botviolstaticsignaturelogprofile,omitempty"`
	/**
	* Number of static signatutre violations logged by the Bot profile.
	*/
	Botviolstaticsignaturelogprofilerate int32 `json:"botviolstaticsignaturelogprofilerate,omitempty"`
	Botviolstaticsignaturedropprofile uint64 `json:"botviolstaticsignaturedropprofile,omitempty"`
	/**
	* Number of static signatutre violations dropped by the Bot profile.
	*/
	Botviolstaticsignaturedropprofilerate int32 `json:"botviolstaticsignaturedropprofilerate,omitempty"`
	Botviolstaticsignatureredirectprofile uint64 `json:"botviolstaticsignatureredirectprofile,omitempty"`
	/**
	* Number of static signatutre violations redirected by the Bot profile to a different Web page or web server.
	*/
	Botviolstaticsignatureredirectprofilerate int32 `json:"botviolstaticsignatureredirectprofilerate,omitempty"`
	Botviolstaticsignatureresetprofile uint64 `json:"botviolstaticsignatureresetprofile,omitempty"`
	/**
	* Number of static signatutre violations reset by the Bot profile to a different Web page or web server.
	*/
	Botviolstaticsignatureresetprofilerate int32 `json:"botviolstaticsignatureresetprofilerate,omitempty"`
	Botvioltpsprofile uint64 `json:"botvioltpsprofile,omitempty"`
	/**
	* Number of tps violations seen by the Bot profile.
	*/
	Botvioltpsprofilerate int32 `json:"botvioltpsprofilerate,omitempty"`
	Botvioltpslogprofile uint64 `json:"botvioltpslogprofile,omitempty"`
	/**
	* Number of tps violations logged by the Bot profile.
	*/
	Botvioltpslogprofilerate int32 `json:"botvioltpslogprofilerate,omitempty"`
	Botvioltpsdropprofile uint64 `json:"botvioltpsdropprofile,omitempty"`
	/**
	* Number of tps violations dropped by the Bot profile.
	*/
	Botvioltpsdropprofilerate int32 `json:"botvioltpsdropprofilerate,omitempty"`
	Botvioltpsredirectprofile uint64 `json:"botvioltpsredirectprofile,omitempty"`
	/**
	* Number of tps violations requests redirected by the Bot profile to a different Web page or web server.
	*/
	Botvioltpsredirectprofilerate int32 `json:"botvioltpsredirectprofilerate,omitempty"`
	Botvioltpsresetprofile uint64 `json:"botvioltpsresetprofile,omitempty"`
	/**
	* Number of tps violations reset by the Bot profile.
	*/
	Botvioltpsresetprofilerate int32 `json:"botvioltpsresetprofilerate,omitempty"`
	Botvioltpscaptchaprofile uint64 `json:"botvioltpscaptchaprofile,omitempty"`
	/**
	* Number of tps violation requests for which CAPTCHA challenge was sent due to Bot profile.
	*/
	Botvioltpscaptchaprofilerate int32 `json:"botvioltpscaptchaprofilerate,omitempty"`
	Botviolcaptchaprofile uint64 `json:"botviolcaptchaprofile,omitempty"`
	/**
	* Number of Captcha challenge failures seen by the Bot profile.
	*/
	Botviolcaptchaprofilerate int32 `json:"botviolcaptchaprofilerate,omitempty"`
	Botviolcaptchalogprofile uint64 `json:"botviolcaptchalogprofile,omitempty"`
	/**
	* Number of Captcha challenge failures logged by the Bot profile.
	*/
	Botviolcaptchalogprofilerate int32 `json:"botviolcaptchalogprofilerate,omitempty"`
	Botviolcaptchadropprofile uint64 `json:"botviolcaptchadropprofile,omitempty"`
	/**
	* Number of Captcha challenge failures dropped by the Bot profile.
	*/
	Botviolcaptchadropprofilerate int32 `json:"botviolcaptchadropprofilerate,omitempty"`
	Botviolcaptcharedirectprofile uint64 `json:"botviolcaptcharedirectprofile,omitempty"`
	/**
	* Number of Captcha challenge failures redirected by the Bot profile.
	*/
	Botviolcaptcharedirectprofilerate int32 `json:"botviolcaptcharedirectprofilerate,omitempty"`
	Botviolcaptcharesetprofile uint64 `json:"botviolcaptcharesetprofile,omitempty"`
	/**
	* Number of Captcha challenge failures reset by the Bot profile.
	*/
	Botviolcaptcharesetprofilerate int32 `json:"botviolcaptcharesetprofilerate,omitempty"`
	Botvioltrapprofile uint64 `json:"botvioltrapprofile,omitempty"`
	/**
	* Number of trap violations seen by the Bot profile.
	*/
	Botvioltrapprofilerate int32 `json:"botvioltrapprofilerate,omitempty"`
	Botvioltraplogprofile uint64 `json:"botvioltraplogprofile,omitempty"`
	/**
	* Number of trap violations logged by the Bot profile.
	*/
	Botvioltraplogprofilerate int32 `json:"botvioltraplogprofilerate,omitempty"`
	Botvioltrapdropprofile uint64 `json:"botvioltrapdropprofile,omitempty"`
	/**
	* Number of trap violations dropped by the Bot profile.
	*/
	Botvioltrapdropprofilerate int32 `json:"botvioltrapdropprofilerate,omitempty"`
	Botvioltrapredirectprofile uint64 `json:"botvioltrapredirectprofile,omitempty"`
	/**
	* Number of trap violations requests redirected by the Bot profile to a different Web page or web server.
	*/
	Botvioltrapredirectprofilerate int32 `json:"botvioltrapredirectprofilerate,omitempty"`
	Botvioltrapresetprofile uint64 `json:"botvioltrapresetprofile,omitempty"`
	/**
	* Number of trap violations reset by the Bot profile.
	*/
	Botvioltrapresetprofilerate int32 `json:"botvioltrapresetprofilerate,omitempty"`

}
