# 阿里云 SDK 公共包

这是一个用于创建阿里云各种服务 API 客户端的 Go 语言公共包，提供了统一的客户端创建方式，简化了阿里云 SDK 的使用流程。

## 功能特性

- **统一接口**：所有服务客户端的创建遵循相同的模式，降低学习成本
- **AccessKey 认证**：支持使用阿里云 AccessKey 进行身份认证
- **模块化设计**：按服务类型组织代码，便于维护和扩展
- **开箱即用**：快速集成到现有的 Go 项目中

## 支持的服务

本包支持以下阿里云服务的客户端创建：

| 服务类型 | 包路径 | 说明 |
|---------|-------|------|
| ECS | `aliyun/ecs` | 弹性计算服务（Elastic Compute Service） |
| VPC | `aliyun/vpc` | 专有网络（Virtual Private Cloud） |
| SLB | `aliyun/loadbalance/slb` | 传统型负载均衡（Server Load Balancer） |
| ALB | `aliyun/loadbalance/alb` | 应用型负载均衡（Application Load Balancer） |
| NLB | `aliyun/loadbalance/nlb` | 网络型负载均衡（Network Load Balancer） |
| ALIDNS | `aliyun/alidns` | 云解析 DNS（Aliyun DNS） |
| ARMS | `aliyun/arms` | 应用实时监控服务（Application Real-Time Monitoring Service） |
| SAE | `aliyun/sae` | Serverless 应用引擎（Serverless App Engine） |

## 安装

使用 Go 模块进行安装：

```bash
go get bal043-aliyun-sdk/aliyun/...
```

确保在项目中已经初始化 Go 模块：

```bash
go mod init your-module-name
```

## 快速开始

### 1. 配置凭证

在使用 SDK 之前，您需要准备好阿里云 AccessKey 凭证。您可以通过以下方式获取：

