FROM golang:1.23.8-bookworm AS builder

RUN set -ex && \
	apt-get install -y curl && \
	curl -L https://github.com/harelba/q/releases/download/latest/linux-q \
		-o /usr/local/bin/q && \
	chmod +x /usr/local/bin/q

WORKDIR /go/src/app
COPY go.mod go.sum .
RUN --mount=type=cache,target=/go/pkg \
	set -ex && \
	go mod download
COPY . .
RUN --mount=type=cache,target=/go/pkg \
	set -ex && \
	go build -o app .


FROM debian:bookworm-slim AS executor

COPY --from=builder /usr/local/bin/q /usr/local/bin/q
COPY --from=builder /go/src/app/app /opt/app/

WORKDIR /opt/app

ENTRYPOINT ["./app"]
