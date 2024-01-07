up: 
	docker-compose up --build -d

test: up
	docker-compose exec app go test -race -v ./...

cover: up
	docker-compose exec app go test -race -cover ./...

testTokenSuccess:
	docker-compose exec app ./test_token_successs.sh 

testTokenError:
	docker-compose exec app ./test_token_error.sh

testIpSuccess:
	docker-compose exec app ./test_ip_successs.sh 

testIpError:
	docker-compose exec app ./test_ip_error.sh 