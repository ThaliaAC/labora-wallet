package controllers

import (
	"net/http"

	"github.com/ThaliaAC/labora-wallet/services"
)

type WalletController struct {
	WalletService services.WalletService
}

func ResponseJson() {

}

func (c *WalletController) CreateWallet(response http.ResponseWriter, request *http.Request) {
}

func (c *WalletController) UpdateWallet(response http.ResponseWriter, request *http.Request) {

}

func (c *WalletController) DeleteWallet(response http.ResponseWriter, request *http.Request) {
}

func (c *WalletController) WalletStatus(w http.ResponseWriter, r *http.Request) {
}

func (c *WalletController) GetLogs(w http.ResponseWriter, r *http.Request) {
}
