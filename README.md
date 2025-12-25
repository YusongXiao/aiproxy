# AIProxy

这是一个简单的反向代理服务，用于转发请求到 Google Gemini API (`https://generativelanguage.googleapis.com`)。

## 功能

- 将请求代理到 Google Gemini API。
- 自动处理 Host header 以适配 Google 的服务器要求。
- 提供简单的静态页面 (`index.html`) 和 `robots.txt`。

## 快速开始 (Docker)

你可以直接使用已构建好的 Docker 镜像来运行此服务：

```bash
docker run -d -p 8080:8080 uswccr.ccs.tencentyun.com/songhappy/aiproxy
```

服务启动后，将在本地的 `8080` 端口监听。

## 安全警告 ⚠️

**请务必注意 API Key 的安全性！**

此代理服务会原样转发包含 API Key 的请求。为了确保您的 API Key 不会在传输过程中被截获，**请务必在部署此服务时配置 TLS (HTTPS) 证书**。

如果不使用 HTTPS，您的 API Key 将以明文形式在网络中传输，存在极大的安全风险。建议使用 Nginx、Caddy 或云厂商的负载均衡器在前端进行 SSL/TLS 卸载。
