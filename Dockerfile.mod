# dependencies
FROM    golang:1.18-alpine AS dependencies
RUN     mkdir -p /opt/app-code
COPY    ./go.mod /opt/app-code
COPY    ./go.sum /opt/app-code
WORKDIR /opt/app-code
RUN     go mod download