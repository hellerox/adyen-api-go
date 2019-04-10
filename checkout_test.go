package adyen

import (
	"encoding/json"
	"testing"
)

func TestPaymentMethodsResponse_ParseMerchantAccount(t *testing.T) {
	rawResponse := `{"paymentMethods":[{"details":[{"key":"additionalData.card.encrypted.json","type":"cardToken"}],"name":"Credit Card","type":"scheme"},{"details":[{"items":[{"id":"1121","name":"Test Issuer"},{"id":"1154","name":"Test Issuer 5"},{"id":"1153","name":"Test Issuer 4"},{"id":"1152","name":"Test Issuer 3"},{"id":"1151","name":"Test Issuer 2"},{"id":"1162","name":"Test Issuer Cancelled"},{"id":"1161","name":"Test Issuer Pending"},{"id":"1160","name":"Test Issuer Refused"},{"id":"1159","name":"Test Issuer 10"},{"id":"1158","name":"Test Issuer 9"},{"id":"1157","name":"Test Issuer 8"},{"id":"1156","name":"Test Issuer 7"},{"id":"1155","name":"Test Issuer 6"}],"key":"idealIssuer","type":"select"}],"name":"iDEAL","type":"ideal"},{"name":"Pay later with Klarna.","type":"klarna"},{"details":[{"key":"sepa.ownerName","type":"text"},{"key":"sepa.ibanNumber","type":"text"}],"name":"SEPA Direct Debit","type":"sepadirectdebit"},{"name":"UnionPay","type":"unionpay"}]}`

	var response PaymentMethodsResponse
	if err := json.Unmarshal([]byte(rawResponse), &response); err != nil {
		t.Fatalf("error unmarshalling json: %v", err)
	}

	exp := PaymentMethodsResponse{
		PaymentMethods: []PaymentMethodDetails{
			{
				Details: []PaymentMethodDetailsInfo{
					{
						Key:  "additionalData.card.encrypted.json",
						Type: "cardToken"}},
				Name: "Credit Card",
				Type: "scheme"},
			{
				Details: []PaymentMethodDetailsInfo{
					{
						Items: []PaymentMethodItems{
							{ID: "1121", Name: "Test Issuer"},
							{ID: "1154", Name: "Test Issuer 5"},
							{ID: "1153", Name: "Test Issuer 4"},
							{ID: "1152", Name: "Test Issuer 3"},
							{ID: "1151", Name: "Test Issuer 2"},
							{ID: "1162", Name: "Test Issuer Cancelled"},
							{ID: "1161", Name: "Test Issuer Pending"},
							{ID: "1160", Name: "Test Issuer Refused"},
							{ID: "1159", Name: "Test Issuer 10"},
							{ID: "1158", Name: "Test Issuer 9"},
							{ID: "1157", Name: "Test Issuer 8"},
							{ID: "1156", Name: "Test Issuer 7"},
							{ID: "1155", Name: "Test Issuer 6"}},
						Key: "idealIssuer", Type: "select"},
				},
				Name: "iDEAL",
				Type: "ideal",
			},
			{
				Name: "Pay later with Klarna.",
				Type: "klarna"},
			{
				Details: []PaymentMethodDetailsInfo{
					{
						Key:  "sepa.ownerName",
						Type: "text",
					},
					{
						Key:  "sepa.ibanNumber",
						Type: "text",
					},
				},
				Name: "SEPA Direct Debit",
				Type: "sepadirectdebit",
			},
			{
				Name: "UnionPay",
				Type: "unionpay",
			},
		},
	}

	equals(t, exp, response)
}

func TestPaymentMethodsResponse_ParseCountryAmount(t *testing.T) {
	rawResponse := `{"paymentMethods":[{"details":[{"items":[{"id":"1121","name":"Test Issuer"},{"id":"1154","name":"Test Issuer 5"},{"id":"1153","name":"Test Issuer 4"},{"id":"1152","name":"Test Issuer 3"},{"id":"1151","name":"Test Issuer 2"},{"id":"1162","name":"Test Issuer Cancelled"},{"id":"1161","name":"Test Issuer Pending"},{"id":"1160","name":"Test Issuer Refused"},{"id":"1159","name":"Test Issuer 10"},{"id":"1158","name":"Test Issuer 9"},{"id":"1157","name":"Test Issuer 8"},{"id":"1156","name":"Test Issuer 7"},{"id":"1155","name":"Test Issuer 6"}],"key":"idealIssuer","type":"select"}],"name":"iDEAL","type":"ideal"},{"details":[{"key":"additionalData.card.encrypted.json","type":"cardToken"}],"name":"Credit Card","type":"scheme"},{"name":"Pay later with Klarna.","type":"klarna"},{"details":[{"key":"sepa.ownerName","type":"text"},{"key":"sepa.ibanNumber","type":"text"}],"name":"SEPA Direct Debit","type":"sepadirectdebit"},{"name":"UnionPay","type":"unionpay"}]}`

	var response PaymentMethodsResponse
	if err := json.Unmarshal([]byte(rawResponse), &response); err != nil {
		t.Fatalf("error unmarshalling json: %v", err)
	}

	exp := PaymentMethodsResponse{
		PaymentMethods: []PaymentMethodDetails{
			{
				Details: []PaymentMethodDetailsInfo{
					{
						Items: []PaymentMethodItems{
							{ID: "1121", Name: "Test Issuer"},
							{ID: "1154", Name: "Test Issuer 5"},
							{ID: "1153", Name: "Test Issuer 4"},
							{ID: "1152", Name: "Test Issuer 3"},
							{ID: "1151", Name: "Test Issuer 2"},
							{ID: "1162", Name: "Test Issuer Cancelled"},
							{ID: "1161", Name: "Test Issuer Pending"},
							{ID: "1160", Name: "Test Issuer Refused"},
							{ID: "1159", Name: "Test Issuer 10"},
							{ID: "1158", Name: "Test Issuer 9"},
							{ID: "1157", Name: "Test Issuer 8"},
							{ID: "1156", Name: "Test Issuer 7"},
							{ID: "1155", Name: "Test Issuer 6"},
						},
						Key:  "idealIssuer",
						Type: "select",
					},
				},
				Name: "iDEAL",
				Type: "ideal",
			},
			{
				Details: []PaymentMethodDetailsInfo{
					{
						Key:  "additionalData.card.encrypted.json",
						Type: "cardToken"},
				},
				Name: "Credit Card",
				Type: "scheme",
			},
			{
				Name: "Pay later with Klarna.",
				Type: "klarna",
			},
			{
				Details: []PaymentMethodDetailsInfo{
					{
						Key:  "sepa.ownerName",
						Type: "text",
					},
					{
						Key:  "sepa.ibanNumber",
						Type: "text",
					},
				},
				Name: "SEPA Direct Debit",
				Type: "sepadirectdebit",
			},
			{
				Name: "UnionPay",
				Type: "unionpay",
			},
		},
	}

	equals(t, exp, response)
}

