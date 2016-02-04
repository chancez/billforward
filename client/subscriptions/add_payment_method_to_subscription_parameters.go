package subscriptions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/strfmt"

	"github.com/authclub/billforward/models"
)

// NewAddPaymentMethodToSubscriptionParams creates a new AddPaymentMethodToSubscriptionParams object
// with the default values initialized.
func NewAddPaymentMethodToSubscriptionParams() *AddPaymentMethodToSubscriptionParams {
	var ()
	return &AddPaymentMethodToSubscriptionParams{}
}

/*AddPaymentMethodToSubscriptionParams contains all the parameters to send to the API endpoint
for the add payment method to subscription operation typically these are written to a http.Request
*/
type AddPaymentMethodToSubscriptionParams struct {

	/*PaymentMethod*/
	PaymentMethod *models.AddPaymentMethodRequest
	/*SubscriptionID*/
	SubscriptionID string
}

// WithPaymentMethod adds the paymentMethod to the add payment method to subscription params
func (o *AddPaymentMethodToSubscriptionParams) WithPaymentMethod(paymentMethod *models.AddPaymentMethodRequest) *AddPaymentMethodToSubscriptionParams {
	o.PaymentMethod = paymentMethod
	return o
}

// WithSubscriptionID adds the subscriptionId to the add payment method to subscription params
func (o *AddPaymentMethodToSubscriptionParams) WithSubscriptionID(subscriptionId string) *AddPaymentMethodToSubscriptionParams {
	o.SubscriptionID = subscriptionId
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *AddPaymentMethodToSubscriptionParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	if o.PaymentMethod == nil {
		o.PaymentMethod = new(models.AddPaymentMethodRequest)
	}

	if err := r.SetBodyParam(o.PaymentMethod); err != nil {
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
