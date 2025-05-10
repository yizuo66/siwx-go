package siwx

import "testing"

func TestVerify(t *testing.T) {
	// 示例测试用例，可以扩展更多实际测试
	req := VerifyRequest{
		CAIP122Message: "caip122://eip155:1/0x1234abcd?nonce=abc123",
		Signature:      "0x...",
		PublicKey:      "",
	}
	result := Verify(req)
	if result.Valid {
		t.Log("验证成功")
	} else {
		t.Logf("验证失败: %v", result.Error)
	}
}
