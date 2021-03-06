package addresses

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/strfmt"
)

// New creates a new addresses API client.
func New(transport client.Transport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for addresses API
*/
type Client struct {
	transport client.Transport
	formats   strfmt.Registry
}

/*Create

{"nickname":"Create a new address","response":"createAddressResponse.html","request":"createAddressRequest.html"}
*/
func (a *Client) CreateAddress(params *CreateAddressParams) (*CreateAddressOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateAddressParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:          "createAddress",
		Method:      "POST",
		PathPattern: "/addresses",
		Schemes:     []string{"https"},
		Params:      params,
		Reader:      &CreateAddressReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateAddressOK), nil
}

/*Update

{"nickname":"Update an address","response":"updateAddressResponse.html","request":"updateAddressRequest.html"}
*/
func (a *Client) UpdateAddress(params *UpdateAddressParams) (*UpdateAddressOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateAddressParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:          "updateAddress",
		Method:      "PUT",
		PathPattern: "/addresses",
		Schemes:     []string{"https"},
		Params:      params,
		Reader:      &UpdateAddressReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateAddressOK), nil
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport client.Transport) {
	a.transport = transport
}

// NewAPIError creates a new API error
func NewAPIError(opName string, response interface{}, code int) APIError {
	return APIError{
		OperationName: opName,
		Response:      response,
		Code:          code,
	}
}

// APIError wraps an error model and captures the status code
type APIError struct {
	OperationName string
	Response      interface{}
	Code          int
}

func (a APIError) Error() string {
	return fmt.Sprintf("%s (status %d): %+v ", a.OperationName, a.Code, a.Response)
}
