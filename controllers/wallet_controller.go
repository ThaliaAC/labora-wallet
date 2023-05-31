package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/ThaliaAC/labora-wallet/models"
	"github.com/ThaliaAC/labora-wallet/services"
	"github.com/gorilla/mux"
)

type PostgresDBHandler struct {
	db *sql.DB
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

func (p *PostgresDBHandler) CreateWallet(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	var newPerson models.Person

	err := json.NewDecoder(request.Body).Decode(&newPerson)
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		if err := json.NewEncoder(response).Encode(err); err != nil {
			fmt.Println("Error encoding the error: ", err)
		}
		return
	}

	wallet, err := CreateRequest(newPerson)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		fmt.Println("Error creating wallet: ", err)
		return
	}
	response.WriteHeader(http.StatusCreated)
	ResponseJson(response, http.StatusOK, wallet)
	fmt.Println("Wallet successfully created")

}

func (p *PostgresDBHandler) CreateRequest(Body_request models.Api_Request_To_Truora) (int, models.Wallet, error) {
	var wallet models.Wallet
	var truoraGetResponse models.TruoraGetResponse
	autorization, err := GetApproval(truoraGetResponse.Check.Score)
	if err != nil {
		return http.StatusInternalServerError, models.Wallet{}, fmt.Errorf("API request failed %w", err)
	}
	wallet.National_id = Body_request.National_id
	wallet.Country = Body_request.Country
	wallet.RequestDate = time.Now()
	wallet.Balance = 0

	if !autorization {
		err := p.db.CreateLog(wallet.National_id, wallet.Country, "Denied", "CREATE WALLET")
		if err != nil {
			return http.StatusInternalServerError, models.Wallet{}, fmt.Errorf("Error creating the log: %w", err)
		}

		return http.StatusConflict, models.Wallet{}, nil
	}

	wallet, err = p.db.CreateWallet(wallet)
	if err != nil {

		return http.StatusInternalServerError, models.Wallet{}, fmt.Errorf("Error creating the wallet %w", err)
	}

	return http.StatusOK, wallet, nil
}

func (p *PostgresDBHandler) UpdateWallet(response http.ResponseWriter, request *http.Request) {
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

	wallet, err = services.UpdateWallet(id, wallet)
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)

		return
	}

	ResponseJson(response, http.StatusOK, wallet)
}

func (p *PostgresDBHandler) DeleteWallet(response http.ResponseWriter, request *http.Request) {
	var log models.Log
	parameters := mux.Vars(request)
	id, err := strconv.Atoi(parameters["id"])

	if err != nil {
		http.Error(response, "it must be a number", http.StatusBadRequest)
	}

	err = services.DeleteWallet(id, log)
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}

	ResponseJson(response, http.StatusOK, models.Wallet{})
}

func (p *PostgresDBHandler) WalletStatus(w http.ResponseWriter, r *http.Request) {
}

func (p *PostgresDBHandler) CreateLog(w http.ResponseWriter, r *http.Request) {
}
func (p *PostgresDBHandler) GetLog(w http.ResponseWriter, r *http.Request) {
}
