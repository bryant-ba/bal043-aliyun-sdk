package arms

import (
	arms20190808 "github.com/alibabacloud-go/arms-20190808/v11/client"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"

	"bal043-aliyun-sdk/aliyun/credential"
)

// CreateClient 创建阿里云应用实时监控服务客户端实例
// 用于初始化阿里云应用实时监控服务（ARMS）的API客户端
//
// 参数说明:
//   - accessKeyid: 阿里云访问密钥ID（RAM用户的AccessKey ID）
//   - accessKeySecret: 阿里云访问密钥（RAM用户的AccessKey Secret）
//   - endpoint: 阿里云ARMS服务的访问端点（如 "arms.cn-hangzhou.aliyuncs.com"）
//
// 返回值说明:
//   - result: ARMS客户端实例指针，用于调用应用实时监控API
//   - err: 客户端创建失败时返回错误信息，成功时返回nil
//
// 使用示例:
//
//	client, err := arms.CreateClient("your-access-key-id", "your-access-key-secret", "arms.cn-hangzhou.aliyuncs.com")
//	if err != nil {
//	    log.Fatalf("创建ARMS客户端失败: %v", err)
//	}
//	// 使用client调用ARMS API，如查询应用监控数据
func CreateClient(accessKeyid, accessKeySecret, endpoint string) (result *arms20190808.Client, err error) {
	c, err := credential.CreateCredentials(accessKeyid, accessKeySecret)
	if err != nil {
		return nil, err
	}
	config := &openapi.Config{
		Credential: c,
		Endpoint:   &endpoint,
	}
	r, err := arms20190808.NewClient(config)
	if err != nil {
		return nil, err
	}
	return r, nil
}
