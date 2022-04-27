# dependencies
FROM    golang:1.16-alpine AS dependencies
RUN     mkdir -p /opt/app-code
COPY    ./go.mod /opt/app-code
COPY    ./go.sum /opt/app-code
WORKDIR /opt/app-code
RUN     go mod download

#builder
FROM    dependencies AS builder
ARG     app
COPY    ./ /opt/app-code
WORKDIR /opt/app-code/${app}
RUN     go build -v -o ./dist/app *.go

# dist 
FROM    alpine:3.14 AS dist
ARG     app
ARG     install_utilities
COPY    --from=builder /opt/app-code/${app}/dist/. /opt/app-code
WORKDIR /opt/app-code
RUN     touch .env
#EXPOSE  3000
CMD     ["./app"]