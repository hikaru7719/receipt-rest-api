package datastore

import (
	"fmt"
	"github.com/hikaru7719/receipt-rest-api/domain/model"
	"github.com/hikaru7719/receipt-rest-api/domain/repository"
	"os"
	"reflect"
	"testing"
	"github.com/DATA-DOG/go-sqlmock"
	"database/sql"
)

var (
	TestRepository repository.ReceiptRepository
	mock sqlmock.Sqlmock
	db *sql.DB
)

func TestMain(m *testing.M) {
	db,mock,_ =sqlmock.New()
	CreateConnection(db)
	TestRepository = NewReceiptRepository()
	code := m.Run()
	os.Exit(code)
}

func TestReceiptRepository_FindOne(t *testing.T) {
	testData := &model.Receipt{ID:1,Name: "test", Kind: "日用品", Date: "2018-08-09", Memo: "test"}
	cols := []string{"id","name","kind","date","memo"}
	mock.ExpectQuery("SELECT (.+) FROM `receipts` WHERE").WillReturnRows(sqlmock.NewRows(cols).AddRow(testData.ID,testData.Name,testData.Kind,testData.Date,testData.Memo))
	actual, err := TestRepository.FindOne(1)

	if !reflect.DeepEqual(actual, testData) {
		fmt.Println(err)
		fmt.Println(actual.Date)
		t.Errorf("actual %v\nwant %v", actual, testData)
	}
}

func TestReceiptRepository_Create(t *testing.T) {
	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO `receipts`").WillReturnResult(sqlmock.NewResult(1,1))
	mock.ExpectCommit()
	expected := &model.Receipt{Name: "test", Kind: "日用品", Date: "2018-08-09", Memo: "test"}
	actual, err := TestRepository.Create(expected)

	if err != nil {
		t.Error(err)
	}

	if actual.ID == 0 {
		t.Error("could not get id")
	}
}
