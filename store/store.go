package store

import (
	"gvadmin_v3/core/config"
	"sync"
)

var (
	soClient SoClient
	once     sync.Once
)

type SoClient interface {
	UploadFile(dstFileName string, localFilePath string) (string, error)
	DeleteFile(dstFileName string) error
}

func Instance() SoClient {
	if soClient == nil {
		once.Do(func() {
			switch config.Instance().Store.StoreType {
			case "minio":
				soClient = newMinioClient()
			case "oss":
				soClient = newOssClient()
			case "local":
				soClient = newLocalClient()
			default:
				soClient = newLocalClient()
			}
		})
	}
	return soClient
}
