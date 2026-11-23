FROM golang:1.27-bookworm AS build

WORKDIR /src

COPY go.mod go.sum* ./
RUN go mod download

COPY . .

ARG TARGETOS=linux
ARG TARGETARCH=amd64
RUN CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH go build -trimpath -ldflags="-s -w" -o /out/nora-node ./cmd/nora-node

FROM debian:bookworm-slim AS xray

RUN apt-get update \
    && apt-get install -y --no-install-recommends ca-certificates curl unzip bash \
    && curl -L https://raw.githubusercontent.com/norapanel/Nora/master/scripts/nora/install_latest_xray.sh | bash \
    && rm -rf /var/lib/apt/lists/*

FROM debian:bookworm-slim

RUN apt-get update \
    && apt-get install -y --no-install-recommends ca-certificates haproxy \
    && rm -rf /var/lib/apt/lists/*

COPY --from=build /out/nora-node /usr/local/bin/nora-node
COPY --from=xray /usr/local/bin/xray /usr/local/bin/xray
COPY --from=xray /usr/local/share/xray /usr/local/share/xray

WORKDIR /code

CMD ["nora-node"]
