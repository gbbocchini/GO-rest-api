FROM golang:1.18
WORKDIR /home/gbbocchini/Documents/GoLang/GO-rest-api
COPY . .
COPY go.mod .
RUN go mod download
COPY *.go .
RUN go build -v
CMD ["./rest"]