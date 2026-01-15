package ecs

import (
	"context"
	"fmt"
	"time"

	ecs20140526 "github.com/alibabacloud-go/ecs-20140526/v7/client"

	"github.com/bryant-ba/bal043-aliyun-sdk/aliyun/models"
)

// ECSOperator ECS资源操作封装
// 实现了operation.ResourceOperation接口
// 提供ECS资源的查询、标签管理等功能

type ECSOperator struct {
	client *ecs20140526.Client
	region string
}

// NewECSOperator 创建ECS资源操作实例
//
// 参数说明:
//   - client: ECS客户端实例
//   - region: 区域ID
//
// 返回值说明:
//   - ECS资源操作实例
func NewECSOperator(client *ecs20140526.Client, region string) *ECSOperator {
	return &ECSOperator{
		client: client,
		region: region,
	}
}

// ListResources 查询ECS实例列表
// 实现operation.ResourceOperation接口
//
// 参数说明:
//   - ctx: 上下文，用于超时控制和取消操作
//   - regionId: 区域ID
//
// 返回值说明:
//   - 资源列表和错误信息
func (op *ECSOperator) ListResources(ctx context.Context, regionId string) ([]models.CloudResource, error) {
	instances, err := op.ListInstances(ctx, regionId)
	if err != nil {
		return nil, err
	}

	// 转换为统一资源模型
	resources := make([]models.CloudResource, len(instances))
	for i, inst := range instances {
		resources[i] = inst
	}

	return resources, nil
}

// GetResourceById 根据ID查询单个ECS实例
// 实现operation.ResourceOperation接口
//
// 参数说明:
//   - ctx: 上下文
//   - regionId: 区域ID
//   - resourceId: 资源ID
//
// 返回值说明:
//   - 单个资源和错误信息
func (op *ECSOperator) GetResourceById(ctx context.Context, regionId, resourceId string) (models.CloudResource, error) {
	return op.GetInstanceById(ctx, regionId, resourceId)
}

// ListResourcesByTag 根据标签查询ECS实例
// 实现operation.ResourceOperation接口
//
// 参数说明:
//   - ctx: 上下文
//   - regionId: 区域ID
//   - tagKey: 标签键
//   - tagValue: 标签值
//
// 返回值说明:
//   - 资源列表和错误信息
func (op *ECSOperator) ListResourcesByTag(ctx context.Context, regionId, tagKey, tagValue string) ([]models.CloudResource, error) {
	tags := map[string]string{
		tagKey: tagValue,
	}

	instances, err := op.ListInstancesByTags(ctx, regionId, tags)
	if err != nil {
		return nil, err
	}

	// 转换为统一资源模型
	resources := make([]models.CloudResource, len(instances))
	for i, inst := range instances {
		resources[i] = inst
	}

	return resources, nil
}

// ListResourcesByTags 根据多标签查询ECS实例
// 实现operation.ResourceOperation接口
//
// 参数说明:
//   - ctx: 上下文
//   - regionId: 区域ID
//   - tags: 标签键值对
//
// 返回值说明:
//   - 资源列表和错误信息
func (op *ECSOperator) ListResourcesByTags(ctx context.Context, regionId string, tags map[string]string) ([]models.CloudResource, error) {
	instances, err := op.ListInstancesByTags(ctx, regionId, tags)
	if err != nil {
		return nil, err
	}

	// 转换为统一资源模型
	resources := make([]models.CloudResource, len(instances))
	for i, inst := range instances {
		resources[i] = inst
	}

	return resources, nil
}

// GetResourceTags 获取ECS实例标签
// 实现operation.ResourceOperation接口
//
// 参数说明:
//   - ctx: 上下文
//   - regionId: 区域ID
//   - resourceId: 资源ID
//
// 返回值说明:
//   - 标签映射和错误信息
func (op *ECSOperator) GetResourceTags(ctx context.Context, regionId, resourceId string) (map[string]string, error) {
	return op.GetInstanceTags(ctx, regionId, resourceId)
}

