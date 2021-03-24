FROM archlinux
LABEL maintainer Plankiton <pl4nk1ton@gmail.com>

RUN pacman -Syu community/go --needed --noconfirm
RUN mkdir /api/bin -p
ADD . /api

EXPOSE 8000

WORKDIR /api

CMD test -e ./bin/diamond || go build -o ./bin/diamond . && ./bin/diamond > api.log 2> api.err.log
