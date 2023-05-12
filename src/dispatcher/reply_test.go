package dispatcher

import (
	"context"
	"fmt"
	lark "github.com/larksuite/oapi-sdk-go/v3"
	larkcontact "github.com/larksuite/oapi-sdk-go/v3/service/contact/v3"
	larkim "github.com/larksuite/oapi-sdk-go/v3/service/im/v1"
	"testing"
)

func TestReply(t *testing.T) {
	client := lark.NewClient("cli_a4d63a6006edd00e", "kCI3j3QtL25KHAlMgzgr2e1jrLvyYcKE")

	req := larkcontact.NewGetUserReqBuilder().
		UserId(`ou_eee3bf38b1e9536bd6d6453ab7f8a945`).
		UserIdType(`open_id`).
		DepartmentIdType(`open_department_id`).
		Build()
	resp, err := client.Contact.User.Get(context.Background(), req)
	fmt.Println(err, resp)

	user := resp.Data.User
	card := ReplyCardMessage("img_v2_e650d363-0b5f-4ed3-8dbe-e10df7567dbg", "https://i-x.oss-cn-hangzhou.aliyuncs.com/img_v2_e650d363-0b5f-4ed3-8dbe-e10df7567dbg.jpg", *user.OpenId, *user.Name)
	ctx := context.TODO()
	cardContent, _ := card.String()
	chatId := "oc_53c204135c6f6e1f44041699aa5aa533"
	respMessage2, err2 := client.Im.Message.Create(ctx, larkim.NewCreateMessageReqBuilder().
		ReceiveIdType(larkim.ReceiveIdTypeChatId).
		Body(larkim.NewCreateMessageReqBodyBuilder().
			MsgType(larkim.MsgTypeInteractive).
			ReceiveId(chatId).
			Content(cardContent).
			Build()).
		Build())
	fmt.Println(err2, respMessage2)
}
