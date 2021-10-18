FROM golang:1.16.6 as builder
WORKDIR /Users/vincentlafosse/ileossa/bank
COPY main.go .
COPY go.mod .
COPY go.sum .

RUN go mod download github.com/yuin/goldmark
RUN go get github.com/ileossa/go-bank-backend/controllers
RUN go get github.com/ileossa/go-bank-backend/database
RUN go get github.com/ileossa/go-bank-backend/docs

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bank .

FROM scratch
COPY --from=builder /Users/vincentlafosse/ileossa/bank/bank .
EXPOSE 8080
ENTRYPOINT ["./bank"]