package datastore

import (
	"github.com/hikaru7719/receipt-rest-api/domain/model"
	"github.com/hikaru7719/receipt-rest-api/domain/repository"
	"github.com/jinzhu/gorm"
	"os"
	"reflect"
	"testing"
)

var (
	db             *gorm.DB
	tx             *gorm.DB
	TestRepository repository.ReceiptRepository
)

func TestMain(m *testing.M) {
	db = GetConnection()
	tx = db.Begin()
	TestRepository = NewReceiptRepository(tx)
	code := m.Run()
	tx.Rollback()
	os.Exit(code)
}

func TestReceiptRepository_FindOne(t *testing.T) {
	expected := &model.Receipt{ID: 1, Name: "test", Kind: "日用品", Date: "2018-08-08", Memo: "test"}
	tx.Create(&expected)
	actual, _ := TestRepository.FindOne(1)

	if !reflect.DeepEqual(actual, expected) {
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
