FROM golang:1.15-buster
LABEL maintainer Plankiton <pl4nk1ton@gmail.com>

RUN mkdir /api/bin -p
ADD . /api

EXPOSE 8000

WORKDIR /api

CMD test -e ./bin/diamond || go build -o ./bin/diamond . && ./bin/diamond > api.log 2> api.err.log
