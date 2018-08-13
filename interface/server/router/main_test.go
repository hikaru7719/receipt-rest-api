package main

import (
	"testing"
	"github.com/hikaru7719/receipt-rest-api/domain/model"
	"net/http/httptest"
	"net/http"
	"encoding/json"
	"github.com/hikaru7719/receipt-rest-api/infrastructure/datastore"
	"reflect"
	"fmt"
	"strconv"
)


func TestRouter(t *testing.T){
	db := datastore.GetConnection()
	tx := db.Begin()
	testData,_ := model.NewReceipt("test","日用品", "2018-08-08","memo")
	tx.Create(testData)

	r := router()
	testServer := httptest.NewServer(r)
	client := new(http.Client)
	fmt.Println(testData.ID)
	id := strconv.Itoa(testData.ID)
	fmt.Println(id)
	req, _ := http.NewRequest("GET", testServer.URL+"/v1/receipt/"+id, nil)
	res, _ := client.Do(req)
	receipt := &model.Receipt{}
	json.NewDecoder(res.Body).Decode(receipt)

	fmt.Println(receipt)
	if !reflect.DeepEqual(testData, receipt){
		t.Error("Get Request Not Working")
	}

	tx.Rollback()
	db.Close()
}