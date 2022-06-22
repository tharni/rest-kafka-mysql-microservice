package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/segmentio/kafka-go"
)

type Job struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Company     string `json:"company"`
	Salary      string `json:"salary"`
}

const (
	topic         = "test2"
	brokerAddress = "localhost:9092"
)

func main() {
	router := gin.Default()
	router.POST("/jobs", jobsPostHandler)

	server := &http.Server{
		Addr:    ":9090",
		Handler: router,
	}
	server.ListenAndServe()
}
func jobsPostHandler(c *gin.Context) {
	var job Job

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&job); err != nil {
		return
	}
	ctx := context.Background()
	saveJobToKafka(job, ctx)
	c.IndentedJSON(http.StatusCreated, job)
}

func saveJobToKafka(job Job, ctx context.Context) {

	fmt.Println("save to kafka")

	jsonString, err := json.Marshal(job)
	if err != nil {
		return
	}
	jobString := string(jsonString)
	fmt.Print(jobString)

	// initialize a counter
	i := 0

	l := log.New(os.Stdout, "kafka writer: ", 0)
	// intialize the writer with the broker addresses, and the topic
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{brokerAddress},
		Topic:   topic,
		// assign the logger to the writer
		Logger: l,
	})

	// each kafka message has a key and value. The key is used
	// to decide which partition (and consequently, which broker)
	// the message gets published on
	err1 := w.WriteMessages(ctx, kafka.Message{
		Key: []byte(strconv.Itoa(0)),
		// create an arbitrary message payload for the value
		Value: []byte(jobString),
	})
	if err1 != nil {
		panic("could not write message " + err.Error())
	}

	// log a confirmation once the message is written
	fmt.Println("writes:", i)

	if err != nil {
		panic("could not write message " + err.Error())
	}

}
