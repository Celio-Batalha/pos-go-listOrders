package main

import (
	"database/sql"
	"fmt"
	"net"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/Celio-Batalha/pos-go-listOrders/configs"
	"github.com/Celio-Batalha/pos-go-listOrders/internal/event/_handler"
	"github.com/Celio-Batalha/pos-go-listOrders/internal/infra/graphql/graph"
	"github.com/Celio-Batalha/pos-go-listOrders/internal/infra/grpc/pb"
	"github.com/Celio-Batalha/pos-go-listOrders/internal/infra/grpc/service"
	"github.com/Celio-Batalha/pos-go-listOrders/internal/infra/web/webserver"
	"github.com/Celio-Batalha/pos-go-listOrders/pkg/events"
	"github.com/streadway/amqp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	// mysql
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	db, err := sql.Open(configs.DBDriver, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", configs.DBUser, configs.DBPassword, configs.DBHost, configs.DBPort, configs.DBName))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rabbitMQChannel := getRabbitMQChannel()

	eventDispatcher := events.NewEventDispatcher()
	eventDispatcher.Register("OrderCreated", &_handler.OrderCreatedHandler{
		RabbitMQChannel: rabbitMQChannel,
	})

	createOrderUseCase := NewCreateOrderUseCase(db, eventDispatcher)
	listOrderUseCase := NewListOrderUseCase(db)
	// createOrderUseCase := NewCreateOrderUseCase(db, eventDispatcher)

	// web server setup
	webserver := webserver.NewWebServer(configs.WebServerPort)

	webOderHandler := NewWebOrderHandler(db, eventDispatcher)
	webserver.AddHandler("POST /orders", webOderHandler.Create)
	webserver.AddHandler("GET /orders", webOderHandler.List)
	fmt.Println("Starting web server on port", configs.WebServerPort)
	go webserver.Start()

	// gRPC server setup
	grpcServer := grpc.NewServer()
	createdOrderService := service.NewOrderService(*createOrderUseCase, *listOrderUseCase)
	pb.RegisterOrderServiceServer(grpcServer, createdOrderService)
	reflection.Register(grpcServer)

	fmt.Println("Starting gRPC server on port", configs.GRPCServerPort)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", configs.GRPCServerPort))
	if err != nil {
		panic("Error starting gRPC server: " + err.Error())
	}
	go grpcServer.Serve(lis)

	// GraphQL server setup
	srv := handler.New(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		CreateOrderUseCase: *createOrderUseCase,
		ListOrderUseCase:   *listOrderUseCase,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	fmt.Println("Starting GraphQL server on port", configs.GraphqlServerPort)
	fmt.Printf("GraphQL playground: http://localhost:%s/\n", configs.GraphqlServerPort)
	go http.ListenAndServe(":"+configs.GraphqlServerPort, nil)

	// Manter o programa rodando
	select {}
}

func getRabbitMQChannel() *amqp.Channel {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic("Error connecting to RabbitMQ: " + err.Error())
	}
	channel, err := conn.Channel()
	if err != nil {
		panic("Error creating RabbitMQ channel: " + err.Error())
	}
	return channel
}
