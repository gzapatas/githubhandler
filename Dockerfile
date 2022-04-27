#builder
FROM    golang:1.18-alpine AS builder
COPY    ./ /opt/app-code
WORKDIR /opt/app-code
RUN     go mod download
RUN     go build -v -o ./dist/app *.go

# dist 
FROM    alpine:3.14 AS dist
COPY    --from=builder /opt/app-code/dist/. /opt/app-code
WORKDIR /opt/app-code
RUN     touch .env
#EXPOSE  5555
CMD     ["./app"]