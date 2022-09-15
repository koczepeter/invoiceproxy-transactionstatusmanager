package transactionStatusManager

import (
	"database/sql"
	"log"
	"os"
	"testing"

	lf "github.com/koczepeter/invoiceproxy-loggerfactory"
	_ "github.com/lib/pq"
)

var tsm TransactionStatusManager

func TestMain(m *testing.M) {

	var err error
	var db *sql.DB

	connStr := os.Getenv("APP_DB_CONNECTION_STRING")
	appLog := lf.New("", "")
	appLog.SetTestMode()

	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	tsm.Init(db, appLog)
	tsm.EnsureTableExists()
	tsm.ClearTable()

	code := m.Run()
	os.Exit(code)
}

//
//
func TestAddTransaction(t *testing.T) {

	id := "test_id"
	transactionId := "test_transaction_id"

	tsm.Add(id)
	tsm.SaveTransactionId(id, transactionId)
	result, _ := tsm.GetTransactionIdById(id)

	if result != transactionId {
		t.Fatalf("The transactionId want to be %s instead of %s", transactionId, result)
	}
}

//
//
func TestGetMissingTransactionId1(t *testing.T) {

	id := "test_id_2"
	transactionId := ""

	tsm.Add(id)
	result, createdAt := tsm.GetTransactionIdById(id)

	if result != transactionId {
		t.Fatalf("The transactionId want to be %s instead of %s", transactionId, result)
	}

	if createdAt == 0 {
		t.Fatal("The createdAt want to be greater than zero")
	}
}

//
//
func TestGetMissingTransactionId2(t *testing.T) {

	id := "test_id_3"
	transactionId := ""

	result, createdAt := tsm.GetTransactionIdById(id)

	if result != transactionId {
		t.Fatalf("The transactionId want to be %s instead of %s", transactionId, result)
	}
	if createdAt != 0 {
		t.Fatal("The createdAt want to be zero")
	}
}
