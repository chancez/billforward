package subscriptions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/authclub/billforward/models"
)

// NewReviveSubscriptionParams creates a new ReviveSubscriptionParams object
// with the default values initialized.
func NewReviveSubscriptionParams() *ReviveSubscriptionParams {
	var ()
	return &ReviveSubscriptionParams{}
}

/*ReviveSubscriptionParams contains all the parameters to send to the API endpoint
for the revive subscription operation typically these are written to a http.Request
*/
type ReviveSubscriptionParams struct {

	/*Request
	  The revive request

	*/
	Request *models.ReviveSubscriptionRequest
	/*SubscriptionID*/
	SubscriptionID string
}

// WithRequest adds the request to the revive subscription params
func (o *ReviveSubscriptionParams) WithRequest(Request *models.ReviveSubscriptionRequest) *ReviveSubscriptionParams {
	o.Request = Request
	return o
}

// WithSubscriptionID adds the subscriptionId to the revive subscription params
func (o *ReviveSubscriptionParams) WithSubscriptionID(SubscriptionID string) *ReviveSubscriptionParams {
	o.SubscriptionID = SubscriptionID
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *ReviveSubscriptionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	var res []error

	if o.Request == nil {
		o.Request = new(models.ReviveSubscriptionRequest)
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
