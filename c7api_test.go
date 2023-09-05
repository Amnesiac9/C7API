package c7api

import (
	"encoding/base64"
	"fmt"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

var _ = godotenv.Load(".env")
var (
	AppAuthEncoded = "Basic " + base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s",
		os.Getenv("appid"),
		os.Getenv("appkey"),
	)))
	testTenant = os.Getenv("testTenant")
)

func TestGetJSONFromC7(t *testing.T) {

	urlString := "https://api.commerce7.com/v1/order?orderPaidDate=btw:2023-07-29T07:00:00.000Z|2023-07-31T06:59:59.999Z"
	tenant := testTenant
	goodAuth := AppAuthEncoded
	badAuth := "Basic " + base64.StdEncoding.EncodeToString([]byte("bad:auth"))

	jsonBytes, err := GetJsonFromC7(&urlString, &tenant, &goodAuth, 0)
	if err != nil {
		t.Error("Error getting JSON from C7: ", err.Error())
		return
	}

	if jsonBytes == nil {
		t.Error("JSON from C7 is empty")
		return
	}

	jsonBytes, err = GetJsonFromC7(&urlString, &tenant, &badAuth, 3)
	if err == nil {
		t.Error("Error, did not get err with bad auth: ", err.Error())
		return
	}
	fmt.Println(err.Error())

	if err.Error() != `status code: 401, error: {"statusCode":401,"type":"unauthorized","message":"Unauthenticated User","errors":[]}` {
		t.Error("Error, expected: ", `Status Code: 401, C7 Error: {"statusCode":401,"type":"unauthorized","message":"Unauthenticated User","errors":[]}`, " got: ", err.Error())
		return
	}

	if err.(C7Error).StatusCode != 401 {
		t.Error("Error, expected status code 401, got: ", err.(C7Error).StatusCode)
		return
	}

	if jsonBytes == nil {
		t.Error("JSON Bytes should not be nil with bad auth")
		return
	}

	_, err = GetJsonFromC7(nil, nil, nil, 0)
	if err == nil {
		t.Error("Error should not be nil with nil params.")
		return
	}

}

func TestPostJsonToC7(t *testing.T) {

	//urlString := "https://api.commerce7.com/v1/order"
	urlStringFulfillment := "https://api.commerce7.com/v1/order/034e6096-429d-452c-b258-5d37a1522934/fulfillment/all"
	//urlStringRemoveFulfillment := "https://api.commerce7.com/v1/order/034e6096-429d-452c-b258-5d37a1522934/fulfillment/f1439243-ceee-4ed6-b08a-4bd12f36c63e"
	tenant := testTenant
	goodAuth := AppAuthEncoded
	badAuth := "Basic " + base64.StdEncoding.EncodeToString([]byte("bad:auth"))

	blankBytes := []byte("")
	goodBytes := []byte(`{
		"sendTransactionEmail": false,
		"type": "Shipped",
		"fulfillmentDate": "2023-07-30T10:59:32.000Z",
		"shipped": {
			"trackingNumbers": ["1Z525R5EA803600000"],
			"carrier": "UPS"
		},
		"packageCount": 1
	}`)

	testCases := []struct {
		urlString     string
		tenant        string
		bytes         []byte
		auth          string
		attempts      int
		expectedCode  int
		expectedBytes []byte
	}{
		// 1 Blank Bytes
		{urlStringFulfillment, tenant, blankBytes, goodAuth, 2, 422, []byte(`{"statusCode":422,"type":"validationError","message":"One or more elements is missing or invalid","errors":[{"field":"type","message":"required"},{"field":"fulfillmentDate","message":"required"},{"field":"packageCount","message":"required"}]}`)},
		// 2 Bad Auth
		{urlStringFulfillment, tenant, goodBytes, badAuth, 2, 401, []byte(`{"statusCode":401,"type":"unauthorized","message":"Unauthenticated User","errors":[]}`)},
		// 3 Good Auth, Good Bytes
		{urlStringFulfillment, tenant, goodBytes, goodAuth, 2, 422, []byte(`{"statusCode":422,"type":"validationError","message":"Can not fulfill an order that is marked Fulfilled"}`)},
	}

	for i, testCase := range testCases {
		jsonBytes, err := PostJsonToC7(&testCase.urlString, &testCase.tenant, &testCase.bytes, &testCase.auth, testCase.attempts)
		if err.(C7Error).StatusCode != testCase.expectedCode {
			t.Error("TestPostJsonToC7, test case: ", i+1, " Expected status code: ", testCase.expectedCode, " got: ", err.(C7Error).StatusCode)
		}
		if string(*jsonBytes) != string(testCase.expectedBytes) {
			t.Error("TestPostJsonToC7, test case: ", i+1, "Expected: ", string(testCase.expectedBytes), " got: ", string(*jsonBytes))
			return
		}
	}

	// TODO: Test deleting fulfillment and posting fulfillment

}

func TestDeleteFromC7(t *testing.T) {

	urlString := "https://api.commerce7.com/v1/order/034e6096-429d-452c-b258-5d37a1522934/fulfillment/f1439243-ceee-4ed6-b08a-4bd12f36c63e"
	tenant := testTenant
	goodAuth := AppAuthEncoded
	badAuth := "Basic " + base64.StdEncoding.EncodeToString([]byte("bad:auth"))

	// Test 1
	bytes, err := DeleteFromC7(&urlString, &tenant, &badAuth, 2)
	if err == nil {
		t.Error("Error should not be nil with bad auth.")
		return
	}

	if err.(C7Error).StatusCode != 401 {
		t.Error("Status code should be 401 with bad auth, got: ", err.(C7Error).StatusCode)
		return
	}

	if bytes == nil {
		t.Error("bytes from C7 should not be nil with bad auth")
		return
	}

	// Test 2
	bytes2, err := DeleteFromC7(&urlString, &tenant, &goodAuth, 0)
	if err == nil {
		t.Error("Error should not be nil with good auth.")
		return
	}

	if err.(C7Error).StatusCode != 422 {
		t.Error("Status code should be 422 with good auth, got: ", err.(C7Error).StatusCode)
		return
	}

	if bytes2 == nil {
		t.Error("bytes2 from C7 should not be nil")
		return
	}

	_, err = DeleteFromC7(nil, nil, nil, 0)
	if err == nil {
		t.Error("Error should not be nil with nil params.")
		return
	}

}
func TestFormatDatesForC7(t *testing.T) {
	testParams := []string{"01/02/2006 15:04", "01/02/2025 15:04", "01/02/2006", ""}
	expected := []string{"2006-01-02T15:04:00.000Z", "2025-01-02T15:04:00.000Z", "01/02/2006", ""}

	for i, param := range testParams {
		result, err := FormatDatesForC7(param)
		if result != expected[i] {
			t.Error("Expected ", expected[i], " got ", result)
		}
		if err != nil {
			if err.Error() != `parsing time "01/02/2006" as "01/02/2006 15:04": cannot parse "" as "15"` && err.Error() != `date is empty` {
				t.Error("Expected error, got ", err.Error())
			}
		}
	}
}

func TestGetFulfillmentId(t *testing.T) {
	testParams := []int{1232, 1235, 999}
	expected := []string{"9475723a-8f11-4111-9234-852d85813581", "075baea0-34fe-4ac4-815d-57bd79bedf5e", ""}

	for i, param := range testParams {
		result, _ := GetFulfillmentId(param, testTenant, AppAuthEncoded)
		if result != expected[i] {
			t.Error("Expected ", expected[i], " got ", result)
		}
	}
}
