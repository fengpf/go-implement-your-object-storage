#!/bin/bash

LISTEN_ADDRESS=10.29.1.1:12345 STORAGE_ROOT=/tmp/1 RABBITMQ_SERVER=amqp://test:test@10.29.102.173:5672 go run $1/dataServer/dataServer.go &
LISTEN_ADDRESS=10.29.1.2:12345 STORAGE_ROOT=/tmp/2 RABBITMQ_SERVER=amqp://test:test@10.29.102.173:5672 go run $1/dataServer/dataServer.go &
LISTEN_ADDRESS=10.29.1.3:12345 STORAGE_ROOT=/tmp/3 RABBITMQ_SERVER=amqp://test:test@10.29.102.173:5672 go run $1/dataServer/dataServer.go &
LISTEN_ADDRESS=10.29.1.4:12345 STORAGE_ROOT=/tmp/4 RABBITMQ_SERVER=amqp://test:test@10.29.102.173:5672 go run $1/dataServer/dataServer.go &
LISTEN_ADDRESS=10.29.1.5:12345 STORAGE_ROOT=/tmp/5 RABBITMQ_SERVER=amqp://test:test@10.29.102.173:5672 go run $1/dataServer/dataServer.go &
LISTEN_ADDRESS=10.29.1.6:12345 STORAGE_ROOT=/tmp/6 RABBITMQ_SERVER=amqp://test:test@10.29.102.173:5672 go run $1/dataServer/dataServer.go &
LISTEN_ADDRESS=10.29.2.1:12345 ES_SERVER=10.29.102.173:9200 RABBITMQ_SERVER=amqp://test:test@10.29.102.173:5672 go run $1/apiServer/apiServer.go &
LISTEN_ADDRESS=10.29.2.2:12345 ES_SERVER=10.29.102.173:9200 RABBITMQ_SERVER=amqp://test:test@10.29.102.173:5672 go run $1/apiServer/apiServer.go &