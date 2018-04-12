// Automatically generated by Jenny. DO NOT EDIT!

// Package v1 as generated by Jenny
// Please read about it https://localhost:8080/_spec
package v1

import (
	"context"

	"github.com/Typeform/jenny/options"
	"github.com/go-kit/kit/endpoint"
)

// Net as generated by Jenny
// Please read more at https://localhost:8080/_spec
type Net interface {
	Whois(ctx context.Context, Name string) (*Domain, error)
}

// Domain is generated from a swagger definition
type Domain struct {
	Available bool   `json:"available"` // Available is generated from a swagger definition
	Name      string `json:"name"`      // Name is generated from a swagger definition
	Whois     string `json:"whois"`     // Whois is generated from a swagger definition
}

// _whoisRequest is not to be used outside of this file.
// see https://gokit.io/examples/stringsvc.html#requests-and-responses for more detail
type _whoisRequest struct {
	Name string `json:"name"` // the domain name

}

// _whoisResponse is not to be used outside of this file.
// see https://gokit.io/examples/stringsvc.html#requests-and-responses for more detail
type _whoisResponse struct {
	Body *Domain `json:"body,omitempty"` // Body is generated from a swagger definition

}

// endpoints as used in https://gokit.io/examples/stringsvc.html#endpoints
func makeWhoisEndpoint(svc Net, opts *options.Options) endpoint.Endpoint {
	whoisEndpoint := func(ctx context.Context, request interface{}) (interface{}, error) {

		req := request.(_whoisRequest)

		resp := _whoisResponse{}
		var err error

		resp.Body, err = svc.Whois(ctx, req.Name)

		return resp, err
	}

	whoisMiddleware := opts.OpMiddlewares("Whois")

	return whoisMiddleware(whoisEndpoint)
}