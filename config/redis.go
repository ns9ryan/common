package config

// RedisConf 是一个 Redis 配置
type RedisConf struct {
	Host         string `json:",env=REDIS_HOST"`
	Db           int    `json:",default=0,env=REDIS_DB"`
	Mode         string `json:",optional,default=single,env=REDIS_MODE"`
	Username     string `json:",optional,env=REDIS_USERNAME"`
	Pass         string `json:",optional,env=REDIS_PASSWORD"`
	Tls          bool   `json:",optional,env=REDIS_TLS"`
	Master       string `json:",optional,env=REDIS_MASTER"`
	PoolSize     int    `json:",optional,default=48,env=REDIS_POOL_SIZE"`
	MaxIdleConns int    `json:",optional,default=12,env=REDIS_MAX_IDLE_CONNS"`
}

//func (r RedisConf) NewUniversalRedis() (redis.UniversalClient, error) {
//
//}
//
//func (r RedisConf) MustNewUniversalRedis() redis.UniversalClient {
//	rds, err := r.NewUniversalRedis()
//	logx.Must(err)
//
//	return rds
//}
