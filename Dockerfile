FROM ubuntu:18.04

RUN mkdir /app
COPY main /bin/monorepo-operator

RUN apt-get update
RUN apt-get install -y \
        wget \
        git
