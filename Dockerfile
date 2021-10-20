FROM golang:1.16.6 as builder
WORKDIR /Users/vincentlafosse/go/src/github.com/ileossa/go-bank-backend

RUN #go clean -modcache

COPY go.mod .
COPY main.go .

ADD http/handlers handlers
ADD http/service service
ADD http/utils utils

RUN go mod download github.com/gin-gonic/gin
RUN go mod download github.com/yuin/goldmark
RUN go mod download github.com/ileossa/go-bank-backend/
#RUN go get github.com/ileossa/go-bank-backend/http/handlers
#RUN go get github.com/ileossa/go-bank-backend/http/service

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bank .

FROM scratch
COPY --from=builder /Users/vincentlafosse/ileossa/bank/bank .
EXPOSE 8080
ENTRYPOINT ["./bank"]