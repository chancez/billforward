package types

import (
	"time"
)

type Product struct {
	Id             string    `json:"id,omitempty"`
	AccountId      string    `json:"accountID,omitempty"`
	Created        time.Time `json:"created,omitempty"`
	Deleted        bool      `json:"deleted"`
	Description    string    `json:"description"`
	Duration       int32     `json:"duration"`
	DurationPeriod string    `json:"durationPeriod"`
	Name           string    `json:"name"`
	ProductType    string    `json:"productType"`
	PublicName     string    `json:"publicName,omitempty"`
	Trial          int32     `json:"trial"`
	TrialPeriod    string    `json:"trialPeriod"`
	Updated        time.Time `json:"updated,omitempty"`
}

type ProductsResponse struct {
	PaginatedResponse
	Results []*Product `json:"results"`
}

type ProductsFilter struct {
	IncludeRetired bool
}
