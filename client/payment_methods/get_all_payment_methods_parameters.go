package payment_methods

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetAllPaymentMethodsParams creates a new GetAllPaymentMethodsParams object
// with the default values initialized.
func NewGetAllPaymentMethodsParams() *GetAllPaymentMethodsParams {
	var (
		includeRetiredDefault bool   = bool(true)
		offsetDefault         int32  = int32(0)
		orderDefault          string = string("DESC")
		orderByDefault        string = string("created")
		recordsDefault        int32  = int32(10)
	)
	return &GetAllPaymentMethodsParams{
		IncludeRetired: &includeRetiredDefault,
		Offset:         &offsetDefault,
		Order:          &orderDefault,
		OrderBy:        &orderByDefault,
		Records:        &recordsDefault,
	}
}

/*GetAllPaymentMethodsParams contains all the parameters to send to the API endpoint
for the get all payment methods operation typically these are written to a http.Request
*/
type GetAllPaymentMethodsParams struct {

	/*IncludeRetired
	  Whether retired products should be returned.

	*/
	IncludeRetired *bool
	/*Offset
	  The offset from the first payment-method to return.

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
	  The maximum number of payment-methods to return.

	*/
	Records *int32
}

// WithIncludeRetired adds the includeRetired to the get all payment methods params
func (o *GetAllPaymentMethodsParams) WithIncludeRetired(IncludeRetired *bool) *GetAllPaymentMethodsParams {
	o.IncludeRetired = IncludeRetired
	return o
}

// WithOffset adds the offset to the get all payment methods params
func (o *GetAllPaymentMethodsParams) WithOffset(Offset *int32) *GetAllPaymentMethodsParams {
	o.Offset = Offset
	return o
}

// WithOrder adds the order to the get all payment methods params
func (o *GetAllPaymentMethodsParams) WithOrder(Order *string) *GetAllPaymentMethodsParams {
	o.Order = Order
	return o
}

// WithOrderBy adds the orderBy to the get all payment methods params
func (o *GetAllPaymentMethodsParams) WithOrderBy(OrderBy *string) *GetAllPaymentMethodsParams {
	o.OrderBy = OrderBy
	return o
}

// WithOrganizations adds the organizations to the get all payment methods params
func (o *GetAllPaymentMethodsParams) WithOrganizations(Organizations []string) *GetAllPaymentMethodsParams {
	o.Organizations = Organizations
	return o
}

// WithRecords adds the records to the get all payment methods params
func (o *GetAllPaymentMethodsParams) WithRecords(Records *int32) *GetAllPaymentMethodsParams {
	o.Records = Records
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *GetAllPaymentMethodsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	var res []error

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
