package config

import (
	"io/ioutil"
	"log"
	"path/filepath"

	"gopkg.in/ini.v1"
	"gopkg.in/yaml.v3"
)

// Config 阿里云SDK配置结构
// 支持多账号、多区域配置，方便CMDB系统管理不同阿里云账号下的资源

type Config struct {
	DefaultRegion string          `yaml:"default_region"` // 默认区域ID
	Accounts      []AccountConfig `yaml:"accounts"`       // 阿里云账号列表
	Cache         CacheConfig     `yaml:"cache"`          // 缓存配置
	ResourceTypes []string        `yaml:"resource_types"` // 关注的资源类型列表
}

// AccountConfig 阿里云账号配置
// 存储单个阿里云账号的认证信息和关注的区域

type AccountConfig struct {
	Name            string   `yaml:"name"`              // 账号名称，用于区分不同账号
	AccessKeyId     string   `yaml:"access_key_id"`     // AccessKey ID
	AccessKeySecret string   `yaml:"access_key_secret"` // AccessKey Secret
	Regions         []string `yaml:"regions"`           // 关注的区域列表
	Enabled         bool     `yaml:"enabled"`           // 账号是否启用
}

// CacheConfig 缓存配置
// 控制资源数据的缓存策略

type CacheConfig struct {
	Enabled       bool   `yaml:"enabled"`        // 是否启用缓存
	ExpireTime    int    `yaml:"expire_time"`    // 缓存过期时间（秒）
	StorageType   string `yaml:"storage_type"`   // 存储类型：memory/redis
	RedisAddress  string `yaml:"redis_address"`  // Redis地址（当storage_type为redis时）
	RedisPassword string `yaml:"redis_password"` // Redis密码
	RedisDB       int    `yaml:"redis_db"`       // Redis数据库
}

// ConfigManager 配置管理器
// 负责配置的加载、管理和访问

type ConfigManager struct {
	config *Config
}

// NewConfigManager 创建配置管理器
// 初始化并加载配置文件
func NewConfigManager(configPath string) (*ConfigManager, error) {
	// 支持YAML和INI两种配置文件格式
	ext := filepath.Ext(configPath)

	var config Config
	var err error

	if ext == ".yaml" || ext == ".yml" {
		config, err = loadYAMLConfig(configPath)
	} else if ext == ".ini" {
		config, err = loadINIConfig(configPath)
	} else {
		config, err = loadYAMLConfig(configPath) // 默认尝试YAML
	}

	if err != nil {
		return nil, err
	}

	return &ConfigManager{
		config: &config,
	}, nil
}

// loadYAMLConfig 加载YAML格式配置文件
func loadYAMLConfig(configPath string) (Config, error) {
	var config Config

	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Printf("读取配置文件失败: %v", err)
		return config, err
	}

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Printf("解析YAML配置失败: %v", err)
		return config, err
	}

	return config, nil
}

// loadINIConfig 加载INI格式配置文件
// 支持传统INI格式配置，方便老系统迁移
func loadINIConfig(configPath string) (Config, error) {
	var config Config

	cfg, err := ini.Load(configPath)
	if err != nil {
		log.Printf("读取INI配置文件失败: %v", err)
		return config, err
	}

	// 加载默认区域
	config.DefaultRegion = cfg.Section("default").Key("region").String()

	// 加载资源类型
	resourceTypes := cfg.Section("resource").Key("types").Strings(",")
	config.ResourceTypes = resourceTypes

	// 加载缓存配置
	cacheSection := cfg.Section("cache")
	storageType := cacheSection.Key("storage_type").String()
	if storageType == "" {
		storageType = "memory"
	}
	config.Cache = CacheConfig{
		Enabled:       cacheSection.Key("enabled").MustBool(false),
		ExpireTime:    cacheSection.Key("expire_time").MustInt(3600),
		StorageType:   storageType,
		RedisAddress:  cacheSection.Key("redis_address").String(),
		RedisPassword: cacheSection.Key("redis_password").String(),
		RedisDB:       cacheSection.Key("redis_db").MustInt(0),
	}

	// 加载账号配置（支持多个账号）
	for _, section := range cfg.Sections() {
		if section.Name() == "default" || section.Name() == "cache" || section.Name() == "resource" {
			continue
		}

		// 每个非默认section代表一个账号
		account := AccountConfig{
			Name:            section.Name(),
			AccessKeyId:     section.Key("access_key_id").String(),
			AccessKeySecret: section.Key("access_key_secret").String(),
			Regions:         section.Key("regions").Strings(","),
			Enabled:         section.Key("enabled").MustBool(true),
		}

		config.Accounts = append(config.Accounts, account)
	}

	return config, nil
}

