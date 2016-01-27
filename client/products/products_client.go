package products

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/strfmt"
)

// New creates a new products API client.
func New(transport client.Transport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for products API
*/
type Client struct {
	transport client.Transport
	formats   strfmt.Registry
}

/*
DeleteMetadataForProduct removes any associated metadata

{"nickname":"Clear metadata from product","request" :"deleteProductMetadataRequest.html","response":"deleteProductMetadataResponse.html"}
*/
func (a *Client) DeleteMetadataForProduct(params *DeleteMetadataForProductParams) (*DeleteMetadataForProductOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteMetadataForProductParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:                 "deleteMetadataForProduct",
		Method:             "DELETE",
		PathPattern:        "/products/{product-ID}/metadata",
		ProducesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteMetadataForProductReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteMetadataForProductOK), nil
}

/*
GetAllProducts returns a collection of products by default 10 values are returned records are returned in natural order

{"nickname":"Get all products","response":"getProductAll.html"}
*/
func (a *Client) GetAllProducts(params *GetAllProductsParams) (*GetAllProductsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAllProductsParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:                 "getAllProducts",
		Method:             "GET",
		PathPattern:        "/products",
		ProducesMediaTypes: []string{"application/json", "text/plain"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAllProductsReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetAllProductsOK), nil
}

/*
GetMetadataForProduct retrieves any associated metadata

{"nickname":"Retrieve metadata on product","request":"getProductMetadataRequest.html","response":"getProductMetadataResponse.html"}
*/
func (a *Client) GetMetadataForProduct(params *GetMetadataForProductParams) (*GetMetadataForProductOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetMetadataForProductParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:                 "getMetadataForProduct",
		Method:             "GET",
		PathPattern:        "/products/{product-ID}/metadata",
		ProducesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetMetadataForProductReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetMetadataForProductOK), nil
}

/*
SetMetadataForProduct removes any existing metadata keys and create the provided data

{"nickname":"Set metadata on product","request":"setProductMetadataRequest.html","response":"setProductMetadataResponse.html"}
*/
func (a *Client) SetMetadataForProduct(params *SetMetadataForProductParams) (*SetMetadataForProductOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSetMetadataForProductParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:                 "setMetadataForProduct",
		Method:             "POST",
		PathPattern:        "/products/{product-ID}/metadata",
		ProducesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SetMetadataForProductReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*SetMetadataForProductOK), nil
}

/*
UpsertMetadataForProduct updates any existing metadata key values and insert any new key values no keys will be removed

{"nickname":"Upsert metadata on product","request":"upsertProductMetadataRequest.html","response":"upsertProductMetadataResponse.html"}
*/
func (a *Client) UpsertMetadataForProduct(params *UpsertMetadataForProductParams) (*UpsertMetadataForProductOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpsertMetadataForProductParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:                 "upsertMetadataForProduct",
		Method:             "PUT",
		PathPattern:        "/products/{product-ID}/metadata",
		ProducesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpsertMetadataForProductReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpsertMetadataForProductOK), nil
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
