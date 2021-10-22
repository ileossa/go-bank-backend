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

publish-images:
	docker push bank:latest

doc:
	swag i

publish-doc:
	echo "Not implemented"

