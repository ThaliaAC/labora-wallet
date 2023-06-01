package models

type DbInterface interface {
	CreateWallet(wallet Wallet, log Log) (Wallet, error)
	UpdateWallet(id int, wallet Wallet) (Wallet, error)
	DeleteWallet(id int, log Log) error
	WalletStatus(id int) (string, error)
	CreateLog(log Log) error
	GetLogs(log Log) (Log, error)
}
