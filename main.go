package main

import (
	"log"

	"github.com/ThaliaAC/labora-wallet/config"
	"github.com/ThaliaAC/labora-wallet/controllers"
	"github.com/ThaliaAC/labora-wallet/services"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	var dbHandler, error = services.Connect_DB()
	if error != nil {
		log.Fatal(error)
	}
	walletService := &services.WalletService{DbHandler: dbHandler}
	controller := &controllers.WalletController{WalletService: *walletService}

	router := mux.NewRouter()

	router.HandleFunc("/CreateWallet", controller.CreateWallet).Methods("POST")
	router.HandleFunc("/UpdateWallet/{id}", controller.UpdateWallet).Methods("PUT")
	router.HandleFunc("/DeleteWallet/{id}", controller.DeleteWallet).Methods("DELETE")
	router.HandleFunc("/WalletStatus", controller.WalletStatus).Methods("GET")
	//router.HandleFunc("/CreateLogs", controller.CreateLog).Methods("GET")
	router.HandleFunc("/GetLogs", controller.GetLogs).Methods("GET")

	corsOptions := handlers.CORS(
		handlers.AllowedOrigins([]string{"http://localhost:4000/"}),
		handlers.AllowedMethods([]string{"GET", "POST"}),
	)
	handler := corsOptions(router)

	if err := config.StartServer(handler); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
