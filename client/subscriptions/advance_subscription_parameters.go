package subscriptions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"

	strfmt "github.com/go-swagger/go-swagger/strfmt"

	"github.com/authclub/billforward/models"
)

// NewAdvanceSubscriptionParams creates a new AdvanceSubscriptionParams object
// with the default values initialized.
func NewAdvanceSubscriptionParams() *AdvanceSubscriptionParams {
	var ()
	return &AdvanceSubscriptionParams{}
}

/*AdvanceSubscriptionParams contains all the parameters to send to the API endpoint
for the advance subscription operation typically these are written to a http.Request
*/
type AdvanceSubscriptionParams struct {

	/*Request
	  The request

	*/
	Request *models.TimeRequest
	/*SubscriptionID
	  ID of the subscription.

	*/
	SubscriptionID string
}

// WithRequest adds the request to the advance subscription params
func (o *AdvanceSubscriptionParams) WithRequest(request *models.TimeRequest) *AdvanceSubscriptionParams {
	o.Request = request
	return o
}

// WithSubscriptionID adds the subscriptionId to the advance subscription params
func (o *AdvanceSubscriptionParams) WithSubscriptionID(subscriptionId string) *AdvanceSubscriptionParams {
	o.SubscriptionID = subscriptionId
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *AdvanceSubscriptionParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	if o.Request == nil {
		o.Request = new(models.TimeRequest)
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
