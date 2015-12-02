package invoices

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/strfmt"
	"github.com/go-swagger/go-swagger/swag"
)

/*
GetInvoiceAsPDFParams contains all the parameters to send to the API endpoint
for the get invoice as p d f operation typically these are written to a http.Request
*/
type GetInvoiceAsPDFParams struct {
	/*
	  The ID of the invoice.
	*/
	ID string
	/*
	  Whether to provide a breakdown of charge tiering.
	*/
	TierBreakdown bool
}

// WriteToRequest writes these params to a swagger request
func (o *GetInvoiceAsPDFParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	// path param ID
	if err := r.SetPathParam("ID", o.ID); err != nil {
		return err
	}

	// query param tier_breakdown
	if err := r.SetQueryParam("tier_breakdown", swag.FormatBool(o.TierBreakdown)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
