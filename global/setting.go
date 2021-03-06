package global

import (
	"github.com/Zioyi/blog-service/pkg/logger"
	"github.com/Zioyi/blog-service/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	JWTSetting      *setting.JWTSettingS
	Logger          *logger.Logger
)
