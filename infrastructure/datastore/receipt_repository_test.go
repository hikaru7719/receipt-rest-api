package datastore

import (
	"github.com/hikaru7719/receipt-rest-api/domain/model"
	"reflect"
	"testing"
)

func TestReceiptRepository_FindOne(t *testing.T) {
	db := GetConnection()
	expected := &model.Receipt{ID: 1, Name: "test", Kind: "日用品", Date: "2018-08-08", Memo: "test"}
	tx := db.Begin()
	repositpory := NewReceiptRepository(tx)
	tx.Create(&expected)
	actual, _ := repositpory.FindOne(1)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("actual %v\nwant %v", actual, expected)
	}

	defer tx.Rollback()
}
