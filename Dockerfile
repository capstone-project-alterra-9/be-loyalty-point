FROM golang:alpine

WORKDIR /app

COPY . /app

RUN go mod tidy
RUN go mod download
RUN go build -o dist

EXPOSE 8080

ENTRYPOINT [ "./dist" ]