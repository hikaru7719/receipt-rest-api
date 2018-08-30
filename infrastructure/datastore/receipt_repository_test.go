package datastore

import (
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/hikaru7719/receipt-rest-api/domain/model"
	"reflect"
	"testing"
)

func TestReceiptRepository_FindOne(t *testing.T) {
	testData := &model.Receipt{ID: 1, UserID: 1, Name: "test", Price: 1000, Kind: "日用品", Date: "2018-08-09", Memo: "test", CreditID: 1}
	cols := []string{"id", "user_id", "name", "price", "kind", "date", "memo", "credit_id"}
	mock.ExpectQuery("SELECT (.+) FROM `receipts` WHERE").WillReturnRows(sqlmock.NewRows(cols).AddRow(testData.ID, testData.UserID, testData.Name, testData.Price, testData.Kind, testData.Date, testData.Memo, testData.CreditID))
	actual, err := TestReceiptRepository.FindOne(1)

	if !reflect.DeepEqual(actual, testData) {
		fmt.Println(err)
		fmt.Println(actual.Date)
		t.Errorf("actual %v\nwant %v", actual, testData)
	}
}

func TestReceiptRepository_Create(t *testing.T) {
	mock.ExpectExec("INSERT INTO `receipts`").WillReturnResult(sqlmock.NewResult(1, 1))
	expected := &model.Receipt{UserID: 1, Name: "test", Kind: "日用品", Date: "2018-08-09", Memo: "test", CreditID: 1}
	actual, err := TestReceiptRepository.Create(expected)

	if err != nil {
		t.Error(err)
	}

	if actual.ID == 0 {
		t.Error("could not get id")
	}
}

func TestReceiptRepository_Delete(t *testing.T) {
	mock.ExpectExec("DELETE FROM `receipts`").WillReturnResult(sqlmock.NewResult(1, 1))

	err := TestReceiptRepository.Delete(1000)

	if err != nil {
		t.Error(err)
	}
}
