package sae

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	sae20190506 "github.com/alibabacloud-go/sae-20190506/v4/client"

	"bal043-aliyun-sdk/aliyun/credential"
)

// CreateClient 创建阿里云Serverless应用引擎客户端实例
// 用于初始化阿里云Serverless应用引擎（SAE）的API客户端
//
// 参数说明:
//   - accessKeyid: 阿里云访问密钥ID（RAM用户的AccessKey ID）
//   - accessKeySecret: 阿里云访问密钥（RAM用户的AccessKey Secret）
//   - endpoint: 阿里云SAE服务的访问端点（如 "sae.cn-hangzhou.aliyuncs.com"）
//
// 返回值说明:
//   - result: SAE客户端实例指针，用于调用Serverless应用引擎API
//   - err: 客户端创建失败时返回错误信息，成功时返回nil
//
// 使用示例:
//
//	client, err := sae.CreateClient("your-access-key-id", "your-access-key-secret", "sae.cn-hangzhou.aliyuncs.com")
//	if err != nil {
//	    log.Fatalf("创建SAE客户端失败: %v", err)
//	}
//	// 使用client调用SAE API，如部署应用
func CreateClient(accessKeyid, accessKeySecret, endpoint string) (result *sae20190506.Client, err error) {
	c, err := credential.CreateCredentials(accessKeyid, accessKeySecret)
	if err != nil {
		return nil, err
	}
	config := &openapi.Config{
		Credential: c,
		Endpoint:   &endpoint,
	}
	r, err := sae20190506.NewClient(config)
	if err != nil {
		return nil, err
	}
	return r, nil
}
