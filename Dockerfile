FROM golang:1.23.3-alpine

WORKDIR /opt/code

# creating go.mod and go.sum copies after copying source
COPY go.mod go.sum ./
RUN go mod download

# all code copying
COPY . .

# installing dependencies
RUN apk add --no-cache git

# make binary file
RUN go build -o bin/ready-server cmd/ready-server/main.go

# clearing temporary files
RUN rm -rf /var/cache/apk/*

ENTRYPOINT [ "/opt/code/bin/ready-server" ]
