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
	Botrequestsrate float64 `json:"botrequestsrate,omitempty"`
	Botreqbytes uint64 `json:"botreqbytes,omitempty"`
	/**
	* Number of bytes transfered for requests
	*/
	Botreqbytesrate float64 `json:"botreqbytesrate,omitempty"`
	Botresponses uint64 `json:"botresponses,omitempty"`
	/**
	* HTTP/HTTPS responses sent by your protected web servers via the Bot Management.
	*/
	Botresponsesrate float64 `json:"botresponsesrate,omitempty"`
	Botresbytes uint64 `json:"botresbytes,omitempty"`
	/**
	* Number of bytes transfered for responses
	*/
	Botresbytesrate float64 `json:"botresbytesrate,omitempty"`
	Bottotallog uint64 `json:"bottotallog,omitempty"`
	/**
	* Total number of logs by the bot management.
	*/
	Botlograte float64 `json:"botlograte,omitempty"`
	Bottotaldrop uint64 `json:"bottotaldrop,omitempty"`
	/**
	* Total number of drops by the bot management.
	*/
	Botdroprate float64 `json:"botdroprate,omitempty"`
	Bottotalredirect uint64 `json:"bottotalredirect,omitempty"`
	/**
	* Total number of redirects by the bot management.
	*/
	Botredirectrate float64 `json:"botredirectrate,omitempty"`
	Bottotalreset uint64 `json:"bottotalreset,omitempty"`
	/**
	* Total number of resets by the bot management.
	*/
	Botresetrate float64 `json:"botresetrate,omitempty"`
	Botvioldevicefingerprint uint64 `json:"botvioldevicefingerprint,omitempty"`
	/**
	* Number of device fingerprint violations seen by the Bot Management.
	*/
	Botvioldevicefingerprintrate float64 `json:"botvioldevicefingerprintrate,omitempty"`
	Botvioldevicefingerprintlog uint64 `json:"botvioldevicefingerprintlog,omitempty"`
	/**
	* Number of device fingerprint violations logged by the Bot Management.
	*/
	Botvioldevicefingerprintlograte float64 `json:"botvioldevicefingerprintlograte,omitempty"`
	Botvioldevicefingerprintdrop uint64 `json:"botvioldevicefingerprintdrop,omitempty"`
	/**
	* Number of device fingerprint violations dropped by the Bot Management.
	*/
	Botvioldevicefingerprintdroprate float64 `json:"botvioldevicefingerprintdroprate,omitempty"`
	Botvioldevicefingerprintredirect uint64 `json:"botvioldevicefingerprintredirect,omitempty"`
	/**
	* Number of device fingerprint violations requests redirected by the Bot Management to a different Web page or web server.
	*/
	Botvioldevicefingerprintredirectrate float64 `json:"botvioldevicefingerprintredirectrate,omitempty"`
	Botvioldevicefingerprintcaptcha uint64 `json:"botvioldevicefingerprintcaptcha,omitempty"`
	/**
	* Number of device fingerprint violations requests for which CAPTCHA challenge was sent.
	*/
	Botvioldevicefingerprintcaptcharate float64 `json:"botvioldevicefingerprintcaptcharate,omitempty"`
	Botvioldevicefingerprintreset uint64 `json:"botvioldevicefingerprintreset,omitempty"`
	/**
	* Number of device fingerprint violations reset by the Bot Management.
	*/
	Botvioldevicefingerprintresetrate float64 `json:"botvioldevicefingerprintresetrate,omitempty"`
	Botviolipreputation uint64 `json:"botviolipreputation,omitempty"`
	/**
	* Number of ip reputation violations seen by the Bot Management.
	*/
	Botviolipreputationrate float64 `json:"botviolipreputationrate,omitempty"`
	Botviolipreputationlog uint64 `json:"botviolipreputationlog,omitempty"`
	/**
	* Number of ip reputation violations logged by the Bot Management.
	*/
	Botviolipreputationlograte float64 `json:"botviolipreputationlograte,omitempty"`
	Botviolipreputationdrop uint64 `json:"botviolipreputationdrop,omitempty"`
	/**
	* Number of ip reputation violations dropped by the Bot Management.
	*/
	Botviolipreputationdroprate float64 `json:"botviolipreputationdroprate,omitempty"`
	Botviolipreputationredirect uint64 `json:"botviolipreputationredirect,omitempty"`
	/**
	* Number of ip reputation violations requests redirected by the Bot Management to a different Web page or web server.
	*/
	Botviolipreputationredirectrate float64 `json:"botviolipreputationredirectrate,omitempty"`
	Botviolipreputationcaptcha uint64 `json:"botviolipreputationcaptcha,omitempty"`
	/**
	* Number of ip reputation violations requests for which CAPTCHA challenge was sent.
	*/
	Botviolipreputationcaptcharate float64 `json:"botviolipreputationcaptcharate,omitempty"`
	Botviolipreputationreset uint64 `json:"botviolipreputationreset,omitempty"`
	/**
	* Number of ip reputation violations reset by the Bot Management.
	*/
	Botviolipreputationresetrate float64 `json:"botviolipreputationresetrate,omitempty"`
	Botviolwhitelist uint64 `json:"botviolwhitelist,omitempty"`
	/**
	* Number of white list violations seen by the Bot Management.
	*/
	Botviolwhitelistrate float64 `json:"botviolwhitelistrate,omitempty"`
	Botviolwhitelistlog uint64 `json:"botviolwhitelistlog,omitempty"`
	/**
	* Number of white list violations logged by the Bot Management.
	*/
	Botviolwhitelistlograte float64 `json:"botviolwhitelistlograte,omitempty"`
	Botviolblacklist uint64 `json:"botviolblacklist,omitempty"`
	/**
	* Number of black list violations seen by the Bot Management.
	*/
	Botviolblacklistrate float64 `json:"botviolblacklistrate,omitempty"`
	Botviolblacklistlog uint64 `json:"botviolblacklistlog,omitempty"`
	/**
	* Number of black list violations logged by the Bot Management.
	*/
	Botviolblacklistlograte float64 `json:"botviolblacklistlograte,omitempty"`
	Botviolblacklistdrop uint64 `json:"botviolblacklistdrop,omitempty"`
	/**
	* Number of black list violations dropped by the Bot Management.
	*/
	Botviolblacklistdroprate float64 `json:"botviolblacklistdroprate,omitempty"`
	Botviolblacklistreset uint64 `json:"botviolblacklistreset,omitempty"`
	/**
	* Number of black list violations reset by the Bot Management.
	*/
	Botviolblacklistresetrate float64 `json:"botviolblacklistresetrate,omitempty"`
	Botviolblacklistredirect uint64 `json:"botviolblacklistredirect,omitempty"`
	/**
	* Number of black list violations redirected by the Bot Management to a different Web page or web server.
	*/
	Botviolblacklistredirectrate float64 `json:"botviolblacklistredirectrate,omitempty"`
	Botviolratelimit uint64 `json:"botviolratelimit,omitempty"`
	/**
	* Number of rate limiting violations seen by the Bot Management.
	*/
	Botviolratelimitrate float64 `json:"botviolratelimitrate,omitempty"`
	Botviolratelimitlog uint64 `json:"botviolratelimitlog,omitempty"`
	/**
	* Number of rate limiting violations logged by the Bot Management.
	*/
	Botviolratelimitlograte float64 `json:"botviolratelimitlograte,omitempty"`
	Botviolratelimitdrop uint64 `json:"botviolratelimitdrop,omitempty"`
	/**
	* Number of rate limiting violations dropped by the Bot Management.
	*/
	Botviolratelimitdroprate float64 `json:"botviolratelimitdroprate,omitempty"`
	Botviolratelimitredirect uint64 `json:"botviolratelimitredirect,omitempty"`
	/**
	* Number of rate limiting violations requests redirected by the Bot Management to a different Web page or web server.
	*/
	Botviolratelimitredirectrate float64 `json:"botviolratelimitredirectrate,omitempty"`
	Botviolratelimitreset uint64 `json:"botviolratelimitreset,omitempty"`
	/**
	* Number of rate limiting violations reset by the Bot Management.
	*/
	Botviolratelimitresetrate float64 `json:"botviolratelimitresetrate,omitempty"`
	Botviolstaticsignature uint64 `json:"botviolstaticsignature,omitempty"`
	/**
	* Number of static signature violations seen by the Bot Management.
	*/
	Botviolstaticsignaturerate float64 `json:"botviolstaticsignaturerate,omitempty"`
	Botviolstaticsignaturelog uint64 `json:"botviolstaticsignaturelog,omitempty"`
	/**
	* Number of static signature violations logged by the Bot Management.
	*/
	Botviolstaticsignaturelograte float64 `json:"botviolstaticsignaturelograte,omitempty"`
	Botviolstaticsignaturedrop uint64 `json:"botviolstaticsignaturedrop,omitempty"`
	/**
	* Number of static signature violations dropped by the Bot Management.
	*/
	Botviolstaticsignaturedroprate float64 `json:"botviolstaticsignaturedroprate,omitempty"`
	Botviolstaticsignatureredirect uint64 `json:"botviolstaticsignatureredirect,omitempty"`
	/**
	* Number of static signature violations requests redirected by the Bot Management to a different Web page or web server.
	*/
	Botviolstaticsignatureredirectrate float64 `json:"botviolstaticsignatureredirectrate,omitempty"`
	Botviolstaticsignaturereset uint64 `json:"botviolstaticsignaturereset,omitempty"`
	/**
	* Number of static signature violations requests reset by the Bot Management to a different Web page or web server.
	*/
	Botviolstaticsignatureresetrate float64 `json:"botviolstaticsignatureresetrate,omitempty"`
	Botvioltps uint64 `json:"botvioltps,omitempty"`
	/**
	* Number of tps violations seen by the Bot Management.
	*/
	Botvioltpsrate float64 `json:"botvioltpsrate,omitempty"`
	Botvioltpslog uint64 `json:"botvioltpslog,omitempty"`
	/**
	* Number of tps violations logged by the Bot Management.
	*/
	Botvioltpslograte float64 `json:"botvioltpslograte,omitempty"`
	Botvioltpsdrop uint64 `json:"botvioltpsdrop,omitempty"`
	/**
	* Number of tps violations dropped by the Bot Management.
	*/
	Botvioltpsdroprate float64 `json:"botvioltpsdroprate,omitempty"`
	Botvioltpsredirect uint64 `json:"botvioltpsredirect,omitempty"`
	/**
	* Number of tps violations requests redirected by the Bot Management to a different Web page or web server.
	*/
	Botvioltpsredirectrate float64 `json:"botvioltpsredirectrate,omitempty"`
	Botvioltpsreset uint64 `json:"botvioltpsreset,omitempty"`
	/**
	* Number of tps violations reset by the Bot Management.
	*/
	Botvioltpsresetrate float64 `json:"botvioltpsresetrate,omitempty"`
	Botvioltpscaptcha uint64 `json:"botvioltpscaptcha,omitempty"`
	/**
	* Number of TPS violations requests for which CAPTCHA challenge was sent.
	*/
	Botvioltpscaptcharate float64 `json:"botvioltpscaptcharate,omitempty"`
	Botviolcaptcha uint64 `json:"botviolcaptcha,omitempty"`
	/**
	* Number of Captcha challenge failures seen by the Bot Management.
	*/
	Botviolcaptcharate float64 `json:"botviolcaptcharate,omitempty"`
	Botviolcaptchalog uint64 `json:"botviolcaptchalog,omitempty"`
	/**
	* Number of Captcha challenge failures logged by the Bot Management.
	*/
	Botviolcaptchalograte float64 `json:"botviolcaptchalograte,omitempty"`
	Botviolcaptchadrop uint64 `json:"botviolcaptchadrop,omitempty"`
	/**
	* Number of Captcha challenge failures dropped by the Bot Management.
	*/
	Botviolcaptchadroprate float64 `json:"botviolcaptchadroprate,omitempty"`
	Botviolcaptcharedirect uint64 `json:"botviolcaptcharedirect,omitempty"`
	/**
	* Number of Captcha challenge failures redirected by the Bot Management.
	*/
	Botviolcaptcharedirectrate float64 `json:"botviolcaptcharedirectrate,omitempty"`
	Botviolcaptchareset uint64 `json:"botviolcaptchareset,omitempty"`
	/**
	* Number of Captcha challenge failures reset by the Bot Management.
	*/
	Botviolcaptcharesetrate float64 `json:"botviolcaptcharesetrate,omitempty"`
	Botvioltrap uint64 `json:"botvioltrap,omitempty"`
	/**
	* Number of trap violations seen by the Bot Management.
	*/
	Botvioltraprate float64 `json:"botvioltraprate,omitempty"`
	Botvioltraplog uint64 `json:"botvioltraplog,omitempty"`
	/**
	* Number of trap violations logged by the Bot Management.
	*/
	Botvioltraplograte float64 `json:"botvioltraplograte,omitempty"`
	Botvioltrapdrop uint64 `json:"botvioltrapdrop,omitempty"`
	/**
	* Number of trap violations dropped by the Bot Management.
	*/
	Botvioltrapdroprate float64 `json:"botvioltrapdroprate,omitempty"`
	Botvioltrapredirect uint64 `json:"botvioltrapredirect,omitempty"`
	/**
	* Number of trap violations requests redirected by the Bot Management to a different Web page or web server.
	*/
	Botvioltrapredirectrate float64 `json:"botvioltrapredirectrate,omitempty"`
	Botvioltrapreset uint64 `json:"botvioltrapreset,omitempty"`
	/**
	* Number of trap violations reset by the Bot Management.
	*/
	Botvioltrapresetrate float64 `json:"botvioltrapresetrate,omitempty"`

}
