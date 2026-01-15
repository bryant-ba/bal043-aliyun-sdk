package models

import "time"

// CloudResource 阿里云资源通用接口
// 定义了所有阿里云资源应该实现的基本方法
// CMDB系统通过此接口统一管理不同类型的云资源

type CloudResource interface {
	// GetResourceId 获取资源唯一标识符
	GetResourceId() string
	// GetResourceType 获取资源类型
	GetResourceType() string
	// GetResourceName 获取资源名称
	GetResourceName() string
	// GetRegionId 获取资源所在区域ID
	GetRegionId() string
}

// ECSInstance 阿里云ECS实例资源模型
// 封装了ECS实例的核心属性，方便CMDB系统存储和管理

type ECSInstance struct {
	InstanceId         string            // 实例ID
	InstanceName       string            // 实例名称
	RegionId           string            // 区域ID
	Status             string            // 实例状态
	InstanceType       string            // 实例规格
	PublicIp           []string          // 公网IP地址列表
	PrivateIp          []string          // 内网IP地址列表
	VpcId              string            // 所属VPC ID
	VSwitchId          string            // 所属交换机ID
	CreationTime       time.Time         // 创建时间
	ExpiredTime        time.Time         // 到期时间
	Tags               map[string]string // 标签
	CpuCoreCount       int               // CPU核心数
	MemorySize         int               // 内存大小（GB）
	OsType             string            // 操作系统类型
	OsName             string            // 操作系统名称
	ZoneId             string            // 可用区ID
	SecurityGroupIds   []string          // 安全组ID列表
	HostName           string            // 主机名
	ImageId            string            // 镜像ID
	InstanceChargeType string            // 计费方式
}

// GetResourceId 获取ECS实例ID
func (e *ECSInstance) GetResourceId() string {
	return e.InstanceId
}

// GetResourceType 获取资源类型
func (e *ECSInstance) GetResourceType() string {
	return "ECS"
}

// GetResourceName 获取ECS实例名称
func (e *ECSInstance) GetResourceName() string {
	return e.InstanceName
}

// GetRegionId 获取ECS实例所在区域ID
func (e *ECSInstance) GetRegionId() string {
	return e.RegionId
}

// VPC 阿里云VPC资源模型
// 封装了VPC的核心属性

type VPC struct {
	VpcId        string            // VPC ID
	VpcName      string            // VPC名称
	RegionId     string            // 区域ID
	Status       string            // VPC状态
	CidrBlock    string            // VPC网段
	Description  string            // 描述
	CreationTime time.Time         // 创建时间
	Tags         map[string]string // 标签
	IsDefault    bool              // 是否默认VPC
}

// GetResourceId 获取VPC ID
func (v *VPC) GetResourceId() string {
	return v.VpcId
}

// GetResourceType 获取资源类型
func (v *VPC) GetResourceType() string {
	return "VPC"
}

// GetResourceName 获取VPC名称
func (v *VPC) GetResourceName() string {
	return v.VpcName
}

// GetRegionId 获取VPC所在区域ID
func (v *VPC) GetRegionId() string {
	return v.RegionId
}

// VSwitch 阿里云交换机资源模型
// 封装了交换机的核心属性

type VSwitch struct {
	VSwitchId               string            // 交换机ID
	VSwitchName             string            // 交换机名称
	RegionId                string            // 区域ID
	ZoneId                  string            // 可用区ID
	VpcId                   string            // 所属VPC ID
	Status                  string            // 状态
	CidrBlock               string            // 网段
	CreationTime            time.Time         // 创建时间
	Tags                    map[string]string // 标签
	AvailableIpAddressCount int               // 可用IP数量
}

// GetResourceId 获取交换机ID
func (vs *VSwitch) GetResourceId() string {
	return vs.VSwitchId
}

// GetResourceType 获取资源类型
func (vs *VSwitch) GetResourceType() string {
	return "VSwitch"
}

// GetResourceName 获取交换机名称
func (vs *VSwitch) GetResourceName() string {
	return vs.VSwitchName
}

// GetRegionId 获取交换机所在区域ID
func (vs *VSwitch) GetRegionId() string {
	return vs.RegionId
}

// RDSInstance 阿里云RDS实例资源模型
// 封装了RDS实例的核心属性

type RDSInstance struct {
	InstanceId               string            // 实例ID
	InstanceName             string            // 实例名称
	RegionId                 string            // 区域ID
	Engine                   string            // 数据库引擎
	EngineVersion            string            // 数据库版本
	Status                   string            // 实例状态
	ConnectionString         string            // 连接地址
	Port                     int               // 端口
	CreationTime             time.Time         // 创建时间
	ExpiredTime              time.Time         // 到期时间
	Tags                     map[string]string // 标签
	DbInstanceClass          string            // 实例规格
	DbInstanceNetType        string            // 网络类型
	VpcId                    string            // 所属VPC ID
	VSwitchId                string            // 所属交换机ID
	ReadonlyConnectionString string            // 只读连接地址
	ReadOnlyDBInstanceIds    []string          // 只读实例ID列表
	InstanceChargeType       string            // 计费方式
}

