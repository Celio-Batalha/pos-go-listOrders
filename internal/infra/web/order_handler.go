package web

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Celio-Batalha/pos-go-listOrders/internal/entity"
	"github.com/Celio-Batalha/pos-go-listOrders/internal/usecase"
	"github.com/Celio-Batalha/pos-go-listOrders/pkg/events"
)

type WebOrderHandler struct {
	EventDispacher    events.EventDispatcherInterface
	OrderRepository   entity.OrderRepositoryInterface
	OrderCreatedEvent events.EventInterface
	// ListOrderUseCase  usecase.ListOrderUseCase
}

func NewWebOrderHandler(
	EventDispatcher events.EventDispatcherInterface,
	OrderRepository entity.OrderRepositoryInterface,
	OrderCreatedEvent events.EventInterface,
	// ListOrderUseCase usecase.ListOrderUseCase,
) *WebOrderHandler {
	return &WebOrderHandler{
		EventDispacher:    EventDispatcher,
		OrderRepository:   OrderRepository,
		OrderCreatedEvent: OrderCreatedEvent,
		// ListOrderUseCase:  ListOrderUseCase,
	}
}

func (h *WebOrderHandler) Create(w http.ResponseWriter, r *http.Request) {
	var dto usecase.OrderInputDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createOrder := usecase.NewCreateOrderUseCase(h.OrderRepository, h.OrderCreatedEvent, h.EventDispacher)
	output, err := createOrder.Execute(dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *WebOrderHandler) List(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	listOrder := usecase.NewListOrderUseCase(h.OrderRepository)
	log.Println("Calling ListOrderUseCase.Execute()")
	orders, err := listOrder.Execute()
	if err != nil {
		log.Printf("Error in ListOrderUseCase: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Printf("Found %d orders", len(orders))
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(orders)
}
