package caip122

import (
	"crypto/ed25519"
	"encoding/base64"
	"fmt"

	"github.com/mr-tron/base58"
)

// VerifySolanaSignature 验证 Solana ed25519 签名
// 参数:
//   - pubKeyBase58: Solana 公钥，base58 编码
//   - signatureBase64: 签名，base64 编码
//   - message: 原始消息，用于验证签名
//
// 返回:
//   - bool: 验证是否成功
//   - error: 错误信息，如公钥格式错误、签名格式错误、公钥长度不符合 ed25519 要求等
func VerifySolanaSignature(pubKeyBase58, signatureBase64, message string) (bool, error) {
	// 将 base58 编码的公钥转换为字节数组
	pubKeyBytes, err := base58.Decode(pubKeyBase58)
	if err != nil {
		return false, fmt.Errorf("invalid solana public key: %v", err)
	}
	// 将 base64 编码的签名转换为字节数组
	sigBytes, err := base64.StdEncoding.DecodeString(signatureBase64)
	if err != nil {
		return false, fmt.Errorf("invalid solana signature: %v", err)
	}
	// 检查公钥长度是否符合 ed25519 要求
	if len(pubKeyBytes) != ed25519.PublicKeySize {
		return false, fmt.Errorf("invalid solana public key size")
	}
	// 使用 ed25519 验证签名
	return ed25519.Verify(pubKeyBytes, []byte(message), sigBytes), nil
}
