package vpc

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	vpc20160428 "github.com/alibabacloud-go/vpc-20160428/v6/client"

	"github.com/bryant-ba/bal043-aliyun-sdk/aliyun/credential"
)

// CreateClient 创建阿里云专有网络客户端实例
// 用于初始化阿里云专有网络（VPC）的API客户端
//
// 参数说明:
//   - accessKeyid: 阿里云访问密钥ID（RAM用户的AccessKey ID）
//   - accessKeySecret: 阿里云访问密钥（RAM用户的AccessKey Secret）
//   - endpoint: 阿里云VPC服务的访问端点（如 "vpc.cn-hangzhou.aliyuncs.com"）
//
// 返回值说明:
//   - result: VPC客户端实例指针，用于调用专有网络API
//   - err: 客户端创建失败时返回错误信息，成功时返回nil
//
// 使用示例:
//
//	client, err := vpc.CreateClient("your-access-key-id", "your-access-key-secret", "vpc.cn-hangzhou.aliyuncs.com")
//	if err != nil {
//	    log.Fatalf("创建VPC客户端失败: %v", err)
//	}
//	// 使用client调用VPC API，如查询VPC列表、创建交换机等
func CreateClient(accessKeyid, accessKeySecret, endpoint string) (result *vpc20160428.Client, err error) {
	c, err := credential.CreateCredentials(accessKeyid, accessKeySecret)
	if err != nil {
		return nil, err
	}
	config := &openapi.Config{
		Credential: c,
		Endpoint:   &endpoint,
	}
	r, err := vpc20160428.NewClient(config)
	if err != nil {
		return nil, err
	}
	return r, nil
}