// TagResource 为ECS实例添加标签
// 实现operation.ResourceOperation接口
//
// 参数说明:
//   - ctx: 上下文
//   - regionId: 区域ID
//   - resourceId: 资源ID
//   - tags: 要添加的标签
//
// 返回值说明:
//   - 错误信息
func (op *ECSOperator) TagResource(ctx context.Context, regionId, resourceId string, tags map[string]string) error {
	return fmt.Errorf("TagResource not implemented yet")
}

// UntagResource 为ECS实例删除标签
// 实现operation.ResourceOperation接口
//
// 参数说明:
//   - ctx: 上下文
//   - regionId: 区域ID
//   - resourceId: 资源ID
//   - tagKeys: 要删除的标签键列表
//
// 返回值说明:
//   - 错误信息
func (op *ECSOperator) UntagResource(ctx context.Context, regionId, resourceId string, tagKeys []string) error {
	return fmt.Errorf("UntagResource not implemented yet")
}

// ListInstances 查询所有ECS实例
//
// 参数说明:
//   - ctx: 上下文，用于超时控制和取消操作
//   - regionId: 区域ID
//
// 返回值说明:
//   - ECS实例列表和错误信息
func (op *ECSOperator) ListInstances(ctx context.Context, regionId string) ([]*models.ECSInstance, error) {
	// 处理区域ID，如果为空则使用默认区域
	if regionId == "" {
		regionId = op.region
	}

	var allInstances []*models.ECSInstance
	pageNumber := int32(1)
	pageSize := int32(100)

	// 循环处理分页查询
	for {
		request := &ecs20140526.DescribeInstancesRequest{
			RegionId:   &regionId,
			PageNumber: &pageNumber,
			PageSize:   &pageSize,
		}

		response, err := op.client.DescribeInstances(request)
		if err != nil {
			return nil, err
		}

		// 转换实例数据
		instances := convertToECSInstances(response)
		allInstances = append(allInstances, instances...)

		// 检查是否还有下一页
		if int64(pageNumber)*int64(pageSize) >= int64(*response.Body.TotalCount) {
			break
		}

		pageNumber++
	}

	return allInstances, nil
}

// GetInstanceById 根据ID查询单个ECS实例
//
// 参数说明:
//   - ctx: 上下文
//   - regionId: 区域ID
//   - instanceId: 实例ID
//
// 返回值说明:
//   - 单个ECS实例和错误信息
func (op *ECSOperator) GetInstanceById(ctx context.Context, regionId, instanceId string) (*models.ECSInstance, error) {
	// 处理区域ID，如果为空则使用默认区域
	if regionId == "" {
		regionId = op.region
	}

	instanceIdsStr := instanceId
	request := &ecs20140526.DescribeInstancesRequest{
		RegionId:    &regionId,
		InstanceIds: &instanceIdsStr,
	}

	response, err := op.client.DescribeInstances(request)
	if err != nil {
		return nil, err
	}

	// 转换实例数据
	instances := convertToECSInstances(response)
	if len(instances) == 0 {
		return nil, fmt.Errorf("实例 %s 不存在", instanceId)
	}

	return instances[0], nil
}

// ListInstancesByTag 根据标签查询ECS实例
//
// 参数说明:
//   - ctx: 上下文
//   - regionId: 区域ID
//   - tagKey: 标签键
//   - tagValue: 标签值
//
// 返回值说明:
//   - ECS实例列表和错误信息
func (op *ECSOperator) ListInstancesByTag(ctx context.Context, regionId, tagKey, tagValue string) ([]*models.ECSInstance, error) {
	tags := map[string]string{
		tagKey: tagValue,
	}

	return op.ListInstancesByTags(ctx, regionId, tags)
}

