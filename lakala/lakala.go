package lakala

import (
	"context"
	"go.uber.org/zap"
	"os"
	"sync"
	"whgo_xjd/library/lakala/common"
	"whgo_xjd/library/log"
	v8 "whgo_xjd/library/redis/v8"
	"whgo_xjd/library/setting"
)

var lakalaClientOnce sync.Once
var lakala *Lakala

func GetLakalaClient() (*Lakala, error) {
	if lakala == nil {
		log.InfoWithCtx(context.Background(), "初始lakla", zap.Any("setting", setting.LakalaSetting))
		var debug bool
		if os.Getenv("env") != "prod" {
			debug = true
		}
		lakalaClientOnce.Do(func() {
			redisCli, _ := v8.GetRedis(context.Background())
			redisInstance := common.NewRedisCache(redisCli)
			lakala = NewLakalaClient()
			lakala.Cache = redisInstance
			lakala.Debug = debug
		})
	}
	return lakala, nil
}

type Lakala struct {
	Cache  common.ICache
	Config *LakalaConfig
	Debug  bool
}

func NewLakalaClient() *Lakala {
	app := &Lakala{}
	return app
}
