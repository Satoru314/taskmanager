package services

import (
	"taskmanager/models"
	"taskmanager/repository"
)

func (s *MyAppServices) GetAccountService(accountID int) (resAccount models.Account, err error) {
	resAccount, err = repository.SelectAccount(s.db, accountID)
	if err != nil {
		return models.Account{}, err
	}
	return resAccount, nil
}

func (s *MyAppServices) PostAccountService(reqAccount models.Account) (resAccount models.Account, err error) {
	resAccount, err = repository.InsertAccount(s.db, reqAccount)
	if err != nil {
		return models.Account{}, err
	}
	return resAccount, nil
}

func (s *MyAppServices) PutAccountService(reqAccount models.Account) (resAccount models.Account, err error) {
	resAccount, err = repository.UpdateAccount(s.db, reqAccount)
	if err != nil {
		return models.Account{}, err
	}
	return resAccount, nil
}

func (s *MyAppServices) DeleteAccountService(accountID int) (err error) {
	err = repository.DeleteAccount(s.db, accountID)
	return err
}
