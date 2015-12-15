package models

import (
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/strfmt"
)

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

/*
TieredPricingComponent tiered pricing component

swagger:model TieredPricingComponent
*/
type TieredPricingComponent struct {
	PricingComponent
}

// Validate validates this tiered pricing component
func (m *TieredPricingComponent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.PricingComponent.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}