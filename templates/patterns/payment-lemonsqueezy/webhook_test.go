package ls

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"testing"
)

func TestLemonSqueezyVerifier(t *testing.T) {
	secret := "secret-key"
	body := []byte(`{"event":"order_created"}`)

	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write(body)
	sig := hex.EncodeToString(mac.Sum(nil))

	verifier := NewLemonSqueezyVerifier(secret)
	if !verifier.Verify(sig, body) {
		t.Error("expected valid HMAC signature to pass")
	}

	if verifier.Verify("invalid-sig", body) {
		t.Error("expected invalid HMAC signature to fail")
	}
}

func TestMockPaymentVerifier(t *testing.T) {
	verifier := NewMockPaymentVerifier()
	if !verifier.Verify("", []byte("any body")) {
		t.Error("mock verifier should always return true")
	}
}
