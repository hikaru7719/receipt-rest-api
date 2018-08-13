package main

import (
	"github.com/hikaru7719/receipt-rest-api/domain/model"
	"github.com/hikaru7719/receipt-rest-api/infrastructure/datastore"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
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

	defer tx.Rollback()
	defer db.Close()
}
