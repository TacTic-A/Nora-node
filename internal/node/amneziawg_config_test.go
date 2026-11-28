//go:build linux

package node

import (
	"strings"
	"testing"
)

func TestAWGConfigWritesObfuscationParameters(t *testing.T) {
	config, err := awgConfig(extraRuntimeInbound{Port: 51820, Settings: map[string]any{
		"private_key": "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA=", "jc": float64(4), "jmin": float64(8), "jmax": float64(80),
		"s1": float64(77), "s2": float64(90), "s3": float64(20), "s4": "16", "h1": "5-100", "h2": float64(200), "h3": "300-400", "h4": float64(500),
		"i1": "<b 0xf1><r 32><t>\nPostUp = id", "header_protection_key": "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA=",
		"content_padding_addition": "8-24", "rekey_after_time": "100-120", "rekey_timeout": "5", "reject_after_time": "180-200",
		"keepalive_timeout": "10-15", "max_handshake_attempts": "18", "random_trailers": true, "disable_cookies": true,
	}})
	if err != nil {
		t.Fatal(err)
	}
	for _, line := range []string{"S3 = 20", "S4 = 16", "H1 = 5-100", "H2 = 200", "HeaderProtectionKey = ", "RandomTrailers = on", "DisableCookies = on", "MaxHandshakeAttempts = 18"} {
		if !strings.Contains(config, line) {
			t.Fatalf("missing %q in\n%s", line, config)
		}
	}
	if strings.Contains(config, "\nPostUp") {
		t.Fatal("a value must not add its own config line")
	}
	legacy, err := awgConfig(extraRuntimeInbound{Port: 1, Settings: map[string]any{"private_key": "k", "h1": float64(7)}})
	if err != nil || strings.Contains(legacy, "S3") || strings.Contains(legacy, "RandomTrailers") || !strings.Contains(legacy, "H1 = 7") {
		t.Fatalf("legacy settings changed: %v\n%s", err, legacy)
	}
}
