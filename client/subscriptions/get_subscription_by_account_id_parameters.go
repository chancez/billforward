package subscriptions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/strfmt"
	"github.com/go-swagger/go-swagger/swag"
)

// NewGetSubscriptionByAccountIDParams creates a new GetSubscriptionByAccountIDParams object
// with the default values initialized.
func NewGetSubscriptionByAccountIDParams() *GetSubscriptionByAccountIDParams {
	var (
		orderDefault   string = string("DESC")
		orderByDefault string = string("id")
		recordsDefault int32  = int32(10)
	)
	return &GetSubscriptionByAccountIDParams{
		Order:   &orderDefault,
		OrderBy: &orderByDefault,
		Records: &recordsDefault,
	}
}

/*GetSubscriptionByAccountIDParams contains all the parameters to send to the API endpoint
for the get subscription by account ID operation typically these are written to a http.Request
*/
type GetSubscriptionByAccountIDParams struct {

	/*AccountID*/
	AccountID string
	/*IncludeRetired
	  Whether retired subscriptions should be returned.

	*/
	IncludeRetired *bool
	/*Offset
	  The offset from the first subscription to return.

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
	  The maximum number of subscriptions to return.

	*/
	Records *int32
}

// WithAccountID adds the accountId to the get subscription by account ID params
func (o *GetSubscriptionByAccountIDParams) WithAccountID(accountId string) *GetSubscriptionByAccountIDParams {
	o.AccountID = accountId
	return o
}

// WithIncludeRetired adds the includeRetired to the get subscription by account ID params
func (o *GetSubscriptionByAccountIDParams) WithIncludeRetired(includeRetired *bool) *GetSubscriptionByAccountIDParams {
	o.IncludeRetired = includeRetired
	return o
}

// WithOffset adds the offset to the get subscription by account ID params
func (o *GetSubscriptionByAccountIDParams) WithOffset(offset *int32) *GetSubscriptionByAccountIDParams {
	o.Offset = offset
	return o
}

// WithOrder adds the order to the get subscription by account ID params
func (o *GetSubscriptionByAccountIDParams) WithOrder(order *string) *GetSubscriptionByAccountIDParams {
	o.Order = order
	return o
}

// WithOrderBy adds the orderBy to the get subscription by account ID params
func (o *GetSubscriptionByAccountIDParams) WithOrderBy(orderBy *string) *GetSubscriptionByAccountIDParams {
	o.OrderBy = orderBy
	return o
}

// WithOrganizations adds the organizations to the get subscription by account ID params
func (o *GetSubscriptionByAccountIDParams) WithOrganizations(organizations []string) *GetSubscriptionByAccountIDParams {
	o.Organizations = organizations
	return o
}

// WithRecords adds the records to the get subscription by account ID params
func (o *GetSubscriptionByAccountIDParams) WithRecords(records *int32) *GetSubscriptionByAccountIDParams {
	o.Records = records
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *GetSubscriptionByAccountIDParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	// path param account-ID
	if err := r.SetPathParam("account-ID", o.AccountID); err != nil {
		return err
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
