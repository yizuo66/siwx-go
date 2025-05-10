package main

import (
	"fmt"

	"github.com/yizuo998/siwx-go"
)

func main() {
	req := siwx.VerifyRequest{
		CAIP122Message: "caip122://eip155:1/0x1234abcd?nonce=abc123",
		Signature:      "0x...",
		PublicKey:      "",
	}

	result := siwx.Verify(req)

	if result.Valid {
		fmt.Println("✅ 验证成功:", result.Address, result.ChainID)
	} else {
		fmt.Println("❌ 验证失败:", result.Error)
	}
}
