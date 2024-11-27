FROM golang:1.23 AS build
WORKDIR /app
COPY go.mod go.sum /app/       
RUN go mod download
ADD cmd/store-api/* /app/
ADD pkg/clients/* /app/pkg/clients/
ADD pkg/common/* /app/pkg/common/

RUN go build -o /store-api

FROM gcr.io/distroless/base-debian11 AS run-container

WORKDIR /

COPY --from=build /store-api /store-api

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/store-api"]