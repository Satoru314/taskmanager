package repository

import (
	"database/sql"
	"taskmanager/models"
)

func SelectAccount(db *sql.DB, accountID int) (resAccount models.Account, err error) {
	const sqlStr = `SELECT account_name, path_hash FROM accounts WHERE account_id = $1;`
	row := db.QueryRow(sqlStr, accountID)
	err = row.Scan(&resAccount.AccountName, &resAccount.PathHash)
	if err != nil {
		return models.Account{}, err
	}
	resAccount.AccountID = accountID
	return resAccount, nil
}

func InsertAccount(db *sql.DB, account models.Account) (models.Account, error) {
	const sqlStr = `INSERT INTO accounts (account_name, path_hash) VALUES ($1, $2) RETURNING account_id;`
	var accountID int
	err := db.QueryRow(sqlStr, account.AccountName, account.PathHash).Scan(&accountID)
	if err != nil {
		return models.Account{}, err
	}
	account.AccountID = accountID
	return account, nil
}

func UpdateAccount(db *sql.DB, account models.Account) (models.Account, error) {
	const sqlStr = `UPDATE accounts SET account_name = $1, path_hash = $2 WHERE account_id = $3;`
	_, err := db.Exec(sqlStr, account.AccountName, account.PathHash, account.AccountID)
	if err != nil {
		return models.Account{}, err
	}
	return account, nil
}

func DeleteAccount(db *sql.DB, accountID int) error {
	const sqlStr = `DELETE FROM accounts WHERE account_id = $1;`
	_, err := db.Exec(sqlStr, accountID)
	return err
}
