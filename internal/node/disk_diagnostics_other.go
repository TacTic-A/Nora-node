//go:build !linux

package node

import nodev1 "github.com/norapanel/nora-node/internal/proto/node/v1"

func diskDiagnostics(string) []*nodev1.ProtocolStatus { return nil }
