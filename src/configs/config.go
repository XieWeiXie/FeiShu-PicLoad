package configs

import (
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

var (
	DefaultFeiShuConfig    FeiShuConfig
	DefaultAliYunOssConfig AliYunOssConfig
)

type FeiShuConfig struct {
	AppId             string `yaml:"appid"`
	AppSecret         string `yaml:"appsecret"`
	VerificationToken string `yaml:"verificationtoken"`
	EventEncryptKey   string `yaml:"eventencryptkey"`
}

type AliYunOssConfig struct {
	AccessId   string `yaml:"accessid"`
	AccessKey  string `yaml:"accesskey"`
	BucketName string `yaml:"bucketname"`
}

func init() {
	dir, _ := os.Getwd()
	vip := viper.New()
	vip.AddConfigPath(dir)
	vip.SetConfigName("config")
	vip.SetConfigType("yaml")
	path := filepath.Join(dir, "config.yaml")
	vip.SetConfigFile(path)
	if err := vip.ReadInConfig(); err != nil {
		panic("读取配置文件失败")
	}

	DefaultFeiShuConfig = FeiShuConfig{
		AppId:             vip.GetString("feishu.appid"),
		AppSecret:         vip.GetString("feishu.appsecret"),
		VerificationToken: vip.GetString("feishu.verificationtoken"),
		EventEncryptKey:   vip.GetString("feishu.eventencryptkey"),
	}
	DefaultAliYunOssConfig = AliYunOssConfig{
		AccessId:   vip.GetString("aliyun.accessid"),
		AccessKey:  vip.GetString("aliyun.accesskey"),
		BucketName: vip.GetString("aliyun.bucketname"),
	}
}
