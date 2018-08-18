package datastore

import (
	"fmt"
	"github.com/hikaru7719/receipt-rest-api/domain/model"
	"github.com/hikaru7719/receipt-rest-api/domain/repository"
	"os"
	"reflect"
	"testing"
)

var (
	TestRepository repository.ReceiptRepository
)

func TestMain(m *testing.M) {
	CreateConnection(GetDBEnv())
	TestRepository = NewReceiptRepository()
	code := m.Run()
	os.Exit(code)
}

func TestReceiptRepository_FindOne(t *testing.T) {
	expected := &model.Receipt{Name: "test", Kind: "日用品", Date: "2018-08-08", Memo: "test"}
	DB.Create(&expected)
	fmt.Println(expected.ID)
	actual, err := TestRepository.FindOne(expected.ID)

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
