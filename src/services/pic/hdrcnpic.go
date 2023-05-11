package pic

import (
	"bytes"
	"io"
	"net/http"
)

const (
	hdrcnRootUrl = "https://upload.hd-r.cn/"
)

type HdRCnUploader struct {
}

func (h HdRCnUploader) Upload(f string, content []byte) (string, error) {
	var (
		remoteUrl string
		err       error
	)
	payload := &bytes.Buffer{}
	payload.Write(content)
	req, err := http.NewRequest(http.MethodPost, hdrcnRootUrl, payload)
	if err != nil {
		return remoteUrl, err
	}
	var (
		res string
		try int = 3
	)
	for try > 0 {
		client := http.DefaultClient
		response, err := client.Do(req)
		if err != nil {
			try -= 1
			continue
		}
		defer response.Body.Close()
		body, err := io.ReadAll(response.Body)
		if len(body) == 0 {
			try -= 1
			continue
		}
		res = string(body)
		break
	}
	return res, nil
}
