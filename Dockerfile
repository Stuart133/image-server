FROM golang:bullseye as build

RUN apt-get update && apt-get install -y libvips-dev

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY ./ ./

RUN go build -o /image-server

FROM ubuntu:22.04
RUN apt-get update && apt-get install -y libvips
COPY --from=build /image-server /usr/local/bin/image-server

EXPOSE 80

CMD [ "/usr/local/bin/image-server" ]