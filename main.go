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

	//TODO: Debe estar en Services
	dbHandler, err := services.Connect_DB()
	if err != nil {
		log.Fatal(err)
	}
	walletService := &services.WalletService{DbInterface: dbHandler}
	walletController := &controllers.WalletController{WalletService: *walletService}

	router := mux.NewRouter()

	router.HandleFunc("/CreateWallet", walletController.CreateWallet).Methods("POST")
	router.HandleFunc("/UpdateWallet/{id}", walletController.UpdateWallet).Methods("PUT")
	router.HandleFunc("/DeleteWallet/{id}", walletController.DeleteWallet).Methods("DELETE")
	router.HandleFunc("/WalletStatus", walletController.WalletStatus).Methods("GET")
	router.HandleFunc("/GetLogs", walletController.GetLogs).Methods("GET")

	//TODO: Crear una pequeña app en React
	corsOptions := handlers.CORS(
		handlers.AllowedOrigins([]string{"http://localhost:4000/"}),
		handlers.AllowedMethods([]string{"GET", "POST"}),
	)
	handler := corsOptions(router)

	if err := config.StartServer(handler); err != nil {
		log.Fatalf("Server startup failed: %v", err)
		panic(err)
	}
}
