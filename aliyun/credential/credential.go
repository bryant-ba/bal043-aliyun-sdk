package credential

import (
	"github.com/aliyun/credentials-go/credentials"
)

// CreateCredentials 创建阿里云访问凭证
// 用于生成阿里云SDK所需的认证凭证，支持AccessKey认证方式
//
// 参数说明:
//   - accessKeyid: 阿里云访问密钥ID（RAM用户的AccessKey ID）
//   - accessKeySecret: 阿里云访问密钥（RAM用户的AccessKey Secret）
//
// 返回值说明:
//   - credential: 阿里云凭证对象，用于SDK认证
//   - err: 凭证创建失败时返回错误信息，成功时返回nil
//
// 使用示例:
//
//	credential, err := credential.CreateCredentials("your-access-key-id", "your-access-key-secret")
//	if err != nil {
//	    log.Fatalf("创建凭证失败: %v", err)
//	}
func CreateCredentials(accessKeyid, accessKeySecret string) (credential credentials.Credential, err error) {
	AKconfig := new(credentials.Config).SetType("access_key").SetAccessKeyId(accessKeyid).SetAccessKeySecret(accessKeySecret)
	akCredential, err := credentials.NewCredential(AKconfig)
	if err != nil {
		return nil, err
	}
	return akCredential, nil
}
