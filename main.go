package main

import (
	"log"
	"net/http"

	"github.com/samluiz/delivery-service/api/http/handlers"
	"github.com/samluiz/delivery-service/config/db"
	"github.com/samluiz/delivery-service/config/server"
	"github.com/samluiz/delivery-service/internal/delivery"
)

func main() {
	db := db.OpenMySQLConnection()

	srv := server.NewServer(db)

	deliveryRepository := delivery.NewDeliveryRepository(db)
	deliveryService := delivery.NewDeliveryService(deliveryRepository)
	deliveryHandler := handlers.NewDeliveryHandler(deliveryService)

	srv.Router.HandleFunc("POST /deliveries", deliveryHandler.HandleCreateDelivery)
	srv.Router.HandleFunc("GET /deliveries", deliveryHandler.HandleGetDeliveries)
	srv.Router.HandleFunc("GET /deliveries/{id}", deliveryHandler.HandleGetDelivery)
	srv.Router.HandleFunc("PUT /deliveries/{id}", deliveryHandler.HandleUpdateDelivery)
	srv.Router.HandleFunc("DELETE /deliveries/{id}", deliveryHandler.HandleDeleteDelivery)
	srv.Router.HandleFunc("DELETE /deliveries", deliveryHandler.HandleDeleteAllDeliveries)

	srv.Router.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	log.Println("Server iniciando na porta 8080...")
	log.Fatal(http.ListenAndServe(":8080", srv))
}
