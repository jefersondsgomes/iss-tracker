FROM golang

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . ./

RUN go build -o /iss-tracker

CMD ["/iss-tracker"]