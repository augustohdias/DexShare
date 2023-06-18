FROM golang:bookworm
WORKDIR /app
COPY . .

RUN go build -o app ./src
ENV GIN_MODE release
EXPOSE 8080
CMD ./app
