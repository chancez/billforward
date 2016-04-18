package profiles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new profiles API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for profiles API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetProfile returns a single profile specified by the ID parameter

{"nickname":"Retrieve an existing profile","response":"getProfileByID.html"}
*/
func (a *Client) GetProfile(params *GetProfileParams) (*GetProfileOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetProfileParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getProfile",
		Method:             "GET",
		PathPattern:        "/profiles/{profile-ID}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"text/plain"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetProfileReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetProfileOK), nil
}

/*
UpdateProfile updates a profile

{"nickname":"Update a profile","request":"updateProfileRequest.html","response":"updateProfileResponse.html"}
*/
func (a *Client) UpdateProfile(params *UpdateProfileParams) (*UpdateProfileOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateProfileParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateProfile",
		Method:             "PUT",
		PathPattern:        "/profiles",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateProfileReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateProfileOK), nil
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
