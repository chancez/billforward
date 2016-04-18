package subscriptions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/authclub/billforward/models"
)

// NewUpdateSubscriptionV2Params creates a new UpdateSubscriptionV2Params object
// with the default values initialized.
func NewUpdateSubscriptionV2Params() *UpdateSubscriptionV2Params {
	var ()
	return &UpdateSubscriptionV2Params{}
}

/*UpdateSubscriptionV2Params contains all the parameters to send to the API endpoint
for the update subscription v2 operation typically these are written to a http.Request
*/
type UpdateSubscriptionV2Params struct {

	/*Request*/
	Request *models.UpdateSubscriptionRequest
}

// WithRequest adds the request to the update subscription v2 params
func (o *UpdateSubscriptionV2Params) WithRequest(Request *models.UpdateSubscriptionRequest) *UpdateSubscriptionV2Params {
	o.Request = Request
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateSubscriptionV2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	var res []error

	if o.Request == nil {
		o.Request = new(models.UpdateSubscriptionRequest)
	}

	if err := r.SetBodyParam(o.Request); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
