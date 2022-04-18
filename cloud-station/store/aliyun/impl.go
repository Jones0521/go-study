package aliyun

import (
	"cloud_station/store"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

func NewAliYunOssUpload(endpoint, ak, sk string) store.OSSUpload {
	return &impl{
		endpoint: endpoint,
		ak:       ak,
		sk:       sk,
	}
}

type impl struct {
	endpoint string
	ak       string
	sk       string
}

func (i *impl) Upload(bucketName, objectKey, filename string) (downloadUrl string, err error) {
	client, err := oss.New(i.endpoint, i.ak, i.sk)
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
