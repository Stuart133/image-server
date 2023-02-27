FROM golang:bullseye

RUN apt-get update && apt-get install -y libvips-dev

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY ./ ./

RUN go build -o /image-server

EXPOSE 80

CMD [ "/image-server" ]