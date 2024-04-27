run:
	CompileDaemon --command="./rabbitmq-message-queue"

build:
	go build -o rabbitmq-message-queue

start:
	./rabbitmq-message-queue