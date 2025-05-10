// 使用 ethers.js 进行签名
import { ethers } from 'ethers';

// 生成 SIWX 消息
function generateSIWXMessage(address, nonce) {
    return `caip122://eip155:1/${address}?nonce=${nonce}`;
}

// 签名消息
async function signMessage(message, privateKey) {
    const wallet = new ethers.Wallet(privateKey);
    const signature = await wallet.signMessage(message);
    return signature;
}

// 验证签名
async function verifySignature(message, signature, publicKey) {
    try {
        const response = await fetch('http://your-backend-url/verify', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                caip122Message: message,
                signature: signature,
                publicKey: publicKey
            })
        });
        
        const result = await response.json();
        return result.valid;
    } catch (error) {
        console.error('验证失败:', error);
        return false;
    }
}

// 使用示例
async function example() {
    // 用户钱包地址
    const address = '0x1234...';
    // 随机生成的 nonce
    const nonce = 'abc123';
    // 用户私钥（实际使用时应该从钱包获取）
    const privateKey = 'your-private-key';
    
    // 1. 生成消息
    const message = generateSIWXMessage(address, nonce);
    console.log('生成的消息:', message);
    
    // 2. 签名消息
    const signature = await signMessage(message, privateKey);
    console.log('签名结果:', signature);
    
    // 3. 验证签名
    const isValid = await verifySignature(message, signature, address);
    console.log('验证结果:', isValid);
}

// React 组件示例
import React, { useState } from 'react';
import { ethers } from 'ethers';

function SIWXLogin() {
    const [address, setAddress] = useState('');
    const [signature, setSignature] = useState('');
    
    const handleLogin = async () => {
        try {
            // 检查是否安装了 MetaMask
            if (typeof window.ethereum === 'undefined') {
                alert('请安装 MetaMask!');
                return;
            }
            
            // 请求用户连接钱包
            const provider = new ethers.providers.Web3Provider(window.ethereum);
            await provider.send("eth_requestAccounts", []);
            const signer = provider.getSigner();
            const userAddress = await signer.getAddress();
            setAddress(userAddress);
            
            // 生成随机 nonce
            const nonce = Math.random().toString(36).substring(7);
            
            // 生成消息
            const message = generateSIWXMessage(userAddress, nonce);
            
            // 请求用户签名
            const signature = await signer.signMessage(message);
            setSignature(signature);
            
            // 验证签名
            const isValid = await verifySignature(message, signature, userAddress);
            if (isValid) {
                alert('登录成功！');
            } else {
                alert('验证失败！');
            }
        } catch (error) {
            console.error('登录失败:', error);
            alert('登录失败: ' + error.message);
        }
    };
    
    return (
        <div>
            <button onClick={handleLogin}>
                使用钱包登录
            </button>
            {address && <p>钱包地址: {address}</p>}
            {signature && <p>签名: {signature}</p>}
        </div>
    );
}

export default SIWXLogin; 