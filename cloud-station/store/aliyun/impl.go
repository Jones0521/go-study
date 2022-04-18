package aliyun

import (
	"cloud_station/store"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	validator "github.com/go-playground/validator/v10"
)

var (
	validate = validator.New()
)

func NewAliYunOssUpload(endpoint, ak, sk string) (store.OSSUpload, error) {
	uploader := &impl{
		Endpoint: endpoint,
		Ak:       ak,
		Sk:       sk,
	}
	if err := uploader.Validate(); err != nil {
		return nil, fmt.Errorf("vaildate params error: %s", err)
	}
	return uploader, nil
}

type impl struct {
	Endpoint string `validate:"required"`
	Ak       string `validate:"required"`
	Sk       string `validate:"required"`
}

func (i *impl) Validate() error {
	return validate.Struct(i)
}

func (i *impl) Upload(bucketName, objectKey, filename string) (downloadUrl string, err error) {
	client, err := oss.New(i.Endpoint, i.Ak, i.Sk)
	if err != nil {
		err = fmt.Errorf("new client error: %s", err)
		return
	}
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		err = fmt.Errorf("get bucker %s error error: %s", bucketName, err)
		return
	}
	err = bucket.PutObjectFromFile(objectKey, filename)
	if err != nil {
		err = fmt.Errorf("upload file %s error error: %s", filename, err)
		return
	}
	// 生成下载链接
	return bucket.SignURL(objectKey, oss.HTTPGet, 60*60*24*3)

}
