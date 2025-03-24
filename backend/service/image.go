package service

import (
	"context"
	"fmt"
	"github.com/HCH1212/blog/backend/conf"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"mime/multipart"
)

// UploadImage 上传文件到七牛云
func UploadService(file multipart.File, fileName string) (string, error) {
	// 生成上传凭证
	mac := qbox.NewMac(conf.GetConf().QiNiu.AccessKey, conf.GetConf().QiNiu.SecretKey)
	putPolicy := storage.PutPolicy{
		Scope: conf.GetConf().QiNiu.Bucket,
	}
	upToken := putPolicy.UploadToken(mac)

	cfg := storage.Config{
		Zone:          &storage.Zone_as0, // 根据 bucket 区域修改
		UseHTTPS:      true,
		UseCdnDomains: true,
	}

	// 初始化表单上传
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}

	// 读取文件内容
	data := make([]byte, 1024*1024) // 1MB 缓冲区
	n, err := file.Read(data)
	if err != nil {
		return "", fmt.Errorf("file read error: %v", err)
	}

	// 上传文件
	err = formUploader.Put(context.Background(), &ret, upToken, fileName, file, int64(n), nil)
	if err != nil {
		return "", fmt.Errorf("upload error: %v", err)
	}

	// 拼接 URL 并返回
	imageURL := fmt.Sprintf("%s/%s", conf.GetConf().QiNiu.Domain, ret.Key)
	return imageURL, nil
}
