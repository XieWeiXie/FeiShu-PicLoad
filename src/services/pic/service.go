package pic

import (
	"context"
	"fmt"
	"github.com/XieWeiXie/feishuPicLoad/src/api/v1"
	"github.com/XieWeiXie/feishuPicLoad/src/configs"
	"github.com/XieWeiXie/feishuPicLoad/src/handlers"
)

type Service struct {
}

func (s Service) UploadPic(ctx context.Context, req *v1.UploadPicReq) (*v1.UploadPicReply, error) {

	var (
		res = new(v1.UploadPicReply)
		err error
	)

	var uploader handlers.Uploader
	switch req.Channel {
	case "hdrcn":
		uploader = HdRCnUploader{}
	default:
		uploader = AliYunOSSUploader{
			bucket:    configs.DefaultAliYunOssConfig.BucketName,
			endPoint:  "https://oss-cn-hangzhou.aliyuncs.com",
			publicUrl: fmt.Sprintf("https://%s.oss-cn-hangzhou.aliyuncs.com", configs.DefaultAliYunOssConfig.BucketName),
		}
	}
	response, err := uploader.Upload(req.ImgKey, []byte(req.File))
	if err != nil {
		return res, err
	}
	res.Img = response
	res.UserName = req.UserName
	res.ImgKey = req.ImgKey
	return res, nil
}
