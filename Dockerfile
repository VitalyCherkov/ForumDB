# 1. Build go binary bundle

FROM golang:alpine as builder

WORKDIR /src

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
	go build -mod vendor -a -installsuffix cgo -ldflags="-w -s" -o go-boundle


# 2. Build PG + go-bundle (without sources) together

FROM ubuntu:18.04

ENV PG_VERSION=10

RUN apt-get update -y && apt-get install -y postgresql-$PG_VERSION postgresql-contrib

USER postgres

RUN    /etc/init.d/postgresql start &&\
    psql --command "CREATE USER docker WITH SUPERUSER PASSWORD 'docker';" &&\
    createdb -O docker docker

RUN echo "host all  all    0.0.0.0/0  md5" >> /etc/postgresql/$PG_VERSION/main/pg_hba.conf

RUN echo "listen_addresses='*'" >> /etc/postgresql/$PG_VERSION/main/postgresql.conf

EXPOSE 5432

# Add VOLUMEs to allow backup of config, logs and databases
VOLUME  ["/etc/postgresql", "/var/log/postgresql", "/var/lib/postgresql"]

USER root

WORKDIR /app

COPY --from=builder /src/go-boundle .

EXPOSE 5000
CMD service postgresql start && ./go-boundle
