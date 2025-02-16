package video

import (
	"context"
	"os"
	"path/filepath"

	"go-template/service/svc"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func GenerateVideoStreamLogic(r *svc.TemplateHandler) app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		filename := ctx.Param("filename")
		if filename == "" {
			ctx.JSON(consts.StatusBadRequest, map[string]interface{}{
				"code": 400,
				"msg":  "Missing filename",
			})
			return
		}

		// 获取视频文件路径
		videoPath := filepath.Join(r.TemplateService.GetVideoStoragePath(), filename)

		// 检查文件是否存在
		if _, err := os.Stat(videoPath); os.IsNotExist(err) {
			ctx.JSON(consts.StatusNotFound, map[string]interface{}{
				"code": 404,
				"msg":  "Video not found",
			})
			return
		}

		// 设置响应头
		ctx.Header("Content-Type", "video/mp4")
		ctx.Header("Content-Disposition", "inline")

		// 发送文件
		ctx.File(videoPath)
	}
}
