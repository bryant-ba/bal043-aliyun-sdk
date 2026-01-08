package ecs

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	ecs20140526 "github.com/alibabacloud-go/ecs-20140526/v7/client"

	"bal043-aliyun-sdk/aliyun/credential"
)

// CreateClient 创建阿里云ECS客户端实例
// 用于初始化阿里云弹性计算服务（ECS）的API客户端
//
// 参数说明:
//   - accessKeyid: 阿里云访问密钥ID（RAM用户的AccessKey ID）
//   - accessKeySecret: 阿里云访问密钥（RAM用户的AccessKey Secret）
//   - endpoint: 阿里云ECS服务的访问端点（如 "ecs.cn-hangzhou.aliyuncs.com"）
//
// 返回值说明:
//   - result: ECS客户端实例指针，用于调用ECS API
//   - err: 客户端创建失败时返回错误信息，成功时返回nil
//
// 使用示例:
//
//	client, err := ecs.CreateClient("your-access-key-id", "your-access-key-secret", "ecs.cn-hangzhou.aliyuncs.com")
//	if err != nil {
//	    log.Fatalf("创建ECS客户端失败: %v", err)
//	}
//	// 使用client调用ECS API，如查询实例列表
func CreateClient(accessKeyid, accessKeySecret, endpoint string) (result *ecs20140526.Client, err error) {
	c, err := credential.CreateCredentials(accessKeyid, accessKeySecret)
	if err != nil {
		return nil, err
	}
	config := &openapi.Config{
		Credential: c,
		Endpoint:   &endpoint,
	}
	r, err := ecs20140526.NewClient(config)
	if err != nil {
		return nil, err
	}
	return r, nil
}
