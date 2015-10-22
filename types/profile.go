package types

import (
	"time"
)

/*

Represents an Account Profile.

Docs Example:

{
  "id": "PRF-05AA-508B-4BC8-A30F-BEA976F04535",
  "accountID": "ACC-4CA-9D68-467A-A685-2E805A1AAE7C",
  "email": "john.smith@test.com",
  "firstName": "John",
  "lastName": "Smith",
  "companyName": "Acme Inc.",
  "addresses": [
    {
      "id": "ADD-05AA-508B-4BC8-A30F-BEA976F04535",
      "addressLine1": "Test Address",
      "addressLine2": "Test Address",
      "addressLine3": "Test Address",
      "city": "London",
      "province": "London",
      "country": "UK",
      "postcode": "SW1 1AA",
      "landline": "02000000000",
      "primaryAddress": true
    }
  ],
  "landline": "02000000000",
  "mobile": "07700000000",
  "dob": "1970-01-01T00:00:00Z"
}

*/
type Profile struct {
	Id                    string     `json:"id,omitempty"`
	AccountID             string     `json:"accountID,omitempty"`
	Email                 string     `json:"email"`
	FirstName             string     `json:"firstName"`
	LastName              string     `json:"lastName"`
	CompanyName           string     `json:"companyName,omitempty"`
	LogoURL               string     `json:"logoURL,omitempty"`
	Addresses             []*Address `json:"addresses,omitempty"`
	Mobile                string     `json:"mobile,omitempty"`
	Landline              string     `json:"landline,omitempty"`
	Fax                   string     `json:"fax,omitempty"`
	Dob                   time.Time  `json:"dob,omitempty"`
	VatNumber             string     `json:"vatNumber,omitempty"`
	AdditionalInformation string     `json:"additionalInformation,omitempty"`
}

/*

Repersents an Address, part of a Porfile.

Docs Example:

    {
      "id": "ADD-34C6-0C8A-11E4-98D0-2F1CFEA9529D",
      "profileID": "723A380F-0AD4-4CE6-9B7A-3EFB037C7A3E",
      "addressLine1": "address line 1",
      "addressLine2": "address line 2",
      "addressLine3": "address line 3",
      "city": "London",
      "province": "London",
      "country": "United Kingdom",
      "postcode": "SW1 1AS",
      "landline": "02000000000",
      "primaryAddress": true,
      "deleted": false
    }
*/
type Address struct {
	Id             string `json:"id,omitempty"`
	ProfileId      string `json:"profileID,omitempty"`
	AddressLine1   string `json:"addressLine1"`
	AddressLine2   string `json:"addressLine1,omitempty"`
	AddressLine3   string `json:"addressLine1,omitempty"`
	City           string `json:"city"`
	Province       string `json:"province"`
	Country        string `json:"country"`
	Postcode       string `json:"postcode"`
	Landline       string `json:"landline,omitempty"`
	PrimaryAddress bool   `json:"primaryAddress,omitempty"`
	Deleted        bool   `json:"primaryAddress,omitempty"`
}
