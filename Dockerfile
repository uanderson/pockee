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
    openjdk17-jre-headless \
    postgresql-client

RUN curl -L https://github.com/liquibase/liquibase/releases/download/v4.8.0/liquibase-4.8.0.zip -o liquibase.zip \
  && unzip liquibase.zip -d /usr/local/bin/liquibase \
  && rm -f liquibase.zip

RUN chmod +x /app/Dockerrun

ENTRYPOINT ["/app/Dockerrun", "/go/bin/app"]