package services

import (
	"github.com/ThaliaAC/labora-wallet/models"
)

type WalletService struct {
	dbHandler models.DBHandler
}

func (s *WalletService) CreateWallet(wallet models.Wallet, log models.Log) (models.Wallet, error) {

	return s.dbHandler.CreateWallet(wallet, log)
}

func (s *WalletService) UpdateWallet(id int, wallet models.Wallet) (models.Wallet, error) {

	return s.dbHandler.UpdateWallet(id, wallet)
}

func (s *WalletService) DeleteWallet(id int, log models.Log) error {

	return s.dbHandler.DeleteWallet(id, log)
}

func (s *WalletService) WalletStatus(id int) (bool, error) {

	return s.dbHandler.WalletStatus(id)
}

func (s *WalletService) CreateLog(log models.Log) error {

	return s.dbHandler.CreateLog(log)
}

func (s *WalletService) GetLogs(log models.Log) (models.Log, error) {

	return s.dbHandler.GetLogs(log)
}
