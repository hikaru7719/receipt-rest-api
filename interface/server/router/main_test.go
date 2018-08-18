package main

import (
	"github.com/hikaru7719/receipt-rest-api/domain/model"
	"github.com/hikaru7719/receipt-rest-api/infrastructure/datastore"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"encoding/json"
	"bytes"
	"github.com/hikaru7719/receipt-rest-api/interface/server/form"
)

func TestRouter(t *testing.T) {
	datastore.CreateConnection()
	testData, _ := model.NewReceipt("test", "日用品", "2018-08-08", "memo")
	datastore.DB.Create(testData)
	r := router()
	testServer := httptest.NewServer(r)
	client := new(http.Client)
	id := strconv.Itoa(testData.ID)
	req, _ := http.NewRequest("GET", testServer.URL+"/v1/receipt/"+id, nil)
	res, _ := client.Do(req)

	if res.StatusCode != 200 {
		t.Error("Get Request Not Working")
	}

	test := &form.ReceiptForm{Name:"test",Kind:"日用品",Date:"2018-08-08",Memo:"test"}
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(test)
	postReq, _ := http.NewRequest("POST",testServer.URL+"/v1/receipt",buf)
	postReq.Header.Set("Content-Type","application/json")
	postRes, _ := client.Do(postReq)

	if postRes.StatusCode != 201 {
		t.Error("Post Request Not Working")
	}

}
