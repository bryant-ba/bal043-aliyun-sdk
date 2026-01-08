package alb

import (
	alb "github.com/alibabacloud-go/alb-20200616/v2/client"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"

	"bal043-aliyun-sdk/aliyun/credential"
)

// CreateClient 创建阿里云应用型负载均衡客户端实例
// 用于初始化阿里云应用型负载均衡（ALB）的API客户端
//
// 参数说明:
//   - accessKeyid: 阿里云访问密钥ID（RAM用户的AccessKey ID）
//   - accessKeySecret: 阿里云访问密钥（RAM用户的AccessKey Secret）
//   - endpoint: 阿里云ALB服务的访问端点（如 "alb.cn-hangzhou.aliyuncs.com"）
//
// 返回值说明:
//   - result: ALB客户端实例指针，用于调用应用型负载均衡API
//   - err: 客户端创建失败时返回错误信息，成功时返回nil
//
// 使用示例:
//
//	client, err := alb.CreateClient("your-access-key-id", "your-access-key-secret", "alb.cn-hangzhou.aliyuncs.com")
//	if err != nil {
//	    log.Fatalf("创建ALB客户端失败: %v", err)
//	}
//	// 使用client调用ALB API，如创建ALB实例、配置转发规则等
func CreateClient(accessKeyid, accessKeySecret, endpoint string) (result *alb.Client, err error) {
	c, err := credential.CreateCredentials(accessKeyid, accessKeySecret)
	if err != nil {
		return nil, err
	}
	config := &openapi.Config{
		Credential: c,
		Endpoint:   &endpoint,
	}
	r, err := alb.NewClient(config)
	if err != nil {
		return nil, err
	}
	return r, nil
}
