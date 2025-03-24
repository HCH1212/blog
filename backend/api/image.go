package api

import (
	"context"
	"fmt"
	"github.com/HCH1212/blog/backend/service"
	"github.com/HCH1212/utils/resp"
	"github.com/cloudwego/hertz/pkg/app"
	"path"
	"time"
)

func Upload(ctx context.Context, c *app.RequestContext) {
	// 解析上传的文件
	form, err := c.MultipartForm()
	if err != nil {
		resp.Fail(c, "input error", nil)
		return
	}

	files := form.File["file"]
	if len(files) == 0 {
		resp.Fail(c, "input error", nil)
		return
	}

	var urls []string

	for _, fileHeader := range files {
		// 打开文件
		file, err := fileHeader.Open()
		if err != nil {
			resp.FailButServer(c, "server error", err.Error())
			return
		}
		defer file.Close()

		// 生成唯一文件名
		fileExt := path.Ext(fileHeader.Filename)
		newFileName := fmt.Sprintf("uploads/%d%s", time.Now().UnixNano(), fileExt)

		// 上传到七牛云
		imageURL, err := service.UploadService(file, newFileName)
		if err != nil {
			resp.FailButServer(c, "server error", err.Error())
			return
		}

		urls = append(urls, imageURL)
	}

	resp.Success(c, "upload success", urls)
}
