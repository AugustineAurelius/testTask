FROM golang:alpine
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go get -u -v ./...
RUN go install -v ./...
RUN go build -o main src/main.go
CMD /app/main

