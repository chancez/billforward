package subscriptions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/strfmt"

	"github.com/authclub/billforward/models"
)

/*
CancelSubscriptionParams contains all the parameters to send to the API endpoint
for the cancel subscription operation typically these are written to a http.Request
*/
type CancelSubscriptionParams struct {
	/*
	  The cancellation request
	*/
	Request *models.MutableBillingEntity

	SubscriptionID string
}

// WriteToRequest writes these params to a swagger request
func (o *CancelSubscriptionParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	if o.Request == nil {
		o.Request = new(models.MutableBillingEntity)
	}

	if err := r.SetBodyParam(o.Request); err != nil {
		return err
	}

	// path param subscription-ID
	if err := r.SetPathParam("subscription-ID", o.SubscriptionID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
