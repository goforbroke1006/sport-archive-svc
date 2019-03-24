FROM debian:jessie

MAINTAINER Sergey Cherkesov <go.for.broke1006@gmail.com>

ARG BINARY_LOCATION
ARG BINARY_NAME

#COPY build/Realise/sport-archive-svc /app/bin/

COPY ${BINARY_LOCATION} /usr/local/bin/${BINARY_NAME}

RUN mkdir -p /app/bin/
RUN mkdir -p /app/logs/ && chmod -R 0777 /app/logs/
RUN ln -sf /usr/local/bin/${BINARY_NAME} /app/bin/application

WORKDIR /app/

VOLUME /app/logs
VOLUME /app/conf
VOLUME /app/data

EXPOSE 8080

CMD /app/bin/application --db-conn=/app/data/sport-archive.db --log-path=/app/logs/access.log
