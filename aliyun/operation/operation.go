package operation

import (
	"context"

	"github.com/bryant-ba/bal043-aliyun-sdk/aliyun/models"
)

// ResourceOperation 资源操作统一接口
// 定义了所有资源操作应该实现的基本方法
// 包括资源的查询、过滤、标签管理等

type ResourceOperation interface {
	// ListResources 查询资源列表
	// ctx: 上下文，用于超时控制和取消操作
	// regionId: 区域ID
	// 返回值: 资源列表和错误信息
	ListResources(ctx context.Context, regionId string) ([]models.CloudResource, error)

	// GetResourceById 根据ID查询单个资源
	// ctx: 上下文
	// regionId: 区域ID
	// resourceId: 资源ID
	// 返回值: 单个资源和错误信息
	GetResourceById(ctx context.Context, regionId, resourceId string) (models.CloudResource, error)

	// ListResourcesByTag 根据标签查询资源
	// ctx: 上下文
	// regionId: 区域ID
	// tagKey: 标签键
	// tagValue: 标签值
	// 返回值: 资源列表和错误信息
	ListResourcesByTag(ctx context.Context, regionId, tagKey, tagValue string) ([]models.CloudResource, error)

	// ListResourcesByTags 根据多标签查询资源
	// ctx: 上下文
	// regionId: 区域ID
	// tags: 标签键值对
	// 返回值: 资源列表和错误信息
	ListResourcesByTags(ctx context.Context, regionId string, tags map[string]string) ([]models.CloudResource, error)

	// GetResourceTags 获取资源标签
	// ctx: 上下文
	// regionId: 区域ID
	// resourceId: 资源ID
	// 返回值: 标签映射和错误信息
	GetResourceTags(ctx context.Context, regionId, resourceId string) (map[string]string, error)

	// TagResource 为资源添加标签
	// ctx: 上下文
	// regionId: 区域ID
	// resourceId: 资源ID
	// tags: 要添加的标签
	// 返回值: 错误信息
	TagResource(ctx context.Context, regionId, resourceId string, tags map[string]string) error

	// UntagResource 为资源删除标签
	// ctx: 上下文
	// regionId: 区域ID
	// resourceId: 资源ID
	// tagKeys: 要删除的标签键列表
	// 返回值: 错误信息
	UntagResource(ctx context.Context, regionId, resourceId string, tagKeys []string) error
}

// OperationFactory 资源操作工厂
// 根据资源类型创建对应的资源操作实例

type OperationFactory struct {
	// 可以添加一些工厂配置，如客户端创建器等
}

// NewOperationFactory 创建资源操作工厂
func NewOperationFactory() *OperationFactory {
	return &OperationFactory{}
}

// ResourceOperationOption 资源操作选项
// 用于配置资源操作实例的创建

type ResourceOperationOption struct {
	AccessKeyId     string
	AccessKeySecret string
	RegionId        string
	Endpoint        string
}

// CommonParams 公共参数
// 定义了阿里云API请求的公共参数

type CommonParams struct {
	AccessKeyId     string
	AccessKeySecret string
	RegionId        string
	Endpoint        string
}

// PaginationParams 分页参数
// 用于控制资源查询的分页

type PaginationParams struct {
	PageNumber int // 页码，从1开始
	PageSize   int // 每页大小，最大100
}

// FilterParams 过滤参数
// 用于资源查询时的条件过滤

type FilterParams struct {
	Name             string            // 资源名称过滤
	Status           []string          // 状态过滤，支持多个状态
	InstanceType     []string          // 实例类型过滤，支持多个类型
	VpcId            string            // VPC ID过滤
	ZoneId           string            // 可用区ID过滤
	CreationTimeFrom string            // 创建时间起始，格式：yyyy-MM-ddTHH:mm:ssZ
	CreationTimeTo   string            // 创建时间结束，格式：yyyy-MM-ddTHH:mm:ssZ
	ExtraFilters     map[string]string // 其他过滤条件，根据资源类型而定
}

// ListRequest 资源列表查询请求
// 包含了资源查询的所有参数

type ListRequest struct {
	CommonParams
	FilterParams
	PaginationParams
	TagFilters map[string]string // 标签过滤
}

// ListResponse 资源列表查询响应
// 包含了资源列表和分页信息

type ListResponse struct {
	Resources  []models.CloudResource // 资源列表
	TotalCount int                    // 总数量
	PageNumber int                    // 当前页码
	PageSize   int                    // 每页大小
	NextPage   bool                   // 是否有下一页
}

// TagRequest 标签操作请求
// 用于资源标签的添加和删除

type TagRequest struct {
	CommonParams
	ResourceId string            // 资源ID
	Tags       map[string]string // 标签键值对
}

// TagResponse 标签操作响应
// 标签操作的结果

type TagResponse struct {
	Success   bool   // 操作是否成功
	RequestId string // 请求ID
}

// ResourceInfo 资源信息结构体
// 包含资源的基本信息和扩展信息

type ResourceInfo struct {
	Resource       models.CloudResource   // 资源基本信息
	Tags           map[string]string      // 资源标签
	ExtraInfo      map[string]interface{} // 扩展信息，根据资源类型而定
	LastUpdateTime string                 // 最后更新时间
}

// ResourceManager 资源管理器
// 用于管理不同类型的资源操作实例
// 可以根据资源类型获取对应的操作实例

type ResourceManager struct {
	operations map[string]ResourceOperation
	factory    *OperationFactory
}

// NewResourceManager 创建资源管理器
func NewResourceManager() *ResourceManager {
	return &ResourceManager{
		operations: make(map[string]ResourceOperation),
		factory:    NewOperationFactory(),
	}
}

// RegisterOperation 注册资源操作实例
// resourceType: 资源类型
// operation: 资源操作实例
func (rm *ResourceManager) RegisterOperation(resourceType string, operation ResourceOperation) {
	rm.operations[resourceType] = operation
}

// GetOperation 获取资源操作实例
// resourceType: 资源类型
// 返回值: 资源操作实例和是否存在
func (rm *ResourceManager) GetOperation(resourceType string) (ResourceOperation, bool) {
	operation, ok := rm.operations[resourceType]
	return operation, ok
}

// ListResourceTypes 获取所有支持的资源类型
// 返回值: 支持的资源类型列表
func (rm *ResourceManager) ListResourceTypes() []string {
	var resourceTypes []string
	for resourceType := range rm.operations {
		resourceTypes = append(resourceTypes, resourceType)
	}
	return resourceTypes
}

// IsResourceTypeSupported 检查资源类型是否支持
// resourceType: 资源类型
// 返回值: 是否支持
func (rm *ResourceManager) IsResourceTypeSupported(resourceType string) bool {
	_, ok := rm.operations[resourceType]
	return ok
}
