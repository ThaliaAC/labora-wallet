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
	var dbHandler, err = services.Connect_DB()
	if err != nil {
		log.Fatal(err)
	}
	walletService := &services.WalletService{DbHandler: dbHandler}
	controllers := &controllers.WalletController{WalletService: *walletService}

	router := mux.NewRouter()

	router.HandleFunc("/CreateWallet", controllers.CreateWallet).Methods("POST")
	router.HandleFunc("/UpdateWallet/{id}", controllers.UpdateWallet).Methods("PUT")
	router.HandleFunc("/DeleteWallet/{id}", controllers.DeleteWallet).Methods("DELETE")
	router.HandleFunc("/WalletStatus", controllers.WalletStatus).Methods("GET")
	router.HandleFunc("/CreateLog", controllers.CreateLog).Methods("GET")
	router.HandleFunc("/GetLog", controllers.GetLog).Methods("GET")

	corsOptions := handlers.CORS(
		handlers.AllowedOrigins([]string{"http://localhost:4000/"}),
		handlers.AllowedMethods([]string{"GET", "POST"}),
	)
	handler := corsOptions(router)

	if err := config.StartServer(handler); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
