package chain

import (
	"crypto/ed25519"
	"encoding/base64"
	"fmt"

	"github.com/mr-tron/base58"
)

// Solana ed25519 验证
func VerifySolanaSignature(pubKeyBase58, signatureBase64, message string) (bool, error) {
	pubKeyBytes, err := base58.Decode(pubKeyBase58)
	if err != nil {
		return false, fmt.Errorf("invalid solana public key: %v", err)
	}
	sigBytes, err := base64.StdEncoding.DecodeString(signatureBase64)
	if err != nil {
		return false, fmt.Errorf("invalid solana signature: %v", err)
	}
	if len(pubKeyBytes) != ed25519.PublicKeySize {
		return false, fmt.Errorf("invalid solana public key size")
	}
	return ed25519.Verify(pubKeyBytes, []byte(message), sigBytes), nil
}
