package file_store

// DriverType 驱动類型
type DriverType string

const (
	// HuaweiOBS 华为云OBS
	HuaweiOBS DriverType = "HuaweiOBS"
	// AliYunOSS 阿里云OSS
	AliYunOSS DriverType = "AliYunOSS"
	// QiNiuKodo 七牛云kodo
	QiNiuKodo DriverType = "QiNiuKodo"
)

type ClientOption map[string]interface{}

// TODO: FileStoreType名稱待定

// FileStoreType OXS
type FileStoreType interface {
	// Setup 装载 endpoint sss
	Setup(endpoint, accessKeyID, accessKeySecret, BucketName string, options ...ClientOption) error
	// UpLoad 上傳
	UpLoad(yourObjectName string, localFile interface{}) error
	// GetTempToken 取得临時Token
	GetTempToken() (string, error)
}
