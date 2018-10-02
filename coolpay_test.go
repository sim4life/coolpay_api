package main

import (
	"fmt"
	"net/http"
	//"log"
	//"io/ioutil"
	//"encoding/json"
	"testing"
)

func TestUserAuth(t *testing.T) {
	credentials := &Credentials{"ShahzadI", "CDA8777865C7CC3C"}
	token, httpStatusCode := getAuthToken(*credentials, authURL)
	if http.StatusOK != httpStatusCode {
		t.Errorf("User Authentication returned http StatusCode was incorrect, got: %d, want: %d.", httpStatusCode, http.StatusOK)
	}
	fmt.Println("Token = ", token)
	
}
func TestWrongUserAuthUsername(t *testing.T) {
	credentials := &Credentials{"Shahzad", "CDA8777865C7CC3C"}
	token, httpStatusCode := getAuthToken(*credentials, authURL)
	if http.StatusNotFound != httpStatusCode {
		t.Errorf("User Authentication returned http StatusCode was incorrect, got: %d, want: %d.", httpStatusCode, http.StatusNotFound)
	}
	fmt.Println("Token = ", token)
	
}
func TestWrongUserAuthApikey(t *testing.T) {
	credentials := &Credentials{"ShahzadI", "CDA8777865C7CC3D"}
	token, httpStatusCode := getAuthToken(*credentials, authURL)
	if http.StatusNotFound != httpStatusCode {
		t.Errorf("User Authentication returned http StatusCode was incorrect, got: %d, want: %d.", httpStatusCode, http.StatusNotFound)
	}
	fmt.Println("Token = ", token)
	
}

func TestAddRecipient(t *testing.T) {
	credentials := &Credentials{"ShahzadI", "CDA8777865C7CC3C"}
	token, httpStatusCode := getAuthToken(*credentials, authURL)
	if http.StatusOK != httpStatusCode {
		t.Errorf("User Authentication returned http StatusCode was incorrect, got: %d, want: %d.", httpStatusCode, http.StatusOK)
	}
	recipientInfo := &RecipientInfo{Name:"Will Cox"}
	returnedRecipientInfo, httpStatusCode := addRecipient(*recipientInfo, token, addRecipientURL)
	if http.StatusCreated != httpStatusCode {
		t.Errorf("Add recipient returned http StatusCode was incorrect, got: %d, want: %d.", httpStatusCode, http.StatusCreated)
	}
	// fmt.Printf("Returned recipientInfo Id is:%s\n", returnedRecipientInfo.Id)
	if recipientInfo.Name != returnedRecipientInfo.Name {
		t.Errorf("Add recipient returned Recipient Name was incorrect, got: %s, want: %s.", returnedRecipientInfo.Name, recipientInfo.Name)
	}
}
func TestAddRecipientWrongToken(t *testing.T) {
	wrongToken := Token("")
	recipientInfo := &RecipientInfo{Name:"Will Cox"}
	_, httpStatusCode := addRecipient(*recipientInfo, &wrongToken, addRecipientURL)
	if http.StatusInternalServerError != httpStatusCode {
		t.Errorf("Add recipient returned http StatusCode was incorrect, got: %d, want: %d.", httpStatusCode, http.StatusInternalServerError)
	}
}

