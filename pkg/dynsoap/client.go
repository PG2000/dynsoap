package dynsoap

import (
	"github.com/hooklift/gowsdl/soap"
	"net/http"
	"time"
)

// Returns a Dynect Client
// If no http.client is given. A default one will be used
// The default settings for the http.client has a timeout of
// 10 seconds and reads proxy variables from environment http.ProxyFromEnvironment
func NewCustomDynectClient(url string, client *http.Client) Dynect {
	if client == nil {
		client = &http.Client{
			Timeout: time.Second * 10,
			Transport: &http.Transport{
				Proxy: http.ProxyFromEnvironment,
			},
		}
	}
	soapClient := soap.NewClient(url, soap.WithHTTPClient(client))
	return NewDynect(soapClient)
}
