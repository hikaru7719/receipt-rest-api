package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/hikaru7719/receipt-rest-api/domain/model"
	"github.com/hikaru7719/receipt-rest-api/infrastructure/datastore"
	"github.com/hikaru7719/receipt-rest-api/interface/server/form"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestAPI(t *testing.T) {
	datastore.CreateConnection(datastore.GetDBEnv())
	r := router()
	testServer := httptest.NewServer(r)
	client := new(http.Client)
	testGetReceipt(t, client, testServer)
	testPostReceipt(t, client, testServer)
	testDeleteReceipt(t, client, testServer)
	testPostCredit(t, client, testServer)
}

func testGetReceipt(t *testing.T, client *http.Client, testServer *httptest.Server) {
	testData, _ := model.NewReceipt(1, "test", 1000, "日用品", "2018-08-08", "memo", 1)
	datastore.DB.Create(testData)
	id := strconv.Itoa(testData.ID)
	req, _ := http.NewRequest("GET", testServer.URL+"/v1/receipts/"+id, nil)
	res, _ := client.Do(req)

	if res.StatusCode != 200 {
		t.Error("Get Request /v1/receipts/:id Not Working")
	}

	body, _ := ioutil.ReadAll(res.Body)
	actual := string(body)
	fmt.Println(actual)
}

func testPostReceipt(t *testing.T, client *http.Client, testServer *httptest.Server) {
	testData := &form.ReceiptForm{UserID: 2, Name: "test", Price: 1000, Kind: "日用品", Date: "2018-08-08", Memo: "test", CreditID: 2}
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(testData)
	req, _ := http.NewRequest("POST", testServer.URL+"/v1/receipts", buf)
	req.Header.Set("Content-Type", "application/json")
	res, _ := client.Do(req)

	if res.StatusCode != 201 {
		t.Error("Post Request /v1/receipts Not Working")
	}
	body, _ := ioutil.ReadAll(res.Body)
	actual := string(body)
	fmt.Println(actual)
}

func testDeleteReceipt(t *testing.T, client *http.Client, testServer *httptest.Server) {
	testData, _ := model.NewReceipt(3, "test", 1000, "日用品", "2018-08-08", "memo", 1)
	datastore.DB.Create(testData)
	id := strconv.Itoa(testData.ID)
	req, _ := http.NewRequest("DELETE", testServer.URL+"/v1/receipts/"+id, nil)
	res, _ := client.Do(req)

	if res.StatusCode != 202 {
		t.Error("Delete Request /v1/receipts/:id Not Working")
	}
	body, _ := ioutil.ReadAll(res.Body)
	actual := string(body)
	fmt.Println(actual)

}

func testPostCredit(t *testing.T, client *http.Client, testServer *httptest.Server) {
	testData := &form.CreditForm{UserID: 1, CardName: "アメリカンエクスプレス", FinishDate: 10, WithdrawalDate: 4, LaterPaymentMonth: 1}
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(testData)
	req, _ := http.NewRequest("POST", testServer.URL+"/v1/credits", buf)
	req.Header.Set("Content-Type", "application/json")
	res, _ := client.Do(req)

	if res.StatusCode != 201 {
		t.Error("Post Request /v1/credits Not Working")
	}
	body, _ := ioutil.ReadAll(res.Body)
	actual := string(body)
	fmt.Println(actual)
}
