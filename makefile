dev:
	docker-compose up -d && go run http/main.go

test:
	go test

package:
	cd http && docker build -t bank:latest .

publish:
	docker push

doc:
	swag i
