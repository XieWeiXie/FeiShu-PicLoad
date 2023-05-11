package pic

import (
	"io"
	"os"
	"testing"
)

func TestAliYun(t *testing.T) {
	a := AliYunOSSUploader{
		bucket:   "xwimages",
		endPoint: "https://oss-cn-hangzhou.aliyuncs.com",
		raw:      true,
	}
	f, _ := os.Open("./wechat.jpg")
	b, _ := io.ReadAll(f)
	_, _ = a.Upload("wechat.jpg", b)
}
