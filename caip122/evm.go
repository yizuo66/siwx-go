package caip122

import (
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/crypto"
)

// VerifyEthereumSignature 验证以太坊 secp256k1 签名
// 参数:
//   - address: 以太坊钱包地址，用于比对恢复出的公钥地址
//   - signature: 签名，十六进制字符串，可选带 0x 前缀
//   - message: 原始消息，用于构造以太坊签名前缀
//
// 返回:
//   - bool: 验证是否成功
//   - error: 错误信息，如签名格式错误、恢复公钥失败等
func VerifyEthereumSignature(address, signature, message string) (bool, error) {
	// 去掉签名中的 0x 前缀
	sig := strings.TrimPrefix(signature, "0x")
	// 将十六进制签名转换为字节数组
	sigBytes, err := hex.DecodeString(sig)
	if err != nil {
		return false, fmt.Errorf("invalid hex signature: %v", err)
	}
	// 构造以太坊签名前缀，格式为 "\x19Ethereum Signed Message:\n" + 消息长度 + 消息内容
	prefixed := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(message), message)
	// 计算 Keccak256 哈希
	hash := crypto.Keccak256Hash([]byte(prefixed))
	// 如果签名的 v 值大于等于 27，则减去 27 以适配以太坊签名格式
	if sigBytes[64] >= 27 {
		sigBytes[64] -= 27
	}
	// 从签名和哈希中恢复出公钥
	pubKey, err := crypto.SigToPub(hash.Bytes(), sigBytes)
	if err != nil {
		return false, err
	}
	// 从公钥中恢复出以太坊地址
	recovered := crypto.PubkeyToAddress(*pubKey)
	// 比对恢复出的地址与传入的地址是否一致（忽略大小写）
	return strings.EqualFold(recovered.Hex(), address), nil
}
