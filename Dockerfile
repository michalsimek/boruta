FROM golang:1.11.1-stretch
LABEL maintainer="Alexander Mazuruk <a.mazuruk@samsung.com>"

ENV PROJECT="github.com/SamsungSLAV/boruta"

RUN go get -d "${PROJECT}/..."

WORKDIR "${GOPATH}/src/${PROJECT}"

# Build Boruta server.
RUN go build -o /boruta cmd/boruta/boruta.go

# Build Dryad agents.
RUN GOOS=linux GOARCH=arm GOARM=7 go build -o /dryad_armv7 cmd/dryad/dryad.go
RUN GOOS=linux GOARCH=amd64 go build -o /dryad_amd64 cmd/dryad/dryad.go
