package dynsoap

import (
	"github.com/hooklift/gowsdl/soap"
	"net/http"
	"time"
)

// Returns a Dynect Client with a configured http.Client
// The default settings for the http.client has a timeout of
// 10 seconds and reads proxy variables from environment http.ProxyFromEnvironment
func NewDynectClient(url string) Dynect {
	client := &http.Client{
		Timeout: time.Second * 10,
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
		},
	}
	soapClient := soap.NewClient(url, soap.WithHTTPClient(client))
	return NewDynect(soapClient)
}

// Returns a Dynect Client without http.Client
func NewCustomDynectClient(url string, client http.Client) Dynect {
	soapClient := soap.NewClient(url, soap.WithHTTPClient(&client))
	return NewDynect(soapClient)
}
