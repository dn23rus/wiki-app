# Dockerfile References: https://docs.docker.com/engine/reference/builder/
FROM golang:alpine as builder

# Used for command: docker image prune --filter label=stage=intermediate
LABEL stage=intermediate

# Install git to habe ability to download dependencies
RUN apk update && apk add git

WORKDIR /go/src/github.com/dn23rus/wiki-v2
COPY . .

# Download dependencies
RUN go get -d -v ./...
#RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/wiki-app .
RUN go install -v ./...

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /app/

COPY --from=builder /go/bin/wiki-app .
COPY --from=builder /go/bin/wiki-app-setup .
COPY ./configs ./configs

ADD deployments/docker/app-optimized/entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

EXPOSE 8001

ENTRYPOINT ["/entrypoint.sh"]
CMD ["/app/wiki-app"]
