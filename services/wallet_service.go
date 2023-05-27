package services

import (
	"github.com/ThaliaAC/labora-wallet/models"
)

type WalletService struct {
	DbHandler models.DBHandler
}

func (s *WalletService) CreateWallet(wallet models.Wallet, log models.Log) (models.Wallet, error) {

	return s.DbHandler.CreateWallet(wallet, log)
}

func (s *WalletService) UpdateWallet(id int, wallet models.Wallet) (models.Wallet, error) {

	return s.DbHandler.UpdateWallet(id, wallet)
}

func (s *WalletService) DeleteWallet(id int, log models.Log) error {

	return s.DbHandler.DeleteWallet(id, log)
}

func (s *WalletService) WalletStatus(pages, walletsPerPage int) ([]models.Wallet, int, error) {

	return s.DbHandler.WalletStatus(pages, walletsPerPage)
}

func (s *WalletService) CreateLog(log models.Log) error {

	return s.DbHandler.CreateLog(log)
}

func (s *WalletService) GetLogs(pages, logsPerPage int) ([]models.Log, int, error) {

	return s.DbHandler.GetLogs(pages, logsPerPage)
}