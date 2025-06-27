package store

import (
	"context"
	"gvadmin_v3/core/global/E"
	"os"
	"path/filepath"
)

type localClient struct {
	ctx context.Context
}

func newLocalClient() *localClient {
	return &localClient{
		ctx: context.Background(),
	}
}

func (l *localClient) UploadFile(dstFileName string, localFilePath string) (string, error) {
	backFilePath := ""
	//本地存储时, 仅转换文件显示地址(local)
	_, err := os.Stat(localFilePath)
	if err == nil || os.IsExist(err) {
		backFilePath = filepath.Join(E.ShowFilePrefix, dstFileName)
	} else {
		return "", err
	}

	//if sysos.IsWindows() {
	//	backFilePath = strings.ReplaceAll(backFilePath, "\\", "/")
	//}

	return backFilePath, nil
}

func (l *localClient) DeleteFile(dstFileName string) error {
	return nil
}
