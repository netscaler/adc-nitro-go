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
	Clearstats            string `json:"clearstats,omitempty"`
	Botrequestsperprofile int    `json:"botrequestsperprofile,omitempty"`
	/**
	* HTTP/HTTPS requests sent to your protected web servers via the Bot profile.
	 */
	Botrequestsperprofilerate float64 `json:"botrequestsperprofilerate,omitempty"`
	Botreqbytesperprofile     int     `json:"botreqbytesperprofile,omitempty"`
	/**
	* Number of bytes transfered for requests
	 */
	Botreqbytesperprofilerate float64 `json:"botreqbytesperprofilerate,omitempty"`
	Botresponsesperprofile    int     `json:"botresponsesperprofile,omitempty"`
	/**
	* HTTP/HTTPS responses sent by your protected web servers via the Bot profile.
	 */
	Botresponsesperprofilerate float64 `json:"botresponsesperprofilerate,omitempty"`
	Botresbytesperprofile      int     `json:"botresbytesperprofile,omitempty"`
	/**
	* Number of bytes transfered for responses
	 */
	Botresbytesperprofilerate float64 `json:"botresbytesperprofilerate,omitempty"`
	Bottotallogprofile        int     `json:"bottotallogprofile,omitempty"`
	/**
	* Total number of logs by the Bot profile.
	 */
	Botlogprofilerate   float64 `json:"botlogprofilerate,omitempty"`
	Bottotaldropprofile int     `json:"bottotaldropprofile,omitempty"`
	/**
	* Total number of drops by the Bot profile.
	 */
	Botdropprofilerate      float64 `json:"botdropprofilerate,omitempty"`
	Bottotalredirectprofile int     `json:"bottotalredirectprofile,omitempty"`
	/**
	* Total number of redirects by the Bot profile.
	 */
	Botredirectprofilerate float64 `json:"botredirectprofilerate,omitempty"`
	Bottotalresetprofile   int     `json:"bottotalresetprofile,omitempty"`
	/**
	* Total number of resets by the Bot profile.
	 */
	Botresetprofilerate             float64 `json:"botresetprofilerate,omitempty"`
	Botvioldevicefingerprintprofile int     `json:"botvioldevicefingerprintprofile,omitempty"`
	/**
	* Number of device fingerprint violations seen by the Bot profile.
	 */
	Botvioldevicefingerprintprofilerate float64 `json:"botvioldevicefingerprintprofilerate,omitempty"`
	Botvioldevicefingerprintlogprofile  int     `json:"botvioldevicefingerprintlogprofile,omitempty"`
	/**
	* Number of device fingerprint violations logged by the Bot profile.
	 */
	Botvioldevicefingerprintlogprofilerate float64 `json:"botvioldevicefingerprintlogprofilerate,omitempty"`
	Botvioldevicefingerprintdropprofile    int     `json:"botvioldevicefingerprintdropprofile,omitempty"`
	/**
	* Number of device fingerprint violations dropped by the Bot profile.
	 */
	Botvioldevicefingerprintdropprofilerate float64 `json:"botvioldevicefingerprintdropprofilerate,omitempty"`
	Botvioldevicefingerprintredirectprofile int     `json:"botvioldevicefingerprintredirectprofile,omitempty"`
	/**
	* Number of device fingerprint violations requests redirected by the Bot profile to a different Web page or web server.
	 */
	Botvioldevicefingerprintredirectprofilerate float64 `json:"botvioldevicefingerprintredirectprofilerate,omitempty"`
	Botvioldevicefingerprintcaptchaprofile      int     `json:"botvioldevicefingerprintcaptchaprofile,omitempty"`
	/**
	* Number of device fingerprint violation requests for which CAPTCHA challenge was sent due to Bot profile.
	 */
	Botvioldevicefingerprintcaptchaprofilerate float64 `json:"botvioldevicefingerprintcaptchaprofilerate,omitempty"`
	Botvioldevicefingerprintresetprofile       int     `json:"botvioldevicefingerprintresetprofile,omitempty"`
	/**
	* Number of device fingerprint violations reset by the Bot profile.
	 */
	Botvioldevicefingerprintresetprofilerate float64 `json:"botvioldevicefingerprintresetprofilerate,omitempty"`
	Botviolipreputationprofile               int     `json:"botviolipreputationprofile,omitempty"`
	/**
	* Number of ip reputation violations seen by the Bot profile.
	 */
	Botviolipreputationprofilerate float64 `json:"botviolipreputationprofilerate,omitempty"`
	Botviolipreputationlogprofile  int     `json:"botviolipreputationlogprofile,omitempty"`
	/**
	* Number of ip reputation violations logged by the Bot Profile.
	 */
	Botviolipreputationlogprofilerate float64 `json:"botviolipreputationlogprofilerate,omitempty"`
	Botviolipreputationdropprofile    int     `json:"botviolipreputationdropprofile,omitempty"`
	/**
	* Number of ip reputation violations dropped by the Bot profile.
	 */
	Botviolipreputationdropprofilerate float64 `json:"botviolipreputationdropprofilerate,omitempty"`
	Botviolipreputationredirectprofile int     `json:"botviolipreputationredirectprofile,omitempty"`
	/**
	* Number of ip reputation violations requests redirected by the Bot profile to a different Web page or web server.
	 */
	Botviolipreputationredirectprofilerate float64 `json:"botviolipreputationredirectprofilerate,omitempty"`
	Botviolipreputationcaptchaprofile      int     `json:"botviolipreputationcaptchaprofile,omitempty"`
	/**
	* Number of ip reputation violation requests for which CAPTCHA challenge was sent due to Bot profile.
	 */
	Botviolipreputationcaptchaprofilerate float64 `json:"botviolipreputationcaptchaprofilerate,omitempty"`
	Botviolipreputationresetprofile       int     `json:"botviolipreputationresetprofile,omitempty"`
	/**
	* Number of ip reputation violations reset by the Bot profile.
	 */
	Botviolipreputationresetprofilerate float64 `json:"botviolipreputationresetprofilerate,omitempty"`
	Botviolwhitelistprofile             int     `json:"botviolwhitelistprofile,omitempty"`
	/**
	* Number of white list violations seen by the Bot profile.
	 */
	Botviolwhitelistprofilerate float64 `json:"botviolwhitelistprofilerate,omitempty"`
	Botviolwhitelistlogprofile  int     `json:"botviolwhitelistlogprofile,omitempty"`
	/**
	* Number of white list violations logged by the Bot profile.
	 */
	Botviolwhitelistlogprofilerate float64 `json:"botviolwhitelistlogprofilerate,omitempty"`
	Botviolblacklistprofile        int     `json:"botviolblacklistprofile,omitempty"`
	/**
	* Number of black list violations seen by the Bot profile.
	 */
	Botviolblacklistprofilerate float64 `json:"botviolblacklistprofilerate,omitempty"`
	Botviolblacklistlogprofile  int     `json:"botviolblacklistlogprofile,omitempty"`
	/**
	* Number of black list violations logged by the Bot profile.
	 */
	Botviolblacklistlogprofilerate float64 `json:"botviolblacklistlogprofilerate,omitempty"`
	Botviolblacklistdropprofile    int     `json:"botviolblacklistdropprofile,omitempty"`
	/**
	* Number of black list violations dropped by the Bot profile.
	 */
	Botviolblacklistdropprofilerate float64 `json:"botviolblacklistdropprofilerate,omitempty"`
	Botviolblacklistresetprofile    int     `json:"botviolblacklistresetprofile,omitempty"`
	/**
	* Number of black list violations reset by the Bot profile.
	 */
	Botviolblacklistresetprofilerate float64 `json:"botviolblacklistresetprofilerate,omitempty"`
	Botviolblacklistredirectprofile  int     `json:"botviolblacklistredirectprofile,omitempty"`
	/**
	* Number of black list violations redirected by the Bot profile to a different Web page or web server.
	 */
	Botviolblacklistredirectprofilerate float64 `json:"botviolblacklistredirectprofilerate,omitempty"`
	Botviolratelimitprofile             int     `json:"botviolratelimitprofile,omitempty"`
	/**
	* Number of rate limiting violations seen by the Bot profile.
	 */
	Botviolratelimitprofilerate float64 `json:"botviolratelimitprofilerate,omitempty"`
	Botviolratelimitlogprofile  int     `json:"botviolratelimitlogprofile,omitempty"`
	/**
	* Number of rate limiting violations logged by the Bot profile.
	 */
	Botviolratelimitlogprofilerate float64 `json:"botviolratelimitlogprofilerate,omitempty"`
	Botviolratelimitdropprofile    int     `json:"botviolratelimitdropprofile,omitempty"`
	/**
	* Number of rate limiting violations dropped by the Bot profile.
	 */
	Botviolratelimitdropprofilerate float64 `json:"botviolratelimitdropprofilerate,omitempty"`
	Botviolratelimitredirectprofile int     `json:"botviolratelimitredirectprofile,omitempty"`
	/**
	* Number of rate limiting violations requests redirected by the Bot profile to a different Web page or web server.
	 */
	Botviolratelimitredirectprofilerate float64 `json:"botviolratelimitredirectprofilerate,omitempty"`
	Botviolratelimitresetprofile        int     `json:"botviolratelimitresetprofile,omitempty"`
	/**
	* Number of rate limiting violations reset by the Bot profile.
	 */
	Botviolratelimitresetprofilerate float64 `json:"botviolratelimitresetprofilerate,omitempty"`
	Botviolstaticsignatureprofile    int     `json:"botviolstaticsignatureprofile,omitempty"`
	/**
	* Number of static signatutre violations seen by the Bot profile.
	 */
	Botviolstaticsignatureprofilerate float64 `json:"botviolstaticsignatureprofilerate,omitempty"`
	Botviolstaticsignaturelogprofile  int     `json:"botviolstaticsignaturelogprofile,omitempty"`
	/**
	* Number of static signatutre violations logged by the Bot profile.
	 */
	Botviolstaticsignaturelogprofilerate float64 `json:"botviolstaticsignaturelogprofilerate,omitempty"`
	Botviolstaticsignaturedropprofile    int     `json:"botviolstaticsignaturedropprofile,omitempty"`
	/**
	* Number of static signatutre violations dropped by the Bot profile.
	 */
	Botviolstaticsignaturedropprofilerate float64 `json:"botviolstaticsignaturedropprofilerate,omitempty"`
	Botviolstaticsignatureredirectprofile int     `json:"botviolstaticsignatureredirectprofile,omitempty"`
	/**
	* Number of static signatutre violations redirected by the Bot profile to a different Web page or web server.
	 */
	Botviolstaticsignatureredirectprofilerate float64 `json:"botviolstaticsignatureredirectprofilerate,omitempty"`
	Botviolstaticsignatureresetprofile        int     `json:"botviolstaticsignatureresetprofile,omitempty"`
	/**
	* Number of static signatutre violations reset by the Bot profile to a different Web page or web server.
	 */
	Botviolstaticsignatureresetprofilerate float64 `json:"botviolstaticsignatureresetprofilerate,omitempty"`
	Botvioltpsprofile                      int     `json:"botvioltpsprofile,omitempty"`
	/**
	* Number of tps violations seen by the Bot profile.
	 */
	Botvioltpsprofilerate float64 `json:"botvioltpsprofilerate,omitempty"`
	Botvioltpslogprofile  int     `json:"botvioltpslogprofile,omitempty"`
	/**
	* Number of tps violations logged by the Bot profile.
	 */
	Botvioltpslogprofilerate float64 `json:"botvioltpslogprofilerate,omitempty"`
	Botvioltpsdropprofile    int     `json:"botvioltpsdropprofile,omitempty"`
	/**
	* Number of tps violations dropped by the Bot profile.
	 */
	Botvioltpsdropprofilerate float64 `json:"botvioltpsdropprofilerate,omitempty"`
	Botvioltpsredirectprofile int     `json:"botvioltpsredirectprofile,omitempty"`
	/**
	* Number of tps violations requests redirected by the Bot profile to a different Web page or web server.
	 */
	Botvioltpsredirectprofilerate float64 `json:"botvioltpsredirectprofilerate,omitempty"`
	Botvioltpsresetprofile        int     `json:"botvioltpsresetprofile,omitempty"`
	/**
	* Number of tps violations reset by the Bot profile.
	 */
	Botvioltpsresetprofilerate float64 `json:"botvioltpsresetprofilerate,omitempty"`
	Botvioltpscaptchaprofile   int     `json:"botvioltpscaptchaprofile,omitempty"`
	/**
	* Number of tps violation requests for which CAPTCHA challenge was sent due to Bot profile.
	 */
	Botvioltpscaptchaprofilerate float64 `json:"botvioltpscaptchaprofilerate,omitempty"`
	Botviolcaptchaprofile        int     `json:"botviolcaptchaprofile,omitempty"`
	/**
	* Number of Captcha challenge failures seen by the Bot profile.
	 */
	Botviolcaptchaprofilerate float64 `json:"botviolcaptchaprofilerate,omitempty"`
	Botviolcaptchalogprofile  int     `json:"botviolcaptchalogprofile,omitempty"`
	/**
	* Number of Captcha challenge failures logged by the Bot profile.
	 */
	Botviolcaptchalogprofilerate float64 `json:"botviolcaptchalogprofilerate,omitempty"`
	Botviolcaptchadropprofile    int     `json:"botviolcaptchadropprofile,omitempty"`
	/**
	* Number of Captcha challenge failures dropped by the Bot profile.
	 */
	Botviolcaptchadropprofilerate float64 `json:"botviolcaptchadropprofilerate,omitempty"`
	Botviolcaptcharedirectprofile int     `json:"botviolcaptcharedirectprofile,omitempty"`
	/**
	* Number of Captcha challenge failures redirected by the Bot profile.
	 */
	Botviolcaptcharedirectprofilerate float64 `json:"botviolcaptcharedirectprofilerate,omitempty"`
	Botviolcaptcharesetprofile        int     `json:"botviolcaptcharesetprofile,omitempty"`
	/**
	* Number of Captcha challenge failures reset by the Bot profile.
	 */
	Botviolcaptcharesetprofilerate float64 `json:"botviolcaptcharesetprofilerate,omitempty"`
	Botvioltrapprofile             int     `json:"botvioltrapprofile,omitempty"`
	/**
	* Number of trap violations seen by the Bot profile.
	 */
	Botvioltrapprofilerate float64 `json:"botvioltrapprofilerate,omitempty"`
	Botvioltraplogprofile  int     `json:"botvioltraplogprofile,omitempty"`
	/**
	* Number of trap violations logged by the Bot profile.
	 */
	Botvioltraplogprofilerate float64 `json:"botvioltraplogprofilerate,omitempty"`
	Botvioltrapdropprofile    int     `json:"botvioltrapdropprofile,omitempty"`
	/**
	* Number of trap violations dropped by the Bot profile.
	 */
	Botvioltrapdropprofilerate float64 `json:"botvioltrapdropprofilerate,omitempty"`
	Botvioltrapredirectprofile int     `json:"botvioltrapredirectprofile,omitempty"`
	/**
	* Number of trap violations requests redirected by the Bot profile to a different Web page or web server.
	 */
	Botvioltrapredirectprofilerate float64 `json:"botvioltrapredirectprofilerate,omitempty"`
	Botvioltrapresetprofile        int     `json:"botvioltrapresetprofile,omitempty"`
	/**
	* Number of trap violations reset by the Bot profile.
	 */
	Botvioltrapresetprofilerate float64 `json:"botvioltrapresetprofilerate,omitempty"`
}
