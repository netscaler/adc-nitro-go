# Go client for configuring Citrix ADC

## About

The Citrix®ADC® NITRO for GO allows you to configure and monitor the NetScaler appliance programmatically in GO based applications.

## Installation
- From github

```
go get github.com/citrix/adc-nitro-go
```
- When downloaded from Citrix ADC GUI

Once you have a copy of the source archive unpacked into a similarly-named directory, run following command:
```
./install.sh
```
`install.sh` will copy the lib to `$GOPATH/src/github.com/citrix/adc-nitro-go`.

## Usage
Import the SDK from github.com/citrix/adc-nitro-go/service. Config objects are available at github.com/citrix/adc-nitro-go/resource/config.
Instantiate a client using `NewNitroClient`. To initialize the client from environment variables:

```
export NS_URL=http://<ip-address>
export NS_LOGIN=<adc-username>
export NS_PASSWORD=<adc-password>
```

Config object types can be passed in as strings ("lbvserver"), or looked up from `service.<config object type>.Type()`.
The general pattern for NetScaler config objects is some combination of  `AddResource`, `UpdateResource`, `BindResource`, `UnbindResource` and `DeleteResource`.

Logging level can be set using env variable `NS_LOG` or calling function `client.SetLogLevel("debug")`.
## Example

```
package main

import (
        "github.com/citrix/adc-nitro-go/resource/config/lb"
        "github.com/citrix/adc-nitro-go/service"
)

func main() {
        client, _ := service.NewNitroClientFromEnv()
        lb1 := lb.Lbvserver{
                Name:        "sample_lb",
                Ipv46:       "10.71.136.50",
                Lbmethod:    "ROUNDROBIN",
                Servicetype: "HTTP",
                Port:        8000,
        }
        _, err := client.AddResource(service.Lbvserver.Type(), "sample_lb", &lb1)
        if err == nil {
            client.SaveConfig()
        }
}

```

### Unit Tests
The unit tests are invoked with `make unit`. Note that they are actually functional tests and need a running NetScaler. The tests also need the environment variables `NS_URL`, `NS_LOGIN` and `NS_PASSWORD` to be set. Additionaly `NS_LOG` can be set for setting log level and `ADC_PLATFORM` to `CPX` for skipping some tests which don't work for CPX.

### Using HTTPS
If you specify `https` in the URL then the client will use HTTPS. By default it will verify the presented certificate. If you want to use the default or self-signed certificates without verification, specify `sslVerify=false` in the constructor `NewNitroClientFromParams` or set the environment variable `NS_SSLVERIFY` to `false` and use the `NewNitroClientFromEnv` constructor
