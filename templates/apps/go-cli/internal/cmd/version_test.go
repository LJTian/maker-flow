package cmd

import (
	"bytes"
	"testing"
)

func TestVersionCmd(t *testing.T) {
	b := new(bytes.Buffer)
	rootCmd.SetOut(b)
	rootCmd.SetArgs([]string{"version"})

	if err := rootCmd.Execute(); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !bytes.Contains(b.Bytes(), []byte("go-cli")) {
		t.Errorf("expected version output to contain 'go-cli', got %s", b.String())
	}
}
