FROM golang:1.23 AS build
WORKDIR /app
COPY go.mod go.sum ./       
RUN go mod download
COPY cmd/store-api/* internals/* pkg/* ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /store-api

FROM gcr.io/distroless/base-debian11 AS run-container

WORKDIR /

COPY --from=build /store-api /store-api

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/store-api"]