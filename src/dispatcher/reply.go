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
	larkcard "github.com/larksuite/oapi-sdk-go/v3/card"
	"github.com/larksuite/oapi-sdk-go/v3/core"
	"github.com/larksuite/oapi-sdk-go/v3/event/dispatcher"
	larkcontact "github.com/larksuite/oapi-sdk-go/v3/service/contact/v3"
	"github.com/larksuite/oapi-sdk-go/v3/service/im/v1"
	"io"
	"regexp"
	"time"
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
				imageKey, _ := parseContent("image_key", *event.Event.Message.Content)
				resp, err := client.Im.Image.Get(ctx, larkim.NewGetImageReqBuilder().ImageKey(imageKey).Build())
				if err != nil {
					return err
				}
				if !resp.Success() {
					fmt.Println(resp.Code, resp.Msg, resp.RequestId())
					return errors.New("获取图片信息错误")
				}

				b, _ := io.ReadAll(resp.File)
				pic := pic2.Service{}
				remote, err := pic.UploadPic(ctx, &v1.UploadPicReq{
					File:     string(b),
					UserName: "",
					Channel:  "",
					ImgKey:   imageKey,
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
					fmt.Println(respMessage.Code, respMessage.Msg, respMessage.RequestId())
					return errors.New("发送图片回调错误")
				}

				req := larkcontact.NewGetUserReqBuilder().
					UserId(*event.Event.Sender.SenderId.OpenId).
					UserIdType(`open_id`).
					DepartmentIdType(`open_department_id`).
					Build()
				// 发起请求
				respUser, err := client.Contact.User.Get(context.Background(), req)
				if err != nil {
					return err
				}
				user := respUser.Data.User
				card := ReplyCardMessage(imageKey, respMessage.Msg, *event.Event.Sender.SenderId.OpenId, *user.Name)
				cardContent, _ := card.String()
				respMessage2, err2 := client.Im.Message.Create(ctx, larkim.NewCreateMessageReqBuilder().
					ReceiveIdType(larkim.ReceiveIdTypeChatId).
					Body(larkim.NewCreateMessageReqBodyBuilder().
						MsgType(larkim.MsgTypeInteractive).
						ReceiveId(*event.Event.Message.ChatId).
						Content(cardContent).
						Build()).
					Build())
				if err2 != nil {
					return err
				}
				if !respMessage2.Success() {
					fmt.Println(respMessage2.Code, respMessage2.Msg, respMessage2.RequestId())
					return errors.New("发送卡片消息错误")
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

func ReplyCardMessage(imKey string, remoteUrl string, userId string, userName string) *larkcard.MessageCard {

	cfg := larkcard.NewMessageCardConfig().UpdateMulti(false).WideScreenMode(true).EnableForward(true).Build()

	header := larkcard.NewMessageCardHeader().
		Template(larkcard.TemplateBlue).
		Title(larkcard.NewMessageCardPlainText().Content("🍿🍿🍿").Build()).
		Build()

	hr := larkcard.NewMessageCardHr().Build()

	image := larkcard.NewMessageCardImage().
		ImgKey(imKey).
		Mode(larkcard.MessageCardImageModelFitHorizontal).
		Alt(larkcard.NewMessageCardPlainText().Content("图床助手@谢伟")).
		Build()

	var layout larkcard.MessageCardActionLayout = larkcard.MessageCardActionLayoutBisected
	actions := larkcard.NewMessageCardAction().
		Actions(
			[]larkcard.MessageCardActionElement{
				larkcard.NewMessageCardEmbedButton().
					Type(larkcard.MessageCardButtonType(larkcard.MessageCardButtonTypePrimary)).
					Confirm(larkcard.NewMessageCardActionConfirm().Text(larkcard.NewMessageCardPlainText().Content(remoteUrl)).Title(larkcard.NewMessageCardPlainText().Content("🍊 远程地址"))).
					Value(map[string]interface{}{
						"key": "yes",
					}).
					Text(larkcard.NewMessageCardPlainText().Content("🍊 Copy").Build()).Build(),
			},
		).Layout(&layout).Build()

	note := larkcard.NewMessageCardNote().Elements([]larkcard.MessageCardNoteElement{
		larkcard.NewMessageCardPlainText().Content(fmt.Sprintf("🍊 上传时间 %s \n🍊 来源阿里云存储\n🍊 <at user_id=\"%s\"> %s </at>", time.Now().Format("15:04:05"), userId, userName)).Build(),
	}).Build()

	messageCard := larkcard.NewMessageCard().
		Config(cfg).
		Header(header).
		Elements([]larkcard.MessageCardElement{image, actions, hr, note}).
		Build()
	return messageCard
}
