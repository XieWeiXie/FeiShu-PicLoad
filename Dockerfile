FROM ubuntu:20.04

ARG NAME=FeiShuPicLoad

ENV DEBIAN_FRONTEND=noninteractive

RUN sed -i 's#http://deb.debian.org#https://mirrors.163.com#g' /etc/apt/sources.list

RUN apt-get update \
    && apt-get install -y --no-install-recommends tzdata ca-certificates \
    && rm -rf /var/lib/apt/lists/*

RUN ln -fs /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && dpkg-reconfigure --frontend noninteractive tzdata

RUN update-ca-certificates

RUN export GO111MODULE=on && echo go version &&go build -o ${NAME} main.go

COPY ./${NAME} app/${NAME}

CMD ./${NAME}

