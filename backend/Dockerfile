FROM golang:1.23-alpine

WORKDIR /go-app-dir

RUN go install github.com/air-verse/air@latest

# COPY go.mod go.sum ./
# COPY ./golang/go.mod ./
COPY ./golang ./

RUN go mod download

# COPY ./golang ./


#production
#CMD ["go", "run", "main.go"]
#development
CMD ["air", "-c", ".air.toml"]