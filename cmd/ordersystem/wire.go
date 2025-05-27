//go:build wireinject
// +build wireinject

package main

import (
	"database/sql"

	"github.com/Celio-Batalha/pos-go-listOrders/internal/entity"
	"github.com/Celio-Batalha/pos-go-listOrders/internal/event"
	"github.com/Celio-Batalha/pos-go-listOrders/internal/infra/database"
	"github.com/Celio-Batalha/pos-go-listOrders/internal/infra/web"
	"github.com/Celio-Batalha/pos-go-listOrders/internal/usecase"
	"github.com/Celio-Batalha/pos-go-listOrders/pkg/events"
	"github.com/google/wire"
)

var setOrderRepositoryDependency = wire.NewSet(
	database.NewOrderRepository,
	wire.Bind(new(entity.OrderRepositoryInterface), new(*database.OrderRepository)),
)

var setEventDependency = wire.NewSet(
	event.NewOrderCreated,
	wire.Bind(new(events.EventInterface), new(*event.OrderCreated)),
)

var setOrderCreatedEvent = wire.NewSet(
	event.NewOrderCreated,
	wire.Bind(new(events.EventInterface), new(*event.OrderCreated)),
)

func NewCreateOrderUseCase(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *usecase.CreateOrderUseCase {
	wire.Build(
		setOrderRepositoryDependency,
		setOrderCreatedEvent,
		usecase.NewCreateOrderUseCase,
	)
	return &usecase.CreateOrderUseCase{}
}

func NewListOrderUseCase(db *sql.DB) *usecase.ListOrderUseCase {
	wire.Build(
		setOrderRepositoryDependency,
		usecase.NewListOrderUseCase,
	)
	return &usecase.ListOrderUseCase{}
}

func NewWebOrderHandler(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *web.WebOrderHandler {
	wire.Build(
		setOrderRepositoryDependency,
		setOrderCreatedEvent,
		// usecase.NewListOrderUseCase,
		web.NewWebOrderHandler,
	)
	return &web.WebOrderHandler{}
}
