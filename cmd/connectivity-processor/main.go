package main

import (
	"context"
	"fmt"

	http_infrastructure "connectivity-processor/internal/http/infrastructure"

	processor_application "connectivity-processor/internal/processor/application"
	redis_client "connectivity-processor/pkg/redis"
)

func main() {
	fmt.Println("starting connectivity processor")

	var ctx = context.Background()

	redisClient := redis_client.NewRedisClient()
	connectivityRepository := redis_client.NewRedisConnectivityRepository(ctx, redisClient)

	// TESTS
	// fmt.Println(connectivityRepository.IsConnected("1234"))
	// fmt.Println(connectivityRepository.GetConnectionId("1234"))

	// fmt.Println(connectivityRepository.SaveConnection("1234", "connection_id"))
	// fmt.Println(connectivityRepository.IsConnected("1234"))
	// fmt.Println(connectivityRepository.GetConnectionId("1234"))

	// fmt.Println(connectivityRepository.SaveConnection("1234", "new_connection_id"))
	// fmt.Println(connectivityRepository.IsConnected("1234"))
	// fmt.Println(connectivityRepository.GetConnectionId("1234"))

	// fmt.Println(connectivityRepository.RemoveConnection("1234"))
	// fmt.Println(connectivityRepository.IsConnected("1234"))
	// fmt.Println(connectivityRepository.GetConnectionId("1234"))

	ceProcessor := processor_application.NewConnectedEventProcessor(connectivityRepository)
	deProcessor := processor_application.NewDisconnectedEventProcessor(connectivityRepository)

	http_infrastructure.StartHttpServer("3000", 15, 15, ceProcessor, deProcessor)
}