func TestPaymentMethodsResponse_ParseOneClick(t *testing.T) {
	rawResponse := `{"oneClickPaymentMethods":[{"details":[{"key":"cardDetails.cvc","type":"cvc"}],"name":"VISA","type":"visa","storedDetails":{"card":{"expiryMonth":"8","expiryYear":"2018","holderName":"John Smith","number":"1111"}}}],"paymentMethods":[{"details":[{"items":[{"id":"1121","name":"Test Issuer"},{"id":"1154","name":"Test Issuer 5"},{"id":"1153","name":"Test Issuer 4"},{"id":"1152","name":"Test Issuer 3"},{"id":"1151","name":"Test Issuer 2"},{"id":"1162","name":"Test Issuer Cancelled"},{"id":"1161","name":"Test Issuer Pending"},{"id":"1160","name":"Test Issuer Refused"},{"id":"1159","name":"Test Issuer 10"},{"id":"1158","name":"Test Issuer 9"},{"id":"1157","name":"Test Issuer 8"},{"id":"1156","name":"Test Issuer 7"},{"id":"1155","name":"Test Issuer 6"}],"key":"idealIssuer","type":"select"}],"name":"iDEAL","type":"ideal"},{"details":[{"key":"additionalData.card.encrypted.json","type":"cardToken"},{"key":"storeDetails","optional":"true","type":"boolean"}],"name":"Credit Card","type":"scheme"},{"name":"Pay later with Klarna.","type":"klarna"},{"details":[{"key":"sepa.ownerName","type":"text"},{"key":"sepa.ibanNumber","type":"text"}],"name":"SEPA Direct Debit","type":"sepadirectdebit"},{"name":"UnionPay","type":"unionpay"}]}`

	var response PaymentMethodsResponse
	if err := json.Unmarshal([]byte(rawResponse), &response); err != nil {
		t.Fatalf("error unmarshalling json: %v", err)
	}

	exp := PaymentMethodsResponse{
		PaymentMethods: []PaymentMethodDetails{
			{
				Details: []PaymentMethodDetailsInfo{
					{
						Items: []PaymentMethodItems{
							{ID: "1121", Name: "Test Issuer"},
							{ID: "1154", Name: "Test Issuer 5"},
							{ID: "1153", Name: "Test Issuer 4"},
							{ID: "1152", Name: "Test Issuer 3"},
							{ID: "1151", Name: "Test Issuer 2"},
							{ID: "1162", Name: "Test Issuer Cancelled"},
							{ID: "1161", Name: "Test Issuer Pending"},
							{ID: "1160", Name: "Test Issuer Refused"},
							{ID: "1159", Name: "Test Issuer 10"},
							{ID: "1158", Name: "Test Issuer 9"},
							{ID: "1157", Name: "Test Issuer 8"},
							{ID: "1156", Name: "Test Issuer 7"},
							{ID: "1155", Name: "Test Issuer 6"},
						},
						Key:  "idealIssuer",
						Type: "select",
					},
				},
				Name: "iDEAL",
				Type: "ideal"},
			{
				Details: []PaymentMethodDetailsInfo{
					{
						Key:  "additionalData.card.encrypted.json",
						Type: "cardToken",
					},
					{
						Key:  "storeDetails",
						Type: "boolean",
					},
				},
				Name: "Credit Card",
				Type: "scheme",
			},
			{
				Name: "Pay later with Klarna.",
				Type: "klarna",
			},
			{
				Details: []PaymentMethodDetailsInfo{
					{
						Key:  "sepa.ownerName",
						Type: "text",
					},
					{
						Key:  "sepa.ibanNumber",
						Type: "text",
					},
				},
				Name: "SEPA Direct Debit",
				Type: "sepadirectdebit",
			},
			{
				Name: "UnionPay",
				Type: "unionpay",
			},
		},
		OneClickPaymentMethods: []OneClickPaymentMethodDetails{
			{
				Details: []PaymentMethodTypes{
					{
						Key:  "cardDetails.cvc",
						Type: "cvc",
					},
				},
				Name: "VISA",
				Type: "visa",
				StoredDetails: PaymentMethodStoredDetails{
					Card: PaymentMethodCard{
						ExpiryMonth: "8",
						ExpiryYear:  "2018",
						HolderName:  "John Smith",
						Number:      "1111",
					},
				},
			},
		},
	}

	equals(t, exp, response)
}
