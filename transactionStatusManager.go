package transactionStatusManager

import (
	"database/sql"
	"fmt"
	"time"

	lf "github.com/koczepeter/invoiceproxy-loggerfactory"
	_ "github.com/lib/pq"
)

type TransactionStatusManager struct {
	Log *lf.Logger
	DB  *sql.DB
}

//
//
func (tsm *TransactionStatusManager) Init(db *sql.DB, log *lf.Logger) {

	tsm.Log = log
	tsm.DB = db
}

//
//
func (tsm *TransactionStatusManager) SaveTransactionId(id, transactionId string) {

	if _, err := tsm.DB.Exec(setTransactionQuery, transactionId, id); err != nil {
		tsm.Log.Error(fmt.Sprintf("unable to set transactionId to the cache table (id:%s): %s", id, err.Error()))
	}
}

//
//
func (tsm *TransactionStatusManager) Add(id string) {

	now := time.Now().Unix()
	_, err := tsm.DB.Exec(addTransactionQuery, id, now)
	if err != nil {
		tsm.Log.Error(fmt.Sprintf("unable to add new item to the transactionstatus table: %s", err.Error()))
	}
}

//
//
func (tsm *TransactionStatusManager) EnsureTableExists() {
	if _, err := tsm.DB.Exec(createNavTransactionTableQuery); err != nil {
		tsm.Log.Emerg(fmt.Sprintf("unable to create sql table: %s", err.Error()))
	}

}

//
//
func (tsm *TransactionStatusManager) ClearTable() {
	if _, err := tsm.DB.Exec(clearTransactionsQuery); err != nil {
		tsm.Log.Emerg(fmt.Sprintf("unable to truncate sql table: %s", err.Error()))
	}
}

//
//It retrieves the NAV transactionId that belongs to the provided InvoiceProxy transactionId
//The second parameter is greater than zero only if the InvoiceProxy transactionId exists
func (tsm *TransactionStatusManager) GetTransactionIdById(id string) (transactionId string, createdAt int64) {

	var err error

	tsm.Log.Context = ""
	tsm.Log.ReportId = ""

	tsm.Log.Debug(fmt.Sprintf("looking for the transactionId in the transaction's cache (id:%s)", id))

	row := tsm.DB.QueryRow(getTransactionIdByIdQuery, id)
	err = row.Scan(&transactionId, &createdAt)

	if err == sql.ErrNoRows {
		tsm.Log.Debug(fmt.Sprintf("transactionId was not found in the cache table (id:%s)", id))
		return
	}

	if err != nil {
		tsm.Log.Emerg(fmt.Sprintf("failed to get transactionId from the cache table (id: %s): %s", id, err.Error()))
	}

	tsm.Log.Debug(fmt.Sprintf("transactionId was found in the cache table (id:%s, transactionId:%s)", id, transactionId))

	return
}