func TestMakePayment(t *testing.T) {
	credentials := &Credentials{"ShahzadI", "CDA8777865C7CC3C"}
	token, httpStatusCode := getAuthToken(*credentials, authURL)
	if http.StatusOK != httpStatusCode {
		t.Errorf("User Authentication returned http StatusCode was incorrect, got: %d, want: %d.", httpStatusCode, http.StatusOK)
	}
	recipientInfo := &RecipientInfo{Name:"Brian Adams"}
	returnedRecipientInfo, httpStatusCode := addRecipient(*recipientInfo, token, addRecipientURL)
	if http.StatusCreated != httpStatusCode {
		t.Errorf("Add recipient returned http StatusCode was incorrect, got: %d, want: %d.", httpStatusCode, http.StatusCreated)
	}

	paymentInfo := &PaymentInfo{Amount:12.5, Currency:"GBP", RecipientId:returnedRecipientInfo.Id}
	returnedPaymentInfo, httpStatusCode := makePaymentToRecipient(*paymentInfo, token, makePaymentURL)
	if http.StatusCreated != httpStatusCode {
		t.Errorf("Make payment returned http StatusCode was incorrect, got: %d, want: %d.", httpStatusCode, http.StatusCreated)
	}
	// fmt.Printf("Returned paymentInfo Id is:%s\n", returnedPaymentInfo.Id)
	if paymentInfo.Amount != returnedPaymentInfo.Amount {
		t.Errorf("Make payment returned Payment Amount was incorrect, got: %f, want: %f.", returnedPaymentInfo.Amount, paymentInfo.Amount)
	}
	if paymentInfo.Currency != returnedPaymentInfo.Currency {
		t.Errorf("Make payment returned Payment Currency was incorrect, got: %s, want: %s.", returnedPaymentInfo.Currency, paymentInfo.Currency)
	}
	if paymentInfo.RecipientId != returnedPaymentInfo.RecipientId {
		t.Errorf("Make payment returned Payment RecipientId was incorrect, got: %s, want: %s.", returnedPaymentInfo.RecipientId, paymentInfo.RecipientId)
	}
}
func TestMakePaymentWrongToken(t *testing.T) {
	credentials := &Credentials{"ShahzadI", "CDA8777865C7CC3C"}
	token, httpStatusCode := getAuthToken(*credentials, authURL)
	if http.StatusOK != httpStatusCode {
		t.Errorf("User Authentication returned http StatusCode was incorrect, got: %d, want: %d.", httpStatusCode, http.StatusOK)
	}
	recipientInfo := &RecipientInfo{Name:"Brian Adams"}
	returnedRecipientInfo, httpStatusCode := addRecipient(*recipientInfo, token, addRecipientURL)
	if http.StatusCreated != httpStatusCode {
		t.Errorf("Add recipient returned http StatusCode was incorrect, got: %d, want: %d.", httpStatusCode, http.StatusCreated)
	}
	
	wrongToken := Token("")
	paymentInfo := &PaymentInfo{Amount:12.5, Currency:"GBP", RecipientId:returnedRecipientInfo.Id}
	_, httpStatusCode = makePaymentToRecipient(*paymentInfo, &wrongToken, makePaymentURL)
	if http.StatusInternalServerError != httpStatusCode {
		t.Errorf("Make payment returned http StatusCode was incorrect, got: %d, want: %d.", httpStatusCode, http.StatusInternalServerError)
	}
}
func TestMakePaymentWrongRecipient(t *testing.T) {
	credentials := &Credentials{"ShahzadI", "CDA8777865C7CC3C"}
	token, httpStatusCode := getAuthToken(*credentials, authURL)
	if http.StatusOK != httpStatusCode {
		t.Errorf("User Authentication returned http StatusCode was incorrect, got: %d, want: %d.", httpStatusCode, http.StatusOK)
	}

	paymentInfo := &PaymentInfo{Amount:12.5, Currency:"GBP", RecipientId:""}
	_, httpStatusCode = makePaymentToRecipient(*paymentInfo, token, makePaymentURL)
	if http.StatusUnprocessableEntity != httpStatusCode {
		t.Errorf("Make payment returned http StatusCode was incorrect, got: %d, want: %d.", httpStatusCode, http.StatusUnprocessableEntity)
	}
}

func TestVerifyPayment(t *testing.T) {
	credentials := &Credentials{"ShahzadI", "CDA8777865C7CC3C"}
	token, httpStatusCode := getAuthToken(*credentials, authURL)
	if http.StatusOK != httpStatusCode {
		t.Errorf("User Authentication returned http StatusCode was incorrect, got: %d, want: %d.", httpStatusCode, http.StatusOK)
	}
	recipientInfo := &RecipientInfo{Name:"Phil Collins"}
	returnedRecipientInfo, httpStatusCode := addRecipient(*recipientInfo, token, addRecipientURL)
	if http.StatusCreated != httpStatusCode {
		t.Errorf("Add recipient returned http StatusCode was incorrect, got: %d, want: %d.", httpStatusCode, http.StatusCreated)
	}
	paymentInfo := &PaymentInfo{Amount:12.5, Currency:"GBP", RecipientId:returnedRecipientInfo.Id}
	returnedPaymentInfo, httpStatusCode := makePaymentToRecipient(*paymentInfo, token, makePaymentURL)
	if http.StatusCreated != httpStatusCode {
		t.Errorf("Make payment returned http StatusCode was incorrect, got: %d, want: %d.", httpStatusCode, http.StatusCreated)
	}

	status, httpStatusCode := verifyPaymentToRecipient(*returnedPaymentInfo, token, listPaymentsURL)
	if http.StatusOK != httpStatusCode {
		t.Errorf("List all payments returned http StatusCode was incorrect, got: %d, want: %d.", httpStatusCode, http.StatusOK)
	}
	if "paid" == status {
		fmt.Println("\nPayment VERIFIED")
	} else {
		fmt.Printf("\nPayment NOT Verified with status:%s\n", status)
	}
}

