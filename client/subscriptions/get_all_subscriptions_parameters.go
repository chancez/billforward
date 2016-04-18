package subscriptions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetAllSubscriptionsParams creates a new GetAllSubscriptionsParams object
// with the default values initialized.
func NewGetAllSubscriptionsParams() *GetAllSubscriptionsParams {
	var (
		excludeChildrenDefault bool   = bool(true)
		includeRetiredDefault  bool   = bool(false)
		offsetDefault          int32  = int32(0)
		orderDefault           string = string("DESC")
		orderByDefault         string = string("created")
		recordsDefault         int32  = int32(10)
	)
	return &GetAllSubscriptionsParams{
		ExcludeChildren: &excludeChildrenDefault,
		IncludeRetired:  &includeRetiredDefault,
		Offset:          &offsetDefault,
		Order:           &orderDefault,
		OrderBy:         &orderByDefault,
		Records:         &recordsDefault,
	}
}

/*GetAllSubscriptionsParams contains all the parameters to send to the API endpoint
for the get all subscriptions operation typically these are written to a http.Request
*/
type GetAllSubscriptionsParams struct {

	/*AccountID
	  A list of accountIDs to filter subscriptions on

	*/
	AccountID []string
	/*ExcludeChildren
	  Should child subscriptiosn be excluded.

	*/
	ExcludeChildren *bool
	/*ExcludeServiceEnded*/
	ExcludeServiceEnded *bool
	/*IncludeRetired
	  Whether retired subscriptions should be returned.

	*/
	IncludeRetired *bool
	/*Metadata*/
	Metadata *string
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

// WithAccountID adds the accountId to the get all subscriptions params
func (o *GetAllSubscriptionsParams) WithAccountID(AccountID []string) *GetAllSubscriptionsParams {
	o.AccountID = AccountID
	return o
}

// WithExcludeChildren adds the excludeChildren to the get all subscriptions params
func (o *GetAllSubscriptionsParams) WithExcludeChildren(ExcludeChildren *bool) *GetAllSubscriptionsParams {
	o.ExcludeChildren = ExcludeChildren
	return o
}

// WithExcludeServiceEnded adds the excludeServiceEnded to the get all subscriptions params
func (o *GetAllSubscriptionsParams) WithExcludeServiceEnded(ExcludeServiceEnded *bool) *GetAllSubscriptionsParams {
	o.ExcludeServiceEnded = ExcludeServiceEnded
	return o
}

// WithIncludeRetired adds the includeRetired to the get all subscriptions params
func (o *GetAllSubscriptionsParams) WithIncludeRetired(IncludeRetired *bool) *GetAllSubscriptionsParams {
	o.IncludeRetired = IncludeRetired
	return o
}

// WithMetadata adds the metadata to the get all subscriptions params
func (o *GetAllSubscriptionsParams) WithMetadata(Metadata *string) *GetAllSubscriptionsParams {
	o.Metadata = Metadata
	return o
}

// WithOffset adds the offset to the get all subscriptions params
func (o *GetAllSubscriptionsParams) WithOffset(Offset *int32) *GetAllSubscriptionsParams {
	o.Offset = Offset
	return o
}

// WithOrder adds the order to the get all subscriptions params
func (o *GetAllSubscriptionsParams) WithOrder(Order *string) *GetAllSubscriptionsParams {
	o.Order = Order
	return o
}

// WithOrderBy adds the orderBy to the get all subscriptions params
func (o *GetAllSubscriptionsParams) WithOrderBy(OrderBy *string) *GetAllSubscriptionsParams {
	o.OrderBy = OrderBy
	return o
}

// WithOrganizations adds the organizations to the get all subscriptions params
func (o *GetAllSubscriptionsParams) WithOrganizations(Organizations []string) *GetAllSubscriptionsParams {
	o.Organizations = Organizations
	return o
}

// WithRecords adds the records to the get all subscriptions params
func (o *GetAllSubscriptionsParams) WithRecords(Records *int32) *GetAllSubscriptionsParams {
	o.Records = Records
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *GetAllSubscriptionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	var res []error

	valuesAccountID := o.AccountID

	joinedAccountID := swag.JoinByFormat(valuesAccountID, "multi")
	// query array param accountID
	if err := r.SetQueryParam("accountID", joinedAccountID...); err != nil {
		return err
	}

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

	if o.ExcludeServiceEnded != nil {

		// query param exclude_service_ended
		var qrExcludeServiceEnded bool
		if o.ExcludeServiceEnded != nil {
			qrExcludeServiceEnded = *o.ExcludeServiceEnded
		}
		qExcludeServiceEnded := swag.FormatBool(qrExcludeServiceEnded)
		if qExcludeServiceEnded != "" {
			if err := r.SetQueryParam("exclude_service_ended", qExcludeServiceEnded); err != nil {
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

	if o.Metadata != nil {

		// query param metadata
		var qrMetadata string
		if o.Metadata != nil {
			qrMetadata = *o.Metadata
		}
		qMetadata := qrMetadata
		if qMetadata != "" {
			if err := r.SetQueryParam("metadata", qMetadata); err != nil {
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
