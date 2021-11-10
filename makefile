dev:
	go run .

provisioning:
	rm -f go.mod go.sum
	go mod init github.com/ileossa/go-bank-backend
	go mod tidy
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
	echo "Not implemented"
	#swag i

publish-doc:
	echo "Not implemented"

