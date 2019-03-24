FROM debian:jessie

MAINTAINER Sergey Cherkesov

ARG BINARY_NAME

COPY ./build/Realise/sport-archive-svc /usr/local/bin/${BINARY_NAME}

WORKDIR /app/

VOLUME /app/logs
VOLUME /app/conf
VOLUME /app/data

EXPOSE 8080

ENTRYPOINT [ "/usr/local/bin/${BINARY_NAME}", \
    "-db-conn",     "/app/data/sport-archive.db", \
    "-handle-addr", ":8080", \
    "-allow-save",  "true", \
    "-log-path",    "/app/logs/access.log", \
    "" ]
