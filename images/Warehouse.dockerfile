FROM golang:1.23 AS build
WORKDIR /app
COPY go.mod go.sum ./       
RUN go mod download
COPY cmd/warehouse-api/* internals/* pkg/* ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /warehouse-api

FROM gcr.io/distroless/base-debian11 AS run-container

WORKDIR /

COPY --from=build /warehouse-api /warehouse-api

EXPOSE 8081

USER nonroot:nonroot

ENTRYPOINT ["/warehouse-api"]