func TestVerifyPaymentWrongToken(t *testing.T) {
	credentials := &Credentials{"ShahzadI", "CDA8777865C7CC3C"}
	token, httpStatusCode := getAuthToken(*credentials, authURL)
	if http.StatusOK != httpStatusCode {
		t.Errorf("User Authentication returned http StatusCode was incorrect, got: %d, want: %d.", httpStatusCode, http.StatusOK)
	}
	recipientInfo := &RecipientInfo{Name:"Phil Collins"}
	returnedRecipientInfo, httpStatusCode := addRecipient(*recipientInfo, token, addRecipientURL)
	if http.StatusCreated != httpStatusCode {
		t.Errorf("Add recipient returned http StatusCode was incorrect, got: %d, want: %d.", httpStatusCode, http.StatusCreated)
	}
	paymentInfo := &PaymentInfo{Amount:12.5, Currency:"GBP", RecipientId:returnedRecipientInfo.Id}
	returnedPaymentInfo, httpStatusCode := makePaymentToRecipient(*paymentInfo, token, makePaymentURL)
	if http.StatusCreated != httpStatusCode {
		t.Errorf("Make payment returned http StatusCode was incorrect, got: %d, want: %d.", httpStatusCode, http.StatusCreated)
	}

	wrongToken := Token("")
	_, httpStatusCode = verifyPaymentToRecipient(*returnedPaymentInfo, &wrongToken, listPaymentsURL)
	if http.StatusInternalServerError != httpStatusCode {
		t.Errorf("List all payments returned http StatusCode was incorrect, got: %d, want: %d.", httpStatusCode, http.StatusInternalServerError)
	}
}
func TestVerifyPaymentWrongPaymentId(t *testing.T) {
	returnedPaymentInfo, token := getPaymentInfo(t)
	returnedPaymentInfo.Id = ""
	status, httpStatusCode := verifyPaymentToRecipient(returnedPaymentInfo, token, listPaymentsURL)
	if http.StatusOK != httpStatusCode {
		t.Errorf("List all payments returned http StatusCode was incorrect, got: %d, want: %d.", httpStatusCode, http.StatusOK)
	}
	if "NOT found" != status {
		t.Errorf("List all payments returned http StatusCode was incorrect, got: %s, want: %s.", status, "NOT found")
	}
}
func TestVerifyPaymentWrongPaymentAmount(t *testing.T) {
	returnedPaymentInfo, token := getPaymentInfo(t)
	returnedPaymentInfo.Amount = 10
	status, httpStatusCode := verifyPaymentToRecipient(returnedPaymentInfo, token, listPaymentsURL)
	if http.StatusOK != httpStatusCode {
		t.Errorf("List all payments returned http StatusCode was incorrect, got: %d, want: %d.", httpStatusCode, http.StatusOK)
	}
	if "Payment credentials DONOT match" != status {
		t.Errorf("List all payments returned http StatusCode was incorrect, got: %s, want: %s.", status, "Payment credentials DONOT match")
	}
}
func TestVerifyPaymentWrongPaymentCurrency(t *testing.T) {
	returnedPaymentInfo, token := getPaymentInfo(t)
	returnedPaymentInfo.Currency = "USD"
	status, httpStatusCode := verifyPaymentToRecipient(returnedPaymentInfo, token, listPaymentsURL)
	if http.StatusOK != httpStatusCode {
		t.Errorf("List all payments returned http StatusCode was incorrect, got: %d, want: %d.", httpStatusCode, http.StatusOK)
	}
	if "Payment credentials DONOT match" != status {
		t.Errorf("List all payments returned http StatusCode was incorrect, got: %s, want: %s.", status, "Payment credentials DONOT match")
	}
}
func TestVerifyPaymentWrongPaymentRecipientId(t *testing.T) {
	returnedPaymentInfo, token := getPaymentInfo(t)
	returnedPaymentInfo.RecipientId = ""
	status, httpStatusCode := verifyPaymentToRecipient(returnedPaymentInfo, token, listPaymentsURL)
	if http.StatusOK != httpStatusCode {
		t.Errorf("List all payments returned http StatusCode was incorrect, got: %d, want: %d.", httpStatusCode, http.StatusOK)
	}
	if "Payment credentials DONOT match" != status {
		t.Errorf("List all payments returned http StatusCode was incorrect, got: %s, want: %s.", status, "Payment credentials DONOT match")
	}
}

func getPaymentInfo(t *testing.T) (PaymentInfo, *Token) {
	credentials := &Credentials{"ShahzadI", "CDA8777865C7CC3C"}
	token, httpStatusCode := getAuthToken(*credentials, authURL)
	if http.StatusOK != httpStatusCode {
		t.Errorf("User Authentication returned http StatusCode was incorrect, got: %d, want: %d.", httpStatusCode, http.StatusOK)
	}
	recipientInfo := &RecipientInfo{Name:"Phil Collins"}
	returnedRecipientInfo, httpStatusCode := addRecipient(*recipientInfo, token, addRecipientURL)
	if http.StatusCreated != httpStatusCode {
		t.Errorf("Add recipient returned http StatusCode was incorrect, got: %d, want: %d.", httpStatusCode, http.StatusCreated)
	}
	paymentInfo := &PaymentInfo{Amount:12.5, Currency:"GBP", RecipientId:returnedRecipientInfo.Id}
	returnedPaymentInfo, httpStatusCode := makePaymentToRecipient(*paymentInfo, token, makePaymentURL)
	if http.StatusCreated != httpStatusCode {
		t.Errorf("Make payment returned http StatusCode was incorrect, got: %d, want: %d.", httpStatusCode, http.StatusCreated)
	}
	return *returnedPaymentInfo, token
}
