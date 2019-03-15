# See https://hub.docker.com/_/golang
FROM golang:alpine

# Label the application
LABEL Description="Wiki application"

# Install git to habe ability to download dependencies
RUN apk update && apk add git

WORKDIR /go/src/github.com/dn23rus/wiki-v2
COPY . .

# Download dependencies
RUN go get -d -v ./...
RUN go install -v ./...

ADD deployments/docker/app/entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

EXPOSE 8001

ENTRYPOINT ["/entrypoint.sh"]
CMD ["wiki-app"]
