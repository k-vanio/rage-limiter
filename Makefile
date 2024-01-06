up: 
	docker-compose up --build

test:
	docker-compose exec app go test -race -v ./...

cover:
	docker-compose exec app go test -race -cover ./...