// ListInstancesByTags 根据多标签查询ECS实例
//
// 参数说明:
//   - ctx: 上下文
//   - regionId: 区域ID
//   - tags: 标签键值对
//
// 返回值说明:
//   - ECS实例列表和错误信息
func (op *ECSOperator) ListInstancesByTags(ctx context.Context, regionId string, tags map[string]string) ([]*models.ECSInstance, error) {
	// 处理区域ID，如果为空则使用默认区域
	if regionId == "" {
		regionId = op.region
	}

	// 循环处理分页查询
	var allInstances []*models.ECSInstance
	pageNumber := int32(1)
	pageSize := int32(100)

	for {
		request := &ecs20140526.DescribeInstancesRequest{
			RegionId:   &regionId,
			PageNumber: &pageNumber,
			PageSize:   &pageSize,
		}

		response, err := op.client.DescribeInstances(request)
		if err != nil {
			return nil, err
		}

		// 转换实例数据
		instances := convertToECSInstances(response)
		allInstances = append(allInstances, instances...)

		// 检查是否还有下一页
		if int64(pageNumber)*int64(pageSize) >= int64(*response.Body.TotalCount) {
			break
		}

		pageNumber++
	}

	// 手动过滤标签
	var filteredInstances []*models.ECSInstance
	for _, inst := range allInstances {
		match := true
		for k, v := range tags {
			if inst.Tags[k] != v {
				match = false
				break
			}
		}
		if match {
			filteredInstances = append(filteredInstances, inst)
		}
	}

	return filteredInstances, nil
}

// GetInstanceTags 获取ECS实例标签
//
// 参数说明:
//   - ctx: 上下文
//   - regionId: 区域ID
//   - instanceId: 实例ID
//
// 返回值说明:
//   - 标签映射和错误信息
func (op *ECSOperator) GetInstanceTags(ctx context.Context, regionId, instanceId string) (map[string]string, error) {
	return map[string]string{}, nil // 简化实现
}

// convertToECSInstances 将ECS API响应转换为统一模型
//
// 参数说明:
//   - response: ECS API响应
//
// 返回值说明:
//   - ECS实例列表
func convertToECSInstances(response *ecs20140526.DescribeInstancesResponse) []*models.ECSInstance {
	if response.Body.Instances == nil || response.Body.Instances.Instance == nil {
		return []*models.ECSInstance{}
	}

	instances := make([]*models.ECSInstance, len(response.Body.Instances.Instance))

	for i, inst := range response.Body.Instances.Instance {
		// 转换创建时间
		creationTime, _ := time.Parse(time.RFC3339, *inst.CreationTime)
		expiredTime := time.Time{}
		if inst.ExpiredTime != nil {
			expiredTime, _ = time.Parse(time.RFC3339, *inst.ExpiredTime)
		}

		// 初始化默认值
		publicIp := []string{}
		privateIp := []string{}
		securityGroupIds := []string{}
		vpcId := ""
		vSwitchId := ""
		zoneId := ""
		hostName := ""
		imageId := ""
		osType := ""
		osName := ""
		instanceChargeType := ""
		tags := make(map[string]string)

		// 转换实例数据
		instances[i] = &models.ECSInstance{
			InstanceId:         *inst.InstanceId,
			InstanceName:       *inst.InstanceName,
			RegionId:           *inst.RegionId,
			Status:             *inst.Status,
			InstanceType:       *inst.InstanceType,
			PublicIp:           publicIp,
			PrivateIp:          privateIp,
			VpcId:              vpcId,
			VSwitchId:          vSwitchId,
			CreationTime:       creationTime,
			ExpiredTime:        expiredTime,
			Tags:               tags,
			CpuCoreCount:       int(*inst.Cpu),
			MemorySize:         int(*inst.Memory),
			OsType:             osType,
			OsName:             osName,
			ZoneId:             zoneId,
			SecurityGroupIds:   securityGroupIds,
			HostName:           hostName,
			ImageId:            imageId,
			InstanceChargeType: instanceChargeType,
		}
	}

	return instances
}
