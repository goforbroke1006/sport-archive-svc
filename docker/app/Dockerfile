FROM debian:jessie

MAINTAINER Sergey Cherkesov <go.for.broke1006@gmail.com>

ARG BINARY_LOCATION
ARG BINARY_NAME

COPY ${BINARY_LOCATION} /usr/local/bin/${BINARY_NAME}

WORKDIR /app/

VOLUME /app/logs
VOLUME /app/conf
VOLUME /app/data

EXPOSE 8080

ENTRYPOINT [ "/usr/local/bin/${BINARY_NAME}", \
    "-db-conn",     "/app/data/sport-archive.db", \
    "-serve-addr",  ":8080", \
    "-allow-save",  "true", \
    "-log-path",    "/app/logs/access.log", \
    "" ]