// GetResourceId 获取RDS实例ID
func (r *RDSInstance) GetResourceId() string {
	return r.InstanceId
}

// GetResourceType 获取资源类型
func (r *RDSInstance) GetResourceType() string {
	return "RDS"
}

// GetResourceName 获取RDS实例名称
func (r *RDSInstance) GetResourceName() string {
	return r.InstanceName
}

// GetRegionId 获取RDS实例所在区域ID
func (r *RDSInstance) GetRegionId() string {
	return r.RegionId
}

// SLBInstance 阿里云传统型负载均衡实例模型
// 封装了SLB实例的核心属性

type SLBInstance struct {
	LoadBalancerId     string            // 负载均衡ID
	LoadBalancerName   string            // 负载均衡名称
	RegionId           string            // 区域ID
	Status             string            // 状态
	Address            string            // 公网IP地址
	VpcId              string            // 所属VPC ID
	VSwitchId          string            // 所属交换机ID
	NetworkType        string            // 网络类型
	LoadBalancerType   string            // 负载均衡类型
	CreationTime       time.Time         // 创建时间
	Tags               map[string]string // 标签
	ListenerPorts      []int             // 监听端口列表
	InstanceChargeType string            // 计费方式
}

// GetResourceId 获取SLB实例ID
func (s *SLBInstance) GetResourceId() string {
	return s.LoadBalancerId
}

// GetResourceType 获取资源类型
func (s *SLBInstance) GetResourceType() string {
	return "SLB"
}

// GetResourceName 获取SLB实例名称
func (s *SLBInstance) GetResourceName() string {
	return s.LoadBalancerName
}

// GetRegionId 获取SLB实例所在区域ID
func (s *SLBInstance) GetRegionId() string {
	return s.RegionId
}

// ALBInstance 阿里云应用型负载均衡实例模型
// 封装了ALB实例的核心属性

type ALBInstance struct {
	LoadBalancerId     string            // 负载均衡ID
	LoadBalancerName   string            // 负载均衡名称
	RegionId           string            // 区域ID
	Status             string            // 状态
	AddressType        string            // 地址类型
	VpcId              string            // 所属VPC ID
	ZoneIds            []string          // 可用区ID列表
	CreationTime       time.Time         // 创建时间
	Tags               map[string]string // 标签
	InstanceChargeType string            // 计费方式
}

// GetResourceId 获取ALB实例ID
func (a *ALBInstance) GetResourceId() string {
	return a.LoadBalancerId
}

// GetResourceType 获取资源类型
func (a *ALBInstance) GetResourceType() string {
	return "ALB"
}

// GetResourceName 获取ALB实例名称
func (a *ALBInstance) GetResourceName() string {
	return a.LoadBalancerName
}

// GetRegionId 获取ALB实例所在区域ID
func (a *ALBInstance) GetRegionId() string {
	return a.RegionId
}

// NLBInstance 阿里云网络型负载均衡实例模型
// 封装了NLB实例的核心属性

type NLBInstance struct {
	LoadBalancerId     string            // 负载均衡ID
	LoadBalancerName   string            // 负载均衡名称
	RegionId           string            // 区域ID
	Status             string            // 状态
	NetworkType        string            // 网络类型
	VpcId              string            // 所属VPC ID
	ZoneIds            []string          // 可用区ID列表
	CreationTime       time.Time         // 创建时间
	Tags               map[string]string // 标签
	InstanceChargeType string            // 计费方式
}

// GetResourceId 获取NLB实例ID
func (n *NLBInstance) GetResourceId() string {
	return n.LoadBalancerId
}

// GetResourceType 获取资源类型
func (n *NLBInstance) GetResourceType() string {
	return "NLB"
}

// GetResourceName 获取NLB实例名称
func (n *NLBInstance) GetResourceName() string {
	return n.LoadBalancerName
}

// GetRegionId 获取NLB实例所在区域ID
func (n *NLBInstance) GetRegionId() string {
	return n.RegionId
}

// SecurityGroup 阿里云安全组资源模型
// 封装了安全组的核心属性

type SecurityGroup struct {
	SecurityGroupId   string            // 安全组ID
	SecurityGroupName string            // 安全组名称
	RegionId          string            // 区域ID
	VpcId             string            // 所属VPC ID
	Description       string            // 描述
	CreationTime      time.Time         // 创建时间
	Tags              map[string]string // 标签
}

// GetResourceId 获取安全组ID
func (sg *SecurityGroup) GetResourceId() string {
	return sg.SecurityGroupId
}

// GetResourceType 获取资源类型
func (sg *SecurityGroup) GetResourceType() string {
	return "SecurityGroup"
}

// GetResourceName 获取安全组名称
func (sg *SecurityGroup) GetResourceName() string {
	return sg.SecurityGroupName
}

// GetRegionId 获取安全组所在区域ID
func (sg *SecurityGroup) GetRegionId() string {
	return sg.RegionId
}
