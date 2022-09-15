package transactionStatusManager

const createNavTransactionTableQuery = `CREATE TABLE IF NOT EXISTS nav_transactions (
	id 				VARCHAR (20) NOT NULL,
	transaction_id 	VARCHAR (20) NOT NULL DEFAULT '',
	created_at 		BIGINT NOT NULL,

	PRIMARY KEY (id)
  )`

const addTransactionQuery = `INSERT INTO nav_transactions (id, created_at)  VALUES ($1, $2)`
const setTransactionQuery = `UPDATE nav_transactions SET transaction_id=$1 WHERE id=$2`

const getTransactionIdByIdQuery = `SELECT transaction_id, created_at FROM nav_transactions WHERE id=$1`

const clearTransactionsQuery = `DELETE FROM nav_transactions`
