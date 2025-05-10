package chain

import (
	"crypto/ed25519"
	"encoding/base64"
	"fmt"
)

// Sui ed25519 验证
func VerifySuiSignature(pubKeyBase64, signatureBase64, message string) (bool, error) {
	pubKeyBytes, err := base64.StdEncoding.DecodeString(pubKeyBase64)
	if err != nil {
		return false, fmt.Errorf("invalid sui public key: %v", err)
	}
	sigBytes, err := base64.StdEncoding.DecodeString(signatureBase64)
	if err != nil {
		return false, fmt.Errorf("invalid sui signature: %v", err)
	}
	if len(pubKeyBytes) != ed25519.PublicKeySize {
		return false, fmt.Errorf("invalid sui public key size")
	}
	return ed25519.Verify(pubKeyBytes, []byte(message), sigBytes), nil
}
