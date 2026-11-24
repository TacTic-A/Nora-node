# Nora-node

Nora-node is the agent that runs on each node server. It controls Xray and talks to the Nora panel over mutually authenticated gRPC.

## Install

The node connects to the panel with the certificate and token flow in the panel. Run this on the node server:

```bash
curl -sL https://raw.githubusercontent.com/TacTic-A/Nora/master/scripts/nora/nora-node-binary.sh | sudo bash -s -- install
```

The installer creates a systemd service and downloads the release for the server's architecture.

- Dev channel: `curl -sL https://raw.githubusercontent.com/TacTic-A/Nora/dev/scripts/nora/nora-node-binary.sh | sudo bash -s -- install --dev`
- Second node on the same server: add `--name nora-node2` to the install command.
- Commands: `nora-node help`

## Build

CI builds binaries for every push to `dev`, every pull request, and every published release. Release assets cover Linux (`386`, `amd64`, `arm64`, `armv5`, `armv6`, `armv7`, `s390x`) and Windows `amd64`.

To build and test locally:

```bash
go run ./tools/build
go run ./tools/smoke
go test ./...
```
