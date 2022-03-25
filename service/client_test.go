/*
Copyright 2021 Citrix Systems, Inc

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package service

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
	"testing"

	"github.com/citrix/adc-nitro-go/service/internal/testcert"
)

// equals fails the test if exp is not equal to act.
func equals(tb testing.TB, exp, act interface{}) {
	if !reflect.DeepEqual(exp, act) {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d:\n\n\texp: %#v\n\n\tgot: %#v\033[39m\n\n", filepath.Base(file), line, exp, act)
		tb.FailNow()
	}
}

func newTLSTestServer(t *testing.T, handler http.Handler) *httptest.Server {
	// Start a local HTTP server
	server := httptest.NewUnstartedServer(handler)

	serverCert, err := tls.X509KeyPair(testcert.LocalhostCert, testcert.LocalhostKey)
	if err != nil {
		t.Error(err)
	}

	server.TLS = &tls.Config{
		Certificates: []tls.Certificate{serverCert},
	}

	server.StartTLS()

	return server
}

func SaveAndRestore(name string) func() {
	old, ok := os.LookupEnv(name)
	if ok {
		os.Unsetenv(name)
		return func() {
			os.Setenv(name, old)
		}
	}

	return func() {}

}

func TestClientCreate(t *testing.T) {
	t.Log("Create client from environment variables and supplied params")
	defer SaveAndRestore("NS_URL")()
	defer SaveAndRestore("NS_LOGIN")()
	defer SaveAndRestore("NS_PASSWORD")()

	_, err := NewNitroClientFromEnv()
	if err == nil {
		t.Error("Expected to fail in creating client")
	}

	os.Setenv("NS_URL", "http://127.0.0.1:32775")
	_, err = NewNitroClientFromEnv()
	if err == nil {
		t.Error("Expected to fail in creating client")
	}
	os.Setenv("NS_LOGIN", "nsroot")
	_, err = NewNitroClientFromEnv()
	if err == nil {
		t.Error("Expected to fail in creating client")
	}
	os.Setenv("NS_PASSWORD", "nsroot")
	_, err = NewNitroClientFromEnv()
	if err != nil {
		t.Error("Didnt expect to fail in creating client")
	}
	os.Setenv("_MPS_API_PROXY_MANAGED_INSTANCE_IP", "10.221.48.207")
	client, err := NewNitroClientFromEnv()
	if err != nil {
		t.Error("Didnt expect to fail in creating client")
	}
	if client.proxiedNs != "10.221.48.207" {
		t.Error("proxiedNS not set despite being set in the environment")
	}
	os.Setenv("NS_SSLVERIFY", "False")
	_, err = NewNitroClientFromEnv()
	if err != nil {
		t.Error("Didnt expect to fail in creating client with SSL verify option")
	}

}

func TestClientCreateFromParams(t *testing.T) {
	t.Log("Create client from supplied params")

	params := NitroParams{
		Url:      "http://127.0.0.1",
		Username: "user",
		Password: "pass",
	}
	client, err := NewNitroClientFromParams(params)
	if client == nil {
		t.Error("Expected to succeed in creating client")
	}

	params.SslVerify = false
	client, err = NewNitroClientFromParams(params)
	if client == nil {
		t.Error("Expected to succeed in creating client")
	}

	params.Url = "localhost:32770"
	client, err = NewNitroClientFromParams(params)
	if err == nil {
		t.Error("Expected to fail in creating client due to invalid URL")
	}

	nsUrlSplit := strings.Split(os.Getenv("NS_URL"), "://")
	ip := nsUrlSplit[len(nsUrlSplit)-1]
	params.Url = "https://" + ip
	params.SslVerify = true
	params.RootCAPath = os.Getenv("NS_CA_PATH")
	params.ServerName = os.Getenv("NS_SERVER_NAME")
	client, err = NewNitroClientFromParams(params)
	if client == nil {
		t.Error("Expected to succeed in creating client")
	}
}

func Test_newHttpClient(t *testing.T) {
	type args struct {
		params NitroParams
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "NoSslVerify",
			args: args{
				params: NitroParams{
					SslVerify: false,
				},
			},
			wantErr: false,
		},
		{
			name: "SslVerify",
			args: args{
				params: NitroParams{
					SslVerify: true,
				},
			},
			wantErr: true,
		},
		{
			name: "WithRootCA",
			args: args{
				params: NitroParams{
					SslVerify: true,
					RootCAs: [][]byte{
						testcert.LocalcaCert,
					},
				},
			},
			wantErr: false,
		},
		{
			name: "WithInvalidRootCa",
			args: args{
				params: NitroParams{
					SslVerify: true,
					RootCAs: [][]byte{
						testcert.InvalidLocalcaCert,
					},
				},
			},
			wantErr: true,
		},
		{
			name: "WithHostCert",
			args: args{
				params: NitroParams{
					SslVerify: true,
					RootCAs: [][]byte{
						testcert.LocalhostCert,
					},
				},
			},
			wantErr: false,
		},
		{
			name: "WithInvalidHostCert",
			args: args{
				params: NitroParams{
					SslVerify: true,
					RootCAs: [][]byte{
						testcert.InvalidLocalhostCert,
					},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// Start a local HTTP server
			server := newTLSTestServer(t, http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
				// Test request parameters
				equals(t, req.URL.String(), "/some/path")
				// Send response to be tested
				rw.Write([]byte(`OK`))
			}))

			// Close the server when test finishes
			defer server.Close()

			// Get the certificate used

			client, err := newHttpClient(tt.args.params)
			if err != nil {
				t.Error(err)
			}

			_, err = client.Get(server.URL + "/some/path")

			if (err != nil) != tt.wantErr {
				t.Errorf("newHttpClient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

		})
	}
}
