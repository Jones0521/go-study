package host

import (
	"time"

	"github.com/go-playground/validator/v10"
)

var (
	validate = validator.New()
)

func NewDefaultHost() *Host {
	return &Host{
		Resource: &Resource{
			CreateAt: time.Now().UnixNano() / 1000000,
		},
		Describe: &Describe{},
	}
}

type Host struct {
	ResourceHash string
	DescribeHash string
	*Resource
	*Describe
}

func (h *Host) Validate() error {
	return validate.Struct(h)
}

type Vendor int

const (
	ALI_CLOUD Vendor = iota
	TX_CLOUD
	HW_CLOUD
)

type Resource struct {
	Id          string            `json:"id" validate:"required"`         // 全局唯一Id
	Vendor      Vendor            `json:"vendor" validate:"required"`     // 厂商
	Region      string            `json:"region" validate:"required"`     // 地域
	Zone        string            `json:"zone"`                           // 区域
	CreateAt    int64             `json:"create_at" validate:"required"`  // 创建时间
	ExpireAt    int64             `json:"expire_at"`                      // 过期时间
	Category    string            `json:"category"`                       // 种类
	Type        string            `json:"type"`                           // 规格
	InstanceId  string            `json:"instance_id"`                    // 实例ID
	Name        string            `json:"name"`                           // 名称
	Description string            `json:"description"`                    // 描述
	Status      string            `json:"status" validate:"required"`     // 服务商中的状态
	Tags        map[string]string `json:"tags"`                           // 标签
	UpdateAt    int64             `json:"update_at"`                      // 更新时间
	SyncAt      int64             `json:"sync_at"`                        // 同步时间
	SyncAccount string            `json:"sync_accout"`                    // 同步的账号
	PublicIP    string            `json:"public_ip"`                      // 公网IP
	PrivateIP   string            `json:"private_ip" validate:"required"` // 内网IP
	PayType     string            `json:"pay_type"`                       // 实例付费方式
}

type Describe struct {
	ResourceId              string `json:"resource_id"`                // 关联Resource
	CPU                     int    `json:"cpu" validate:"required"`    // 核数
	Memory                  int    `json:"memory"`                     // 内存
	GPUAmount               int    `json:"gpu_amount"`                 // GPU数量
	GPUSpec                 string `json:"gpu_spec"`                   // GPU类型
	OSType                  string `json:"os_type"`                    // 操作系统类型，分为Windows和Linux
	OSName                  string `json:"os_name"`                    // 操作系统名称
	SerialNumber            string `json:"serial_number"`              // 序列号
	ImageID                 string `json:"image_id"`                   // 镜像ID
	InternetMaxBandwidthOut int    `json:"internet_max_bandwidth_out"` // 公网出带宽最大值，单位为 Mbps
	InternetMaxBandwidthIn  int    `json:"internet_max_bandwidth_in"`  // 公网入带宽最大值，单位为 Mbps
	KeyPairName             string `json:"key_pair_name"`              // 秘钥对名称
	SecurityGroups          string `json:"security_groups"`            // 安全组  采用逗号分隔
}

func NewSet() *Set {
	return &Set{
		Items: []*Host{},
	}
}

type Set struct {
	Total int64
	Items []*Host
}

func (s *Set) Add(item *Host) {
	s.Items = append(s.Items, item)
}
