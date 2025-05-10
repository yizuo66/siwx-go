package main

import (
	"fmt"

	"github.com/yizuo998/siwx-go"
)

func main() {
	fmt.Println("--- Ethereum 示例 ---")
	ethReq := siwx.VerifyRequest{
		CAIP122Message: "caip122://eip155:1/0x你的钱包地址?nonce=xxx",
		Signature:      "0x你的签名",
		PublicKey:      "",
	}
	ethResult := siwx.Verify(ethReq)
	if ethResult.Valid {
		fmt.Println("✅ 验证成功:", ethResult.Address, ethResult.ChainID)
	} else {
		fmt.Println("❌ 验证失败:", ethResult.Error)
	}

	fmt.Println("\n--- Solana 示例 ---")
	solReq := siwx.VerifyRequest{
		CAIP122Message: "caip122://solana/mainnet/你的钱包地址?nonce=xxx",
		Signature:      "你的签名（base64）",
		PublicKey:      "你的公钥（base58）",
	}
	solResult := siwx.Verify(solReq)
	if solResult.Valid {
		fmt.Println("✅ 验证成功:", solResult.Address, solResult.ChainID)
	} else {
		fmt.Println("❌ 验证失败:", solResult.Error)
	}

	fmt.Println("\n--- Sui 示例 ---")
	suiReq := siwx.VerifyRequest{
		CAIP122Message: "caip122://sui/mainnet/你的钱包地址?nonce=xxx",
		Signature:      "你的签名（base64）",
		PublicKey:      "你的公钥（base64）",
	}
	suiResult := siwx.Verify(suiReq)
	if suiResult.Valid {
		fmt.Println("✅ 验证成功:", suiResult.Address, suiResult.ChainID)
	} else {
		fmt.Println("❌ 验证失败:", suiResult.Error)
	}
}
