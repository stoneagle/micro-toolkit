FROM golang:1.10

COPY ./release/backend /go/src/toolkit/backend
COPY ./backend/static /go/src/toolkit/static
COPY ./release/tool /go/src/toolkit/tool
COPY ./hack/release/config.yaml /go/src/config/config.yaml

WORKDIR /go/src/toolkit

CMD ["./backend"]
