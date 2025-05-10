package siwx

import (
	"errors"
	"fmt"
	"net/url"
	"strings"

	caip "github.com/yizuo998/siwx-go/caip122"
)

// VerifyRequest 输入参数
type VerifyRequest struct {
	CAIP122Message string // caip122://eip155:1/0x你的钱包地址?nonce=xxx
	Signature      string // EVM: hex; Solana/Sui: base64
	PublicKey      string // Solana: base58; Sui: base64
}

// VerifyResult 返回结果
type VerifyResult struct {
	ChainID string // e.g., eip155:1
	Address string // e.g., 0x你的钱包地址
	Nonce   string // e.g., xxx
	Valid   bool   // 是否验证成功
	Error   error  // 错误信息（可为空）
}

// Verify 验证签名主入口
func Verify(req VerifyRequest) VerifyResult {
	chainID, address, nonce, err := parseCAIP122Message(req.CAIP122Message)
	if err != nil {
		return VerifyResult{Error: fmt.Errorf("parse error: %v", err)}
	}

	var ok bool
	switch {
	case strings.HasPrefix(chainID, "eip155"):
		ok, err = caip.VerifyEthereumSignature(address, req.Signature, req.CAIP122Message)
	case strings.HasPrefix(chainID, "solana"):
		ok, err = caip.VerifySolanaSignature(req.PublicKey, req.Signature, req.CAIP122Message)
	case strings.HasPrefix(chainID, "sui"):
		ok, err = caip.VerifySuiSignature(req.PublicKey, req.Signature, req.CAIP122Message)
	default:
		err = errors.New("unsupported chain")
	}

	return VerifyResult{
		ChainID: chainID,
		Address: address,
		Nonce:   nonce,
		Valid:   ok && err == nil,
		Error:   err,
	}
}

// parseCAIP122Message 解析 caip122:// 链地址消息
func parseCAIP122Message(message string) (chainID, address, nonce string, err error) {
	u, err := url.Parse(message)
	if err != nil {
		return "", "", "", err
	}
	chainID = u.Host
	address = strings.TrimPrefix(u.Path, "/")
	nonce = u.Query().Get("nonce")
	return
}
