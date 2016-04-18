package products

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteMetadataForProductParams creates a new DeleteMetadataForProductParams object
// with the default values initialized.
func NewDeleteMetadataForProductParams() *DeleteMetadataForProductParams {
	var ()
	return &DeleteMetadataForProductParams{}
}

/*DeleteMetadataForProductParams contains all the parameters to send to the API endpoint
for the delete metadata for product operation typically these are written to a http.Request
*/
type DeleteMetadataForProductParams struct {

	/*Organizations
	  A list of organization-IDs used to restrict the scope of API calls.

	*/
	Organizations []string
	/*ProductID*/
	ProductID string
}

// WithOrganizations adds the organizations to the delete metadata for product params
func (o *DeleteMetadataForProductParams) WithOrganizations(Organizations []string) *DeleteMetadataForProductParams {
	o.Organizations = Organizations
	return o
}

// WithProductID adds the productId to the delete metadata for product params
func (o *DeleteMetadataForProductParams) WithProductID(ProductID string) *DeleteMetadataForProductParams {
	o.ProductID = ProductID
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteMetadataForProductParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	var res []error

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
