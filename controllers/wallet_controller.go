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

func (p *WalletController) CreateWallet(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	var newPerson models.Person
	var Body_request models.Api_Request_To_Truora
	err := json.NewDecoder(request.Body).Decode(&newPerson)
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		if err := json.NewEncoder(response).Encode(err); err != nil {
			fmt.Println("Error encoding the error: ", err)
		}
		return
	}

	wallet, err := p.WalletService.CreateRequest(Body_request)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		fmt.Println("Error creating wallet: ", err)
		return
	}
	response.WriteHeader(http.StatusCreated)
	ResponseJson(response, http.StatusOK, wallet)
	fmt.Println("Wallet successfully created")

}

func (p *WalletController) UpdateWallet(response http.ResponseWriter, request *http.Request) {
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

	wallet, err = p.WalletService.UpdateWallet(id, wallet)
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)

		return
	}

	ResponseJson(response, http.StatusOK, wallet)
}

func (p *WalletController) DeleteWallet(response http.ResponseWriter, request *http.Request) {
	var log models.Log
	parameters := mux.Vars(request)
	id, err := strconv.Atoi(parameters["id"])

	if err != nil {
		http.Error(response, "it must be a number", http.StatusBadRequest)
	}

	err = p.WalletService.DeleteWallet(id, log)
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}

	ResponseJson(response, http.StatusOK, models.Wallet{})
}

func (p *WalletController) WalletStatus(response http.ResponseWriter, request *http.Request) {
	var log models.Log
	parameters := mux.Vars(request)
	id, err := strconv.Atoi(parameters["id"])

	if err != nil {
		http.Error(response, "it must be a number", http.StatusBadRequest)
	}

	statusMsg, err := services.WalletStatusFromLog(id, log)
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}

	ResponseJson(response, http.StatusOK, statusMsg)
}

func (p *WalletController) CreateLog(response http.ResponseWriter, request *http.Request) {
}

func (p *WalletController) GetLogs(response http.ResponseWriter, request *http.Request) {
}
