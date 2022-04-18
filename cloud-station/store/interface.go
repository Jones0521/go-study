package store

// OSSUpload 存储适配器
// 阿里云oss/腾讯云cos/aws S3
type OSSUpload interface {
	Upload(bucket, objectKey, filename string) (downloadUrl string, err error)
}
