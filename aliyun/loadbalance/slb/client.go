package slb

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	slb20140515 "github.com/alibabacloud-go/slb-20140515/v4/client"

	"github.com/bryant-ba/bal043-aliyun-sdk/aliyun/credential"
)

// CreateClient 创建阿里云传统型负载均衡客户端实例
// 用于初始化阿里云传统型负载均衡（SLB）的API客户端
//
// 参数说明:
//   - accessKeyid: 阿里云访问密钥ID（RAM用户的AccessKey ID）
//   - accessKeySecret: 阿里云访问密钥（RAM用户的AccessKey Secret）
//   - endpoint: 阿里云SLB服务的访问端点（如 "slb.cn-hangzhou.aliyuncs.com"）
//
// 返回值说明:
//   - result: SLB客户端实例指针，用于调用传统型负载均衡API
//   - err: 客户端创建失败时返回错误信息，成功时返回nil
//
// 使用示例:
//
//	client, err := slb.CreateClient("your-access-key-id", "your-access-key-secret", "slb.cn-hangzhou.aliyuncs.com")
//	if err != nil {
//	    log.Fatalf("创建SLB客户端失败: %v", err)
//	}
//	// 使用client调用SLB API，如创建负载均衡实例、添加监听等
func CreateClient(accessKeyid, accessKeySecret, endpoint string) (result *slb20140515.Client, err error) {
	c, err := credential.CreateCredentials(accessKeyid, accessKeySecret)
	if err != nil {
		return nil, err
	}
	config := &openapi.Config{
		Credential: c,
		Endpoint:   &endpoint,
	}
	r, err := slb20140515.NewClient(config)
	if err != nil {
		return nil, err
	}
	return r, nil
}
