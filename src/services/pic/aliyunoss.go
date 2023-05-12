package pic

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"io"

	"github.com/XieWeiXie/feishuPicLoad/src/configs"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

type AliYunOSSUploader struct {
	bucket    string
	endPoint  string
	raw       bool
	publicUrl string
}

func (a AliYunOSSUploader) Upload(f string, bs []byte) (string, error) {
	var (
		remoteUrl string
		err       error
	)
	client, _ := oss.New(a.endPoint, configs.DefaultAliYunOssConfig.AccessId, configs.DefaultAliYunOssConfig.AccessKey)
	bucketS, _ := client.Bucket(a.bucket)
	v := md5.New()
	_, _ = io.WriteString(v, f)
	nf := fmt.Sprintf("%s.jpg", f)
	switch a.raw {
	case true:
		if err = bucketS.PutObjectFromFile(nf, f); err != nil {
			return remoteUrl, err
		}
	default:
		if err = bucketS.PutObject(nf, bytes.NewReader(bs)); err != nil {
			return remoteUrl, err
		}
	}

	remoteUrl = fmt.Sprintf("%s/%s", a.publicUrl, nf)
	return remoteUrl, nil
}
