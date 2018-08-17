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
	"fmt"
	"github.com/hikaru7719/receipt-rest-api/interface/server/form"
)

func TestRouter(t *testing.T) {
	db := datastore.GetConnection()
	tx := db.Begin()
	testData, _ := model.NewReceipt("test", "日用品", "2018-08-08", "memo")
	tx.Create(testData)

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
	jsonStr := `{"name":"test","kind":"日用品","date":"2018-08-09","memo":"memo"}`

	//buf := new(bytes.Buffer)
	//json.NewEncoder(buf).Encode(test)
	//fmt.Println(buf)
	//postReq, _ := http.NewRequest("POST",testServer.URL+"/v1/receipt",buf)
	postReq, _ := http.NewRequest("POST",testServer.URL+"/v1/receipt",bytes.NewBuffer([]byte(jsonStr)))
	postReq.Header.Set("Content-Type","application/json")
	postRes, _ := client.Do(postReq)

	fmt.Println(req)
	fmt.Println(postReq)
	fmt.Println(postRes.StatusCode)

	if postRes.StatusCode != 201 {
		t.Error("Post Request Not Working")
	}

	//defer tx.Rollback()
	//defer db.Close()
}
