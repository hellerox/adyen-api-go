package adyen

import "fmt"

// CheckoutGateway - allows you to accept all of Adyen's payment
// methods and flows.
type CheckoutGateway struct {
	*Adyen
}

const (
	paymentMethodsURL = "paymentMethods"
	paymentsURL       = "payments"
)

// PaymentMethods - Perform paymentMethods request in Adyen.
//
// Used to get a collection of available payment methods for a merchant.
func (a *CheckoutGateway) PaymentMethods(req *PaymentMethods) (*PaymentMethodsResponse, error) {
	url := a.checkoutURL(paymentMethodsURL, CheckoutAPIVersion)

	resp, err := a.execute(url, req)
	if err != nil {
		return nil, err
	}

	return resp.paymentMethods()
}

// Payments - Perform payments request in Adyen.
//
// Used to get a collection of available payment methods for a merchant.
func (a *CheckoutGateway) Payments(req *PaymentRequest) (*PaymentResponse, error) {
	url := a.checkoutURL(paymentsURL, CheckoutAPIVersion)

	resp, err := a.execute(url, req)
	if err != nil {
		return nil, err
	}

	paymentResponse, err := resp.payments()
	if err != nil {
		return nil, err
	}
	if paymentResponse.Error != nil {
		return paymentResponse, fmt.Errorf(paymentResponse.Error.Message)
	}

	return paymentResponse, err
}
