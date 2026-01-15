# 阿里云 SDK 公共包

这是一个用于创建阿里云各种服务 API 客户端的 Go 语言公共包，提供了统一的客户端创建方式、资源模型和操作封装，简化了阿里云 SDK 的使用流程，特别适合作为 CMDB 系统的依赖库。

## 功能特性

- **统一资源模型**：定义了统一的云资源接口和多种资源模型，便于 CMDB 系统进行资源管理和分析
- **统一操作接口**：所有资源操作遵循相同的接口规范，降低学习成本
- **配置管理**：支持多账号、多区域配置，支持 YAML 和 INI 两种配置格式
- **模块化设计**：按服务类型组织代码，便于维护和扩展
- **AccessKey 认证**：支持使用阿里云 AccessKey 进行身份认证
- **缓存支持**：可配置的资源缓存机制，提高查询性能
- **开箱即用**：快速集成到现有的 Go 项目中

## 架构设计

```
aliyun-sdk/
├── aliyun/
│   ├── credential/       # 凭证管理（已有）
│   ├── ecs/              # ECS 客户端和操作封装
│   ├── rds/              # RDS 客户端（已有）
│   ├── vpc/              # VPC 客户端（已有）
│   ├── loadbalance/      # 负载均衡（已有）
│   ├── alidns/           # DNS（已有）
│   ├── arms/             # 监控（已有）
│   ├── sae/              # SAE（已有）
│   ├── models/           # 统一资源模型
│   ├── config/           # 配置管理
│   └── operation/        # 资源操作封装
├── go.mod
├── README.md
└── LICENSE
```

### 核心模块说明

| 模块 | 功能 | 说明 |
|------|------|------|
| models | 统一资源模型 | 定义了 CloudResource 接口和多种云资源模型 |
| config | 配置管理 | 支持多账号、多区域配置，支持 YAML 和 INI 格式 |
| operation | 资源操作封装 | 定义了资源操作的统一接口和相关组件 |
| credential | 凭证管理 | 负责创建阿里云访问凭证 |
| 各服务包 | 客户端实现 | 各云服务的客户端创建和操作封装 |

## 支持的服务

本包支持以下阿里云服务的客户端创建：

| 服务类型 | 包路径 | 说明 |
|---------|-------|------|
| ECS | `aliyun/ecs` | 弹性计算服务（Elastic Compute Service） |
| VPC | `aliyun/vpc` | 专有网络（Virtual Private Cloud） |
| RDS | `aliyun/rds` | 关系型数据库（Relational Database Service） |
| SLB | `aliyun/loadbalance/slb` | 传统型负载均衡（Server Load Balancer） |
| ALB | `aliyun/loadbalance/alb` | 应用型负载均衡（Application Load Balancer） |
| NLB | `aliyun/loadbalance/nlb` | 网络型负载均衡（Network Load Balancer） |
| ALIDNS | `aliyun/alidns` | 云解析 DNS（Aliyun DNS） |
| ARMS | `aliyun/arms` | 应用实时监控服务（Application Real-Time Monitoring Service） |
| SAE | `aliyun/sae` | Serverless 应用引擎（Serverless App Engine） |

## 安装

使用 Go 模块进行安装：

```bash
go get github.com/bryant-ba/bal043-aliyun-sdk/aliyun/...
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

#### 2.1 传统客户端使用方式

```go
package main

