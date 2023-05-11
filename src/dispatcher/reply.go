package dispatcher

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	v1 "github.com/XieWeiXie/feishuPicLoad/src/api/v1"
	"github.com/XieWeiXie/feishuPicLoad/src/configs"
	pic2 "github.com/XieWeiXie/feishuPicLoad/src/services/pic"
	"github.com/larksuite/oapi-sdk-go/v3"
	"github.com/larksuite/oapi-sdk-go/v3/core"
	"github.com/larksuite/oapi-sdk-go/v3/event/dispatcher"
	"github.com/larksuite/oapi-sdk-go/v3/service/im/v1"
	"io"
	"regexp"
)

func ReplyMessage(client *lark.Client) *dispatcher.EventDispatcher {
	return dispatcher.NewEventDispatcher(configs.DefaultFeiShuConfig.VerificationToken, configs.DefaultFeiShuConfig.EventEncryptKey).
		OnP2MessageReceiveV1(func(ctx context.Context, event *larkim.P2MessageReceiveV1) error {
			fmt.Println(larkcore.Prettify(event))

			parseContent := func(key string, content string) (string, error) {
				var contentMap map[string]interface{}
				err := json.Unmarshal([]byte(content), &contentMap)
				if err != nil {
					fmt.Println(err)
				}
				text := contentMap[key].(string)
				regex := regexp.MustCompile(`@[^ ]*`)
				return regex.ReplaceAllString(text, ""), nil
			}

			switch *event.Event.Message.MessageType {
			case larkim.MsgTypeText:
				// 文本处理
				content, _ := parseContent(larkim.MsgTypeText, *event.Event.Message.Content)
				req := larkim.NewTextMsgBuilder().Text(content).Build()
				resp, err := client.Im.Message.Create(ctx, larkim.NewCreateMessageReqBuilder().
					ReceiveIdType(larkim.ReceiveIdTypeChatId).
					Body(larkim.NewCreateMessageReqBodyBuilder().
						MsgType(larkim.MsgTypeText).
						ReceiveId(*event.Event.Message.ChatId).
						Content(req).
						Build()).
					Build())
				if err != nil {
					return err
				}
				if !resp.Success() {
					fmt.Println(resp.Code, resp.Msg, resp.RequestId())
				}
			case larkim.MsgTypePost:
			case larkim.MsgTypeImage:
				// 图片处理
				content, _ := parseContent("image_key", *event.Event.Message.Content)
				req := larkim.NewGetImageReqBuilder().ImageKey(content).Build()
				resp, err := client.Im.Image.Get(ctx, req)
				if err != nil {
					return err
				}
				if !resp.Success() {
					fmt.Println(resp.Code, resp.Msg, resp.RequestId())
					return errors.New("内部错误")
				}

				b, _ := io.ReadAll(resp.File)

				pic := pic2.Service{}
				remote, err := pic.UploadPic(ctx, &v1.UploadPicReq{
					File:     string(b),
					UserName: "",
					Channel:  "",
					ImgKey:   content,
				})
				if err != nil {
					return err
				}

				respMessage, err := client.Im.Message.Create(ctx, larkim.NewCreateMessageReqBuilder().
					ReceiveIdType(larkim.ReceiveIdTypeChatId).
					Body(larkim.NewCreateMessageReqBodyBuilder().
						MsgType(larkim.MsgTypeText).
						ReceiveId(*event.Event.Message.ChatId).
						Content(larkim.NewTextMsgBuilder().Text(remote.Img).Build()).
						Build()).
					Build())
				if err != nil {
					return err
				}
				if !respMessage.Success() {
					fmt.Println(resp.Code, resp.Msg, resp.RequestId())
					return errors.New("内部错误")
				}
			case larkim.MsgTypeFile:
			case larkim.MsgTypeAudio:
			case larkim.MsgTypeMedia:
			case larkim.MsgTypeSticker:
			case larkim.MsgTypeInteractive:
			case larkim.MsgTypeShareChat:
			case larkim.MsgTypeShareUser:

			}

			return nil
		})
}
