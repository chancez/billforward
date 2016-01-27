package product_rate_plans

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/strfmt"
	"github.com/go-swagger/go-swagger/swag"

	"github.com/authclub/billforward/models"
)

// NewUpsertMetadataForRatePlanParams creates a new UpsertMetadataForRatePlanParams object
// with the default values initialized.
func NewUpsertMetadataForRatePlanParams() *UpsertMetadataForRatePlanParams {
	var ()
	return &UpsertMetadataForRatePlanParams{}
}

/*UpsertMetadataForRatePlanParams contains all the parameters to send to the API endpoint
for the upsert metadata for rate plan operation typically these are written to a http.Request
*/
type UpsertMetadataForRatePlanParams struct {

	/*Metadata*/
	Metadata *models.DynamicMetadata
	/*Organizations
	  A list of organization-IDs used to restrict the scope of API calls.

	*/
	Organizations []string
	/*ProductRatePlanID*/
	ProductRatePlanID string
}

// WithMetadata adds the metadata to the upsert metadata for rate plan params
func (o *UpsertMetadataForRatePlanParams) WithMetadata(metadata *models.DynamicMetadata) *UpsertMetadataForRatePlanParams {
	o.Metadata = metadata
	return o
}

// WithOrganizations adds the organizations to the upsert metadata for rate plan params
func (o *UpsertMetadataForRatePlanParams) WithOrganizations(organizations []string) *UpsertMetadataForRatePlanParams {
	o.Organizations = organizations
	return o
}

// WithProductRatePlanID adds the productRatePlanId to the upsert metadata for rate plan params
func (o *UpsertMetadataForRatePlanParams) WithProductRatePlanID(productRatePlanId string) *UpsertMetadataForRatePlanParams {
	o.ProductRatePlanID = productRatePlanId
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *UpsertMetadataForRatePlanParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	if err := r.SetBodyParam(o.Metadata); err != nil {
		return err
	}

	valuesOrganizations := o.Organizations

	joinedOrganizations := swag.JoinByFormat(valuesOrganizations, "multi")
	// query array param organizations
	if err := r.SetQueryParam("organizations", joinedOrganizations...); err != nil {
		return err
	}

	// path param product-rate-plan-ID
	if err := r.SetPathParam("product-rate-plan-ID", o.ProductRatePlanID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
