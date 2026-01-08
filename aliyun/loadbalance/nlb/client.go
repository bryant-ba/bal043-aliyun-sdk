package nlb

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	nlb "github.com/alibabacloud-go/nlb-20220430/v4/client"

	"github.com/bryant-ba/bal043-aliyun-sdk/aliyun/credential"
)

// CreateClient 创建阿里云网络型负载均衡客户端实例
// 用于初始化阿里云网络型负载均衡（NLB）的API客户端
//
// 参数说明:
//   - accessKeyid: 阿里云访问密钥ID（RAM用户的AccessKey ID）
//   - accessKeySecret: 阿里云访问密钥（RAM用户的AccessKey Secret）
//   - endpoint: 阿里云NLB服务的访问端点（如 "nlb.cn-hangzhou.aliyuncs.com"）
//
// 返回值说明:
//   - result: NLB客户端实例指针，用于调用网络型负载均衡API
//   - err: 客户端创建失败时返回错误信息，成功时返回nil
//
// 使用示例:
//
//	client, err := nlb.CreateClient("your-access-key-id", "your-access-key-secret", "nlb.cn-hangzhou.aliyuncs.com")
//	if err != nil {
//	    log.Fatalf("创建NLB客户端失败: %v", err)
//	}
//	// 使用client调用NLB API，如创建NLB实例、配置监听和服务器组等
func CreateClient(accessKeyid, accessKeySecret, endpoint string) (result *nlb.Client, err error) {
	c, err := credential.CreateCredentials(accessKeyid, accessKeySecret)
	if err != nil {
		return nil, err
	}
	config := &openapi.Config{
		Credential: c,
		Endpoint:   &endpoint,
	}
	r, err := nlb.NewClient(config)
	if err != nil {
		return nil, err
	}
	return r, nil
}
