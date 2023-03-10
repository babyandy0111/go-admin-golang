package file_store

import (
	"fmt"
	"github.com/huaweicloud/huaweicloud-sdk-go-obs/obs"
	"log"
)

type HuaWeiOBS struct {
	Client     interface{}
	BucketName string
}

func (e *HuaWeiOBS) Setup(endpoint, accessKeyID, accessKeySecret, BucketName string, options ...ClientOption) error {
	// 建立ObsClient結構体
	client, err := obs.New(accessKeyID, accessKeySecret, endpoint)
	if err != nil {
		log.Println("Error:", err)
		return err
	}
	e.Client = client
	e.BucketName = BucketName
	return nil
}

// UpLoad 文件上傳
// yourObjectName 文件路径名稱，与objectKey是同一概念，表示断点续傳上傳文件到OSS時需要指定包含文件後缀在内的完整路径，例如abc/efg/123.jpg
func (e *HuaWeiOBS) UpLoad(yourObjectName string, localFile interface{}) error {
	// 取得存储空间。
	input := &obs.PutFileInput{}
	input.Bucket = e.BucketName
	input.Key = yourObjectName
	input.SourceFile = localFile.(string)
	output, err := e.Client.(*obs.ObsClient).PutFile(input)

	if err == nil {
		fmt.Printf("RequestId:%s\n", output.RequestId)
		fmt.Printf("ETag:%s, StorageClass:%s\n", output.ETag, output.StorageClass)
	} else {
		if obsError, ok := err.(obs.ObsError); ok {
			fmt.Println(obsError.Code)
			fmt.Println(obsError.Message)
		} else {
			fmt.Println(err)
		}
	}
	return nil
}

func (e *HuaWeiOBS) GetTempToken() (string, error) {
	return "", nil
}
