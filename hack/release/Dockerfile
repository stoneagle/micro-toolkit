FROM golang:1.10

COPY ./release/backend /go/src/toolkit/backend
COPY ./backend/static /go/src/toolkit/static
COPY ./release/tool /go/src/toolkit/tool

WORKDIR /go/src/toolkit

CMD ["./backend"]
