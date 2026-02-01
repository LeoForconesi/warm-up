package server

import (
	"log"
	"warm-up/internal/adapters/in/http"
	"warm-up/internal/adapters/out/inMemory"
	"warm-up/internal/application/orders"
)

func New() *Server {
	//cfg := config.LoadConfig() // Esto configura la aplicación con variables de entorno o archivos de configuración...
	server, err := BuildDependencies() //(cfg)
	if err != nil {
		log.Fatal(err)
	}
	return server
}

func BuildDependencies() (*Server, error) {
	logg := NewLogger("development") //cfg.Env)

	orderRepo := configDb() //cfg)
	createOrderUC := orders.NewCreateOrderUC(orderRepo)
	getOrderByIDUC := orders.NewGetOrdersByOrderIDUC(orderRepo)

	orderHandler := http.NewOrdersHandler(createOrderUC, getOrderByIDUC)
	r := http.NewRouter(orderHandler)

	return NewServer(r, logg), nil
}

func configDb() *inMemoryPersistence.InMemoryOrderRepository {
	return inMemoryPersistence.NewInMemoryOrderRepository()
}
