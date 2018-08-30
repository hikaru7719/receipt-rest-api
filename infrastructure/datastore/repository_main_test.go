package datastore

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/hikaru7719/receipt-rest-api/domain/repository"
	"os"
	"testing"
)

var (
	TestReceiptRepository repository.ReceiptRepository
	TestCreditRepository  repository.CreditRepository
	mock                  sqlmock.Sqlmock
	db                    *sql.DB
)

func TestMain(m *testing.M) {
	db, mock, _ = sqlmock.New()
	CreateConnection(db)
	TestReceiptRepository = NewReceiptRepository()
	TestCreditRepository = NewCreditRepository()
	code := m.Run()
	os.Exit(code)
}
