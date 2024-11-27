FROM golang:1.23 AS build
WORKDIR /app
COPY go.mod go.sum /app/       
RUN go mod download
ADD cmd/warehouse-api/* /app/
ADD pkg/clients/* /app/pkg/clients/
ADD pkg/common/* /app/pkg/common/

RUN CGO_ENABLED=0 GOOS=linux go build -o /warehouse-api

FROM gcr.io/distroless/base-debian11 AS run-container

WORKDIR /

COPY --from=build /warehouse-api /warehouse-api

EXPOSE 8081

USER nonroot:nonroot

ENTRYPOINT ["/warehouse-api"]