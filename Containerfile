ARG VERSION=latest

FROM golang:1.20-alpine as BUILD
ARG VERSION

RUN apk add --no-cache gcc g++ make

WORKDIR /app/

COPY src/go.mod src/go.sum ./
RUN go mod verify && go mod download
COPY src/ .

RUN GOOOS=linux GOARCH=amd64 go build -o demo-service -ldflags="-X main.version=$VERSION" .

FROM alpine:3.18.0

RUN apk add --no-cache ca-certificates curl wget bash

RUN addgroup -g 1000 -S demo && \
    adduser -u 1000 -S demo -G demo

RUN mkdir -p /demo/

COPY --from=BUILD /app/demo-service /bin/demo-service

RUN chgrp -R 0 /demo  && \
    chmod -R g=u /demo/ && \
    chgrp -R 0 /bin/demo-service &&  \
    chmod -R g=u /bin/demo-service

USER 1000
EXPOSE 8080
ENTRYPOINT [ "/bin/demo-service" ]
CMD [ "server" ]