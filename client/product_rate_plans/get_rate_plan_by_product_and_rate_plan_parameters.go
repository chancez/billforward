package product_rate_plans

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetRatePlanByProductAndRatePlanParams creates a new GetRatePlanByProductAndRatePlanParams object
// with the default values initialized.
func NewGetRatePlanByProductAndRatePlanParams() *GetRatePlanByProductAndRatePlanParams {
	var (
		includeRetiredDefault bool   = bool(false)
		offsetDefault         int32  = int32(0)
		orderDefault          string = string("DESC")
		orderByDefault        string = string("created")
		recordsDefault        int32  = int32(10)
	)
	return &GetRatePlanByProductAndRatePlanParams{
		IncludeRetired: &includeRetiredDefault,
		Offset:         &offsetDefault,
		Order:          &orderDefault,
		OrderBy:        &orderByDefault,
		Records:        &recordsDefault,
	}
}

/*GetRatePlanByProductAndRatePlanParams contains all the parameters to send to the API endpoint
for the get rate plan by product and rate plan operation typically these are written to a http.Request
*/
type GetRatePlanByProductAndRatePlanParams struct {

	/*IncludeRetired
	  Whether retired products should be returned.

	*/
	IncludeRetired *bool
	/*Offset
	  The offset from the first product-rate-plan to return.

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
	/*ProductID*/
	ProductID string
	/*RatePlanID*/
	RatePlanID string
	/*Records
	  The maximum number of product-rate-plans to return.

	*/
	Records *int32
}

// WithIncludeRetired adds the includeRetired to the get rate plan by product and rate plan params
func (o *GetRatePlanByProductAndRatePlanParams) WithIncludeRetired(IncludeRetired *bool) *GetRatePlanByProductAndRatePlanParams {
	o.IncludeRetired = IncludeRetired
	return o
}

// WithOffset adds the offset to the get rate plan by product and rate plan params
func (o *GetRatePlanByProductAndRatePlanParams) WithOffset(Offset *int32) *GetRatePlanByProductAndRatePlanParams {
	o.Offset = Offset
	return o
}

// WithOrder adds the order to the get rate plan by product and rate plan params
func (o *GetRatePlanByProductAndRatePlanParams) WithOrder(Order *string) *GetRatePlanByProductAndRatePlanParams {
	o.Order = Order
	return o
}

// WithOrderBy adds the orderBy to the get rate plan by product and rate plan params
func (o *GetRatePlanByProductAndRatePlanParams) WithOrderBy(OrderBy *string) *GetRatePlanByProductAndRatePlanParams {
	o.OrderBy = OrderBy
	return o
}

// WithOrganizations adds the organizations to the get rate plan by product and rate plan params
func (o *GetRatePlanByProductAndRatePlanParams) WithOrganizations(Organizations []string) *GetRatePlanByProductAndRatePlanParams {
	o.Organizations = Organizations
	return o
}

// WithProductID adds the productId to the get rate plan by product and rate plan params
func (o *GetRatePlanByProductAndRatePlanParams) WithProductID(ProductID string) *GetRatePlanByProductAndRatePlanParams {
	o.ProductID = ProductID
	return o
}

// WithRatePlanID adds the ratePlanId to the get rate plan by product and rate plan params
func (o *GetRatePlanByProductAndRatePlanParams) WithRatePlanID(RatePlanID string) *GetRatePlanByProductAndRatePlanParams {
	o.RatePlanID = RatePlanID
	return o
}

// WithRecords adds the records to the get rate plan by product and rate plan params
func (o *GetRatePlanByProductAndRatePlanParams) WithRecords(Records *int32) *GetRatePlanByProductAndRatePlanParams {
	o.Records = Records
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *GetRatePlanByProductAndRatePlanParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param product-ID
	if err := r.SetPathParam("product-ID", o.ProductID); err != nil {
		return err
	}

	// path param rate-plan-ID
	if err := r.SetPathParam("rate-plan-ID", o.RatePlanID); err != nil {
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
