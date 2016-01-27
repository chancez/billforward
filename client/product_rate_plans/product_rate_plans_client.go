package product_rate_plans

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/strfmt"
)

// New creates a new product rate plans API client.
func New(transport client.Transport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for product rate plans API
*/
type Client struct {
	transport client.Transport
	formats   strfmt.Registry
}

/*
DeleteMetadataForRatePlan removes any associated metadata

{"nickname":"Clear metadata from rate-plan","request" :"deleteRatePlanMetadataRequest.html","response":"deleteRatePlanMetadataResponse.html"}
*/
func (a *Client) DeleteMetadataForRatePlan(params *DeleteMetadataForRatePlanParams) (*DeleteMetadataForRatePlanOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteMetadataForRatePlanParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:                 "deleteMetadataForRatePlan",
		Method:             "DELETE",
		PathPattern:        "/product-rate-plans/{product-rate-plan-ID}/metadata",
		ProducesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteMetadataForRatePlanReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteMetadataForRatePlanOK), nil
}

/*
GetAllRatePlans returns a collection of product rate plans by default 10 values are returned records are returned in natural order

{"nickname":"Get all rate-plans","response":"getPRPAll.html"}
*/
func (a *Client) GetAllRatePlans(params *GetAllRatePlansParams) (*GetAllRatePlansOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAllRatePlansParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:                 "getAllRatePlans",
		Method:             "GET",
		PathPattern:        "/product-rate-plans",
		ProducesMediaTypes: []string{"application/json", "text/plain"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAllRatePlansReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetAllRatePlansOK), nil
}

/*
GetMetadataForRatePlan retrieves any associated metadata

{"nickname":"Retrieve metadata on rate-plan","request":"getRatePlanMetadataRequest.html","response":"getRatePlanMetadataResponse.html"}
*/
func (a *Client) GetMetadataForRatePlan(params *GetMetadataForRatePlanParams) (*GetMetadataForRatePlanOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetMetadataForRatePlanParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:                 "getMetadataForRatePlan",
		Method:             "GET",
		PathPattern:        "/product-rate-plans/{product-rate-plan-ID}/metadata",
		ProducesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetMetadataForRatePlanReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetMetadataForRatePlanOK), nil
}

/*
GetProductRatePlanByID returns product rate plans specified by the product rate plan id or name

{"nickname":"Retrieve an existing rate-plan","response":"getPRPByID.html"}
*/
func (a *Client) GetProductRatePlanByID(params *GetProductRatePlanByIDParams) (*GetProductRatePlanByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetProductRatePlanByIDParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:                 "getProductRatePlanByID",
		Method:             "GET",
		PathPattern:        "/product-rate-plans/{product-rate-plan-ID}",
		ProducesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetProductRatePlanByIDReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetProductRatePlanByIDOK), nil
}

/*
GetRatePlanByProduct returns a collection of product rate plans specified by the product ID parameter by default 10 values are returned records are returned in natural order

{"nickname":"Retrieve by product","response":"getPRPByProductID.html"}
*/
func (a *Client) GetRatePlanByProduct(params *GetRatePlanByProductParams) (*GetRatePlanByProductOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRatePlanByProductParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:                 "getRatePlanByProduct",
		Method:             "GET",
		PathPattern:        "/product-rate-plans/product/{product-ID}",
		ProducesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetRatePlanByProductReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetRatePlanByProductOK), nil
}

/*
GetRatePlanByProductAndRatePlan returns a collection of product rate plans specified by the product ID parameter by default 10 values are returned records are returned in natural order

{"nickname":"Retrieve by name","response":"getPRPByNameAndProduct.html"}
*/
func (a *Client) GetRatePlanByProductAndRatePlan(params *GetRatePlanByProductAndRatePlanParams) (*GetRatePlanByProductAndRatePlanOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRatePlanByProductAndRatePlanParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:                 "getRatePlanByProductAndRatePlan",
		Method:             "GET",
		PathPattern:        "/product-rate-plans/product/{product-ID}/rate-plan/{rate-plan-ID}",
		ProducesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetRatePlanByProductAndRatePlanReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetRatePlanByProductAndRatePlanOK), nil
}

/*
SetMetadataForRatePlan removes any existing metadata keys and create the provided data

{"nickname":"Set metadata on rate-plan","request":"setRatePlanMetadataRequest.html","response":"setRatePlanMetadataResponse.html"}
*/
func (a *Client) SetMetadataForRatePlan(params *SetMetadataForRatePlanParams) (*SetMetadataForRatePlanOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSetMetadataForRatePlanParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:                 "setMetadataForRatePlan",
		Method:             "POST",
		PathPattern:        "/product-rate-plans/{product-rate-plan-ID}/metadata",
		ProducesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SetMetadataForRatePlanReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*SetMetadataForRatePlanOK), nil
}

/*
UpsertMetadataForRatePlan updates any existing metadata key values and insert any new key values no keys will be removed

{"nickname":"Upsert metadata on rate-plan","request":"upsertRatePlanMetadataRequest.html","response":"upsertRatePlanMetadataResponse.html"}
*/
func (a *Client) UpsertMetadataForRatePlan(params *UpsertMetadataForRatePlanParams) (*UpsertMetadataForRatePlanOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpsertMetadataForRatePlanParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:                 "upsertMetadataForRatePlan",
		Method:             "PUT",
		PathPattern:        "/product-rate-plans/{product-rate-plan-ID}/metadata",
		ProducesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpsertMetadataForRatePlanReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpsertMetadataForRatePlanOK), nil
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
