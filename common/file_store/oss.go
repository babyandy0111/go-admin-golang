package file_store

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"log"
)

type ALiYunOSS struct {
	Client     interface{}
	BucketName string
}

// Setup 装载
// endpoint sss
func (e *ALiYunOSS) Setup(endpoint, accessKeyID, accessKeySecret, BucketName string, options ...ClientOption) error {
	client, err := oss.New(endpoint, accessKeyID, accessKeySecret)
	if err != nil {
		log.Println("Error:", err)
		return err
	}
	e.Client = client
	e.BucketName = BucketName

	return nil
}

// UpLoad 文件上傳
func (e *ALiYunOSS) UpLoad(yourObjectName string, localFile interface{}) error {
	// 取得存储空间。
	bucket, err := e.Client.(*oss.Client).Bucket(e.BucketName)
	if err != nil {
		log.Println("Error:", err)
		return err
	}
	// 設定分片大小為100 KB，指定分片上傳併發数為3，并開啟断点续傳上傳。
	// 其中<yourObjectName>与objectKey是同一概念，表示断点续傳上傳文件到OSS時需要指定包含文件後缀在内的完整路径，例如abc/efg/123.jpg。
	// "LocalFile"為filePath，100*1024為partSize。
	err = bucket.UploadFile(yourObjectName, localFile.(string), 100*1024, oss.Routines(3), oss.Checkpoint(true, ""))
	if err != nil {
		log.Println("Error:", err)
		return err
	}
	return nil
}

func (e *ALiYunOSS) GetTempToken() (string, error) {
	return "", nil
}
