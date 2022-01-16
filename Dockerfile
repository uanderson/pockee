FROM golang:alpine as build
RUN apk --update add ca-certificates
RUN apk --update add git

# All these steps will be cached
RUN mkdir /app
WORKDIR /app

ARG ACCESS_USERNAME
ARG ACCESS_TOKEN

RUN git config --global url."https://$ACCESS_USERNAME:$ACCESS_TOKEN@github.com".insteadOf "https://github.com"

COPY go.mod .
COPY go.sum .

RUN go mod download
COPY . .

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /go/bin/app

FROM scratch
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=build /go/bin/app /go/bin/app

ENTRYPOINT ["/go/bin/app"]