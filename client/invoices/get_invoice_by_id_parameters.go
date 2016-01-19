package invoices

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/strfmt"
	"github.com/go-swagger/go-swagger/swag"
)

// NewGetInvoiceByIDParams creates a new GetInvoiceByIDParams object
// with the default values initialized.
func NewGetInvoiceByIDParams() *GetInvoiceByIDParams {
	var (
		excludeChildrenDefault bool   = bool(true)
		orderDefault           string = string("DESC")
		orderByDefault         string = string("created")
		recordsDefault         int32  = int32(10)
	)
	return &GetInvoiceByIDParams{
		ExcludeChildren: &excludeChildrenDefault,
		Order:           &orderDefault,
		OrderBy:         &orderByDefault,
		Records:         &recordsDefault,
	}
}

/*GetInvoiceByIDParams contains all the parameters to send to the API endpoint
for the get invoice by ID operation typically these are written to a http.Request
*/
type GetInvoiceByIDParams struct {

	/*ExcludeChildren
	  Should child invoices be excluded.

	*/
	ExcludeChildren *bool
	/*IncludeRetired
	  Whether retired products should be returned.

	*/
	IncludeRetired *bool
	/*InvoiceID
	  The ID of the invoice.

	*/
	InvoiceID string
	/*Offset
	  The offset from the first invoice to return.

	*/
	Offset *int32
	/*Order
	  Ihe direction of any ordering, either ASC or DESC.

	*/
	Order *string
	/*OrderBy
	  Specify a field used to order the result set.

	*/
	OrderBy *string
	/*Organizations
	  A list of organization-IDs used to restrict the scope of API calls.

	*/
	Organizations []string
	/*Records
	  The maximum number of invoices to return.

	*/
	Records *int32
}

// WithExcludeChildren adds the excludeChildren to the get invoice by ID params
func (o *GetInvoiceByIDParams) WithExcludeChildren(excludeChildren *bool) *GetInvoiceByIDParams {
	o.ExcludeChildren = excludeChildren
	return o
}

// WithIncludeRetired adds the includeRetired to the get invoice by ID params
func (o *GetInvoiceByIDParams) WithIncludeRetired(includeRetired *bool) *GetInvoiceByIDParams {
	o.IncludeRetired = includeRetired
	return o
}

// WithInvoiceID adds the invoiceId to the get invoice by ID params
func (o *GetInvoiceByIDParams) WithInvoiceID(invoiceId string) *GetInvoiceByIDParams {
	o.InvoiceID = invoiceId
	return o
}

// WithOffset adds the offset to the get invoice by ID params
func (o *GetInvoiceByIDParams) WithOffset(offset *int32) *GetInvoiceByIDParams {
	o.Offset = offset
	return o
}

// WithOrder adds the order to the get invoice by ID params
func (o *GetInvoiceByIDParams) WithOrder(order *string) *GetInvoiceByIDParams {
	o.Order = order
	return o
}

// WithOrderBy adds the orderBy to the get invoice by ID params
func (o *GetInvoiceByIDParams) WithOrderBy(orderBy *string) *GetInvoiceByIDParams {
	o.OrderBy = orderBy
	return o
}

// WithOrganizations adds the organizations to the get invoice by ID params
func (o *GetInvoiceByIDParams) WithOrganizations(organizations []string) *GetInvoiceByIDParams {
	o.Organizations = organizations
	return o
}

// WithRecords adds the records to the get invoice by ID params
func (o *GetInvoiceByIDParams) WithRecords(records *int32) *GetInvoiceByIDParams {
	o.Records = records
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *GetInvoiceByIDParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	if o.ExcludeChildren != nil {

		// query param exclude_children
		var qrExcludeChildren bool
		if o.ExcludeChildren != nil {
			qrExcludeChildren = *o.ExcludeChildren
		}
		qExcludeChildren := swag.FormatBool(qrExcludeChildren)
		if qExcludeChildren != "" {
			if err := r.SetQueryParam("exclude_children", qExcludeChildren); err != nil {
				return err
			}
		}

	}

	if o.IncludeRetired != nil {

		// query param include_retired
		var qrIncludeRetired bool
		if o.IncludeRetired != nil {
			qrIncludeRetired = *o.IncludeRetired
		}
		qIncludeRetired := swag.FormatBool(qrIncludeRetired)
		if qIncludeRetired != "" {
			if err := r.SetQueryParam("include_retired", qIncludeRetired); err != nil {
				return err
			}
		}

	}

	// path param invoice-ID
	if err := r.SetPathParam("invoice-ID", o.InvoiceID); err != nil {
		return err
	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int32
		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt32(qrOffset)
		if qOffset != "" {
			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}

	}

	if o.Order != nil {

		// query param order
		var qrOrder string
		if o.Order != nil {
			qrOrder = *o.Order
		}
		qOrder := qrOrder
		if qOrder != "" {
			if err := r.SetQueryParam("order", qOrder); err != nil {
				return err
			}
		}

	}

	if o.OrderBy != nil {

		// query param order_by
		var qrOrderBy string
		if o.OrderBy != nil {
			qrOrderBy = *o.OrderBy
		}
		qOrderBy := qrOrderBy
		if qOrderBy != "" {
			if err := r.SetQueryParam("order_by", qOrderBy); err != nil {
				return err
			}
		}

	}

	valuesOrganizations := o.Organizations

	joinedOrganizations := swag.JoinByFormat(valuesOrganizations, "multi")
	// query array param organizations
	if err := r.SetQueryParam("organizations", joinedOrganizations...); err != nil {
		return err
	}

	if o.Records != nil {

		// query param records
		var qrRecords int32
		if o.Records != nil {
			qrRecords = *o.Records
		}
		qRecords := swag.FormatInt32(qrRecords)
		if qRecords != "" {
			if err := r.SetQueryParam("records", qRecords); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}