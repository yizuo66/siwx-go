package chain

import (
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/crypto"
)

// Ethereum secp256k1 验证
func VerifyEthereumSignature(address, signature, message string) (bool, error) {
	sig := strings.TrimPrefix(signature, "0x")
	sigBytes, err := hex.DecodeString(sig)
	if err != nil {
		return false, fmt.Errorf("invalid hex signature: %v", err)
	}
	prefixed := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(message), message)
	hash := crypto.Keccak256Hash([]byte(prefixed))
	if sigBytes[64] >= 27 {
		sigBytes[64] -= 27
	}
	pubKey, err := crypto.SigToPub(hash.Bytes(), sigBytes)
	if err != nil {
		return false, err
	}
	recovered := crypto.PubkeyToAddress(*pubKey)
	return strings.EqualFold(recovered.Hex(), address), nil
}
