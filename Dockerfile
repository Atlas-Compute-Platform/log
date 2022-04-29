# Atlas Dockerfile
FROM alpine:latest
RUN apk add go
COPY ./log /usr/local/bin/log
EXPOSE 8800/tcp
CMD ["/usr/local/bin/log"]
