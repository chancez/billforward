package invoices

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/strfmt"
)

// New creates a new invoices API client.
func New(transport client.Transport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for invoices API
*/
type Client struct {
	transport client.Transport
	formats   strfmt.Registry
}

/*Retrieves a single invoice specified by the ID parameter.

{ "nickname" : "PDF Invoice","response" : "getInvoiceByID.pdf"}
*/
func (a *Client) GetInvoiceAsPDF(params GetInvoiceAsPDFParams) (*GetInvoiceAsPDFOK, error) {
	// TODO: Validate the params before sending

	result, err := a.transport.Submit(&client.Operation{
		ID:     "getInvoiceAsPDF",
		Params: &params,
		Reader: &GetInvoiceAsPDFReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetInvoiceAsPDFOK), nil
}

/*Retrieves a collection of invoices specified by the account-ID parameter. By default 10 values are returned. Records are returned in natural order.

{ "nickname" : "Retrieve by account","response" : "getInvoiceByAccountID.html"}
*/
func (a *Client) GetInvoicesByAccountID(params GetInvoicesByAccountIDParams) (*GetInvoicesByAccountIDOK, error) {
	// TODO: Validate the params before sending

	result, err := a.transport.Submit(&client.Operation{
		ID:     "getInvoicesByAccountID",
		Params: &params,
		Reader: &GetInvoicesByAccountIDReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetInvoicesByAccountIDOK), nil
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
