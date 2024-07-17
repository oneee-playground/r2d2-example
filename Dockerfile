FROM golang:1.22 as build

WORKDIR /build

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o app .

FROM alpine:latest

COPY --from=build /build .

ENTRYPOINT [ "./app" ]