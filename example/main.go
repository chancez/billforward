package main

import (
	"log"
	"os"

	"github.com/authclub/billforward/models"

	"github.com/authclub/billforward/client/accounts"

	"github.com/authclub/billforward/client"
	httpclient "github.com/go-swagger/go-swagger/httpkit/client"
	"github.com/go-swagger/go-swagger/spec"
)

func main() {
	log.Println("starting")
	swaggerSpec, err := spec.New(client.SwaggerJSON, "")
	if err != nil {
		log.Fatal(err)
	}
	transport := httpclient.New(swaggerSpec)
	// transport.Debug = true
	transport.Host = "api-sandbox.billforward.net:443"
	transport.BasePath = "/v1"
	transport.DefaultAuthentication = httpclient.BearerToken(os.Getenv("BILLFORWARD_API_KEY"))
	bfClient := client.New(transport, nil)

	knownAccountID := "ACC-4C23D41A-F547-4402-A0B3-DDF8455B"

	var acct *models.Account

	log.Println("getting all accounts")
	acctsResp, err := bfClient.Accounts.GetAllAccounts(accounts.GetAllAccountsParams{
		Order:   "DESC",
		OrderBy: "created",
	})
	if err != nil {
		gerr := err.(accounts.APIError).Response.(*accounts.GetAllAccountsDefault).Payload
		log.Fatal(gerr)
	}
	for _, result := range acctsResp.Payload.Results {
		if !result.Deleted && result.ID == knownAccountID {
			log.Println("found existing account", result)
			acct = result
			break
		}
	}

	// Only create the account if it doesnt exist
	if acct == nil {
		log.Println("creating new account")
		createAcctResp, err := bfClient.Accounts.CreateAccount(accounts.CreateAccountParams{
			Account: &models.Account{
				Profile: &models.Profile{
					FirstName:   "chancetest",
					LastName:    "tester",
					Email:       "chance.zibolski+test@coreos.com",
					CompanyName: "enterprise company",
				},
			},
		})
		if err != nil {
			gerr := err.(accounts.APIError).Response.(*accounts.CreateAccountDefault).Payload
			log.Fatal(gerr)
		}
		acct = createAcctResp.Payload.Results[0]
	}

	log.Println("getting account by ID", acct.ID)
	// just for fun
	getAcctResp, err := bfClient.Accounts.GetAccountByID(accounts.GetAccountByIDParams{
		AccountID: acct.ID,
	})
	if err != nil {
		log.Fatal(err)
	}
	acct = getAcctResp.Payload.Results[0]
	log.Println("acct:", acct)

	// productsResp, err := bfClient.Products.GetAllProducts(products.GetAllProductsParams{
	// 	Records: 10,
	// 	Order:   "DESC",
	// 	OrderBy: "created",
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// for _, product := range productsResp.Payload.Results {
	// 	log.Println("product", product.Name)
	// }

	// dob := time.Date(1993, time.July, 11, 0, 0, 0, 0, time.UTC)
	// updateProfileResp, err := bfClient.Profiles.UpdateProfile(profiles.UpdateProfileParams{
	// 	Profile: &models.UpdateProfileRequest{
	// 		AccountID: acct.ID,
	// 		// FirstName: "chancetest",
	// 		// LastName: "another tester",
	// 		// Dob:      &strfmt.DateTime{dob},
	// 	},
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Println("update profile response:", updateProfileResp.Payload.Results[0])

	// createAddrResp, err := bfClient.Addresses.CreateAddress(addresses.CreateAddressParams{
	// 	Request: &models.CreateAddressRequest{
	// 		ProfileID:    acct.Profile.ID,
	// 		AddressLine1: "555 nw kings st",
	// 		City:         "san francisco",
	// 		Postcode:     "94110",
	// 		Province:     "CA",
	// 		Country:      "US",
	// 	},
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Println("create address response:", createAddrResp.Payload.Results[0])

	// updateAddrResp, err := bfClient.Addresses.UpdateAddress(addresses.UpdateAddressParams{
	// 	Request: &models.UpdateAddressRequest{
	// 		ProfileID:    acct.Profile.ID,
	// 		AddressLine1: "555 nw kings st",
	// 		City:         "san francisco",
	// 		Postcode:     "99033",
	// 		Province:     "CA",
	// 		Country:      "US",
	// 	},
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Println("update address response:", updateAddrResp.Payload.Results[0])

	// tokenResp, err := bfClient.Tokenization.AuthCapture(tokenization.AuthCaptureParams{
	// 	Request: &models.AuthCaptureRequest{
	// 		AccountID: acct.ID,
	// 		Gateway:   "stripe",
	// 	},
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Println("tokenResp:", tokenResp.Payload.Results)

	// log.Println("No subscriptions, creating one")
	// // TODO (remove). This can be automatically created when we create accounts
	// createSubResp, err := bfClient.Subscriptions.CreateAggregatingSubscription(subscriptions.CreateAggregatingSubscriptionParams{
	// 	Request: &models.CreateAggregatingSubscriptionRequest{
	// 		AccountID:      acct.ID,
	// 		Duration:       30,
	// 		DurationPeriod: "days",
	// 		ProductType:    "recurring",
	// 		Currency:       "USD",
	// 		AggregateAllSubscriptionsOnAccount: true,
	// 		Description:                        fmt.Sprintf("Aggregating subscription for %s", acct.Profile.CompanyName),
	// 	},
	// })
	// if err != nil {
	// 	gerr := err.(subscriptions.APIError).Response.(*subscriptions.CreateAggregatingSubscriptionDefault).Payload
	// 	log.Fatal(gerr)
	// }

	// result := createSubResp.Payload.Results[0]
	// log.Println("create aggregating subscription result:", result)

	// log.Println("creating premium managed linux subscription")
	// _, err = bfClient.Subscriptions.CreateSubscription(subscriptions.CreateSubscriptionParams{
	// 	Request: &models.CreateSubscriptionRequest{
	// 		AccountID:       acct.ID,
	// 		Name:            "sub for enterprisecompany.com",
	// 		ProductRatePlan: "PRP-84EE41D7-09E9-4CAC-B0C1-C898D718", // Premium Managed Linux
	// 		State:           "AwaitingPayment",
	// 		Description:     "testing",
	// 	},
	// })
	// if err != nil {
	// 	gerr := err.(subscriptions.APIError).Response.(*subscriptions.CreateSubscriptionDefault).Payload
	// 	log.Fatal(gerr)
	// }

	// log.Println("creating pro services subscription")
	// _, err = bfClient.Subscriptions.CreateSubscription(subscriptions.CreateSubscriptionParams{
	// 	Request: &models.CreateSubscriptionRequest{
	// 		AccountID:       acct.ID,
	// 		Name:            "sub for enterprisecompany.com",
	// 		ProductRatePlan: "PRP-F1E1C03E-4C41-47ED-BB9B-CABADC72", // pro services
	// 		State:           "AwaitingPayment",
	// 		Description:     "test",
	// 	},
	// })
	// if err != nil {
	// 	gerr := err.(subscriptions.APIError).Response.(*subscriptions.CreateSubscriptionDefault).Payload
	// 	log.Fatal(gerr)
	// }
}

func String(s string) *string {
	return &s
}