import (
	"log"

	"github.com/bryant-ba/bal043-aliyun-sdk/aliyun/ecs"
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

#### 2.2 使用统一资源模型和操作封装

```go
package main

import (
	"context"
	"log"

	"github.com/bryant-ba/bal043-aliyun-sdk/aliyun/ecs"
)

func main() {
	// 1. 创建 ECS 客户端
	client, err := ecs.CreateClient(
		"your-access-key-id",
		"your-access-key-secret",
		"ecs.cn-hangzhou.aliyuncs.com",
	)
	if err != nil {
		log.Fatalf("创建 ECS 客户端失败: %v", err)
	}

	// 2. 创建 ECS 操作实例
	ecsOperator := ecs.NewECSOperator(client, "cn-hangzhou")

	// 3. 查询 ECS 实例列表
	ctx := context.Background()
	resources, err := ecsOperator.ListResources(ctx, "cn-hangzhou")
	if err != nil {
		log.Fatalf("查询 ECS 实例失败: %v", err)
	}

	// 4. 处理查询结果
	log.Printf("共查询到 %d 个 ECS 实例", len(resources))
	for _, resource := range resources {
		instance, ok := resource.(*ecs.ECSInstance)
		if ok {
			log.Printf("实例 ID: %s, 名称: %s, 状态: %s", 
				instance.InstanceId, instance.InstanceName, instance.Status)
		}
	}
}
```

#### 2.3 使用配置管理

```go
package main

import (
	"log"

	"github.com/bryant-ba/bal043-aliyun-sdk/aliyun/config"
)

func main() {
	// 1. 加载配置文件
	configManager, err := config.NewConfigManager("config.yaml")
	if err != nil {
		log.Fatalf("加载配置文件失败: %v", err)
	}

	// 2. 获取配置信息
	cfg := configManager.GetConfig()
	log.Printf("默认区域: %s", cfg.DefaultRegion)

	// 3. 获取账号配置
	accounts := configManager.GetEnabledAccounts()
	log.Printf("启用的账号数量: %d", len(accounts))

	for _, account := range accounts {
		log.Printf("账号名称: %s, 区域列表: %v", account.Name, account.Regions)
	}

	// 4. 获取资源类型配置
	resourceTypes := configManager.GetResourceTypes()
	log.Printf("关注的资源类型: %v", resourceTypes)
}
```

**配置文件示例（YAML格式）：**

```yaml
# 阿里云SDK配置文件
# 默认区域ID
default_region: cn-hangzhou

# 关注的资源类型列表
resource_types:
  - ECS
  - RDS
  - VPC
  - SLB
  - ALB
  - NLB
  - SecurityGroup

# 阿里云账号列表
accounts:
  - name: prod
    access_key_id: your-prod-access-key-id
    access_key_secret: your-prod-access-key-secret
    regions:
      - cn-hangzhou
      - cn-shanghai
      - cn-beijing
    enabled: true

  - name: dev
    access_key_id: your-dev-access-key-id
    access_key_secret: your-dev-access-key-secret
    regions:
      - cn-hangzhou
    enabled: true

# 缓存配置
cache:
  enabled: true
  expire_time: 3600
  storage_type: memory
  # redis配置（当storage_type为redis时）
  redis_address: localhost:6379
  redis_password: ""
  redis_db: 0
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

	"github.com/bryant-ba/bal043-aliyun-sdk/aliyun/ecs"
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

建议对客户端创建和API调用过程中的错误进行处理：

```go
client, err := ecs.CreateClient(accessKeyId, accessKeySecret, endpoint)
if err != nil {
	// 记录错误日志
	log.Errorf("创建 ECS 客户端失败: %v", err)
	// 根据业务需求进行处理
	return err
}

ctx := context.Background()
resources, err := ecsOperator.ListResources(ctx, regionId)
if err != nil {
	// 记录错误日志
	log.Errorf("查询 ECS 实例失败: %v", err)
	// 根据业务需求进行处理
	return err
}
```

### 3. 区域选择

根据您的业务需求选择合适的区域：

- 选择距离用户最近的区域以降低网络延迟
- 考虑数据合规性要求（如某些数据必须存储在特定区域）
- 参考阿里云各区域的资源配额和价格差异

## API 参考

### 核心接口

#### CloudResource 接口

```go
// CloudResource 阿里云资源通用接口
type CloudResource interface {
	GetResourceId() string
	GetResourceType() string
	GetResourceName() string
	GetRegionId() string
}
```

#### ResourceOperation 接口

```go
// ResourceOperation 资源操作统一接口
type ResourceOperation interface {
	ListResources(ctx context.Context, regionId string) ([]models.CloudResource, error)
	GetResourceById(ctx context.Context, regionId, resourceId string) (models.CloudResource, error)
	ListResourcesByTag(ctx context.Context, regionId, tagKey, tagValue string) ([]models.CloudResource, error)
	ListResourcesByTags(ctx context.Context, regionId string, tags map[string]string) ([]models.CloudResource, error)
	GetResourceTags(ctx context.Context, regionId, resourceId string) (map[string]string, error)
	TagResource(ctx context.Context, regionId, resourceId string, tags map[string]string) error
	UntagResource(ctx context.Context, regionId, resourceId string, tagKeys []string) error
}
```

### 官方 SDK 文档

创建客户端后，您可以调用各服务提供的 API 方法。以下是官方 SDK 文档链接：

- [ECS SDK 文档](https://github.com/alibabacloud-go/ecs-20140526)
- [VPC SDK 文档](https://github.com/alibabacloud-go/vpc-20160428)
- [RDS SDK 文档](https://github.com/alibabacloud-go/rds-20140815)
- [SLB SDK 文档](https://github.com/alibabacloud-go/slb-20140515)
- [ALB SDK 文档](https://github.com/alibabacloud-go/alb-20200616)
- [NLB SDK 文档](https://github.com/alibabacloud-go/nlb-20220430)
- [Alidns SDK 文档](https://github.com/alibabacloud-go/alidns-20150109)
- [ARMS SDK 文档](https://github.com/alibabacloud-go/arms-20190808)
- [SAE SDK 文档](https://github.com/alibabacloud-go/sae-20190506)

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

## 测试

### 运行单元测试

```bash
go test -v ./aliyun/...
```

### 运行集成测试

```bash
go test -v -tags integration ./aliyun/...
```

## 贡献指南

如果您希望为本项目贡献代码，请：

1. Fork 本仓库
2. 创建您的特性分支（`git checkout -b feature/AmazingFeature`）
3. 提交您的更改（`git commit -m 'Add some AmazingFeature'`）
4. 推送到分支（`git push origin feature/AmazingFeature`）
5. 创建一个 Pull Request

## 版本信息

### v1.0.0

- 初始版本
- 支持 ECS、VPC、RDS、SLB、ALB、NLB、ALIDNS、ARMS、SAE 等服务的客户端创建
- 实现了统一资源模型和操作封装
- 实现了配置管理功能

## 许可证

本项目采用 Apache License 2.0 许可证。

## 联系

如有问题或建议，请通过 Issues 页面反馈。

[![Go Reference](https://pkg.go.dev/badge/github.com/bryant-ba/bal043-aliyun-sdk.svg)](https://pkg.go.dev/github.com/bryant-ba/bal043-aliyun-sdk)
