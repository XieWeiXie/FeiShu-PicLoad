package main

import (
	"net/http"

	"github.com/XieWeiXie/feishuPicLoad/src/configs"
	"github.com/XieWeiXie/feishuPicLoad/src/dispatcher"
	"github.com/gin-gonic/gin"
	"github.com/larksuite/oapi-sdk-gin"
	"github.com/larksuite/oapi-sdk-go/v3"
)

func main() {
	s := gin.New()
	s.Use(gin.Recovery(), gin.Logger(), gin.ErrorLogger())

	v1 := s.Group("/picLoad/v1")

	heartBeatGroup := v1.Group("ping")

	heartBeatGroup.Any("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "",
			"data":    "pong",
		})
	})

	client := lark.NewClient(configs.DefaultFeiShuConfig.AppId, configs.DefaultFeiShuConfig.AppSecret)
	webhook := v1.Group("feiShu")
	webhook.POST("webhook/event", sdkginext.NewEventHandlerFunc(dispatcher.ReplyMessage(client)))

	_ = s.Run(":9091")
}
