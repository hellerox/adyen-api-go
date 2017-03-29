package adyen

import "encoding/json"

// AuthoriseGateway - Adyen payment authorisation logic
type AuthoriseGateway struct {
	*Adyen
}

// RequestType - authorise type request, @TODO: move to enums
const RequestType = "authorise"

// Payment - Perform authorise payment in Adyen
func (a *AuthoriseGateway) Payment(encryptedData string, reference string, amount float32) (*AuthoriseResponse, error) {
	auth := AuthoriseRequest{
		Amount:          Amount{Value: amount, Currency: "EUR"},
		MerchantAccount: a.MerchantAccount,
		AdditionalData:  AdditionalData{Content: encryptedData},
		Reference:       reference,
	}

	resp, err := a.execute(RequestType, auth)

	if err != nil {
		return nil, err
	}

	var val AuthoriseResponse
	json.NewDecoder(resp.Body).Decode(&val)

	return &val, nil
}