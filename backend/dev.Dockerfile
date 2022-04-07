FROM cosmtrek/air

LABEL maintaner="Fernando Constantino <const.fernando@gmail.com>"

WORKDIR /go/src/github.com/ferdn4ndo/userver-logger-slim/src

COPY ./entrypoint.dev.sh /entrypoint.dev.sh

ENTRYPOINT ["/entrypoint.dev.sh"]
