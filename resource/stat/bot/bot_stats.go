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


type Botstats struct {
	/**
	* Clear the statsistics / counters
	*/
	Clearstats string `json:"clearstats,omitempty"`
	Botrequests uint64 `json:"botrequests,omitempty"`
	/**
	* HTTP/HTTPS requests sent to your protected web servers via the Bot Management.
	*/
	Botrequestsrate int32 `json:"botrequestsrate,omitempty"`
	Botreqbytes uint64 `json:"botreqbytes,omitempty"`
	/**
	* Number of bytes transfered for requests
	*/
	Botreqbytesrate int32 `json:"botreqbytesrate,omitempty"`
	Botresponses uint64 `json:"botresponses,omitempty"`
	/**
	* HTTP/HTTPS responses sent by your protected web servers via the Bot Management.
	*/
	Botresponsesrate int32 `json:"botresponsesrate,omitempty"`
	Botresbytes uint64 `json:"botresbytes,omitempty"`
	/**
	* Number of bytes transfered for responses
	*/
	Botresbytesrate int32 `json:"botresbytesrate,omitempty"`
	Bottotallog uint64 `json:"bottotallog,omitempty"`
	/**
	* Total number of logs by the bot management.
	*/
	Botlograte int32 `json:"botlograte,omitempty"`
	Bottotaldrop uint64 `json:"bottotaldrop,omitempty"`
	/**
	* Total number of drops by the bot management.
	*/
	Botdroprate int32 `json:"botdroprate,omitempty"`
	Bottotalredirect uint64 `json:"bottotalredirect,omitempty"`
	/**
	* Total number of redirects by the bot management.
	*/
	Botredirectrate int32 `json:"botredirectrate,omitempty"`
	Bottotalreset uint64 `json:"bottotalreset,omitempty"`
	/**
	* Total number of resets by the bot management.
	*/
	Botresetrate int32 `json:"botresetrate,omitempty"`
	Botvioldevicefingerprint uint64 `json:"botvioldevicefingerprint,omitempty"`
	/**
	* Number of device fingerprint violations seen by the Bot Management.
	*/
	Botvioldevicefingerprintrate int32 `json:"botvioldevicefingerprintrate,omitempty"`
	Botvioldevicefingerprintlog uint64 `json:"botvioldevicefingerprintlog,omitempty"`
	/**
	* Number of device fingerprint violations logged by the Bot Management.
	*/
	Botvioldevicefingerprintlograte int32 `json:"botvioldevicefingerprintlograte,omitempty"`
	Botvioldevicefingerprintdrop uint64 `json:"botvioldevicefingerprintdrop,omitempty"`
	/**
	* Number of device fingerprint violations dropped by the Bot Management.
	*/
	Botvioldevicefingerprintdroprate int32 `json:"botvioldevicefingerprintdroprate,omitempty"`
	Botvioldevicefingerprintredirect uint64 `json:"botvioldevicefingerprintredirect,omitempty"`
	/**
	* Number of device fingerprint violations requests redirected by the Bot Management to a different Web page or web server.
	*/
	Botvioldevicefingerprintredirectrate int32 `json:"botvioldevicefingerprintredirectrate,omitempty"`
	Botvioldevicefingerprintcaptcha uint64 `json:"botvioldevicefingerprintcaptcha,omitempty"`
	/**
	* Number of device fingerprint violations requests for which CAPTCHA challenge was sent.
	*/
	Botvioldevicefingerprintcaptcharate int32 `json:"botvioldevicefingerprintcaptcharate,omitempty"`
	Botvioldevicefingerprintreset uint64 `json:"botvioldevicefingerprintreset,omitempty"`
	/**
	* Number of device fingerprint violations reset by the Bot Management.
	*/
	Botvioldevicefingerprintresetrate int32 `json:"botvioldevicefingerprintresetrate,omitempty"`
	Botviolipreputation uint64 `json:"botviolipreputation,omitempty"`
	/**
	* Number of ip reputation violations seen by the Bot Management.
	*/
	Botviolipreputationrate int32 `json:"botviolipreputationrate,omitempty"`
	Botviolipreputationlog uint64 `json:"botviolipreputationlog,omitempty"`
	/**
	* Number of ip reputation violations logged by the Bot Management.
	*/
	Botviolipreputationlograte int32 `json:"botviolipreputationlograte,omitempty"`
	Botviolipreputationdrop uint64 `json:"botviolipreputationdrop,omitempty"`
	/**
	* Number of ip reputation violations dropped by the Bot Management.
	*/
	Botviolipreputationdroprate int32 `json:"botviolipreputationdroprate,omitempty"`
	Botviolipreputationredirect uint64 `json:"botviolipreputationredirect,omitempty"`
	/**
	* Number of ip reputation violations requests redirected by the Bot Management to a different Web page or web server.
	*/
	Botviolipreputationredirectrate int32 `json:"botviolipreputationredirectrate,omitempty"`
	Botviolipreputationcaptcha uint64 `json:"botviolipreputationcaptcha,omitempty"`
	/**
	* Number of ip reputation violations requests for which CAPTCHA challenge was sent.
	*/
	Botviolipreputationcaptcharate int32 `json:"botviolipreputationcaptcharate,omitempty"`
	Botviolipreputationreset uint64 `json:"botviolipreputationreset,omitempty"`
	/**
	* Number of ip reputation violations reset by the Bot Management.
	*/
	Botviolipreputationresetrate int32 `json:"botviolipreputationresetrate,omitempty"`
	Botviolwhitelist uint64 `json:"botviolwhitelist,omitempty"`
	/**
	* Number of white list violations seen by the Bot Management.
	*/
	Botviolwhitelistrate int32 `json:"botviolwhitelistrate,omitempty"`
	Botviolwhitelistlog uint64 `json:"botviolwhitelistlog,omitempty"`
	/**
	* Number of white list violations logged by the Bot Management.
	*/
	Botviolwhitelistlograte int32 `json:"botviolwhitelistlograte,omitempty"`
	Botviolblacklist uint64 `json:"botviolblacklist,omitempty"`
	/**
	* Number of black list violations seen by the Bot Management.
	*/
	Botviolblacklistrate int32 `json:"botviolblacklistrate,omitempty"`
	Botviolblacklistlog uint64 `json:"botviolblacklistlog,omitempty"`
	/**
	* Number of black list violations logged by the Bot Management.
	*/
	Botviolblacklistlograte int32 `json:"botviolblacklistlograte,omitempty"`
	Botviolblacklistdrop uint64 `json:"botviolblacklistdrop,omitempty"`
	/**
	* Number of black list violations dropped by the Bot Management.
	*/
	Botviolblacklistdroprate int32 `json:"botviolblacklistdroprate,omitempty"`
	Botviolblacklistreset uint64 `json:"botviolblacklistreset,omitempty"`
	/**
	* Number of black list violations reset by the Bot Management.
	*/
	Botviolblacklistresetrate int32 `json:"botviolblacklistresetrate,omitempty"`
	Botviolblacklistredirect uint64 `json:"botviolblacklistredirect,omitempty"`
	/**
	* Number of black list violations redirected by the Bot Management to a different Web page or web server.
	*/
	Botviolblacklistredirectrate int32 `json:"botviolblacklistredirectrate,omitempty"`
	Botviolratelimit uint64 `json:"botviolratelimit,omitempty"`
	/**
	* Number of rate limiting violations seen by the Bot Management.
	*/
	Botviolratelimitrate int32 `json:"botviolratelimitrate,omitempty"`
	Botviolratelimitlog uint64 `json:"botviolratelimitlog,omitempty"`
	/**
	* Number of rate limiting violations logged by the Bot Management.
	*/
	Botviolratelimitlograte int32 `json:"botviolratelimitlograte,omitempty"`
	Botviolratelimitdrop uint64 `json:"botviolratelimitdrop,omitempty"`
	/**
	* Number of rate limiting violations dropped by the Bot Management.
	*/
	Botviolratelimitdroprate int32 `json:"botviolratelimitdroprate,omitempty"`
	Botviolratelimitredirect uint64 `json:"botviolratelimitredirect,omitempty"`
	/**
	* Number of rate limiting violations requests redirected by the Bot Management to a different Web page or web server.
	*/
	Botviolratelimitredirectrate int32 `json:"botviolratelimitredirectrate,omitempty"`
	Botviolratelimitreset uint64 `json:"botviolratelimitreset,omitempty"`
	/**
	* Number of rate limiting violations reset by the Bot Management.
	*/
	Botviolratelimitresetrate int32 `json:"botviolratelimitresetrate,omitempty"`
	Botviolstaticsignature uint64 `json:"botviolstaticsignature,omitempty"`
	/**
	* Number of static signature violations seen by the Bot Management.
	*/
	Botviolstaticsignaturerate int32 `json:"botviolstaticsignaturerate,omitempty"`
	Botviolstaticsignaturelog uint64 `json:"botviolstaticsignaturelog,omitempty"`
	/**
	* Number of static signature violations logged by the Bot Management.
	*/
	Botviolstaticsignaturelograte int32 `json:"botviolstaticsignaturelograte,omitempty"`
	Botviolstaticsignaturedrop uint64 `json:"botviolstaticsignaturedrop,omitempty"`
	/**
	* Number of static signature violations dropped by the Bot Management.
	*/
	Botviolstaticsignaturedroprate int32 `json:"botviolstaticsignaturedroprate,omitempty"`
	Botviolstaticsignatureredirect uint64 `json:"botviolstaticsignatureredirect,omitempty"`
	/**
	* Number of static signature violations requests redirected by the Bot Management to a different Web page or web server.
	*/
	Botviolstaticsignatureredirectrate int32 `json:"botviolstaticsignatureredirectrate,omitempty"`
	Botviolstaticsignaturereset uint64 `json:"botviolstaticsignaturereset,omitempty"`
	/**
	* Number of static signature violations requests reset by the Bot Management to a different Web page or web server.
	*/
	Botviolstaticsignatureresetrate int32 `json:"botviolstaticsignatureresetrate,omitempty"`
	Botvioltps uint64 `json:"botvioltps,omitempty"`
	/**
	* Number of tps violations seen by the Bot Management.
	*/
	Botvioltpsrate int32 `json:"botvioltpsrate,omitempty"`
	Botvioltpslog uint64 `json:"botvioltpslog,omitempty"`
	/**
	* Number of tps violations logged by the Bot Management.
	*/
	Botvioltpslograte int32 `json:"botvioltpslograte,omitempty"`
	Botvioltpsdrop uint64 `json:"botvioltpsdrop,omitempty"`
	/**
	* Number of tps violations dropped by the Bot Management.
	*/
	Botvioltpsdroprate int32 `json:"botvioltpsdroprate,omitempty"`
	Botvioltpsredirect uint64 `json:"botvioltpsredirect,omitempty"`
	/**
	* Number of tps violations requests redirected by the Bot Management to a different Web page or web server.
	*/
	Botvioltpsredirectrate int32 `json:"botvioltpsredirectrate,omitempty"`
	Botvioltpsreset uint64 `json:"botvioltpsreset,omitempty"`
	/**
	* Number of tps violations reset by the Bot Management.
	*/
	Botvioltpsresetrate int32 `json:"botvioltpsresetrate,omitempty"`
	Botvioltpscaptcha uint64 `json:"botvioltpscaptcha,omitempty"`
	/**
	* Number of TPS violations requests for which CAPTCHA challenge was sent.
	*/
	Botvioltpscaptcharate int32 `json:"botvioltpscaptcharate,omitempty"`
	Botviolcaptcha uint64 `json:"botviolcaptcha,omitempty"`
	/**
	* Number of Captcha challenge failures seen by the Bot Management.
	*/
	Botviolcaptcharate int32 `json:"botviolcaptcharate,omitempty"`
	Botviolcaptchalog uint64 `json:"botviolcaptchalog,omitempty"`
	/**
	* Number of Captcha challenge failures logged by the Bot Management.
	*/
	Botviolcaptchalograte int32 `json:"botviolcaptchalograte,omitempty"`
	Botviolcaptchadrop uint64 `json:"botviolcaptchadrop,omitempty"`
	/**
	* Number of Captcha challenge failures dropped by the Bot Management.
	*/
	Botviolcaptchadroprate int32 `json:"botviolcaptchadroprate,omitempty"`
	Botviolcaptcharedirect uint64 `json:"botviolcaptcharedirect,omitempty"`
	/**
	* Number of Captcha challenge failures redirected by the Bot Management.
	*/
	Botviolcaptcharedirectrate int32 `json:"botviolcaptcharedirectrate,omitempty"`
	Botviolcaptchareset uint64 `json:"botviolcaptchareset,omitempty"`
	/**
	* Number of Captcha challenge failures reset by the Bot Management.
	*/
	Botviolcaptcharesetrate int32 `json:"botviolcaptcharesetrate,omitempty"`
	Botvioltrap uint64 `json:"botvioltrap,omitempty"`
	/**
	* Number of trap violations seen by the Bot Management.
	*/
	Botvioltraprate int32 `json:"botvioltraprate,omitempty"`
	Botvioltraplog uint64 `json:"botvioltraplog,omitempty"`
	/**
	* Number of trap violations logged by the Bot Management.
	*/
	Botvioltraplograte int32 `json:"botvioltraplograte,omitempty"`
	Botvioltrapdrop uint64 `json:"botvioltrapdrop,omitempty"`
	/**
	* Number of trap violations dropped by the Bot Management.
	*/
	Botvioltrapdroprate int32 `json:"botvioltrapdroprate,omitempty"`
	Botvioltrapredirect uint64 `json:"botvioltrapredirect,omitempty"`
	/**
	* Number of trap violations requests redirected by the Bot Management to a different Web page or web server.
	*/
	Botvioltrapredirectrate int32 `json:"botvioltrapredirectrate,omitempty"`
	Botvioltrapreset uint64 `json:"botvioltrapreset,omitempty"`
	/**
	* Number of trap violations reset by the Bot Management.
	*/
	Botvioltrapresetrate int32 `json:"botvioltrapresetrate,omitempty"`

}