1. 登录 [阿里云控制台](https://console.aliyun.com/)
2. 进入 **RAM 控制台**（访问控制）
3. 在 **用户管理** 中创建 RAM 用户或使用已有用户
4. 为用户授予必要的权限
5. 在 **安全设置** 中创建或查看 AccessKey

**安全提示**：请妥善保管您的 AccessKey，不要将其硬编码在代码中或提交到版本控制系统。建议使用环境变量或配置文件来管理敏感信息。

### 2. 使用示例

以下是各服务的使用示例：

#### ECS 客户端

```go
package main

import (
	"log"

	"bal043-aliyun-sdk/aliyun/ecs"
)

func main() {
	// 创建 ECS 客户端
	client, err := ecs.CreateClient(
		"your-access-key-id",
		"your-access-key-secret",
		"ecs.cn-hangzhou.aliyuncs.com",
	)
	if err != nil {
		log.Fatalf("创建 ECS 客户端失败: %v", err)
	}

	// 使用客户端调用 ECS API
	// 例如查询实例列表：
	// response, err := client.DescribeInstances(&request)
	_ = client
}
```

#### VPC 客户端

```go
package main

import (
	"log"

	"bal043-aliyun-sdk/aliyun/vpc"
)

func main() {
	// 创建 VPC 客户端
	client, err := vpc.CreateClient(
		"your-access-key-id",
		"your-access-key-secret",
		"vpc.cn-hangzhou.aliyuncs.com",
	)
	if err != nil {
		log.Fatalf("创建 VPC 客户端失败: %v", err)
	}

	// 使用客户端调用 VPC API
	_ = client
}
```

#### 负载均衡客户端

```go
package main

import (
	"log"

	"bal043-aliyun-sdk/aliyun/loadbalance/slb"
	"bal043-aliyun-sdk/aliyun/loadbalance/alb"
	"bal043-aliyun-sdk/aliyun/loadbalance/nlb"
)

func main() {
	// 传统型负载均衡（SLB）
	slbClient, err := slb.CreateClient(
		"your-access-key-id",
		"your-access-key-secret",
		"slb.cn-hangzhou.aliyuncs.com",
	)
	if err != nil {
		log.Fatalf("创建 SLB 客户端失败: %v", err)
	}
	_ = slbClient

	// 应用型负载均衡（ALB）
	albClient, err := alb.CreateClient(
		"your-access-key-id",
		"your-access-key-secret",
		"alb.cn-hangzhou.aliyuncs.com",
	)
	if err != nil {
		log.Fatalf("创建 ALB 客户端失败: %v", err)
	}
	_ = albClient

	// 网络型负载均衡（NLB）
	nlbClient, err := nlb.CreateClient(
		"your-access-key-id",
		"your-access-key-secret",
		"nlb.cn-hangzhou.aliyuncs.com",
	)
	if err != nil {
		log.Fatalf("创建 NLB 客户端失败: %v", err)
	}
	_ = nlbClient
}
```

#### 其他服务客户端

```go
package main

import (
	"log"

	"bal043-aliyun-sdk/aliyun/alidns"
	"bal043-aliyun-sdk/aliyun/arms"
	"bal043-aliyun-sdk/aliyun/sae"
)

func main() {
	// 云解析 DNS
	dnsClient, err := alidns.CreateClient(
		"your-access-key-id",
		"your-access-key-secret",
		"alidns.cn-hangzhou.aliyuncs.com",
	)
	if err != nil {
		log.Fatalf("创建 Alidns 客户端失败: %v", err)
	}
	_ = dnsClient

	// 应用实时监控
	armsClient, err := arms.CreateClient(
		"your-access-key-id",
		"your-access-key-secret",
		"arms.cn-hangzhou.aliyuncs.com",
	)
	if err != nil {
		log.Fatalf("创建 ARMS 客户端失败: %v", err)
	}
	_ = armsClient

	// Serverless 应用引擎
	saeClient, err := sae.CreateClient(
		"your-access-key-id",
		"your-access-key-secret",
		"sae.cn-hangzhou.aliyuncs.com",
	)
	if err != nil {
		log.Fatalf("创建 SAE 客户端失败: %v", err)
	}
	_ = saeClient
}
```

### 3. 端点说明

在创建客户端时，需要指定服务的访问端点（Endpoint）。端点的格式通常为：

```
<service-id>.<region-id>.aliyuncs.com
```

常用区域对应的端点示例：

| 区域 | 区域 ID | ECS 端点 | VPC 端点 | SLB 端点 |
|-----|--------|---------|---------|---------|
| 杭州 | cn-hangzhou | ecs.cn-hangzhou.aliyuncs.com | vpc.cn-hangzhou.aliyuncs.com | slb.cn-hangzhou.aliyuncs.com |
| 上海 | cn-shanghai | ecs.cn-shanghai.aliyuncs.com | vpc.cn-shanghai.aliyuncs.com | slb.cn-shanghai.aliyuncs.com |
| 北京 | cn-beijing | ecs.cn-beijing.aliyuncs.com | vpc.cn-beijing.aliyuncs.com | slb.cn-beijing.aliyuncs.com |
| 广州 | cn-guangzhou | ecs.cn-guangzhou.aliyuncs.com | vpc.cn-guangzhou.aliyuncs.com | slb.cn-guangzhou.aliyuncs.com |
| 深圳 | cn-shenzhen | ecs.cn-shenzhen.aliyuncs.com | vpc.cn-shenzhen.aliyuncs.com | slb.cn-shenzhen.aliyuncs.com |

完整的服务端点列表，请参考 [阿里云官方文档](https://help.aliyun.com/document_detail/40654.html)。

## 最佳实践

### 1. 凭证管理

推荐使用环境变量来管理 AccessKey 凭证：

```go
package main

import (
	"os"

	"bal043-aliyun-sdk/aliyun/ecs"
)

func main() {
	accessKeyId := os.Getenv("ALIBABACLOUD_ACCESS_KEY_ID")
	accessKeySecret := os.Getenv("ALIBABACLOUD_ACCESS_KEY_SECRET")

	if accessKeyId == "" || accessKeySecret == "" {
		panic("请设置环境变量 ALIBABACLOUD_ACCESS_KEY_ID 和 ALIBABACLOUD_ACCESS_KEY_SECRET")
	}

	client, err := ecs.CreateClient(
		accessKeyId,
		accessKeySecret,
		"ecs.cn-hangzhou.aliyuncs.com",
	)
	if err != nil {
		panic(err)
	}

	_ = client
}
```

设置环境变量：

**Linux/macOS：**

```bash
export ALIBABACLOUD_ACCESS_KEY_ID="your-access-key-id"
export ALIBABACLOUD_ACCESS_KEY_SECRET="your-access-key-secret"
```

**Windows（PowerShell）：**

```powershell
$env:ALIBABACLOUD_ACCESS_KEY_ID="your-access-key-id"
$env:ALIBABACLOUD_ACCESS_KEY_SECRET="your-access-key-secret"
```

### 2. 错误处理

建议对客户端创建过程中的错误进行处理：

```go
client, err := ecs.CreateClient(accessKeyId, accessKeySecret, endpoint)
if err != nil {
	// 记录错误日志
	log.Errorf("创建 ECS 客户端失败: %v", err)
	// 根据业务需求进行处理
	return err
}
defer func() {
	// 如果需要，可以在这里释放资源
}()
```

### 3. 区域选择

根据您的业务需求选择合适的区域：

- 选择距离用户最近的区域以降低网络延迟
- 考虑数据合规性要求（如某些数据必须存储在特定区域）
- 参考阿里云各区域的资源配额和价格差异

## API 参考

创建客户端后，您可以调用各服务提供的 API 方法。以下是官方 SDK 文档链接：

- [ECS SDK 文档](https://github.com/alibabacloud-go/ecs-20140526)
- [VPC SDK 文档](https://github.com/alibabacloud-go/vpc-20160428)
- [SLB SDK 文档](https://github.com/alibabacloud-go/slb-20140515)
- [ALB SDK 文档](https://github.com/alibabacloud-go/alb-20200616)
- [NLB SDK 文档](https://github.com/alibabacloud-go/nlb-20220430)
- [Alidns SDK 文档](https://github.com/alibabacloud-go/alidns-20150109)
- [ARMS SDK 文档](https://github.com/alibabacloud-go/arms-20190808)
- [SAE SDK 文档](https://github.com/alibabacloud-go/sae-20190506)

完整的 API 调用示例和参数说明，请参考各服务的官方文档。

## 常见问题

### 1. 权限不足

如果遇到权限错误，请检查：
- AccessKey 是否有效
- RAM 用户是否被授予了相应的权限策略
- 策略是否正确作用于当前用户

### 2. 端点错误

如果遇到端点相关的错误，请：
- 确认端点格式正确
- 确认该服务是否在指定区域开放
- 参考[服务端点列表](https://help.aliyun.com/document_detail/40654.html)

### 3. 网络连接

如果遇到网络连接问题，请：
- 检查网络连接是否正常
- 确认是否需要配置代理
- 检查防火墙规则是否允许访问阿里云服务

## 贡献指南

如果您希望为本项目贡献代码，请：

1. Fork 本仓库
2. 创建您的特性分支（`git checkout -b feature/AmazingFeature`）
3. 提交您的更改（`git commit -m 'Add some AmazingFeature'`）
4. 推送到分支（`git push origin feature/AmazingFeature`）
5. 创建一个 Pull Request

## 许可证

本项目采用 Apache License 2.0 许可证。

## 联系

如有问题或建议，请通过 Issues 页面反馈。

[![Go Reference](https://pkg.go.dev/badge/github.com/bryant-ba/bal043-aliyun-sdk.svg)](https://pkg.go.dev/github.com/bryant-ba/bal043-aliyun-sdk)