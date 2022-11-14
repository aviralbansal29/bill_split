FROM golang:1.18.0-alpine3.14

RUN apk add build-base bash git

WORKDIR /app

# copy dependency 
COPY go.mod . 
COPY go.sum .
RUN go mod download

COPY . /app

WORKDIR /app

# Go Build
RUN go build -o /bill_split

# Port
EXPOSE 9000

# Run
CMD ["/bill_split"]
