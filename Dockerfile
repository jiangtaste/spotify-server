FROM golang:alpine AS builder

WORKDIR /opt

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY service ./service
COPY core ./core
COPY router ./router
COPY main.go ./

RUN CGO_ENABLED=0 go build -ldflags "-s -w" -o app
# RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine
WORKDIR /opt

COPY --from=builder /opt/app /opt/app
EXPOSE 8080

CMD ["/opt/app"]
