package cmd

import (
	"cloud_station/store"
	"cloud_station/store/aliyun"
	"fmt"
	"net"
	"os"
	"path"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

const (
	// BuckName todo
	defaultBuckName = ""
	defaultEndpoint = ""
	defaultALIAK    = ""
	defaultALISK    = ""
)

var (
	buckName       string
	uploadFilePath string
	bucketEndpoint string
)

// uploadCmd represents the start command
var uploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "上传文件到中转站",
	Long:  `上传文件到中转站`,
	RunE: func(cmd *cobra.Command, args []string) error {
		p, err := getProvider()
		if err != nil {
			return err
		}
		if uploadFilePath == "" {
			return fmt.Errorf("upload file path is missing")
		}

		// 为了防止文件都堆在一个文件夹里面 无法查看
		// 我们采用日期进行编码
		day := time.Now().Format("20060102")

		// 为了防止不同用户同一时间上传相同的文件
		// 我们采用用户的主机名作为前置
		hn, err := os.Hostname()
		if err != nil {
			ipAddr := getOutBindIp()
			if ipAddr == "" {
				hn = "unknown"
			} else {
				hn = ipAddr
			}
		}

		fn := path.Base(uploadFilePath)
		ok := fmt.Sprintf("%s/%s/%s", day, hn, fn)
		_, err = p.Upload(buckName, ok, uploadFilePath)
		if err != nil {
			return err
		}
		return nil
	},
}

func getOutBindIp() string {
	conn, err := net.Dial("udp", "baidu.com:80")
	if err != nil {
		return ""
	}
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			_ = fmt.Errorf("connect close  error: %s", err)
		}
	}(conn)

	addr := strings.Split(conn.LocalAddr().String(), ":")
	if len(addr) == 0 {
		return ""
	}

	return addr[0]
}

func getProvider() (p store.OSSUpload, err error) {
	switch ossProvider {
	case "aliyun":
		fmt.Printf("上传云商: 阿里云[%s]\n", defaultEndpoint)
		if aliAccessKey == "" {
			aliAccessKey = defaultALIAK
		}
		if aliSecretKey == "" {
			aliSecretKey = defaultALISK
		}
		fmt.Printf("上传用户: %s\n", aliAccessKey)
		p, err = aliyun.NewAliYunOssUpload(bucketEndpoint, aliAccessKey, aliSecretKey)
		return
	case "qcloud":
		return nil, fmt.Errorf("not impl")
	default:
		return nil, fmt.Errorf("unknown oss privier options [aliyun/qcloud]")
	}
}

func init() {
	uploadCmd.PersistentFlags().StringVarP(&uploadFilePath, "file_path", "f", "", "upload file path")
	uploadCmd.PersistentFlags().StringVarP(&buckName, "bucket_name", "b", defaultBuckName, "upload oss bucket name")
	uploadCmd.PersistentFlags().StringVarP(&bucketEndpoint, "bucket_endpoint", "e", defaultEndpoint, "upload oss endpoint")
	RootCmd.AddCommand(uploadCmd)
}
