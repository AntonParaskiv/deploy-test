FROM golang:latest

# mount current local dir to containter workdir /go
ADD . /go

# build
RUN go build -o app

# run
CMD "/go/app"

# mapping external to internal ports
EXPOSE 80:8081