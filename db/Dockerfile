FROM mysql:latest

ENV MYSQL_ROOT_PASSWORD=rootpass
ENV MYSQL_DATABASE=test

COPY ./init.sql /docker-entrypoint-initdb.d/
