FROM golang:1.22 as build

WORKDIR /build

COPY . .

RUN GOOS=linux go build -o ./app .

FROM alpine:latest

COPY --from=build /build/app ./app

ENTRYPOINT [ "./app" ]