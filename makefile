dev:
	go run main.go

provisionnig:
	docker-compose up -d

destroy:
	docker-compose down

test:
	go test

package:
	docker build -t bank:latest .

publish:
	docker push bank:latest

doc:
	swag i
