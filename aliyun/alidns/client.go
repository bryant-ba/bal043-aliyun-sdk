package alidns

import (
	alidns20150109 "github.com/alibabacloud-go/alidns-20150109/v4/client"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"

	"github.com/bryant-ba/bal043-aliyun-sdk/aliyun/credential"
)

// CreateClient 创建阿里云云解析DNS客户端实例
// 用于初始化阿里云云解析服务（Aliyun DNS）的API客户端
//
// 参数说明:
//   - accessKeyid: 阿里云访问密钥ID（RAM用户的AccessKey ID）
//   - accessKeySecret: 阿里云访问密钥（RAM用户的AccessKey Secret）
//   - endpoint: 阿里云Alidns服务的访问端点（如 "alidns.cn-hangzhou.aliyuncs.com"）
//
// 返回值说明:
//   - result: Alidns客户端实例指针，用于调用云解析DNS API
//   - err: 客户端创建失败时返回错误信息，成功时返回nil
//
// 使用示例:
//
//	client, err := alidns.CreateClient("your-access-key-id", "your-access-key-secret", "alidns.cn-hangzhou.aliyuncs.com")
//	if err != nil {
//	    log.Fatalf("创建Alidns客户端失败: %v", err)
//	}
//	// 使用client调用云解析DNS API，如查询域名解析记录

func CreateClient(accessKeyid, accessKeySecret, endpoint string) (result *alidns20150109.Client, err error) {
	c, err := credential.CreateCredentials(accessKeyid, accessKeySecret)
	if err != nil {
		return nil, err
	}
	config := &openapi.Config{
		Credential: c,
		Endpoint:   &endpoint,
	}
	r, err := alidns20150109.NewClient(config)
	if err != nil {
		return nil, err
	}
	return r, nil
}
