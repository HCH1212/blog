package service

import (
	"blog/global"
	"blog/utils"
	"errors"
	"github.com/gin-gonic/gin"
	"path"
	"strings"
)

type Response struct {
	FileName  string `json:"file_name"`  // 文件名
	IsSuccess bool   `json:"is_success"` // 是否上传成功
	Msg       string `json:"msg"`        // 返回消息
}

// 上传图片，返回图片的path
func ImageUpdateService(c *gin.Context) (resp []Response, err error) {
	//fileHeader, err := c.FormFile("image")
	//if err != nil {
	//	return "", err
	//}
	form, err := c.MultipartForm()
	if err != nil {
		global.Log.Error(err.Error())
		return nil, errors.New("service error")
	}
	fileList, ok := form.File["image"]
	if !ok {
		return nil, errors.New("no images")
	}

	var resList []Response

	for _, file := range fileList {
		// 判断文件后缀是否在白名单
		if !utils.InList(strings.ToLower(path.Ext(file.Filename)), global.Config.Upload.Suffix) {
			resList = append(resList, Response{
				FileName:  file.Filename,
				IsSuccess: false,
				Msg:       "file suffix is error",
			})
			continue
		}

		// 设置文件存储路径
		filePath := path.Join(global.Config.Upload.Path, file.Filename) // Path不存在的话会自动创建

		// 判断文件大小是否合规
		size := float64(file.Size) / float64(1024*1024)
		if size > global.Config.Upload.Size {
			resList = append(resList, Response{
				FileName:  file.Filename,
				IsSuccess: false,
				Msg:       "such big size",
			})
			continue
		}

		// 保存文件
		err = c.SaveUploadedFile(file, filePath)
		if err != nil {
			global.Log.Error("save file failed:", err)
			resList = append(resList, Response{
				FileName:  file.Filename,
				IsSuccess: false,
				Msg:       "save file failed",
			})
			continue
		}
		resList = append(resList, Response{
			FileName:  filePath,
			IsSuccess: true,
			Msg:       "update success",
		})
	}
	return resList, nil
}
