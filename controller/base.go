package controller

import (
	"github.com/go-redis/redis/v8"
	"github.com/rapando/ussd/utils"
	"github.com/gorilla/mux"
)
type Base struct {
	RedisClient *redis.Client

	Router *mux.Router
}

func (b *Base) Init() {
	utils.InitLogger()
	b.RedisClient = utils.ConnectToRedis()
	b.InitializeRouter()
}
