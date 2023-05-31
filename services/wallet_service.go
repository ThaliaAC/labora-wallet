package services

import (
	"fmt"
	"time"

	"github.com/ThaliaAC/labora-wallet/models"
)

// Inyección de dependencias
type WalletService struct {
	DbHandler models.DbHandler
}

func (s *WalletService) CreateRequest(Body_request models.Api_Request_To_Truora) (models.Wallet, error) {
	var wallet models.Wallet
	var log models.Log
	var truoraGetResponse models.TruoraGetResponse
	autorization, err := GetApproval(truoraGetResponse.Check.Score)
	if err != nil {
		return models.Wallet{}, fmt.Errorf("API request failed %w", err)
	}
	wallet.National_id = Body_request.National_id
	wallet.Country = Body_request.Country
	wallet.RequestDate = time.Now()
	wallet.Balance = 0

	log.National_id = Body_request.National_id
	log.Country = Body_request.Country
	log.RequestDate = time.Now()
	log.RequestType = "CREATE WALLET"

	if !autorization {
		log.Status = "REJECTED"
		err := s.DbHandler.CreateLog(log)
		if err != nil {
			return models.Wallet{}, fmt.Errorf("error creating the log: %w", err)
		}
		return models.Wallet{}, nil
	}

	log.Status = "APPROVED"
	wallet, err = s.DbHandler.CreateWallet(wallet, log)
	if err != nil {

		return models.Wallet{}, fmt.Errorf("error creating the wallet %w", err)
	}

	return wallet, nil
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

func (s *WalletService) WalletStatus(id int) (bool, error) {

	return s.DbHandler.WalletStatus(id)
}

func (s *WalletService) CreateLog(log models.Log) error {

	return s.DbHandler.CreateLog(log)
}

func (s *WalletService) GetLogs(log models.Log) (models.Log, error) {

	return s.DbHandler.GetLogs(log)
}
