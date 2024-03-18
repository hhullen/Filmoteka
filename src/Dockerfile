FROM golang:1.22.1-alpine3.19 AS builder

WORKDIR /application
COPY . .
RUN go work sync && go build -o /application/server .


FROM alpine:3.19

WORKDIR /application
COPY --from=builder /application/server .
ENTRYPOINT [ "./server" ]
