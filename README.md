Here is the high-level architecture of this simple asynchronous processing example wtih 2 microservices.
![rest-kafka-mongo-microservice-draw-io]

**Microservice 1** - is a REST microservice which receives data from a /POST http call to it. After receiving the request, it retrieves the data from the http request and saves it to Kafka. After saving, it responds to the caller with the same data sent via /POST

**Microservice 2** - is a microservice which subscribes to a topic in Kafka where Microservice 1 saves the data. Once a message is consumed by the microservice, it then saves the data to MysqlDB.

Before you proceed, we need a few things to be able to run these microservices:
1. Start kafka (Kafka listening on port 9092 and my confluent url is http://localhost:9021/)
2. Install & run mysql db. 

Run Microservice 1. Make sure Kafka is running.
```
$ go run rest-kafka-sample.go
```

I used Postman to send data to Microservice 1

Since we are not running Microservice 2 yet, the data saved by Microservice 1 will just be in Kafka. Let's consume it and save to MysqlDB by running Microservice 2.
```
$ go run kafka-mysql-sample.go
```

Now you'll see that Microservice 2 consumes the data and saves it to MysqlDB

Check if data is saved in MysqlDB. If it is there, we're good!
