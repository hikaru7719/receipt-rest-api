package datastore

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/hikaru7719/receipt-rest-api/domain/model"
	"testing"
)

func TestCreditRepository_Create(t *testing.T) {
	mock.ExpectExec("INSERT INTO `credits`").WillReturnResult(sqlmock.NewResult(1, 1))
	expected := &model.Credit{ID: 1, UserID: 1, CardName: "アメリカンエクスプレス", FinishDate: 10, WithdrawalDate: 4, LaterPaymentMonth: 1}
	actual, err := TestCreditRepository.Create(expected)

	if err != nil {
		t.Error(err)
	}

	if actual.ID == 0 {
		t.Error("could not get id")
	}
}
