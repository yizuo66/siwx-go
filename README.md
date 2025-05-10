# siwx-go

**siwx-go** 是一个用于 Go 语言的 CAIP-122 钱包登录校验模块，支持多链签名验证，包括 Ethereum (EIP-155)、Solana 和 Sui，帮助你快速实现基于钱包的去中心化身份登录。

---

## 功能亮点

- ✅ **CAIP-122 消息解析**  
- ✅ **Ethereum (EIP-155, secp256k1) 签名验证**  
- ✅ **Solana (ed25519) 签名验证**  
- ✅ **Sui (ed25519) 签名验证**  
- ✅ **模块化设计，易于集成**  
- ✅ **清晰的错误信息与返回结构**

---

## 安装

通过 `go get` 安装 `siwx-go`：

```bash 
go get github.com:yizuo998/siwx-go.git
```

## 快速开始

### 前端示例代码
详细的示例代码请参考 [examples/frontend/SIWXLogin.jsx](examples/frontend/SIWXLogin.jsx)

### 后端示例代码
详细的示例代码请参考 [examples/backend/main.go](examples/backend/main.go)

### 参数说明

| 参数            | 类型   | 描述                                                     |
|-----------------|--------|----------------------------------------------------------|
| `CAIP122Message`| string | CAIP-122 格式消息，例如 `caip122://eip155:1/0x1234abcd?nonce=abc123` |
| `Signature`     | string | Ethereum 签名为 hex 格式（带 `0x`）；Solana/Sui 签名为 base64 格式 |
| `PublicKey`     | string | Solana 公钥为 base58 编码，Sui 公钥为 base64 编码，Ethereum 留空 |

### 返回值结构

| 字段     | 类型   | 描述                   |
|----------|--------|------------------------|
| `ChainID`| string | 链 ID，例如 `eip155:1` |
| `Address`| string | 钱包地址               |
| `Nonce`  | string | 非法请求的 nonce 值    |
| `Valid`  | bool   | 签名是否验证成功       |
| `Error`  | error  | 错误信息（如有）       |

---

## 支持的链

| 链名称   | ChainID 格式  | 签名算法    | 公钥格式      |
|----------|---------------|-------------|---------------|
| Ethereum | `eip155:1`    | `secp256k1` | -             |
| Solana   | `solana:mainnet` | `ed25519`  | base58        |
| Sui      | `sui:mainnet` | `ed25519`   | base64        |

---

## 开发与测试

### 运行测试

```bash
go test ./...
```

### 构建项目
```bash
make build
```

### 清理和整理模块
```bash
make tidy
```


## 贡献与支持
欢迎提交 Issue 和 Pull Request，贡献你的代码或改进建议！

GitHub: https://github.com/yizuo998/siwx-go#

Email: yizuo998@gmail.com
