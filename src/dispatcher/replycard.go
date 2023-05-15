package dispatcher

import (
	"context"
	"fmt"
	"github.com/XieWeiXie/feishuPicLoad/src/configs"
	"github.com/larksuite/oapi-sdk-go/v3"
	"github.com/larksuite/oapi-sdk-go/v3/card"
	"github.com/larksuite/oapi-sdk-go/v3/core"
)

func ReplyCardMessageDispatcher(client *lark.Client) *larkcard.CardActionHandler {
	return larkcard.NewCardActionHandler(configs.DefaultFeiShuConfig.VerificationToken, configs.DefaultFeiShuConfig.EventEncryptKey, func(ctx context.Context, cardAction *larkcard.CardAction) (interface{}, error) {
		fmt.Println(larkcore.Prettify(cardAction))
		fmt.Println(cardAction.RequestId())
		return nil, nil
	})
}
