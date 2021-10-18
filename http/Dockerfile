FROM golang:1.16.6 as builder
WORKDIR /Users/vincentlafosse/ileossa/bank
COPY go.mod .
COPY go.sum .
COPY main.go .

ADD handlers handlers
ADD service service
ADD utils utils

RUN go mod download github.com/yuin/goldmark
RUN go get github.com/ileossa/go-bank-backend/http/docs
RUN go get github.com/ileossa/go-bank-backend/http/handlers
RUN go get github.com/ileossa/go-bank-backend/http/service

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bank .

FROM scratch
COPY --from=builder /Users/vincentlafosse/ileossa/bank/bank .
EXPOSE 8080
ENTRYPOINT ["./bank"]