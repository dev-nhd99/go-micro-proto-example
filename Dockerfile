FROM alpine:latest

WORKDIR /apps
COPY go-micro-proto-example . 


ENTRYPOINT [ "./go-micro-proto-example" ]     