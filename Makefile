build:
	cd ./client; npm run build
	go build
dev:
	go run main.go & cd ./client;npm run start
	
test:
	go test ./api/topics/test