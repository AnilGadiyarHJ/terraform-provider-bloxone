/*
BloxOne Redirect API

You can configure BloxOne Threat Defense Cloud to redirect traffic to the Infoblox server that displays the default or customized redirect page. You can redirect traffic to a custom destination using custom redirects.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package redirect

import (
	"github.com/infobloxopen/bloxone-go-client/internal"
	"github.com/infobloxopen/bloxone-go-client/option"
)

const serviceBasePath = "/api/atcfw/v1"

// APIClient manages communication with the BloxOne Redirect API v1
// In most cases there should be only one, shared, APIClient.
type APIClient struct {
	*internal.APIClient

	// API Services
	CertificateAPI     CertificateAPI
	CustomRedirectsAPI CustomRedirectsAPI
	RedirectPageAPI    RedirectPageAPI
}

// NewAPIClient creates a new API client.
// The client can be configured with a variadic option. The following options are available:
// - WithClientName(string) sets the name of the client using the SDK.
// - WithCSPUrl(string) sets the URL for BloxOne Cloud Services Portal.
// - WithAPIKey(string) sets the APIKey for accessing the BloxOne API.
// - WithHTTPClient(*http.Client) sets the HTTPClient to use for the SDK.
// - WithDefaultTags(map[string]string) sets the tags the client can set by default for objects that has tags support.
// - WithDebug() sets the debug mode.
func NewAPIClient(options ...option.ClientOption) *APIClient {
	cfg := internal.NewConfiguration()
	for _, o := range options {
		o(cfg)
	}

	c := &APIClient{}
	c.APIClient = internal.NewAPIClient(serviceBasePath, cfg)

	// API Services
	c.CertificateAPI = (*CertificateAPIService)(&c.Common)
	c.CustomRedirectsAPI = (*CustomRedirectsAPIService)(&c.Common)
	c.RedirectPageAPI = (*RedirectPageAPIService)(&c.Common)

	return c
}
