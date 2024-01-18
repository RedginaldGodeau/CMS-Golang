FROM golang:1.20-alpine

RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
RUN go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.54.2
