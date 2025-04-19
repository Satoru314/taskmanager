package services

import (
	"taskmanager/models"
	"taskmanager/repositorys"
)

func (s *MyAppServices) GetAccountService(accountID int) (resAccount models.Account, err error) {
	resAccount, err = repositorys.SelectAccount(s.db, accountID)
	if err != nil {
		return models.Account{}, err
	}
	return resAccount, nil
}

func (s *MyAppServices) PostAccountService(reqAccount models.Account) (resAccount models.Account, err error) {
	resAccount, err = repositorys.InsertAccount(s.db, reqAccount)
	if err != nil {
		return models.Account{}, err
	}
	return resAccount, nil
}

func (s *MyAppServices) PutAccountService(reqAccount models.Account) (resAccount models.Account, err error) {
	resAccount, err = repositorys.UpdateAccount(s.db, reqAccount)
	if err != nil {
		return models.Account{}, err
	}
	return resAccount, nil
}

func (s *MyAppServices) DeleteAccountService(accountID int) (err error) {
	err = repositorys.DeleteAccount(s.db, accountID)
	return err
}
