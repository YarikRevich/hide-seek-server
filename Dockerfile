FROM golang:latest
RUN mkdir src/GameServer
WORKDIR /go/src/GameServer
ENV GOBIN /go/bin
COPY . .
RUN go get github.com/mbndr/figlet4go
RUN go build main.go
CMD ["./main"]
