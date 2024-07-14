FROM postgres:latest

COPY ./init.sql /docker-entrypoint-initdb.d/init.sql

ENV POSTGRES_USER=user
ENV POSTGRES_PASSWORD=password
ENV POSTGRES_DB=mybarbershopdb