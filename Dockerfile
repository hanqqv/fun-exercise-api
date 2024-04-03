FROM golang:1.22.1-alpine as build-base

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go test -tags=unit ./...

RUN go build -o ./out/funx .


### ------------

FROM alpine:3.19
COPY --from=build-base /app/out/funx /app/funx

EXPOSE 1323

CMD ["/app/funx"]