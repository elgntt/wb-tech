compose-up: ### Run docker-compose
	docker-compose up --build -d && docker-compose logs -f
.PHONY: compose-up

compose-down: ### Down docker-compose
	docker-compose down --remove-orphans
.PHONY: compose-down

test: ### run test
	go test -v ./...

cover-html: ### run test with coverage and open html report
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out
	rm coverage.out
.PHONY: coverage-html

cover: ### run test with coverage
	go test -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out
	rm coverage.out
.PHONY: coverage

mockgen: ### generate mock
	mockgen -source=internal/service/deps.go 	-destination=internal/service/mocks_test.go -package=service
	mockgen -source=internal/api/deps.go	    -destination=internal/api/mocks_test.go -package=api
.PHONY: mockgen