package types

/**
*

{
  "id": "PMD-3E6E-F0F5-4D72-88FF-9B54482E8AB2",
  "accountID": "ACC-FBB05-2C4C-4A94-85F3-058D02B2CE0E",
  "name": "My Card",
  "description": "############1111",
  "cardHolderName": "John Smith",
  "expiryDate": "0116",
  "cardType": "Visa",
  "linkID": "832C7172-FD9B-45B9-AC33-4BB82B584058",
  "gateway": "stripe"
}

*/

type PaymentMethod struct {
	Id                   string `json:"id,omitempty"`
	AccountID            string `json:"accountID,omitempty"`
	Name                 string `json:"name"`
	Description          string `json:"description,omitempty"`
	CardHolderName       string `json:"cardHolderName,omitempty"`
	ExpiryDate           string `json:"expiryDate"`
	CardType             string `json:"cardType,omitempty"`
	Gateway              string `json:"gateway,omitempty"`
	LinkID               string `json:"linkID,omitempty"`
	DefaultPaymentMethod bool   `json:"defaultPaymentMethod,omitempty"`
	Deleted              bool   `json:"deleted,omitempty"`
}
