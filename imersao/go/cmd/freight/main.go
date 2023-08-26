package main

import (
	"database/sql"
	"encoding/json"
	"fmt"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joaovian06/fullstack-imersao/go/internal/routes/entity"
	"github.com/joaovian06/fullstack-imersao/go/internal/routes/infra/repository"
	"github.com/joaovian06/fullstack-imersao/go/internal/routes/usecase"
	"github.com/joaovian06/fullstack-imersao/go/pkg/kafka"
)

func main() {
	msgChan := make(chan *ckafka.Message)
	topics := []string{"routes"}
	servers := "host.docker.internal:9094"
	go kafka.Consume(topics, servers, msgChan)

	db, err := sql.Open("mysql", "root:root@tcp(host.docker.internal:3307)/routes?parseTime=true")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	repository := repository.NewRouteRepositoryMysql(db)
	freight := entity.NewFreight(10)
	createRouteUseCase := usecase.NewCreateRouteUseCase(repository, freight)
	changeRouteStatusUseCase := usecase.NewChangeRouteStatusUseCase(repository)

	for msg := range msgChan {
		input := usecase.CreateRouteInput{}
		json.Unmarshal(msg.Value, &input)

		switch input.Event {
		case "RouteCreated":
			_, err := createRouteUseCase.Execute(input)
			if err != nil {
				fmt.Println(err)
			}

		case "RouteStarted", "RouteFinished":
			input := usecase.ChangeRouteStatusInput{}
			json.Unmarshal(msg.Value, &input)
			_, err := changeRouteStatusUseCase.Execute(input)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}