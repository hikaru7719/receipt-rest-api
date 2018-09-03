package datastore

import (
	"github.com/hikaru7719/receipt-rest-api/domain/model"
	"reflect"
	"testing"
)

func TestCreditRepository_Create(t *testing.T) {
	CreateConnection(GetDBEnv())
	testData := &model.Credit{UserID: 1, CardName: "アメリカンエクスプレス", FinishDate: 10, WithdrawalDate: 4, LaterPaymentMonth: 1}
	expected, _ := TestCreditRepository.Create(testData)
	actual := &model.Credit{}
	DB.Where("id=?", expected.ID).Find(actual)
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("actual %v,want %v", actual, expected)
	}
	DB.Delete(actual)
}
