FROM golang:alpine as build
RUN apk --update add ca-certificates
RUN apk --update add git

# All these steps will be cached
RUN mkdir /app
WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /go/bin/app

FROM alpine:latest

COPY --from=build /app/.migrations /app/.migrations
COPY --from=build /app/Dockerrun /app/Dockerrun
COPY --from=build /go/bin/app /go/bin/app

RUN apk --update --no-cache add \
    bash \
    curl \
    postgresql-client

RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-amd64.tar.gz -o migrate.tar.gz \
  && mkdir -p /usr/local/bin/migrate \
  && tar -zxf migrate.tar.gz --directory /usr/local/bin/migrate \
  && rm -f migrate.tar.gz

RUN chmod +x /app/Dockerrun

ENTRYPOINT ["/app/Dockerrun", "/go/bin/app"]