// GetConfig 获取完整配置
func (m *ConfigManager) GetConfig() *Config {
	return m.config
}

// GetAccount 获取指定名称的账号配置
func (m *ConfigManager) GetAccount(accountName string) (*AccountConfig, error) {
	for _, account := range m.config.Accounts {
		if account.Name == accountName {
			return &account, nil
		}
	}
	return nil, nil // 找不到返回nil，不报错
}

// GetDefaultAccount 获取默认账号配置（第一个启用的账号）
func (m *ConfigManager) GetDefaultAccount() (*AccountConfig, error) {
	for _, account := range m.config.Accounts {
		if account.Enabled {
			return &account, nil
		}
	}
	return nil, nil // 没有启用的账号返回nil
}

// GetEnabledAccounts 获取所有启用的账号
func (m *ConfigManager) GetEnabledAccounts() []AccountConfig {
	var enabledAccounts []AccountConfig
	for _, account := range m.config.Accounts {
		if account.Enabled {
			enabledAccounts = append(enabledAccounts, account)
		}
	}
	return enabledAccounts
}

// GetDefaultRegion 获取默认区域
func (m *ConfigManager) GetDefaultRegion() string {
	return m.config.DefaultRegion
}

// GetResourceTypes 获取关注的资源类型列表
func (m *ConfigManager) GetResourceTypes() []string {
	return m.config.ResourceTypes
}

// IsResourceTypeEnabled 检查资源类型是否启用
func (m *ConfigManager) IsResourceTypeEnabled(resourceType string) bool {
	if len(m.config.ResourceTypes) == 0 {
		return true // 默认全部启用
	}

	for _, rt := range m.config.ResourceTypes {
		if rt == resourceType {
			return true
		}
	}

	return false
}

// GetCacheConfig 获取缓存配置
func (m *ConfigManager) GetCacheConfig() CacheConfig {
	return m.config.Cache
}

// IsCacheEnabled 检查是否启用缓存
func (m *ConfigManager) IsCacheEnabled() bool {
	return m.config.Cache.Enabled
}

// ExampleYAMLConfig 返回YAML配置示例
// 用于生成配置文件模板
func ExampleYAMLConfig() string {
	return `# 阿里云SDK配置文件
# 默认区域ID
default_region: cn-hangzhou

# 关注的资源类型列表
resource_types:
  - ECS
  - RDS
  - VPC
  - SLB
  - ALB
  - NLB
  - SecurityGroup

# 阿里云账号列表
accounts:
  - name: prod
    access_key_id: your-prod-access-key-id
    access_key_secret: your-prod-access-key-secret
    regions:
      - cn-hangzhou
      - cn-shanghai
      - cn-beijing
    enabled: true

  - name: dev
    access_key_id: your-dev-access-key-id
    access_key_secret: your-dev-access-key-secret
    regions:
      - cn-hangzhou
    enabled: true

# 缓存配置
cache:
  enabled: true
  expire_time: 3600
  storage_type: memory
  # redis配置（当storage_type为redis时）
  redis_address: localhost:6379
  redis_password: ""
  redis_db: 0
`
}

// ExampleINIConfig 返回INI配置示例
// 用于生成配置文件模板
func ExampleINIConfig() string {
	return `[default]
region = cn-hangzhou

[resource]
types = ECS,RDS,VPC,SLB,ALB,NLB,SecurityGroup

[cache]
enabled = true
expire_time = 3600
storage_type = memory

[prod]
access_key_id = your-prod-access-key-id
access_key_secret = your-prod-access-key-secret
regions = cn-hangzhou,cn-shanghai,cn-beijing
enabled = true

[dev]
access_key_id = your-dev-access-key-id
access_key_secret = your-dev-access-key-secret
regions = cn-hangzhou
enabled = true
`
}
