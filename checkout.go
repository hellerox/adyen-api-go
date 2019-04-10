package adyen

import "fmt"

// PaymentMethods contains the fields required by the checkout
// API's /paymentMethods endpoint.  See the following for more
// information:
//
// https://docs.adyen.com/api-explorer/#/PaymentSetupAndVerificationService/v32/paymentMethods
type PaymentMethods struct {
	Amount           *Amount `json:"amount"`
	Channel          string  `json:"channel"`
	CountryCode      string  `json:"countryCode"`
	MerchantAccount  string  `json:"merchantAccount"`
	ShopperLocale    string  `json:"shopperLocale"`
	ShopperReference string  `json:"shopperReference"`
}

// PaymentMethodsResponse is returned by Adyen in response to
// a PaymentMethods request.
type PaymentMethodsResponse struct {
	PaymentMethods         []PaymentMethodDetails         `json:"paymentMethods"`
	OneClickPaymentMethods []OneClickPaymentMethodDetails `json:"oneClickPaymentMethods,omitempty"`
}

// PaymentMethodDetails describes the PaymentMethods part of
// a PaymentMethodsResponse.
type PaymentMethodDetails struct {
	Details []PaymentMethodDetailsInfo `json:"details,omitempty"`
	Name    string                     `json:"name"`
	Type    string                     `json:"type"`
}

// PaymentMethodDetailsInfo describes the collection of all
// payment methods.
type PaymentMethodDetailsInfo struct {
	Items []PaymentMethodItems `json:"items"`
	Key   string               `json:"key"`
	Type  string               `json:"type"`
}

// PaymentMethodItems describes a single payment method.
type PaymentMethodItems struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// OneClickPaymentMethodDetails describes the OneClickPayment part of
// a PaymentMethods response.
type OneClickPaymentMethodDetails struct {
	Details       []PaymentMethodTypes       `json:"details"`
	Name          string                     `json:"name"`
	Type          string                     `json:"type"`
	StoredDetails PaymentMethodStoredDetails `json:"storedDetails"`
}

// PaymentMethodTypes describes any additional information associated
// with a OneClick payment.
type PaymentMethodTypes struct {
	Key  string `json:"key"`
	Type string `json:"type"`
}

// PaymentMethodStoredDetails describes the information stored for a
// OneClick payment.
type PaymentMethodStoredDetails struct {
	Card PaymentMethodCard `json:"card"`
}

// PaymentMethodCard describes the card information associated with a
// OneClick payment.
type PaymentMethodCard struct {
	ExpiryMonth string `json:"expiryMonth"`
	ExpiryYear  string `json:"expiryYear"`
	HolderName  string `json:"holderName"`
	Number      string `json:"number"`
}

// PaymentRequest receives the request for a payment
type PaymentRequest struct {
	AccountInfo              *AccountInfo                `json:"accountInfo"`
	AdditionalData           *AdditionalData             `json:"additionalData"`
	Amount                   *Amount                     `json:"amount"`
	BillingAddress           *Address                    `json:"billingAddress,omitempty"`
	BrowserInfo              *BrowserInfo                `json:"browserInfo,omitempty"`
	CaptureDelayHours        int                         `json:"captureDelayHours,omitempty"`
	Channel                  string                      `json:"channel"`
	Origin                   string                      `json:"origin"`
	Company                  *Company                    `json:"company"`
	CountryCode              string                      `json:"countryCode"`
	DateOfBirth              string                      `json:"dateOfBirth"`
	DccQuote                 interface{}                 `json:"dccQuote"`
	DeliveryAddress          *Address                    `json:"deliveryAddress,omitempty"`
	DeliveryDate             string                      `json:"deliveryDate,omitempty"`
	DeviceFingerprint        string                      `json:"DeviceFingerprint"`
	EnableOneClick           string                      `json:"enableOneClick"`
	EntityType               string                      `json:"entityType"`
	FraudOffset              int                         `json:"fraudOffset"`
	Installments             interface{}                 `json:"installments"`
	LineItems                interface{}                 `json:"lineItems"`
	Mcc                      string                      `json:"mcc"`
	MerchantAccount          string                      `json:"merchantAccount"`
	MerchantData             string                      `json:"merchantData"`           //Holds different merchant data points like product, purchase, customer, and so on. It takes data in a JSON string
	MerchantOrderReference   string                      `json:"merchantOrderReference"` //This reference allows linking multiple transactions to each other
	MerchantRiskIndicator    interface{}                 `json:"merchantRiskIndicator"`
	Metadata                 interface{}                 `json:"metadata"`
	OrderReference           string                      `json:"orderReference"`
	PaymentMethod            *PaymentMethodCardEncrypted `json:"paymentMethod"`
	RecurringProcessingModel string                      `json:"recurringProcessingModel"`
	RedirectFromIssuerMethod string                      `json:"redirectFromIssuerMethod"`
	RedirectToIssuerMethod   string                      `json:"redirectToIssuerMethod"`
	Reference                string                      `json:"reference"`
	ReturnURL                string                      `json:"returnUrl"`
	SessionValidity          string                      `json:"sessionValidity"`
	ShopperEmail             string                      `json:"shopperEmail,omitempty"`
	ShopperIP                string                      `json:"shopperIP,omitempty"`
	ShopperInteraction       string                      `json:"shopperInteraction,omitempty"` // Possible values Ecommerce, ContAuth, Moto, POS
	ShopperLocale            string                      `json:"shopperLocale,omitempty"`
	ShopperName              *Name                       `json:"shopperName,omitempty"`
	ShopperReference         string                      `json:"shopperReference,omitempty"`
	ShopperStatement         string                      `json:"shopperStatement,omitempty"`
	SocialSecurityNumber     string                      `json:"SocialSecurityNumber,omitempty"`
	Splits                   interface{}                 `json:"splits,omitempty"`
	TelephoneNumber          string                      `json:"telephoneNumber,omitempty"`
	ThreeDS2RequestData      ThreeDS2RequestData         `json:"threeDS2RequestData,omitempty"`
	TrustedShopper           bool                        `json:"trustedShopper,omitempty"`
	EnablePayout             bool                        `json:"enablePayout"`
	EnableRecurring          bool                        `json:"enableRecurring"`
}

