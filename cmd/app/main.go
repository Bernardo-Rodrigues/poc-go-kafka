package main

import (
	"database/sql"
	"encoding/json"
	"test/internal/infra/akafka"
	"test/internal/infra/repository"
	"test/internal/usecase"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(host.docker.internal):3306/products")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	msgChan := make(chan *kafka.Message)
	go akafka.Consume([]string{"products"}, "host.docker.internal:9094", msgChan)

	repository := repository.NewProductRepositoryMySql(db)
	createProductUseCase := usecase.NewCreateProductUseCase(repository)

	for msg := range msgChan {
		dto := usecase.CreateProductInputDto{}
		json.Unmarshal(msg.Value, &dto)
		_, err = createProductUseCase.Execute(dto)
	}
}
