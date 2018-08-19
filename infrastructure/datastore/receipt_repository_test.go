package datastore

import (
	"github.com/hikaru7719/receipt-rest-api/domain/model"
	"github.com/hikaru7719/receipt-rest-api/domain/repository"
	"os"
	"reflect"
	"testing"
	"fmt"
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
	expected := &model.Receipt{Name: "test", Kind: "日用品", Date: "2018-08-08", Memo: "test"}
	actual, err := TestRepository.FindOne(1)

	if !reflect.DeepEqual(actual, expected) {
		fmt.Println(err)
		t.Errorf("actual %v\nwant %v", actual, expected)
	}
}

func TestReceiptRepository_Create(t *testing.T) {
	expected := &model.Receipt{Name: "test", Kind: "日用品", Date: "2018-08-09", Memo: "test"}
	actual, err := TestRepository.Create(expected)

	if err != nil {
		t.Error(err)
	}

	if actual.ID == 0 {
		t.Error("could not get id")
	}
}
