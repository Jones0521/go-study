package aliyun_test

import (
	"cloud_station/store/aliyun"
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var (
	ep string
	ak string
	sk string
	bn string
)

func TestUpload(t *testing.T) {
	fmt.Println(ep, ak, sk, bn)
	// 断言对象
	should := assert.New(t)

	upload, err := aliyun.NewAliYunOssUpload(ep, ak, sk)
	if should.NoError(err) {
		downloadUrl, err := upload.Upload(bn, "main.go", "main.go")
		wd, _ := os.Getwd()
		fmt.Println("work dir: ", wd)
		if should.NoError(err) {
			should.NotEmpty(downloadUrl)
		}
	}

}

// init 通过环境变量加载参数
func init() {
	ep = os.Getenv("ALI_OSS_ENDPOINT")
	ak = os.Getenv("ALI_AK")
	sk = os.Getenv("ALI_SK")
	bn = os.Getenv("ALI_BUCKET_NAME")

}
