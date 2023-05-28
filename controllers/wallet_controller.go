package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/ThaliaAC/labora-wallet/models"
	"github.com/ThaliaAC/labora-wallet/services"
	"github.com/gorilla/mux"
)

type WalletController struct {
	WalletService services.WalletService
}

func ResponseJson(response http.ResponseWriter, status int, data interface{}) error {
	bytes, err := json.Marshal(data)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)

		return fmt.Errorf("error while marshalling object %v, trace: %+v", data, err)
	}
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(status)
	_, err = response.Write(bytes)
	if err != nil {

		return fmt.Errorf("error while writing bytes to response writer: %+v", err)
	}

	return nil
}

func (c *WalletController) CreateWallet(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	var newPerson models.Person

	err := json.NewDecoder(request.Body).Decode(&newPerson)
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(response).Encode(err)
		return
	}
	status, wallet, err := services.CreateRequest(newPerson)
	response.WriteHeader(http.StatusCreated)
	ResponseJson(response, status, wallet)
	fmt.Println("Wallet successfully created")

}

func (c *WalletController) UpdateWallet(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	parameters := mux.Vars(request)
	var wallet models.Wallet

	id, err := strconv.Atoi(parameters["id"])
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte("ID must be a number"))

		return
	}

	err = json.NewDecoder(request.Body).Decode(&wallet)
	defer request.Body.Close()
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)

		return
	}

	wallet, err = c.WalletService.UpdateWallet(id, wallet)
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)

		return
	}

	ResponseJson(response, http.StatusOK, wallet)
}

func (c *WalletController) DeleteWallet(response http.ResponseWriter, request *http.Request) {
	var log models.Log
	parameters := mux.Vars(request)
	id, err := strconv.Atoi(parameters["id"])

	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte("ID must be a number"))

		return
	}

	err = c.WalletService.DeleteWallet(id, log)
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)

		return
	}

	ResponseJson(response, http.StatusOK, models.Wallet{})
}

func (c *WalletController) WalletStatus(w http.ResponseWriter, r *http.Request) {
}

func (c *WalletController) CreateLog(w http.ResponseWriter, r *http.Request) {
}
func (c *WalletController) GetLog(w http.ResponseWriter, r *http.Request) {
}