// ThreeDS2RequestData - Request fields for 3D Secure 2.0
type ThreeDS2RequestData struct {
	AuthenticationOnly  bool   `json:"authenticationOnly,omitempty"`
	ChallengeIndicator  string `json:"challengeIndicator,omitempty"`
	NotificationURL     string `json:"notificationURL,omitempty"`
	ThreeDSRequestorURL string `json:"threeDSRequestorURL,omitempty"`
}

// AccountInfo receives the request for a payment
type AccountInfo struct {
	AccountAgeIndicator           string `json:"accountAgeIndicator,omitempty"`
	AccountChangeDate             string `json:"accountChangeDate,omitempty"`
	AccountChangeIndicator        string `json:"accountChangeIndicator,omitempty"`
	AccountCreationDate           string `json:"accountCreationDate,omitempty"`
	AddCardAttemptsDay            string `json:"addCardAttemptsDay,omitempty"`
	DeliveryAddressUsageDate      string `json:"deliveryAddressUsageDate,omitempty"`
	DeliveryAddressUsageIndicator string `json:"deliveryAddressUsageIndicator,omitempty"`
	HomePhone                     string `json:"homePhone,omitempty"`
	MobilePhone                   string `json:"mobilePhone,omitempty"`
	PasswordChangeDate            string `json:"passwordChangeDate,omitempty"`
	PasswordChangeIndicator       string `json:"passwordChangeIndicator,omitempty"`
	PastTransactionsDay           int    `json:"pastTransactionsDay,omitempty"`
	PastTransactionsYear          int    `json:"pastTransactionsYear,omitempty"`
	PaymentAccountAge             string `json:"paymentAccountAge"`
	PaymentAccountIndicator       string `json:"paymentAccountIndicator"`
	PurchasesLast6Months          int    `json:"purchasesLast6Months"`
	SuspiciousActivity            string `json:"suspiciousActivity"`
	WorkPhone                     string `json:"workPhone"`
}

// PaymentMethodCardEncrypted describes the card information associated with a checkout payment
type PaymentMethodCardEncrypted struct {
	Type                  string `json:"type"`
	EncryptedCardNumber   string `json:"encryptedCardNumber"`
	EncryptedExpiryMonth  string `json:"encryptedExpiryMonth"`
	EncryptedExpiryYear   string `json:"encryptedExpiryYear"`
	EncryptedSecurityCode string `json:"encryptedSecurityCode"`
	HolderName            string `json:"holderName"`
}

// Company Information regarding the compan
type Company struct {
	Homepage           string `json:"homepage"`
	Name               string `json:"name"`
	RegistrationNumber string `json:"registrationNumber"`
	RegistryLocation   string `json:"registryLocation"`
	TaxID              string `json:"taxId"`
	Type               string `json:"type"`
}

// PaymentResponse is a response structure for Adyen
//
// Link - https://docs.adyen.com/developers/api-reference/payments-api#paymentresult
type PaymentResponse struct {
	PspReference   string                  `json:"pspReference"`
	ResultCode     string                  `json:"resultCode"`
	AuthCode       string                  `json:"authCode"`
	RefusalReason  string                  `json:"refusalReason"`
	IssuerURL      string                  `json:"issuerUrl"`
	MD             string                  `json:"md"`
	PaRequest      string                  `json:"paRequest"`
	AdditionalData *CheckoutAdditionalData `json:"additionalData,omitempty"`
	Error          *AError                 `json:"error"`
	// Added for Oxxo response, maybe only available in México
	Details []struct {
		Key  string `json:"key"`
		Type string `json:"type"`
	} `json:"details"`
	PaymentData string `json:"paymentData"`
	Redirect    struct {
		Method string `json:"method"`
		URL    string `json:"url"`
	} `json:"redirect"`
}

// CheckoutAdditionalData includes additional data about the payment
type CheckoutAdditionalData struct {
	CardSummary              string `json:"cardSummary,omitempty"`
	PaymentMethod            string `json:"paymentMethod,omitempty"`
	AuthorisedAmountValue    int    `json:"authorisedAmountValue,omitempty"`
	AuthorisedAmountCurrency string `json:"authorisedAmountCurrency,omitempty"`
	Allow3DS2                bool   `json:"allow3DS2,omitempty"`
}

// AError type received
type AError struct {
	Code         int    `json:"code"`
	Message      string `json:"message"`
	RequestedURI string `json:"requested URI"`
}

func (e *AError) Error() string {
	return fmt.Sprintf("Code:%d, Message:%s, RequestedURI:%s", e.Code, e.Message, e.RequestedURI)
}

// Authorised verifies if the payment was Authorised, validating the ResultCode
func (p *PaymentResponse) Authorised() bool {
	return p.ResultCode == "Authorised"
}
