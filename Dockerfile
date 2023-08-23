FROM golang:1.21-bookworm as builder

WORKDIR /app

ARG BINARY
ARG GIT_TAG
ARG GIT_COMMIT

ENV CGO_ENABLED=0

COPY . .

RUN make build


FROM scratch

ARG BINARY

COPY --from=builder /build  /usr/local/bin/act
