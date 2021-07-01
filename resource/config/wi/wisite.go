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

package wi

/**
* Configuration for WI site resource.
*/
type Wisite struct {
	/**
	* Path to the Web Interface site being created on the Citrix ADC.
	*/
	Sitepath string `json:"sitepath,omitempty"`
	/**
	* Call back URL of the Gateway.
	*/
	Agurl string `json:"agurl,omitempty"`
	/**
	* URL of the Secure Ticket Authority (STA) server.
	*/
	Staurl string `json:"staurl,omitempty"`
	/**
	* URL of the second Secure Ticket Authority (STA) server.
	*/
	Secondstaurl string `json:"secondstaurl,omitempty"`
	/**
	* Enable session reliability through Access Gateway.
	*/
	Sessionreliability string `json:"sessionreliability,omitempty"`
	/**
	* Request tickets issued by two separate Secure Ticket Authorities (STA) when a resource is accessed.
	*/
	Usetwotickets string `json:"usetwotickets,omitempty"`
	/**
	* Authentication point for the Web Interface site.
	*/
	Authenticationpoint string `json:"authenticationpoint,omitempty"`
	/**
	* Method for authenticating a Web Interface site if you have specified Web Interface as the authentication point.
		Available settings function as follows:
		* Explicit - Users must provide a user name and password to log on to the Web Interface.
		* Anonymous - Users can log on to the Web Interface without providing a user name and password. They have access to resources published for anonymous users.
	*/
	Agauthenticationmethod string `json:"agauthenticationmethod,omitempty"`
	/**
	* The method of authentication to be used at Web Interface
	*/
	Wiauthenticationmethods []string `json:"wiauthenticationmethods,omitempty"`
	/**
	* Default language for the Web Interface site.
	*/
	Defaultcustomtextlocale string `json:"defaultcustomtextlocale,omitempty"`
	/**
	* Time-out, in minutes, for idle Web Interface browser sessions. If a client's session is idle for a time that exceeds the time-out value, the Citrix ADC terminates the connection.
	*/
	Websessiontimeout int `json:"websessiontimeout,omitempty"`
	/**
	* Default access method for clients accessing the Web Interface site.
		Note: Before you configure an access method based on the client IP address, you must enable USIP mode on the Web Interface service to make the client's IP address available with the Web Interface.
		Depending on whether the Web Interface site is configured to use an HTTP or HTTPS virtual server or to use access gateway, you can send clients or access gateway the IP address, or the alternate address, of a XenApp or XenDesktop server. Or, you can send the IP address translated from a mapping entry, which defines mapping of an internal address and port to an external address and port.
		Note: In the Citrix ADC command line, mapping entries can be created by using the bind wi site command.
	*/
	Defaultaccessmethod string `json:"defaultaccessmethod,omitempty"`
	/**
	* A custom login page title for the Web Interface site.
	*/
	Logintitle string `json:"logintitle,omitempty"`
	/**
	* Specifies localized text to appear at the top of the main content area of the Applications screen. LanguageCode is en, de, es, fr, ja, or any other supported language identifier.
	*/
	Appwelcomemessage string `json:"appwelcomemessage,omitempty"`
	/**
	* Localized welcome message that appears on the welcome area of the login screen.
	*/
	Welcomemessage string `json:"welcomemessage,omitempty"`
	/**
	* Localized text that appears in the footer area of all pages.
	*/
	Footertext string `json:"footertext,omitempty"`
	/**
	* Localized text that appears at the bottom of the main content area of the login screen.
	*/
	Loginsysmessage string `json:"loginsysmessage,omitempty"`
	/**
	* Localized text that appears as the name of the pre-login message confirmation button.
	*/
	Preloginbutton string `json:"preloginbutton,omitempty"`
	/**
	* Localized text that appears on the pre-login message page.
	*/
	Preloginmessage string `json:"preloginmessage,omitempty"`
	/**
	* Localized text that appears as the title of the pre-login message page.
	*/
	Prelogintitle string `json:"prelogintitle,omitempty"`
	/**
	* Domain names listed on the login screen for explicit authentication.
	*/
	Domainselection string `json:"domainselection,omitempty"`
	/**
	* Type of access to the Web Interface site. Available settings function as follows:
		* XenApp/XenDesktop web site - Configures the Web Interface site for access by a web browser.
		* XenApp/XenDesktop services site - Configures the Web Interface site for access by the XenApp plug-in.
	*/
	Sitetype string `json:"sitetype,omitempty"`
	/**
	* Specifies whether the site is focused towards users accessing applications or desktops. Setting the parameter to Desktops changes the functionality of the site to improve the experience for XenDesktop users. Citrix recommends using this setting for any deployment that includes XenDesktop.
	*/
	Userinterfacebranding string `json:"userinterfacebranding,omitempty"`
	/**
	* Method for accessing the published XenApp and XenDesktop resources. 
		Available settings function as follows:
		* Online - Allows applications to be launched on the XenApp and XenDesktop servers. 
		* Offline - Allows streaming of applications to the client. 
		* DualMode - Allows both online and offline modes.
	*/
	Publishedresourcetype string `json:"publishedresourcetype,omitempty"`
	/**
	* User settings do not persist from one session to another.
	*/
	Kioskmode string `json:"kioskmode,omitempty"`
	/**
	* Enables search option on XenApp websites
	*/
	Showsearch string `json:"showsearch,omitempty"`
	/**
	* Provides the Refresh button on the applications screen.
	*/
	Showrefresh string `json:"showrefresh,omitempty"`
	/**
	* Appearance of the login screen.
		* Simple - Only the login fields for the selected authentication method are displayed.
		* Advanced - Displays the navigation bar, which provides access to the pre-login messages and preferences screens.
	*/
	Wiuserinterfacemodes string `json:"wiuserinterfacemodes,omitempty"`
	/**
	* Specifies whether or not to use the compact user interface.
	*/
	Userinterfacelayouts string `json:"userinterfacelayouts,omitempty"`
	/**
	* The RestrictDomains setting is used to enable/disable domain restrictions. If domain restriction is enabled, the LoginDomains list is used for validating the login domain. It is applied to all the authentication methods except Anonymous for XenApp Web and XenApp Services sites
	*/
	Restrictdomains string `json:"restrictdomains,omitempty"`
	/**
	* [List of NetBIOS domain names], Domain names to use for access restriction.
		Only takes effect when used in conjunction with the RestrictDomains setting.
	*/
	Logindomains string `json:"logindomains,omitempty"`
	/**
	* The HideDomainField setting is used to control whether the domain field is displayed on the logon screen.
	*/
	Hidedomainfield string `json:"hidedomainfield,omitempty"`
	/**
	* Callback AGURL to which Web Interface contacts. 
	*/
	Agcallbackurl string `json:"agcallbackurl,omitempty"`

}
