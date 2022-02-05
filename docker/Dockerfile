FROM golang:1.18beta2-alpine as dev

ENV ROOT=/go/src/github.com/kichikawa
ENV PATH /usr/local/go/bin:$PATH
ENV CGO_ENABLED 0
WORKDIR ${ROOT}

RUN apk update && apk add git
COPY go.mod go.sum ./
RUN go mod download
RUN go get -u github.com/cosmtrek/air && \
    go build -o /go/bin/air github.com/cosmtrek/air
RUN go get -u entgo.io/ent/cmd/ent
EXPOSE 8080

CMD ["air", "-c", ".air.toml"]


FROM golang:1.15.7-alpine as builder

ENV ROOT=/go/src/github.com/kichikawa
WORKDIR ${ROOT}

RUN apk update && apk add git
COPY go.mod go.sum ./
RUN go mod download

COPY . ${ROOT}
RUN CGO_ENABLED=0 GOOS=linux go build -o $ROOT/binary


FROM scratch as prod

ENV ROOT=/go/src/github.com/kichikawa
WORKDIR ${ROOT}
COPY --from=builder ${ROOT}/binary ${ROOT}

EXPOSE 8080
CMD ["/go/src/github.com/kichikawa/binary"]